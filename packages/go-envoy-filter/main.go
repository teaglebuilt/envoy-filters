package main

import (
	"github.com/buger/jsonparser"
	"github.com/gomodule/redigo/redis"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

var (
	authHeader string
	connection *redis.Conn
)

const clusterName = "gateway_authz"

type vmContext struct {
	types.DefaultVMContext
}

func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	types.DefaultPluginContext
}

func (*pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpContext{}
}

type httpContext struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
}

func (ctx *httpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogInfo("OnHttpRequestHeaders")
	return types.ActionContinue
}

func (ctx *httpContext) OnHttpResponseHeaders(int, bool) types.Action {
	proxywasm.AddHttpResponseHeader("x-wasm-filter", "youve been filtered")
	proxywasm.AddHttpResponseHeader("x-auth", authHeader)
	return types.ActionContinue
}

func (ctx *pluginContext) OnHttpRequestBody(int, bool) types.Action {
	body, err := proxywasm.GetHttpRequestBody(0, 4096)
	if err != nil {
		proxywasm.LogCriticalf("failed to parse body: %v", err)
	}
	proxywasm.LogDebugf("body: %v", string(body))

	hs := [][2]string{
		{":authority", "authz:5004"},
		{":path", "/oauth2/login"},
		{":method", "POST"},
		{"accept", "*/*"},
		{"user-agent", "proxy-wasm"},
	}

	// authz, err := proxywasm.GetHttpRequestHeader("authorization")
	if err != nil {
		proxywasm.LogCriticalf("failed to get authz from headers: %v", err)
		return types.ActionContinue
	}
	if _, err := proxywasm.DispatchHttpCall(
		clusterName, hs, body, [][2]string{}, 5000, httpCallResponseCallback,
	); err != nil {
		proxywasm.LogCriticalf("dispatch httpcall failed: %v", err)
		return types.ActionContinue
	}
	proxywasm.ResumeHttpRequest()
	return types.ActionPause
}

// func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {

// 	return types.OnPluginStartStatusOK
// }

func (ctx *httpContext) OnHttpStreamDone() {
	proxywasm.LogInfof("finished")
}

func httpCallResponseCallback(numHeaders, bodySize, numTrailers int) {

	b, err := proxywasm.GetHttpCallResponseBody(0, bodySize)
	if err != nil {
		proxywasm.LogCriticalf("failed to get response body: %v", err)
		proxywasm.ResumeHttpRequest()
		return
	}

	allowed, err := jsonparser.GetBoolean(b, "result", "allow")
	if err != nil {
		proxywasm.LogCriticalf("failed to parse allowed filed from response body: %v", err)
		proxywasm.ResumeHttpRequest()
		return
	}

	proxywasm.LogInfof("allowed: %v", allowed)

	proxywasm.ResumeHttpRequest()
}

func main() {
	proxywasm.SetVMContext(&vmContext{})
}
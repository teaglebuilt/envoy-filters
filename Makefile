-include filters/go-envoy-filter/Makefile
-include filters/rust-envoy-filter/Makefile

up:
	docker compose up -d --build

build_sandbox:
	docker build -t envoy_sandbox -f Dockerfile.envoy .

run_sandbox: build_sandbox
	docker run -e ENVOY_FILTERS="router,go-envoy-filter" \
		-e WASM_FILTER_NAME="go-envoy-filter" \
		-p 8080:8080 envoy_sandbox
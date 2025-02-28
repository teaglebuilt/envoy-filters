#!/bin/sh

set -e

echo "Starting Envoy with the following filters: $ENVOY_FILTERS"

export LISTENER_PORT=${LISTENER_PORT:-8080}
export UPSTREAM_CLUSTER=${UPSTREAM_CLUSTER:-"backend_service"}

export WASM_RUNTIME=${WASM_RUNTIME:-"envoy.wasm.runtime.v8"}
export WASM_FILTER_NAME=${WASM_FILTER_NAME:-"go-envoy-filter"}
export WASM_FILE=${WASM_FILE:-"go-envoy-filter.wasm"}

echo "Checking available filters in /etc/envoy/filters/"
ls -l /etc/envoy/filters/

SELECTED_FILTERS=${ENVOY_FILTERS:-"router"}  # Default to built-in router filter

FILTER_CONFIGS=""
for FILTER in $(echo $SELECTED_FILTERS | tr "," "\n"); do
  if [ "$FILTER" = "wasm" ]; then
    WASM_CONFIG=$(cat /etc/envoy/filters/wasm-filter.yaml.tpl | \
                  sed "s|{{ .WASM_RUNTIME }}|$WASM_RUNTIME|g" | \
                  sed "s|{{ .WASM_FILTER_NAME }}|$WASM_FILTER_NAME|g" | \
                  sed "s|{{ .WASM_FILE }}|$WASM_FILE|g")
    FILTER_CONFIGS="${FILTER_CONFIGS}\n$WASM_CONFIG"
  else
    FILTER_FILE="/etc/envoy/filters/${FILTER}.yaml"
    if [ -f "$FILTER_FILE" ]; then
      echo "✅ Adding filter: $FILTER"
      FILTER_CONFIGS="${FILTER_CONFIGS}$(cat $FILTER_FILE)\n"
    else
      echo "⚠️ Warning: Filter config not found for $FILTER"
    fi
  fi
done

echo "Generating Envoy configuration..."
envsubst < /etc/envoy/envoy.yaml.tpl > /etc/envoy/envoy.yaml
cat /etc/envoy/envoy.yaml

echo "🚀 Starting Envoy..."
exec envoy -c /etc/envoy/envoy.yaml --log-level debug

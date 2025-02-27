---
- name: envoy.filters.http.wasm
  typed_config:
    "@type": type.googleapis.com/envoy.extensions.filters.http.wasm.v3.Wasm
    config:
      name: {{ .WASM_FILTER_NAME }}
      root_id: wasm_filter_root
      vm_config:
        runtime: "{{ .WASM_RUNTIME }}"
        code:
          local:
            filename: "/lib/{{ .WASM_FILE }}"

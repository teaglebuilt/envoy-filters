-include filters/go-envoy-filter/Makefile
-include filters/rust-envoy-filter/Makefile

up:
	docker compose up -d --build

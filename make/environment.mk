.PHONY: setup-devenv
setup-devenv: $(PROTOC) $(GRPC_GATEWAY) $(MODVENDOR) $(PROTOC_SWAGGER_GEN) deps-vendor modvendor

.PHONY: setup-cienv
setup-cienv: deps-vendor modvendor

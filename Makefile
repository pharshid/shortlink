# Contracts
PROTO_DIR=./pb

# API Service
API_PROTO=$(PROTO_DIR)/api.proto
API_OUT=./svc/api/api
API_SHORTCUT_OUT=./svc/api/shortcut

# Shortcut Service
SHORTCUT_PROTO=$(PROTO_DIR)/shortcut.proto
SHORTCUT_OUT=./svc/shortcut/shortcut
SHORTCUT_API_OUT=./svc/shortcut/api

# Go protoc options
GO_OPT=import
GRPC_OPT=source_relative

up: codegen
	@echo "Starting Services ..."
	docker-compose up -d --build
	@echo "Done!"

up_nogen:
	@echo "Starting Services (no codegen) ..."
	docker-compose up -d --build
	@echo "Done!"

up_nobuild:
	@echo "Starting Services (no build / no codegen)  ..."
	docker-compose up -d
	@echo "Done!"

down:
	@echo "Stopping Services ..."
	docker-compose down
	@echo "Done!"

codegen: codegen_api codegen_shortcut
	@echo "Generated All Services ..."

# Generate Code for API Service
codegen_api:
	@echo "Generating Code for API Service ..."
	# APIService 
	protoc -I=$(PROTO_DIR) --go_out=$(API_OUT) --go_opt=paths=source_relative \
    --go-grpc_out=$(API_OUT) --go-grpc_opt=paths=source_relative \
    $(API_PROTO)

	# ShortcutService (Client)
	protoc -I=$(PROTO_DIR) --go_out=$(API_SHORTCUT_OUT) --go_opt=paths=source_relative \
    --go-grpc_out=$(API_SHORTCUT_OUT) --go-grpc_opt=paths=source_relative \
    $(SHORTCUT_PROTO)

# Generate Code for Shortcut Service	
codegen_shortcut:
	@echo "Generating Code for Shortcut Service ..."
	# ShortcutService 
	protoc -I=$(PROTO_DIR) --go_out=$(SHORTCUT_OUT) --go_opt=paths=source_relative \
    --go-grpc_out=$(SHORTCUT_OUT) --go-grpc_opt=paths=source_relative \
    $(SHORTCUT_PROTO)

	# ApiService (.proto import dependency)
	protoc -I=$(PROTO_DIR) --go_out=$(SHORTCUT_API_OUT) --go_opt=paths=source_relative \
    --go-grpc_out=$(SHORTCUT_API_OUT) --go-grpc_opt=paths=source_relative \
    $(API_PROTO)
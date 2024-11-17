# Change this to your own Go module
GO_MODULE := my-protobuf

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	@echo "Cleaning on Windows..."
	# Cek apakah ada folder 'protogen' dan hapus di Windows menggunakan cmd
	@rd /s /q protogen 2> nul || echo "No protogen directory to clean"
else
	@echo "Cleaning on Unix-like OS..."
	# Di Unix-like, menggunakan perintah bash standar
	rm -rf ./protogen || echo "No protogen directory to clean"
endif

.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
		./proto/basic/*.proto

.PHONY: build
build: clean protoc tidy

.PHONY: run
run:
	go run main.go

.PHONY: execute
execute: build run

.PHONY: protoc-validate
protoc-validate:
	protoc --validate_out="lang=go:./generated" --go_opt=module=${GO_MODULE} --go_out=. ./proto/car/*.proto

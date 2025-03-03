.PHONY: lint install-lint clean

# Linter 版本
GOLANGCI_LINT_VERSION ?= v1.61.0

# Linter 执行路径
GOLANGCI_LINT := $(shell which golangci-lint || echo "")

# 检查是否安装了 golangci-lint
ifeq ($(GOLANGCI_LINT),)
GOLANGCI_LINT = $(GOBIN)/golangci-lint
endif

# 安装 golangci-lint
install-lint:
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

# 运行 lint 检查
lint: install-lint
	@echo "Running golangci-lint..."
	@$(GOLANGCI_LINT) run ./...

# 清理临时文件
clean:
	@rm -rf ./tmp


.PHONY: goimports install-goimports format clean

# goimports 执行路径
GOIMPORTS := $(shell which goimports || echo "")

# 检查是否安装了 goimports
ifeq ($(GOIMPORTS),)
GOIMPORTS = $(GOBIN)/goimports
endif

# 安装 goimports
install-goimports:
	@echo "Installing goimports..."
	@go install golang.org/x/tools/cmd/goimports@latest

# 运行 goimports 格式化代码
format: install-goimports
	@echo "Running goimports..."
	@$(GOIMPORTS) -w .

# 清理临时文件
clean:
	@rm -rf ./tmp

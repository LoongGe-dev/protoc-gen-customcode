.PHONY: build install example clean help

# 构建插件
build:
	@echo "构建插件..."
	go build -o bin/protoc-gen-customcode ./cmd/protoc-gen-customcode
	@echo "构建完成: bin/protoc-gen-customcode"

# 安装插件到 GOPATH/bin
install:
	@echo "安装插件..."
	go install ./cmd/protoc-gen-customcode
	@echo "安装完成"

# 生成示例代码
example: install
	@echo "\n=== 生成示例代码 ==="
	@mkdir -p examples/proto
	protoc --proto_path=examples/proto \
		--go_out=examples/proto \
		--go_opt=paths=source_relative \
		--customcode_out=examples/proto \
		--customcode_opt=paths=source_relative \
		examples/proto/echo.proto
	@echo "\n生成的错误码文件:"
	@ls -la examples/proto/*_customcode.go
	@echo "\n=== 运行示例 ==="
	go run examples/main.go

# 清理生成的文件
clean:
	@echo "清理生成的文件..."
	rm -f examples/proto/*.pb.go
	rm -f examples/proto/*_customcode.go
	rm -f bin/protoc-gen-customcode
	@echo "清理完成"

# 测试
test: example
	go test ./...

# 创建发布包
dist: clean
	@echo "创建发布包..."
	@mkdir -p dist
	tar -czf dist/protoc-gen-customcode.tar.gz --exclude=".git" --exclude="dist" .
	@echo "发布包已创建: dist/protoc-gen-customcode.tar.gz"

# 显示帮助
help:
	@echo "可用的命令:"
	@echo "  make build   - 构建插件"
	@echo "  make install - 安装插件到 GOPATH/bin"
	@echo "  make example - 运行示例"
	@echo "  make clean   - 清理生成的文件"
	@echo "  make dist    - 创建发布包"
	@echo "  make help    - 显示此帮助"
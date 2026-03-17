#!/bin/bash
# 测试插件脚本

# 设置插件路径
export PATH="./bin:$PATH"

# 清理旧文件
rm -f examples/proto/*_customcode.go

# 运行 protoc
protoc --proto_path=examples/proto \
    --go_out=examples/proto \
    --go_opt=paths=source_relative \
    --customcode_out=examples/proto \
    --customcode_opt=paths=source_relative \
    examples/proto/echo.proto

# 显示生成的文件
echo "生成的文件:"
ls -la examples/proto/*_customcode.go 2>/dev/null || echo "没有生成 customcode 文件"

# 显示文件内容
if [ -f "examples/proto/echo_customcode.go" ]; then
    echo "文件内容:"
    cat examples/proto/echo_customcode.go
fi

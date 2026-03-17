# protoc-gen-customcode

Protocol Buffers 自定义错误码生成器

## 功能特性

- 根据 proto 文件中的 //start 注释生成自定义错误码
- 错误码格式：`start<<16 + proto_value`
- 生成 Go 代码，包含错误码常量和映射函数
- 支持枚举值名称的驼峰转换

## 安装

### 前提条件

- Go 1.19+
- protoc (Protocol Buffers编译器)
- protoc-gen-go

### 安装步骤

1. **下载项目**
```bash
git clone <repository-url>
cd protoc-gen-customcode
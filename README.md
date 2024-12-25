# IT666 多平台统一管理系统

## 项目介绍
IT666 是一个基于 Golang 开发的后端多平台统一管理项目模板，采用现代化的技术栈，提供完整的 CRUD 基础功能实现。

## 技术栈
- 编程语言: Go 1.20.x
- Web 框架: Gin
- ORM 框架: Gorm
- 缓存数据库: Redis
- 日志框架: ZAP
- 配置管理: Viper
- 配置格式: YAML/JSON

## 项目特性
- 遵循 RESTful API 设计规范
- 基于 Gofmt 的标准化代码格式
- 集成常用中间件
- 统一的错误处理
- 完整的数据库 CRUD 示例
- 规范的路由分组管理
- 集成参数验证器
- AI 辅助开发（50%注释和30%代码由 AI 生成）

## 项目结构
```
├── config/         # 配置文件目录
├── controllers/    # 控制器层
├── models/         # 数据模型层
├── routes/         # 路由配置
├── services/       # 业务逻辑层
├── utils/          # 工具函数
└── main.go         # 程序入口
```

## 快速开始

### 环境要求
- Go 1.20.x
- MySQL
- Redis

### 安装步骤
1. 克隆项目
```bash
git clone https://github.com/yourusername/it666-server.git
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置文件
复制 `config.example.yaml` 为 `config.yaml` 并修改相关配置

4. 运行项目
```bash
go run main.go
```

## 开发指南
1. 本项目已完成基础 CRUD 框架搭建
2. 参考已有示例，可快速实现新的业务模块
3. 遵循项目既定的代码规范和架构设计

## 贡献指南
欢迎提交 Issue 和 Pull Request

## 许可证
MIT License

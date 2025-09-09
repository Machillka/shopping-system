# 练手项目

实现订单系统（老掉牙的东西）
尝试采用 Clean Architecture 分层设计，支持 HTTP/gRPC 接口、异步消息、可扩展的领域模型。

## 实现内容

1. 商品数据
   1. 名称
   2. 价格
   3. 折扣
2. 仓库
   1. 管理商品数据的内容
   2. 模拟进货、出库
3. 订单管理
   1. 下单
   2. 支付
   3. 取消支付
   4. 订单查询

## 目录结构

- cmd/order: 服务启动入口
- internal/domain: 领域模型与接口
- internal/application: 用例层
- internal/adapters: HTTP/GRPC/MQ 适配层
- internal/infra: 基础设施实现
- pkg: 公共工具
- configs: 配置文件
- scripts: 辅助脚本

├── cmd/
│   └── order/              # 可执行启动文件
├── internal/
│   ├── domain/             # 领域实体、接口
│   ├── application/        # 用例交互器
│   ├── adapters/
│   │   ├── http/           # Rest API 实现
│   │   ├── grpc/           # gRPC 服务实现
│   │   └── mq/             # MQ 消费者/生产者
│   └── infra/              # 仓储、外部系统、配置、日志
├── pkg/                    # 公共工具库（错误包装、日志等）
├── configs/                # YAML/JSON 配置文件
├── scripts/                # 数据库迁移、构建脚本
├── go.mod
└── Makefile

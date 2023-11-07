# micro-shop-services
## 概述
这是一个实现商城基本功能的底层业务处理逻辑的服务模块集，使用 gRPC 来为[其上层Web层](https://github.com/Yifangmo/micro-shop-web)提供服务。该层服务模块间的代码相互解耦，但可能会需要在服务上相互调用。系统具体架构如下：
![商城后端系统架构.png](https://s2.loli.net/2023/11/07/FvBnQXA8Ob2DzWk.png)

## 功能模块
| 子目录 | 服务 | 描述 |
| ---  | ---- | ---- |
| user | 用户服务 | 实现用户个人信息、商品收藏、收货人信息、留言的增删改查服务 |
| goods | 商品服务 | 实现商品及其分类、品牌及其分类、轮播图的增删改查服务 |
| inventory | 库存服务 | 实现库存的设置、扣减和归还等库存操作服务 |
| order  | 订单服务 | 实现对订单的新增和查询、购物车增删改查操作服务 |

## 目录结构
    |-- go.mod
    |-- gen.go 存放 go generate 所执行的命令
    |-- common 存放共用的proto文件、服务注册逻辑等
    |-- user   对应一个服务模块
        |-- main.go
        |-- run.sh   简单的运行命令
        |-- configs  存放配置文件、Nacos和服务的配置结构体
        |-- global   存放全局使用的变量，如服务器配置、需要调用其他服务的 gRPC 客户端、是否为debug等
        |-- services 存放服务 rpc 接口的实现方法、或 mq 的消费逻辑
        |-- initialize 存放全局变量的初始化或功能注册逻辑
        |-- proto 存放对应本服务所定义的 rpc 接口和消息结构的 proto 文件，以及由 protoc 编译生成的源码文件，或者包含所需的本层其他服务的 proto 文件及生成文件
        |-- models 存放数据库表对应模型及其钩子方法、模型转换方法等
        |-- utils 存放通用的工具函数，如错误处理或转换、服务注册
        |-- scripts 存放一次运行的脚本，如数据库迁移
    |-- ... 其他服务模块(同上)

## 服务启动
1. 在 `${service}/configs/dev.yaml` 配置好 Nacos 服务器信息，并在 Nacos 中配置好 `ServerConfig`
2. 在 `go.mod` 同级目录运行 `go generate`
3. 进入各服务的子目录，如 `cd user/`，运行 `./run.sh`

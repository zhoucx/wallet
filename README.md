# wallet
    wallet server
# 包说明
## cmd
    程序命令入口
## internal
### app
    构建http路由，启动http服务
### model
    钱包存储的数据结构和逻
    * Wallet 用户的钱包信息
    * TransferLog 用户的交易流水
## pkg
    一些公共函数
    * errors 定义钱包返回的错误码
    * log 记录程序日志
    * rand_str 生成随机码
    * time 时间的一些封装
    * uuid 内部生成唯一ID的封装，为了能直接运行，没有使用第三方包
## test
    测试文件存放的地方
    * wallet_test.go 钱包的测试用例
    * benchmark_test.go 钱包的压力测试用例

# 程序执行
    go run ./cmd/server , 默认监听端口 9091

# 其他说明
    * 如果要接入GPRC，grpc服务可以直接调用model的钱包接口，实现逻辑共享
    * 水平扩展问题，在业务服务前面加一个网关进行分发即可，钱包服务在数据库操作的时候对用户的钱包进行加锁即可（加锁的时候，两个钱包ID排一下序，防止死锁）
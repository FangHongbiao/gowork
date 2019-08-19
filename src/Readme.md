### golang

### 框架
1. 直接执行go run main.go 提示未定义文件, 因为其他文件的编译也需要手动指定 `go run handler.go server.go  main.go`
2. error running client: dial tcp 127.0.0.1:9090: connectex: No connection could be made because the target machine actively refused it.
1. thrift
2. govendor
   1. 目录结构
   2. 本地包import不到: `govendor add +e ; govendor add +l`
   3. 改动了本地的文件需要更新: `govendor update rpc/userservice`
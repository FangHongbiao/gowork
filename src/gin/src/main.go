package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc/userservice"
	"github.com/apache/thrift/lib/go/thrift"
	"context"
)


var protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

var transportFactory = thrift.NewTTransportFactory()

// var transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

var secure = false

var defaultCtx = context.Background()

var addr = "localhost:8080"
var userService, rpcErr = userservice.GetClient(transportFactory, protocolFactory, addr, secure)

func main() {
	
	if rpcErr != nil {
		fmt.Println("error running rpc service:", rpcErr)
	}
	fmt.Println(*userService)
	gin.SetMode(gin.DebugMode)

	//获得路由实例
	router := gin.Default()    

	//添加中间件
	router.Use(Middleware)

	//注册接口
	router.GET("/mall/user/:userId", UserInfoHandler)
	router.POST("/mall/stock/:goodsId", StockInfoHandler)
	router.PUT("/mall/goods/:userId/:stockId", BuyGoodsHandler)

	//监听端口
	http.ListenAndServe(":8005", router)
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

func UserInfoHandler(c *gin.Context) {
	userId := c.Param("userId")
	fmt.Println(userId)
	userInfo, err := userService.GetUserInfo(defaultCtx, 200)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Username: ",userInfo.Username)
	fmt.Println("Credits: ", userInfo.Credits)
	fmt.Println("Level: ", userInfo.Level)
	fmt.Println("Sex: ", userInfo.Sex)
	fmt.Println("Phone: ", userInfo.Phone)
	fmt.Println("Address: ", userInfo.Address)

	fmt.Println(err)
	c.JSON(http.StatusOK, userInfo)
	return
}
func StockInfoHandler(c *gin.Context) {
	stockId := c.Param("stockId")
	fmt.Println(stockId)
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, "")
	return
}

func BuyGoodsHandler(c *gin.Context) {

	userId := c.Param("userId")
	stockId := c.Param("stockId")
	fmt.Println(userId)
	fmt.Println(stockId)

	fmt.Println("------商品模块中购买商品接口-----")
	fmt.Println("调用会员服务")
	fmt.Println("调用库存服务")
	fmt.Println("调用订单服务")
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}
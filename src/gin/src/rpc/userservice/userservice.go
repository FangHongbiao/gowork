package userservice

import (
	"context"
	"crypto/tls"
	"fmt"
	"mall/user"

	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

// func handleClient(client *user.UserManagerClient) (err error) {
// 	client.Ping(defaultCtx)
// 	fmt.Println("ping()")

// 	userInfo, err := client.GetUserInfo(defaultCtx, 200)
	
// 	if err != nil {
// 		fmt.Print(err)
// 		return err
// 	}

// 	fmt.Println("Username: ",userInfo.Username)
// 	fmt.Println("Credits: ", userInfo.Credits)
// 	fmt.Println("Level: ", userInfo.Level)
// 	fmt.Println("Sex: ", userInfo.Sex)
// 	fmt.Println("Phone: ", userInfo.Phone)
// 	fmt.Println("Address: ", userInfo.Address)
// 	return nil
// }

func GetClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) (client *user.UserManagerClient, e error) {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil, err
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return nil, err
	}
	// defer transport.Close()
	if err := transport.Open(); err != nil {
		return nil, err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return user.NewUserManagerClient(thrift.NewTStandardClient(iprot, oprot)), nil
}

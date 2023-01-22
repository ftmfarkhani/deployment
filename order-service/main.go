package main

import (
	"flag"
	//"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	authAddr := flag.String("authAddr", "localhost:50051", "Authentication service address (GRPC)")
	prodAddr := flag.String("prodAddr", "localhost:50052", "Product service address(GRPC)")

	flag.Parse()

	svc := NewInmemservice(*prodAddr)
	handler := NewHandler(svc)
	authMW := NewAuthMiddlaware(*authAddr)

	router := gin.Default()

	router.GET("/order-service", authMW.hassAcces, handler.getBasket)
	router.POST("/order-service/:pid", authMW.hassAcces, handler.addProductTobasket)
	router.PUT("/order-service/:pid", authMW.hassAcces, handler.modifyBasket)

	router.Run(":3000")
}

package main

import (
	"fmt"

	"github.com/heyubani/go-template/connections"
	"github.com/heyubani/go-template/database"
	"github.com/heyubani/go-template/server"
)

func main() {
	fmt.Println("About Starting Cloudsania logging Service")

	mongoDb := database.ConnectMongo()
	defer mongoDb()
	fmt.Println("DB Connected")

	redisConn := connections.GetRedisConnection()
	defer redisConn.Close()
	fmt.Println("Redis Connected")

	fmt.Println("About Starting HttpServer Service")
	server.InitHttpServer()

}

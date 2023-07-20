package main


import (
	"log"
	"github.com/turamant/ModernGo/config"
	"github.com/turamant/ModernGo/server"	
	_ "github.com/lib/pq"
)
func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing configuration")
	config := config.InitConfig("runners")
	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
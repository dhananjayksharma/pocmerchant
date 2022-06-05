package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dkgosql.com/pacenow-service/pkg/adapter/mysql"
	"dkgosql.com/pacenow-service/pkg/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

const (
	serviceName = "dkgosql.com/pacenow-service"
)

//go:embed config.yml
var config embed.FS

func startService() {
	// Set the file name of the configurations file
	if os.Getenv("MICROSERVICECAUTHNEWAPI") == "local" {
		viper.SetConfigName("config-local")
	} else if os.Getenv("MICROSERVICECAUTHNEWAPI") == "pre" {
		viper.SetConfigName("config-pre")
	} else if os.Getenv("MICROSERVICECAUTHNEWAPI") == "beta" {
		viper.SetConfigName("config-beta")
	} else {
		viper.SetConfigName("config")
	}

	log.Println("Current Config :", os.Getenv("MICROSERVICECAUTHNEWAPI"))

	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	dbReadWrite := viper.GetString("ENV_VAR_RW")
	dbConnection, err := mysql.DBConn(dbReadWrite)
	if err != nil {
		log.Fatalf("MySQL connection error , %v", err)
	} else {
		fmt.Printf("dbConnection connected: %v, %T", dbConnection, dbConnection)
	}
	ssr := handlers.ServiceSetupRouter{DB: dbConnection}
	router := ssr.SetupRouter()
	serverPort := viper.GetString("CONS_WEB_PORT")
	log.Printf("API environment :%v", viper.GetString("ENV_RUN_ENV"))
	listenAndServe(router, serverPort)
}

func main() {
	log.Println("Started in main func")
	flag.Parse()

	startService()
}

func listenAndServe(router *gin.Engine, port string) {
	log.Println("In listenAndServe start")
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Printf("Listening on address: %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Printf("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Printf("Server exiting")
}
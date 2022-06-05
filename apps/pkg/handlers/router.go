package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (ssr *ServiceSetupRouter) SetupRouter() *gin.Engine {
	r := gin.Default()
	corsConfig := CORS()

	r.Use(corsConfig)
	r.GET("/health", ssr.Health)

	// NewUserController
	usersHandler := NewUserController(ssr)

	// NewMerchantController
	merchantsHandler := NewMerchantController(ssr)
	{
		v1Group := r.Group("/merchants")
		{
			v1Group.POST("/merchants", merchantsHandler.Create)
			v1Group.GET("/merchants", merchantsHandler.List)
			v1Group.GET("/merchant/:code", merchantsHandler.ListByID)
			v1Group.PUT("/merchant/:code", merchantsHandler.UpdateByID)

			v1Group.POST("/members", usersHandler.Create)
			v1Group.GET("/member/:code/:email", usersHandler.ListByCode)
			v1Group.PUT("/member/:code/:email", usersHandler.UpdateByID)

			v1Group.GET("/members/:code", usersHandler.ListMembersByCode)
		}
	}
	return r
}

func (ssr *ServiceSetupRouter) Health(c *gin.Context) {
	pingDb(ssr.DB)
	c.JSON(200, gin.H{"action": "Health", "status": "success", "message": "Healtch Check OK"})
}

func pingDb(dbConnection *gorm.DB) {
	sqlDB, err := dbConnection.DB()
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = sqlDB.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return
	}
	log.Println("Connected to DB successfully")
}

func CORS() gin.HandlerFunc {
	config := cors.Config{}
	config.AllowHeaders = []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
	// config.AllowAllOrigins = true
	config.AllowBrowserExtensions = true
	config.AllowCredentials = true
	config.AllowWildcard = true
	config.AllowOrigins = []string{"*"}
	config.MaxAge = time.Hour * 12
	return cors.New(config)
}

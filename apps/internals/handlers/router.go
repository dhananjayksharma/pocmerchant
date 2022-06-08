package handlers

import (
	"dkgosql-merchant-service-v3/pkg/v1/models/merchants"
	"dkgosql-merchant-service-v3/pkg/v1/models/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(merchantService merchants.MerchantService, userService users.UserService) *gin.Engine {
	r := gin.Default()
	corsConfig := CORS()

	r.Use(corsConfig)
	healthHandler := NewHealthHandler()
	r.GET("/health", healthHandler.Health)

	// NewMerchantHandler
	merchantHandler := NewMerchantHandler(merchantService)
	// NewUserHandler
	userHandler := NewUserHandler(userService)
	{
		v1Group := r.Group("/merchants")
		{

			v1Group.PUT("/merchant/:code", merchantHandler.UpdateMerchantByID)
			v1Group.POST("/merchants", merchantHandler.CreateMerchant)
			v1Group.GET("/merchants", merchantHandler.GetMerchantList)

			v1Group.GET("/members/:code", userHandler.ListMembersByCode)
			v1Group.POST("/:code/member", userHandler.CreateMerchantMember)

		}
	}
	return r
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

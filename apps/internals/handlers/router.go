package handlers

import (
	"dkgosql-merchant-service-v4/internals/middleware"
	"dkgosql-merchant-service-v4/pkg/v1/models/merchants"
	"dkgosql-merchant-service-v4/pkg/v1/models/users"
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
			secured := v1Group.Group("/secured").Use(middleware.Auth())
			{
				secured.PUT("/merchant/:code", merchantHandler.UpdateMerchantByID)
				secured.POST("/merchants", merchantHandler.CreateMerchant)
				secured.GET("/merchants", merchantHandler.GetMerchantList) //.Use(auth.GetClaim(c))
				secured.GET("/members/:code", userHandler.ListMembersByCode)
				secured.POST("/:code/member", userHandler.CreateMerchantMember)
			}

			v1Group.POST("/member/login", userHandler.LoginMember)
			v1Group.GET("/member/refresh", userHandler.RefreshToken)

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

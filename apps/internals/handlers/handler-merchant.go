package handlers

import (
	"dkgosql-merchant-service-v4/internals/middleware"
	"dkgosql-merchant-service-v4/internals/util"
	"dkgosql-merchant-service-v4/pkg/v1/models/merchants"

	"github.com/gin-gonic/gin"
)

// MerchantHandler
type MerchantHandler interface {
	GetMerchantList(c *gin.Context)
	CreateMerchant(c *gin.Context)
	UpdateMerchantByID(c *gin.Context)
}

// merchantHandler
type merchantHandler struct {
	service merchants.MerchantService
}

// NewMerchantHandler
func NewMerchantHandler(service merchants.MerchantService) MerchantHandler {
	return &merchantHandler{service: service}
}

// GetMerchantList
func (srv *merchantHandler) GetMerchantList(c *gin.Context) {

	err := middleware.Claim(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	resp, err := srv.service.GetMerchantList(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

// CreateMerchant
func (srv *merchantHandler) CreateMerchant(c *gin.Context) {
	resp, err := srv.service.CreateMerchant(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

// UpdateMerchantByID
func (srv *merchantHandler) UpdateMerchantByID(c *gin.Context) {
	resp, err := srv.service.UpdateMerchantByID(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

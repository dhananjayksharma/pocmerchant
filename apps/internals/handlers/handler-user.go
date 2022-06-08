package handlers

import (
	"dkgosql-merchant-service-v3/internals/util"
	"dkgosql-merchant-service-v3/pkg/v1/models/users"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	ListMembersByCode(c *gin.Context)
	CreateMerchantMember(c *gin.Context)
}

type userHandler struct {
	service users.UserService
}

func NewUserHandler(service users.UserService) UserHandler {
	return &userHandler{service: service}
}

func (srv *userHandler) ListMembersByCode(c *gin.Context) {
	resp, err := srv.service.ListMembersByCode(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

func (srv *userHandler) CreateMerchantMember(c *gin.Context) {
	resp, err := srv.service.CreateMerchantMember(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

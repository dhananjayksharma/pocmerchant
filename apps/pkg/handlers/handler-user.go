package handlers

import (
	"dkgosql.com/pacenow-service/pkg/util"
	"dkgosql.com/pacenow-service/pkg/v1/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Create(c *gin.Context)

	ListByCode(c *gin.Context)
	UpdateByID(c *gin.Context)
	ListMembersByCode(c *gin.Context)
}

type userController struct {
	DB *gorm.DB
}

func NewUserController(svr *ServiceSetupRouter) UserController {
	return &userController{DB: svr.DB}
}

func (svr *userController) Create(c *gin.Context) {

	resp, err := models.NewUserService(svr.DB).Create(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.JSON(c, resp.Data, resp.Message)
}

func (svr *userController) ListByCode(c *gin.Context) {
	resp, err := models.NewUserService(svr.DB).ListByCode(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.JSON(c, resp.Data, resp.Message)
}

func (svr *userController) ListMembersByCode(c *gin.Context) {
	resp, err := models.NewUserService(svr.DB).ListMembersByCode(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.JSON(c, resp.Data, resp.Message)
}

func (svr *userController) UpdateByID(c *gin.Context) {
	resp, err := models.NewUserService(svr.DB).UpdateByID(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.JSON(c, resp.Data, resp.Message)
}

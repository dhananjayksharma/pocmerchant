package handlers

import (
	"dkgosql.com/pacenow-service/pkg/util"
	"dkgosql.com/pacenow-service/pkg/v1/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MerchantController interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	ListByID(c *gin.Context)
	UpdateByID(c *gin.Context)
}

type merchantController struct {
	DB *gorm.DB
}

func NewMerchantController(svr *ServiceSetupRouter) MerchantController {
	return &merchantController{DB: svr.DB}
}

func (svr *merchantController) Create(c *gin.Context) {

	resp, err := models.NewMerchantService(svr.DB).Create(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.JSON(c, resp.Data, resp.Message)
}

func (svr *merchantController) List(c *gin.Context) {
	resp, err := models.NewMerchantService(svr.DB).List(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.JSON(c, resp.Data, resp.Message)
}

func (svr *merchantController) ListByID(c *gin.Context) {
	resp, err := models.NewMerchantService(svr.DB).ListByID(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.JSON(c, resp.Data, resp.Message)
}

func (svr *merchantController) UpdateByID(c *gin.Context) {
	resp, err := models.NewMerchantService(svr.DB).UpdateByID(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	util.JSON(c, resp.Data, resp.Message)
}

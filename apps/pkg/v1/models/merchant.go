package models

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"dkgosql.com/pacenow-service/pkg/adapter/mysql/entities"
	"dkgosql.com/pacenow-service/pkg/adapter/mysql/query"
	"dkgosql.com/pacenow-service/pkg/consts"
	"dkgosql.com/pacenow-service/pkg/util"
	"dkgosql.com/pacenow-service/pkg/v1/models/request"
	"dkgosql.com/pacenow-service/pkg/v1/models/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MerchantService interface {
	Create(c *gin.Context) (Response, error)
	List(c *gin.Context) (Response, error)
	ListByID(c *gin.Context) (Response, error)
	UpdateByID(c *gin.Context) (Response, error)
}

type merchantService struct {
	DB *gorm.DB
}

func NewMerchantService(db *gorm.DB) MerchantService {
	return &merchantService{DB: db}
}

// Create merchant
func (service merchantService) Create(c *gin.Context) (Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var resp = Response{}
	var addMerchantRequest request.AddMerchantInputRequest
	if err := c.BindJSON(&addMerchantRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	resp.Data = "DATA DATA"
	resp.Message = "MESSAGE MESSAGE"
	currTime := time.Now()
	var status uint8
	status = uint8(consts.ActiveStatus)
	addMerchant := entities.Merchant{
		UpdatedAt: currTime,
		CreatedAt: currTime,
		Code:      addMerchantRequest.Code,
		Status:    &status,
		Name:      addMerchantRequest.Name,
		Address:   addMerchantRequest.Address,
	}

	var ms = query.NewMerchantDbConfig(service.DB)
	err := ms.Add(ctx, &addMerchant)
	if err != nil {
		return resp, err
	}

	var newMerchantMaster []response.MerchantResponse
	newMerchantMaster = append(newMerchantMaster, response.MerchantResponse{
		Name:      addMerchantRequest.Name,
		Code:      addMerchantRequest.Code,
		CreatedAt: currTime.String(),
		Address:   addMerchantRequest.Address,
	})

	resp.Data = newMerchantMaster
	resp.Message = consts.MerchantAddedSuccess
	return resp, nil
}

// List merchants
func (service merchantService) List(c *gin.Context) (Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var clientData []response.MerchantResponse
	var resp = Response{}
	if err != nil {
		return resp, err
	}
	var ms = query.NewMerchantDbConfig(service.DB)
	err = ms.List(ctx, &clientData)
	if err != nil {
		return resp, err
	}

	var outMSM []response.MerchantResponse
	for _, row := range clientData {
		outMSM = append(outMSM, response.MerchantResponse{
			Name:      row.Name,
			Code:      row.Code,
			Address:   row.Address,
			CreatedAt: row.CreatedAt,
		})
	}
	resp.Data = outMSM
	return resp, nil
}

// ListByID merchant
func (service merchantService) ListByID(c *gin.Context) (Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var clientData []response.MerchantResponse
	var resp = Response{}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err = errors.New(consts.InvalidMerchantCode)
		return resp, err
	}

	var ms = query.NewMerchantDbConfig(service.DB)
	err = ms.ListByID(ctx, &clientData, code)
	if err != nil {
		return resp, err
	}
	var responseMerchant []response.MerchantResponse
	for _, row := range clientData {
		responseMerchant = append(responseMerchant, response.MerchantResponse{
			Code:      row.Code,
			Name:      row.Name,
			CreatedAt: row.CreatedAt,
			Address:   row.Address,
		})
	}

	resp.Data = responseMerchant
	return resp, nil
}

func (service merchantService) UpdateByID(c *gin.Context) (Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var clientData entities.Merchant
	var resp = Response{}

	var updateMerchantRequest request.UpdateMerchantInputRequest
	if err := c.BindJSON(&updateMerchantRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err = errors.New(consts.InvalidCode)
		return resp, err
	}

	var ms = query.NewMerchantDbConfig(service.DB)
	var responseMerchant []response.MerchantResponse
	err = ms.ListByID(ctx, &responseMerchant, code)
	if err != nil {
		return resp, err
	}
	if len(responseMerchant) == 0 {
		err = errors.New(fmt.Sprintf(consts.ErrorDataNotFoundCode, code))
		return resp, err
	}

	var updateTypeData = make(map[string]interface{})
	updateTypeData["address"] = updateMerchantRequest.Address
	updateTypeData["name"] = updateMerchantRequest.Name

	err = ms.UpdateByID(ctx, &clientData, updateTypeData, code)
	if err != nil {
		return resp, err
	}

	err = ms.ListByID(ctx, &responseMerchant, code)
	if err != nil {
		return resp, err
	}
	resp.Data = responseMerchant
	resp.Message = fmt.Sprintf(consts.MerchantUpdatedSuccess, code)
	return resp, nil
}

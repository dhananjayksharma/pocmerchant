package models

import (
	"context"
	"errors"
	"fmt"
	"strconv"
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

type UserService interface {
	Create(c *gin.Context) (Response, error)
	ListByCode(c *gin.Context) (Response, error)
	UpdateByID(c *gin.Context) (Response, error)
	ListMembersByCode(c *gin.Context) (Response, error)
}

type userService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}

func (service userService) Create(c *gin.Context) (Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var resp = Response{}
	var addUserRequest request.AddUserInputRequest
	if err := c.BindJSON(&addUserRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	resp.Data = "DATA DATA"
	resp.Message = "MESSAGE MESSAGE"
	var status uint8
	status = uint8(consts.ActiveStatus)
	addUser := entities.Users{
		IsActive:  &status,
		FirstName: addUserRequest.FirstName,
		LastName:  addUserRequest.LastName,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		FkCode:    addUserRequest.Code,
		Email:     addUserRequest.Email,
	}

	var ms = query.NewUserDbConfig(service.DB)
	err := ms.Add(ctx, &addUser)
	if err != nil {
		return resp, err
	}

	var newSpotlightMaster []response.UsersResponse
	newSpotlightMaster = append(newSpotlightMaster, response.UsersResponse{
		IsActive:       &status,
		FirstName:      addUserRequest.FirstName,
		LastName:       addUserRequest.LastName,
		FkMerchantCode: addUserRequest.Code,
		Email:          addUserRequest.Email,
	})

	resp.Data = newSpotlightMaster
	resp.Message = consts.UserAddedSuccess
	return resp, nil
}

func (service userService) ListMembersByCode(c *gin.Context) (Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var userData []response.MerchantsMembersResponse
	var resp = Response{}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err = errors.New(consts.InvalidUserDocId)
		return resp, err
	}
	skip_number, err := strconv.ParseUint(c.Query("skip"), 10, 64)
	if skip_number < 0 || err != nil {

		if err != nil {
			return resp, err
		}
		err = errors.New(consts.SkipMessage)
		return resp, err
	}

	page_limit, _ := strconv.ParseUint(c.Query("limit"), 10, 64)

	if page_limit < 1 {
		err = errors.New(consts.PageLimitMessage)
		return resp, err
	}

	var queryParams = request.QueryMembersInputRequest{Code: code, Limit: int(page_limit), Skip: int(skip_number)}

	var ms = query.NewUserDbConfig(service.DB)
	err = ms.ListMembersByCode(ctx, &userData, queryParams)
	if err != nil {
		return resp, err
	}
	var responseUser []response.MerchantsMembersResponse
	for _, row := range userData {
		responseUser = append(responseUser, response.MerchantsMembersResponse{IsActive: row.IsActive,
			FirstName:    row.FirstName,
			LastName:     row.LastName,
			Email:        row.Email,
			Mobile:       row.Mobile,
			MerchantName: row.MerchantName,
			FkCode:       row.FkCode,
			CreatedAt:    row.CreatedAt,
		})
	}

	resp.Data = responseUser
	return resp, nil
}

func (service userService) ListByCode(c *gin.Context) (Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var userData []response.UsersResponse
	var resp = Response{}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err = errors.New(consts.InvalidUserDocId)
		return resp, err
	}

	var ms = query.NewUserDbConfig(service.DB)
	err = ms.ListByCode(ctx, &userData, code)
	if err != nil {
		return resp, err
	}
	var responseUser []response.UsersResponse
	for _, row := range userData {
		responseUser = append(responseUser, response.UsersResponse{IsActive: row.IsActive,
			FirstName:      row.FirstName,
			LastName:       row.LastName,
			FkMerchantCode: row.FkMerchantCode,
			Email:          row.Email})
	}

	resp.Data = responseUser
	return resp, nil
}

func (service userService) UpdateByID(c *gin.Context) (Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var userData entities.Users
	var resp = Response{}

	email := strings.Trim(c.Param("email"), "")
	if len(email) == 0 {
		err = errors.New(consts.InvalidCode)
		return resp, err
	}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err = errors.New(consts.InvalidCode)
		return resp, err
	}

	var updateUserInputRequest request.UpdateUserInputRequest
	if err := c.BindJSON(&updateUserInputRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	var ms = query.NewUserDbConfig(service.DB)
	var responseMovieSpotlight []response.UsersResponse
	err = ms.ListByCode(ctx, &responseMovieSpotlight, code)
	if err != nil {
		return resp, err
	}
	if len(responseMovieSpotlight) == 0 {
		err = errors.New(fmt.Sprintf(consts.ErrorDataNotFoundCode, code))
		return resp, err
	}

	var updateTypeData map[string]interface{}
	updateTypeData["first_name"] = updateUserInputRequest.FirstName
	updateTypeData["last_name"] = updateUserInputRequest.LastName
	updateTypeData["mobile"] = updateUserInputRequest.Mobile

	var queryParams = request.ListUserInputRequest{Code: code, Email: email}

	err = ms.UpdateByID(ctx, &userData, updateTypeData, queryParams)
	if err != nil {
		return resp, err
	}

	err = ms.ListByCode(ctx, &responseMovieSpotlight, code)
	if err != nil {
		return resp, err
	}
	resp.Data = responseMovieSpotlight
	resp.Message = "Update successfully!"
	return resp, nil
}

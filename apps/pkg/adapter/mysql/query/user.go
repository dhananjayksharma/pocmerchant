package query

import (
	"context"
	"fmt"
	"log"
	"strings"

	"dkgosql.com/pacenow-service/pkg/adapter/mysql/entities"
	"dkgosql.com/pacenow-service/pkg/consts"
	"dkgosql.com/pacenow-service/pkg/util"
	"dkgosql.com/pacenow-service/pkg/v1/models/request"
	"dkgosql.com/pacenow-service/pkg/v1/models/response"
	"gorm.io/gorm"
)

type UserDbAccess interface {
	Add(ctx context.Context, user *entities.Users) error
	List(ctx context.Context, user *[]response.UsersResponse) error
	ListByCode(ctx context.Context, user *[]response.UsersResponse, code string) error
	UpdateByID(ctx context.Context, user *entities.Users, updateTypeData map[string]interface{}, queryParams request.ListUserInputRequest) error
	// ListByEmailID(ctx context.Context, user *[]response.UsersResponse, queryParams request.ListUserInputRequest) error

	ListMembersByCode(ctx context.Context, user *[]response.MerchantsMembersResponse, queryParams request.QueryMembersInputRequest) error
}

type UserDbConfig struct {
	DB *gorm.DB
}

func NewUserDbConfig(db *gorm.DB) UserDbAccess {
	return &UserDbConfig{DB: db}
}
func (m *UserDbConfig) Add(ctx context.Context, user *entities.Users) error {

	log.Println("Add now")
	result := m.DB.Debug().WithContext(ctx).Create(&user)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			_userMsg := fmt.Sprintf(consts.ErrUserAlreadyExists, user.FkCode, user.Email)
			return &util.BadRequest{ErrMessage: _userMsg}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

func (m *UserDbConfig) List(ctx context.Context, user *[]response.UsersResponse) error {
	log.Println("List all ")
	result := m.DB.WithContext(ctx).Model(&response.UsersResponse{}).Select("fk_code, email, first_name, last_name").Find(&user)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

func (m *UserDbConfig) ListMembersByCode(ctx context.Context, user *[]response.MerchantsMembersResponse, queryParams request.QueryMembersInputRequest) error {

	log.Println("ListMembersByCode ")

	result := m.DB.Debug().WithContext(ctx).Model(&response.MerchantsMembersResponse{}).Select("users.fk_code, users.first_name, users.last_name, users.email, users.mobile, users.is_active, users.created_at, merchants.name as MerchantName").Joins("left join merchants on merchants.code = users.fk_code").Where("fk_code=?", queryParams.Code).Limit(queryParams.Limit).Offset(queryParams.Skip).Scan(&user)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorDataNotFoundCode, queryParams.Code)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

func (m *UserDbConfig) ListByCode(ctx context.Context, user *[]response.UsersResponse, code string) error {

	log.Println("ListByCode ") //.Where or
	result := m.DB.WithContext(ctx).Model(&response.UsersResponse{}).Select("fk_code, first_name, last_name, email, mobile, is_active, created_at, updated_at").Where("fk_code=?", code).Scan(&user)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorDataNotFoundCode, code)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

func (m *UserDbConfig) UpdateByID(ctx context.Context, user *entities.Users, updateTypeData map[string]interface{}, queryParams request.ListUserInputRequest) error {

	var updateFields = make(map[string]interface{})
	for key, val := range updateTypeData {
		updateFields[key] = val
	}

	result := m.DB.Debug().WithContext(ctx).Model(&user).Where("fk_code=? AND email=?", queryParams.Code, queryParams.Email).Updates(updateFields)

	log.Println("UpdateByID updated rows: ", result.RowsAffected)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	} else if result.RowsAffected == 0 {
		err := fmt.Sprintf(consts.ErrorUpdateMember, queryParams.Code, queryParams.Email)
		return &util.InternalServer{ErrMessage: err}
	}
	return nil
}

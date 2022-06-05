package query

import (
	"context"
	"fmt"
	"log"
	"strings"

	"dkgosql.com/pacenow-service/pkg/adapter/mysql/entities"
	"dkgosql.com/pacenow-service/pkg/consts"
	"dkgosql.com/pacenow-service/pkg/util"
	"dkgosql.com/pacenow-service/pkg/v1/models/response"
	"gorm.io/gorm"
)

type MerchantDbAccess interface {
	Add(ctx context.Context, user *entities.Merchant) error
	List(ctx context.Context, user *[]response.MerchantResponse) error
	ListByID(ctx context.Context, user *[]response.MerchantResponse, code string) error
	UpdateByID(ctx context.Context, user *entities.Merchant, updateTypeData map[string]interface{}, code string) error
}

type MerchantDbConfig struct {
	DB *gorm.DB
}

func NewMerchantDbConfig(db *gorm.DB) MerchantDbAccess {
	return &MerchantDbConfig{DB: db}
}
func (m *MerchantDbConfig) Add(ctx context.Context, user *entities.Merchant) error {

	log.Println("Add now")
	result := m.DB.Debug().WithContext(ctx).Create(&user)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			_userMsg := fmt.Sprintf(consts.ErrMerchantAlreadyExists, user.Code)
			return &util.BadRequest{ErrMessage: _userMsg}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

func (m *MerchantDbConfig) List(ctx context.Context, user *[]response.MerchantResponse) error {
	log.Println("List all ")
	result := m.DB.WithContext(ctx).Model(&response.MerchantResponse{}).Select("code,  name, address, status, created_at, updated_at").Find(&user)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

func (m *MerchantDbConfig) ListByID(ctx context.Context, user *[]response.MerchantResponse, code string) error {

	log.Println("ListByID ")
	result := m.DB.WithContext(ctx).Model(&response.MerchantResponse{}).Select("code, name, address, status, created_at, updated_at").Where("code=?", code).Scan(&user)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorDataNotFoundCode, code)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

func (m *MerchantDbConfig) UpdateByID(ctx context.Context, user *entities.Merchant, updateTypeData map[string]interface{}, code string) error {

	var updateFields = make(map[string]interface{})
	for key, val := range updateTypeData {
		updateFields[key] = val
	}

	result := m.DB.Debug().WithContext(ctx).Model(&user).Where("code=?", code).Omit("code", "id").Updates(updateFields)

	//result := m.DB.Debug().WithContext(ctx).Model(&user).Where("code=?", code).Select("address", "name").Updates(updateFields)

	log.Println("UpdateByID updated rows: ", result.RowsAffected)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	} else if result.RowsAffected == 0 {
		err := fmt.Sprintf(consts.ErrorUpdateType, code)
		return &util.InternalServer{ErrMessage: err}
	}
	return nil
}

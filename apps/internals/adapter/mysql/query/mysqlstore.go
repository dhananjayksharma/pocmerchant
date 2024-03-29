package query

import (
	"context"
	"dkgosql-merchant-service-v4/internals/adapter/mysql/entities"
	"dkgosql-merchant-service-v4/internals/consts"
	"dkgosql-merchant-service-v4/internals/util"
	"dkgosql-merchant-service-v4/pkg/v1/models/request"
	"dkgosql-merchant-service-v4/pkg/v1/models/response"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type mySQLDBStore struct {
	db *gorm.DB
}

func NewMySQLDBStore(db *gorm.DB) MySQLDBStoreAccess {
	return &mySQLDBStore{db: db}
}

type MySQLDBStoreAccess interface {
	GetMerchantList(ctx context.Context, merchantData *[]response.MerchantResponse) error
	CreateMerchant(ctx context.Context, merchantData *entities.Merchant) error
	ListMerchantByID(ctx context.Context, merchantData *[]response.MerchantResponse, code string) error
	UpdateMerchantByID(ctx context.Context, user *entities.Merchant, updateTypeData map[string]interface{}, code string) error

	CreateMerchantMember(ctx context.Context, user *entities.Users) error

	ListMembersByCode(ctx context.Context, user *[]response.MerchantsMembersResponse, queryParams request.QueryMembersInputRequest) error

	LoginUserByEmailID(ctx context.Context, userData *[]response.UserLoginResponse, queryParams request.LoginUserInputRequest) error
}

// CreateMerchantMember
func (ms *mySQLDBStore) CreateMerchantMember(ctx context.Context, user *entities.Users) error {
	result := ms.db.Debug().WithContext(ctx).Create(&user)
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

// UpdateMerchantByID
func (ms *mySQLDBStore) UpdateMerchantByID(ctx context.Context, user *entities.Merchant, updateTypeData map[string]interface{}, code string) error {

	var updateFields = make(map[string]interface{})
	for key, val := range updateTypeData {
		updateFields[key] = val
	}

	result := ms.db.Debug().WithContext(ctx).Model(&user).Where("code=?", code).Omit("code", "id").Updates(updateFields)

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

// CreateMerchant
func (ms *mySQLDBStore) CreateMerchant(ctx context.Context, merchant *entities.Merchant) error {
	result := ms.db.Debug().WithContext(ctx).Create(&merchant)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			_userMsg := fmt.Sprintf(consts.ErrMerchantAlreadyExists, merchant.Code)
			return &util.BadRequest{ErrMessage: _userMsg}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

// ListMerchantByID
func (ms *mySQLDBStore) ListMerchantByID(ctx context.Context, merchantData *[]response.MerchantResponse, code string) error {

	log.Println("ListMerchantByID ")
	result := ms.db.Debug().WithContext(ctx).Model(&response.MerchantResponse{}).Select("code, name, address, status, created_at, updated_at").Where("code=?", code).Scan(&merchantData)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorDataNotFoundCode, code)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

// ListMembersByCode
func (ms *mySQLDBStore) LoginUserByEmailID(ctx context.Context, userData *[]response.UserLoginResponse, queryParams request.LoginUserInputRequest) error {

	result := ms.db.Debug().WithContext(ctx).Model(&response.UserLoginResponse{}).Select("users.fk_code, users.first_name, users.last_name, users.email, users.mobile, users.password, users.is_active, users.created_at, merchants.name as MerchantName").Joins("left join merchants on merchants.code = users.fk_code").Where("fk_code=? AND users.email=?", queryParams.Code, queryParams.Email).Scan(&userData)

	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorUserNotFoundCode, queryParams.Code)}
	}
	
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

// ListMembersByCode
func (ms *mySQLDBStore) ListMembersByCode(ctx context.Context, merchant *[]response.MerchantsMembersResponse, queryParams request.QueryMembersInputRequest) error {

	result := ms.db.Debug().WithContext(ctx).Model(&response.MerchantsMembersResponse{}).Select("users.fk_code, users.first_name, users.last_name, users.email, users.mobile, users.is_active, users.created_at, merchants.name as MerchantName").Joins("left join merchants on merchants.code = users.fk_code").Where("fk_code=?", queryParams.Code).Limit(queryParams.Limit).Offset(queryParams.Skip).Scan(&merchant)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorDataNotFoundCode, queryParams.Code)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

// GetMerchantList
func (ms *mySQLDBStore) GetMerchantList(ctx context.Context, merchantData *[]response.MerchantResponse) error {
	result := ms.db.WithContext(ctx).Model(&response.MerchantResponse{}).Select("code,  name, address, status, created_at, updated_at").Find(&merchantData)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

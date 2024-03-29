package query

import (
	"context"
	"dkgosql-merchant-service-v4/internals/adapter/mysql/entities"
	"dkgosql-merchant-service-v4/pkg/v1/models/request"
	"dkgosql-merchant-service-v4/pkg/v1/models/response"
)

type MockMySQLDBStore struct {
}

var _ MySQLDBStoreAccess = (*MockMySQLDBStore)(nil)

// UpdateMerchantByID
func (ms *MockMySQLDBStore) ListMerchantByID(ctx context.Context, merchantData *[]response.MerchantResponse, code string) error {
	return nil
}

// UpdateMerchantByID
func (ms *MockMySQLDBStore) UpdateMerchantByID(ctx context.Context, user *entities.Merchant, updateTypeData map[string]interface{}, code string) error {
	return nil
}

// CreateMerchantMember
func (ms *MockMySQLDBStore) CreateMerchantMember(ctx context.Context, user *entities.Users) error {
	return nil
}

// GetMerchantList
func (ms *MockMySQLDBStore) GetMerchantList(ctx context.Context, merchantData *[]response.MerchantResponse) error {
	data := []response.MerchantResponse{
		{
			Code:      "1454dddd",
			Name:      "TestMerchant",
			CreatedAt: "2022-06-04 16:40:28",
			Address:   "Mumbai",
		}, {
			Code:      "124578d3e",
			Name:      "TestMerchant2",
			CreatedAt: "2022-06-04 16:40:28",
			Address:   "Mumbai",
		},
	}
	*merchantData = data
	return nil
}

// ListMembersByCode
func (ms *MockMySQLDBStore) ListMembersByCode(ctx context.Context, user *[]response.MerchantsMembersResponse, queryParams request.QueryMembersInputRequest) error {
	data := []response.MerchantsMembersResponse{
		{
			MerchantName: "TestMerchant",
		}, {
			MerchantName: "TestMerchant2",
		},
	}
	*user = data
	return nil
}

// CreateMerchant
func (ms *MockMySQLDBStore) CreateMerchant(ctx context.Context, user *entities.Merchant) error {
	var data = entities.Merchant{
		Name:    "Sony New ltd",
		Address: "Mumbai, Ville Parle",
		Code:    "cadjq02gqpmvljdra98",
	}
	*user = data
	return nil
}

// LoginUserByEmailID
func (ms *MockMySQLDBStore) LoginUserByEmailID(ctx context.Context, userData *[]response.UserLoginResponse, queryParams request.LoginUserInputRequest) error {
	return nil
}

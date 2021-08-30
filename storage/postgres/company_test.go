package postgres_test

import (
	"testing"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"github.com/stretchr/testify/assert"
)

func createCompany(t *testing.T) *pb.Company {

	company := &pb.Company{
		Id:        "1",
		Name:      "smth",
		Inn:       "sdsds",
		OwnerName: "man",
	}

	res, err := strg.Company().Create(company)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return company
}

func deleteCompany(t *testing.T, id string) {
	err := strg.Company().Delete(id)

	assert.NoError(t, err)
}

func TestNewCompanyRepo(t *testing.T) {
	company := createCompany(t)
	deleteCompany(t, company.Id)
}

func TestUpdateCompany(t *testing.T) {
	company := createCompany(t)

	company.Name = fakeData.CompanyName()

	id, err := strg.Company().Update(company)

	assert.NotEmpty(t, id)
	assert.NoError(t, err)
	deleteCompany(t, company.Id)
}

func TestCompanyGet(t *testing.T) {
	company := createCompany(t)

	_, err := strg.Company().Get(company.Id)

	assert.NoError(t, err)
	deleteCompany(t, company.Id)
}

func TestCompanyGetAll(t *testing.T) {
	company := createCompany(t)
	resp, err := strg.Company().GetAll(&pb.GetAllCompanyRequest{
		Offset: 0,
		Limit:  10,
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Companies)
	deleteCompany(t, company.Id)
}

func TestCompanyDelete(t *testing.T) {
	company := createCompany(t)
	deleteCompany(t, company.Id)
}

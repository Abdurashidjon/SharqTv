package postgres_test

import (
	"fmt"
	"testing"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"github.com/stretchr/testify/assert"
)

func createResearcher(t *testing.T) *pb.Researcher {
	company := createCompany(t)
	researcher := &pb.Researcher{
		Id:              createRandomId(t),
		Name:            fakeData.Name(),
		Email:           fakeData.Email(),
		Phone:           fakeData.PhoneNumber(),
		CompanyId:       company.Id,
		RoleId:          createRandomId(t),
		ProfessionTitle: fakeData.Characters(10),
	}

	res, err := strg.Researcher().Create(researcher)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return researcher
}

func deleteResearcher(t *testing.T, id string) {
	err := strg.Researcher().Delete(id)

	assert.NoError(t, err)
}

func TestCreateResearcher(t *testing.T) {
	id := createResearcher(t)
	deleteResearcher(t, id.Id)
}

func TestUpdateResearcher(t *testing.T) {
	researcher := createResearcher(t)

	researcher.Name = fakeData.Name()
	researcher.Phone = fakeData.PhoneNumber()
	researcher.ProfessionTitle = fakeData.Characters(10)

	id, err := strg.Researcher().Update(researcher)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	deleteResearcher(t, researcher.Id)
}

func TestResearcherGet(t *testing.T) {
	id := createResearcher(t).Id

	_, err := strg.Researcher().Get(id)

	assert.NoError(t, err)

	deleteResearcher(t, id)
}

func TestResearcherGetAll(t *testing.T) {
	company := createCompany(t)
	researcher := createResearcher(t)
	resp, err := strg.Researcher().GetAll(&pb.GetAllResearcherRequest{
		Offset:    0,
		Limit:     10,
		CompanyId: company.Id,
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	deleteResearcher(t, researcher.Id)
}

func TestResearcherDelete(t *testing.T) {
	id := createResearcher(t).Id
	deleteResearcher(t, id)
}

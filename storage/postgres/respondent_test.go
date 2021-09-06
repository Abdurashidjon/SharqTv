package postgres_test

import (
	"testing"

	pb "bitbucket.org/udevs/ur_go_user_service/genproto/user_service"
	"github.com/stretchr/testify/assert"
)

func createRespondent(t *testing.T) *pb.Respondent {
	company := createCompany(t)
	respondent := &pb.Respondent{
		Id:      createRandomId(t),
		Name:    fakeData.Name(),
		Email:   fakeData.Email(),
		Phone:   fakeData.PhoneNumber(),
		Company: company.Id,
	}

	res, err := strg.Respondent().Create(respondent)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return respondent
}

func deleteRespondent(t *testing.T, id string) {
	err := strg.Respondent().Delete(id)

	assert.NoError(t, err)
}

func TestCreateRespondent(t *testing.T) {
	id := createRespondent(t).Id
	deleteRespondent(t, id)
}

func TestUpdateRespondent(t *testing.T) {
	respondent := createRespondent(t)

	respondent.Name = fakeData.Name()
	respondent.Phone = fakeData.PhoneNumber()

	id, err := strg.Respondent().Update(respondent)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	deleteResearcher(t, respondent.Id)
}

func TestRespondentGet(t *testing.T) {
	id := createRespondent(t).Id

	_, err := strg.Respondent().Get(id)

	assert.NoError(t, err)

	deleteRespondent(t, id)
}

func TestRespondentGetAll(t *testing.T) {
	respondent := createRespondent(t)
	resp, err := strg.Respondent().GetAll(&pb.GetAllRespondentRequest{
		Offset: 0,
		Limit:  10,
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Respondents)
	deleteRespondent(t, respondent.Id)
}

func TestRespondentDelete(t *testing.T) {
	id := createRespondent(t).Id
	deleteRespondent(t, id)
}

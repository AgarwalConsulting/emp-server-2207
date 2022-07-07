package service_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := service.NewV1(mockRepo)

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockRepo.EXPECT().ListAll().Return(expectedEmps, nil)

	actualEmployees, err := sut.Index()

	assert.Nil(t, err)
	assert.NotNil(t, actualEmployees)

	assert.Equal(t, expectedEmps, actualEmployees)
}

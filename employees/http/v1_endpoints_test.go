package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/service"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	mockSvc.EXPECT().Index().Return(nil, nil)

	req := httptest.NewRequest("GET", "/v1/employees", nil)
	resRec := httptest.NewRecorder()

	sut.IndexV1(resRec, req)

	result := resRec.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)
}

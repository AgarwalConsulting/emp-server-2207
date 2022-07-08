package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestIndexV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockSvc.EXPECT().Index().Return(expectedEmps, nil)

	req := httptest.NewRequest("GET", "/v1/employees", nil)
	resRec := httptest.NewRecorder()

	// sut.IndexV1(resRec, req)
	sut.ServeHTTP(resRec, req)

	result := resRec.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	var actualEmps []entities.Employee

	json.NewDecoder(result.Body).Decode(&actualEmps)

	assert.NotNil(t, actualEmps)
	assert.Equal(t, 1, len(actualEmps))
	assert.Equal(t, "Gaurav", actualEmps[0].Name)
}

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	newEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 1001}
	createdEmp := newEmp
	createdEmp.ID = 1
	mockSvc.EXPECT().Create(newEmp).Return(&createdEmp, nil)

	resRec := httptest.NewRecorder()

	jsonBody := `{"name": "Gaurav", "speciality": "LnD", "project": 1001}`
	reqBody := strings.NewReader(jsonBody)
	req := httptest.NewRequest("POST", "/v1/employees", reqBody)

	// sut.CreateV1(resRec, req)
	sut.ServeHTTP(resRec, req)

	result := resRec.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	var actualEmp entities.Employee

	json.NewDecoder(result.Body).Decode(&actualEmp)

	// assert.Equal(t, "Gaurav", actualEmp.Name)
	assert.Equal(t, newEmp, actualEmp)
}

func FuzzCreateV1(f *testing.F) {
	sampleJSONBody := `{"name": "Gaurav`

	f.Add(sampleJSONBody)

	f.Fuzz(func(t *testing.T, jsonBody string) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := service.NewMockEmployeeService(ctrl)

		sut := empHTTP.NewHandler(mockSvc)

		resRec := httptest.NewRecorder()

		reqBody := strings.NewReader(jsonBody)
		req := httptest.NewRequest("POST", "/v1/employees", reqBody)

		// sut.CreateV1(resRec, req)
		sut.ServeHTTP(resRec, req)

		result := resRec.Result()

		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})
}

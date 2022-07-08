package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	existingEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, existingEmps)

	existingCount := len(existingEmps)
	assert.NotEqual(t, 0, existingCount)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		newEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 1001}

		go func() {
			defer wg.Done()
			_, err := sut.Create(newEmp)
			assert.Nil(t, err)

			_, err = sut.ListAll()
			assert.Nil(t, err)
		}()
	}

	wg.Wait()

	allEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, allEmps)

	newCount := len(allEmps)

	assert.Equal(t, existingCount+100, newCount)
}

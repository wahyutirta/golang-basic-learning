package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserDB struct {
	mock.Mock
}

func (m *MockUserDB) FindByID(ID string) (string, error) {
	args := m.Called(ID)
	return args.Get(0).(string), args.Error(1)
}

func TestGetuser(t *testing.T) {
	t.Run("can get user", func(t *testing.T) {
		userDBMock := new(MockUserDB)
		userDBMock.On("FindByID", "1").Return("wahyu", nil)

		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := GetUser(userDBMock)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusOK, rr.Code)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body is what we expect.
		expected := `{"name": "wahyu"}`
		assert.JSONEq(t, expected, rr.Body.String())
		// if rr.Body.String() != expected {
		// 	t.Errorf("handler returned unexpected body: got %v want %v",
		// 		rr.Body.String(), expected)
		// }
		userDBMock.AssertExpectations(t)
		// test PutObject func to expect PutObjectInput to receive specific object args and return specific value
	})

	t.Run("can not get user, db error", func(t *testing.T) {
		userDBMock := new(MockUserDB)
		userDBMock.On("FindByID", "1").Return("", fmt.Errorf("fail connect to db"))

		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := GetUser(userDBMock)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusInternalServerError)
		}
	})

}

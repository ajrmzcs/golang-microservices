package models

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	// Traditional way
	if user != nil {
		t.Error("we were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("we were expecting an error for user with id 0")
	}

	if err != nil && err.Status != http.StatusNotFound {
		t.Error("we were expecting a 404 status when user is not found")
	}

	// Package assert
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	if err != nil {
		assert.EqualValues(t, http.StatusNotFound, err.Status)
		assert.EqualValues(t, "not found", err.Code)
		assert.EqualValues(t, "user 0 not found", err.Message)
	}
}

func TestGetUserNotError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	if user != nil {
		assert.EqualValues(t, 1, user.Id)
		assert.EqualValues(t, "Antonio", user.FirstName)
		assert.EqualValues(t, "Ramirez", user.LastName)
		assert.EqualValues(t, "ajrmzcs@gmail.com", user.Email)
	}
}
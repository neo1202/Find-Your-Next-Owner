package test

import (
	"fyno/server/internal/models"
	"fyno/server/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserInterface := mock.NewMockUserInterface(ctrl)

	// Set up expected call to Get
	expectedID := uuid.New()
	expectedUser := &models.User{ID: expectedID}
	mockUserInterface.EXPECT().Get(expectedID).Return(expectedUser, nil)

	// Call the function being tested
	user, err := mockUserInterface.Get(expectedID)

	// Verify that the mocked interface was called with the expected parameters
	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}
	if user == nil {
		t.Errorf("Get returned an unexpected user: %v", user)
	}
	if user.ID != expectedID {
		t.Errorf("Get returned an unexpected user: %v", user)
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserInterface := mock.NewMockUserInterface(ctrl)

	// Set up expected call to Create
	expectedUser := models.User{ID: uuid.New()}
	mockUserInterface.EXPECT().Create(expectedUser).Return(uuid.New(), nil)

	// Call the function being tested
	id, err := mockUserInterface.Create(expectedUser)

	// Verify that the mocked interface was called with the expected parameters
	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}
	if id == uuid.Nil {
		t.Errorf("Create returned an empty UUID")
	}
}

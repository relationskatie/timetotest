package server

import (
	"context"
	"github.com/golang/mock/gomock"
	mockstorage "github.com/relationskatie/timetotest/internal/storage/mock/storage_mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestController_Run(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mockstorage.NewMockInterface(mockCtrl)
	ctrl := testController(t, mockStore)

	go func() {
		err := ctrl.Run(context.Background())
		assert.NoError(t, err)
	}()
}
func TestController_Shutdown(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mockstorage.NewMockInterface(mockCtrl)
	ctrl := testController(t, mockStore)

	err := ctrl.Shutdown(context.Background())
	assert.NoError(t, err)
}

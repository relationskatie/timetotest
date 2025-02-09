package server

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/relationskatie/timetotest/internal/modles"
	mockstorage "github.com/relationskatie/timetotest/internal/storage/mock/storage_mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*func TestRight(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockStore := mockstorage.NewMockInterface(mockCtrl)
	mockUser := mockstorage.NewMockUserStorage(mockCtrl)

	mockStore.EXPECT().User().Return(mockUser)

	ctrl := testController(t, mockStore)

	userID := uuid.New()
	user := modles.UserDTO{
		ID:        userID,
		Name:      "Test User",
		Username:  "Test User",
		Age:       19,
		Telephone: "234567890",
	}

	mockUser.EXPECT().GetUserByID(gomock.Any(), userID).Return(user, nil)

	req := httptest.NewRequest(http.MethodGet, "/user/"+userID.String(), nil)

	rec := httptest.NewRecorder()

	ctx := ctrl.srv.NewContext(req, rec)

	ctx.SetParamNames("id")
	ctx.SetParamValues(userID.String())

	err := ctrl.handleGetUserByID(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}*/

func TestGetUserByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ID := uuid.New()

	tt := []struct {
		name          string
		ID            string
		user          modles.UserDTO
		expectedCode  int
		positive      bool
		userIsCreated bool
	}{
		{
			name: "Test 1: valid UUID with exist user",
			ID:   ID.String(),
			user: modles.UserDTO{
				ID:        ID,
				Name:      "Test User",
				Username:  "Test User",
				Age:       19,
				Telephone: "234567890",
			},
			expectedCode:  http.StatusOK,
			positive:      true,
			userIsCreated: true,
		},
		{
			name:          "Test 2: valid UUID with don't exist user",
			ID:            ID.String(),
			user:          modles.UserDTO{},
			expectedCode:  http.StatusNotFound,
			positive:      false,
			userIsCreated: true,
		},
		{
			name:          "Test 3: invalid UUID",
			ID:            "efuhdwkl;1323",
			user:          modles.UserDTO{},
			expectedCode:  http.StatusBadRequest,
			positive:      false,
			userIsCreated: false,
		},
		{
			name:          "Negative: ID is empty",
			ID:            "",
			user:          modles.UserDTO{},
			expectedCode:  http.StatusBadRequest,
			positive:      true,
			userIsCreated: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockstorage.NewMockInterface(mockCtrl)
			mockUser := mockstorage.NewMockUserStorage(mockCtrl)

			mockStore.EXPECT().User().Return(mockUser).AnyTimes()

			ctrl := testController(t, mockStore)

			if tc.userIsCreated {
				if tc.positive {
					mockUser.EXPECT().GetUserByID(gomock.Any(), ID).Return(tc.user, nil)
				} else {
					mockUser.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(tc.user, errors.New("error"))
				}
			}

			req := httptest.NewRequest(http.MethodGet, "/user/"+tc.ID, nil)
			rec := httptest.NewRecorder()

			ctx := ctrl.srv.NewContext(req, rec)
			ctx.SetParamNames("id")
			ctx.SetParamValues(tc.ID)

			_ = ctrl.handleGetUserByID(ctx)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestHandleAddNewUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tt := []struct {
		name         string
		req          modles.AddUserRequest
		expectedCode int
		positive     bool
	}{
		{
			name: "Test 1: successful user addition",
			req: modles.AddUserRequest{
				Name:      "John Doe",
				Username:  "johndoe",
				Age:       30,
				Telephone: "1234567890",
			},
			expectedCode: http.StatusCreated,
			positive:     true,
		},
		{
			name: "Test 2: invalid request (missing field)",
			req: modles.AddUserRequest{
				Username:  "johndoe",
				Age:       30,
				Telephone: "1234567890",
			},
			expectedCode: http.StatusInternalServerError,
			positive:     false,
		},
		{
			name: "Test 3: error adding user to storage",
			req: modles.AddUserRequest{
				Name:      "John Doe",
				Username:  "johndoe",
				Age:       30,
				Telephone: "1234567890",
			},
			expectedCode: http.StatusInternalServerError,
			positive:     false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockstorage.NewMockInterface(mockCtrl)
			mockUser := mockstorage.NewMockUserStorage(mockCtrl)

			mockStore.EXPECT().User().Return(mockUser).AnyTimes()

			ctrl := testController(t, mockStore)

			if tc.positive {
				mockUser.EXPECT().AddNewUser(gomock.Any(), gomock.Any()).Return(nil)
			} else {
				mockUser.EXPECT().AddNewUser(gomock.Any(), gomock.Any()).Return(errors.New("error adding user"))
			}

			reqBody := fmt.Sprintf(`{"name":"%s","username":"%s","age":%d,"telephone":"%s"}`,
				tc.req.Name, tc.req.Username, tc.req.Age, tc.req.Telephone)
			req := httptest.NewRequest(http.MethodPost, "/add_user/", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			ctx := ctrl.srv.NewContext(req, rec)

			err := ctrl.handleAddNewUser(ctx)

			if tc.positive {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
func TestHandleGetAllUsers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tt := []struct {
		name         string
		expectedCode int
		users        []modles.UserDTO
		positive     bool
	}{
		{
			name: "Test 1: successful retrieval of users",
			users: []modles.UserDTO{
				{
					ID:        uuid.New(),
					Name:      "User1",
					Username:  "user1",
					Age:       25,
					Telephone: "123456789",
				},
				{
					ID:        uuid.New(),
					Name:      "User2",
					Username:  "user2",
					Age:       30,
					Telephone: "987654321",
				},
			},
			expectedCode: http.StatusOK,
			positive:     true,
		},
		{
			name:         "Test 2: users don't exist",
			users:        nil,
			expectedCode: http.StatusOK,
			positive:     true,
		},
		{
			name:         "Test 3: error in db",
			users:        nil,
			expectedCode: http.StatusInternalServerError,
			positive:     false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockstorage.NewMockInterface(mockCtrl)
			mockUser := mockstorage.NewMockUserStorage(mockCtrl)

			mockStore.EXPECT().User().Return(mockUser).AnyTimes()

			ctrl := testController(t, mockStore)

			if tc.positive {
				mockUser.EXPECT().GetAllUsers(gomock.Any()).Return(tc.users, nil)
			} else {
				mockUser.EXPECT().GetAllUsers(gomock.Any()).Return(nil, errors.New("error"))
			}

			req := httptest.NewRequest(http.MethodGet, "/return_all_users/", nil)
			rec := httptest.NewRecorder()

			ctx := ctrl.srv.NewContext(req, rec)

			err := ctrl.handleGetAllUsers(ctx)

			if tc.positive {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
func TestHandleChangeUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tt := []struct {
		name         string
		requestBody  string
		expectedCode int
		positive     bool
	}{
		{
			name: "Test 1: Successful user update",
			requestBody: `{
				"name": "Updated User",
				"age": 25,
				"telephone": "123456789",
				"username": "updated_user"
			}`,
			expectedCode: http.StatusOK,
			positive:     true,
		},
		{
			name:         "Test 2: Invalid JSON",
			requestBody:  `{ "name": "Test User", "age": "invalid" }`,
			expectedCode: http.StatusBadRequest,
			positive:     false,
		},
		{
			name: "Test 3: Database error",
			requestBody: `{
				"name": "Test User",
				"age": 30,
				"telephone": "987654321",
				"username": "test_user"
			}`,
			expectedCode: http.StatusInternalServerError,
			positive:     false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockstorage.NewMockInterface(mockCtrl)
			mockUser := mockstorage.NewMockUserStorage(mockCtrl)
			mockStore.EXPECT().User().Return(mockUser).AnyTimes()

			ctrl := testController(t, mockStore)

			if tc.positive {
				mockUser.EXPECT().ChangeUser(gomock.Any(), gomock.Any()).Return(nil)
			} else if tc.expectedCode == http.StatusInternalServerError {
				mockUser.EXPECT().ChangeUser(gomock.Any(), gomock.Any()).Return(errors.New("internal server error"))
			}

			req := httptest.NewRequest(http.MethodPost, "/user/change", strings.NewReader(tc.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			ctx := ctrl.srv.NewContext(req, rec)
			_ = ctrl.handleChangeUser(ctx)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
func TestHandleDeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tt := []struct {
		name         string
		username     string
		expectedCode int
		positive     bool
	}{
		{
			name:         "Test 1: successful user deletion",
			username:     "test1",
			expectedCode: http.StatusNoContent,
			positive:     true,
		},
		{
			name:         "Test 2: database error",
			username:     "test2",
			expectedCode: http.StatusInternalServerError,
			positive:     false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockstorage.NewMockInterface(mockCtrl)
			mockUser := mockstorage.NewMockUserStorage(mockCtrl)

			mockStore.EXPECT().User().Return(mockUser).AnyTimes()

			ctrl := testController(t, mockStore)

			if tc.positive {
				mockUser.EXPECT().DeleteUserByUsername(gomock.Any(), tc.username).Return(nil)
			} else {
				mockUser.EXPECT().DeleteUserByUsername(gomock.Any(), tc.username).Return(errors.New("error deleting user"))
			}

			req := httptest.NewRequest(http.MethodDelete, "/delete_user/"+tc.username, nil)
			rec := httptest.NewRecorder()

			ctx := ctrl.srv.NewContext(req, rec)
			ctx.SetParamNames("name")
			ctx.SetParamValues(tc.username)

			err := ctrl.handleDeleteUser(ctx)

			if tc.positive {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

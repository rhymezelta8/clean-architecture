package user

import (
	"arquitectura/domain/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	//userUseCaseMock := new(usecasemocks.UseCaseUser)
	//userUseCaseMock.On("Create", mock.Anything).Return(nil)
	//
	userHandler := newHandler(userUseCaseMock)
	//
	//gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/v1/users/user", handler.Create)
	//userHandler.Create()
	//t.Run("given an invalid request it returns 400", func(t *testing.T) {
	//	userReq := model.User{ID: uuid.New()}
	//
	//	b, err := json.Marshal(userReq)
	//	require.NoError(t, err)
	//
	//	req, err := http.NewRequest(http.MethodPost, "v1/users/user", bytes.NewReader(b))
	//	require.NoError(t, err)
	//
	//	rec := httptest.NewRecorder()
	//	r.ServeHTTP(rec, req)
	//
	//	res := rec.Result()
	//	defer res.Body.Close()
	//
	//	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	//
	//})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		userReq := model.User{
			ID:   uuid.New(),
			Name: "Melany",
		}

		b, err := json.Marshal(userReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/v1/users/user", bytes.NewReader(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)

	})
}

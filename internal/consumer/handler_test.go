package consumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
	"xyz_golang/internal/domain"
	"xyz_golang/internal/mocks"
)

func TestHttpHandler_GetByID(t *testing.T) {
	var mockConsumer domain.Consumer
	err := faker.FakeData(&mockConsumer)
	assert.NoError(t, err)
	mockService := new(mocks.ConsumerService)

	t.Run("success", func(t *testing.T) {
		mockService.On("GetByID", mockConsumer.ID).
			Return(&mockConsumer, nil).Once()

		app := fiber.New()
		NewHttpHandler(app, mockService)
		id := mockConsumer.ID

		resp, err := app.Test(httptest.NewRequest("GET", "/"+id.String(), nil))
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockService.On("GetByID", mockConsumer.ID).
			Return(nil, fiber.ErrNotFound).Once()
		app := fiber.New()
		NewHttpHandler(app, mockService)
		id := mockConsumer.ID
		resp, err := app.Test(httptest.NewRequest("GET", "/"+id.String(), nil))
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockService.On("GetByID", mockConsumer.ID).
			Return(nil, fiber.ErrInternalServerError).Once()
		app := fiber.New()
		NewHttpHandler(app, mockService)
		id := mockConsumer.ID
		resp, err := app.Test(httptest.NewRequest("GET", "/"+id.String(), nil))
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestHttpHandler_GetLimit(t *testing.T) {
	var mockLimits []domain.Limit
	var mockLimit domain.Limit
	err := faker.FakeData(&mockLimit)
	assert.NoError(t, err)
	mockLimits = append(mockLimits, mockLimit)
	mockService := new(mocks.ConsumerService)

	t.Run("success", func(t *testing.T) {
		mockService.On("GetLimit", mockLimit.ConsumerID).
			Return(&mockLimits, nil).Once()

		app := fiber.New()
		NewHttpHandler(app, mockService)
		id := mockLimit.ConsumerID

		resp, err := app.Test(httptest.NewRequest("GET", "/"+id.String()+"/limit", nil))
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockService.On("GetLimit", mockLimit.ConsumerID).
			Return(nil, fiber.ErrNotFound).Once()
		app := fiber.New()
		NewHttpHandler(app, mockService)
		id := mockLimit.ConsumerID
		resp, err := app.Test(httptest.NewRequest("GET", "/"+id.String()+"/limit", nil))
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockService.On("GetLimit", mockLimit.ConsumerID).
			Return(nil, fiber.ErrInternalServerError).Once()
		app := fiber.New()
		NewHttpHandler(app, mockService)
		id := mockLimit.ConsumerID
		resp, err := app.Test(httptest.NewRequest("GET", "/"+id.String()+"/limit", nil))
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestHttpHandler_Store(t *testing.T) {
	var mockConsumerStoreRequest domain.ConsumerRequest
	err := faker.FakeData(&mockConsumerStoreRequest)
	assert.NoError(t, err)

	dateBirth, _ := time.Parse("2006-01-02", mockConsumerStoreRequest.DateOfBirth)
	mockConsumer := domain.Consumer{
		NIK:          mockConsumerStoreRequest.NIK,
		FullName:     mockConsumerStoreRequest.FullName,
		LegalName:    mockConsumerStoreRequest.LegalName,
		PlaceOfBirth: mockConsumerStoreRequest.PlaceOfBirth,
		DateOfBirth:  dateBirth,
		Salary:       mockConsumerStoreRequest.Salary,
		PhotoKTP:     mockConsumerStoreRequest.PhotoKTP,
		PhotoSelfie:  mockConsumerStoreRequest.PhotoSelfie,
	}

	mockService := new(mocks.ConsumerService)

	t.Run("success", func(t *testing.T) {
		mockService.On("Store", &mockConsumer).
			Return(nil).Once()
		app := fiber.New()

		NewHttpHandler(app, mockService)
		body, _ := json.Marshal(mockConsumerStoreRequest)
		fmt.Println(string(body))
		req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockService.On("Store", &mockConsumer).
			Return(fmt.Errorf("error")).Once()
		app := fiber.New()
		NewHttpHandler(app, mockService)
		body, _ := json.Marshal(mockConsumerStoreRequest)
		req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

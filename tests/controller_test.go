package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/FabricioCosati/onfly-test/internal/controllers"
	"github.com/FabricioCosati/onfly-test/internal/domain/dao"
	"github.com/FabricioCosati/onfly-test/internal/domain/dto"
	ce "github.com/FabricioCosati/onfly-test/internal/errors"
	cval "github.com/FabricioCosati/onfly-test/internal/validator"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrderService(ctx *gin.Context, orderRequest dto.OrderRequestPost) (dao.Order, ce.CustomError) {
	mockArgs := m.Called(ctx, orderRequest)
	return mockArgs.Get(0).(dao.Order), mockArgs.Get(1).(ce.CustomError)
}

func (m *MockOrderService) UpdateOrderStatus(ctx *gin.Context, order dto.OrderRequestPatch) (dao.Order, ce.CustomError) {
	mockArgs := m.Called(ctx, order)
	return mockArgs.Get(0).(dao.Order), mockArgs.Get(1).(ce.CustomError)
}

func (m *MockOrderService) GetOrderById(ctx *gin.Context) (dao.Order, ce.CustomError) {
	mockArgs := m.Called(ctx)
	return mockArgs.Get(0).(dao.Order), mockArgs.Get(1).(ce.CustomError)
}

func (m *MockOrderService) GetOrders(ctx *gin.Context, status string) (dao.OrderCollection, ce.CustomError) {
	mockArgs := m.Called(ctx)
	return mockArgs.Get(0).(dao.OrderCollection), mockArgs.Get(1).(ce.CustomError)
}

func CompareOrderRequest(order1, order2 dto.OrderRequestPost) bool {
	validRequester := order1.RequesterName == order2.RequesterName
	validDestination := order1.Destination == order2.Destination
	validGoingDate := order1.GoingDate.Time.Equal(order2.GoingDate.Time)
	validReturnDate := order1.ReturnDate.Time.Equal(order2.ReturnDate.Time)
	return validRequester && validDestination && validGoingDate && validReturnDate
}

func TestPostOrderController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockOrderService := new(MockOrderService)
	orderControllerImpl := controllers.OrderControllerInit(mockOrderService)

	engine := gin.New()
	api := engine.Group("/api")
	api.POST("/order-service", orderControllerImpl.CreateOrderService)

	t.Run("Success create order - 201", func(t *testing.T) {
		body := bytes.NewBufferString(`{
			"requester": "Teste Sucesso",
			"destination": "Buenos Aires, Argentina",
			"goingDate": "2024-11-15",
			"returnDate": "2024-12-15"
		}`)

		request := httptest.NewRequest("POST", "/api/order-service", body)
		request.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		goingTime, _ := time.Parse("2006-01-02", "2024-11-15")
		returnTime, _ := time.Parse("2006-01-02", "2024-12-15")
		goingDate := cval.Datetime{Time: goingTime}
		returnDate := cval.Datetime{Time: returnTime}

		mockedOrderRequest := dto.OrderRequestPost{
			RequesterName: "Teste Sucesso",
			Destination:   "Buenos Aires, Argentina",
			GoingDate:     goingDate,
			ReturnDate:    returnDate,
		}
		mockedOrderDao := dao.Order{
			OrderId:       1,
			RequesterName: "Teste Sucesso",
			Destination:   "Buenos Aires, Argentina",
			GoingDate:     goingTime,
			ReturnDate:    returnTime,
			Status:        "requested",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		mockOrderService.On(
			"CreateOrderService",
			mock.AnythingOfType("*gin.Context"),
			mock.MatchedBy(func(req dto.OrderRequestPost) bool {
				return CompareOrderRequest(req, mockedOrderRequest)
			})).
			Return(mockedOrderDao, ce.CustomError{}).
			Once()

		engine.ServeHTTP(recorder, request)

		var expectedOrder dao.Order
		err := json.Unmarshal(recorder.Body.Bytes(), &expectedOrder)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(t, 1, expectedOrder.OrderId)
		assert.Equal(t, "Teste Sucesso", expectedOrder.RequesterName)
		assert.Equal(t, "Buenos Aires, Argentina", expectedOrder.Destination)
		assert.Equal(t, "requested", expectedOrder.Status)
	})

	t.Run("Error on validate JSON - 422", func(t *testing.T) {
		body := bytes.NewBufferString(`{"test": "error"}`)

		request := httptest.NewRequest("POST", "/api/order-service", body)
		request.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		engine.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})
}

func TestPatchOrderController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockOrderService := new(MockOrderService)
	orderControllerImpl := controllers.OrderControllerInit(mockOrderService)

	engine := gin.New()
	api := engine.Group("/api")
	api.PATCH("/order-service/:id", orderControllerImpl.UpdateOrderStatus)

	t.Run("Success update order status - 200", func(t *testing.T) {
		body := bytes.NewBufferString(`{"status": "approved"}`)

		request := httptest.NewRequest("PATCH", "/api/order-service/1", body)
		request.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		mockedOrderPatch := dto.OrderRequestPatch{
			Status: "approved",
		}
		mockedOrderDao := dao.Order{
			Status: "approved",
		}

		mockOrderService.On(
			"UpdateOrderStatus",
			mock.AnythingOfType("*gin.Context"),
			mockedOrderPatch,
		).
			Return(mockedOrderDao, ce.CustomError{}).
			Once()

		engine.ServeHTTP(recorder, request)

		var expectedOrder dao.Order
		err := json.Unmarshal(recorder.Body.Bytes(), &expectedOrder)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "approved", expectedOrder.Status)
	})

	t.Run("Error on validate JSON - 422", func(t *testing.T) {
		body := bytes.NewBufferString(`{"status": "requested"}`)

		request := httptest.NewRequest("PATCH", "/api/order-service/1", body)
		request.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		engine.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})
}

func TestGetByIdOrderController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockOrderService := new(MockOrderService)
	orderControllerImpl := controllers.OrderControllerInit(mockOrderService)

	engine := gin.New()
	api := engine.Group("/api")
	api.GET("/order-service/:id", orderControllerImpl.GetOrderById)

	t.Run("Success on get order - 200", func(t *testing.T) {
		body := &bytes.Buffer{}
		request := httptest.NewRequest("GET", "/api/order-service/1", body)

		recorder := httptest.NewRecorder()

		mockedOrderDao := dao.Order{
			OrderId:       1,
			RequesterName: "Teste Sucesso",
			Destination:   "Buenos Aires, Argentina",
			Status:        "requested",
		}

		mockOrderService.On(
			"GetOrderById",
			mock.AnythingOfType("*gin.Context"),
		).
			Return(mockedOrderDao, ce.CustomError{}).
			Once()

		engine.ServeHTTP(recorder, request)

		var expectedOrder dao.Order
		err := json.Unmarshal(recorder.Body.Bytes(), &expectedOrder)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, 1, expectedOrder.OrderId)
		assert.Equal(t, "Teste Sucesso", expectedOrder.RequesterName)
		assert.Equal(t, "Buenos Aires, Argentina", expectedOrder.Destination)
		assert.Equal(t, "requested", expectedOrder.Status)
	})
}

func TestGetAllOrdersController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockOrderService := new(MockOrderService)
	orderControllerImpl := controllers.OrderControllerInit(mockOrderService)

	engine := gin.New()
	api := engine.Group("/api")
	api.GET("/order-services", orderControllerImpl.GetOrders)

	t.Run("Success on get all orders - 200", func(t *testing.T) {
		body := &bytes.Buffer{}
		request := httptest.NewRequest("GET", "/api/order-services", body)

		recorder := httptest.NewRecorder()

		mockedOrderDao := dao.OrderCollection{
			dao.Order{
				OrderId:       1,
				RequesterName: "Teste1",
				Destination:   "Buenos Aires, Argentina",
				Status:        "requested",
			},
			dao.Order{
				OrderId:       2,
				RequesterName: "Teste2",
				Destination:   "Tóquio, Japão",
				Status:        "approved",
			},
		}

		mockOrderService.On(
			"GetOrders",
			mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("string"),
		).
			Return(mockedOrderDao, ce.CustomError{}).
			Once()

		engine.ServeHTTP(recorder, request)

		var expectedOrder dao.OrderCollection
		err := json.Unmarshal(recorder.Body.Bytes(), &expectedOrder)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, 1, expectedOrder[0].OrderId)
		assert.Equal(t, 2, expectedOrder[1].OrderId)
		assert.Equal(t, "Teste1", expectedOrder[0].RequesterName)
		assert.Equal(t, "Teste2", expectedOrder[1].RequesterName)
		assert.Equal(t, "Buenos Aires, Argentina", expectedOrder[0].Destination)
		assert.Equal(t, "Tóquio, Japão", expectedOrder[1].Destination)
		assert.Equal(t, "requested", expectedOrder[0].Status)
		assert.Equal(t, "approved", expectedOrder[1].Status)
	})
}

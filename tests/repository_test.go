package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/FabricioCosati/onfly-test/internal/domain/dao"
	"github.com/FabricioCosati/onfly-test/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDatabaseMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error on create mock: %s", err)
	}

	mysqlConf := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gorm, err := gorm.Open(mysqlConf, &gorm.Config{})
	if err != nil {
		t.Fatalf("error on initialize GORM: %s", err)
	}

	return gorm, mock
}

func TestPostOrder(t *testing.T) {
	db, mock := getDatabaseMock(t)
	conn, err := db.DB()
	if err != nil {
		t.Fatalf("error on getting database connection")
	}
	defer conn.Close()

	goingDate, _ := time.Parse("2006-01-02", "2024-12-15")
	returnDate, _ := time.Parse("2006-01-02", "2025-01-24")
	expectedOrder := dao.Order{
		RequesterName: "Light Yagami",
		Destination:   "Tóquio, Japão",
		GoingDate:     goingDate,
		ReturnDate:    returnDate,
		Status:        "requested",
	}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `orders`*").
		WithArgs(
			"Light Yagami",
			"Tóquio, Japão",
			goingDate,
			returnDate,
			"requested",
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	orderRep := repository.OrderRepositoryInit(db)
	orderResponse, err := orderRep.CreateOrder(&expectedOrder)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, orderResponse)
}

func TestPatchOrder(t *testing.T) {
	db, mock := getDatabaseMock(t)
	conn, err := db.DB()
	if err != nil {
		t.Fatalf("error on getting database connection")
	}
	defer conn.Close()

	expectedOrder := dao.Order{
		Status: "approved",
	}

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `orders`.*").
		WithArgs(
			"approved",
			sqlmock.AnyArg(),
			25,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectQuery("^SELECT \\* FROM `orders`*").
		WithArgs(
			25,
			1,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id_order"}).AddRow(1))

	orderRep := repository.OrderRepositoryInit(db)
	orderResponse, err := orderRep.UpdateStatusOrder(&expectedOrder, 25)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, orderResponse)
}

func TestGetByIdOrder(t *testing.T) {
	db, mock := getDatabaseMock(t)
	conn, err := db.DB()
	if err != nil {
		t.Fatalf("error on getting database connection")
	}
	defer conn.Close()

	orderId := 1
	expectedOrder := dao.Order{
		OrderId: 1,
	}
	mock.ExpectQuery("^SELECT \\* FROM `orders`*").
		WithArgs(
			orderId,
			1,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id_order"}).AddRow(orderId))

	orderRep := repository.OrderRepositoryInit(db)
	orderResponse, err := orderRep.GetOrderById(orderId)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, orderResponse)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

func TestGetAllOrders(t *testing.T) {
	db, mock := getDatabaseMock(t)
	conn, err := db.DB()
	if err != nil {
		t.Fatalf("error on getting database connection")
	}
	defer conn.Close()

	expectedOrders := dao.OrderCollection{
		dao.Order{OrderId: 25},
		dao.Order{OrderId: 26},
	}
	mock.ExpectQuery("^SELECT \\* FROM `orders`*").
		WillReturnRows(sqlmock.NewRows([]string{"id_order"}).
			AddRow(25).
			AddRow(26),
		)

	orderRep := repository.OrderRepositoryInit(db)
	ordersResponse, err := orderRep.GetOrders()

	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, ordersResponse)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	repository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
	"gitlab.ozon.dev/daker255/homework-8/tests/fixtures"
)

func TestCreateOrder(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("Product").
			Quantity(1).
			UserID(1).
			ID(1).
			Status("approved").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		actualOrderID, err := orderRepo.CreateOrder(context.Background(), 1, "Product", 1)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, order.ID, actualOrderID)
	})
	t.Run("no user ID", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("Product").
			Quantity(1).
			ID(1).
			Status("approved").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		actualOrderID, err := orderRepo.CreateOrder(context.Background(), 0, "Product", 1)

		//assert
		assert.Error(t, err)
		assert.NotEqual(t, order.ID, actualOrderID)
	})
	t.Run("empty order productName", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("Product").
			Quantity(1).
			UserID(1).
			ID(1).
			Status("approved").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		actualOrderID, _ := orderRepo.CreateOrder(context.Background(), 1, "", 1)
		//assert
		assert.NotEqual(t, order.ID, actualOrderID)
	})
}

func TestGetOrderByID(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("Product").
			Quantity(1).
			UserID(1).
			ID(1).
			OrderDate(time.Now().UTC()).
			Status("created").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		orderID, _ := orderRepo.CreateOrder(context.Background(), user.ID, order.ProductName, order.Quantity)
		actualOrder, err := orderRepo.GetByID(context.Background(), orderID)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, order.ID, actualOrder.ID)
		assert.Equal(t, order.UserID, actualOrder.UserID)
		assert.Equal(t, order.ProductName, actualOrder.ProductName)
		assert.Equal(t, order.Status, actualOrder.Status)
		assert.Equal(t, order.Quantity, actualOrder.Quantity)
	})
	t.Run("no order with this ID", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("").
			Quantity(0).
			UserID(0).
			ID(0).
			OrderDate(time.Now().UTC()).
			Status("").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		orderID, _ := orderRepo.CreateOrder(context.Background(), user.ID, order.ProductName, order.Quantity)
		actualOrder, err := orderRepo.GetByID(context.Background(), orderID)
		//assert
		assert.Error(t, err)
		assert.Equal(t, order.ID, actualOrder.ID)
		assert.Equal(t, order.UserID, actualOrder.UserID)
		assert.Equal(t, order.ProductName, actualOrder.ProductName)
		assert.Equal(t, order.Status, actualOrder.Status)
		assert.Equal(t, order.Quantity, actualOrder.Quantity)
	})
}

func TestGetAllOrders(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {

		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order1 := fixtures.Order().
			ProductName("Product1").
			Quantity(1).
			UserID(1).
			ID(1).
			OrderDate(time.Now()).
			Status("created").P()

		order2 := fixtures.Order().
			ProductName("Product2").
			Quantity(1).
			UserID(1).
			ID(2).
			OrderDate(time.Now().UTC()).
			Status("created").P()

		orderSlice := make([]*models.Order, 0)
		orders := append(orderSlice, order1, order2)
		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)

		_, _ = orderRepo.CreateOrder(context.Background(), user.ID, order1.ProductName, order1.Quantity)
		_, _ = orderRepo.CreateOrder(context.Background(), user.ID, order2.ProductName, order2.Quantity)

		actualOrders, err := orderRepo.GetAll(context.Background())

		//assert
		assert.NoError(t, err)
		// check 1st order
		assert.Equal(t, actualOrders[0].ID, orders[0].ID)
		assert.Equal(t, actualOrders[0].ProductName, orders[0].ProductName)
		assert.Equal(t, actualOrders[0].Quantity, orders[0].Quantity)
		assert.Equal(t, actualOrders[0].Status, orders[0].Status)
		// check 1st order
		assert.Equal(t, actualOrders[1].ID, orders[1].ID)
		assert.Equal(t, actualOrders[1].ProductName, orders[1].ProductName)
		assert.Equal(t, actualOrders[1].Quantity, orders[1].Quantity)
		assert.Equal(t, actualOrders[1].Status, orders[1].Status)
	})
	t.Run("no entities in database", func(t *testing.T) {

		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order1 := fixtures.Order().
			ProductName("Product1").
			Quantity(1).
			UserID(1).
			ID(1).
			OrderDate(time.Now()).
			Status("created").P()

		order2 := fixtures.Order().
			ProductName("Product2").
			Quantity(1).
			UserID(1).
			ID(2).
			OrderDate(time.Now().UTC()).
			Status("created").P()

		orderSlice := make([]*models.Order, 0)
		_ = append(orderSlice, order1, order2)
		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)

		actualOrders, err := orderRepo.GetAll(context.Background())

		//assert
		assert.NoError(t, err)
		// check 1st order
		assert.Equal(t, len(orderSlice), len(actualOrders))
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("Product").
			Quantity(1).
			UserID(1).
			ID(1).
			OrderDate(time.Now().UTC()).
			Status("created").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		orderID, _ := orderRepo.CreateOrder(context.Background(), user.ID, order.ProductName, order.Quantity)
		actualOrder, _ := orderRepo.GetByID(context.Background(), orderID)

		isUpdated, err := orderRepo.UpdateOrderStatus(context.Background(), actualOrder.ID, "in progress")

		//assert
		assert.NoError(t, err)
		assert.Equal(t, true, isUpdated)
	})
	t.Run("no order with specific id found", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("Product").
			Quantity(1).
			UserID(1).
			ID(1).
			OrderDate(time.Now().UTC()).
			Status("created").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		orderID, _ := orderRepo.CreateOrder(context.Background(), user.ID, order.ProductName, order.Quantity)
		_, _ = orderRepo.GetByID(context.Background(), orderID)

		isUpdated, err := orderRepo.UpdateOrderStatus(context.Background(), 2, "in progress")

		//assert
		assert.NoError(t, err)
		assert.Equal(t, false, isUpdated)
	})
}

func TestDeleteOrder(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("Product").
			Quantity(1).
			UserID(1).
			ID(1).
			OrderDate(time.Now().UTC()).
			Status("created").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		orderID, _ := orderRepo.CreateOrder(context.Background(), user.ID, order.ProductName, order.Quantity)

		isDeleted, err := orderRepo.DeleteOrder(context.Background(), orderID)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, true, isDeleted)
	})
	t.Run("order with this ID not found and not deleted", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		orderRepo := repository.NewPostgresqlOrderRepo(DbInstance.DB)

		user := fixtures.User().Valid().P()

		order := fixtures.Order().
			ProductName("Product").
			Quantity(1).
			UserID(1).
			ID(1).
			OrderDate(time.Now().UTC()).
			Status("created").P()

		//act
		_, _ = userRepo.CreateUser(context.Background(), user.Username, user.Email)
		_, _ = orderRepo.CreateOrder(context.Background(), user.ID, order.ProductName, order.Quantity)

		isDeleted, err := orderRepo.DeleteOrder(context.Background(), 2)

		//assert
		assert.Error(t, err)
		assert.Equal(t, false, isDeleted)
	})
}

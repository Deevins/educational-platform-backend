package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	models_dto "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/models"
	repository "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/postgresql"
	"gitlab.ozon.dev/daker255/homework-8/tests/fixtures"
	_ "gitlab.ozon.dev/daker255/homework-8/tests/pgmock"
)

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)

		user := fixtures.User().Username("Deevins").Email("div@div.com").P()

		res, err := userRepo.CreateUser(context.Background(), user.Username, user.Email)

		assert.NoError(t, err)
		assert.Equal(t, res, fixtures.User().ID(1).V().ID)
	})
	t.Run("empty email error", func(t *testing.T) {
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)

		user := fixtures.User().Username("Deevins").Email("div@div.com").P()

		res, err := userRepo.CreateUser(context.Background(), user.Username, "")

		assert.Error(t, err)
		assert.Equal(t, res, models.UserID(0))
	})
	t.Run("empty username error", func(t *testing.T) {
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()
		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)

		user := fixtures.User().Username("Deevins").Email("div@div.com").P()

		res, err := userRepo.CreateUser(context.Background(), "", user.Email)

		assert.Error(t, err)
		assert.Equal(t, res, models.UserID(0))
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		ctx := context.Background()

		mockedUser := fixtures.User().Username("Deevins").Email("div@div.com").ID(1).P()

		createdUserID, _ := userRepo.CreateUser(context.Background(), mockedUser.Username, mockedUser.Email)

		//act
		result, err := userRepo.GetByID(ctx, createdUserID)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, mockedUser, result)
	})
	t.Run("user by this ID not found", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		ctx := context.Background()

		mockedUser := fixtures.User().Username("Deevins").Email("div@div.com").ID(2).P()

		createdUserID, _ := userRepo.CreateUser(context.Background(), mockedUser.Username, mockedUser.Email)

		//act
		result, err := userRepo.GetByID(ctx, createdUserID)

		//assert
		assert.NoError(t, err)
		assert.NotEqual(t, mockedUser, result)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)

		userSlice := make([]*models_dto.UserDTO, 0)

		user1 := fixtures.User().Username("Deevins1").Email("div@div.com").P()
		//user2 := fixtures.User().Username("Deevins2").Email("div@div.com").P()

		//Create mock users
		createdUserID1, _ := userRepo.CreateUser(context.Background(), user1.Username, user1.Email)
		//createdUserID2, _ := userRepo.CreateUser(context.Background(), user2.Username, user2.Email)

		//Get mock users
		fetchedMockUser1, _ := userRepo.GetByID(context.Background(), createdUserID1)
		//fetchedMockUser2, _ := userRepo.GetByID(context.Background(), createdUserID2)

		fmt.Println(fetchedMockUser1)
		//fmt.Println(fetchedMockUser2)

		mockUsers := append(userSlice, fetchedMockUser1)
		//act
		result, err := userRepo.GetAll(context.Background())

		//assert
		assert.NoError(t, err)
		assert.Equal(t, mockUsers, result)
	})
	t.Run("no users found in DB", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)

		expectedUserSlice := make([]*models_dto.UserDTO, 0)

		user1 := fixtures.User().Username("Deevins1").Email("div@div.com").P()
		//user2 := fixtures.User().Username("Deevins2").Email("div@div.com").P()

		//Create mock users
		//createdUserID1, _ := userRepo.CreateUser(context.Background(), user1.Username, user1.Email)
		//createdUserID2, _ := userRepo.CreateUser(context.Background(), user2.Username, user2.Email)

		//Get mock users
		//fetchedMockUser1, _ := userRepo.GetByID(context.Background(), createdUserID1)
		//fetchedMockUser2, _ := userRepo.GetByID(context.Background(), createdUserID2)

		mockUsers := append(expectedUserSlice, user1)
		//act
		result, err := userRepo.GetAll(context.Background())

		//assert
		assert.NoError(t, err)
		assert.NotEqual(t, mockUsers, result)
	})
}

func TestUpdateEmail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		ctx := context.Background()

		mockedUser := fixtures.User().Username("Deevins").Email("div@div.com").ID(1).V()

		createdUserID, _ := userRepo.CreateUser(context.Background(), mockedUser.Username, mockedUser.Email)

		//act

		isChanged, err := userRepo.UpdateEmail(ctx, createdUserID, "div2@div.com")

		//assert
		assert.NoError(t, err)
		assert.Equal(t, true, isChanged)
	})
	t.Run("user not found and email not updated", func(t *testing.T) {

		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		ctx := context.Background()

		mockedUser := fixtures.User().Username("Deevins").Email("div@div.com").ID(1).V()

		_, _ = userRepo.CreateUser(context.Background(), mockedUser.Username, mockedUser.Email)

		//act
		isChanged, err := userRepo.UpdateEmail(ctx, 2, "div2@div.com")

		//assert
		assert.Error(t, err)
		assert.Equal(t, false, isChanged)
	})
}

func TestUpdateUsername(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		ctx := context.Background()

		mockedUser := fixtures.User().Username("Deevins").Email("div@div.com").ID(1).V()

		createdUserID, _ := userRepo.CreateUser(context.Background(), mockedUser.Username, mockedUser.Email)

		//act

		isChanged, err := userRepo.UpdateEmail(ctx, createdUserID, "div2@div.com")

		//assert
		assert.NoError(t, err)
		assert.Equal(t, true, isChanged)
	})
	t.Run("user not found and username not updated", func(t *testing.T) {

		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		ctx := context.Background()

		mockedUser := fixtures.User().Username("Deevins").Email("div@div.com").ID(1).V()

		_, _ = userRepo.CreateUser(context.Background(), mockedUser.Username, mockedUser.Email)

		//act
		isChanged, err := userRepo.UpdateUsername(ctx, 2, "Deevins2011")

		//assert
		assert.Error(t, err)
		assert.Equal(t, false, isChanged)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		ctx := context.Background()

		mockedUser := fixtures.User().Username("Deevins").Email("div@div.com").ID(1).P()

		_, _ = userRepo.CreateUser(context.Background(), mockedUser.Username, mockedUser.Email)
		//act
		isDeleted, err := userRepo.DeleteUser(ctx, 1)
		//fmt.Println(deletedCount, err)
		//assert
		assert.NoError(t, err)
		assert.Equal(t, true, isDeleted)
	})
	t.Run("user by this ID not found and not deleted", func(t *testing.T) {
		//arrange
		DbInstance.SetUp(t)
		defer DbInstance.TearDown()

		userRepo := repository.NewPostgresqlUserRepo(DbInstance.DB)
		ctx := context.Background()

		mockedUser := fixtures.User().Username("Deevins").Email("div@div.com").ID(1).P()

		_, _ = userRepo.CreateUser(context.Background(), mockedUser.Username, mockedUser.Email)
		//act
		isDeleted, err := userRepo.DeleteUser(ctx, 23)
		//fmt.Println(deletedCount, err)
		//assert
		assert.NoError(t, err)
		assert.Equal(t, false, isDeleted)
	})
}

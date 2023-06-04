package sql

import (
	"context"
	"errors"
	"meli-product-miner/app/supplier"
	"meli-product-miner/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestReaderMySQL_GetSupplier(t *testing.T) {
	ctx := context.Background()
	mock, gormDB := mockDB(t)

	reader := NewReaderMySQL(gormDB)

	t.Run("should return success when get supplier", func(t *testing.T) {
		expectedSupplier := createValidSupplier()
		expectedSupplierID := expectedSupplier.ID

		mockQueryToGet(mock, expectedSupplierID, expectedSupplier)

		supplier, err := reader.GetSupplier(ctx, expectedSupplierID)
		assert.NoError(t, err)
		AssertGetSupplier(t, expectedSupplier, supplier)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("should return error when get supplier", func(t *testing.T) {
		expectedSupplier := createValidSupplier()
		expectedSupplierID := expectedSupplier.ID

		mockQueryToGet(mock, expectedSupplierID, expectedSupplier).WillReturnError(errors.New("error to get supplier"))

		supplier, err := reader.GetSupplier(ctx, expectedSupplierID)
		assert.Error(t, err)
		assert.Nil(t, supplier)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("should return error when supplier not found", func(t *testing.T) {
		expectedSupplierID := uint64(1)

		mockQueryToGet(mock, expectedSupplierID, nil)

		supplier, err := reader.GetSupplier(ctx, expectedSupplierID)
		assert.NoError(t, err)
		assert.Nil(t, supplier)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

}

func mockQueryToGet(mock sqlmock.Sqlmock, expectedSupplierID uint64, expectedSupplier *supplier.SupplierEntity) *sqlmock.ExpectedQuery {
	query := "SELECT (.+) FROM `suppliers` WHERE id = ?"

	if expectedSupplier == nil {
		return mock.ExpectQuery(query).WithArgs(expectedSupplierID).WillReturnError(gorm.ErrRecordNotFound)
	}

	return mock.ExpectQuery("SELECT (.+) FROM `suppliers`").WithArgs(expectedSupplierID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "Name", "Email", "Phone", "ContactName",
			"ContactEmail", "ContactPhone", "SiteURL"}).AddRow(expectedSupplierID, expectedSupplier.Name,
			expectedSupplier.Email, expectedSupplier.Phone, expectedSupplier.ContactName,
			expectedSupplier.ContactEmail, expectedSupplier.ContactPhone, expectedSupplier.SiteURL))
}

func mockDB(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error to create mock: %s", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("error to create gormDB: %s", err)
	}
	return mock, gormDB
}

func AssertGetSupplier(t *testing.T, expectedSupplier *supplier.SupplierEntity, supplier *supplier.SupplierEntity) {
	assert.Equal(t, expectedSupplier.ID, supplier.ID)
	assert.Equal(t, expectedSupplier.Name, supplier.Name)
	assert.Equal(t, expectedSupplier.Email, supplier.Email)
	assert.Equal(t, expectedSupplier.Phone, supplier.Phone)
	assert.Equal(t, expectedSupplier.ContactName, supplier.ContactName)
	assert.Equal(t, expectedSupplier.ContactEmail, supplier.ContactEmail)
	assert.Equal(t, expectedSupplier.ContactPhone, supplier.ContactPhone)
	assert.Equal(t, expectedSupplier.SiteURL, supplier.SiteURL)
}

func createValidSupplier() *supplier.SupplierEntity {
	return &supplier.SupplierEntity{
		ID:           mock.RandomUint64(),
		Name:         mock.RandomString(),
		Email:        mock.RandomString(),
		Phone:        mock.RandomString(),
		ContactName:  mock.RandomString(),
		ContactEmail: mock.RandomString(),
		ContactPhone: mock.RandomString(),
		SiteURL:      mock.RandomString(),
	}
}

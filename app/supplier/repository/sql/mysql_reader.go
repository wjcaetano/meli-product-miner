package sql

import (
	"context"
	"errors"
	entity "meli-product-miner/app/supplier"

	"gorm.io/gorm"
)

type (
	ReaderMySQL struct {
		db *gorm.DB
	}
)

func NewReaderMySQL(db *gorm.DB) ReaderMySQL {
	return ReaderMySQL{db: db}
}

func (r ReaderMySQL) GetSupplier(ctx context.Context, supplierID uint64) (result *entity.SupplierEntity, err error) {
	var supplier *supplierDTO
	tx := r.db.
		WithContext(ctx).
		Where("id = ?", supplierID)

	err = tx.First(&supplier).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, nil
		}
		return nil, err
	}

	if supplier != nil {
		convertedDomain := toSupplierDomain(*supplier)
		result = &convertedDomain
	}
	return result, nil
}

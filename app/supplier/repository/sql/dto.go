package sql

import (
	entity "meli-product-miner/app/supplier"
)

type (
	supplierDTO struct {
		ID           uint64 `gorm:"primaryKey"`
		Name         string `gorm:"column:name not null"`
		Email        string `gorm:"column:email not null"`
		Phone        string `gorm:"column:phone not null"`
		ContactName  string `gorm:"column:contact_name not null"`
		ContactEmail string `gorm:"column:contact_email not null"`
		ContactPhone string `gorm:"column:contact_phone not null"`
		SiteURL      string `gorm:"column:site_url not null"`
	}

	productDTO struct {
		ID          uint64  `gorm:"primaryKey"`
		SupplierID  uint64  `gorm:"column:supplier_id not null"`
		SupplierSKU string  `gorm:"column:supplier_sku not null"`
		Name        string  `gorm:"column:name not null"`
		Description string  `gorm:"column:description not null"`
		Price       float64 `gorm:"column:price not null"`
		Stock       int     `gorm:"column:stock not null"`
	}
)

func (supplierDTO) TableName() string {
	return "suppliers"
}

func toSupplier(supplier entity.SupplierEntity) supplierDTO {
	return supplierDTO{
		ID:           supplier.ID,
		Name:         supplier.Name,
		Email:        supplier.Email,
		Phone:        supplier.Phone,
		ContactName:  supplier.ContactName,
		ContactEmail: supplier.ContactEmail,
		ContactPhone: supplier.ContactPhone,
		SiteURL:      supplier.SiteURL,
	}
}

func toProduct(product entity.ProductEntity) productDTO {
	return productDTO{
		ID:          product.ID,
		SupplierID:  product.SupplierID,
		SupplierSKU: product.SupplierSKU,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func toSupplierDomain(supplier supplierDTO) entity.SupplierEntity {
	return entity.SupplierEntity{
		ID:           supplier.ID,
		Name:         supplier.Name,
		Email:        supplier.Email,
		Phone:        supplier.Phone,
		ContactName:  supplier.ContactName,
		ContactEmail: supplier.ContactEmail,
		ContactPhone: supplier.ContactPhone,
		SiteURL:      supplier.SiteURL,
	}
}

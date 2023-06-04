package supplier

type (
	SupplierEntity struct {
		ID           uint64
		Name         string
		Email        string
		Phone        string
		ContactName  string
		ContactEmail string
		ContactPhone string
		SiteURL      string
	}

	ProductEntity struct {
		ID          uint64
		SupplierID  uint64
		SupplierSKU string
		Name        string
		Description string
		Price       float64
		Stock       int
	}
)

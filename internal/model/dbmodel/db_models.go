package dbmodel

import "time"

type Customer struct {
	CustomerID      string `gorm:"primaryKey" json:"customer_id"`
	CustomerName    string `gorm:"size:255;not null" json:"customer_name"`
	CustomerEmail   string `gorm:"size:255;not null;unique" json:"customer_email"`
	CustomerAddress string `gorm:"type:text" json:"customer_address"`
}

type Region struct {
	RegionID   int    `gorm:"primaryKey;autoIncrement" json:"region_id"`
	RegionName string `gorm:"size:255;not null" json:"region_name"`
}

type Category struct {
	CategoryID   int    `gorm:"unique;primaryKey" json:"category_id"`
	CategoryName string `gorm:"size:255;not null" json:"category_name"`
}

type Product struct {
	ProductID   string   `gorm:"primaryKey;size:255;not null" json:"product_id"`
	ProductName string   `gorm:"size:255;not null" json:"product_name"`
	Category    Category `gorm:"foreignKey:categoryid"`
	Categoryid  int      `JSON:"categoryid, omitempty"`
	UnitPrice   float64  `gorm:"type:decimal(10,2)" json:"unit_price"`
}

type Order struct {
	OrderID       int       `gorm:"primaryKey" json:"order_id"`
	Customer      Customer  `gorm:"foreignKey:customersid"`
	Customersid   string    `JSON:"customersid, omitempty"`
	OrderDate     time.Time `gorm:"not null" json:"order_date"`
	TotalAmount   float64   `gorm:"not null" json:"total_amount"`
	ShippingCost  float64   `gorm:"not null" json:"shipping_cost"`
	PaymentMethod string    `gorm:"size:255" json:"payment_method"`
	Discount      float64   `gorm:"type:decimal(10,2)" json:"discount"`
	Region        Region    `gorm:"foreignKey:regionid"`
	Regionid      int       `JSON:"regionid, omitempty"`
}

type OrderItem struct {
	OrderItemId uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Order       Order   `gorm:"foreignKey:ordersid"`
	Ordersid    int     `JSON:"ordersid, omitempty"`
	Product     Product `gorm:"foreignKey:productsid"`
	Productsid  string  `JSON:"productsid, omitempty"`
	Quantity    int     `gorm:"not null" json:"quantity"`
	UnitPrice   float64 `gorm:"type:decimal(10,2)" json:"unit_price"`
}

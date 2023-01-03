package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Products []Product

type PriceInTime struct {
	Timestamp          primitive.DateTime `json:"timestamp" bson:"timestamp" validate:"required"`
	IsOnPromotion      bool               `json:"is-on-promotion" bson:"is-on-promotion" validate:"required"`
	Price              float64            `json:"price" bson:"price" validate:"required"`
	PricePerUnit       string             `json:"price-per-unit" bson:"price-per-unit" validate:"required"`
	RegularPrice       float64            `json:"regular-price" bson:"regular-price" validate:"required"`
	PricePerUnitNumber float64            `json:"price-per-unit-number" bson:"price-per-unit-number" validate:"required"`
	BestPrice          float64            `json:"best-price" bson:"best-price" validate:"required"`
	StockStatus        string             `json:"stock-status" bson:"stock-status" validate:"required"`
	IsNew              bool               `json:"is-new" bson:"is-new" validate:"required"`
}

type Product struct {
	Id   primitive.ObjectID `json:"id" bson:"_id" validate:"required"`
	Name string             `json:"name" bson:"name" validate:"required"`

	CategoryNames   []string `json:"category-names" bson:"category-names" validate:"required"`
	CategoryName    string   `json:"category-name" bson:"category-name" validate:"required"`
	AllergensFilter []string `json:"allergens-filter" bson:"allergens-filter" validate:"required"`

	SalesUnit           string             `json:"sales-unit" bson:"sales-unit" validate:"required"`
	Title               string             `json:"title" bson:"title" validate:"required"`
	CodeInternal        uint64             `json:"code-internal" bson:"code-internal" validate:"required"`
	CreatedAt           primitive.DateTime `json:"created-at" bson:"created-at" validate:"required"`
	ImageURL            string             `json:"image-url" bson:"image-url" validate:"required"`
	ApproxWeightProduct bool               `json:"approx-weight-product" bson:"approx-weight-product" validate:"required"`
	URL                 string             `json:"url" bson:"url" validate:"required"`
	Brand               string             `json:"brand" bson:"brand" validate:"required"`

	PriceInTime []PriceInTime `json:"price-in-time" bson:"price-in-time" validate:"required"`
}

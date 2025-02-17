package domain

// CarBrand enum (sample values)
type CarBrand string

const (
	Toyota CarBrand = "Toyota"
	Honda  CarBrand = "Honda"
	Ford   CarBrand = "Ford"
)

// CarModel enum (sample values)
type CarModel string

const (
	Corolla CarModel = "Corolla"
	Civic   CarModel = "Civic"
	Mustang CarModel = "Mustang"
)

// Transmission enum
type Transmission string

const (
	Automatic Transmission = "Automatic"
	Manual    Transmission = "Manual"
)

// FuelType enum
type FuelType string

const (
	Petrol   FuelType = "Petrol"
	Diesel   FuelType = "Diesel"
	Electric FuelType = "Electric"
)

type Car struct {
	ID           int64        `db:"id" json:"id"`
	Brand        CarBrand     `db:"brand" json:"brand"`
	Model        CarModel     `db:"model" json:"model"`
	Transmission Transmission `db:"transmission" json:"transmission"`
	FuelType     FuelType     `db:"fuel_type" json:"fuel_type"`
	PricePerDay  float64      `db:"price_per_day" json:"price_per_day"`
	Available    bool         `db:"available" json:"available"`
	Photo        string       `db:"photo" json:"photo"`
}

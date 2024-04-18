package domain

type Equipment struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `json:"name"`
	Availability bool    `json:"availability"`
	RentalCosts  float64 `json:"rentalCosts"`
	Category     string  `json:"category"`
}

package domain

import "time"

type Order struct {
	OrderID     int       `json:"orderID"`
	UserID      int       `json:"userID"`
	EquipmentID int       `json:"equipmentID"`
	RentalDate  time.Time `json:"rentalDate"`
	ReturnDate  time.Time `json:"returnDate,omitempty"`
}

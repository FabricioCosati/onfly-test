package dto

import cval "github.com/FabricioCosati/onfly-test/internal/validator"

type OrderStatus string

const (
	REQUESTED OrderStatus = "requested"
	APPROVED  OrderStatus = "approved"
	CANCELED  OrderStatus = "canceled"
)

type OrderRequestPost struct {
	RequesterName string        `json:"requester" validate:"required"`
	Destination   string        `json:"destination" validate:"required"`
	GoingDate     cval.Datetime `json:"goingDate" validate:"dateRequired"`
	ReturnDate    cval.Datetime `json:"returnDate" validate:"dateRequired,gteDate=GoingDate"`
}

type OrderRequestPatch struct {
	Status string `json:"status" validate:"required,oneof=approved canceled"`
}

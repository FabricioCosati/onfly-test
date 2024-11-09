package dto

import cval "github.com/FabricioCosati/onfly-test/internal/validator"

type OrderStatus string

const (
	REQUESTED OrderStatus = "requested"
	APPROVED  OrderStatus = "approved"
	CANCELED  OrderStatus = "canceled"
)

type OrderRequest struct {
	RequesterName string        `json:"requester" validate:"required"`
	Destination   string        `json:"destination" validate:"required"`
	GoingDate     cval.Datetime `json:"goingDate" validate:"required"`
	ReturnDate    cval.Datetime `json:"returnDate" validate:"required,gteDate=GoingDate"`
	Status        string        `json:"status" validate:"required,oneof=requested approved canceled"`
}
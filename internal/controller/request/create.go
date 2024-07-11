package request

import (
	"test-task-go/internal/model"
	"time"
)

type CreateService struct {
	Name        string `json:"name" binding:"required,max=255"`
	Type        string `json:"type" binding:"required,max=255,type_validation"`
	PaymentType string `json:"paymentType" binding:"required,max=255"`
	Price       uint64 `json:"price" binding:"required,min=0"`
}

func (r CreateService) ToModel() model.Service {
	return model.Service{
		Id:          0,
		CreatedAt:   time.Time{},
		Name:        r.Name,
		Type:        r.Type,
		PaymentType: r.PaymentType,
		Price:       r.Price,
	}
}

var ServiceTypes = [...]string{"VDS", "Dedicated_Server", "Hosting"}
var ServicePaymentTypes = [...]string{"year", "half-year", "year"}

package dto

type CreateTransaksi struct {
	Customer		uint32		`json:"customer" form:"customer"`
	TotalPrice		uint64		`json:"total_price" form:"total_price"`
	PaymentMethods	string		`json:"payment_methods" form:"payment_methods"`
	Status			string		`json:"status" form:"status"`
}

type UpdateTransaksi struct {
	Customer		uint32		`json:"customer" form:"customer"`
	TotalPrice		uint64		`json:"total_price" form:"total_price"`
	PaymentMethods	string		`json:"payment_methods" form:"payment_methods"`
	Status			string		`json:"status" form:"status"`
	IsReview		string		`json:"is_review" form:"is_review"`
}

package cart

import (
	"inventory/ent/entgen"
	"inventory/internal/inventory"
)

type RQCart struct {
	ProductId string `json:"productId" validate:"required"`
	UserID    string `json:"-"`
	Status    string `json:"status"`
}

func (cart *RQCart) MapFromPayment(payment *entgen.TblPayment) {
	cart.ProductId = payment.InventoryId
	cart.UserID = payment.UserId
	cart.Status = payment.Status
}

type Cart struct {
	Id        string
	ProductId inventory.Inventory
	UserID    string
	Status    string
	ImageUrl  string
}

func (cart *Cart) MapFrom(cartEnt *entgen.TblCart) {
	cart.Id = cartEnt.ID
	var productId inventory.Inventory

	productId.Name = cartEnt.Edges.Inventory.Name
	productId.Description = cartEnt.Edges.Inventory.Description
	productId.Price = cartEnt.Edges.Inventory.Price
	productId.Id = cartEnt.Edges.Inventory.ID
	productId.IsActive = cartEnt.Edges.Inventory.IsActive
	//productId.Tags = cartEnt.Edges.Inventory.Tags

	cart.ProductId = productId
	cart.UserID = cartEnt.UserId
	cart.Status = cartEnt.Status.String()
	cart.ImageUrl = cartEnt.Edges.Inventory.ImageLink

}

type RSCart struct {
	Id        string                `json:"id"`
	ProductId inventory.RSInventory `json:"productId"`
	UserID    string                `json:"userId"`
	Status    string                `json:"status"`
	ImageUrl  string                `json:"imageUrl"`
}

func (rs *RSCart) MapFrom(cart *Cart) {
	rs.Id = cart.Id
	var productId inventory.RSInventory

	productId.Name = cart.ProductId.Name
	productId.Description = cart.ProductId.Description
	productId.Price = cart.ProductId.Price
	productId.Id = cart.ProductId.Id
	productId.Tags = cart.ProductId.Tags

	rs.ProductId = productId
	rs.UserID = cart.UserID
	rs.Status = cart.Status
	rs.ImageUrl = cart.ImageUrl
}

type Payment struct {
	Id              string
	ProductId       string
	UserID          string
	RazorPayOrderId string
	OrderId         string
	Amount          float32
	Status          string
}

func (payment *Payment) MapFrom(entPayment *entgen.TblPayment) {
	payment.Id = entPayment.ID
	payment.UserID = entPayment.UserId
	payment.Amount = entPayment.Amount
	payment.Status = entPayment.Status
	payment.Status = entPayment.Status
	payment.ProductId = entPayment.InventoryId
	payment.RazorPayOrderId = entPayment.RazorpayOrderId
}

type RSPayment struct {
	Id              string  `json:"id"`
	ProductId       string  `json:"productId"`
	RazorPayOrderId string  `json:"razorPayOrderId"`
	Amount          float32 `json:"amount"`
	Status          string  `json:"status"`
}

func (rs *RSPayment) MapFrom(payment *Payment) {
	rs.Id = payment.Id
	//rs.UserID=entPayment.UserId
	rs.Amount = payment.Amount
	rs.Status = payment.Status
	rs.RazorPayOrderId = payment.RazorPayOrderId
	rs.ProductId = payment.ProductId
}

type RQPayment struct {
	RazorpayOrderID   string `json:"razorpay_order_id"`
	RazorpayPaymentID string `json:"razorpay_payment_id"`
	// RazorpaySignature string `json:"razorpay_signature"` // optional for now
}

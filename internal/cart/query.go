package cart

import (
	"fmt"
	"inventory/ent/entgen"
	"inventory/ent/entgen/tblcart"
	"inventory/ent/entgen/tblinventory"
	"inventory/ent/entgen/tblpayment"
	"inventory/internal/util"
	"time"
)

func AddStatus(client *entgen.Tx, cart RQCart) *entgen.TblCartCreate {
	fmt.Println(cart)
	query := client.TblCart.Create().
		SetID(util.GetUuid()).
		SetUserID(cart.UserID).
		SetProductId(cart.ProductId).
		SetUpdatedAt(time.Now().UTC()).
		SetCreatedAt(time.Now().UTC())

	if cart.Status == "cart" {
		query.SetStatus(tblcart.StatusCart)
	}

	if cart.Status == "brought" {
		query.SetStatus(tblcart.StatusBrought)
	}

	if cart.Status == "canceled" {
		query.SetStatus(tblcart.StatusCanceled)
	}

	return query
}

func Get(client *entgen.Client, id string) *entgen.TblCartQuery {
	return client.TblCart.
		Query().
		Where(tblcart.ID(id)).
		WithInventory()

}

func getOrderId(client *entgen.Client, userId, inventoryId string) *entgen.TblPaymentQuery {
	return client.TblPayment.Query().
		Where(tblpayment.UserId(userId)).
		Where(tblpayment.InventoryId(inventoryId)).
		Where(tblpayment.StatusEQ("paid"))
}

func List(client *entgen.Client, userId string, status string) *entgen.TblCartQuery {
	query := client.TblCart.
		Query().
		Where(tblcart.UserId(userId)).
		Where(tblcart.DeletedAtIsNil()).
		WithInventory()

	if status != "" {
		if status == "cart" {
			query.Where(tblcart.StatusEQ(tblcart.StatusCart))
		}

		if status == "brought" {
			query.Where(tblcart.StatusEQ(tblcart.StatusBrought))
		}

		if status == "canceled" {
			query.Where(tblcart.StatusEQ(tblcart.StatusCanceled))
		}
	}

	return query
}

func GetProduct(client *entgen.Client, id string) *entgen.TblInventoryQuery {
	return client.TblInventory.Query().Where(tblinventory.ID(id))
}

func buy(client *entgen.Tx, payment Payment) *entgen.TblPaymentCreate {
	query := client.TblPayment.Create().
		SetID(payment.Id).
		SetUserId(payment.UserID).
		SetInventoryId(payment.ProductId).
		SetRazorpayOrderId(payment.RazorPayOrderId).
		SetAmount(payment.Amount).
		SetStatus(payment.Status).
		SetUpdatedAt(time.Now().UTC()).
		SetCreatedAt(time.Now().UTC())

	return query
}

func getPayment(client *entgen.Client, id string) *entgen.TblPaymentQuery {
	return client.TblPayment.Query().Where(tblpayment.ID(id))
}

func getPaymentByOrder(client *entgen.Client, orderId string) *entgen.TblPaymentQuery {
	return client.TblPayment.Query().Where(tblpayment.RazorpayOrderId(orderId))
}

func PaymentDone(client *entgen.Client, id string) *entgen.TblPaymentUpdateOne {
	return client.TblPayment.UpdateOneID(id).SetStatus("paid").SetUpdatedAt(time.Now().UTC())
}

func BuyCount(client *entgen.Client, id string) *entgen.TblCartQuery {
	return client.TblCart.Query().
		Where(tblcart.UserId(id)).
		Where(tblcart.StatusEQ(tblcart.StatusBrought))
}

func listCompletedPayment(client *entgen.Client, userId string) *entgen.TblPaymentQuery {
	return client.TblPayment.Query().
		Where(tblpayment.UserId(userId)).
		WithInventory()
}

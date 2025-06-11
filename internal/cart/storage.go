package cart

import (
	"context"
	"fmt"
	"inventory/ent/entgen"
	"inventory/internal/util"
)

type ICartStorage interface {
	Add(ctx context.Context, cart RQCart) (string, error)
	Get(ctx context.Context, id string) (Cart, error)
	List(ctx context.Context, userId string, status string) ([]Cart, error)
	GetInventory(ctx context.Context, id string) (*entgen.TblInventory, error)
	Buy(ctx context.Context, payment Payment) (string, error)
	getPaymentByOrder(ctx context.Context, orderId string) (*entgen.TblPayment, error)
	changeStatus(ctx context.Context, id string) error
	getPayment(ctx context.Context, id string) (Payment, error)
	getBuyCount(ctx context.Context, userId string) (int, error)
	getInvoice(ctx context.Context, userId, inventoryId string) (Payment, error)
	listCompletedPayment(ctx context.Context, userId string) ([]Cart, error)
}

type SCartStorage struct {
	client *entgen.Client
}

func NewCartStorage(client *entgen.Client) ICartStorage {
	return &SCartStorage{
		client: client,
	}
}

func (t *SCartStorage) Get(ctx context.Context, id string) (Cart, error) {
	cartEnt, err := Get(t.client, id).First(ctx)
	if err != nil {
		return Cart{}, util.WrapperForDatabaseError("get", err)
	}

	var cart Cart
	cart.MapFrom(cartEnt)
	return cart, nil

}

func (t *SCartStorage) Add(ctx context.Context, cart RQCart) (string, error) {
	var id string
	err := util.ExecTx(ctx, t.client, func(tx *entgen.Tx) error {
		cartEnt, err := AddStatus(tx, cart).Save(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return util.WrapperForDatabaseError("add", err)
		}
		id = cartEnt.ID
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return id, nil
}

func (t *SCartStorage) List(ctx context.Context, userId string, status string) ([]Cart, error) {
	entCarts, err := List(t.client, userId, status).All(ctx)
	if err != nil {
		return nil, util.WrapperForDatabaseError("list", err)
	}

	carts := make([]Cart, len(entCarts))
	for index, value := range entCarts {
		carts[index].MapFrom(value)
	}
	return carts, nil
}

func (t *SCartStorage) GetInventory(ctx context.Context, id string) (*entgen.TblInventory, error) {
	return GetProduct(t.client, id).First(ctx)
}

func (t *SCartStorage) Buy(ctx context.Context, payment Payment) (string, error) {
	var id string
	err := util.ExecTx(ctx, t.client, func(tx *entgen.Tx) error {
		cartEnt, err := buy(tx, payment).Save(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return util.WrapperForDatabaseError("add", err)
		}
		id = cartEnt.ID
		return nil
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t *SCartStorage) getPayment(ctx context.Context, id string) (Payment, error) {
	entPayment, err := getPayment(t.client, id).First(ctx)
	if err != nil {
		return Payment{}, util.WrapperForDatabaseError("get", err)
	}

	var payment Payment
	payment.MapFrom(entPayment)
	return payment, nil
}

func (t *SCartStorage) getPaymentByOrder(ctx context.Context, orderId string) (*entgen.TblPayment, error) {
	return getPaymentByOrder(t.client, orderId).First(ctx)
}

func (t *SCartStorage) changeStatus(ctx context.Context, id string) error {
	return PaymentDone(t.client, id).Exec(ctx)
}

func (t *SCartStorage) getBuyCount(ctx context.Context, userId string) (int, error) {
	return BuyCount(t.client, userId).Count(ctx)
}

func (t *SCartStorage) getInvoice(ctx context.Context, userId, inventoryId string) (Payment, error) {
	entPayment, err := getOrderId(t.client, userId, inventoryId).First(ctx)
	if err != nil {
		return Payment{}, util.WrapperForDatabaseError("orderId", err)
	}

	var payment Payment
	payment.MapFrom(entPayment)
	return payment, nil
}

func (t *SCartStorage) listCompletedPayment(ctx context.Context, userId string) ([]Cart, error) {
	entPayment, err := listCompletedPayment(t.client, userId).All(ctx)
	if err != nil {
		return nil, util.WrapperForDatabaseError("list payment", err)
	}

	carts := make([]Cart, len(entPayment))
	for index, payment := range entPayment {
		carts[index].MapFrom(payment.Edges.Inventory)
	}
}

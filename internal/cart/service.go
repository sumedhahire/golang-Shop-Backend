package cart

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/razorpay/razorpay-go"
	"inventory/config"
	"inventory/internal/util"
	"os"
	"time"
)

type ICartService interface {
	Add(ctx context.Context, rq RQCart) (RSCart, error)
	Get(ctx context.Context, id string) (RSCart, error)
	List(ctx context.Context, userId string, status string) ([]RSCart, error)
	Buy(ctx context.Context, rq RQCart) (RSPayment, error)
	Verify(ctx context.Context, rq RQPayment) error
	BuyCount(ctx context.Context, userId string) (int, error)
	GetInvoice(ctx context.Context, inventoryId, cartId string) error
}

type SCartService struct {
	razor *razorpay.Client
	//b2Bucket *b2.Bucket
	minio   *minio.Client
	storage ICartStorage
}

func NewCartService(client *config.AppConfig) ICartService {
	return &SCartService{
		//b2Bucket: client.B2Bucket,
		minio:   client.MinioClient,
		razor:   client.Razor,
		storage: NewCartStorage(client.Client),
	}
}

func (s SCartService) BuyCount(ctx context.Context, userId string) (int, error) {
	return s.storage.getBuyCount(ctx, userId)
}

func (s SCartService) Get(ctx context.Context, id string) (RSCart, error) {
	mapCart, err := s.storage.Get(ctx, id)
	if err != nil {
		return RSCart{}, err
	}

	var rs RSCart
	rs.MapFrom(&mapCart)
	rs.ImageUrl, err = s.getMinioLink(rs.ImageUrl)
	if err != nil {
		return RSCart{}, err
	}
	return rs, nil
}

func (s SCartService) getMinioLink(filename string) (string, error) {
	bucketName := os.Getenv("MINIO_BUCKET")
	expiry := time.Minute * 10

	presignedURL, err := s.minio.PresignedGetObject(context.Background(), bucketName, filename, expiry, nil)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil

}

//func (s SCartService) GetAuthorizedDownloadURL(ctx context.Context, filename string, duration time.Duration) (string, error) {
//	// Generate download auth token valid for duration
//	authToken, err := s.b2Bucket.AuthToken(ctx, filename, duration)
//	if err != nil {
//		return "", err
//	}
//	// Construct full URL with token
//	url := fmt.Sprintf("https://f005.backblazeb2.com/file/%s/%s?Authorization=%s",
//		s.b2Bucket.Name(), filename, authToken)
//	return url, nil
//}

func (s SCartService) Add(ctx context.Context, rq RQCart) (RSCart, error) {
	//cart := rq.MapTo()
	id, err := s.storage.Add(ctx, rq)
	if err != nil {
		return RSCart{}, err
	}

	mapCart, err := s.storage.Get(ctx, id)
	if err != nil {
		return RSCart{}, err
	}

	var rs RSCart
	rs.MapFrom(&mapCart)
	return rs, nil
}

func (s SCartService) List(ctx context.Context, userId string, status string) ([]RSCart, error) {
	mapCart, err := s.storage.List(ctx, userId, status)
	if err != nil {
		return nil, err
	}
	carts := make([]RSCart, len(mapCart))
	for index, value := range mapCart {
		value.ImageUrl, err = s.getMinioLink(value.ImageUrl)
		if err != nil {
			return nil, err
		}
		carts[index].MapFrom(&value)

	}
	return carts, nil
}

func (s SCartService) Buy(ctx context.Context, rq RQCart) (RSPayment, error) {
	var err error
	product, err := s.storage.GetInventory(ctx, rq.ProductId)
	if err != nil {
		return RSPayment{}, err
	}

	data := map[string]interface{}{
		"amount":   int(product.Price) * 100,
		"currency": "INR",
		"receipt":  util.GetUuid(),
	}

	rzpOrder, err := s.razor.Order.Create(data, nil)
	if err != nil {
		return RSPayment{}, err
	}

	rzpOrderID := rzpOrder["id"].(string)

	payment := Payment{
		Id:              util.GetUuid(),
		ProductId:       product.ID,
		UserID:          rq.UserID,
		RazorPayOrderId: rzpOrderID,
		Amount:          product.Price,
		Status:          "created",
	}

	id, err := s.storage.Buy(ctx, payment)
	if err != nil {
		return RSPayment{}, err
	}

	paymentRs, err := s.storage.getPayment(ctx, id)
	if err != nil {
		return RSPayment{}, err
	}

	var rs RSPayment
	rs.MapFrom(&paymentRs)

	return rs, nil
}

func (s SCartService) Verify(ctx context.Context, rq RQPayment) error {
	paymentData, err := s.razor.Payment.Fetch(rq.RazorpayPaymentID, nil, nil)
	if err != nil {
		return err
	}

	// Step 2: Check if status is captured
	status := paymentData["status"].(string)
	if status != "captured" {
		return errors.New("payment not captured")
	}

	pmt, err := s.storage.getPaymentByOrder(ctx, rq.RazorpayOrderID)
	if err != nil {
		return err
	}

	err = s.storage.changeStatus(ctx, pmt.ID)
	if err != nil {
		return err
	}

	var rqCart RQCart
	rqCart.MapFromPayment(pmt)
	rqCart.Status = "brought"
	_, err = s.storage.Add(ctx, rqCart)
	if err != nil {
		return err
	}

	return nil
}

func (s SCartService) GetInvoice(ctx context.Context, inventoryId, cartId string) error {
	payment, err := s.storage.getInvoice(ctx, inventoryId, cartId)
	if err != nil {
		return err
	}

	orderId := payment.OrderId
	order, err := s.razor.Order.Fetch(orderId, nil, nil)
	if err != nil {
		return err
	}

	receipt, ok := order["receipt"].(string)
	if !ok {
		receipt = "No receipt number set"
	}

	fmt.Printf("Order ID: %s\nReceipt: %s\nAmount: %v\nCurrency: %s\nStatus: %s\n",
		order["id"], receipt, order["amount"], order["currency"], order["status"])

	return nil
}

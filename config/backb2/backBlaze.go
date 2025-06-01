package backb2

import (
	"github.com/kurin/blazer/b2"
	"golang.org/x/net/context"
	"os"
	"sync"
)

var (
	b2Bucket *b2.Bucket
	once     sync.Once
)

func getB2() {
	ctx := context.Background()
	client, err := b2.NewClient(ctx, os.Getenv("B2_ACCOUNT_ID"), os.Getenv("B2_ACCOUNT_KEY"))
	if err != nil {
		panic(err)
	}

	bucket, err := client.Bucket(ctx, "kkwCollege")
	if err != nil {
		panic(err)
	}

	b2Bucket = bucket
}

func InitB2() *b2.Bucket {
	once.Do(func() {
		getB2()
	})
	return b2Bucket
}

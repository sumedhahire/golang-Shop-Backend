package config

import (
	"github.com/kurin/blazer/b2"
	"github.com/razorpay/razorpay-go"
	"inventory/config/backb2"
	"inventory/config/dbConfig"
	"inventory/ent/entgen"
	"os"
	"sync"
)

var (
	appConfig *AppConfig
	once      sync.Once
)

type AppConfig struct {
	Client   *entgen.Client
	Razor    *razorpay.Client
	B2Bucket *b2.Bucket
}

func InitServices() *AppConfig {
	once.Do(func() {
		appConfig = &AppConfig{
			B2Bucket: backb2.InitB2(),
			Client:   dbConfig.InitDB(),
			Razor:    razorpay.NewClient(os.Getenv("RAZOR_ID"), os.Getenv("RAZOR_SECRET")),
		}
	})
	return appConfig
}

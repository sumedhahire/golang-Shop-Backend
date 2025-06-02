package config

import (
	"github.com/minio/minio-go/v7"
	"github.com/razorpay/razorpay-go"
	"inventory/config/dbConfig"
	minio2 "inventory/config/minio"
	"inventory/ent/entgen"
	"os"
	"sync"
)

var (
	appConfig *AppConfig
	once      sync.Once
)

type AppConfig struct {
	Client *entgen.Client
	Razor  *razorpay.Client
	//B2Bucket *b2.Bucket
	MinioClient *minio.Client
}

func InitServices() *AppConfig {
	once.Do(func() {
		appConfig = &AppConfig{
			//B2Bucket: backb2.InitB2(),
			MinioClient: minio2.InitMinio(),
			Client:      dbConfig.InitDB(),
			Razor:       razorpay.NewClient(os.Getenv("RAZOR_ID"), os.Getenv("RAZOR_SECRET")),
		}
	})
	return appConfig
}

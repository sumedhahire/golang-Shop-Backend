package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"sync"
)

var (
	MiniClient *minio.Client
	once       sync.Once
)

func getMini() {
	endPoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyId := os.Getenv("MINIO_ID")
	accessKeySecret := os.Getenv("MINIO_SECRET")
	useSSL := false

	client, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, accessKeySecret, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}

	MiniClient = client // ← THIS is the missing line

	bucketName := os.Getenv("MINIO_BUCKET")
	ctx := context.Background()

	err = MiniClient.SetBucketPolicy(ctx, bucketName, "")
	if err != nil {
		panic(err)
	}
	log.Println("Bucket policy set to public read")
}

func InitMinio() *minio.Client {
	once.Do(func() {
		getMini()
	})
	return MiniClient
}

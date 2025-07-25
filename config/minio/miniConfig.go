package minio

import (
	"context"
	"fmt"
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
	location := "us-east-1"

	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("Failed to check bucket: %v", err)
	}

	if !exists {
		err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
		log.Printf("✅ Created bucket: %s\n", bucketName)
	} else {
		log.Printf("✅ Bucket already exists: %s\n", bucketName)
	}

	publicPolicy := fmt.Sprintf(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Action": ["s3:GetObject"],
					"Effect": "Allow",
					"Principal": {"AWS": ["*"]},
					"Resource": ["arn:aws:s3:::%s/*"]
				}
			]
		}`, bucketName)
	err = MiniClient.SetBucketPolicy(ctx, bucketName, publicPolicy)
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

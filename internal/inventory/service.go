package inventory

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"inventory/config"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

type IService interface {
	Get(ctx context.Context, id string) (RSInventory, error)
	List(ctx context.Context, filter Filter) ([]RSInventory, error)
	Add(ctx context.Context, rq RQInventory, file *multipart.FileHeader) (RSInventory, error)
	Update(ctx context.Context, rq RQInventory, updateId string) (RSInventory, error)
}

type SService struct {
	minio *minio.Client
	//b2Bucket *b2.Bucket
	storage IStorage
}

func NewService(client *config.AppConfig) IService {
	return SService{
		minio: client.MinioClient,
		//b2Bucket: client.B2Bucket,
		storage: NewStorage(client.Client),
	}
}

func (s SService) Get(ctx context.Context, id string) (RSInventory, error) {
	inventory, err := s.storage.Get(ctx, id)
	if err != nil {
		return RSInventory{}, err
	}
	var rs RSInventory
	rs.MapTo(inventory)
	//rs.ImageUrl, err = s.GetAuthorizedDownloadURL(ctx, inventory.ImageUrl, time.Minute*30)
	//if err != nil {
	//	return RSInventory{}, err
	//}
	//rs.ImageUrl, err = s.getMinioLink(rs.ImageUrl)
	rs.ImageUrl = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PUBLIC_ENDPOINT"), rs.ImageUrl)

	return rs, nil
}

//--------------------------------------------------------
//func (s SService) GetAuthorizedDownloadURL(ctx context.Context, filename string, duration time.Duration) (string, error) {
//	// Generate download auth token valid for duration
//	authToken, err := s.b2Bucket.AuthToken(ctx, filename, duration)
//	if err != nil {
//		return "", err
//	}
//
//	// Construct full URL with token
//	url := fmt.Sprintf("https://f005.backblazeb2.com/file/%s/%s?Authorization=%s",
//		s.b2Bucket.Name(), filename, authToken)
//	return url, nil
//}
//-----------------------------------------------------

func (s SService) List(ctx context.Context, filter Filter) ([]RSInventory, error) {
	inventories, err := s.storage.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	rsArr := make([]RSInventory, len(inventories))
	for i, inv := range inventories {
		var rs RSInventory
		rs.MapTo(&inv)
		//rs.ImageUrl, err = s.GetAuthorizedDownloadURL(ctx, inv.ImageUrl, time.Minute*30)
		//if err != nil {
		//	return nil, err
		//	}
		//rs.ImageUrl, err = s.getMinioLink(rs.ImageUrl)
		rs.ImageUrl = fmt.Sprintf("%s/%s", os.Getenv("MINIO_PUBLIC_ENDPOINT"), rs.ImageUrl)
		//if err != nil {
		//	return nil, err
		//}
		rsArr[i] = rs
	}

	return rsArr, nil
}

func (s SService) Add(ctx context.Context, rq RQInventory, fileHeader *multipart.FileHeader) (RSInventory, error) {
	var err error
	inv := rq.MapFrom()

	file, err := fileHeader.Open()
	if err != nil {
		return RSInventory{}, err
	}
	defer file.Close()

	// Optional: rename file to avoid collisions
	ext := ""
	parts := strings.Split(fileHeader.Filename, ".")
	if len(parts) > 1 {
		ext = parts[len(parts)-1]
	}
	filename := fmt.Sprintf("%d.%s", time.Now().UnixNano(), ext)

	url, err := s.uploadToMinio(file, fileHeader, filename)
	if err != nil {
		return RSInventory{}, err
	}

	inv.ImageUrl = url

	id, err := s.storage.Add(ctx, &inv)
	if err != nil {
		return RSInventory{}, err
	}

	inventory, err := s.storage.Get(ctx, id)
	if err != nil {
		return RSInventory{}, err
	}

	var rs RSInventory
	rs.MapTo(inventory)

	return rs, nil
}

func (s SService) Update(ctx context.Context, rq RQInventory, updateId string) (RSInventory, error) {
	err := s.storage.Update(ctx, updateId, rq)
	if err != nil {
		return RSInventory{}, err
	}

	inventory, err := s.storage.Get(ctx, updateId)
	if err != nil {
		return RSInventory{}, err
	}

	var rs RSInventory
	rs.MapTo(inventory)
	return rs, nil
}

//func (s SService) uploadToB2(file io.Reader, filename string) (string, error) {
//	ctx := context.Background()
//
//	obj := s.b2Bucket.Object(filename)
//	fmt.Println(obj)
//	writer := obj.NewWriter(ctx)
//	defer writer.Close()
//
//	if _, err := io.Copy(writer, file); err != nil {
//		return "", err
//	}
//
//	return filename, nil
//}

func (s SService) uploadToMinio(file multipart.File, fileHeader *multipart.FileHeader, filename string) (string, error) {
	bucketName := os.Getenv("MINIO_BUCKET")
	_, err := s.minio.PutObject(
		context.Background(),
		bucketName,      // bucket name
		filename,        // object name in bucket
		file,            // file reader (stream)
		fileHeader.Size, // file size
		minio.PutObjectOptions{ // content-type
			ContentType: fileHeader.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func (s SService) getMinioLink(filename string) (string, error) {
	bucketName := os.Getenv("MINIO_BUCKET")
	expiry := time.Minute * 10

	presignedURL, err := s.minio.PresignedGetObject(context.Background(), bucketName, filename, expiry, nil)
	if err != nil {
		return "", err
	}

	publicURL := strings.Replace(presignedURL.String(), "host.docker.internal", "localhost", 1)
	return publicURL, nil

}

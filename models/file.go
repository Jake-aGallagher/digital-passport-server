package models

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/joho/godotenv"
)

var client *azblob.Client
var clientErr error
var once sync.Once

func GetFilesClient() (*azblob.Client, error) {
	var err error
	once.Do(func() {
		err = godotenv.Load(".env")
		if err != nil {
			clientErr = fmt.Errorf("error loading .env file: %w", err)
			return
		}
		accountName := os.Getenv("STORAGEACOUNTNAME")
		accountKey := os.Getenv("STORAGEACOUNTKEY")

		cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
		if err != nil {
			clientErr = fmt.Errorf("error creating shared key credential: %w", err)
			return
		}
		url := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)
		client, err = azblob.NewClientWithSharedKeyCredential(url, cred, nil)
		if err != nil {
			clientErr = fmt.Errorf("error creating client: %w", err)
			return
		}
	})
	return client, clientErr

}

func Download(blobName string, destination string) error {
	containerName := "documents"
	client, err := GetFilesClient()
	if err != nil {
		return err
	}
	target := path.Join(destination, blobName)
	d, err := os.Create(target)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = client.DownloadFile(context.Background(), containerName, blobName, d, nil)
	return err
}

func Upload(data []byte, blobName string) error {
	client, err := GetFilesClient()
	if err != nil {
		fmt.Println(err)
		return err
	}
	res, err := client.UploadBuffer(context.Background(), "documents", blobName, data, &azblob.UploadBufferOptions{})
	fmt.Println("response: ", res)
	if err != nil {
		return err
	}
	return nil
}

func GetFileUrl(fileId string) (string, error) {
	client, err := GetFilesClient()
	if err != nil {
		return "", err
	}
	url := client.URL()

	return url + "documents/" + fileId, nil
}

func DownloadStream(fileId string) (blob.DownloadStreamResponse, error) {
	client, err := GetFilesClient()
	if err != nil {
		return blob.DownloadStreamResponse{}, err
	}

	res, err := client.DownloadStream(context.Background(), "documents", fileId, nil)
	if err != nil {
		return blob.DownloadStreamResponse{}, err
	}

	return res, nil
}

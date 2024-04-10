package routes

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"example.com/digital-passport/models"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func addFile(context *gin.Context) {
	fileHeader, err := context.FormFile("file")
	if err != nil {
		context.JSON(400, gin.H{"message": "unable to parse file"})
		return
	}

	fmt.Printf("Received a file named %s with size %d", fileHeader.Filename, fileHeader.Size)

	// add uuid to filename
	rand := uuid.NewString()
	fileHeader.Filename = rand + "___" + fileHeader.Filename

	fmt.Println("new filename: ", fileHeader.Filename)

	file, err := fileHeader.Open()
	if err != nil {
		context.JSON(500, gin.H{"message": "unable to open file"})
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		context.JSON(500, gin.H{"message": "unable to read file"})
		return
	}

	err = models.Upload(data, fileHeader.Filename)
	if err != nil {
		context.JSON(500, gin.H{"message": "unable to save file"})
		return
	}

	context.JSON(201, gin.H{"message": "file uploaded successfully", "fileName": fileHeader.Filename})
}

func getFile(c *gin.Context) {
	fileId := c.Param("id")
	if fileId == "" {
		c.JSON(400, gin.H{"message": "Could not parse file id"})
		return
	}

	file, err := models.DownloadStream(fileId)
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not get file"})
		return
	}

	downloadedData := bytes.Buffer{}
	retryReader := file.NewRetryReader(context.Background(), &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not get file"})
		return
	}

	err = retryReader.Close()
	if err != nil {
		c.JSON(400, gin.H{"message": "Could not get file"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=file-name.txt")
	c.Data(200, *file.ContentType, downloadedData.Bytes())
}

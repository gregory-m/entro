package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

var uploader *s3manager.Uploader
var bucket = "122"

func postRoot(w http.ResponseWriter, r *http.Request) {

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	id := uuid.New().String()
	fileName := fmt.Sprintf("%s-%s", timestamp, id)

	fmt.Printf("Uploading to: %s\n", fileName)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Body:   r.Body,
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't upload %s\n", err), http.StatusInternalServerError)
	}

	io.WriteString(w, "OK\n")
}

func main() {

	bucket = os.Getenv("S3_BUCKET")
	if bucket == "" {
		bucket = "g-test2"
	}

	fmt.Printf("Uploading to bucket: %s\n", bucket)

	//

	http.HandleFunc("/", postRoot)

	sess, err := session.NewSession(&aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		// Region: aws.String("eu-north-1"),
	})
	if err != nil {
		fmt.Printf("Can't create aws session: %s\n", err)
		os.Exit(1)
	}

	uploader = s3manager.NewUploader(sess)

	fmt.Printf("Starring server on port 3333\n")
	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("Can't listen: %s\n", err)
		os.Exit(1)
	}
}

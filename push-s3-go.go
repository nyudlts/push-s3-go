// This code is pulled from this web page:
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/nyudlts/push-s3-go/config"
	"github.com/nyudlts/push-s3-go/util"
)

func main() {

	if len(os.Args) != 3 {
		util.ExitErrorf("path-to-file required\nUsage: %s <path-to-file> <S3 key>\ne.g., %s ../foo.wav a/b/c/d/foo.wav",
			os.Args[0], os.Args[0])
	}

	conf, err := config.GetConf()
	if err != nil {
		util.ExitErrorf("Problem obtaining config: %v", err)
	}

	os.Setenv("AWS_ACCESS_KEY_ID", conf.S3AwsAccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", conf.S3AwsSecretAccessKey)
	os.Setenv("AWS_REGION", conf.S3AwsRegion)

	filePath := os.Args[1]
	s3key := os.Args[2]
	bucket := conf.S3BucketId

	file, err := os.Open(filePath)
	if err != nil {
		util.ExitErrorf("Unable to open file %q, %v", err)
	}
	defer file.Close()

	sess, err := session.NewSession()
	if err != nil {
		util.ExitErrorf("Problem creating AWS session %q, %v", err)
	}

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3key),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		util.ExitErrorf("Unable to upload %q to %q, as %q %v", filePath, bucket, s3key, err)
	}

	fmt.Printf("Successfully uploaded %q to %q as %q\n", filePath, bucket, s3key)
}

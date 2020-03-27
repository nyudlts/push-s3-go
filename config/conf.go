package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Conf struct {
	S3BucketId           string
	S3BaseUrl            string
	S3KeyPrefix          string
	S3AwsAccessKeyId     string
	S3AwsSecretAccessKey string
	S3AwsRegion          string
}

func GetConf() (Conf, error) {

	var conf Conf

	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.push-s3-go")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return conf, fmt.Errorf("Fatal error config file: %s \n", err)
	}

	conf.S3BaseUrl = viper.GetString("s3.base_url")
	conf.S3BucketId = viper.GetString("s3.bucket")
	conf.S3KeyPrefix = viper.GetString("s3.key_prefix")
	conf.S3AwsAccessKeyId = viper.GetString("s3.aws_access_key_id")
	conf.S3AwsSecretAccessKey = viper.GetString("s3.aws_secret_access_key")
	conf.S3AwsRegion = viper.GetString("s3.aws_region")

	return conf, nil
}

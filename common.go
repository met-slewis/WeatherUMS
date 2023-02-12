package WeatherUMS

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("coredev-dev"))
	if err != nil {
		panic(err)
	}
	s3Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = "ap-southeast-2"
	})
}

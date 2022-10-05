// Encoding: UTF-8

package main

import (
	"context"
	"flag"
	"os"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cw_types "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

var Debug, OneShot bool

var Interval time.Duration

func init() {
	flag.BoolVar(&OneShot, "once", false, "Run only once")
	flag.BoolVar(&Debug, "debug", false, "Enable debug logging")
	flag.DurationVar(&Interval, "interval", 15*time.Minute, "Reporting interval for CloudWatch logs")
}

func main() {
	flag.Parse()

	if versionFlag {
		showVersion()
		os.Exit(0)
	}

	if _, ok := os.LookupEnv("AWS_CW_UPTIME_DEBUG"); ok || Debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}

	if runtime.GOOS == "windows" {
		RunWindows()
	} else {
		Run()
	}
}

func Run() {
	ctx := context.Background()

	cfgctx, cancel := context.WithTimeout(ctx, time.Duration(10*time.Second))
	defer cancel()

	cfg, err := config.LoadDefaultConfig(cfgctx,
		config.WithEC2IMDSRegion(),
		config.WithSharedConfigFiles([]string{}),
		config.WithSharedCredentialsFiles([]string{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	metadata := imds.New(imds.Options{})
	identity, err := metadata.GetInstanceIdentityDocument(ctx, &imds.GetInstanceIdentityDocumentInput{})
	if err != nil {
		log.Fatal(err)
	}
	cw := cloudwatch.NewFromConfig(cfg)

	for {
		input := &cloudwatch.PutMetricDataInput{
			Namespace: aws.String("Custom"),
			MetricData: []cw_types.MetricDatum{
				{
					MetricName: aws.String("Uptime"),
					Unit:       cw_types.StandardUnitSeconds,
					Value:      aws.Float64(GetUptime().Seconds()),
					Dimensions: []cw_types.Dimension{
						{
							Name:  aws.String("InstanceId"),
							Value: aws.String(identity.InstanceID),
						},
					},
				},
			},
		}

		if _, err := cw.PutMetricData(ctx, input); err != nil {
			log.Fatal(err)
		}

		if OneShot {
			break
		}

		log.Infof("Waiting %s before updating Uptime metric...", Interval)
		time.Sleep(Interval)
	}
}

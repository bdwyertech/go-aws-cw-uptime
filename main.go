// Encoding: UTF-8

package main

import (
	"flag"
	"os"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var OneShotFlag bool

func init() {
	flag.BoolVar(&OneShotFlag, "once", false, "Run only once")

	if _, ok := os.LookupEnv("AWS_CW_UPTIME_DEBUG"); ok {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}
}

func main() {
	flag.Parse()
	if versionFlag {
		showVersion()
		os.Exit(0)
	}

	if runtime.GOOS == "windows" {
		RunWindows()
	} else {
		Run()
	}
}

func Run() {
	// AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            *aws.NewConfig().WithCredentialsChainVerboseErrors(true),
		SharedConfigState: session.SharedConfigDisable,
	}))

	metadata := ec2metadata.New(sess)

	if !metadata.Available() {
		log.Fatal("EC2 Metadata is not available... Are we running on an EC2 instance?")
	}

	identity, err := metadata.GetInstanceIdentityDocument()
	if err != nil {
		log.Fatal(err)
	}
	instanceID := identity.InstanceID
	sess.Config = sess.Config.WithRegion(identity.Region)

	cw := cloudwatch.New(sess)

	interval := 30 * time.Second
	if durationString, ok := os.LookupEnv("AWS_CW_UPTIME_INTERVAL"); ok {
		if interval, err = time.ParseDuration(durationString); err != nil {
			log.Fatal(err)
		}
	}

	for {
		input := &cloudwatch.PutMetricDataInput{
			Namespace: aws.String("Custom"),
			MetricData: []*cloudwatch.MetricDatum{
				&cloudwatch.MetricDatum{
					MetricName: aws.String("Uptime"),
					Unit:       aws.String("Seconds"),
					Value:      aws.Float64(GetUptime().Seconds()),
					Dimensions: []*cloudwatch.Dimension{
						&cloudwatch.Dimension{
							Name:  aws.String("InstanceId"),
							Value: aws.String(instanceID),
						},
					},
				},
			},
		}

		if _, err := cw.PutMetricData(input); err != nil {
			log.Fatal(err)
		}

		if OneShotFlag {
			break
		}

		log.Debugf("Waiting %s before updating metric...", interval)
		time.Sleep(interval)
	}
}

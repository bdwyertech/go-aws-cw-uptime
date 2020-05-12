// Encoding: UTF-8

package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func init() {
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

	log.Error(GetUptime())

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

	log.Error(instanceID)

	cw := cloudwatch.New(sess)

	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Custom"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("Uptime"),
				Unit:       aws.String("Count"),
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
		log.Error(err)
	}
}

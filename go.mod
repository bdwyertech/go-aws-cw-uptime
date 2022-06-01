module aws-cloudwatch-uptime

go 1.16

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2 v1.16.4
	github.com/aws/aws-sdk-go-v2/config v1.15.9
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.5
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.8.1
	github.com/kardianos/service v1.2.0
	github.com/shirou/gopsutil v3.21.8+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	golang.org/x/sys v0.0.0-20210930141918-969570ce7c6c
)

# AWS CloudWatch Uptime

[![Build Status](https://github.com/bdwyertech/go-aws-asg-status/workflows/Go/badge.svg?branch=master)](https://github.com/bdwyertech/go-aws-asg-status/actions?query=workflow%3AGo+branch%3Amaster)

This is a tool to report an EC2 instances uptime to CloudWatch as a custom metric.  By default, it reports every 15 minutes.

#### CloudWatch Metric Parameters
  * Namespace: `Custom`
  * MetricName: `Uptime`
  * Unit: `Seconds`
  * Dimension: `InstanceId`



### Installing & Running as a Service

#### Linux
```ini
[Unit]
Description=CloudWatch - System Uptime Daemon
Requires=amazon-cloudwatch-agent.service
After=amazon-cloudwatch-agent.service

[Service]
ExecStart=/usr/local/bin/aws-cloudwatch-uptime
```

#### Windows
```sh
aws-cloudwatch-uptime.exe -service install
aws-cloudwatch-uptime.exe -service start
```

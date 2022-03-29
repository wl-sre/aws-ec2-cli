package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	awskey    *string
	awssecret *string
)

func init() {
	awskey = flag.String("key", "", "the AWS_ACCESS_KEY_ID")
	awssecret = flag.String("secret", "", "the AWS_SECRET_ACCESS_KEY")
}

func main() {
	if len(os.Args) < 2 {
		exitErrorf("Please input the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY ")
	}

	flag.Parse()

	result, err := GetInstances()
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return
	}
	var instanceName string
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			for _, t := range i.Tags {
				if *t.Key == "Name" {
					instanceName = *t.Value
					break
				}
			}
			fmt.Printf("%-20s %-20s\n", *i.PrivateIpAddress, instanceName)
			fmt.Println()
		}
	}

}

func GetInstances() (*ec2.DescribeInstancesOutput, error) {
	sess, _ := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(*awskey, *awssecret, ""),
		Region:      aws.String("cn-northwest-1"),
	})
	svc := ec2.New(sess)
	// result, err := svc.DescribeInstances(nil)
	result, err := svc.DescribeInstances(nil)

	if err != nil {
		return nil, err
	}
	return result, nil
}
func exitErrorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

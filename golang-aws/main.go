package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	ctx := context.Background()

	// Load AWS config (reads from ~/.aws or environment)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	// 2. Launch EC2 instance
	instanceID, err := launchEC2Instance(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to launch EC2 instance: %v", err)
	}
	fmt.Println("Launched EC2 instance:", instanceID)
}

func launchEC2Instance(ctx context.Context, cfg aws.Config) (string, error) {
	ec2Client := ec2.NewFromConfig(cfg)

	runInput := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-021a584b49225376d"),
		InstanceType: ec2types.InstanceTypeT3Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
		KeyName:      aws.String("test"),
		SecurityGroupIds: []string{
			"sg-025293188426f6397",
		},
	}

	runResp, err := ec2Client.RunInstances(ctx, runInput)
	if err != nil {
		return "", fmt.Errorf("failed to launch instance: %w", err)
	}

	if len(runResp.Instances) == 0 {
		return "", fmt.Errorf("no instances launched")
	}

	return aws.ToString(runResp.Instances[0].InstanceId), nil
}

func createAndSaveKeyPair(ctx context.Context, cfg aws.Config, keyName string) error {
	ec2Client := ec2.NewFromConfig(cfg)

	out, err := ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
		KeyName: aws.String(keyName),
		KeyType: ec2types.KeyTypeRsa,
	})
	if err != nil {
		return fmt.Errorf("create key pair failed: %w", err)
	}

	// Save the private key material to a .pem file
	filename := keyName + ".pem"
	err = os.WriteFile(filename, []byte(aws.ToString(out.KeyMaterial)), 0600)
	if err != nil {
		return fmt.Errorf("failed to write key file: %w", err)
	}

	fmt.Println("Saved private key to", filename)
	return nil
}

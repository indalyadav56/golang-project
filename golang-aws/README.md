# Golang AWS EC2 Instance Creator

This application uses the AWS SDK for Go v2 to create an EC2 instance with 2 CPUs and 4GB memory (t3.medium).

## Prerequisites

1. Go 1.24 or later
2. AWS credentials configured on your machine

## AWS Credentials Setup

Before running this application, you need to set up your AWS credentials. You can do this in one of the following ways:

1. **Using AWS CLI**:
   ```
   aws configure
   ```
   
2. **Using Environment Variables**:
   ```
   export AWS_ACCESS_KEY_ID="your-access-key"
   export AWS_SECRET_ACCESS_KEY="your-secret-key"
   export AWS_REGION="us-east-1"
   ```

## Configuration

The application is configured to:
- Launch a t3.medium instance (2 vCPUs, 4GB memory)
- Use Amazon Linux 2 AMI (you may need to update the AMI ID)
- Deploy in the us-east-1 region (configurable)
- Tag the instance with the name "GoSDK-EC2-Instance"

## Running the Application

```
go run main.go
```

## Important Notes

1. The AMI ID in the code is an example and may be outdated. In a production environment, you should query for the latest AMI ID or use a known current one for your region.

2. This application will create actual AWS resources that may incur charges to your AWS account. Remember to terminate any instances you don't need.

3. Make sure you have the necessary permissions in your AWS account to create EC2 instances.

## Code Structure

- `main.go` - Contains the main application code
- `go.mod` - Go module definition with dependencies

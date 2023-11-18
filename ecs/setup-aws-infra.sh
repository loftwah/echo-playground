#!/bin/bash

# Set AWS default region
AWS_REGION="ap-southeast-2"
SECURITY_GROUP_NAME="echo-playground-sg"

# Set AWS region
aws configure set default.region $AWS_REGION

# Get the default VPC
DEFAULT_VPC_ID=$(aws ec2 describe-vpcs --filters "Name=is-default,Values=true" --query "Vpcs[0].VpcId" --output text)

if [ $DEFAULT_VPC_ID = "None" ]; then
    echo "No default VPC found."
    exit 1
fi

echo "Default VPC ID: $DEFAULT_VPC_ID"

# Get subnets of the default VPC
SUBNET_IDS=$(aws ec2 describe-subnets --filters "Name=vpc-id,Values=$DEFAULT_VPC_ID" --query "Subnets[*].SubnetId" --output text)
echo "Subnet IDs: $SUBNET_IDS"

# Create a new security group
SECURITY_GROUP_ID=$(aws ec2 create-security-group --group-name $SECURITY_GROUP_NAME --description "Security group for Echo Playground" --vpc-id $DEFAULT_VPC_ID --query 'GroupId' --output text)
echo "Security Group Created with ID: $SECURITY_GROUP_ID"

# Add a rule to allow inbound HTTP traffic (port 1323)
aws ec2 authorize-security-group-ingress --group-id $SECURITY_GROUP_ID --protocol tcp --port 1323 --cidr 0.0.0.0/0

echo "Inbound HTTP rule added to Security Group."

# Output the details
echo "Update your ECS task and service definition files with the following details:"
echo "VPC ID: $DEFAULT_VPC_ID"
echo "Subnet IDs: $SUBNET_IDS"
echo "Security Group ID: $SECURITY_GROUP_ID"

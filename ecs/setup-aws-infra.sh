#!/bin/bash

# Set AWS default region
AWS_REGION="ap-southeast-2"
SECURITY_GROUP_NAME="echo-playground-sg"
ECS_TASK_EXECUTION_ROLE_NAME="ecsTaskExecutionRole"

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

# Check if the security group already exists
EXISTING_SG_ID=$(aws ec2 describe-security-groups --filters "Name=group-name,Values=$SECURITY_GROUP_NAME" --query 'SecurityGroups[0].GroupId' --output text)

if [ $EXISTING_SG_ID = "None" ]; then
    # Create a new security group
    SECURITY_GROUP_ID=$(aws ec2 create-security-group --group-name $SECURITY_GROUP_NAME --description "Security group for Echo Playground" --vpc-id $DEFAULT_VPC_ID --query 'GroupId' --output text)
    echo "Security Group Created with ID: $SECURITY_GROUP_ID"
else
    SECURITY_GROUP_ID=$EXISTING_SG_ID
    echo "Security Group Already Exists with ID: $EXISTING_SG_ID"
fi

# Add a rule to allow inbound HTTP traffic (port 1323) if not already present
aws ec2 authorize-security-group-ingress --group-id $SECURITY_GROUP_ID --protocol tcp --port 1323 --cidr 0.0.0.0/0 2>/dev/null || echo "Inbound HTTP rule already exists."

# Check if ECS Task Execution Role exists
ECS_TASK_EXECUTION_ROLE_ARN=$(aws iam get-role --role-name $ECS_TASK_EXECUTION_ROLE_NAME --query 'Role.Arn' --output text 2>/dev/null)

if [ -z "$ECS_TASK_EXECUTION_ROLE_ARN" ] || [ $ECS_TASK_EXECUTION_ROLE_ARN = "None" ]; then
    # Create the role and attach AmazonECSTaskExecutionRolePolicy
    ECS_TASK_EXECUTION_ROLE_ARN=$(aws iam create-role --role-name $ECS_TASK_EXECUTION_ROLE_NAME --assume-role-policy-document file://ecs-task-trust-policy.json --query 'Role.Arn' --output text)
    aws iam attach-role-policy --role-name $ECS_TASK_EXECUTION_ROLE_NAME --policy-arn arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy
    echo "ECS Task Execution Role Created with ARN: $ECS_TASK_EXECUTION_ROLE_ARN"
else
    echo "ECS Task Execution Role Already Exists with ARN: $ECS_TASK_EXECUTION_ROLE_ARN"
fi

# Output the details
echo "Update your ECS task and service definition files with the following details:"
echo "VPC ID: $DEFAULT_VPC_ID"
echo "Subnet IDs: $SUBNET_IDS"
echo "Security Group ID: $SECURITY_GROUP_ID"
echo "ECS Task Execution Role ARN: $ECS_TASK_EXECUTION_ROLE_ARN"

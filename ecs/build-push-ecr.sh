#!/bin/bash

# Set variables
AWS_REGION="ap-southeast-2"
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query "Account" --output text)
ECR_REPO_NAME="echo-playground"
IMAGE_TAG="latest"
DOCKERFILE_DIR="../" # Directory where Dockerfile is located
CLUSTER_NAME="echo-playground"
SERVICE_NAME="echo-playground-service"

# Configure AWS region
aws configure set default.region $AWS_REGION

# Check if the ECR repository exists
REPO_EXISTS=$(aws ecr describe-repositories --repository-names $ECR_REPO_NAME --region $AWS_REGION --output text --query 'repositories[0].repositoryName' 2>/dev/null)
if [ "$REPO_EXISTS" != "$ECR_REPO_NAME" ]; then
    aws ecr create-repository --repository-name $ECR_REPO_NAME --region $AWS_REGION
    echo "ECR repository $ECR_REPO_NAME created."
else
    echo "ECR repository $ECR_REPO_NAME already exists."
fi

# Change to the directory containing the Dockerfile
cd $DOCKERFILE_DIR

# Build and push Docker image to ECR
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com
docker build -t $ECR_REPO_NAME:latest -f Dockerfile .
docker tag $ECR_REPO_NAME:latest $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG

# Check if docker push was successful
if [ $? -ne 0 ]; then
    echo "Docker push failed."
    exit 1
fi

echo "Image ready."

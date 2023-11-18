#!/bin/bash

# Set variables
AWS_REGION="ap-southeast-2"
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query "Account" --output text)
ECR_REPO_NAME="echo-playground"
IMAGE_TAG="latest"

# Configure AWS region
aws configure set default.region $AWS_REGION

# 1. Create ECR repository (ignore if it already exists)
aws ecr create-repository --repository-name $ECR_REPO_NAME --region $AWS_REGION || true

# 2. Build and push Docker image to ECR
$(aws ecr get-login --no-include-email --region $AWS_REGION)
docker build -t $ECR_REPO_NAME:latest .
docker tag $ECR_REPO_NAME:latest $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG
docker push $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/$ECR_REPO_NAME:$IMAGE_TAG

# 3. Register ECS task definition for Fargate
aws ecs register-task-definition --cli-input-json file://task-def.json

# 4. Create ECS Fargate service
# Ensure that ecs/service-def.json is set up for Fargate
aws ecs create-service --cli-input-json file://service-def.json

echo "Deployment complete."

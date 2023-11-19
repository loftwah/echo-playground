# Infrastructure as Code for Echo Playground

Welcome to the Echo Playground project infrastructure setup! This guide outlines the steps to deploy your environment on AWS using ECS with Fargate in the `ap-southeast-2` region.

## Prerequisites

Before starting, ensure you have the following:

- **AWS CLI:** Configured with the correct credentials.
- **Docker:** For building and pushing the Docker image.
- **Docker Image:** Available in Docker Hub or AWS ECR.

## Infrastructure Components

Our AWS setup includes:

- **ECS (Elastic Container Service):** Using Fargate for serverless task execution.
- **ECR (Elastic Container Registry):** To store Docker images.
- **VPC (Virtual Private Cloud):** Utilizing the default VPC.
- **Subnets:** Using at least two for high availability.
- **Security Group:** Defining ECS task network rules.
- **Application Load Balancer (ALB):** To distribute incoming traffic.

## Step-by-Step Setup

### 1. Build and Push the Docker Image

Navigate to the `ecs` directory and run the `build-push-ecr.sh` script:

```bash
cd ecs
./build-push-ecr.sh
```

This script creates the ECR repository and pushes your Docker image.

### 2. Deploy AWS Infrastructure with Terraform

Navigate to the `terraform` subdirectory:

```bash
cd terraform
```

Execute the following Terraform commands:

```bash
terraform init
terraform plan
terraform apply
```

### 3. Testing and Verification

After Terraform successfully applies:

- **Access the Application**: Visit the ALB DNS URL (`echo-playground-alb-635238929.ap-southeast-2.elb.amazonaws.com`) in your browser.
- **ECS Console Check**: Verify the ECS service and tasks are running correctly.
- **ALB Monitoring**: In AWS Console, check the target group for proper ECS task routing.

### DNS Configuration

Optionally, configure a DNS record to point to the ALB's DNS name for easier access.

## Considerations

- **Security Settings**: Review and maintain security group settings for ALB and ECS.
- **Health Checks**: Regularly monitor the target group's health checks.
- **SSL/TLS**: Utilize AWS Certificate Manager for secure HTTPS communication.

## Notes

- The setup utilizes the default VPC and subnets.
- The ALB ensures the application can handle traffic directed to the ECS service.
- Monitor AWS costs and usage, particularly for Fargate tasks and the ALB.

## Conclusion

Your Echo Playground application is now live on AWS ECS using Fargate. This infrastructure-as-code approach ensures a manageable and scalable environment for your application.

# Infrastructure as Code for Echo Playground

Welcome to the infrastructure setup for the Echo Playground project! Here, we'll guide you through the necessary steps to get your environment up and running on AWS with ECS and Fargate. We're targeting AWS's `ap-southeast-2` region for this setup.

- Todo - Fargate doesn't have a public IP address so I need to add a load balancer to the setup.

## Prerequisites

To make this process smooth, you'll need:

* **AWS CLI:** Configured with the appropriate credentials.
* **Docker:** To build and push the Docker image.
* **A Docker Image:** Available in Docker Hub or AWS ECR.

## Infrastructure Components

Our setup involves a few key AWS components:

* **ECS (Elastic Container Service):** We'll deploy our application as a task in ECS using the Fargate launch type, which abstracts away the underlying server infrastructure.
* **ECR (Elastic Container Registry):** Our Docker image will be stored here. If you're using Docker Hub, you can skip this part.
* **VPC (Virtual Private Cloud):** We'll use the default VPC configured in your AWS account.
* **Subnets:** We need at least two subnets for high availability. These subnets are typically available in your default VPC.
* **Security Group:** We'll create a security group to define the network access rules for our ECS tasks.

## Step-by-Step Setup

### 1. Building and Pushing the Docker Image

Build your Docker image and push it to ECR or Docker Hub. If using ECR, the process typically involves:

```bash
docker build -t echo-playground-prod .
docker tag echo-playground-prod <aws_account_id>.dkr.ecr.ap-southeast-2.amazonaws.com/echo-playground-prod:latest
aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin <aws_account_id>.dkr.ecr.ap-southeast-2.amazonaws.com
docker push <aws_account_id>.dkr.ecr.ap-southeast-2.amazonaws.com/echo-playground-prod:latest
```

### 2. Setting Up the AWS Infrastructure

Run the provided bash script (`setup-aws-infra.sh`) to configure your AWS environment. This script will create a security group and identify the default VPC and subnets for you.

### 3. Updating ECS Configuration Files

With the output from the bash script, update the `ecs/service-def.json` and `ecs/task-def.json` files with the correct VPC, subnet, and security group IDs.

### 4. Deploying to ECS

Finally, deploy your application to ECS using the second bash script (`deploy-to-ecs.sh`), which automates the ECS task and service creation.

## Conclusion

Once you've completed these steps, your Echo Playground application will be live on AWS ECS, leveraging the power and simplicity of Fargate. This setup ensures a scalable, secure, and manageable deployment, letting you focus on the fun part - coding!

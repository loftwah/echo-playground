# Infrastructure as Code for Echo Playground

Welcome to the infrastructure setup for the Echo Playground project! Here, we'll guide you through the necessary steps to get your environment up and running on AWS with ECS and Fargate. We're targeting AWS's `ap-southeast-2` region for this setup.

- Todo - Fargate doesn't have a public IP address so I need to add a load balancer to the setup.

## Prerequisites

To make this process smooth, you'll need:

- **AWS CLI:** Configured with the appropriate credentials.
- **Docker:** To build and push the Docker image.
- **A Docker Image:** Available in Docker Hub or AWS ECR.

## Infrastructure Components

Our setup involves a few key AWS components:

- **ECS (Elastic Container Service):** We'll deploy our application as a task in ECS using the Fargate launch type, which abstracts away the underlying server infrastructure.
- **ECR (Elastic Container Registry):** Our Docker image will be stored here. If you're using Docker Hub, you can skip this part.
- **VPC (Virtual Private Cloud):** We'll use the default VPC configured in your AWS account.
- **Subnets:** We need at least two subnets for high availability. These subnets are typically available in your default VPC.
- **Security Group:** We'll create a security group to define the network access rules for our ECS tasks.

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

## Create an Application Load Balancer

1. **Navigate to the EC2 Dashboard**: Go to the EC2 section in the AWS Management Console.

2. **Load Balancers**: Under the "Load Balancing" section in the navigation pane, click on "Load Balancers".

3. **Create Load Balancer**: Click the “Create Load Balancer” button.

4. **Select Type**: Choose “Application Load Balancer”, then click “Create”.

5. **Configure Load Balancer**:

   - Set a name for your load balancer (e.g., `echo-playground-alb`).
   - Select the VPC and subnets (choose the subnets you are using for your ECS service).
   - Configure the security settings (HTTP/HTTPS).
   - Set up security groups (use the same security group as your ECS tasks or create a new one that allows inbound HTTP/HTTPS).

6. **Configure Routing**: Create a new target group.

   - Set a name (e.g., `echo-playground-tg`).
   - Choose "IP" as the target type.
   - Specify protocol (HTTP), port (the port your application listens on, e.g., 1323), and VPC.

7. **Register Targets**: You can skip this step because ECS will manage the targets.

8. **Review and Create**: Review the settings and create the load balancer.

### Integrate with ECS

1. **Navigate to the ECS Dashboard**: In the AWS Management Console, go to the Elastic Container Service (ECS) section.

2. **Clusters**: Click on “Clusters” and select your existing cluster (`echo-playground`).

3. **Services**: Go to the “Services” tab.

4. **Create or Update Service**:

   - If creating a new service, click “Create” and select the task definition (`echo-playground`).
   - If updating an existing service, select the service and click “Update”.

5. **Configure Network**:

   - In the load balancing section, choose the load balancer type (Application Load Balancer).
   - Select the load balancer name you created (`echo-playground-alb`).
   - Set the container name and port (`1323`).
   - Under “Load balancer listener”, choose the listener port you set up on your ALB (usually 80 or 443).
   - For the target group, select the target group you created (`echo-playground-tg`).

6. **Review and Deploy**: Review your configuration and click “Create Service” or “Update Service”.

### Test and Verify

- Once the service is updated and the tasks are running, go to the ALB’s DNS URL. You should be able to access your application via the load balancer’s DNS name.
- Monitor the target group in the EC2 console to ensure your tasks are healthy and properly registered.

### DNS Configuration

- Manually set up a DNS record in your DNS provider to point to the ALB's DNS name.

### Considerations

- **Security Settings**: Ensure your security groups are correctly configured to allow traffic from the internet (or your specific IP range) to the ALB, and from the ALB to the ECS tasks.
- **Health Checks**: Configure health checks in your target group to ensure that unhealthy tasks are replaced.
- **SSL/TLS**: If you want to use HTTPS, you'll need an SSL certificate. You can request and manage SSL certificates via AWS Certificate Manager.

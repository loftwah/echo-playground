# echo-playground

![Echo Playground](https://github.com/loftwah/echo-playground/assets/19922556/a85c1e83-acd9-4c0d-a960-fd0260b3bbcb)

Learning Echo and Golang. The goal of this project is to build something that will connect to OpenAI's GPT-3.5 Turbo API and generate text based on a prompt. It should be built in Docker and deployed to AWS on ECS.

## Getting Started

### Prerequisites

- Docker
- AWS CLI
- AWS Account

### Installing

1. Clone the repo

   ```bash
   git clone https://github.com/loftwah/echo-playground.git
   ```

2. Build the Docker image

   ```bash
    docker build -t echo-playground .
   ```

3. Run the Docker image

   ```bash
    docker run -p 1323:1323 echo-playground
   ```

4. Open your browser and navigate to [http://localhost:1323](http://localhost:1323)

## Deployment

### Prerequisites

- Docker
- AWS CLI
- AWS Account
- AWS ECR Repository
- AWS ECS Cluster

### Installing

1. Clone the repo

   ```bash
   git clone
   ```

2. Build the Docker image

   ```bash
   docker build -t echo-playground .
   ```

3. Tag the Docker image

   ```bash
   docker tag echo-playground:latest <aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground:latest
   ```

4. Push the Docker image to ECR
   ```bash
   docker push <aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground:latest
   ```
5. Create a new task definition in ECS

   ```bash
   aws ecs register-task-definition --cli-input-json file://task-def.json
   ```

   task-def.json

   ```json
   {
     "family": "echo-playground",
     "containerDefinitions": [
       {
         "name": "echo-playground",
         "image": "<aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground:latest",
         "cpu": 128,
         "memory": 128,
         "essential": true,
         "portMappings": [
           {
             "containerPort": 1323,
             "hostPort": 1323
           }
         ]
       }
     ],
     "requiresCompatibilities": ["EC2"],
     "networkMode": "bridge",
     "cpu": "128",
     "memory": "128"
   }
   ```

6. Create a new service in ECS
   ```bash
   aws ecs create-service --cli-input-json file://service-def.json
   ```
   service-def.json
   ```json
   {
     "cluster": "echo-playground",
     "serviceName": "echo-playground",
     "taskDefinition": "echo-playground",
     "desiredCount": 1,
     "launchType": "EC2",
     "networkConfiguration": {
       "awsvpcConfiguration": {
         "subnets": ["subnet-xxxxxxxx", "subnet-xxxxxxxx"],
         "securityGroups": ["sg-xxxxxxxx"],
         "assignPublicIp": "ENABLED"
       }
     }
   }
   ```

## Built With

- [Echo](https://echo.labstack.com/) - Web framework for Go
- [Golang](https://golang.org/) - Programming language
- [Docker](https://www.docker.com/) - Containerization platform
- [AWS](https://aws.amazon.com/) - Cloud platform

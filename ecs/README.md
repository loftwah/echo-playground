# Infrastructure as Code

For this to work you need. Default region is ap-southeast-2.

- Docker image in Docker Hub or ECR
- 2x subnets
- 1x security group

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
# 🚀 Echo Playground: Where Code Meets Creativity! 🌟

![Echo Playground](https://github.com/loftwah/echo-playground/assets/19922556/a85c1e83-acd9-4c0d-a960-fd0260b3bbcb)

🎉 Welcome to the **Echo Playground**, an adventure in the world of coding where Echo's web framework and OpenAI's GPT-3.5 Turbo API come together in a dazzling display of tech wizardry! 🧙‍♂️

Imagine a place where your words turn into code magic, where AI becomes your brainstorming buddy, and Golang isn't just a language, but a playground for the curious minds. That's what Echo Playground is all about!

## What's Cooking in Echo Playground? 🍳

* **AI-Powered Text Generation:** Throw in a prompt, and watch the AI whip up contextually rich text faster than you can say "abracadabra"! 📜✨
* **Attendance Alchemy:** Our AI isn't just creative; it's got an eye for details too! It magically keeps track of who's attending school and who's not – a bit like a friendly neighborhood watch, but for schools! 🕵️‍♂️🏫
* **Docker & AWS ECS Harmony:** Everything’s neatly packed in Docker containers, floating gracefully in the cloud with AWS ECS. It’s like having your cake and eating it too, but with cloud computing! ☁️🍰
* **Echo Web Framework Wonders:** We’re maneuvering through the Echo web framework with the agility of a ninja, crafting a server that’s not only robust but also as smooth as a jazz tune. 🎷🥋

## Ready for an Adventure? 🚀

Whether you're a code newbie or a seasoned wizard, Echo Playground is your gateway to exploring how AI can transform the way we think about web development. It’s a place to experiment, learn, and maybe even have a little fun along the way. So, grab your hat, bring your wand, and let's dive into a world where code meets creativity, and the possibilities are as limitless as your imagination! 🌈

## Features

- Utilizes the Echo web framework for efficient Go-based web server setup.
- Connects to OpenAI's GPT-3.5 Turbo API for dynamic text generation.
- Processes CSV data to manage and analyze student attendance records.
- Implements detailed error reporting for robust application performance.
- Provides Markdown formatted outputs for user-friendly display.
- Adapts dynamic server port configuration for flexible deployments.

## Getting Started

### Prerequisites

- Docker
- AWS CLI
- AWS Account
- `.env` file with necessary environment variables (e.g., `OPENAI_KEY`)

### Installing and Running Locally

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/loftwah/echo-playground.git
   ```

2. **Build the Docker Image for Development:**

   ```bash
   docker build -f Dockerfile.dev -t echo-playground-dev .
   ```

3. **Run the Docker Container:**

   ```bash
   docker run -p 1323:1323 --env-file .env echo-playground-dev
   ```

   > **Note:** You should probably use Docker Compose to make things simpler.

4. **Run the Docker Container with Docker Compose:**

   ```bash
   docker compose up
   ```

5. **Access the Application:** Open your browser and navigate to <http://localhost:1323>.

### Deployment on AWS ECS

> **Note:** This guide assumes you have an AWS ECR repository set up and AWS CLI configured with the necessary credentials. Eventually, we aim to automate this process using CI/CD with GitHub Actions.

#### 1. **Building and Tagging the Docker Image for Production**

First, we build the Docker image optimized for production. This image is a lean version of our application, free from development tools and configurations, ensuring efficiency and security in a production setting.

```bash
docker build -t echo-playground-prod .
docker tag echo-playground-prod <aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground-prod:latest
```

Replace `<aws_account_id>` and `<region>` with your AWS account ID and region, respectively. This step tags the built image, making it ready for deployment in our AWS environment.

#### 2. **Pushing the Docker Image to AWS ECR**

Next, we push the tagged image to AWS ECR (Elastic Container Registry), a Docker container registry that makes it easy for us to manage, store, and deploy Docker container images.

```bash
docker push <aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground-prod:latest
```

This command uploads our production-ready Docker image to the cloud, making it accessible for deployment on AWS ECS.

#### 3. **Continuing with ECS Deployment**

With the image now in ECR, we proceed to deploy it on AWS ECS (Elastic Container Service), a highly scalable and high-performance container orchestration service.

* **Create a Task Definition in ECS:** We define a task in `ecs/task-def.json`, which includes our container configuration, resource allocation (CPU and memory), and network settings. It serves as a blueprint for how our Docker container should run on ECS.
* **Set Up the ECS Service:** The `ecs/service-def.json` configures the ECS service. Here, we specify the cluster, service name, and network configuration. The service ensures our application is maintained at the desired state, handling tasks like starting new instances if needed and managing public IP assignment.
* **Deploy Using ECS:** Finally, we use the AWS CLI to register the task definition and create the service in ECS. This process tells ECS how to run our application, where to run it, and how to manage networking aspects.

```bash
aws ecs register-task-definition --cli-input-json file://ecs/task-def.json
aws ecs create-service --cli-input-json file://ecs/service-def.json
```

These commands set everything in motion, deploying our Echo Playground application in the AWS cloud, ready to serve users at scale.

## Built With

- [Echo](https://echo.labstack.com/) - Web framework for Go
- [Golang](https://golang.org/) - Programming language
- [Docker](https://www.docker.com/) - Containerization platform
- [AWS](https://aws.amazon.com/) - Cloud platform

## Project Roadmap and Checklist

The following checklist outlines the planned enhancements and features to be implemented:

- Optimize CSV file processing for efficiency.
- Refine complex data handling and logic.
- Enhance Markdown output formatting.
- Upgrade and maintain OpenAI service layer.
- Expand detailed error reporting.
- Integrate additional middleware for security and request management.
- Implement secure authentication and authorization system.
- Integrate databases to bring users existing data into the mix.
- Integrate advanced logging and monitoring tools.
- Develop a comprehensive suite of unit and integration tests.
- Consider developing a front-end interface or API endpoints.
- Implement asynchronous processing capabilities.

## Contributing

Contributions to the Echo Playground project are welcome. Please ensure to follow the project's code of conduct and contribution guidelines.

## License

This project is licensed under the [MIT License](LICENSE). See the LICENSE file for details.

## Acknowledgments

- Thanks to the Echo framework team for their excellent web framework.
- Appreciation to OpenAI for providing the GPT-3.5 Turbo API.
- Gratitude to all

### Notes

1. **Development and Production Docker Builds:** The instructions now differentiate between building a development image (`Dockerfile.dev`) and a production image (`Dockerfile`). Use Docker Compose for development to make it simpler. Production instructions are in the `/ecs` directory.

2. **Environment Variables:** I didn't actually know this before but to run a Docker container with a `.env` file present include the `--env-file .env` flag to ensure that your environment variables are correctly passed into the container. This is probably why all of my ECS stuff failed for ages.

3. **Deployment Steps:** The AWS ECS deployment steps remain as previously described, assuming the production setup doesn’t rely on the `.env` file.

Remember to keep the `.env` file secure and not include any sensitive data in your version control.

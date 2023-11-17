# Echo Playground

![Echo Playground](https://github.com/loftwah/echo-playground/assets/19922556/a85c1e83-acd9-4c0d-a960-fd0260b3bbcb)

The Echo Playground project is an educational endeavor aimed at exploring the capabilities of the Echo web framework in Golang, particularly in conjunction with OpenAI's GPT-3.5 Turbo API. The primary objective is to create an application that can generate contextually relevant text based on user prompts, showcasing the integration of AI in web development. This project is designed for Docker deployment and is optimized for AWS ECS.

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

### Installing and Running Locally

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/loftwah/echo-playground.git
   ```

2. **Build the Docker Image:**

   ```bash
   docker build -t echo-playground .
   ```

3. **Run the Docker Container:**

   ```bash
   docker run -p 1323:1323 echo-playground
   ```

4. **Access the Application:**

   Open your browser and navigate to <http://localhost:1323>.

5. **View Output in User-Friendly Format:**

   ```bash
   # Fetches and formats data from the application
   # ... Bash commands as in your original README ...
   ```

## Deployment on AWS ECS

### Prerequisites

- Docker
- AWS CLI
- AWS Account
- AWS ECR Repository
- AWS ECS Cluster

### Steps

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/loftwah/echo-playground.git
   ```

2. **Build and Tag the Docker Image:**

   ```bash
   docker build -t echo-playground .
   docker tag echo-playground:latest <aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground:latest
   ```

3. **Push the Docker Image to AWS ECR:**

   ```bash
   docker push <aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground:latest
   ```

4. **Register a Task Definition in ECS:**

   ```bash
   aws ecs register-task-definition --cli-input-json file://task-def.json
   ```

5. **Create a New Service in ECS:**

   ```bash
   aws ecs create-service --cli-input-json file://service-def.json
   ```

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
- Transition to a database system from CSV file processing.
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

# 🚀 Echo Playground: Where Code Meets Creativity! 🌟

![Echo Playground](https://github.com/loftwah/echo-playground/assets/19922556/a85c1e83-acd9-4c0d-a960-fd0260b3bbcb)

🎉 Welcome to the **Echo Playground**, an adventure in the world of coding where Echo's web framework and OpenAI's GPT-3.5 Turbo API come together in a dazzling display of tech wizardry! 🧙‍♂️

Imagine a place where your words turn into code magic 🪄, where AI becomes your brainstorming buddy, and Golang isn't just a language, but a playground for the curious minds. That's what Echo Playground is all about!

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

4. **Access the Application:** Open your browser and navigate to <http://localhost:1323>.

### Deployment on AWS ECS

1. **Build and Tag the Docker Image for Production:**

   ```bash
   docker build -t echo-playground-prod .
   docker tag echo-playground-prod <aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground-prod:latest
   ```

2. **Push the Docker Image to AWS ECR:**

   ```bash
   docker push <aws_account_id>.dkr.ecr.<region>.amazonaws.com/echo-playground-prod:latest
   ```

3. **Continue with ECS Deployment:** Follow steps 4-5 under the 'Deployment on AWS ECS' section as previously described.

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

### Notes

1. **Development and Production Docker Builds:** The instructions now differentiate between building a development image (`Dockerfile.dev`) and a production image (`Dockerfile`).

2. **Environment Variables:** Instructions for running the Docker container now include the `--env-file .env` flag to ensure that your environment variables are correctly passed into the container.

3. **Deployment Steps:** The AWS ECS deployment steps remain as previously described, assuming the production setup doesn’t rely on the `.env` file.

Remember to keep the `.env` file secure and not include any sensitive data in your version control.

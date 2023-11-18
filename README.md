# 🚀 Echo Playground: Where Code Meets Creativity! 🌟

![Echo Playground](https://github.com/loftwah/echo-playground/assets/19922556/a85c1e83-acd9-4c0d-a960-fd0260b3bbcb)

🎉 Welcome to the **Echo Playground**, an adventure in the world of coding where Echo's web framework and OpenAI's GPT-3.5 Turbo API come together in a dazzling display of tech wizardry! 🧙‍♂️

Imagine a place where your words turn into code magic, where AI becomes your brainstorming buddy, and Golang isn't just a language, but a playground for the curious minds. That's what Echo Playground is all about!

## What's Cooking in Echo Playground? 🍳

- **AI-Powered Text Generation:** Throw in a prompt, and watch the AI whip up contextually rich text faster than you can say "abracadabra"! 📜✨
- **Attendance Alchemy:** Our AI isn't just creative; it's got an eye for details too! It magically keeps track of who's attending school and who's not – a bit like a friendly neighborhood watch, but for schools! 🕵️‍♂️🏫
- **Docker & AWS ECS Harmony:** Everything’s neatly packed in Docker containers, floating gracefully in the cloud with AWS ECS. It’s like having your cake and eating it too, but with cloud computing! ☁️🍰
- **Echo Web Framework Wonders:** We’re maneuvering through the Echo web framework with the agility of a ninja, crafting a server that’s not only robust but also as smooth as a jazz tune. 🎷🥋

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

See [ecs/README.md](ecs/README.md) for detailed instructions on deploying the application on AWS ECS.

## Built With

- [Echo](https://echo.labstack.com/) - Web framework for Go
- [Golang](https://golang.org/) - Programming language
- [Docker](https://www.docker.com/) - Containerization platform
- [AWS](https://aws.amazon.com/) - Cloud platform

## Project Roadmap and Checklist

| Priority | Feature                                                             | Score | Description                                                                                      |
| -------- | ------------------------------------------------------------------- | ----- | ------------------------------------------------------------------------------------------------ |
| 1        | Upgrade and Maintain OpenAI Service Layer                           | 2     | Central to the project, ensuring AI components are up-to-date and functional.                    |
| 2        | Implement Secure Authentication and Authorization                   | 2     | Crucial for data privacy and security, especially with sensitive student data.                   |
| 3        | Develop a Comprehensive Suite of Unit and Integration Tests         | 2     | Critical for application reliability and stability.                                              |
| 4        | Optimize CSV File Processing for Efficiency                         | 3     | Vital for effectively handling and analyzing the enriched CSV data.                              |
| 5        | Refine Complex Data Handling and Logic                              | 2     | Important for efficiently processing the new, more complex dataset.                              |
| 6        | Implement Asynchronous Processing Capabilities                      | 3     | Enhances performance for data-intensive operations and user experience.                          |
| 7        | Integrate Advanced Logging and Monitoring Tools (e.g., Prometheus)  | 3     | Key for long-term maintenance, scalability, and real-time insights into application performance. |
| 8        | Expand Detailed Error Reporting                                     | 3     | Essential for improving application quality through effective debugging.                         |
| 9        | Integrate Additional Middleware for Security and Request Management | 3     | Enhances the application's security and efficiency in managing requests.                         |
| 10       | Enhance Markdown Output Formatting                                  | 4     | Improves user experience, especially in presenting data and insights.                            |
| 11       | Consider Developing a Front-End Interface or API Endpoints          | 4     | Increases usability and accessibility; significant development effort needed.                    |

### Additional Features for Consideration:

| Priority | Feature                                         | Score | Description                                                                                       |
| -------- | ----------------------------------------------- | ----- | ------------------------------------------------------------------------------------------------- |
| 1        | AI-Driven Personalized Learning Recommendations | 2     | Use AI to create tailored learning recommendations for each student based on CSV data.            |
| 2        | Automated Student Wellbeing Checks              | 2     | Leverage AI to assess student wellbeing regularly using engagement and attendance data from CSV.  |
| 3        | Behavioral Pattern Analysis Tool                | 2     | Analyze and report on patterns in student behavior and its impact on performance using CSV data.  |
| 4        | Interactive Student Engagement Platform         | 3     | Develop an interactive platform for students to engage with AI-driven learning materials.         |
| 5        | Parent-Teacher Communication Enhancer           | 3     | Implement tools to facilitate and enhance communication between parents and teachers.             |
| 6        | Emotional Wellbeing Index Analysis              | 3     | Analyze and interpret the emotional wellbeing index for student support and counseling.           |
| 7        | Technology Proficiency Improvement Programs     | 3     | Create programs to enhance students' technology proficiency, using data for personalized content. |
| 8        | Social Engagement Activity Planner              | 4     | Design activities and programs to boost social engagement among students.                         |
| 9        | Comprehensive Student Achievement Dashboard     | 4     | Develop a dashboard to display comprehensive student achievement and progress data.               |

This roadmap prioritizes crucial updates and enhancements while incorporating new features that leverage the enriched dataset to support and enhance the student learning experience.

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

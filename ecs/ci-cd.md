# GitHub Actions CI/CD for Echo Playground

This guide walks you through setting up a Continuous Integration (CI) and Continuous Deployment (CD) pipeline using GitHub Actions for the Echo Playground project. Our CI/CD pipeline automates the process of building, testing, and deploying your application to AWS ECS using Fargate.

## What This GitHub Action Does

The GitHub Action defined in `.github/workflows/main.yml` performs the following tasks:

1. **Trigger:** It triggers automatically on every push to the `main` branch of the repository.
2. **Build Docker Image:** Builds a Docker image from your project's Dockerfile.
3. **Push to AWS ECR:** Tags and pushes the built image to your AWS Elastic Container Registry (ECR).
4. **Update ECS Task Definition:** Updates the Amazon ECS task definition with the new image URI.
5. **Deploy to ECS:** Deploys the updated task definition to your specified Amazon ECS service.

## Prerequisites

* **AWS Account:** Ensure you have an AWS account with ECR and ECS configured.
* **GitHub Repository:** The Echo Playground code should be hosted in a GitHub repository.
* **AWS IAM User:** An IAM user with permissions to push images to ECR and update ECS services.

## Setup Guide

### Step 1: Configure AWS Credentials in GitHub

1. **Generate AWS Credentials:**

   * Create a new IAM user in AWS with programmatic access.
   * Attach policies that grant permission to interact with ECR and ECS.

2. **Add Secrets to GitHub:**

   * Go to your GitHub repository.

   * Navigate to `Settings` > `Secrets`.

   * Click on `New repository secret` and add the following secrets:

     * `AWS_ACCESS_KEY_ID`: The access key ID of the IAM user.
     * `AWS_SECRET_ACCESS_KEY`: The secret access key of the IAM user.

### Step 2: Update the GitHub Actions Workflow

1. **Edit the Workflow File:**

   * Locate the `.github/workflows/main.yml` file in your repository.
   * Ensure the file contains the correct workflow as described in the section above.

2. **Customize for Your AWS Setup:**

   * Modify the environment variables in the workflow file to match your AWS ECR repository and ECS service names.

### Step 3: Push Changes to Trigger the Workflow

* Push a change to the `main` branch in your GitHub repository.
* This will trigger the GitHub Actions workflow.
* You can monitor the progress under the `Actions` tab in your GitHub repository.

### Step 4: Verify Deployment

* Once the workflow completes, check your AWS ECS service.
* Ensure the new version of your application is deployed and running.

## Understanding the Workflow

* **Build and Push Image:** The workflow builds a Docker image and pushes it to AWS ECR, ensuring that the latest version of your application is stored in a secure container registry.
* **Update and Deploy ECS Task:** By updating the ECS task definition and redeploying the service, your latest code changes are automatically reflected in the running application on AWS ECS.
* **Automation:** This CI/CD pipeline reduces the risk of manual errors and streamlines your deployment process.

## Conclusion

With this GitHub Actions setup, you have an automated pipeline that ensures every change to your main branch is built, tested, and deployed seamlessly to AWS ECS. This automation enhances the reliability and efficiency of your development and deployment processes.
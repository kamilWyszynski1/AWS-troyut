# AWS-trout

## Purpose
Repository is created in order to learn AWS using terraform, docker and kubernetes

## Steps
* [x] Github Action for running tests
* [x] Create a simple app, create Github Action that pushes image to ECR after merge to master


## Github action deploys app
1. `aws configure` in order to initialize AWS CLI connection
2. `aws ecr create-repository --repository-name <repository_name> --region <region_name>`

# Aquafarm Management Applications REST API

## Overview
Developing an Aquafarm Management Application API using Go Fiber, GORM, Postgres, and Docker as storage systems. I successfully implemented CRUD operations for farms and ponds within its relationship model. In addition, I also built a tracking system for API usage statistics for each API request.

## Architecture Model
![Architecture](https://github.com/Wordyka/InternDelos-Wordyka/blob/8bd2105dee2aafad7cf2556ece94c313ddcb1c47/img/Architecture.png)

## Entity Relationship Diagram 
![Acces this link](https://github.com/Wordyka/InternDelos-Wordyka/blob/fbd6a4e7f41ae216bf63651d445b304e86a11348/img/ERD%20Aquafarm.png)

### Rows Data Example
- Farms
![Farms](https://github.com/Wordyka/InternDelos-Wordyka/blob/8bd2105dee2aafad7cf2556ece94c313ddcb1c47/img/Farms%20Rows%20Data.png)
  
- Ponds
![Ponds](https://github.com/Wordyka/InternDelos-Wordyka/blob/8bd2105dee2aafad7cf2556ece94c313ddcb1c47/img/Ponds%20Rows%20Data.png)

## Prerequisites
Before you can run the application, make sure you have the following prerequisites installed:

- [Go (version 1.19 or later)](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [API Platform: Postman](https://www.postman.com/downloads/)

## Environment Setup
- Install Go: You can download and install Go from the official website
- Set up your Go environment if you haven't already, including configuring your `GOPATH` and `PATH` variables.
- Make sure Golang is installed on your system by running `go version` to check the go version
- Make sure Docker is installed on your system by running `docker --version` to check the go version

## Running The Docker
- Execute the command `docker compose build` to build the Docker images for all services defined in docker-compose.yml file. 
- Execute the command `docker compose up` to start the containers.
- Use the API with base URL `http://127.0.0.1:3000/v1`

## API Documentation
You can access the Postman API Documentation on this link :   
[Postman API Documentation](https://documenter.getpostman.com/view/19084112/2s9Ye8hFjd)

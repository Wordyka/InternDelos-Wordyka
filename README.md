# Aquafarm Management Applications REST API

## Overview
Developing an Aquafarm Management Application API using Go Fiber, GORM, and Postgres. I successfully implemented CRUD operations for farms and ponds within its relationship model. In addition, I also built a tracking system for API usage statistics for each API request.

## Entity Relationship Diagram 
[Acces this link](https://drive.google.com/file/d/1NOf7G4SRdG0rgQ5bteL-i-9aJfn4WqF-/view?usp=sharing)

## Prerequisites
Before you can run the application, make sure you have the following prerequisites installed:

- [Go (version 1.19 or later)](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [API Platform : Postman](https://www.postman.com/downloads/)

## Environment Setup
- Install Go : You can download and install Go from the official website
- Set up your Go environment if you haven't already, including configuring your `GOPATH` and `PATH` variables.
- Make sure Golang is installed on your system by running `go version` to check the go version
- Make sure Docker is installed on your system by running `docker --version` to check the go version

## Running The Docker
- Execute command `docker compose build` to build the Docker images for all services defined in docker-compose.yml file. 
- Execute command `docker compose up` to start the containers.

## API Documentation
You can acces the Postman API Documentation on this link :   
[Postman API Documentation](https://documenter.getpostman.com/view/19084112/2s9Ye8hFjd)

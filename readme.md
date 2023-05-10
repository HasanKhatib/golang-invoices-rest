# Rest Service

[![Go Report Card](https://goreportcard.com/badge/github.com/hasankhatib/golang-invoices-rest)](https://goreportcard.com/report/github.com/hasankhatib/golang-invoices-rest)
[![Build Status](https://github.com/hasankhatib/golang-invoices-rest/workflows/Build/badge.svg)](https://github.com/hasankhatib/golang-invoices-rest/actions)
[![Docker Build](https://img.shields.io/docker/cloud/build/hasankhatib/golang-invoices-rest)](https://hub.docker.com/r/hasankhatib/golang-invoices-rest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


## Exercise Description

This project is a RESTful service built with Go. It serves as an exercise to learn and practice Go programming language.

## Prerequisites

Before running this application, ensure that you have the following prerequisites installed:

- Go (version 1.17 or higher)
- PostgreSQL database

## Getting Started

To run the application locally, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/hasankhatib/golang-invoices-rest.git
   cd golang-invoices-rest
    ```
2. Install the dependencies:

    ```shell
    go mod download
    ```

3. Set up the database:
    - Create a PostgreSQL database.
    - Update the database configuration in the configs/config.yaml file.

4. Build and run the application:

    ```shell
    go run cmd/invoiceserver/main.go
    ```
The application will be accessible at http://localhost:8080.

## Running with Docker Compose

To run the application using Docker Compose, follow these steps:

1. Make sure you have Docker and Docker Compose installed on your machine.

2. Open the `docker-compose.yml` file and ensure that the configuration matches your requirements. You can modify the port mappings or environment variables if needed.

3. Run the following command to start the application and the PostgreSQL database:

   ```bash
   docker-compose up


## Docker
You can also run the application using Docker. Follow these steps:

1. Build the Docker image:
    ```shell
    docker build -t hasankhatib/golang-invoices-rest .
    ``` 
2. Run the Docker container:
    ```shell
    docker run -p 8080:8080 -d hasankhatib/golang-invoices-rest
    ```
The application will be accessible at http://localhost:8080.


## License
This project is licensed under the [MIT License](LICENSE).





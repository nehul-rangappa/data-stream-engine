# data-stream-engine
This is a real-time data streaming engine with a Kafka consumer service developed using GoLang

## Technology
* Language: Go v1.22
* Dependency: Kafka
* Design Pattern: Layered Architecture

## Features
* Event Streaming Consumer
* Process and filter data from event stream 
* Produce processed data as per requirements to a newly configured topic
* Perform Data analysis and give insights on the data

## Design Choices
* Go is selected for building this service considering its various advantages over other languages like python which are listed below
    1. Support of Concurrency
    2. Spped and Efficiency
    3. Optimized memory management
    4. Lightweight goroutines
    5. Suitable for microservices and event streaming pipelines

* Layered Architecture is used considering separation of concerns and makes every layer responsible for its own functionality.
* Kafka consumer is used for reading messages constantly from a dedicated topic and producer is used to send processed messages to another topic.
* Docker is used to containerize our services and not have a hard dependency on physical requirements.

## Data Flow
Based on the layered architecture design, our handler is responsible for listening to events on a dedicated topic, service layer is responsible for processing, filtering and performing analysis with aggregations. Finally, the handler is again responsible for writing the processed data to a kafka topic.


## Supporting details
* Our application can be easily deployed in production with the help of docker to run all the services in containers. An orchestrator like Kubernetes can be used to manage the containers. On top of this, we can constantly monitor our services and the resources using tools such as Grafana and Prometheus.

* Some components that I feel we can add further to make the product better and production ready are
    1. Data Storage
    2. Secure services
    3. Load balancer to manage and route the incoming requests across multiple replicas
    4. Add a health API for the service
    5. Set up alerts in case of service breakdowns or potential errors.
    6. Better usage of Concurrency in Go

* Kafka is inherently scalable and Go makes sure to handle large traffic with the asynchronous processing in microservices developed that are also capable of independent scaling.
* As the dataset grows, our Kubernetes can be configured to scale our services horizontally.
* We can use the partitions effectively in each topic helps distribute the load across multiple consumers
* Our service is fault-tolerant by constantly listening to events with an efficient retry mechanism such that the service never stops and exits without observing the streaming pipeline. Also, proper logs are incorporated for monitoring and troubleshooting purposes.

## Project Setup
* Clone the repository
* Install the necessary requirements stated in technology section
* Change the environment variables in .env as needed
* Set up the local environment by spinning up all the docker containers for setting up Kafka, Zookeper and data generator using `docker-compose up`
* Run the application using `go run main.go`


## Project Structure
├── handler\
│ ├── kafka.go\
├── model\
│ ├── message.go\
├── service\
│ ├── kafka.go\
├── utils\
│ ├── kafka.go\
├── main.go\
├── dockerfile\
├── docker-compose.yml\
├── .env\
├── README.md


## Future Enhancements
* Support API to manage the event streaming service.
* Extend the layered architecture to have a repository layer that handles interaction with kafka topics

# Kafka Go Example Apps

## Project structure:
```
kafka-go-example-apps/
├── bin/                         # Compiled binaries for producer and consumer
├── cmd/                         # Contains the applications
│   ├── producer/                # Producer application
│   │   └── main.go              # Main entry point for the producer
│   └── consumer/                # Consumer application
│       └── main.go              # Main entry point for the consumer
├── go.mod (should be committed) # Go module file, defines dependencies
├── go.sum (should be committed) # Dependency lock file
├── README.md                    # This file, contains project documentation
```

## Build Instructions:
To build:
* ```go build -o bin/producer ./cmd/producer```
* ```go build -o bin/consumer ./cmd/consumer```

## Usage Instructions:

### Producer:
Example command to run the producer:
* ```./bin/producer --broker=localhost:19092 --topic=demo-topic```

Full usage instructions:
```
Usage of ./bin/producer:
  -broker string
        Kafka broker address (default "localhost:19092")
  -messages int
        Number of messages to produce (default 100)
  -pause int
        Pause in milliseconds between messages (default 1000)
  -topic string
        Kafka topic (default "demo-topic")
```

### Consumer:
Example command to run the consumer:
* ```./bin/consumer --broker=localhost:9092 --topic=demo-topic --group=demo-group```

Full usage instructions:
```
Usage of ./bin/consumer:
  -broker string
        Kafka broker address (default "localhost:19092")
  -group string
        Kafka consumer group (default "demo-group")
  -topic string
        Kafka topic (default "demo-topic")
```

## Dependency management:
* Run `go mod tidy` to ensure go.mod and go.sum are up to date.
* go.sum ensures reproducible builds and should be committed to version control.
* To update all dependencies: `go get -u ./...`

# kafka-sandbox

Repo to play around with Kafka a bit.

## Folder Structure

```
kafka-sandbox
├── apps                            # Contains example applications using Kafka
│   ├── kafka-go-example-apps       # Example apps using Kafka Go client
|   ...
├── cluster-setup                   # Contains setup scripts for different Kafka clusters
│   └── local                       # Local setup for Kafka cluster
│       └── docker-compose          # Local setup using Docker Compose
│           ├── docker-compose.yaml # Docker Compose file to set up a local Kafka cluster
│           └── README.md           # Instructions for local setup
├── README.md                       # Main README file for the repository (this file)
└── tools                           # Contains tools and utilities for working with Kafka
```

## Recommended Tools

* [Offset Explorer](https://www.kafkatool.com/) - GUI for Kafka

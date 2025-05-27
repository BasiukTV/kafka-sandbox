# Local Kafla Cluster with Docker Compose

## Prerequisites

* [Docker](https://docs.docker.com/get-started/introduction/get-docker-desktop/) - Install Docker Desktop
* (Optional) [Offset Explorer](https://www.kafkatool.com/) - GUI for Kafka

## Cluster Startup
Run from the current directory:
```bash
docker compose up -d
docker compose logs -f kafka1 # wait for "Kafka server started"
```

## Testing Kafka

* Create a test topic (3 partitions, RF=3)
```bash
docker exec -it kafka1 \
  /opt/kafka/bin/kafka-topics.sh \
  --create --topic demo \
  --partitions 3 --replication-factor 3 \
  --bootstrap-server kafka1:9092
```

* List all topics
```bash
docker exec -it kafka1 \
  /opt/kafka/bin/kafka-topics.sh --list --bootstrap-server kafka1:9092
```

* Describe the topic
```bash
docker exec -it kafka1 \
  /opt/kafka/bin/kafka-topics.sh \
  --describe --topic demo --bootstrap-server kafka1:9092
```

* Produce messages to the topic
```bash
docker exec -it kafka1 \
  /opt/kafka/bin/kafka-console-producer.sh \
  --topic demo --bootstrap-server kafka1:9092
```

* Consume messages from the using another broker
```bash
docker exec -it kafka2 \
  /opt/kafka/bin/kafka-console-consumer.sh \
  --topic demo --from-beginning --bootstrap-server kafka2:9092
```

## Connecting to Kafka

### Using Offset Explorer
1. Open Offset Explorer
2. Add a new connection
3. Set the connection name (e.g., `local-kafka`)
4. Set the bootstrap server to `localhost:19092`
5. Set the security protocol to `PLAINTEXT`
6. Click "Test" to verify the connection
7. Click "OK" to save the connection
8. You can now browse topics, produce and consume messages

## Cluster Shutdown
Run from the current directory:
```
docker compose down
```

## Cluster Cleanup
Below command will stop the cluster and remove all data volumes and orphaned containers:
```bash
docker compose down --volumes --remove-orphans
```

# learn-kafka

## steps

```bash
# start kakfa
docker run -d --name broker -p 9092:9092 apache/kafka:latest

# attach broker
docker exec --workdir /opt/kafka/bin/ -it broker sh

# create topic
./kafka-topics.sh --bootstrap-server localhost:9092 --create --topic test-topic
```

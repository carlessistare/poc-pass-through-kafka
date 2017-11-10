# GO POC with Docker-Protobuf-KafkaProducer

## Package dependencies

```bash
go get golang.org/x/sys/unix
go get github.com/tools/godep
PATH=$PATH:`pwd`/../../bin
godep save
```

## Build Docker and Run

```bash
docker build -t ws-test .
docker run -p 8001:8001 -ti --rm ws-test
```

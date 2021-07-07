# udp2mysql
Listen UDP and save packet to mysql or mariadb

# gRPC

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/udp2mysql/v1/udp2mysql.proto
```

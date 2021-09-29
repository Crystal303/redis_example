# redis_example

## docker run
```shell
docker pull redis
docker run --name myredis -d -p 6379:6379 redis
docker exec -it myredis redis-cli
```
# list running processes in specific container

```
docker run --name mongo -d mongo
```

```
docker top mongo
```

# show me all running processes

```
ps aux
```

```
ps aux | grep mongo
```
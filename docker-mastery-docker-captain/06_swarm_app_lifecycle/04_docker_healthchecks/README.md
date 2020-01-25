# Docker healthcheck

- basic command using default options
```
HEALTHCHECK curl -f http://localhost/ || false
```

- custom options with the command
```
HEALTHCHECK --timeout=2s --interval=3s --retries=3 \
CMD curl -f http://localhost/ || exit 1
```

## example using elastic search
```
docker run \
--health-cmd="curl -f localhost:9200/_cluster/health || false" \
--health-interval=5s \
--health-retries=3 \
--health-timeout=2s \
--health-start-period=15s \
elasticsearch:2
```

## container example
- run without healthcheck
```
docker container run --name p1 -d postgres
```

- run with healthcheck
```
docker container run --name p2 -d --health-cmd="pg_isready -U postgres || exit 1" postgres
```

- see the STATUS healthy
```
docker container ls
```

- see the Health
```
docker container inspect p2
```

## service example
- create a service without healthcheck
```
docker service create --name p1 postgres
```

- create service with healthcheck
```
docker service create --name p2 --health-cmd="pg_isready -U postgres || exit 1" postgres
```

## other options

```
--interval=DURATION(default:30s)
--timeout=DURATION(default:30s)
--start-period=DURATION(default: 0s) (17.09+)
--retries=N(default: 3)
```
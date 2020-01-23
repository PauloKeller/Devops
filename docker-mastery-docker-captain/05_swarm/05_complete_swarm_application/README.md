# Complete voting app build on command line

- frontend network
```
docker network create -d overlay voting-app-frontend
```

- backend network
```
docker network create -d overlay voting-app-backend
```

- create vote app service with 3 replicas
```
docker service create --name vote --replicas 3 --network voting-app-frontend -p 80:80 bretfisher/examplevotingapp_vote
```

- create redis database service with 1 replicas
```
docker service create --name redis --network voting-app-frontend redis:3.2
```

- create worker service with 1 replicas
```
docker service create --name worker --network voting-app-backend --network voting-app-frontend bretfisher/examplevotingapp_worker:java
```

- create db database service and volume with 1 replicas
```
docker service create --name db --network voting-app-backend --mount type=volume,source=db-data,target=/var/lib/postgresql/data postgres:9.4
```

- create result service just for admins with 1 replicas
```
docker service create --name result --network voting-app-backend -p 5001:80 bretfisher/examplevotingapp_result
```

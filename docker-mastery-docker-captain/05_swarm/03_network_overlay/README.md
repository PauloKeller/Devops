# Swarm Network 

- create a new overlay network
```
docker network create --driver overlay mydrupal
```

- create a new service postgres with overlay network
```
docker service create --name psql --network mydrupal -e POSTGRES_PASSWORD=mypass postgres
```

- create a new service drupal with overlay network
```
docker service create --name drupal --network mydrupal -p 80:80 drupal
```
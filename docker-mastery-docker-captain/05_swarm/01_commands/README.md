# SWARM commands

- init swarm

```
docker swarm init
```

- list nodes

```
docker node ls
```

- create a node with alpine image (service replaces run)

```
docker service create <image-name>
```

```
docker service create alpine ping 8.8.8.8
```

- show the containers in a service

```
docker service ps <service-name>
```

- scale a service

```
docker service update <service-id> --replicas 3 
```

- remove a service

```
docker service rm <service-name> 
```
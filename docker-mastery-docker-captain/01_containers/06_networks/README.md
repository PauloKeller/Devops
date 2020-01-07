# Networks on containers

- run container on default port 80 and expose 80
```
docker container run -p <port>:<expose-port> --name webhost -d nignx
```

```
docker container run -p 80:80 --name webhost -d nignx
```

- show container port
```
docker container port <container>
```

- inspect and format ipaddres on container
```
docker container inspect --format '<property>' webhost
```

```
docker container inspect --format '{{ .NetworkSettings.IPAddress }}' webhost
```

- list all networks
```
docker network ls
```

- create network
```
docker network create --driver <network-name>
```

```
docker network create my_app_net
```

- run container under a network
```
docker container run -d --name new_nginx --network <network-name> nginx
```

```
docker container run -d --name new_nginx --network my_app_net nginx
```

- connect / disconnect to a network

```
docker network connect <network-id> <container-id>
```

```
docker network disconnect <network-id> <container-id>
```
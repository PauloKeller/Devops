# run docker container

```
docker container run --publish <port> <image>
```

```
docker run --publish 80:80 nginx
```

# start docker container 

```
docker container start <image>
```

```
docker start nginx
```

# run in detach mode 

```
docker container run --publish 80:80 --detach <image>
```

```
docker container run --publish 80:80 -d nginx
```

# list all containers 

- running
```
docker ps
```

```
docker container ls
```

- not running
```
docker container ls -a
```

# stop running container

```
docker container stop <id>
```

```
docker container stop nginx
```

# name a container 

```
docker container run --publish 80:80 --detach --name <name> <image>
```

```
docker container run --publish 80:80 -d --name webhost nginx
```
# see container logs

```
docker container logs <container-name>
```

```
docker container logs webhost
```

# see container running processes

```
docker container top <container-name>
```

```
docker container top webhost
```

# remove a container

- not running 
```
docker container rm <container-id>
```

- running
```
docker container rm -f <container-id>
```

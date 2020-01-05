# networks on containers

```
docker container run -p 80:80 --name webhost -d nignx
```

```
docker container port webhost
```

```
docker container inspect --format '{{ .NetworkSettings.IPAddress }}' webhost
```
# Volumes commands

- list volumes
```
docker volume ls
```

- inspect volumes
```
docker volume inspect <volume-id>
```
- create a named volume
```
docker container run -d --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=True -v <volume-name>:<volume-path> mysql 
```

```
docker container run -d --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=True -v mysql-db:/var/lib/mysql mysql 
```

- create volume runtime
```
docker volume create <volume-name>
```
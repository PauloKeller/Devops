# run some containers

```
docker container run -d -p --name db -e MYSQL_RANDOM_ROOT_PASSWORD=yes mysql
```

```
docker container run -d --name proxy -p 80:80 nginx
```

# process running on containers

```
docker container top mysql
```

```
docker container top nginx
```

# container metada

```
docker container inspect mysql
```

```
docker container inspect nginx
```
# container status

```
docker container stats
```
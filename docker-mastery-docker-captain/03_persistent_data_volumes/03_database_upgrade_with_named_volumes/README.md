# Database upgrade

- create name volume with old version of psql
```
docker container run -d --name psql -v psql:/var/lib/postgresql/data postgres:9.6.1
```

- keep watching it run
```
docker container logs -f psql
```

- update psql with new container
```
docker container run -d --name psql2 -v psql:/var/lib/postgresql/data postgres:9.6.2
```
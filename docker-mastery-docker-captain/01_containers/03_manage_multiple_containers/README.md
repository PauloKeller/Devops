# mysql container 

```
docker container run -d -p 3306:3306 --name db -e MYSQL_RANDOM_ROOT_PASSWORD=yes mysql
```

- find random password
```
docker container logs db 
```

- GENERATED ROOT PASSWORD: <password>

# httpd container (apache)

```
docker container run -d --name webserver -p 8080:80 httpd
```

# nginx container (proxy)

```
docker container run -d --name proxy -p 80:80 nginx
```
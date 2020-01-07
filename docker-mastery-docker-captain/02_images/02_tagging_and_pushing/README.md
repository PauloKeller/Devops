# Tagging and Pushing

- pull a image from latest tag
```
docker pull mysql/mysql-server
```

- pull a nginx image from mainline
```
docker pull nginx:mainline
```

- tag a existing image
```
docker image tag <image> <tag>
```

```
docker image tag nginx paulokeller/nginx
```

```
docker image tag paulokeller/nginx paulokeller/nginx:testing
```

- login into docker
```
docker login
```

- logout
```
docker logout
```

- push image
```
docker image push <tag>
```

```
docker image push paulokeller/nginx
```
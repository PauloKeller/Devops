# Private docker registry with swarm

- create registry service
```
docker service create --name registry --publish 5000:5000 registry
```

- pull image
```
docker pull hello-world
```

- tag to a new repo
```
docker tag hello-world 127.0.0.1:5000/hello-world
```

- push the image
```
docker push 127.0.0.1:5000/hello-world
```

- url to see the repos
```
http://<yourip>:5000/v2/_catalog
```

- pull nginx
```
docker pull nginx
```

- tag nginx
```
docker tag nginx 127.0.0.1:5000/nginx
```

- push nginx
```
docker push nginx 127.0.0.1:5000/nginx
```

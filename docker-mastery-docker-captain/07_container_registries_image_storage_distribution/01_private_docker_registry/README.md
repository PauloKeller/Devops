# Private docker registry

- get registry container running
```
docker container run -d -p 5000:5000 --name registry registry
```

- get a docker container
```
docker pull hello-world
```

- tag the image locally 
```
docker tag hello-world 127.0.0.1:5000/hello-world
```

- push the image to the registry
```
docker push 127.0.0.1:5000/hello-world
```

- remove the image
```
docker image rm -f 127.0.0.1:5000/hello-world
```

- pull the image back 
```
docker pull 127.0.0.1:5000/hello-world
```

- remove the registry
```
docker container kill registry
```

```
docker container rm registry
```

- run registry again
```
docker container run -d -p 5000:5000 --name registry -v $(pwd)/registry-data:/var/lib/registry registry
```
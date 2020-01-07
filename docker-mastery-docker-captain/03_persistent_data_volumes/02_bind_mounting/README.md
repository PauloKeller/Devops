# Bind mounting

- setting up the nginx volume to current docker share path
```
docker container run -d --name nginx -p 80:80 -v $(pwd):/usr/share/nginx/html nginx
```

- run the default nginx
```
docker container run -d --name nginx2 -p 8080:80 nginx
```

- enter to binded container
```
 docker container exec -it nginx bash
```

- create a file in current location where image was created and this will show on running container 
```
new.txt
```
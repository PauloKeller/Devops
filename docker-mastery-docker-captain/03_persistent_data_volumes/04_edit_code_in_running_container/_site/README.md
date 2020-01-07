# Edit code in running container

- run the following command to start the container
```
docker run -p 80:4000 -v $(pwd):/site bretfisher/jekyll-serv
```
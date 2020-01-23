# Routing mesh

- create 3 replicas of elastic search service 
```
docker service create --name search --replicas 3 -p 9200:9200 elasticsearch:2
```
- get elastic search basic info
```
curl localhost:9200
```
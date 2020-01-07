# DNS Round Robin test

- create a network
```
docker network create dude
```

- no name needed just attach to dude and find by dns alias "search" (run 2 times this command)
```
docker container run -d --net dude --net-alias search elasticsearch:2
```

- just run alpine in same net and add search dns
```
docker container run --rm --net dude alpine nslookup search
```

- just run centos with curl passing the dns
```
docker container run --rm --net dude centos curl -s search:9200
```
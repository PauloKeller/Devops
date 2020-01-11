# Create 3 nodes with swarm

- see docker swarm manager token
```
docker swarm join-token manager
```

- in node 2 run this command to link nodes
```
docker swarm join --token SWMTKN-1-2jgii8a7x9h7urwivkli4rw1w751xrnaw6yro6j4i1vj066wgj-eiiqh6eo9jpb3jv9duxq9hnyc 172.20.0.1:2377
```

- in node 1 run to update the node 2 role
```
docker node update --role manager node2
```

- in node 3 run this command to link nodes as node 3 as manager
```
docker swarm join --token <manager-token> <ip-address>
```

- create 3 replicas of same service
```
docker service create --replicas 3 alpine ping 8.8.8.8
```

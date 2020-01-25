# Runing service update 

- just update the image used to a newer version
``` 
docker service update --image myapp:1.2.1 <servicename>
```

- adding an enviroment variable and remove a port
```
docker service update --env-add NODE_ENV=production --publish-rm 8080
```

- change number of replicas of two services
```
docker service scale web=8 api=6
```

- same commang just edit the YAML file then
```
docker stack deploy -c file.yml <stackname>
```

## Example

- create the app service
```
docker service create -p 8088:80 --name web nginx:1.13.7
```

- scale the app
```
docker service scale web=5
```

- update the app image
```
docker service update --image nginx:1.13.6 web
```

- change the publish port
```
docker service update --publish-rm 8088 --publish-add 9090:80 web
```

- force update (rebalancing)
```
docker service update --force web
```
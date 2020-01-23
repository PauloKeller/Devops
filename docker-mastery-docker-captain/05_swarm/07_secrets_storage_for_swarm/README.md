# Secrets storage for swarm

- create sercret from file
```
docker secret create <name> <secret-file>
```

```
docker secret create psql_user psql_user.txt
```

- create secret from command line
```
echo "<secret>" | docker secret create <name> - 
```

```
echo "myDBpassWORD" | docker secret create psql_pass -
```

- use secrets on services
```
docker service create --name psql --secret psql_user --secret psql_pass \
-e POSTGRES_PASSWORD_FILE=/run/secrets/psql_pass \
-e POSTGRES_USER_FILE=/run/secrets/psql_user postgres
```
# Secrets on a local docker compose

Only work with file base secrets

- deploy the compose
```
docker-compose up -d
```

- bind mount the compose with secrets
```
docker-compose exec psql cat /run/secrets/psql_user
```
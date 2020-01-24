# Fully feature app build and deploy with single compose design

use docker-compose.yml first and overlay the docker-compose.override.yml on top

- to get a quick look when started
```
docker-compose up -d 
```

- inspect the drupal image, and see the mounts
```
docker inspect 02_build_and_deploy_with_single_compose_drupal_1
```

- for the CI solution how to run docker compose, 

```
docker-compose -f docker-compose.yml -f docker-compose.test.yml up -d
```

```
docker inspect 02_build_and_deploy_with_single_compose_drupal_1
```

the same container there is no bind mounts, in tests we doesnt need to save information

- use the default config to look at the 2 files and push them together and have the a output to a file "output.yml" to create our stack
```
docker-compose -f docker-compose.yml -f docker-compose.prod.yml config > output.yml
```
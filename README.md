# my_docker_project
Simple project to demonstrate running multiple services with docker compose and a reverse proxy (traefik)

### Starting Services
Simply run the following command to start the services:
```sh
docker compose up --build
```

Access service1 at http://localhost/service1

Access service2 at http://localhost/service2

Access service3 at http://localhost/service3

Access Traefik dashboard at http://localhost:8080/dashboard/#/

NB: Traefik takes a couple minutes to come up, allow some time if things aren't opening on the browser.

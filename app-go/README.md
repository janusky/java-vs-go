# app-go 2019/03/22

Aplicación con 3 servicios HTTP, conectada a base de datos Postgres SQL.

Los fuentes de la aplicación tomados de [nikitsenka/bank-go](https://github.com/nikitsenka/bank-go) como punto de partida.

## Ejecutar aplicación

Ejecutar modo desarrollo con docker-compose.yml

```sh
docker-compose up -d --build

# TEST (test/curl-api.sh)
http://localhost:8000
```

Run tests locally

```
  docker pull postgres
  docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=test1234 -d postgres
  go test ./app
```

Build docker image

```sh
docker build --no-cache -t app-go .  
```

Run app in Docker with external postgres

```sh
docker run --name app-go -p 8000:8000 -e POSTGRES_HOST=${host} -d app-go
```

### Ejecutar desde DIST

En el directorio [dist](./dist/README.md) se encuentra el archivo `Dockerfile` para crear la imagen que contiene la aplicación `app-go`. También está el archivo `docker-compose.yml` necesario para ejecutar el contenedor.

- [Ver DIST](./dist/README.md)
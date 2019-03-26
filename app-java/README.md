# app-java 2019/03/26

Aplicación con 3 servicios HTTP, conectada a base de datos Postgres SQL.

Los fuentes de la aplicación tomados de [nikitsenka/bank-java](https://github.com/nikitsenka/bank-java) como punto de partida.

## Ejecutar aplicación

Ejecutar modo desarrollo con docker-compose.yml

```sh
docker-compose up -d --build

# TEST (test/curl-api.sh)
http://localhost:8080
```

### Ejecutar desde DIST

En el directorio [dist](./dist/README.md) se encuentra el archivo `Dockerfile` para crear la imagen que contiene la aplicación `app-java`. También está el archivo `docker-compose.yml` necesario para ejecutar el contenedor.

- [Ver DIST](./dist/README.md)

## Configuración de aplicación

Los archivos de configuración utilizados pueden variar según el ambiente donde se está trabajando.

1. Se podría requerir certificados de confianza SSL, necesarios para acceder a internet.

1. Como repositorio de fuentes se podría utilizar uno propio <https://my.nexus.server.com/nexus/repository/maven-central>

1. Proxy de red <http://my.proxy.com:80>
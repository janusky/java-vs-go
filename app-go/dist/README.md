# dist app-go

Se debe crear la imagen y luego ejecutar `docker-compose`

```sh
# Crear imagen docker
docker build --build-arg http_proxy=$http_proxy --build-arg https_proxy=$https_proxy --no-cache -t app-go .
# Ejecutar
docker-compose up -d --force-recreate
```

IMPORTANTE

> Se espera contar con las tablas requeridas en base de datos. Según lo indicado en archivo [../config/db.sql](../config/db.sql)
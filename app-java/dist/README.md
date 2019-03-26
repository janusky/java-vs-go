# dist app-java

Se dispone de archivos docker para ejecutar la aplicación desde una versión compilada.
Para ejecutar `docker-compose.yml` se debe copiar el archivo ejecutable tipo jar deseado ([ver](#crear-archivo-ejecutable-tipo-jar)).

## Crear archivo ejecutable tipo JAR

Puede generar el archivo `app-java-0.0.1-SNAPSHOT.jar` ejecutando Maven o copiar el generado al ejecutarse la instancia de contenedor en desarrollo ([ir](../README.md#ejecutar-aplicación)).

```sh
# Ejecutar el empaquetado del proyecto y copiar el jar
cd ~/go/src/github.com/janusky/java-vs-go/app-java
mvn clean install -DtestSkip=true
cp app-java-0.0.1-SNAPSHOT.jar ./dist

# ó

# Copiar el jar desde el contenedor lanzado desde archivo docker-compose.yml
cd ~/go/src/github.com/janusky/java-vs-go/app-java
docker-compose up -d
docker cp  appjava_app-java_1:/app-java/target/app-java-0.0.1-SNAPSHOT.jar ./dist
```

## Ejecutar desde dist

Se debe crear la imagen y luego ejecutar `docker-compose`

```sh
# Crear imagen docker
docker build --no-cache -t app-java .
# Ejecutar
docker-compose up -d --force-recreate
```

IMPORTANTE

> Se espera contar con las tablas requeridas en base de datos. Según lo indicado en archivo [../config/db.sql](../config/db.sql)
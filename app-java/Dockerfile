FROM openjdk:11

RUN apt-get update && apt-get install -y git maven

# Usar certificados para acceder a internet
#ADD ./config/RootCA.crt /usr/local/share/ca-certificates/RootCA.crt
#RUN chmod 644 /usr/local/share/ca-certificates/RootCA.crt && update-ca-certificates

COPY . app-java

WORKDIR app-java

ADD ./config/settings.xml settings.xml
RUN mvn -s settings.xml clean install -DtestSkip=true

CMD java -jar target/app-java-0.0.1-SNAPSHOT.jar
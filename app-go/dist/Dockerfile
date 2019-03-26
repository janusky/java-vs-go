FROM golang:1.8

WORKDIR /go/src

RUN git config --global http.sslverify false

# Crea y accede el directorio de la aplicación
RUN mkdir -p github.com/janusky/java-vs-go/app-go && cd $_

# Descarga el proyecto y fija los fuentes apropiados de aplicación Go
RUN curl -L https://github.com/janusky/java-vs-go/archive/master.zip -O
RUN unzip master.zip -d tmp && mv tmp/*/app-go app-go && rm -rf tmp
RUN rm master.zip

# Dependencias
RUN go get -d -v github.com/janusky/java-vs-go/app-go

# Instalar app
RUN go install

CMD ["app-go"]
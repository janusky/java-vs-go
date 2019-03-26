FROM golang:1.8

RUN git config --global http.sslverify false

COPY . /go/src/github.com/janusky/java-vs-go/app-go

WORKDIR /go/src/github.com/janusky/java-vs-go/app-go

# Descarga las dependencias
RUN go get -d -v github.com/janusky/java-vs-go/app-go

RUN go install

CMD ["app-go"]
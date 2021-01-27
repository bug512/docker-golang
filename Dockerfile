FROM golang:1.14

ARG APP_NAME
ARG PORT 

COPY app/src /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./

#RUN go mod init tvil.ru/tvilCrawler
RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o ${APP_NAME}" --command=./${APP_NAME}
#RUN go build -o ${APP_NAME}

#CMD go run main.go

#CMD ./${APP_NAME}

EXPOSE ${PORT}
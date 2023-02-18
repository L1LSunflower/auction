FROM golang:1.19-alpine3.17 as build

RUN apk update && \
    apk upgrade && \
    apk add gcc && \
    apk add musl-dev && \
    apk add make

COPY . /go/src
WORKDIR /go/src

RUN ls

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o auction cmd/app/main.go && \
    go build -o goose cmd/goose/*.go

FROM alpine:latest AS final

RUN apk add --no-cache tzdata libc6-compat make

WORKDIR /go/src

COPY --from=build /go/src/auction .
COPY --from=build /go/src/goose .

COPY ./internal/db/migrations ./migrations
COPY ./Makefile .
COPY .env .
COPY ./deployment/run.sh run.sh

RUN apk update && \
    apk upgrade && \
    apk add gcc && \
    apk add musl-dev && \
    apk add make

CMD chmod +x auction && ./auction

#CMD chmod +x ./run.sh
#CMD ["./run.sh"]
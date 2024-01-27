FROM golang:1.21-alpine3.19

WORKDIR /

COPY ./data /src/data
COPY ./handlers /src/handlers
COPY ./infra /src/infra
COPY ./middlewares /src/middlewares
COPY ./services /src/services
COPY ./utils /src/utils
COPY ./main.go /src
COPY  go.mod go.sum /src/

WORKDIR /src
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ../triniti

EXPOSE 8080
ENV PORT=:8080

WORKDIR /
RUN rm -rf src


CMD [ "/triniti" ]
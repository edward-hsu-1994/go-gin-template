FROM golang:alpine AS builder
ARG APP_NAME=my-app
ENV APP_NAME=${APP_NAME}

RUN apk add --no-cache make
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY . .

RUN make build


FROM alpine:latest
ARG APP_NAME=my-app
ENV APP_NAME=${APP_NAME}

WORKDIR /app

COPY --from=builder /app/build .

EXPOSE 8080

LABEL authors="edward_hsu_1994"

ENTRYPOINT ["sh", "-c", "./$APP_NAME"]
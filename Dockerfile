FROM golang:alpine as builder

RUN apk add --no-cache make
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY . .

RUN make build


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/build .

EXPOSE 3000

LABEL authors="edward_hsu_1994"

ENTRYPOINT ["./my-fiber-app"]
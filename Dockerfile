FROM golang:1.21.1-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN mkdir bin
RUN go build -o ./bin/lmsapp ./main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/.env ./
COPY --from=builder /app/bin/lmsapp ./

RUN mkdir log

EXPOSE 6000

CMD ./lmsapp
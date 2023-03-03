FROM golang:1.17-alpine3.15

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /main .
EXPOSE 8000
CMD [ "/main" ]
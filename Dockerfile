FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-project-starter ./cmd/main.go

EXPOSE 8080

CMD [ "/go-project-starter" ]
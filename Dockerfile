FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY ./msg/*.go ./msg/

RUN CGO_ENABLED=0 GOOS=linux go build -o /director

EXPOSE 50051

CMD ["/director"]
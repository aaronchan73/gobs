FROM golang:1.21.5
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o collector ./collector
RUN chmod +x collector
EXPOSE 8080
ENTRYPOINT ["./collector/collector"]
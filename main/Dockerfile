FROM golang:1.21.5
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main ./main
RUN chmod +x main
EXPOSE 8080
ENTRYPOINT ["./main/main"]
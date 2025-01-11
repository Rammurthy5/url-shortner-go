FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /app/cmd/web/main /app/cmd/web/main.go
EXPOSE 8080
ENV GIN_MODE=release
CMD ["/app/cmd/web/main"]
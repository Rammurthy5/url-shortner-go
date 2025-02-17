FROM golang:1.24-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /cmd/web-app/main /cmd/web-app/main.go
EXPOSE 3200
CMD ["/cmd/web-app/main"]

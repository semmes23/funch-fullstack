FROM golang:1.23

# Install swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

# Copy dependencies & download modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Generate Swagger docs and remove unsupported fields
RUN swag init --output ./docs --parseDependency && \
    sed -i '/LeftDelim/d' ./docs/docs.go && \
    sed -i '/RightDelim/d' ./docs/docs.go

# Build binary
RUN go build -o main .

EXPOSE 8080

CMD ["go", "run", "main.go"]



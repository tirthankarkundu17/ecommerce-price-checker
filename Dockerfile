FROM golang:alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /ecommerce-price-checker
EXPOSE 8000

# Run
CMD ["/ecommerce-price-checker"]
# ecommerce-price-checker
Go App to check price and details of Ecommerce listed products

# Running the app locally
To run the app locally, you can run below command -
```
go run server.go
```

# Build Docker
The app uses Dockerfile to create a Docker Image which can then be containerized.
To create a docker image, run below command -
```
docker build --tag ecommerce-price-checker .
```

# Run Docker Image
After creating docker image, run below command -
```
docker run --name ecommerce-price-checker -d -p 8085:8000 ecommerce-price-checker
```
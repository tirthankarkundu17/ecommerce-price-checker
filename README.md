# ecommerce-price-checker
Go App to check price and details of Ecommerce listed products.
This uses web scraping to fetch the data. With the help of goroutines, the app makes multiple calls in relatime and fetches the data faster.

# Running the app locally
To run the app locally, 
make a rename .env.example to .env
Update the values for - 
```
DB_USER=<MYSQL_USER>
DB_PASS=<MYSQL_PASSWORD>
DB_HOST=<MYSQL_DATABASE_HOST>
DB_SCHEMA=<MYSQL_DATABASE_SCHEMA>
API_SECRET=<JWT_TOKEN_SECRET>
``` 

Now run below command -
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

# Sample Requests

- Register

    <img src="docs\user_create.png" width="500">
- Login

    <img src="docs\user_login.png" width="500">
- Add Product of Interest

    <img src="docs\add_products.png" width="500">
- Fetch Products of Interests

    <img src="docs\fetch_products.png" width="500">
# Inventory Backend
## DEMO 
Successfully deployed at https://fabuzard-fullstack-dev.netlify.app/

## Technology Stack
- **Frontend:** React, Tailwind CSS  
- **Backend:** Go, Echo framework  
- **Database:** MySQL  
- **Deployments:** Netlify (frontend), Heroku (backend), Google Cloud SQL (database)  

## Prerequisites
- Go 1.24+  
- MySQL database  

## Local Setup

1. Clone the repository and navigate to the backend folder:
```bash
git clone <repo_url>
cd backend
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file in the root:
```env
DB_USER=your_db_user
DB_PASS=your_db_password
DB_HOST=your_db_host
DB_PORT=3306
DB_NAME=your_db_name
jwt_secret=verysecretkey
```

4. Run the backend server:
```bash
go run main.go
```

5. Navigate to the frontend folder:
```bash
cd ../frontend
```

6. Install dependencies:
```bash
npm install
```

7. Update the API URL in `src/services` to point to your backend (e.g., `http://localhost:8080`)

8. Run the frontend:
```bash
npm run dev
```

## Features
- User Authentication: Register and log in using email and password (JWT-based)
- Product Management (CRUD): Manage products with fields: name, SKU, quantity, location, and status
- CSV Export: Export full product list for reporting
- Secure Endpoints: JWT-protected routes for products
- CORS Enabled: Supports requests from frontend apps
- Configurable Environment: Easily set DB and JWT via environment variables

## API Endpoints

### Authentication
- `POST /register` - Register a new user
- `POST /login` - User login

### Products (JWT Required)
- `POST /products` - Create a new product
- `GET /products` - Get all products
- `GET /products/:id` - Get product by ID
- `PUT /products/:id` - Update product by ID
- `DELETE /products/:id` - Delete product by ID
- `GET /products/export` - Export products to CSV

## API Documentation

### Authentication Endpoints

#### Register User
```
POST /register
```

Request:
```json
{
  "fullname": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "message": "User created successfully",
  "user": {
    "id": 1,
    "fullname": "John Doe",
    "email": "john@example.com"
  }
}
```

#### Login User
```
POST /login
```

Request:
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "fullname": "John Doe",
  "email": "john@example.com",
  "id": 1
}
```

### Product Endpoints (JWT Required)

#### Create Product
```
POST /products
```

Request:
```json
{
  "product_name": "Laptop",
  "sku": "LAP001",
  "quantity": 10,
  "location": "Warehouse A",
  "status": "in_stock"
}
```

Response:
```json
{
  "id": 1,
  "product_name": "Laptop",
  "sku": "LAP001",
  "quantity": 10,
  "status": "in_stock",
  "location": "Warehouse A"
}
```

#### Get All Products
```
GET /products
```

Response:
```json
[
  {
    "id": 1,
    "product_name": "Laptop",
    "sku": "LAP001",
    "quantity": 10,
    "status": "in_stock",
    "location": "Warehouse A"
  }
]
```

#### Get Product by ID
```
GET /products/:id
```

Response:
```json
{
  "id": 1,
  "product_name": "Laptop",
  "sku": "LAP001",
  "quantity": 10,
  "status": "in_stock",
  "location": "Warehouse A"
}
```

#### Update Product
```
PUT /products/:id
```

Request:
```json
{
  "product_name": "Gaming Laptop",
  "sku": "LAP001",
  "quantity": 8,
  "location": "Warehouse A",
  "status": "in_stock"
}
```

Response:
```json
{
  "id": 1,
  "product_name": "Gaming Laptop",
  "sku": "LAP001",
  "quantity": 8,
  "status": "in_stock",
  "location": "Warehouse A"
}
```

#### Delete Product
```
DELETE /products/:id
```

Response:
```json
{
  "message": "Product deleted successfully"
}
```

#### Export Products to CSV
```
GET /products/export
```

Response: CSV file download

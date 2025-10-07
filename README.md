# 🛒 Go E-Commerce Backend API

A scalable **E-Commerce REST API** built with **Go (Golang)**, **Gin**, and **GORM**.
This project follows a clean architecture design and supports authentication, product management, cart operations, and order processing.

---

## 🚀 Tech Stack

| Layer          | Technology            |
| -------------- | --------------------- |
| Language       | Go (Golang)           |
| Framework      | Gin Gonic             |
| ORM            | GORM                  |
| Database       | MySQL                 |
| Authentication | JWT (JSON Web Tokens) |
| Config         | dotenv                |
---

## 📁 Folder Structure

```
ecommerce-api/
│
├── main.go               # App entry point
│
├── config/
│   └── config.go            # Loads environment variables
│
├── database/
│   └── database.go                # Database connection setup
│
├── models/
│   ├── user.go
│   ├── product.go
│   ├── category.go
│   ├── cart.go
│   ├── cart_item.go
│   ├── order.go
│   ├── order_item.go
│   └── review.go
│
├── controllers/
│   ├── auth_controller.go
│   ├── product_controller.go
│   ├── category_controller.go
│   ├── cart_controller.go
│   └── order_controller.go
│
├── routes/
│   └── routes.go
│
├── middleware/
│   └── auth_middleware.go
│
├── utils/
│   └── token.go             # JWT helper functions
│
├── go.mod
├── go.sum
└── .env.example
```

---

## ⚙️ Environment Variables

Create a `.env` file in your project root with the following variables:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=ecommerce_db
JWT_SECRET=mysecretkey
APP_PORT=8080
```

---

## 🧩 Features

✅ User Authentication (Register / Login with JWT)
✅ Product Management (CRUD + Category Filter)
✅ Category Management
✅ Shopping Cart (Add / Remove / Update Items)
✅ Order Processing (Checkout & Order History)
✅ Soft Deletes and Timestamp Tracking
✅ Database Auto Migration
✅ Modular Folder Structure for Scalability

---

## 🏗️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/<your-username>/go-ecommerce-api.git
cd go-ecommerce-api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Up Database

Make sure your MySQL server is running and update `.env` with your credentials.

---

### 4. Run the API

```bash
go run main.go
```

Server should start at:

```
http://localhost:8080
```

---

## 🧠 API Endpoints Overview

| Method | Endpoint             | Description             | Auth Required |
| ------ | -------------------- | ----------------------- | ------------- |
| POST   | `/api/auth/register` | Register new user       | ❌             |
| POST   | `/api/auth/login`    | Login user              | ❌             |
| GET    | `/api/products`      | Get all products        | ❌             |
| GET    | `/api/products/:id`  | Get product by ID       | ❌             |
| POST   | `/api/products`      | Create product          | ✅             |
| PUT    | `/api/products/:id`  | Update product          | ✅             |
| DELETE | `/api/products/:id`  | Delete product          | ✅             |
| GET    | `/api/categories`    | Get all categories      | ❌             |
| POST   | `/api/categories`    | Add category            | ✅             |
| GET    | `/api/cart`          | Get user cart           | ✅             |
| POST   | `/api/cart`          | Add item to cart        | ✅             |
| PUT    | `/api/cart/:itemId`  | Update cart item        | ✅             |
| DELETE | `/api/cart/:itemId`  | Remove cart item        | ✅             |
| POST   | `/api/orders`        | Checkout / Create order | ✅             |
| GET    | `/api/orders`        | Get user orders         | ✅             |

---

## 🧮 Example: Product Model (GORM)

```go
type Product struct {
    ID          uint           `gorm:"primaryKey"`
    Name        string         `gorm:"not null"`
    Description string
    Price       float64        `gorm:"not null"`
    Stock       int
    CategoryID  uint
    Category    Category
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}
```

---

## 🛠️ Example: Main Entry (main.go)

```go
package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "ecommerce-api/database"
    "ecommerce-api/routes"
)

func main() {
    godotenv.Load()
    database.ConnectDB()

    r := gin.Default()
    routes.RegisterRoutes(r)

    port := os.Getenv("APP_PORT")
    log.Printf("Server running on port %s", port)
    r.Run(":" + port)
}
```

---

## 🔐 Authentication Flow

1. **Register** → POST `/api/auth/register` → stores hashed password
2. **Login** → POST `/api/auth/login` → returns JWT token
3. **Protected Routes** → use `Authorization: Bearer <token>` header
4. Middleware verifies token and attaches `userID` to the request context

---

## 🧩 Future Enhancements

* Payment integration (Stripe / Paystack)
* Product image upload (Cloudinary / S3)
* Email notifications
* Admin dashboard endpoints
* Redis caching for products
* WebSocket for order updates

---

## 🤝 Contributing

1. Fork the repo
2. Create a new branch `feature/your-feature`
3. Commit changes and push
4. Submit a Pull Request

---

## 🗾 License

MIT License © 2025 Abdulqudri Abdullahi

---

## ⭐ Support

If you find this project helpful, please give it a **star** ⭐ and share!

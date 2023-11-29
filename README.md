# Go Bookstore Application

## Overview

This repository contains a simple Go web application for managing a bookstore. The application includes server-side code written in Go, utilizing the Gorilla Mux router and GORM for database operations. The application provides CRUD (Create, Read, Update, Delete) operations for managing books.

## Project Structure

The project is structured into several packages, each serving a specific purpose:

- **main**: Contains the main entry point of the application.
  - `main.go`: Initializes the server and sets up the routes.

- **pkg**: Organizes the code into several packages, each serving a specific purpose.
  - **config**: Manages the database configuration.
    - `app.go`: Defines functions to connect to the database and retrieve the database instance.
  - **controllers**: Handles HTTP request handling and response for each route.
    - `book-controller.go`: Implements functions for handling book-related operations, such as retrieving all books, getting a book by ID, creating, updating, and deleting a book.
  - **models**: Defines the data model for books and contains database-related functions.
    - `book.go`: Defines the `Book` struct and includes functions for CRUD operations on books.
  - **routes**: Registers all routes (URL endpoints) with the router.
    - `bookstore-routes.go`: Registers routes for creating, retrieving, updating, and deleting books.
  - **utils**: Contains utility functions.
    - `utils.go`: Includes a function for parsing the request body as JSON.

## Setup

1. Ensure you have Go installed on your machine.

2. Clone the repository or download the files.

3. Navigate to the project directory.

4. Run the main application:
   ```bash
   go run main.go
   ```

The server will start on `localhost:9010`.

## Usage

The application provides the following API endpoints:

- `POST /book/`: Create a new book.
- `GET /book/`: Retrieve all books.
- `GET /book/{bookId}`: Retrieve a book by ID.
- `PUT /book/{bookId}`: Update a book by ID.
- `DELETE /book/{bookId}`: Delete a book by ID.

Make HTTP requests to these endpoints to perform CRUD operations on the bookstore.

## Database Configuration

The application is configured to use MySQL. Update the database connection string in `config/app.go` to match your database configuration.

```go
// config/app.go
d, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local")
```

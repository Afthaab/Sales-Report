# Sales-Report

This project is a backend system built with Golang and PostgreSQL to process large CSV files containing sales data, store them in a normalized database, and provide RESTful APIs for analysis.

## Features
- Efficiently loads large CSV files into the database.
- Periodic and on-demand data refresh mechanisms.
- RESTful APIs for querying sales data and analytics.
- Logging for data refresh operations.

## Prerequisites
- Golang (>=1.20)
- PostgreSQL (>=13)

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/Afthaab/Sales-Report
   cd your-repo
   ```

2. Add the Environment Variables to the env file:
   ```sh
   DSN="host=localhost user=your_user password=your_db_password dbname=sales port=your_db_port sslmode=disable TimeZone=Asia/Kolkata"
   ```
   Update `.env` with the correct database credentials.
     ```sh
      CREATE DATABASE sales;
   ```
     Create Database in your postgres


4. Install dependencies:
   ```sh
   go mod tidy
   ```

5. Start the application:
   ```sh
   go run ./cmd/app/sales_report/main.go
   ```

## Database Schema
The database is structured in a normalized format with tables for orders, customers, products, and regions.

![Database Schema](![drawSQL-image-export-2025-03-27 (1)](https://github.com/user-attachments/assets/afcfa779-cd6c-4ab4-b4fe-dcc29307185d)
)

## CSV Processing
1. Place CSV files in the `input_csv/` directory.
2. The system will process all `.csv` files in the directory and move:
   - Successful files to `processed_csv/`
   - Failed files to `failed_csv/`

## API Endpoints

### API Documentation
```http
https://documenter.getpostman.com/view/25649054/2sB2cPi5D1
```
### 1. Refresh Data (On-demand)
**Endpoint:**
```http
POST /api/refresh
```
**Description:** Triggers CSV data refresh manually.

### 2. Get Total Number of Customers (Date Range)
**Endpoint:**
```http
GET http://localhost:8000/total/customers?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
```
**Description:** Returns the count of unique customers who placed orders within the specified date range.

### 3. Get Total Number of Orders (Date Range)
**Endpoint:**
```http
GET http://localhost:8000/total/orders?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
```
**Description:** Returns the count of unique orders placed within the specified date range.

### 4. Get Average Order Value (Date Range)
**Endpoint:**
```http
GET http://localhost:8000/average/order_value?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
```
**Description:** Returns the average order value within the specified date range.

More endpoints are documented in `api_documentation.md`.

## Logging
- Refresh events are logged in the `logs/refresh.log` file.
- Errors are logged for troubleshooting.

## Cron Job Setup
To run the scheduled job daily, configure a cron job:
```sh
CRON_JOB_SCHEDULER = "* * * * *"
```
Alternatively, use a background worker like `go-cron`, By default its happening every One Minute.

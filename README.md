# react-go-mvc

**Table of contents**
- [How to run](#how-to-run)
	- [Backend](#backend)
		- [Example Sqlite](#example-sqlite)
		- [Example PostgreSQL](#example-postgresql)
	- [Frontend](#frontend)
- [API Endpoints](#api-endpoints)
	- [Business](#business)
	- [Address](#address)

A straight forward React + Go (mvc) web app.
React part was mostly ChatGPT with some tweaks to get it working propper.
Golang part was constructed from reading a couple of articles about mvc architecture in go, as well as skimming through some public repos.

## How to run

Two parts, the backend service with go and the frontend with react.

### Backend

Build:
```sh
go mod tidy
go build -o main .
```

Run:
```sh
./main [-config-file <file-name>]
```

Config file (defaults to `./config.json`) contains the hosting information, log levels, and database conntections (`sqlite` and `postgresql` currently supported).

#### Example Sqlite:
```json
{
    "host": "",
    "port": "8080",
    "api_prefix": "/api",
    "log_level": "Trace",
    "database": {
    	"log_level": "Info",
    	"sqlite": {
    		"path": "/tmp/react_go_mvc.db"
    	}
    }
}
```

#### Example PostgreSQL:
```json
{
    "host": "",
    "port": "8080",
    "api_prefix": "/api",
    "log_level": "Trace",
    "database": {
    	"log_level": "Info",
    	"postgresql": {
    		"host": "localhost",
    		"port": "5432",
    		"user": "postgres",
    		"password": "postgres",
    		"db_name": "react_go_mvc"
    	}
    }
}
```

### Frontend

All frontend code lives inside the `client` directory, so first `cd client` before starting.

Build & Run:
```sh
npm install
npm start
```

Run:
Navigate to the link in your browser (most likely `http://localhost:3000`)

## API Endpoints

### Business

Enpoint at `/business`

Fetch all:

```http
GET http://localhost:8080/api/business HTTP/1.1
Accept: *
```

Fetch one:

```http
GET http://localhost:8080/api/business/1 HTTP/1.1
Accept: *
```

Create:

```http
POST http://localhost:8080/api/business HTTP/1.1
Content-Type: application/json
Accept: *

{
	"name": "test",
	"vat_number": 123,
	"registration_number": "2025/test"
}

```

Update:

```http
PUT http://localhost:8080/api/business/1 HTTP/1.1
Content-Type: application/json
Accept: *

{
	"name": "test",
	"vat_number": 123456,
	"registration_number": "2025/test"
}
```

Delete:

```http
DELETE http://localhost:8080/api/business/1 HTTP/1.1
Accept: *
```

### Address

Enpoint at `/address`

Fetch all for a business:

```http
GET http://localhost:8080/api/address?business_id=1 HTTP/1.1
Accept: *
```

Fetch one:

```http
GET http://localhost:8080/api/address/1 HTTP/1.1
Accept: *
```

Create:

```http
POST http://localhost:8080/api/address HTTP/1.1
Content-Type: application/json
Accept: *

{
	"business_id": 1,
	"street_number": "0",
	"street": "asdf",
	"town": ";lkj",
	"post_code": "1234"
}
```

Update:

```http
PUT http://localhost:8080/api/address/1 HTTP/1.1
Content-Type: application/json
Accept: *

{
	"street_number": "0",
	"street": "asdf",
	"town": ";lkj",
	"post_code": "0987"
}
```

Delete:

```http
DELETE http://localhost:8080/api/address/2 HTTP/1.1
Accept: *
```

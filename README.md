# Email Client in Go

This project is a email sening client using Go. It provides endpoints for various operations and is designed to be lightweight, efficient, and easy to use.

## Features
- RESTful API endpoints
- Environment-based configuration
- Modular and scalable code structure

## Prerequisites
- Go (version 1.18 or later)
- A terminal or command-line interface
- A `.env` file for environment variables

# Instructions
- Replace .env.production to .env only


## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

### 2. Set Up Environment Variables
Create a `.env` file in the root directory and add the following variables:
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
EMAIL_ID=exampl@gmail.com
EMAIL_USER=john 
EMAIL_PASSWORD=lkhlkhslkfhsldhf

```

### 3. Install Dependencies
Run the following command to install the required Go modules:
```bash
go mod tidy
```

### 4. Run the Application
Start the server using:
```bash
go run main.go
```

The server will start on the port specified in the `.env` file (default: `8080`).

### 5. Test the API
Use tools like [Postman](https://www.postman.com/) or `curl` to test the API endpoints. For example:
```bash
curl http://localhost:8080/health
```


## Contributing
Feel free to fork this repository and submit pull requests. Contributions are welcome!

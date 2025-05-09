# XYZ Multifinance Application

This repository contains a Go-based microservice application with PostgreSQL database integration.

## Prerequisites

- Docker
- Docker Compose
- Postman (optional, for API testing)

## Getting Started

### Running the Application

1. Clone the repository
2. Start the application using Docker Compose
### This command will:
- Start a PostgreSQL database container
- Build and start the application container
- Set up the necessary network connections
- Configure environment variables automatically

The application will be available at `http://localhost:8080`

### Application Components

- **Database**: PostgreSQL running on port 5432
- **API Service**: Go application running on port 8080

## Additional Resources

### API Documentation
The complete API documentation is available as a Postman collection in the `postman` folder. To use it:
1. Open Postman
2. Import the collection from `postman` folder
3. Use the imported collection to test the available endpoints

### Database
- Database initialization scripts are available in the `sql` folder
- Or Just running project using `docker-compose.yml` because there is already have migration and seeder

### Application Structure
Detailed documentation about the application and database structure can be found in the `structure` folder.

## Stopping the Application
To stop: `docker-compose down` or `docker-compose down -v` (removes volumes)

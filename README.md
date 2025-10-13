# EV Chargers Management System (ECMS)

A comprehensive web application for managing physical electric vehicle charging stations and their ports. The system enables creating locations where multiple charging stations can be placed, and automatically generates charging points (EVSE) based on station ports.

## Project Description

EV Chargers Management System is a centralized platform for managing electric vehicle charging infrastructure. It provides efficient management of stations, locations, and charging points within a single platform, designed for charging station owners, infrastructure operators, and system administrators.

### Key Features
- **Station Management**: CRUD operations for charging stations and their ports
- **Location Management**: Create physical locations with address data and assign stations
- **Automatic EVSE Generation**: 1:1 mapping from station ports to charging points
- **User Management**: Role-based access control (Admin/Owner)
- **Audit System**: Comprehensive logging of all operations and changes
- **Search & Filter**: Advanced search capabilities for stations and locations

## Tech Stack

### Backend
- **Language**: Go 1.25.1
- **Framework**: Gin 1.11.0
- **ORM**: GORM 1.31.0
- **Database Driver**: PostgreSQL

### Frontend
- **Framework**: Angular 20.3.0
- **Styling**: TailwindCSS 4.1.14
- **Runtime**: Node.js 22.17.0

### Database
- **Primary Database**: PostgreSQL

### CI/CD
- **Platform**: GitHub Actions

## Getting Started Locally

### Prerequisites
- Go 1.25.1 or higher
- Node.js 22.17.0 or higher
- PostgreSQL database
- Git

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables (create `.env` file):
```bash
# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=ecms

# JWT configuration
JWT_SECRET=your_jwt_secret
```

4. Run database migrations (if applicable):
```bash
# Add migration commands here when available
```

5. Start the backend server:
```bash
go run main.go
```

The backend API will be available at `http://localhost:8080`

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm start
```

The frontend application will be available at `http://localhost:4200`

## Available Scripts

### Frontend Scripts
- `npm start` - Start the development server
- `npm run build` - Build the application for production
- `npm run watch` - Build and watch for changes in development mode
- `npm test` - Run unit tests
- `npm run lint` - Run ESLint
- `npm run lint:fix` - Fix ESLint issues automatically

### Backend Scripts
- `go run main.go` - Start the development server
- `go test ./...` - Run all tests
- `go build` - Build the application

## Project Scope

### MVP Features
- ✅ CRUD operations for charging stations and ports
- ✅ CRUD operations for locations
- ✅ Automatic EVSE generation from station ports (1:1 mapping)
- ✅ User management with two roles (Admin, Owner)
- ✅ Single-tenant architecture
- ✅ Comprehensive audit and logging system
- ✅ REST API in Go + SPA in Angular
- ✅ PostgreSQL database integration

## Project Status

🚧 **Active Development**

This project is currently in active development. The MVP is being implemented according to the Product Requirements Document (PRD).

## Documentation

- [Product Requirements Document (PRD)](docs/prd.md)
- [Technical Stack Overview](docs/tech-stack.md)
- [MVP Specifications](docs/mvp.md)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

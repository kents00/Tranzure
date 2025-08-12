# Database Testing Guide

This document provides instructions on how to test the database functionality in the Tranzure payment application, including MongoDB setup, connection testing, and data validation.

## Prerequisites

- Docker and Docker Compose installed
- MongoDB tools (optional, for manual testing)
- WSL Ubuntu 24.04 or Windows with Go 1.22+

## Database Setup for Testing

### Option 1: Using Docker Compose (Recommended)

1. **Start the database services:**
   ```bash
   # Start MongoDB and admin tools
   make docker-up

   # Or manually with docker-compose
   docker-compose up -d mongodb mongo-express
   ```

2. **Verify services are running:**
   ```bash
   docker-compose ps
   ```

3. **Access MongoDB Admin UI:**
   - Mongo Express: http://localhost:8081
   - Username: `admin`, Password: `admin`

### Option 2: Local MongoDB Installation

1. **Install MongoDB on WSL Ubuntu:**
   ```bash
   # Import MongoDB public GPG Key
   curl -fsSL https://pgp.mongodb.com/server-7.0.asc | sudo gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor

   # Create list file for MongoDB
   echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list

   # Reload package database
   sudo apt-get update

   # Install MongoDB
   sudo apt-get install -y mongodb-org

   # Start MongoDB service
   sudo systemctl start mongod
   sudo systemctl enable mongod
   ```

2. **Initialize the database:**
   ```bash
   mongosh < scripts/mongo-init.js
   ```

## Database Connection Testing

### Basic Connection Test

1. **Test with the main application:**
   ```bash
   # Set environment variables
   export DATABASE_MONGODB_HOST=localhost
   export DATABASE_MONGODB_PORT=27017
   export DATABASE_MONGODB_DATABASE=payment_db

   # Run the application
   go run main.go
   ```

2. **Expected output:**
   ```
   Starting Payment Service on 0.0.0.0:8080
   Environment: development
   Successfully connected to MongoDB!
   Server running on http://0.0.0.0:8080
   ```

### Manual Database Testing

1. **Connect using MongoDB shell:**
   ```bash
   # For Docker setup
   docker exec -it tranzure-mongodb-1 mongosh

   # For local installation
   mongosh
   ```

2. **Test database operations:**
   ```javascript
   // Switch to payment database
   use payment_db

   // List collections
   show collections

   // Test user creation
   db.users.insertOne({
     user_id: "550e8400-e29b-41d4-a716-446655440000",
     email: "test@example.com",
     password_hash: "hashed_password",
     role: "user",
     status: "active",
     created_at: new Date(),
     updated_at: new Date()
   })

   // Test wallet creation
   db.wallets.insertOne({
     wallet_id: "550e8400-e29b-41d4-a716-446655440001",
     user_id: "550e8400-e29b-41d4-a716-446655440000",
     type: "crypto",
     currency: "BTC",
     address: "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh",
     balance: NumberDecimal("1.50000000"),
     is_primary: true
   })

   // Verify indexes
   db.users.getIndexes()
   db.wallets.getIndexes()
   ```

## Automated Database Tests

### Create Database Test Script

1. **Run database integration tests:**
   ```bash
   # Using the test script
   ./scripts/test_database.sh

   # Or directly with Make
   make test-database
   ```

### Test Collections and Schemas

1. **Validate collection schemas:**
   ```bash
   # Test user validation
   go test -v ./internal/database/tests/ -run TestUserCollection

   # Test wallet validation
   go test -v ./internal/database/tests/ -run TestWalletCollection

   # Test transaction validation
   go test -v ./internal/database/tests/ -run TestTransactionCollection
   ```

## Performance Testing

### Connection Pool Testing

1. **Test concurrent connections:**
   ```bash
   # Run load test
   go test -v ./internal/database/tests/ -run TestConnectionPool
   ```

2. **Monitor connection metrics:**
   ```bash
   # View MongoDB metrics
   docker exec -it tranzure-mongodb-1 mongosh --eval "db.serverStatus().connections"
   ```

### Query Performance Testing

1. **Test index efficiency:**
   ```javascript
   // In MongoDB shell
   use payment_db

   // Explain query plans
   db.users.find({email: "test@example.com"}).explain("executionStats")
   db.transactions.find({from_user_id: "user-id"}).explain("executionStats")
   ```

## Data Validation Testing

### Schema Validation Tests

1. **Test valid documents:**
   ```javascript
   // Valid user document
   db.users.insertOne({
     user_id: "valid-uuid",
     email: "valid@email.com",
     password_hash: "valid_hash",
     role: "user",
     status: "active"
   })
   ```

2. **Test invalid documents (should fail):**
   ```javascript
   // Invalid user document (missing required field)
   db.users.insertOne({
     user_id: "valid-uuid",
     email: "invalid-email",  // Invalid email format
     password_hash: "valid_hash"
     // Missing required fields: role, status
   })
   ```

## Backup and Recovery Testing

### Test Database Backup

1. **Create backup:**
   ```bash
   # For Docker setup
   docker exec tranzure-mongodb-1 mongodump --db payment_db --out /backup

   # For local installation
   mongodump --db payment_db --out ./backup
   ```

2. **Test restore:**
   ```bash
   # For Docker setup
   docker exec tranzure-mongodb-1 mongorestore --db payment_db_test /backup/payment_db

   # For local installation
   mongorestore --db payment_db_test ./backup/payment_db
   ```

## Troubleshooting

### Common Issues

1. **Connection Refused:**
   ```bash
   # Check if MongoDB is running
   docker-compose ps mongodb

   # Or for local installation
   sudo systemctl status mongod
   ```

2. **Authentication Failed:**
   ```bash
   # Check environment variables
   echo $DATABASE_MONGODB_USERNAME
   echo $DATABASE_MONGODB_PASSWORD

   # Verify user exists in MongoDB
   docker exec -it tranzure-mongodb-1 mongosh --eval "db.getUsers()"
   ```

3. **Database Not Found:**
   ```bash
   # List available databases
   docker exec -it tranzure-mongodb-1 mongosh --eval "show dbs"

   # Initialize database if needed
   docker exec -it tranzure-mongodb-1 mongosh < /docker-entrypoint-initdb.d/mongo-init.js
   ```

### Performance Issues

1. **Slow Queries:**
   ```javascript
   // Enable profiling
   db.setProfilingLevel(2, { slowms: 100 })

   // View slow operations
   db.system.profile.find().sort({ts: -1}).limit(5)
   ```

2. **Index Issues:**
   ```javascript
   // Check index usage
   db.collection.aggregate([{$indexStats: {}}])
   ```

## Environment-Specific Testing

### Development Environment
- Uses local MongoDB without authentication
- Default database: `payment_db`
- No SSL/TLS required

### Production Environment
- Uses MongoDB with authentication
- SSL/TLS enabled
- Connection pooling configured
- Monitoring enabled

## Test Checklist

- [ ] MongoDB connection successful
- [ ] All collections created with proper indexes
- [ ] Schema validation working
- [ ] CRUD operations functional
- [ ] Connection pooling working
- [ ] Performance within acceptable limits
- [ ] Backup and restore procedures tested
- [ ] Authentication working (if enabled)
- [ ] Monitoring and logging configured

## Additional Resources

- [MongoDB Documentation](https://docs.mongodb.com/)
- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [Docker Compose Documentation](https://docs.docker.com/compose/)

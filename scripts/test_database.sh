#!/bin/bash

# Script to test database functionality
# Usage: ./scripts/test_database.sh

set -e

echo "===== Tranzure Database Tests ====="

# Get directory of this script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

echo "Project root: $PROJECT_ROOT"
cd "$PROJECT_ROOT"

# Check if Docker is available
if command -v docker &> /dev/null; then
    echo "🐳 Docker detected. Starting MongoDB container..."

    # Start MongoDB container
    docker-compose up -d mongodb

    # Wait for MongoDB to be ready
    echo "⏳ Waiting for MongoDB to be ready..."
    sleep 10

    # Test MongoDB connection
    echo "🔍 Testing MongoDB connection..."
    if docker exec tranzure-mongodb-1 mongosh --eval "db.adminCommand('ping')" &> /dev/null; then
        echo "✅ MongoDB connection successful"
    else
        echo "❌ MongoDB connection failed"
        exit 1
    fi

    # Initialize database
    echo "🔧 Initializing database..."
    docker exec tranzure-mongodb-1 mongosh payment_db < scripts/mongo-init.js

    # Test collections
    echo "📋 Testing collections..."
    docker exec tranzure-mongodb-1 mongosh payment_db --eval "
        print('Collections:');
        db.getCollectionNames().forEach(function(collection) {
            print('- ' + collection);
        });

        print('\\nIndexes on users collection:');
        db.users.getIndexes().forEach(function(index) {
            print('- ' + index.name);
        });
    "

    echo "✅ Database tests completed successfully"

else
    echo "⚠️  Docker not available. Please install Docker to run database tests."
    exit 1
fi

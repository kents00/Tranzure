// Switch to the payment database
db = db.getSiblingDB('payment_db');

// Create collections with validation schemas
db.createCollection('users', {
    validator: {
        $jsonSchema: {
            bsonType: 'object',
            required: ['user_id', 'email', 'password_hash', 'role', 'status'],
            properties: {
                user_id: { bsonType: 'string' },
                email: { bsonType: 'string' },
                password_hash: { bsonType: 'string' },
                role: { enum: ['user', 'admin'] },
                status: { enum: ['active', 'banned'] }
            }
        }
    }
});

// Create indexes
db.users.createIndex({ email: 1 }, { unique: true });
db.users.createIndex({ user_id: 1 }, { unique: true });

db.createCollection('wallets');
db.wallets.createIndex({ user_id: 1 });
db.wallets.createIndex({ wallet_id: 1 }, { unique: true });

db.createCollection('transactions');
db.transactions.createIndex({ from_user_id: 1 });
db.transactions.createIndex({ to_user_id: 1 });
db.transactions.createIndex({ tx_id: 1 }, { unique: true });

db.createCollection('kyc_verifications');
db.kyc_verifications.createIndex({ user_id: 1 });
db.kyc_verifications.createIndex({ kyc_id: 1 }, { unique: true });

db.createCollection('sessions');
db.sessions.createIndex({ user_id: 1 });
db.sessions.createIndex({ session_id: 1 }, { unique: true });
db.sessions.createIndex({ expires_at: 1 }, { expireAfterSeconds: 0 });

db.createCollection('audit_logs');
db.audit_logs.createIndex({ user_id: 1 });
db.audit_logs.createIndex({ target_type: 1, target_id: 1 });

print('Database initialization completed');

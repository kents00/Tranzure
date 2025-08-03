### Sprint 0: Preparation & Architecture (1 week)

**Objectives:** Establish foundation, align vision, configure environments.


1. **Requirement Analysis & User Stories**

    - Workshop with stakeholders to refine features:

        - User onboarding: registration, email/phone verification.

        - Authentication: password reset, multi-factor support.

        - Payment flows: add/remove payment methods, send/receive funds, view history.

        - Admin: user management, transaction moderation.

        - Non-functional: TLS, OWASP Top 10 compliance, 99.9% uptime.

    - Document 10+ detailed user stories with acceptance criteria (e.g., “Given a verified email, when I provide a valid password, then I receive a JWT and am redirected”).

2. **Tech Stack & Repo Initialization**

    - **Backend:**

        - Initialize Go module (`go mod init github.com/yourorg/payapp`).

        - Install Fiber, GORM, JWT libraries.

        - Scaffold directory structure:

            ```
            /cmd/server
            /internal/app
            /internal/models
            /internal/controllers
            /internal/services
            /migrations
            /config
            ```

    - **Frontend (Android):**

        - Create Android Studio project with Kotlin, minimum SDK 21.

        - Setup MVC folder structure:

            ```
            /model
            /view
            /controller
            /network
            /utils
            ```

    - Initialize GitHub repositories with README, .gitignore, and branch protection rules.

3. **High‑Level Architecture & Diagrams**

- Draw system context diagram (Android , Admin Dashboard ⇄ API Gateway ⇄ Fiber app ⇄ MongoDB).
```mermaid
graph TD

  %% Clients

  A[Android App]

  B[Admin Dashboard Web]



  %% Gateway

  G[API Gateway]



  %% Backend

  C[Fiber App - Golang]



  %% Core Services

  D[(MongoDB)]

  E[KYC Provider]

  F[Crypto Wallet Service]

  H[Payment Processor]

  I[Email/SMS Service]



  %% Connections

  A <--> G

  B <--> G

  G --> C

  C --> D

  C --> E

  C --> F

  C --> H

  C --> I

```
- Create a sequence diagram for core flows.
```mermaid
sequenceDiagram
  participant User as Android App
  participant Gateway as API Gateway
  participant Backend as Fiber App (Golang)
  participant DB as MongoDB
  participant KYC as KYC Provider
  participant Wallet as Crypto Wallet Service

  %% === Registration Flow ===
  User->>Gateway: POST /register (email, password)
  Gateway->>Backend: Validate input
  Backend-->>Gateway: Validation OK
  Gateway->>Backend: Forward /register
  Backend->>DB: Check if user exists
  DB-->>Backend: Not found
  Backend->>Backend: Hash password
  Backend->>DB: Save user (email, hashed password, status: unverified)
  DB-->>Backend: Success
  Backend->>Backend: Generate JWT token
  Backend-->>Gateway: Registration success + JWT
  Gateway-->>User: 201 Created + token

  %% === KYC Submission Flow ===
  User->>Gateway: POST /submit-kyc (ID, selfie)
  Gateway->>Backend: Validate file + input
  Backend-->>Gateway: OK
  Gateway->>Backend: Forward /submit-kyc
  Backend->>KYC: Send KYC data to provider
  KYC-->>Backend: Submission accepted (status: pending)
  Backend->>DB: Store KYC submission (status: pending)
  Backend-->>Gateway: KYC submitted
  Gateway-->>User: KYC status: pending

  %% === KYC Callback / Polling Flow ===
  KYC-->>Backend: Callback /poll KYC status (approved/rejected)
  Backend->>DB: Update KYC status
  DB-->>Backend: OK

  %% === Wallet Connect Flow ===
  User->>Gateway: GET /wallet-nonce
  Gateway->>Backend: Request nonce
  Backend->>DB: Create + store nonce for user
  DB-->>Backend: Nonce saved
  Backend-->>Gateway: Return nonce
  Gateway-->>User: Show nonce

  User->>Gateway: POST /connect-wallet (wallet_address, signed_nonce)
  Gateway->>Backend: Forward to backend
  Backend->>Wallet: Verify signature with public address
  Wallet-->>Backend: Valid signature
  Backend->>DB: Link wallet to user profile
  DB-->>Backend: Success
  Backend-->>Gateway: Wallet connected
  Gateway-->>User: Success

```

- **Crypto Payment**
```mermaid
sequenceDiagram
  participant User as Android App
  participant Gateway as API Gateway
  participant Backend as Fiber App
  participant Wallet as Crypto Wallet Service
  participant DB as MongoDB

  User->>Gateway: POST /send-crypto (to, amount, token)
  Gateway->>Backend: Forward /send-crypto
  Backend->>Wallet: Broadcast transaction
  Wallet-->>Backend: Transaction hash
  Backend->>DB: Log TX with hash, sender, receiver
  Backend-->>Gateway: Payment success
  Gateway-->>User: Show TX hash and status

```
- **Fiat Payment**
```mermaid

sequenceDiagram
  participant User as Android App
  participant Gateway as API Gateway
  participant Backend as Fiber App (Golang)
  participant Processor as Payment Processor (Stripe/PayPal)
  participant DB as MongoDB
  participant Recipient as Recipient User

  User->>Gateway: POST /send-fiat (to_user_id, amount, currency)
  Gateway->>Backend: Forward request
  Backend->>Processor: Initiate payment (charge user)
  Processor-->>Backend: Payment success + txn_id
  Backend->>DB: Log transaction (sender, recipient, amount, txn_id)
  Backend->>DB: Update sender & recipient fiat balances
  Backend-->>Gateway: Send confirmation
  Gateway-->>User: Show payment success (TXN ID)
  Note over Backend,Recipient: Recipient sees balance updated in app

```
- **Withdrawal to Bank**
```mermaid

sequenceDiagram
  participant User as Android App
  participant Gateway as API Gateway
  participant Backend as Fiber App (Golang)
  participant Processor as Payment Processor (Stripe/PayPal)
  participant DB as MongoDB
  participant Bank as User’s Bank

  User->>Gateway: POST /withdraw-fiat (amount, bank_account_id)
  Gateway->>Backend: Forward request
  Backend->>DB: Check user balance
  alt Sufficient funds
    Backend->>Processor: Initiate payout to bank
    Processor-->>Backend: Payout initiated (txn_id, status)
    Backend->>DB: Deduct balance and log transaction
    Backend-->>Gateway: Withdrawal success (txn_id)
    Gateway-->>User: Show confirmation
  else Insufficient funds
    Backend-->>Gateway: Reject with "Insufficient Balance"
    Gateway-->>User: Show error
  end

```
- **Cash-In Bank**
```mermaid
sequenceDiagram
  participant User as Android App
  participant Gateway as API Gateway
  participant Backend as Fiber App (Golang)
  participant Processor as Payment Processor (Bank/Stripe/GCash)
  participant DB as MongoDB
  participant Bank as User’s Bank

  User->>Gateway: POST /cash-in (amount, bank_account_id)
  Gateway->>Backend: Forward request
  Backend->>Processor: Initiate pull from bank (e.g., debit PHP 1000)
  Processor-->>Bank: Debit from user’s bank account
  Bank-->>Processor: Bank confirms transfer
  Processor-->>Backend: Payment success (txn_id)
  Backend->>DB: Credit user's fiat balance
  Backend->>DB: Log transaction (cash-in)
  Backend-->>Gateway: Cash-in success
  Gateway-->>User: Show updated balance + confirmation
```
- **P2P Transfer (Fiat or Crypto)**
```mermaid
sequenceDiagram
  participant Sender as Sender (Android App)
  participant Gateway as API Gateway
  participant Backend as Fiber App
  participant DB as MongoDB
  participant Receiver as Recipient (Android App)

  Sender->>Gateway: POST /transfer (recipient_id, amount, type)
  Gateway->>Backend: Forward transfer request
  Backend->>DB: Check sender balance
  alt Sufficient balance
    Backend->>DB: Deduct from sender
    Backend->>DB: Credit to recipient
    Backend->>DB: Log transaction (TXN ID, type: P2P)
    Backend-->>Gateway: Transfer success
    Gateway-->>Sender: Show confirmation
    Note over Backend,Receiver: Recipient sees updated balance
  else Insufficient balance
    Backend-->>Gateway: Transfer failed
    Gateway-->>Sender: Show error
  end
```
- Define data flow for secrets: where to store API keys, TLS certs (use Vault or `.env`).
- Types of Secrets in Tranzure

| Type                     | Examples                                              |
| ------------------------ | ----------------------------------------------------- |
| **API Keys**             | KYC Provider (e.g. Sumsub), Email/SMS service, Stripe |
| **Private Keys**         | For custodial crypto wallets (if applicable)          |
| **TLS Certificates**     | For HTTPS (SSL/TLS) termination                       |
| **JWT Signing Secrets**  | Used for generating/verifying tokens                  |
| **Database Credentials** | MongoDB URI                                           |
| **Admin Tokens**         | Dashboard authentication                              |

**Secret Data Flow and Storage Strategy**

**Where to Store Secrets (for Each Component)**

| Component                      | Secret Type                        | Secure Storage Strategy                                                                        |
| ------------------------------ | ---------------------------------- | ---------------------------------------------------------------------------------------------- |
| **Fiber App (Golang backend)** | API keys, JWT secret, DB URI       | `.env` file locally, injected via **Docker env**, or **Vault**                                 |
| **API Gateway**                | TLS certs, API keys for validation | Use **reverse proxy config** (e.g. NGINX + mounted certs), secure Vault                        |
| **MongoDB**                    | Connection string                  | Stored in **Fiber env** only                                                                   |
| **TLS/HTTPS certs**            | SSL cert, private key              | Store in mounted volume or use **Let's Encrypt** auto-renew with **Certbot**                   |
| **Crypto Wallets (custodial)** | Private keys or mnemonic           | **NEVER in source code**, store in **HashiCorp Vault**, **AWS KMS**, or **GCP Secret Manager** |

**Secret Handling Flow :**

```mermaid
flowchart TD
  subgraph Secrets Store
    VAULT[Secrets Manager e.g. HashiCorp Vault, AWS Secrets Manager]
    ENV[.env or Docker Secrets dev/staging]
  end

  subgraph Services
    FIBER[Fiber App - Golang]
    GATEWAY[API Gateway]
    ADMIN[Admin Panel]
    WORKERS[Background Jobs]
  end

  VAULT --> FIBER
  VAULT --> GATEWAY
  ENV --> FIBER
  ENV --> GATEWAY
  FIBER --> DB[(MongoDB)]
  FIBER --> KYC[KYC API]
  FIBER --> EMAIL[Email/SMS Service]
  FIBER --> STRIPE[Stripe/PayPal]
```

**Best Practices (Secrets Handling)**

**Development & Deployment:**

|Best Practice|How|
|---|---|
|**No hardcoded secrets**|Never commit `.env` or config files with secrets to Git|
|**Use `.env` in dev only**|For local dev, store secrets in `.env`, use `godotenv` in Go|
|**Docker Secrets**|Use `docker secrets` or `docker-compose.override.yml` for staging/prod|
|**Kubernetes**|Use `Secrets` with RBAC + encrypted volumes|
|**HashiCorp Vault (recommended for scaling)**|Secure secret lifecycle + dynamic keys per environment|

**Example `.env` File (Local Dev Only):**

```
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=supersecretjwtkey
KYC_API_KEY=your-sumsub-key
STRIPE_SECRET_KEY=your-stripe-secret
EMAIL_API_KEY=your-mailgun-key
```

> Important: Use `.env.example` (without secrets) for version control.

**TLS Certificate Flow**

**For production HTTPS:**

|Option|How to Manage|
|---|---|
|**Reverse proxy (e.g. NGINX)**|TLS termination handled by NGINX, certs in `/etc/ssl/`|
|**Let's Encrypt + Certbot**|Auto-renew + mounted volumes|
|**Cloud provider TLS**|Use AWS ACM, GCP Certificate Manager, etc.|
Always redirect HTTP → HTTPS on the API Gateway or Load Balancer.
Tools Recommendation

| Tool                                         | Purpose                                   |
| -------------------------------------------- | ----------------------------------------- |
| **Dotenv / godotenv**                        | Load `.env` into Go backend               |
| **HashiCorp Vault**                          | Secure, role-based secrets management     |
| **AWS Secrets Manager / GCP Secret Manager** | Cloud-native alternative to Vault         |
| **Docker Secrets**                           | Encrypt secrets for containerized deploys |
| **Certbot**                                  | Automated TLS with Let's Encrypt          |

4. **Development Environment & CI/CD**
    - **Local Setup:** Docker Compose YAML with services:
        - `app` (Go + Fiber), `db` (Mongo db), `adminer`.
    - **CI Pipeline:** GitHub Actions to run:
        - `go fmt`, `go vet`, `golangci-lint` (backend).
        - `./gradlew lint`, `./gradlew test` (Android).
    - Setup staging branch that deploys to a test VPS via Docker.

**Deliverables:** Requirements doc, architecture diagrams, skeleton codebases, CI green badge.

---

### Sprint 1: Core Data Models & Persistence (1–2 weeks)

**Objectives:** Model database, implement ORM, basic CRUD for Users.

1. **Database Design & Migrations**
    - Finalize ERD with fields, types, constraints:
**Users**

| Field           | Type                    | Notes          |
| --------------- | ----------------------- | -------------- |
| `user_id`       | UUID                    | Primary Key    |
| `email`         | String                  | Unique         |
| `password_hash` | String                  | bcrypt         |
| `role`          | Enum(`user`, `admin`)   | Access control |
| `status`        | Enum(`active`,`banned`) | Account state  |
| `created_at`    | Timestamp               |                |
| `updated_at`    | Timestamp               |                |
**kyc verifications**

| Field              | Type                                    | Notes                  |
| ------------------ | --------------------------------------- | ---------------------- |
| `kyc_id`           | UUID                                    | Primary Key            |
| `user_id`          | UUID                                    | Foreign key to `users` |
| `status`           | Enum(`pending`, `approved`, `rejected`) |                        |
| `document_type`    | String                                  | e.g., passport         |
| `submitted_at`     | Timestamp                               |                        |
| `reviewed_at`      | Timestamp                               |                        |
| `rejection_reason` | String                                  | Optional               |

**wallets**

| Field        | Type                   | Notes                              |
| ------------ | ---------------------- | ---------------------------------- |
| `wallet_id`  | UUID                   | Primary Key                        |
| `user_id`    | UUID                   | FK to `users`                      |
| `type`       | Enum(`crypto`, `fiat`) |                                    |
| `currency`   | String                 | e.g., `USD`, `BTC`, `ETH`          |
| `address`    | String                 | Crypto address or bank account ref |
| `balance`    | Decimal                | Optional for fiat                  |
| `is_primary` | Boolean                |                                    |
**transactions**

| Field          | Type                                 | Notes                    |
| -------------- | ------------------------------------ | ------------------------ |
| `tx_id`        | UUID                                 | Primary Key              |
| `from_user_id` | UUID                                 | Sender                   |
| `to_user_id`   | UUID                                 | Recipient                |
| `wallet_id`    | UUID                                 | Wallet used              |
| `type`         | Enum(`fiat`, `crypto`)               |                          |
| `currency`     | String                               |                          |
| `amount`       | Decimal                              |                          |
| `status`       | Enum(`pending`, `success`, `failed`) |                          |
| `tx_hash`      | String                               | On-chain hash for crypto |
| `metadata`     | JSON                                 | Escrow, fees, etc.       |
| `created_at`   | Timestamp                            |                          |

**sessions**

| Field         | Type      | Notes |
| ------------- | --------- | ----- |
| `session_id`  | UUID      |       |
| `user_id`     | UUID      |       |
| `ip_address`  | String    |       |
| `device_info` | String    |       |
| `login_at`    | Timestamp |       |
| `expires_at`  | Timestamp |       |

**audit_logs**

| Field         | Type      | Notes                      |
| ------------- | --------- | -------------------------- |
| `log_id`      | UUID      |                            |
| `user_id`     | UUID      | Who performed the action   |
| `action`      | String    | Description                |
| `target_type` | String    | "kyc", "transaction", etc. |
| `target_id`   | UUID      |                            |
| `created_at`  | Timestamp |                            |


```mermaid

erDiagram
  users ||--o{ kyc_verifications : has
  users ||--o{ wallets : owns
  users ||--o{ transactions : sends
  users ||--o{ sessions : creates
  users ||--o{ audit_logs : generates
  transactions }o--|| wallets : uses

  users {
    UUID user_id PK
    String email
    String password_hash
    String role
    String status
  }

  kyc_verifications {
    UUID kyc_id PK
    UUID user_id FK
    String status
    String document_type
  }

  wallets {
    UUID wallet_id PK
    UUID user_id FK
    String type
    String currency
    String address
  }

  transactions {
    UUID tx_id PK
    UUID from_user_id FK
    UUID to_user_id FK
    UUID wallet_id FK
    String currency
    String type
    Decimal amount
    String status
  }

  sessions {
    UUID session_id PK
    UUID user_id FK
    String ip_address
  }

  audit_logs {
    UUID log_id PK
    UUID user_id FK
    String action
    String target_type
  }

```

- Write initial SQL migration scripts or use GORM AutoMigrate with manual review.
2. **Backend Model Implementation**
    - Create Go structs in `internal/models` with GORM tags and JSON annotations.
    - Set up config loader (Viper) for DB connection strings.
    - Unit test model validation logic (e.g., non-null constraints).
3. **Service & Repository Layer**
    - Define repository interfaces (e.g., `UserRepo.Create(ctx, User) error`).
    - Implement repository using GORM in `internal/services/user_service.go`.
    - Write unit tests using SQLite in-memory for repositories.
4. **Initial Controller & Routing**
    - Setup Fiber app in `cmd/server/main.go`, load config, connect DB.
    - Register `/health` endpoint returning JSON `{status: "ok"}`.
    - Add `/users` POST endpoint to create users; use service layer.
    - Middleware: logging (Fiber logger), recover (Fiber recover).
5. **Frontend Model Mapping**
    - Create Kotlin data classes in `/model`: `User`, `Wallet`, etc., matching JSON.
    - Setup Retrofit + Moshi or Gson for JSON parsing; define `ApiService` interface.
6. **Stub UI Screens**
    - Add LoginActivity, RegisterActivity layouts (XML) with basic form fields.
    - Controller stubs: button click logs to console.
**Checkpoints:**

- DB tables exist and CRUD via Go tests.
- POST /users works (tested via Postman).
- Kotlin Retrofit setup can fetch from `/health`.
---
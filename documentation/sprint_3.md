### 💳 Sprint 3: Payment Method Management (1–2 weeks)

**Objectives:** CRUD operations for payment methods, secure storage.

1. **Backend Model & Migration**

    - Update `PaymentMethod` struct with fields: `provider` (Visa/Mastercard/Bank), `masked_number`, `expiry_date`, `meta JSON`, `user_id FK`, `created_at`.

    - Migration script to add table with `CHECK` constraints on expiry format.

2. **Controllers & Services**

    - Endpoints in `payment_method_controller.go`:

        - `GET /payment-methods` → list current user’s methods.

        - `POST /payment-methods` → validate payload, encrypt sensitive fields using AES-256, store.

        - `DELETE /payment-methods/:id` → soft-delete (set `is_active=false`).

    - Service layer: ensure only owner can CRUD their methods.

3. **Encryption & Key Management**

    - Generate symmetric key stored via environment variable or Vault.

    - Implement util functions for `encrypt()`, `decrypt()` in `internal/utils/crypto.go`.

    - Write tests to ensure decryption matches plaintext.

4. **Android UI & Controller**

    - **Add Payment Method Screen**:

        - Form with provider dropdown, number (input mask), expiry picker, CVV.

        - Client-side validation (length, Luhn check).

    - **Payment Method List**:

        - RecyclerView with CardView items showing masked number & expiry.

        - Swipe-to-delete with confirmation dialog.

5. **Networking & Error Handling**

    - Retrofit calls for GET, POST, DELETE; use sealed classes for response states.

    - Global ErrorHandler: map 400/401/500 to user-friendly messages.


**Acceptance Criteria:**

- CRUD flows tested end-to-end (Postman + Android emulator).

- Encryption verified via DB inspection.
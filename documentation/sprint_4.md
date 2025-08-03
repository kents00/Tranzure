### 💸 Sprint 4: Core Payment Flow (2 weeks)

**Objectives:** Implement the end-to-end send/receive payment mechanics and status tracking.

1. **Model & Migration Updates**

    - Extend `Transaction` struct:

        - Add `reference_id` (UUID), `description` (string), `fee` (decimal).

        - Ensure `status` enum includes `pending`, `completed`, `failed`, `cancelled`.

    - Create migrations for new fields and index on `from_wallet` and `to_wallet`.

2. **Service Layer & Business Logic**

    - Implement `TransactionService.Create(ctx, fromID, toID, amount, methodID)`:

        - Check sender balance; apply reserved holds; calculate fees.

        - Insert record with `status=pending` and publish event to `transaction_queue`.

    - Create worker in `internal/services/transaction_worker.go`:

        - Listen to queue (e.g., Redis or in-memory pub/sub) to settle transactions.

        - On success: deduct/add balances atomically in a DB transaction; update status to `completed`.

        - On failure: release holds; update status to `failed`; log error.

3. **Controller & Routing**

    - Add endpoints in `transaction_controller.go`:

        - `POST /transactions` → validate input, call `TransactionService.Create`, return `reference_id` and initial `status`

        - `GET /transactions/:id/status` → return current status

        - Webhook stub `POST /transactions/webhook` for external settlement callbacks

    - Middleware: ensure only authenticated users access these endpoints.

4. **Frontend UI & Controller (Android)**

    - **Send Payment Screen**:

        - Form fields: payee selection (search/autocomplete), amount input (decimal keyboard), method dropdown.

        - Confirmation dialog summarizing `amount + fee` and `to` details.

    - **Payment Status Screen**:

        - Poll `/transactions/:id/status` every 5s until `completed` or `failed` (max 1 min).

        - Display animated loader; on final status show success/failure with details.

5. **Error Handling & Edge Cases**

    - Handle network timeouts, server errors with retry/backoff.

    - Display clear UX for `insufficient funds`, `invalid method`, `timeout`.

    - Idempotency: client should handle duplicate submissions safely (use `reference_id`).


**Acceptance Criteria:**

- Transactions move from `pending`→`completed` in DB and UI reflects each state.

- Failure paths (insufficient funds, worker error) update status and surface clear messages.

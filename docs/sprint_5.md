### 📊 Sprint 5: Transaction History & Analytics (1–2 weeks)

**Objectives:** Provide users with paginated history and summary metrics.

1. **Backend Pagination & DTOs**

    - Extend `TransactionController`:

        - `GET /transactions` with query params `?page=&limit=&sort=&from=&to=` filters by date range and pagination.

    - Create `TransactionHistoryDTO` and `DashboardSummaryDTO`:

        - History item: `reference_id`, `type` (sent/received), `amount`, `fee`, `status`, `timestamp`, `counterparty`.

        - Summary: `total_sent`, `total_received`, `average_fee` for given period.

    - Optimize queries with `OFFSET`/`LIMIT` and appropriate indexes; consider keyset pagination for large datasets.

2. **Analytics Service**

    - Implement `AnalyticsService.GetSummary(userID, period)`:

        - Aggregate sums and averages

        - Cache results for 5 minutes in Redis

3. **Endpoints & Security**

    - `GET /transactions/history` → returns paginated `TransactionHistoryDTO`

    - `GET /dashboard/summary?period={daily,monthly,yearly}` → returns `DashboardSummaryDTO`

4. **Android UI & Controller**

    - **History Screen**:

        - RecyclerView with `PagedListAdapter` for infinite scrolling.

        - Date filter UI (calendar pickers) and sort toggle (asc/desc).

    - **Dashboard Screen**:

        - Card views showing `total_sent`, `total_received`, `average_fee`, with icons.

        - Pull-to-refresh to update metrics.

5. **Testing & Performance**

    - Backend integration tests for pagination filters and summary accuracy.

    - Android UI tests (Espresso) for scrolling and date filters.


**Acceptance Criteria:**

- Users can browse history by page, filter by date, sort results.

- Dashboard shows correct summary metrics matching DB aggregates.


### 🚀 Sprint 7: Testing, QA & Documentation (2 weeks)

**Objectives:** Achieve high quality via automated tests, manual QA, and thorough docs.

1. **Automated Tests**

    - **Backend:**

        - Unit tests for services: user, auth, transactions, analytics.

        - Integration tests with testcontainers (Postgres + Redis).

        - API contract tests using Postman/Newman.

    - **Frontend:**

        - Unit tests for controllers/ViewModels with JUnit + Mockito.

        - Instrumentation tests (Espresso) for critical flows: login, add method, send payment, view history.

2. **Manual QA**

    - Test on multiple device emulators and real devices (API 21–30).

    - Security review: pen test basic auth flow, SQLi, XSS (in case of embedded webviews).

    - UX review: form layouts on different screen sizes.

3. **Documentation**

    - **Backend:** API reference generated via Swagger/OpenAPI; publish to internal Confluence or GitHub Pages.

    - **Android:** Code docs with KDoc; README with setup and architecture overview.

    - Developer onboarding guide: how to run locally, code style, commit conventions.

4. **Acceptance & Sign‑off**

    - Demo to stakeholders covering all core features.

    - Collect feedback and log new backlog items.


**Deliverables:** Test reports, Swagger docs, QA checklist, signed-off MVP.
### 🔍 Sprint 6: Validation, Error Handling & Localization (1 week)

**Objectives:** Harden validation, unify error flows, support multiple locales.

1. **Backend Validation**

    - Integrate `fiber/middleware/validator` or use `go-playground/validator`:

        - Validate request structs for all endpoints (`required`, `email`, `min`, `max`, `pattern`).

    - Centralize error formatting middleware:

        - Convert all errors to `{code, message, field_errors}` JSON.

2. **Android Error Framework**

    - Create `ErrorHandler` singleton:

        - Maps API error codes to user messages via `strings.xml`.

    - Implement field-level error highlighting in forms.

3. **Localization**

    - Backend: return `error_code` keys.

    - Android: `strings.xml` for English + placeholder for other languages (e.g., Tagalog).

    - Populate Tagalog translations for core UI texts and error messages.

4. **Edge Case Flows**

    - Test invalid JWT, expired tokens → redirect to login.

    - Test invalid input formats (card number, amount, email).

    - simulate network failures → show retry dialogs.


**Acceptance Criteria:**

- All user flows validated; no unhandled errors.

- App displays localized strings for core screens and errors.


### Sprint 2: Authentication & Security (2 weeks)

**Objectives:** Secure signup/login, token management, session security.

1. **Backend Auth Implementation**

    - Add `bcrypt` for password hashing.

    - Endpoints in `controllers/auth_controller.go`:

        - `POST /auth/register` → hash password, create user, return 201.

        - `POST /auth/login` → verify hash, issue access (15m) & refresh (7d) JWT tokens.

        - `POST /auth/refresh` → validate refresh token, issue new access.

    - JWT config: secret rotation strategy, claims including `sub` (user ID), `iat`, `exp`.

    - Middleware: `Authenticate` checks `Authorization: Bearer` header and sets `ctx.Locals("user_id")`.

2. **Token Storage & Revocation**

    - Create `RefreshToken` model/table: token hash, user_id, expires_at.

    - On login, store refresh token; on logout or rotate, delete old tokens.

3. **Frontend Auth Flow**

    - Retrofit `AuthService` with register, login, refresh endpoints.

    - Securely store tokens in EncryptedSharedPreferences.

    - Interceptor for OkHttp to attach `Authorization` header.

    - Activity results: success navigates to MainActivity; failure shows error toast.

4. **Security Hardening**

    - Enforce HTTPS for all requests; disable cleartext in `network_security_config.xml`.

    - Implement rate limiting in Fiber (e.g., `github.com/gofiber/fiber/middleware/limiter`).

    - CORS policy: allow only Android client origin or use API keys.

5. **Testing**

    - Backend unit tests: invalid password, expired token, unauthorized access.

    - Android instrumented tests: login flow, token persistence.


**Deliverables:** Secure auth endpoints, token lifecycle, Android login/register flows.
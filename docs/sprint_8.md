### 📦 Sprint 8: Deployment, Monitoring & Ops (1 week)

**Objectives:** Move to production, establish monitoring, and prepare release.

1. **Containerization & Deployment**

    - Finalize `Dockerfile` for Go app: multi-stage build, healthcheck.

    - Docker Compose for staging and production override files (with real secrets via Vault).

    - Helm chart or Kubernetes manifests if using k8s.

2. **CI/CD to Production**

    - GitHub Actions workflows:

        - On `main` merge: build images, push to registry, deploy to staging.

        - Manual approval gate → deploy to production.

3. **Monitoring & Alerting**

    - Integrate Prometheus exporters in Fiber (metrics middleware).

    - Grafana dashboard for request rates, error rates, latency.

    - Alertmanager rules for high error rate (>5% over 1m), latency (>500ms).

    - Sentry integration for panic/error tracking.

4. **Android Release**

    - Generate signed APK/AAB; configure versioning policy.

    - Upload to Google Play Console (internal test track → closed beta).

    - Release notes and test instructions.

5. **Post‑Launch Support**

    - On-call schedule for rapid response to incidents.

    - Run load test (k6) against production endpoints; compare with staging.

    - Collect user metrics (Google Analytics, Firebase) for crash reporting and usage.


**Acceptance Criteria:**

- Production environment serving live traffic.

- Monitoring alerts configured and tested.

- Android beta build available to test users.

**Ongoing Iterations & Backlog**

- Add refunds, recurring payments, currency conversion.

- Enhance security: HSM for keys, SOC2 audit prep.

- UX polishing: animations, dark mode, offline mode.

- Scale enhancements: read replicas, query optimization, pub/sub migration.


**Next Steps:** Prioritize backlog based on user feedback and compliance requirements.
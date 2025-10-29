# Personal Website

This repo hosts the code and infrastructure behind [shumphreys.com](https://shumphreys.com).

## Structure

- `tf/` — Terraform to create the GCP project, enable APIs, App Engine app, CI/CD service account and state bucket.
- `frontend/` — Vite + React + TypeScript SPA.
- `backend/` — Go HTTP server. Embeds `frontend/dist` into the binary and provides `/api`.
- `backend/app.yaml` — App Engine config.

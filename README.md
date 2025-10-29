# GCP App Engine: Go + Vite React

This repo deploys a Go backend (App Engine Standard, Go 1.22) serving an embedded Vite React frontend.

## Structure

- `tf/` — Terraform to create the GCP project, enable APIs, App Engine app, CI/CD service account and state bucket.
- `frontend/` — Vite + React + TypeScript SPA.
- `backend/` — Go HTTP server. Embeds `frontend/dist` into the binary and provides `/api`.
- `backend/app.yaml` — App Engine config.

## Local development

In one terminal (frontend):

```bash
cd frontend
npm ci
npm run dev
```

In another (backend API only):

```bash
cd backend
PORT=8080 go run .
```

The frontend dev server will proxy requests to `/api/*` directly to the backend if configured; otherwise edit fetch URLs to `http://localhost:8080/api/*` during local dev.

## Build and run locally (single binary)

```bash
cd backend
make build
./server
```

## Deploy (GitHub Actions)

The workflow `.github/workflows/deploy.yml`:

1. Builds the frontend.
2. Copies `frontend/dist` to `backend/web/dist`.
3. Authenticates to GCP using a service account key secret `GCP_SA_KEY`.
4. Runs `gcloud app deploy` in `backend/`.

Set repository secrets:

- `GCP_SA_KEY` — JSON key for the `cicd-app-engine-deployer` service account created by Terraform.

Ensure App Engine is initialized by Terraform (`tf/appengine.tf`), and required APIs enabled via `tf/variables.tf`.

## Notes

- App Engine standard uses the `PORT` env var; the server honors it.
- You can adjust instance class and scaling in `backend/app.yaml`.
- If you prefer Cloud Build instead of GitHub Actions, you can add a `cloudbuild.yaml` to build and deploy from GCP.

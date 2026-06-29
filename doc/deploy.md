# Deployment Guide — SpotSync API

This guide explains how to deploy SpotSync API to **Render** from a **GitHub** repository.

---

## Prerequisites

- GitHub account
- Render account (https://render.com)
- NeonDB database (https://neon.tech)

---

## Step 1: Push Code to GitHub

### 1.1 Initialize Git (if not already done)

```bash
git init
git add .
git commit -m "Initial commit: SpotSync API"
```

### 1.2 Create GitHub Repository

1. Go to https://github.com/new
2. Name it `spotsync-api`
3. Make it **Public** (required for free Render tier)
4. Do NOT add README or .gitignore (you already have them)

### 1.3 Push to GitHub

```bash
git remote add origin https://github.com/YOUR_USERNAME/spotsync-api.git
git branch -M main
git push -u origin main
```

---

## Step 2: Set Up NeonDB Database

### 2.1 Create Database

1. Go to https://neon.tech and sign up/login
2. Create a new project (e.g., `spotsync-db`)
3. Copy the **connection string** from Dashboard → Connection Details

### 2.2 Connection String Format

```
postgresql://user:password@host.neon.tech/dbname?sslmode=require
```

---

## Step 3: Deploy to Render

### 3.1 Create New Web Service

1. Go to https://render.com/dashboard
2. Click **New +** → **Web Service**
3. Connect your GitHub repository
4. Select `spotsync-api` repository

### 3.2 Configure Build Settings

| Setting           | Value                           |
| ----------------- | ------------------------------- |
| **Name**          | `spotsync-api`                  |
| **Runtime**       | Go                              |
| **Build Command** | `go build -o app ./cmd/main.go` |
| **Start Command** | `./app`                         |

### 3.3 Set Environment Variables

Click **Advanced** → **Add Environment Variable**:

| Key          | Value                                                              | Description                       |
| ------------ | ------------------------------------------------------------------ | --------------------------------- |
| `DSN`        | `postgresql://user:password@host.neon.tech/dbname?sslmode=require` | NeonDB connection string          |
| `JWT_SECRET` | `your-random-secret-key`                                           | JWT signing secret                |
| `PORT`       | `10000`                                                            | Render assigns this automatically |

### 3.4 Deploy

1. Click **Create Web Service**
2. Wait for build to complete (2-5 minutes)
3. Your API will be live at: `https://spotsync-api.onrender.com`

---

## Step 4: Verify Deployment

### 4.1 Health Check

```bash
curl https://spotsync-api.onrender.com/health
```

Expected response:

```json
{
  "status": "ok"
}
```

### 4.2 Test API Endpoints

Use Postman to test:

- `POST https://spotsync-api.onrender.com/api/v1/auth/register`
- `POST https://spotsync-api.onrender.com/api/v1/auth/login`
- `GET https://spotsync-api.onrender.com/api/v1/zones`

---

## Step 5: Update Frontend (if needed)

Update your frontend API base URL to:

```
https://spotsync-api.onrender.com/api/v1
```

---

## Troubleshooting

### Build Fails

- Check that `go.mod` and `go.sum` are committed
- Verify Build Command: `go build -o app ./cmd/main.go`

### Database Connection Error

- Verify DSN is correct in Render environment variables
- Ensure NeonDB allows connections from Render's IP

### JWT Errors

- Make sure `JWT_SECRET` is set in Render environment variables

---

## Useful Commands

### View Logs (Render Dashboard)

1. Go to your Web Service in Render
2. Click **Logs** tab
3. View real-time logs

### Redeploy

Push to GitHub main branch → Render auto-deploys:

```bash
git add .
git commit -m "Update feature"
git push origin main
```

---

## Free Tier Limitations

| Resource        | Limit                    |
| --------------- | ------------------------ |
| **Build time**  | 50 minutes/month         |
| **RAM**         | 512 MB                   |
| **CPU**         | Shared                   |
| **Bandwidth**   | 100 GB/month             |
| **Sleep after** | 15 minutes of inactivity |

> 💡 **Tip:** Free tier spins down after 15 minutes of inactivity. First request after spin-down takes ~30 seconds.

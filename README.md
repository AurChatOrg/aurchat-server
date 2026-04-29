# Aurchat Backend
> This is the backend application for Aurchat, built with Next.js.

### 1. Prerequisites
- Bun >= 1.3.12

### 2. Installation & Running
```bash
# Clone the repository
git clone https://github.com/AurChatOrg/aurchat-server.git

# Install dependencies
bun install

# Start the development server
bun run dev

## Environment Variables (.env.local)
`env
# Auth secrret
BETTER_AUTH_SECRET=YOUR_AUTH_SECRET

# Database (SQLite Temp)
DB_FILE_NAME=YOUR_SQLITE_FILE_NAME

# Frontend URL
BETTER_AUTH_URL=YOUR_FRONTEND_URL

# Backend url&port
BACKEND_URL=APP_URL
LISTENING_PORT=APP_LISTENING_PORT
`

## Related Links
- **Frontend**: [Frontend Repo](https://github.com/AurChatOrg/aurchat-frontend)
- **Product Requirements**: [PRD Link](PRD.md)

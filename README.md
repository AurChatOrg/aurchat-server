# Aurchat Backend
> This is the backend application for Aurchat, built with Next.js.

### 1. Prerequisites
- Bun >= 1.3.12
- Ensure the backend service is running (http://localhost:3001)

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
`

## Related Links
- **Frontend**: [Frontend Repo](https://github.com/AurChatOrg/aurchat-frontend)
- **Product Requirements**: [PRD Link](PRD.md)

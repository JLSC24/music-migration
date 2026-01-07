# Frontend

Next.js frontend for Music Migration Tool.

## Setup

```bash
# Install dependencies
npm install

# Copy environment file
cp .env.example .env.local

# Edit .env.local with your API URL
# nano .env.local

# Run development server
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) in your browser.

## Project Structure

```
frontend/
├── src/
│   ├── app/
│   │   ├── layout.tsx          # Root layout
│   │   ├── page.tsx            # Home page
│   │   ├── dashboard/
│   │   │   └── page.tsx        # Dashboard
│   │   └── migrate/
│   │       └── page.tsx        # Migration page
│   ├── components/
│   │   └── Providers.tsx       # React Query provider
│   ├── lib/
│   │   └── api.ts              # API client
│   └── types/
│       └── index.ts            # TypeScript types
├── public/                     # Static assets
├── package.json
├── tsconfig.json
├── tailwind.config.js
└── next.config.js
```

## Available Scripts

```bash
# Development
npm run dev

# Build for production
npm run build

# Start production server
npm start

# Lint
npm run lint

# Type check
npm run type-check
```

## Features

- ✅ Modern UI with Tailwind CSS
- ✅ TypeScript for type safety
- ✅ React Query for data fetching
- ✅ Real-time progress updates (WebSocket ready)
- ✅ Responsive design
- ✅ OAuth integration ready

## Environment Variables

- `NEXT_PUBLIC_API_URL` - Backend API URL (default: http://localhost:8080)
- `NEXT_PUBLIC_WS_URL` - WebSocket URL for real-time updates

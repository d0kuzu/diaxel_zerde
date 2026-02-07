# Diaxel Frontend

## Requirements

- Node.js 20+

## Setup

1. Create `.env.local` based on `.env.example`.

2. Install dependencies.

3. Run dev server.

## Environment

- `NEXT_PUBLIC_API_BASE_URL`

If set, the Analytics demo will try to request:

- `GET /api/analytics/summary`
- `GET /api/analytics/timeseries`

from that base URL.

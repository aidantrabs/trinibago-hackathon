# KULTUR

**A first-timer's guide to Trinidad & Tobago's cultural festivals.**

*TrinbagoTech Hackathon 2026 · January 30 – February 2*

**Live Site**: [kultur-tt.app](https://kultur-tt.app)

---

## The Team

| Name | Role |
|:-----|:-----|
| **Aidan** | Backend Development |
| **Tevin** | Frontend Development |
| **Jahzara** | Content & Cultural Research |

---

## The Problem

> *"On October 30, 1884, British authorities opened fire on a peaceful Hosay procession in San Fernando. 22 people died, 120+ were wounded. It's one of the bloodiest events in Trinidad's colonial history.*
>
> *Most Trinis have never heard of it."*

Trinidad & Tobago is home to **over 20 cultural festivals** spanning African, Indian, Indigenous, European, and Chinese heritage. These festivals represent centuries of tradition, survival, and cultural fusion that make T&T unique in the Caribbean.

**But most people only know Carnival.**

The information about these festivals is:
- **Scattered** across Facebook groups, outdated government sites, and academic archives
- **Inaccessible** to first-timers who don't know the etiquette, dress codes, or what to expect
- **Invisible** to the diaspora trying to reconnect with their heritage
- **Missing** from tourist experiences seeking authenticity beyond the beach

---

## Features

### Festival Calendar
- Browse all festivals by month, region, or heritage type
- Interactive calendar with color-coded heritage types
- "This Week in T&T" section for upcoming festivals

### First-Timer's Guides
Each festival includes:
- History and cultural significance
- What you'll actually see, hear, and experience
- How to participate respectfully (dress code, etiquette)
- Practical info (parking, food, best viewing spots)

### Newsletter & Reminders
- Email subscription for festival updates
- Welcome emails for new subscribers
- Festival reminder notifications
- Weekly digest of upcoming festivals

### Community Memories
- Real stories from people who've attended
- Submit your own festival memories
- No account required

---

## Demo Highlights

Our demo showcases **10 festivals** across multiple heritage types:

| Festival | Heritage | Significance |
|:---------|:---------|:-------------|
| Hosay | Indian/Islamic | Solemn observance with 1884 massacre history |
| Santa Rosa Festival | Indigenous | Last organized First Peoples community |
| First Peoples Heritage Week | Indigenous | October celebration of Amerindian heritage |
| Divali | Indian/Hindu | Largest in English-speaking Caribbean |
| Emancipation Day | African | August 1 commemoration |
| Spiritual Baptist Liberation Day | African | March 30, end of religious persecution |
| Phagwa | Indian/Hindu | Spring festival of colors |
| Ramleela | Indian/Hindu | Open-air theatre at 35+ sites |
| Parang | Mixed | Christmas music tradition |
| Carnival | Mixed | The one everyone knows |

---

## Technical Stack

### Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                       PRODUCTION                            │
│                                                             │
│   Frontend (kultur-tt.app)  ◄── Vercel                      │
│        │                                                    │
│        │  fetch('https://kultur-api.fly.dev/api/...')       │
│        ▼                                                    │
│   Backend (kultur-api.fly.dev)  ◄── Fly.io                  │
│        │                                                    │
│        ▼                                                    │
│   PostgreSQL  ◄── Neon                                      │
│                                                             │
│   Email  ◄── Resend (noreply@kultur-tt.app)                 │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Stack

| Layer | Technology |
|:------|:-----------|
| Frontend | SvelteKit 2, Svelte 5, TailwindCSS 4, shadcn-svelte |
| Frontend Hosting | Vercel |
| Backend | Go 1.24, Echo framework |
| Backend Hosting | Fly.io |
| Database | PostgreSQL (Neon) |
| Email | Resend |
| Linting/Formatting | Biome |

### Key Endpoints

**Public API** (`https://kultur-api.fly.dev`):
- `GET /api/festivals` - List all festivals
- `GET /api/festivals/upcoming` - Upcoming festivals
- `GET /api/festivals/:slug` - Festival details
- `POST /api/subscribe` - Newsletter signup
- `POST /api/memories` - Submit a memory

See [docs/ROUTES.md](docs/ROUTES.md) for full API documentation.

---

## Documentation

| Document | Description |
|:---------|:------------|
| [docs/SETUP.md](docs/SETUP.md) | Local development setup guide |
| [docs/ROUTES.md](docs/ROUTES.md) | API routes reference |
| [docs/DEMO.md](docs/DEMO.md) | Email template testing guide |
| [docs/PLANNING.md](docs/PLANNING.md) | Original hackathon planning document |

---

## Quick Start

### Prerequisites
- Go 1.24+
- Node.js 20+
- PostgreSQL (or Docker)

### Local Development

```bash
# Clone the repo
git clone https://github.com/aidantrabs/trinbago-hackathon.git
cd trinbago-hackathon

# Backend (Terminal 1)
cd backend
cp .env.example .env  # Get real values from codeowner
go run ./cmd/server

# Frontend (Terminal 2)
cd frontend
npm install
npm run dev
```

See [docs/SETUP.md](docs/SETUP.md) for comprehensive setup instructions.

---

## Scripts

### Frontend

```bash
npm run dev       # Start dev server
npm run build     # Production build
npm run deploy    # Deploy to Vercel
npm run format    # Format with Biome
npm run lint      # Lint with Biome
```

### Backend

```bash
go run ./cmd/server  # Start server
fly deploy           # Deploy to Fly.io
```

---

## Why This Matters

### Cultural Preservation

Trinidad & Tobago's cultural landscape is extraordinary:

- **First Peoples Heritage** — The Santa Rosa First Peoples Community in Arima is the last organized Indigenous group in T&T
- **Hosay** — A solemn Islamic observance with tadjahs reaching 30 feet tall
- **Divali Nagar** — The largest Divali celebration in the English-speaking Caribbean
- **Spiritual Baptist Liberation Day** — Commemorating the 1951 repeal of religious persecution

### Diaspora Connection

**374,500 Trinbagonians** live abroad. This platform gives them a way to:
- Learn about festivals their parents attended
- Plan meaningful visits home around cultural events
- Share traditions with the next generation

### Tourism with Authenticity

**158,890 tourists** arrived in T&T between January and May 2025. Instead of just beaches and Carnival, they can experience authentic cultural celebrations.

---

## Vision

> *"Imagine every festival documented like this. Imagine diaspora kids learning about Hosay before visiting home. Imagine tourists experiencing something real.*
>
> *KULTUR. Your guide to Trinidad's cultural festivals."*

#  TravelSphere

 A full-stack travel destination discovery and trip planner built with the **Beego Framework (Go)**. Explore countries, discover attractions, and manage your personal travel wishlist 

---

##  Table of Contents

- [Project Overview](#-project-overview)
- [Tech Stack](#-tech-stack)
- [Features](#-features)
- [Project Structure](#-project-structure)
- [Wishlist Storage — In-Memory](#-wishlist-storage--in-memory)
- [Getting Started](#-getting-started)
- [Environment Variables](#-environment-variables)
- [Running the App](#-running-the-app)
- [Running Tests](#-running-tests)
- [API Routes Reference](#-api-routes-reference)
- [Country Slug Format](#-country-slug-format)
- [Screenshots](#-screenshots)

---

##  Project Overview

TravelSphere is a Beego MVC web application that lets users:

- Discover countries and their details (flag, capital, population, currency, languages)
- Explore tourist attractions and landmarks powered by OpenTripMap
- Manage a personal travel wishlist with statuses: `Planned` or `Visited`


The app uses **Server-Side Rendering (SSR)** for navigable pages and **AJAX partial updates** for dynamic interactions , no full page reloads for search, wishlist actions, or dashboard refresh.

---

##  Tech Stack

| Layer | Technology |
|---|---|
| Language | Go (Golang) |
| Web Framework | [Beego](https://beego.vip/) |
| Live Reload (Dev) | [Bee Tool](https://github.com/beego/bee) |
| Templating | Beego `.tpl` templates (SSR) |
| Frontend | Vanilla JavaScript + Fetch API (AJAX) |
| Countries API | [REST Countries v3.1](https://restcountries.com/) |
| Attractions API | [OpenTripMap](https://dev.opentripmap.org/) |
| Wishlist Storage | In-memory Go map (no database) |
| Testing | Testify |
| Containerization | Docker |

---

##  Features

###  Home Page
- Search destinations with AJAX (updates results without page reload)
- Featured countries listing
- Popular attractions section
- Navigation menu

### Country Explorer (`/countries`)
- Server-rendered default country list on page load
- Search box and region filter
- AJAX-powered search — only `#country-results` updates, no page reload
- Each country card links to its SSR detail page

###  Destination Details (`/countries/:slug`)
- Full SSR page with country info: flag, capital, population, currency, languages
- Tourist attractions, museums, and landmarks via OpenTripMap
- Current weather and travel conditions via WeatherAPI *(bonus)*
- AJAX "Add to Wishlist" — only `#wishlist-feedback` updates

###  Travel Wishlist (`/wishlist`) : Protected
- Add, edit, and remove destinations
- Edit notes per destination
- Mark destinations as `Planned` or `Visited`

###  Dashboard (`/dashboard`) : Protected


---



##  Wishlist Storage : In-Memory

TravelSphere stores wishlist data entirely **in application memory** using a Go map. No database or external storage is used for this assessment.



### Important: Data is temporary

| Behaviour | Detail |
|---|---|
|  Speed | Extremely fast — all reads and writes happen in RAM, no disk I/O |
| Persistence | **None** — all wishlist data is lost on every application restart |
|  `bee run` reload | Hot-reload triggered by file changes will **reset** the wishlist |
| Terminal close | Closing the terminal or stopping the process **clears** all entries |

This is intentional per the assessment requirements (no database). For production use, replace `wishlistStore` with a persistent store (database, Redis, or a JSON file).

---

##  Getting Started

### Prerequisites

Make sure you have the following installed:

```bash
# Go 1.21+
go version

# Bee tool for live reload
go install github.com/beego/bee/v2@latest

# Verify bee is on your PATH
bee version
```

### 1. Clone the repository

```bash
git clone https://github.com/Parisa-Reza/travelSphere
cd travelSphere
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set up environment variables

```bash
cp .env.example .env
```

```bash
cp conf/app.conf.example conf/app.conf
```

---





##  Running the App

### Development (live reload)

```bash
bee run
```

The app will start at: [http://localhost:8080](http://localhost:8080)

Bee watches for file changes and automatically recompiles and restarts the server. Note that a restart **clears** the in-memory wishlist.

### Run with Docker

```bash
docker-compose up --build
```
the app will start at: [http://localhost:8080](http://localhost:8080)

---

##  Running Tests

```bash
# Run all tests
go test ./...

# Run a specific package
go test ./services/... -v

# total code coverage
go test -coverprofile=total_coverage.out ./... && go tool cover -func=total_coverage.out
```


---

##  API Routes Reference

### SSR Page Routes (return `text/html`)

| Method | Route | Template | Description |
|---|---|---|---|
| `GET` | `/` | `home.tpl` | Home page |
| `GET` | `/countries` | `countries.tpl` | Country Explorer |
| `GET` | `/countries/:slug` | `destination.tpl` | Country detail page |
| `GET` | `/wishlist` | `wishlist.tpl` | Wishlist page *(protected)* |
| `GET` | `/dashboard` | `dashboard.tpl` | Dashboard *(protected)* |

### JSON API Routes (return `application/json`)

| Method | Route | Description |
|---|---|---|
| `GET` | `/api/countries` | Country list; supports `?search=` and `?region=` |
| `GET` | `/api/countries/:slug` | Single country detail |
| `GET` | `/api/attractions` | Attractions by country/coordinates |
| `GET` | `/api/wishlist` | Get all wishlist entries |
| `POST` | `/api/wishlist` | Create a wishlist entry |
| `PUT` | `/api/wishlist/:id` | Update note or status |
| `DELETE` | `/api/wishlist/:id` | Delete a wishlist entry |
counts |



---

##  Screenshots

> Screenshots of the running app.

##  Future improvement

- add debounce for searching
- implement dashboard
- increasing code coverage

---

##  License

This project is for learning purpose assigned by W3 Engineers Ltd.
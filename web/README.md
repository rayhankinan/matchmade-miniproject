# Frontend Application

This is a frontend application that provides a simple movie watchlist application.

## Contents

The contents of the README are as follows:

- [Frontend Application](#frontend-application)
  - [Contents](#contents)
  - [Stack](#stack)
  - [File Structure](#file-structure)
  - [Pages](#pages)
    - [Register](#register)
    - [Login](#login)
    - [Discover Movies](#discover-movies)
    - [Search Movies](#search-movies)
    - [Discover Watchlist](#discover-watchlist)
    - [Search Watchlist](#search-watchlist)
    - [Movie Details](#movie-details)

## Stack

The application is built using the following stack:
- React
- Next (App Router)
- TypeScript
- Tailwind CSS
- Shadcn (UI Components)
- Axios
- TanStack React Query
- Docker

## File Structure

The file structure of the service is as follows:

```
Web
├── Dockerfile
├── README.md
├── components.json
├── next-env.d.ts
├── next.config.mjs
├── package.json
├── postcss.config.cjs
├── prettier.config.js
├── public
│   ├── favicon.ico
│   └── images
│       ├── raster
│       │   └── logo.png
│       └── vector
│           └── placeholder.svg
├── src
│   ├── app
│   │   ├── (main)
│   │   │   ├── page.tsx
│   │   │   ├── template.tsx
│   │   │   └── watchlist
│   │   │       └── page.tsx
│   │   ├── (user)
│   │   │   ├── login
│   │   │   │   └── page.tsx
│   │   │   ├── register
│   │   │   │   └── page.tsx
│   │   │   └── template.tsx
│   │   └── layout.tsx
│   ├── client
│   │   ├── api.ts
│   │   └── movie-api.ts
│   ├── components
│   │   ├── app
│   │   │   ├── icon
│   │   │   │   └── spinner.tsx
│   │   │   ├── main
│   │   │   │   ├── home
│   │   │   │   │   ├── discover-page.tsx
│   │   │   │   │   └── search-page.tsx
│   │   │   │   ├── login
│   │   │   │   │   └── index.tsx
│   │   │   │   ├── movie
│   │   │   │   │   ├── add-to-watchlist-button.tsx
│   │   │   │   │   ├── movie-clickable.tsx
│   │   │   │   │   ├── movie-dialog.tsx
│   │   │   │   │   ├── movie-rating.tsx
│   │   │   │   │   └── movie-trailer.tsx
│   │   │   │   ├── register
│   │   │   │   │   └── index.tsx
│   │   │   │   └── watchlist
│   │   │   │       ├── discover-watchlist-page.tsx
│   │   │   │       └── search-watchlist-page.tsx
│   │   │   ├── navbar
│   │   │   │   ├── index.tsx
│   │   │   │   ├── mode-toggle.tsx
│   │   │   │   ├── navigation-menu.tsx
│   │   │   │   ├── profile.tsx
│   │   │   │   └── search-bar.tsx
│   │   │   ├── template
│   │   │   │   ├── AuthTemplate.tsx
│   │   │   │   └── MovieTemplate.tsx
│   │   │   └── warning
│   │   │       ├── login-alert.tsx
│   │   │       └── watchlist-alert.tsx
│   │   └── ui
│   │       ├── alert.tsx
│   │       ├── avatar.tsx
│   │       ├── button.tsx
│   │       ├── card.tsx
│   │       ├── carousel.tsx
│   │       ├── dialog.tsx
│   │       ├── dropdown-menu.tsx
│   │       ├── form.tsx
│   │       ├── input.tsx
│   │       ├── label.tsx
│   │       ├── navigation-menu.tsx
│   │       ├── toast.tsx
│   │       ├── toaster.tsx
│   │       └── use-toast.ts
│   ├── context
│   │   └── auth.ts
│   ├── env.js
│   ├── hooks
│   │   └── auth.ts
│   ├── lib
│   │   ├── get-blur-data.ts
│   │   └── utils.ts
│   ├── providers
│   │   ├── auth.tsx
│   │   ├── client.tsx
│   │   └── theme.tsx
│   ├── server
│   │   └── auth.ts
│   ├── styles
│   │   └── globals.css
│   └── types
│       └── jwt-payload.ts
├── tailwind.config.ts
├── tsconfig.json
└── yarn.lock
```

## Pages

The application provides the following pages:

### Register

### Login

### Discover Movies

### Search Movies

### Discover Watchlist

### Search Watchlist

### Movie Details
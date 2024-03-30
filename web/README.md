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
<img width="1512" alt="Screenshot 2024-03-30 at 21 21 25" src="https://github.com/rayhankinan/matchmade-miniproject/assets/65068642/0f811577-356b-4e57-ac2e-af8fde40f7f2">

### Login
<img width="1512" alt="Screenshot 2024-03-30 at 21 23 07" src="https://github.com/rayhankinan/matchmade-miniproject/assets/65068642/87fef274-182e-421b-b108-e536343ab7d4">

### Discover Movies
<img width="1512" alt="Screenshot 2024-03-30 at 21 23 43" src="https://github.com/rayhankinan/matchmade-miniproject/assets/65068642/f0d9be3f-eade-4086-a579-76dddcc43423">

### Search Movies
<img width="1512" alt="Screenshot 2024-03-30 at 21 24 22" src="https://github.com/rayhankinan/matchmade-miniproject/assets/65068642/45d20da2-79be-448b-8a9f-77e4ce274f3a">

### Movie Details
<img width="1512" alt="Screenshot 2024-03-30 at 21 25 10" src="https://github.com/rayhankinan/matchmade-miniproject/assets/65068642/4c58a84a-7d14-412f-8268-43bae238a22b">

### Discover Watchlist
<img width="1512" alt="Screenshot 2024-03-30 at 21 27 02" src="https://github.com/rayhankinan/matchmade-miniproject/assets/65068642/54d09ba1-befa-489c-ae15-d6cf23b14ba3">

### Search Watchlist
<img width="1512" alt="Screenshot 2024-03-30 at 21 27 51" src="https://github.com/rayhankinan/matchmade-miniproject/assets/65068642/c1c86d41-fa6c-49d7-bec9-8c97cc3e7c02">

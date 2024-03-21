##### DEPENDENCIES

FROM --platform=linux/amd64 node:20-alpine AS deps
RUN apk add --no-cache libc6-compat openssl
WORKDIR /app

# Install dependencies based on the preferred package manager

COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml\* ./

RUN \
    if [ -f yarn.lock ]; then yarn --frozen-lockfile; \
    elif [ -f package-lock.json ]; then npm ci; \
    elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && pnpm i; \
    else echo "Lockfile not found." && exit 1; \
    fi

##### BUILDER

FROM --platform=linux/amd64 node:20-alpine AS builder

# Add server or client variables as an argument - remove if not using environment variables
ARG NEXT_PUBLIC_API_URL
ARG NEXT_PUBLIC_MOVIE_API_URL
ARG NEXT_PUBLIC_MOVIE_API_TOKEN
ARG NEXT_PUBLIC_MOVIE_IMAGE_URL
ARG NEXT_PUBLIC_YOUTUBE_VIDEO_URL
ARG NEXT_PUBLIC_VIMEO_VIDEO_URL
ENV NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_MOVIE_API_URL=$NEXT_PUBLIC_MOVIE_API_URL
ENV NEXT_PUBLIC_MOVIE_API_TOKEN=$NEXT_PUBLIC_MOVIE_API_TOKEN
ENV NEXT_PUBLIC_MOVIE_IMAGE_URL=$NEXT_PUBLIC_MOVIE_IMAGE_URL
ENV NEXT_PUBLIC_YOUTUBE_VIDEO_URL=$NEXT_PUBLIC_YOUTUBE_VIDEO_URL
ENV NEXT_PUBLIC_VIMEO_VIDEO_URL=$NEXT_PUBLIC_VIMEO_VIDEO_URL

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

# ENV NEXT_TELEMETRY_DISABLED 1

RUN \
    if [ -f yarn.lock ]; then SKIP_ENV_VALIDATION=1 yarn build; \
    elif [ -f package-lock.json ]; then SKIP_ENV_VALIDATION=1 npm run build; \
    elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && SKIP_ENV_VALIDATION=1 pnpm run build; \
    else echo "Lockfile not found." && exit 1; \
    fi

##### RUNNER

FROM --platform=linux/amd64 node:20-alpine AS runner
WORKDIR /app

ENV NODE_ENV production

# ENV NEXT_TELEMETRY_DISABLED 1

COPY --from=builder /app/next.config.mjs ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

EXPOSE 3000
ENV PORT 3000

CMD ["server.js"]
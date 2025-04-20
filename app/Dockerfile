# syntax=docker/dockerfile:1.4  # 🔸 Required for RUN --mount
FROM node:20-alpine AS base

# Install pnpm globally
RUN npm install -g pnpm

WORKDIR /app
COPY . .

# ───────────────────────────────────────────────
# Dev Stage – simple, no cache optimization
FROM base AS development
RUN pnpm install
CMD ["pnpm", "dev", "--", "--host", "0.0.0.0"]

# ───────────────────────────────────────────────
# Build Stage – includes cache-friendly install
FROM base AS build
# Enable cache for pnpm store in prod build
RUN --mount=type=cache,id=pnpm,target=/pnpm/store \
  pnpm install --prod --frozen-lockfile
RUN pnpm build

# ───────────────────────────────────────────────
# Production Stage – serve the static site
FROM node:20-alpine AS production
RUN npm install -g serve
WORKDIR /app
COPY --from=build /app/dist .
CMD sh -c "serve -s . -l ${PORT:-5173}"
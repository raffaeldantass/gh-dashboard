# Development stage
FROM node:20-alpine AS dev

WORKDIR /app

# Install dependencies only when needed
COPY package.json package-lock.json* ./
RUN npm install

# Copy the rest of the application code
COPY . .

# Start development server with hot reload
CMD ["npm", "run", "dev"]

# Production builder stage
FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm install

COPY . .

# Build the application
RUN npm run build

# Production stage
FROM node:20-alpine AS prod

WORKDIR /app

# Copy package files
COPY package.json package-lock.json* ./

# Install only production dependencies
RUN npm install --production

# Copy built application from builder
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/package.json ./package.json

# Start production server
CMD ["npm", "start"]
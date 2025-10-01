# Multi-stage build for production
FROM node:18-alpine AS frontend-builder

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci --only=production --silent
COPY frontend/ ./
RUN npm run build

FROM golang:1.23-alpine AS backend-builder

WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata nginx
RUN addgroup -g 1001 -S appgroup && adduser -u 1001 -S appuser -G appgroup

WORKDIR /app/

# Copy backend binary
COPY --from=backend-builder /app/backend/main .

# Copy frontend build to nginx
COPY --from=frontend-builder /app/frontend/build /var/www/html

# Copy nginx config
COPY nginx.conf /etc/nginx/nginx.conf

# Create non-root user for backend
RUN chown -R appuser:appgroup /app/
RUN chown -R nginx:nginx /var/www/html

# Expose both backend and nginx ports
EXPOSE 8080 80

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/v1/health || exit 1

# Start both nginx and backend
CMD ["sh", "-c", "nginx -g 'daemon off;' & ./main"]
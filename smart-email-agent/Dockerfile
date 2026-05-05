FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o agent .

# Node.js is required for the Gmail MCP server (npx @gongrzhe/server-gmail-autoauth-mcp).
FROM node:20-alpine
WORKDIR /app

# Pre-install the MCP server to avoid npx download at runtime.
RUN npm install -g @gongrzhe/server-gmail-autoauth-mcp

COPY --from=builder /app/agent .

EXPOSE 8080

# GMAIL_TOKEN_PATH and GOOGLE_API_KEY are injected at deploy time via Secret Manager.
CMD ["./agent", "web", "api"]

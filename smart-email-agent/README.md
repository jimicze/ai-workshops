# Smart Email Agent — Workshop

Build a Gmail-backed AI agent with [Google ADK](https://google.golang.org/adk) and Gemini that manages your inbox via natural language.

---

## What You'll Build

```
smart_email_agent  (root orchestrator)
├── email_classifier — triage, spam detection, urgency classification
└── email_drafter    — compose and send replies
```

Gmail operations are powered by [`@gongrzhe/server-gmail-autoauth-mcp`](https://github.com/gongrzhe/server-gmail-autoauth-mcp) — an MCP server that handles Gmail OAuth automatically.

---

## Prerequisites

| Tool | Version | Check |
|---|---|---|
| Go | 1.22+ | `go version` |
| Node.js | 18+ | `node --version` |
| npx | bundled with Node | `npx --version` |
| Gmail account | any | — |

---

## Step 1 — Get a Gemini API Key

1. Go to [Google AI Studio](https://aistudio.google.com/apikey)
2. Click **Create API key**
3. Copy the key — you'll need it in Step 3

---

## Step 2 — Clone and Install Dependencies

```bash
git clone <repo-url>
cd smart-email-agent
go mod download
```

Verify the project structure:
```
agent.go       # agent wiring + main()
gmail_mcp.go   # Gmail MCP server subprocess
tools.go       # classify_email rule-based classifier
types.go       # arg/result types
```

---

## Step 3 — Configure Environment

Export your Gemini API key:

```bash
export GOOGLE_API_KEY=your_gemini_api_key_here
```

Or add it to a `.env` file and source it:
```bash
echo "GOOGLE_API_KEY=your_key_here" > .env
source .env
```

> ⚠️ Never commit `.env` — it's in `.gitignore`.

---

## Step 4 — Run the Agent (Console Mode)

```bash
go run . console
```

**First run:** a browser window will open for Gmail OAuth consent. Sign in and grant access. The token is cached automatically — you won't be prompted again.

You should see:
```
> 
```

Try your first query:
```
> list today's emails
```

---

## Step 5 — Try the Agent

Work through these example queries in the console:

### 5a. Read your inbox
```
list my unread emails
```

### 5b. Classify and triage
```
any urgent emails today?
```
```
check for spam in my inbox
```

### 5c. Draft a reply
```
draft a professional reply to the latest email from my boss
```

### 5d. Send an email
```
send an email to alice@example.com with subject "Hello" and body "Just checking in!"
```

### 5e. Search
```
find emails about the Q2 report
```

---

## Step 6 — Run the Web UI (Optional)

For a browser-based chat interface:

```bash
go run . web webui api
```

Open [http://localhost:8080](http://localhost:8080) and interact with the agent through the UI.

---

## Step 7 — Explore the Code

### How Gmail tools reach the agent

`gmail_mcp.go` launches `@gongrzhe/server-gmail-autoauth-mcp` as a subprocess via `npx` and wraps it as an ADK `Toolset`:

```go
cmd := exec.Command("npx", "@gongrzhe/server-gmail-autoauth-mcp")
transport := &mcp.CommandTransport{Command: cmd}
mcptoolset.New(mcptoolset.Config{Transport: transport})
```

The MCP server exposes tools like `list_emails`, `get_email`, `send_email`, `reply_to_email`, `search_emails`, `trash_email`, and more — all available to every agent automatically.

### How classification works

`classify_email` in `tools.go` is a lightweight rule-based classifier. The agent fetches emails from Gmail MCP, then calls this tool with `subject`, `from`, and `body` to get a classification: `urgent`, `important`, `spam`, or `normal`.

### Sub-agent delegation

The root agent delegates to sub-agents based on intent:
- **email_classifier** → triage, spam detection
- **email_drafter** → composing and sending replies

---

## Troubleshooting

| Error | Fix |
|---|---|
| `npx not found` | Install Node.js from [nodejs.org](https://nodejs.org) |
| `Failed to create model` | Check `GOOGLE_API_KEY` is exported |
| Gmail OAuth loop | Delete `~/.gmail-mcp/` and re-run |
| `Error 400 INVALID_ARGUMENT` | Don't mix `geminitool.GoogleSearch` with function-calling tools |

---

---

## Step 8 — Deploy to Cloud Run

### 8a. Generate the Gmail token locally first

Run the agent locally once to complete the OAuth flow:
```bash
go run . console
```
The token is saved to `~/.gmail-mcp/token.json` (or wherever the MCP server caches it).

### 8b. Store secrets in Secret Manager

```bash
PROJECT_ID=your-gcp-project-id

# Gemini API key
echo -n "your_gemini_api_key" | \
  gcloud secrets create GOOGLE_API_KEY \
    --project=$PROJECT_ID --data-file=-

# Gmail OAuth token (pre-generated locally)
gcloud secrets create GMAIL_TOKEN \
  --project=$PROJECT_ID \
  --data-file=$HOME/.gmail-mcp/token.json
```

### 8c. Build and push the container

```bash
REGION=us-central1
IMAGE=gcr.io/$PROJECT_ID/smart-email-agent

gcloud builds submit --tag $IMAGE --project=$PROJECT_ID
```

### 8d. Deploy

```bash
gcloud run deploy smart-email-agent \
  --image $IMAGE \
  --region $REGION \
  --project $PROJECT_ID \
  --platform managed \
  --allow-unauthenticated \
  --set-secrets="GOOGLE_API_KEY=GOOGLE_API_KEY:latest,/secrets/token.json=GMAIL_TOKEN:latest" \
  --set-env-vars="GMAIL_TOKEN_PATH=/secrets/token.json"
```

The `GMAIL_TOKEN_PATH` env var tells `gmail_mcp.go` to pass the pre-generated token file to the MCP server — no browser flow required in Cloud Run.

### 8e. Verify

```bash
SERVICE_URL=$(gcloud run services describe smart-email-agent \
  --region $REGION --format='value(status.url)')

curl -X POST $SERVICE_URL/api/apps/smart_email_agent/sessions \
  -H "Content-Type: application/json" \
  -d '{"userId": "user1"}'
```

---

## Next Steps

- Add a new sub-agent (e.g. `email_summarizer`)
- Set up CI/CD with Cloud Build triggers on push

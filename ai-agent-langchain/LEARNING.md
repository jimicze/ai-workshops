# Key Learnings

## Part 1: Tools & ReAct Agent

### Tools
- Tools are defined with `tool(fn, { name, description, schema })` from `langchain`
- `zod` schemas enforce type safety on tool inputs
- Tools are async — perfect for API calls (e.g., `wttr.in` weather fetch)

### Agent Reasoning
- The agent uses **ReAct** (Reason + Act) to decide which tool to call
- It reads tool `description` to decide which tool fits the query
- No tool call is made when the answer can be reasoned directly

### Model
- `google:gemini-2.5-flash` — fast, capable, low-latency for tool-calling tasks
- Configured via `GOOGLE_API_KEY` in `.env`

### Gotchas
- `"type": "module"` in `package.json` is required for ES module imports
- Use `@dotenvx/dotenvx` instead of `dotenv` for better `.env` loading with `tsx`

---

## Part 2: Memory & Persistence

_(To be filled after completing Part 2)_

### MemorySaver
- 

### Thread IDs
- 

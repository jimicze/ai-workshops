# Key Learnings

## Part 1: Tools & ReAct Agent

### Tools
- Tools are defined with `tool(fn, { name, description, schema })` from `@langchain/core/tools`
- `zod` schemas enforce type safety on tool inputs
- Tools are async — perfect for API calls (e.g., `wttr.in` weather fetch)

### Agent
- Use `createReactAgent` from `@langchain/langgraph/prebuilt` (NOT `createAgent` from `langchain`)
- Model passed as `llm: new ChatGoogleGenerativeAI({ model: "gemini-2.5-flash" })`
- The agent uses **ReAct** (Reason + Act) to decide which tool to call
- It reads tool `description` to decide which tool fits the query

### Model
- `gemini-2.5-flash` via `ChatGoogleGenerativeAI` from `@langchain/google-genai`
- Configured via `GOOGLE_API_KEY` in `.env`, loaded with `@dotenvx/dotenvx`

### Gotchas
- `"type": "module"` in `package.json` is required for ES module imports
- Use `@dotenvx/dotenvx` instead of `dotenv` for reliable `.env` loading with `tsx`
- Import `tool` from `@langchain/core/tools`, NOT from `langchain` directly

---

## Part 2: Memory & Persistence

### MemorySaver vs SqliteSaver
- `MemorySaver` — in-process, lost on restart. Good for testing
- `SqliteSaver` — on-disk, survives restarts. Use `SqliteSaver.fromConnString("./checkpoint.db")`
- **Gotcha**: `SqliteSaver` requires `memory.setup()` before first use to create DB tables — without it, checkpoints silently fail

### Thread IDs
- Every `agent.invoke()` call must pass `{ configurable: { thread_id: "..." } }` as the second argument
- Same `thread_id` = same conversation history loaded from the checkpointer
- Different `thread_id` = isolated context (separate "save file")

### createReactAgent vs createAgent
- `createAgent` from `langchain` does NOT support `checkpointer` — memory won't work
- Use `createReactAgent` from `@langchain/langgraph/prebuilt` with `llm:` and `checkpointer:` options
- Model must be instantiated explicitly: `new ChatGoogleGenerativeAI({ model: "gemini-2.5-flash" })`

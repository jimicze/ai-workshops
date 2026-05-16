# AI Agent Workshop — LangChain + Google Gemini

A hands-on workshop for building AI agents with **LangChain**, **LangGraph**, and **Google Gemini**.

## Stack
| Package | Purpose |
|---|---|
| `langchain` | Agent creation & tool binding |
| `@langchain/google-genai` | Google Gemini model integration |
| `@langchain/langgraph` | Stateful agent graphs & memory |
| `zod` | Tool schema validation |
| `@dotenvx/dotenvx` | Environment variable loading |

## Project Structure
```
ai-agent-langchain/
├── src/
│   ├── part1.ts        # Part 1 — tools + basic ReAct agent
│   └── part2.ts        # Part 2 — memory, threads, interactive REPL
├── .env                # API keys (not committed)
├── package.json
├── tsconfig.json
├── WORKSHOP_PART1.md   # Part 1 step-by-step instructions
├── WORKSHOP_PART2.md   # Part 2 step-by-step instructions
├── PROGRESS.md         # Workshop progress tracker
└── LEARNING.md         # Key learnings & gotchas
```

## Setup
```bash
npm install
# Add your Google AI Studio key to .env:
# GOOGLE_API_KEY=your_key_here
```

## Running

| Command | What it runs |
|---|---|
| `npm run part1` | Part 1 — `squareTool` + `weatherTool` + basic ReAct agent |
| `npm run part2` | Part 2 — persistent memory, multi-thread, interactive REPL |
| `npm start` | Alias for `npm run part2` |

## Workshop Parts

### Part 1 — Tools & Basic Agent ([src/part1.ts](src/part1.ts))
- `squareTool` — calculates the square of a number
- `weatherTool` — fetches real-time weather from `wttr.in`
- ReAct agent with `createReactAgent` + Gemini 2.5 Flash
- No memory — each run starts fresh

### Part 2 — Memory & Persistence ([src/part2.ts](src/part2.ts))
- `SqliteSaver` checkpointer — on-disk memory that survives restarts
- `thread_id` conversation threading — isolated context per user/session
- Multi-thread demo — `user-123` (Aditya) vs `user-456` (blank slate)
- Interactive REPL — type messages live, type `stop` or `exit` to quit

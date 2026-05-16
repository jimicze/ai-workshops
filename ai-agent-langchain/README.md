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
│   └── index.ts        # Main agent entry point
├── .env                # API keys (not committed)
├── package.json
├── tsconfig.json
├── WORKSHOP_PART1.md   # Tools & basic agent
├── WORKSHOP_PART2.md   # Memory & persistence
├── PROGRESS.md         # Workshop progress tracker
└── LEARNING.md         # Key learnings & notes
```

## Setup
```bash
npm install
# Add your Google AI Studio key to .env:
# GOOGLE_API_KEY=your_key_here
npm start
```

## Workshop Parts
- **Part 1** — Custom tools (`square`, `get_weather`) + ReAct agent
- **Part 2** — Persistent memory with `MemorySaver` + `thread_id` conversations

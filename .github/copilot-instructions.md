# GitHub Copilot Instructions

## Project Context
This is a GDG workshop project building AI agents with **LangChain + Google Gemini** in TypeScript.

## Tech Stack
- **Runtime**: Node.js v20+, TypeScript (ESM `"type": "module"`)
- **AI**: `langchain`, `@langchain/google-genai`, `@langchain/langgraph`
- **Validation**: `zod` v4
- **Dev**: `tsx` for running TypeScript directly, `@dotenvx/dotenvx` for env vars
- **Model**: `google:gemini-2.5-flash`

## Conventions
- All source code lives in `ai-agent-langchain/src/`
- Entry point is `src/index.ts`, run with `npm start` (`tsx src/index.ts`)
- Tools are defined using `tool()` from `langchain` with `zod` schemas
- Agents are created with `createAgent()` from `langchain`
- Environment variables loaded via `@dotenvx/dotenvx` (not plain `dotenv`)
- No `require()` — use `import` only

## Current State
- Part 1 complete: `squareTool` + `weatherTool` + basic ReAct agent working
- Part 2 in progress: adding `MemorySaver` persistence and `thread_id` conversation threading

## Key Files
- `src/part1.ts` — Part 1: tools + basic ReAct agent (no memory)
- `src/part2.ts` — Part 2: SqliteSaver memory, multi-thread, interactive REPL
- `.env` — `GOOGLE_API_KEY` (not committed)
- `WORKSHOP_PART1.md` / `WORKSHOP_PART2.md` — step-by-step instructions
- `PROGRESS.md` — task checklist
- `LEARNING.md` — notes and gotchas

## Run Commands
- `npm run part1` — run Part 1
- `npm run part2` / `npm start` — run Part 2

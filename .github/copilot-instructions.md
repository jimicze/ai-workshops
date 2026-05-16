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
- `src/index.ts` — main agent code
- `.env` — `GOOGLE_API_KEY` (not committed)
- `WORKSHOP_PART1.md` / `WORKSHOP_PART2.md` — step-by-step instructions
- `PROGRESS.md` — task checklist
- `LEARNING.md` — notes and gotchas

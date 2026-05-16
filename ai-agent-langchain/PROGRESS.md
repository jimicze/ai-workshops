# Workshop Progress

## Part 1: Tools & Basic Agent ✅
- [x] Project initialized (TypeScript + ES modules)
- [x] Dependencies installed
- [x] `.env` configured with `GOOGLE_API_KEY`
- [x] `squareTool` — calculates the square of a number
- [x] `weatherTool` — fetches real-time weather via `wttr.in`
- [x] ReAct agent with `createReactAgent` + `gemini-2.5-flash`
- [x] Agent runs and correctly selects tools
- [x] Saved as `src/part1.ts`, runnable via `npm run part1`

## Part 2: Memory & Persistence ✅
- [x] Switch from `createAgent` to `createReactAgent` from `@langchain/langgraph/prebuilt`
- [x] Initialize `SqliteSaver` checkpointer (on-disk, survives restarts)
- [x] Call `memory.setup()` before first use
- [x] Attach checkpointer to agent
- [x] Implement `thread_id` conversation threading
- [x] Multi-thread demo: `user-123` (Aditya) vs `user-456` (isolated, blank slate)
- [x] Interactive REPL with `stop`/`exit` commands
- [x] Test multi-turn memory (name + city recalled across turns)
- [x] Test persistence across process restarts (`checkpoint.db`)
- [x] Saved as `src/part2.ts`, runnable via `npm run part2` / `npm start`

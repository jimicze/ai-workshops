# Workshop Progress

## Part 1: Tools & Basic Agent ✅
- [x] Project initialized (TypeScript + ES modules)
- [x] Dependencies installed
- [x] `.env` configured with `GOOGLE_API_KEY`
- [x] `squareTool` — calculates the square of a number
- [x] `weatherTool` — fetches real-time weather via `wttr.in`
- [x] ReAct agent created with `createAgent` + `gemini-2.5-flash`
- [x] Agent runs and correctly selects tools

## Part 2: Memory & Persistence ✅
- [x] Switch from `createAgent` to `createReactAgent` from `@langchain/langgraph/prebuilt`
- [x] Initialize `SqliteSaver` checkpointer (on-disk, survives restarts)
- [x] Call `memory.setup()` before first use
- [x] Attach checkpointer to agent
- [x] Implement `thread_id` conversation threading
- [x] Test multi-turn memory (name recall across turns)
- [x] Test persistence across process restarts (`checkpoint.db`)

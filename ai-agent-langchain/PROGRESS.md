# Workshop Progress

## Part 1: Tools & Basic Agent ✅
- [x] Project initialized (TypeScript + ES modules)
- [x] Dependencies installed
- [x] `.env` configured with `GOOGLE_API_KEY`
- [x] `squareTool` — calculates the square of a number
- [x] `weatherTool` — fetches real-time weather via `wttr.in`
- [x] ReAct agent created with `createAgent` + `gemini-2.5-flash`
- [x] Agent runs and correctly selects tools

## Part 2: Memory & Persistence 🔲
- [ ] Import `MemorySaver` from `@langchain/langgraph`
- [ ] Initialize checkpointer and attach to agent
- [ ] Implement `thread_id` conversation threading
- [ ] Test multi-turn memory (name recall across turns)
- [ ] Test independent threads (isolated conversation contexts)

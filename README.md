# GDG Workshop — AI Agents

A collection of hands-on workshops for building AI agents with Google Gemini.

## Workshops

### 1. [ai-agent-langchain](./ai-agent-langchain/)
**TypeScript · LangChain · LangGraph**

Build a ReAct agent with custom tools and persistent memory.

| Part | Topic | Status |
|---|---|---|
| Part 1 | Custom tools (`square`, `get_weather`) + ReAct agent | ✅ Done |
| Part 2 | Persistent memory with `MemorySaver` + `thread_id` | 🔲 In progress |

**Stack**: Node.js, TypeScript, `langchain`, `@langchain/google-genai`, `@langchain/langgraph`, `zod`

---

### 2. [smart-email-agent](./smart-email-agent/)
**Go · Google ADK · Gmail MCP**

Build a Gmail-backed multi-agent system that manages your inbox via natural language.

| Agent | Role |
|---|---|
| `email_classifier` | Triage, spam detection, urgency classification |
| `email_drafter` | Compose and send replies |

**Stack**: Go 1.22+, Google ADK, Gemini, Gmail MCP server

---

## Prerequisites
- Node.js v20+ (for `ai-agent-langchain`)
- Go 1.22+ (for `smart-email-agent`)
- Google AI Studio API Key — [aistudio.google.com](https://aistudio.google.com/)

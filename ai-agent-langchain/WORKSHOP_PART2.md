# Workshop Part 2: Memory and Persistent Conversations

In Part 1, we built an agent that could use tools. However, it had "goldfish memory"—it forgot everything the moment the execution finished. In Part 2, we will implement **Persistence** using **LangGraph Checkpoints**.

## 🎯 Objectives
- Understand how Agent Memory works.
- Implement a `MemorySaver` (Checkpointer).
- Use `thread_id` to manage multiple independent conversations.
- Build a persistent assistant that remembers user context across multiple runs.

---

## 🚀 Step 1: Why Memory Matters?

Try asking your Part 1 agent: *"My name is Aditya"* and then *"What is my name?"*. It won't know. 

In LangChain, memory is managed by a **Checkpointer**. This saves the "state" of the agent's brain (the message history) so it can be resumed later.

---

## ⚙️ Step 2: Implementation

We need to import `MemorySaver` and update our agent configuration.

Update `src/index.ts`:

```typescript
import * as dotenv from "dotenv";
import { tool, createAgent } from "langchain";
import { MemorySaver } from "@langchain/langgraph"; // New import
import { z } from "zod";

dotenv.config();

// ... (Keep your squareTool and weatherTool from Part 1)

// 1. Initialize Memory
const memory = new MemorySaver();

// 2. Create the Agent with the Checkpointer
const agent = createAgent({
  model: "google:gemini-2.5-flash",
  tools: [squareTool, weatherTool],
  checkpoint: memory, // Enable persistence
});
```

---

## 🧵 Step 3: Managing Threads

To use memory, every interaction must have a `thread_id`. This is how the agent knows which "save file" to load.

Update your `main` function:

```typescript
async function main() {
  const config = { configurable: { thread_id: "user-123" } };

  // Turn 1: Introduce ourselves
  console.log("--- Turn 1 ---");
  await agent.invoke(
    { messages: [{ role: "user", content: "Hi! My name is Aditya." }] },
    config
  );

  // Turn 2: Ask a follow-up question
  console.log("\n--- Turn 2 ---");
  const response = await agent.invoke(
    { messages: [{ role: "user", content: "What is my name?" }] },
    config
  );

  console.log("Agent:", response.messages[response.messages.length - 1].content);
}
```

---

## 🧪 Step 4: Testing Persistence

When you run `npm start`, you will see:
1. The agent acknowledges your name.
2. In the second turn, without you repeating your name, it correctly identifies you as "Aditya" and calculates the square of 9.

---

## 🛠 Step 5: Advanced Challenge - Contextual Reasoning

Now that the agent has memory, try this:
1. Ask for the weather in **Tokyo**.
2. Then ask: *"Is the temperature there higher than the square of 4?"*

The agent now has to:
- Retrieve the weather from the **previous turn's state**.
- Calculate the square of 4 using the **tool**.
- Compare the two values and give you an answer.

---

## 💾 Step 6: Persistent Memory (On-Disk)

The `MemorySaver` we used in Step 2 is **ephemeral**—it loses all data as soon as your Node.js process stops. For real applications, you need **persistent memory** that survives restarts.

### 1. Install SQLite Checkpointer
LangChain provides an official SQLite-based checkpointer for local on-disk storage.

```bash
npm install @langchain/langgraph-checkpoint-sqlite
```

### 2. Update your Agent to use SQLite
Instead of `MemorySaver`, we use `SqliteSaver` and point it to a database file (e.g., `checkpoint.db`).

```typescript
import { SqliteSaver } from "@langchain/langgraph-checkpoint-sqlite";

// Initialize SQLite checkpointer
const memory = SqliteSaver.fromConnString("./checkpoint.db");

// The rest of the agent setup remains exactly the same!
const agent = createAgent({
  model: "google:gemini-2.5-flash",
  tools: [squareTool, weatherTool],
  checkpoint: memory, 
});
```

### 3. Verify it works
1. Run your agent once and tell it your name.
2. Stop the process (Ctrl+C).
3. Start it again and ask: *"What is my name?"*. 
4. The agent will **still remember you** because the history is now safely stored in your local `checkpoint.db` file!

---

## 🌟 Summary of Part 2
- **Persistence**: We moved from "goldfish memory" to long-term storage.
- **Threads**: We used `thread_id` to separate save files.
- **On-Disk Storage**: We learned how to use SQLite to make our agent's memory permanent.

## References
- Checkpointer: https://reference.langchain.com/python/langgraph/checkpoints
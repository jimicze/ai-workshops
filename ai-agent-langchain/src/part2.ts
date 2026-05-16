import * as dotenv from "@dotenvx/dotenvx";
import { tool } from "@langchain/core/tools";
import { createReactAgent } from "@langchain/langgraph/prebuilt";
import { ChatGoogleGenerativeAI } from "@langchain/google-genai";
import { SqliteSaver } from "@langchain/langgraph-checkpoint-sqlite";
import { z } from "zod";
import * as readline from "readline";

dotenv.config();

// --- Model ---
const model = new ChatGoogleGenerativeAI({ model: "gemini-2.5-flash" });

// --- Tool 1: Square a number ---
const squareTool = tool(
  async ({ n }: { n: number }) => {
    return (n * n).toString();
  },
  {
    name: "square",
    description: "Calculates the square of a number",
    schema: z.object({
      n: z.number().describe("The number to square"),
    }),
  }
);

// --- Tool 2: Real-time weather via wttr.in ---
const weatherTool = tool(
  async ({ city }: { city: string }) => {
    const response = await fetch(`https://wttr.in/${city}?format=3`);
    if (!response.ok) return "Could not fetch weather data.";
    return await response.text();
  },
  {
    name: "get_weather",
    description: "Get the real-time weather for a specific city",
    schema: z.object({
      city: z.string().describe("The city name, e.g., 'London'"),
    }),
  }
);

// --- Memory (on-disk SQLite — survives restarts) ---
const memory = SqliteSaver.fromConnString("./checkpoint.db");

// --- Agent ---
const agent = createReactAgent({
  llm: model,
  tools: [squareTool, weatherTool],
  checkpointer: memory,
});

// --- Helper ---
async function ask(content: string, threadId: string, label?: string) {
  if (label) console.log(`\n--- ${label} ---`);
  const res = await agent.invoke(
    { messages: [{ role: "user", content }] },
    { configurable: { thread_id: threadId } }
  );
  const reply = res.messages[res.messages.length - 1].content as string;
  if (label) console.log("Agent:", reply);
  return reply;
}

// --- Thread 1: Aditya ---
async function runThread1() {
  console.log("\n========== THREAD: user-123 (Aditya) ==========");
  await ask("Hi! My name is Aditya.", "user-123", "Turn 1");
  await ask("I live in Prague and I love coffee.", "user-123", "Turn 2");
  await ask("What is the weather like in the city I live in?", "user-123", "Turn 3");
  await ask("What do you know about me so far?", "user-123", "Turn 4");
}

// --- Thread 2: Isolated user — knows nothing about Aditya ---
async function runThread2() {
  console.log("\n========== THREAD: user-456 (isolated) ==========");
  await ask("What is my name?", "user-456", "Turn 1 — new thread, no context");
  await ask("What is the square of 7?", "user-456", "Turn 2");
}

// --- Thread 3: Interactive REPL ---
async function runInteractive() {
  const threadId = `interactive-${Date.now()}`;
  console.log(`\n========== INTERACTIVE CHAT (thread: ${threadId}) ==========`);
  console.log('Type your message and press Enter. Type "exit" to quit.\n');

  const rl = readline.createInterface({ input: process.stdin, output: process.stdout });

  const prompt = () =>
    new Promise<string>((resolve) => rl.question("You: ", resolve));

  while (true) {
    const input = (await prompt()).trim();
    if (!input || input.toLowerCase() === "exit" || input.toLowerCase() === "stop") {
      console.log("Goodbye!");
      rl.close();
      break;
    }
    const reply = await ask(input, threadId);
    console.log(`Agent: ${reply}\n`);
  }
}

// --- Main ---
async function main() {
  memory.setup();

  await runThread1();
  await runThread2();
  await runInteractive();
}

main().catch(console.error);


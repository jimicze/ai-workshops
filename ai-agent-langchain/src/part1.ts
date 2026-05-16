import * as dotenv from "@dotenvx/dotenvx";
import { tool } from "@langchain/core/tools";
import { createReactAgent } from "@langchain/langgraph/prebuilt";
import { ChatGoogleGenerativeAI } from "@langchain/google-genai";
import { z } from "zod";

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

// --- Agent (no memory) ---
const agent = createReactAgent({
  llm: model,
  tools: [squareTool, weatherTool],
});

// --- Main ---
async function main() {
  console.log("--- Querying Square Tool ---");
  const res1 = await agent.invoke({
    messages: [{ role: "user", content: "What is the square of 15?" }],
  });
  console.log("Agent:", res1.messages[res1.messages.length - 1].content);

  console.log("\n--- Querying Real-time Weather ---");
  const res2 = await agent.invoke({
    messages: [{ role: "user", content: "What's the weather like in Paris right now?" }],
  });
  console.log("Agent:", res2.messages[res2.messages.length - 1].content);
}

main().catch(console.error);

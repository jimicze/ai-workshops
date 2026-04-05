# Workshop: Building AI Agents with LangChain and Google Gemini

In this workshop, you'll learn how to build a functional AI agent that can use custom tools to solve problems. We'll be using **LangChain** and **Google Gemini**.

## 🎯 Objectives
- Initialize a TypeScript project for AI development.
- Configure Google Gemini with LangChain.
- Define custom tools using Zod for type safety.
- Create and execute a ReAct agent.
- Integrate real-time external data (Weather API).

---

## 🛠 Prerequisites
- **Node.js** (v20+)
- **Google AI Studio API Key**: Get it at [aistudio.google.com](https://aistudio.google.com/)
- Basic knowledge of TypeScript/JavaScript.

---

## 🚀 Step 1: Project Initialization

First, create a new directory and initialize your Node.js project.

```bash
mkdir ai-agent-workshop
cd ai-agent-workshop
npm init -y
```

Install the required dependencies:

```bash
npm install langchain @langchain/google @langchain/core zod dotenv
npm install --save-dev typescript tsx @types/node
```

Initialize TypeScript:

```bash
npx tsc --init
```

---

## ⚙️ Step 2: Configuration

### 1. Project Type
Update your `package.json` to use ES modules by adding `"type": "module"`.

```json
{
  "type": "module",
  "scripts": {
    "start": "tsx src/index.ts"
  }
}
```

### 2. Environment Variables
Create a `.env` file in the root directory:

```env
GOOGLE_API_KEY=your_api_key_here
```

---

## 🛠 Step 3: Defining Local Tools

Tools are the "hands" of your agent. We use the `tool` function to define what the agent can do.

Create `src/index.ts` and add the following:

```typescript
import * as dotenv from "dotenv";
import { tool, createAgent } from "langchain";
import { z } from "zod";

dotenv.config();

// Define a local tool to calculate the square of a number
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
```

---

## 🌍 Step 4: Real-time Weather Integration

Now, let's add a tool that fetches real-time data from an external API. We'll use `wttr.in`, a zero-setup weather service.

Add this to `src/index.ts`:

```typescript
const weatherTool = tool(
  async ({ city }: { city: string }) => {
    // Fetching real-time weather from wttr.in (returns text format)
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
```

---

## 🤖 Step 5: Building the Multi-Tool Agent

Initialize the agent and provide it with **both** the `squareTool` and the `weatherTool`.

```typescript
// Create the agent with multiple tools
const agent = createAgent({
  model: "google:gemini-2.5-flash", 
  tools: [squareTool, weatherTool],
});
```

---

## 🏃 Step 6: Executing and Testing

Update the `main` function to ask a complex question that requires the agent to reason about which tool to use.

```typescript
async function main() {
  // Scenario 1: Using the square tool
  console.log("--- Querying Square Tool ---");
  const res1 = await agent.invoke({
    messages: [{ role: "user", content: "What is the square of 15?" }],
  });
  console.log("Agent:", res1.messages[res1.messages.length - 1].content);

  // Scenario 2: Using the real-time weather tool
  console.log("\n--- Querying Real-time Weather ---");
  const res2 = await agent.invoke({
    messages: [{ role: "user", content: "What's the weather like in Paris right now?" }],
  });
  console.log("Agent:", res2.messages[res2.messages.length - 1].content);
}

main().catch(console.error);
```

---

## 🧪 Step 7: Running the Workshop

Run your agent with:

```bash
npm start
```

### What to Observe:
1. **Tool Selection**: Notice how the agent correctly chooses `square` for the first query and `get_weather` for the second.
2. **API Interaction**: The agent handles the asynchronous network request to `wttr.in` seamlessly.
3. **Reasoning**: Try asking: *"If the temperature in Paris is the square of 4, what is the weather?"* and watch the agent chain both tools!

---

## 🌟 Extra Credit: Production Weather
In a production app, you would use an API like **OpenWeatherMap**. You would:
1. Add `WEATHER_API_KEY` to your `.env`.
2. Update the `weatherTool` to fetch from `api.openweathermap.org` and parse the JSON response.
3. Handle more complex schemas (e.g., returning humidity, wind speed, and forecast).

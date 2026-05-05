package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx, "gemini-flash-latest", &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	// Gmail MCP toolset — required.
	gmailToolset, err := NewGmailMCPToolset()
	if err != nil {
		log.Fatalf("Gmail MCP toolset: %v", err)
	}
	gmailToolsets := []tool.Toolset{gmailToolset}

	// classify_email — rule-based classification helper on top of Gmail MCP data.
	classifyEmailTool, err := functiontool.New(functiontool.Config{
		Name:        "classify_email",
		Description: "Classify an email as 'urgent', 'important', 'spam', or 'normal' given its subject, from, and body.",
	}, handleClassifyEmail)
	if err != nil {
		log.Fatalf("Failed to create classify_email tool: %v", err)
	}

	// Email Classifier sub-agent.
	classifierAgent, err := llmagent.New(llmagent.Config{
		Name:        "email_classifier",
		Model:       model,
		Description: "Classifies emails as urgent, important, spam, or normal. Use for inbox triage and spam detection.",
		Instruction: `You are an email classification specialist. Classify emails into:
- urgent: time-sensitive, requires immediate action
- important: significant but not time-critical
- spam: unsolicited, fraudulent, or phishing
- normal: regular email

Steps:
1. Use list_emails or search_emails (Gmail MCP) to fetch messages
2. For each email, call classify_email with its subject, from, and body
3. Present findings with red flags or deadlines highlighted`,
		Tools:    []tool.Tool{classifyEmailTool},
		Toolsets: gmailToolsets,
	})
	if err != nil {
		log.Fatalf("Failed to create classifier agent: %v", err)
	}

	// Email Drafter sub-agent.
	drafterAgent, err := llmagent.New(llmagent.Config{
		Name:        "email_drafter",
		Model:       model,
		Description: "Drafts and sends email replies. Use when the user wants to compose or reply to an email.",
		Instruction: `You are an expert email composer. When drafting replies:
1. Fetch the original email with get_email (Gmail MCP)
2. Compose a complete reply matching the requested tone:
   - professional: courteous, structured
   - friendly: warm, personable
   - formal: highly structured, for executives
   - concise: brief, no pleasantries
3. Present the draft (To, Subject, Body) and ask for confirmation
4. Send with reply_to_email or send_email (Gmail MCP) on confirmation`,
		Toolsets: gmailToolsets,
	})
	if err != nil {
		log.Fatalf("Failed to create drafter agent: %v", err)
	}

	// Root orchestrator agent.
	rootAgent, err := llmagent.New(llmagent.Config{
		Name:        "smart_email_agent",
		Model:       model,
		Description: "Smart email assistant for Gmail inbox management.",
		Instruction: `You are a smart email assistant backed by the Gmail API. Delegate to sub-agents:
- email_classifier: inbox triage, spam detection, urgency classification
- email_drafter: composing and sending replies

Handle directly: reading emails, searching, inbox summaries, label management.

Gmail MCP tools available: list_emails, get_email, send_email, reply_to_email,
search_emails, create_draft, list_labels, modify_labels, trash_email, get_attachment.

Always warn users about suspicious or phishing emails.`,
		Tools:     []tool.Tool{classifyEmailTool},
		Toolsets:  gmailToolsets,
		SubAgents: []agent.Agent{classifierAgent, drafterAgent},
	})
	if err != nil {
		log.Fatalf("Failed to create root agent: %v", err)
	}

	l := full.NewLauncher()
	if err = l.Execute(ctx, &launcher.Config{
		AgentLoader: agent.NewSingleLoader(rootAgent),
	}, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}

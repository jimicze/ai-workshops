package main

import (
	"fmt"
	"strings"

	"google.golang.org/adk/tool"
)

// handleClassifyEmail applies rule-based classification to an email fetched via Gmail MCP.
// The agent is expected to pass the subject, from, and body after reading the message.
func handleClassifyEmail(_ tool.Context, args ClassifyEmailArgs) (ClassifyEmailResult, error) {
	subject := strings.ToLower(args.Subject)
	body := strings.ToLower(args.Body)
	from := strings.ToLower(args.From)

	spamIndicators := []string{"won", "prize", "bank details", "click here", "act now", "compromised", "verify your identity", "processing fee"}
	spamCount := 0
	for _, ind := range spamIndicators {
		if strings.Contains(subject, ind) || strings.Contains(body, ind) {
			spamCount++
		}
	}
	if spamCount >= 2 || strings.Contains(from, "scam") || strings.Contains(from, "totallylegit") {
		return ClassifyEmailResult{
			Classification: "spam",
			Confidence:     "high",
			Reason:         fmt.Sprintf("Multiple spam indicators detected (%d matches)", spamCount),
		}, nil
	}

	for _, ind := range []string{"urgent", "asap", "immediately", "critical", "due today", "deadline"} {
		if strings.Contains(subject, ind) || strings.Contains(body, ind) {
			return ClassifyEmailResult{
				Classification: "urgent",
				Confidence:     "high",
				Reason:         fmt.Sprintf("Urgency indicator: '%s'", ind),
			}, nil
		}
	}

	for _, ind := range []string{"important", "announcement", "restructuring", "all-hands", "company-wide"} {
		if strings.Contains(subject, ind) || strings.Contains(body, ind) {
			return ClassifyEmailResult{
				Classification: "important",
				Confidence:     "high",
				Reason:         fmt.Sprintf("Importance indicator: '%s'", ind),
			}, nil
		}
	}

	return ClassifyEmailResult{
		Classification: "normal",
		Confidence:     "medium",
		Reason:         "No urgency, importance, or spam signals detected",
	}, nil
}

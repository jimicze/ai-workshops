package main

// ClassifyEmailArgs and ClassifyEmailResult are used by the classify_email tool,
// which applies rule-based classification on top of emails fetched via Gmail MCP.

type ClassifyEmailArgs struct {
	Subject string `json:"subject"`
	From    string `json:"from"`
	Body    string `json:"body"`
}

type ClassifyEmailResult struct {
	Classification string `json:"classification"`
	Confidence     string `json:"confidence"`
	Reason         string `json:"reason"`
}

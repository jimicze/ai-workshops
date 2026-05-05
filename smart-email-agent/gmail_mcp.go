package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/mcptoolset"
)

// NewGmailMCPToolset creates an ADK Toolset backed by @gongrzhe/server-gmail-autoauth-mcp.
//
// Auth strategy:
//   - Local dev: the server opens a browser for OAuth on first run and caches
//     the token automatically (default behaviour, no config needed).
//   - Cloud Run / CI: set GMAIL_TOKEN_PATH to a pre-generated token file
//     (e.g. mounted from GCP Secret Manager). The server reads the token
//     directly and skips the browser flow.
//
// npx must be available in PATH.
func NewGmailMCPToolset() (tool.Toolset, error) {
	if _, err := exec.LookPath("npx"); err != nil {
		return nil, fmt.Errorf("npx not found in PATH; install Node.js: %w", err)
	}

	cmd := exec.Command("npx", "@gongrzhe/server-gmail-autoauth-mcp")
	env := append(os.Environ(), "npm_config_loglevel=error")

	// When running in Cloud Run, GMAIL_TOKEN_PATH points to a pre-generated
	// token file mounted from Secret Manager (no browser flow required).
	if tokenPath := os.Getenv("GMAIL_TOKEN_PATH"); tokenPath != "" {
		if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
			return nil, fmt.Errorf("GMAIL_TOKEN_PATH set but file not found: %s", tokenPath)
		}
		env = append(env, "GMAIL_TOKEN_PATH="+tokenPath)
	}

	cmd.Env = env

	return mcptoolset.New(mcptoolset.Config{
		Transport: &mcp.CommandTransport{Command: cmd},
	})
}

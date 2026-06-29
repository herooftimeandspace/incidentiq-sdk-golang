//go:build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type badgePayload struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
}

var statusMessages = map[string]string{
	"success":         "passing",
	"failure":         "failing",
	"cancelled":       "cancelled",
	"skipped":         "skipped",
	"neutral":         "neutral",
	"timed_out":       "timed out",
	"action_required": "action required",
	"startup_failure": "startup failure",
	"stale":           "stale",
}

var statusColors = map[string]string{
	"success":         "brightgreen",
	"failure":         "red",
	"cancelled":       "orange",
	"skipped":         "lightgrey",
	"neutral":         "blue",
	"timed_out":       "orange",
	"action_required": "orange",
	"startup_failure": "orange",
	"stale":           "yellow",
}

func main() {
	if len(os.Args) < 2 {
		fatal("usage: go run scripts/build_badge_json.go status --conclusion <value> --label <label> --output <path>")
	}

	switch os.Args[1] {
	case "status":
		writeStatusBadge(os.Args[2:])
	default:
		fatal("unsupported badge command %q", os.Args[1])
	}
}

func writeStatusBadge(args []string) {
	flags := flag.NewFlagSet("status", flag.ExitOnError)
	conclusion := flags.String("conclusion", "", "GitHub Actions conclusion")
	label := flags.String("label", "", "badge label")
	output := flags.String("output", "", "output JSON path")
	if err := flags.Parse(args); err != nil {
		fatal("parse status flags: %v", err)
	}
	if *label == "" || *output == "" {
		fatal("--label and --output are required")
	}

	normalized := strings.ToLower(strings.TrimSpace(*conclusion))
	if normalized == "" {
		normalized = "neutral"
	}
	message, ok := statusMessages[normalized]
	if !ok {
		message = strings.ReplaceAll(normalized, "_", " ")
	}
	color, ok := statusColors[normalized]
	if !ok {
		color = "lightgrey"
	}

	payload := badgePayload{
		SchemaVersion: 1,
		Label:         *label,
		Message:       message,
		Color:         color,
	}
	encoded, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		fatal("encode badge payload: %v", err)
	}
	encoded = append(encoded, '\n')
	if err := os.WriteFile(*output, encoded, 0o644); err != nil {
		fatal("write %s: %v", *output, err)
	}
}

func fatal(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

//go:build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
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
		fatal("usage: go run scripts/build_badge_json.go <status|coverage> [flags]")
	}

	switch os.Args[1] {
	case "coverage":
		writeCoverageBadge(os.Args[2:])
	case "status":
		writeStatusBadge(os.Args[2:])
	default:
		fatal("unsupported badge command %q", os.Args[1])
	}
}

func writeCoverageBadge(args []string) {
	flags := flag.NewFlagSet("coverage", flag.ExitOnError)
	coverageFile := flags.String("coverage-file", "", "Go coverage profile path")
	label := flags.String("label", "", "badge label")
	minimum := flags.Float64("minimum", 0, "minimum allowed coverage percentage")
	minimumBadgeFile := flags.String("minimum-badge-file", "", "optional existing coverage badge JSON whose message sets a higher floor")
	output := flags.String("output", "", "output JSON path")
	if err := flags.Parse(args); err != nil {
		fatal("parse coverage flags: %v", err)
	}
	if *coverageFile == "" || *label == "" || *output == "" {
		fatal("--coverage-file, --label, and --output are required")
	}

	percent, err := coveragePercent(*coverageFile)
	if err != nil {
		fatal("calculate coverage: %v", err)
	}
	effectiveMinimum := *minimum
	if *minimumBadgeFile != "" {
		badgeMinimum, err := coverageMinimumFromBadge(*minimumBadgeFile)
		if err != nil {
			fatal("read coverage floor badge: %v", err)
		}
		if badgeMinimum > effectiveMinimum {
			effectiveMinimum = badgeMinimum
		}
	}
	roundedPercent := roundCoveragePercent(percent)
	if roundedPercent < effectiveMinimum {
		fatal("coverage %.2f%% is below required minimum %.2f%%", roundedPercent, effectiveMinimum)
	}

	color := "red"
	switch {
	case percent >= 90:
		color = "brightgreen"
	case percent >= 80:
		color = "green"
	case percent >= 70:
		color = "yellowgreen"
	case percent >= 60:
		color = "yellow"
	case percent >= 50:
		color = "orange"
	}

	writePayload(*output, badgePayload{
		SchemaVersion: 1,
		Label:         *label,
		Message:       fmt.Sprintf("%.2f%%", roundedPercent),
		Color:         color,
	})
}

func coverageMinimumFromBadge(path string) (float64, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}

	var badge badgePayload
	if err := json.Unmarshal(payload, &badge); err != nil {
		return 0, fmt.Errorf("%s: decode badge JSON: %w", path, err)
	}
	message := strings.TrimSpace(strings.TrimSuffix(badge.Message, "%"))
	if message == "" {
		return 0, fmt.Errorf("%s: coverage badge message is empty", path)
	}
	percent, err := strconv.ParseFloat(message, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: parse coverage badge message %q: %w", path, badge.Message, err)
	}
	return percent, nil
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
	writePayload(*output, payload)
}

func coveragePercent(path string) (float64, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	var totalStatements int64
	var coveredStatements int64
	for lineNumber, rawLine := range strings.Split(string(payload), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "mode:") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 3 {
			return 0, fmt.Errorf("%s:%d: expected coverage record with 3 fields", path, lineNumber+1)
		}
		statements, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("%s:%d: parse statement count: %w", path, lineNumber+1, err)
		}
		count, err := strconv.ParseInt(fields[2], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("%s:%d: parse coverage count: %w", path, lineNumber+1, err)
		}
		totalStatements += statements
		if count > 0 {
			coveredStatements += statements
		}
	}

	if totalStatements == 0 {
		return 0, nil
	}
	return float64(coveredStatements) / float64(totalStatements) * 100, nil
}

func roundCoveragePercent(percent float64) float64 {
	return math.Round(percent*100) / 100
}

func writePayload(output string, payload badgePayload) {
	encoded, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		fatal("encode badge payload: %v", err)
	}
	encoded = append(encoded, '\n')
	if err := os.WriteFile(output, encoded, 0o644); err != nil {
		fatal("write %s: %v", output, err)
	}
}

func fatal(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

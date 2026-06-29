//go:build ignore

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var validSemverLabels = []string{"semver:patch", "semver:minor", "semver:major"}
var versionPattern = regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)$`)

type associatedPull struct {
	Number int `json:"number"`
	Base   struct {
		Ref string `json:"ref"`
	} `json:"base"`
	Labels []struct {
		Name string `json:"name"`
	} `json:"labels"`
}

func main() {
	if len(os.Args) < 2 {
		fatal("usage: go run scripts/release_tools.go <bump|resolve-promotion-semver> [flags]")
	}

	switch os.Args[1] {
	case "bump":
		runBump(os.Args[2:])
	case "resolve-promotion-semver":
		runResolvePromotionSemver(os.Args[2:])
	default:
		fatal("unsupported release tool command %q", os.Args[1])
	}
}

func runBump(args []string) {
	flags := flag.NewFlagSet("bump", flag.ExitOnError)
	part := flags.String("part", "", "semver part: patch, minor, or major")
	currentVersion := flags.String("current-version", "", "current semantic version")
	if err := flags.Parse(args); err != nil {
		fatal("parse bump flags: %v", err)
	}
	if *part == "" || *currentVersion == "" {
		fatal("--part and --current-version are required")
	}
	next, err := bumpVersion(*currentVersion, *part)
	if err != nil {
		fatal("%v", err)
	}
	fmt.Println(next)
}

func runResolvePromotionSemver(args []string) {
	flags := flag.NewFlagSet("resolve-promotion-semver", flag.ExitOnError)
	repo := flags.String("repo", "", "owner/repo")
	sha := flags.String("sha", "", "commit SHA")
	baseRef := flags.String("base-ref", "", "preferred associated PR base branch")
	defaultLabel := flags.String("default-label", "semver:patch", "default semver label")
	if err := flags.Parse(args); err != nil {
		fatal("parse resolve-promotion-semver flags: %v", err)
	}
	if *repo == "" || *sha == "" {
		fatal("--repo and --sha are required")
	}

	pulls, err := associatedPulls(*repo, *sha)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Could not resolve pull requests associated with %s; defaulting to %s. Reason: %v\n", *sha, *defaultLabel, err)
		fmt.Println(*defaultLabel)
		return
	}

	prNumber, label, err := resolvePromotionSemverLabel(pulls, *baseRef, *defaultLabel)
	if err != nil {
		fatal("%v", err)
	}
	if prNumber == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "No associated pull request found for %s; defaulting to %s.\n", *sha, label)
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "Resolved %s from PR #%d.\n", label, prNumber)
	}
	fmt.Println(label)
}

func associatedPulls(repo string, sha string) ([]associatedPull, error) {
	cmd := exec.Command("gh", "api", fmt.Sprintf("repos/%s/commits/%s/pulls", repo, sha))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		detail := strings.TrimSpace(stderr.String())
		if detail == "" {
			detail = err.Error()
		}
		return nil, errors.New(detail)
	}
	var pulls []associatedPull
	if err := json.Unmarshal(output, &pulls); err != nil {
		return nil, err
	}
	return pulls, nil
}

func resolvePromotionSemverLabel(pulls []associatedPull, baseRef string, defaultLabel string) (int, string, error) {
	if !slices.Contains(validSemverLabels, defaultLabel) {
		return 0, "", fmt.Errorf("unsupported default semver label %q", defaultLabel)
	}

	var selected *associatedPull
	if baseRef != "" {
		for i := range pulls {
			if pulls[i].Base.Ref == baseRef {
				selected = &pulls[i]
				break
			}
		}
	}
	if selected == nil && len(pulls) > 0 {
		selected = &pulls[0]
	}
	if selected == nil {
		return 0, defaultLabel, nil
	}

	labels := make([]string, 0, len(selected.Labels))
	for _, label := range selected.Labels {
		labels = append(labels, label.Name)
	}
	resolved, err := selectSemverLabel(labels, defaultLabel)
	if err != nil {
		return 0, "", err
	}
	return selected.Number, resolved, nil
}

func selectSemverLabel(labels []string, defaultLabel string) (string, error) {
	seen := make([]string, 0, 1)
	for _, label := range labels {
		if !slices.Contains(validSemverLabels, label) || slices.Contains(seen, label) {
			continue
		}
		seen = append(seen, label)
	}
	switch len(seen) {
	case 0:
		return defaultLabel, nil
	case 1:
		return seen[0], nil
	default:
		return "", fmt.Errorf("expected at most one semver label, found %s", strings.Join(seen, ", "))
	}
}

func bumpVersion(currentVersion string, part string) (string, error) {
	match := versionPattern.FindStringSubmatch(strings.TrimSpace(currentVersion))
	if match == nil {
		return "", fmt.Errorf("unsupported version format %q", currentVersion)
	}
	major, _ := strconv.Atoi(match[1])
	minor, _ := strconv.Atoi(match[2])
	patch, _ := strconv.Atoi(match[3])

	switch part {
	case "major":
		major++
		minor = 0
		patch = 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		return "", fmt.Errorf("unsupported semver part %q", part)
	}
	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}

func fatal(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

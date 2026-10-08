package incidentiq

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// contractAliases are the per-vendor filenames that must resolve to AGENTS.md.
// They are symlinks so they cannot drift from the contract.
var contractAliases = []string{
	"CLAUDE.md",
	"GEMINI.md",
	"CONVENTIONS.md",
	".clinerules",
	".windsurfrules",
}

// contractPointers cannot be symlinks. Cursor requires YAML frontmatter, and
// GitHub's API serves a symlink blob as its path text rather than the target's
// content, so Copilot would read the string "AGENTS.md" as its instructions.
var contractPointers = []string{
	".cursor/rules/agents.mdc",
	".github/copilot-instructions.md",
}

// TestAgentContractAliasesResolve keeps the one-contract-many-aliases setup
// honest. A broken alias is invisible in normal use: the harness that reads it
// silently gets nothing, and the agent proceeds without the repository rules.
func TestAgentContractAliasesResolve(t *testing.T) {
	contract, err := os.ReadFile("AGENTS.md")
	if err != nil {
		t.Fatalf("read AGENTS.md: %v", err)
	}
	if len(contract) == 0 {
		t.Fatal("AGENTS.md is empty")
	}

	for _, alias := range contractAliases {
		info, err := os.Lstat(alias)
		if err != nil {
			t.Errorf("%s is missing; it must be a symlink to AGENTS.md", alias)
			continue
		}
		if info.Mode()&os.ModeSymlink == 0 {
			t.Errorf("%s is a regular file; it must be a symlink to AGENTS.md so it cannot drift", alias)
			continue
		}
		target, err := os.Readlink(alias)
		if err != nil {
			t.Errorf("readlink %s: %v", alias, err)
			continue
		}
		if target != "AGENTS.md" {
			t.Errorf("%s points at %q; it must point at AGENTS.md", alias, target)
			continue
		}
		resolved, err := os.ReadFile(alias)
		if err != nil {
			t.Errorf("%s does not resolve: %v", alias, err)
			continue
		}
		if string(resolved) != string(contract) {
			t.Errorf("%s resolves to content that is not AGENTS.md", alias)
		}
	}

	for _, pointer := range contractPointers {
		info, err := os.Lstat(pointer)
		if err != nil {
			t.Errorf("%s is missing", pointer)
			continue
		}
		if info.Mode()&os.ModeSymlink != 0 {
			t.Errorf("%s is a symlink; it must be a real file, because its consumer does not follow symlinks", pointer)
			continue
		}
		body, err := os.ReadFile(pointer)
		if err != nil {
			t.Errorf("read %s: %v", pointer, err)
			continue
		}
		if !strings.Contains(string(body), "AGENTS.md") {
			t.Errorf("%s does not point at AGENTS.md", pointer)
		}
	}
}

// TestDocsSiteSkipsContractAliases guards the symlink skip in
// scripts/build_docs_site.go. Without it the published site carries three
// byte-identical copies of the contract beside AGENTS.md.
func TestDocsSiteSkipsContractAliases(t *testing.T) {
	source, err := os.ReadFile(filepath.Join("scripts", "build_docs_site.go"))
	if err != nil {
		t.Fatalf("read scripts/build_docs_site.go: %v", err)
	}
	if !strings.Contains(string(source), "os.ModeSymlink") {
		t.Error("scripts/build_docs_site.go no longer skips symlinks; the contract aliases would be published as duplicate pages")
	}
}

package main

import (
	"strings"
	"testing"
)

// Test if system prompt is not empty
func TestSystemPromptNotEmpty(t *testing.T) {
	if systemPrompt == "" {
		t.Error("System prompt should not be empty")
	}
}

// Test if system prompt contains key themes
func TestSystemPromptContainsKeyThemes(t *testing.T) {
	themes := []string{
		"darkness",
		"fairy",
		"philosophical",
		"adventure",
	}
	for _, theme := range themes {
		if !strings.Contains(strings.ToLower(systemPrompt), theme) {
			t.Errorf("System prompt should contain the theme: %s", theme)
		}
	}
}

// Test if system prompt contains rules
func TestSystemPromptContainsRules(t *testing.T) {
	if !strings.Contains(strings.ToLower(systemPrompt), "rules") {
		t.Errorf("System prompt should contain rules")
	}
}

// Test if system prompt encourages moral complexity
func TestSystemPromptEncouragesComplexity(t *testing.T) {
	complexityIndicators := []string{
		"morally complex",
		"no clearly \"right\" answer",
		"consequences",
	}

	foundCount := 0
	for _, indicator := range complexityIndicators {
		if strings.Contains(strings.ToLower(systemPrompt), strings.ToLower(indicator)) {
			foundCount++
		}
	}

	if foundCount == 0 {
		t.Errorf("systemPrompt should contain indicators of moral complexity")
	}
}

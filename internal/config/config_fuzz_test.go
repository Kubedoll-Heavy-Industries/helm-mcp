package config

import (
	"testing"
)

// FuzzParseCSV exercises the comma-separated value parser with arbitrary input.
// This function processes untrusted input from environment variables and CLI
// flags for the allowed-hosts and denied-hosts configuration.
func FuzzParseCSV(f *testing.F) {
	// Seed corpus: realistic CSV inputs
	seeds := []string{
		"",
		"   ",
		"example.com",
		"a.com,b.com,c.com",
		" a.com , b.com , c.com ",
		"a.com,,b.com,",
		",,,",
		"charts.bitnami.com,ghcr.io,registry.k8s.io",
		"*.example.com,.example.com,example.com",
		",leading",
		"trailing,",
		"  ,  ,  ",
		"a,b,c,d,e,f,g,h,i,j",
		"single",
		"has spaces, in values, here",
		"unicode.example.\u00e9,test.com",
	}

	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, input string) {
		result := parseCSV(input)

		// Invariants that must always hold:
		// 1. Result is either nil or a non-empty slice
		if result != nil && len(result) == 0 {
			t.Error("parseCSV returned non-nil empty slice; should be nil")
		}

		// 2. No result element should be empty after trimming
		for i, item := range result {
			if item == "" {
				t.Errorf("parseCSV returned empty string at index %d", i)
			}
		}

		// 3. No result element should have leading/trailing whitespace
		for i, item := range result {
			trimmed := item
			if len(trimmed) > 0 && (trimmed[0] == ' ' || trimmed[0] == '\t' ||
				trimmed[len(trimmed)-1] == ' ' || trimmed[len(trimmed)-1] == '\t') {
				t.Errorf("parseCSV returned item with whitespace at index %d: %q", i, item)
			}
		}
	})
}

// FuzzFlagEnvName exercises the flag-to-environment-variable name conversion
// with arbitrary flag names. This ensures the mapping is deterministic and
// never panics on unusual input.
func FuzzFlagEnvName(f *testing.F) {
	seeds := []string{
		"transport",
		"helm-timeout",
		"cache-size",
		"allow-private-ips",
		"log-level",
		"log-format",
		"",
		"a",
		"a-b-c-d",
		"--double-dash",
		"under_score",
		"MiXeD-CaSe",
	}

	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, flagName string) {
		result := flagEnvName(flagName)

		// Invariant: result always starts with the env prefix
		if len(result) < len(envPrefix) || result[:len(envPrefix)] != envPrefix {
			t.Errorf("flagEnvName(%q) = %q, does not start with %q", flagName, result, envPrefix)
		}

		// Invariant: result should not contain hyphens (they're replaced with underscores)
		for _, c := range result[len(envPrefix):] {
			if c == '-' {
				t.Errorf("flagEnvName(%q) = %q, contains hyphen after prefix", flagName, result)
				break
			}
		}
	})
}

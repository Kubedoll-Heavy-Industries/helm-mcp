package helm

import (
	"testing"
)

// FuzzMatchesHostList exercises the host pattern matching logic with arbitrary
// host strings and pattern lists. This function is used for security-critical
// allow/deny list enforcement, so robustness against malformed input matters.
func FuzzMatchesHostList(f *testing.F) {
	// Seed corpus: realistic hosts and patterns
	type seed struct {
		host    string
		pattern string
	}
	seeds := []seed{
		// Exact matches
		{"example.com", "example.com"},
		{"charts.bitnami.com", "charts.bitnami.com"},
		// Subdomain matches
		{"sub.example.com", "example.com"},
		{"a.b.c.example.com", "example.com"},
		// Wildcard prefix
		{"sub.example.com", ".example.com"},
		{"example.com", ".example.com"},
		// Wildcard all
		{"anything.com", "*"},
		// No match
		{"other.com", "example.com"},
		{"notexample.com", "example.com"},
		// Edge cases
		{"", "example.com"},
		{"example.com", ""},
		{"", ""},
		// Case sensitivity
		{"EXAMPLE.COM", "example.com"},
		{"Example.Com", "EXAMPLE.COM"},
		// Hostnames with special characters
		{"host-with-dashes.com", "host-with-dashes.com"},
		{"123.456.com", "123.456.com"},
		// Long hostnames
		{"a.very.deeply.nested.subdomain.of.example.com", "example.com"},
		// Pattern with whitespace
		{"example.com", "  example.com  "},
		// Patterns with dots
		{"example.com", ".com"},
		{"a.b.com", ".b.com"},
	}

	for _, s := range seeds {
		f.Add(s.host, s.pattern)
	}

	f.Fuzz(func(t *testing.T, host, pattern string) {
		// Test with single-pattern list
		result := matchesHostList(host, []string{pattern})
		_ = result

		// Test with pattern in a multi-pattern list
		result2 := matchesHostList(host, []string{"safe.example.com", pattern, "other.test.org"})
		_ = result2

		// If single-pattern matches, multi-pattern must also match
		if result && !result2 {
			t.Errorf("matchesHostList(%q, [%q]) = true but matchesHostList(%q, [..., %q, ...]) = false",
				host, pattern, host, pattern)
		}
	})
}

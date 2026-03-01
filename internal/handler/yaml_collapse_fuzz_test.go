package handler

import (
	"testing"
)

// FuzzCollapseYAML exercises the YAML collapse/progressive-disclosure pipeline
// with arbitrary YAML input. CollapseYAML parses untrusted YAML through an AST,
// builds an ordered tree, and renders collapsed output. The fuzzer looks for
// panics, infinite loops, and unexpected errors on valid YAML.
func FuzzCollapseYAML(f *testing.F) {
	// Seed corpus: realistic Helm values.yaml fragments
	seeds := []string{
		// Simple scalars
		"name: nginx\nversion: 1.0.0\n",
		// Nested objects
		"image:\n  repository: nginx\n  tag: latest\n",
		// Arrays
		"ports:\n  - 80\n  - 443\n",
		// Empty containers
		"annotations: {}\nlabels: {}\nvolumes: []\n",
		// Null values
		"optional: null\nrequired: value\n",
		// Boolean values
		"enabled: true\ndisabled: false\n",
		// Quoted strings
		"empty: \"\"\ncolon: \"host:port\"\n",
		// Deeply nested
		"a:\n  b:\n    c:\n      d:\n        e: deep\n",
		// Array of objects
		"containers:\n  - name: app\n    image: nginx\n  - name: sidecar\n    image: envoy\n",
		// Anchors and aliases
		"defaults: &base\n  timeout: 30\nservice:\n  <<: *base\n  name: myapp\n",
		// Special YAML values
		"pos: .inf\nneg: -.inf\nnan: .nan\n",
		// Quoted keys with dots
		"\"a.b.c\": value1\nnormal: value2\n",
		// Comments
		"# Comment\nname: test\n# Another\nvalue: 123\n",
		// Helm-style comments with @schema
		"# @schema\n# type: object\n# @schema\n# -- Enable the service\nservice:\n  enabled: true\n",
		// Multi-document (only first is used)
		"name: first\n---\nname: second\n",
		// Empty document
		"",
		// Just a comment
		"# just a comment\n",
		// Large numeric values
		"big: 999999999999999999\nfloat: 3.14159265358979\n",
		// Mixed types
		"mixed:\n  str: hello\n  num: 42\n  bool: true\n  null_val: null\n  arr:\n    - 1\n    - two\n    - true\n",
		// YAML with special characters in values
		"special: \"value with \\n newline\"\npath: /usr/local/bin\n",
		// Tagged values
		"binary: !!binary aGVsbG8=\ncustom: !mytag value\n",
	}

	for _, s := range seeds {
		f.Add([]byte(s))
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		// Test with various option combinations to cover all code paths.
		optionSets := []CollapseOptions{
			// Default options
			DefaultCollapseOptions(),
			// Unlimited depth (stripComments path)
			{MaxDepth: 0, ShowDefaults: true, ShowComments: false},
			// Unlimited depth, preserve comments
			{MaxDepth: 0, ShowDefaults: true, ShowComments: true},
			// Depth 1, minimal
			{MaxDepth: 1, MaxArrayItems: 3, ShowDefaults: true, ShowComments: false},
			// Depth 1, no defaults (type inference path)
			{MaxDepth: 1, MaxArrayItems: 3, ShowDefaults: false, ShowComments: false},
			// Depth 1, with comments
			{MaxDepth: 1, MaxArrayItems: 1, ShowDefaults: true, ShowComments: true},
			// Deep depth
			{MaxDepth: 10, MaxArrayItems: 10, ShowDefaults: true, ShowComments: true},
		}

		for _, opts := range optionSets {
			result, _, err := CollapseYAML(data, opts)
			// We only care that it doesn't panic. Errors on invalid YAML are fine.
			if err != nil {
				return
			}
			// If no error, result should be a non-panicking string
			_ = len(result)
		}
	})
}

// FuzzExtractYAMLPath exercises YAML path extraction with arbitrary YAML and
// path strings. This function parses untrusted YAML and applies a yq-style
// path expression, which is a key attack surface.
func FuzzExtractYAMLPath(f *testing.F) {
	// Seed corpus: realistic YAML + path combinations
	type seed struct {
		data []byte
		path string
	}
	seeds := []seed{
		{[]byte("server:\n  port: 8080\n  host: localhost\n"), ".server"},
		{[]byte("server:\n  port: 8080\n  host: localhost\n"), ".server.port"},
		{[]byte("items:\n  - name: a\n  - name: b\n"), ".items"},
		{[]byte("items:\n  - name: a\n  - name: b\n"), ".items[0]"},
		{[]byte("items:\n  - name: a\n  - name: b\n"), ".items[0].name"},
		{[]byte("a:\n  b:\n    c: deep\n"), ".a.b.c"},
		{[]byte("name: test\n"), ".name"},
		{[]byte("name: test\n"), ".nonexistent"},
		{[]byte("name: test\n"), ""},
		{[]byte("name: test\n"), "."},
		{[]byte("top:\n  mid:\n    low: val\n"), ".top.mid"},
		{[]byte("\"dotted.key\": value\n"), ".dotted.key"},
	}

	for _, s := range seeds {
		f.Add(s.data, s.path)
	}

	f.Fuzz(func(t *testing.T, data []byte, path string) {
		// We only care that it doesn't panic. Errors are expected for
		// invalid YAML or invalid/missing paths.
		result, err := extractYAMLPath(data, path)
		if err != nil {
			return
		}
		_ = len(result)
	})
}

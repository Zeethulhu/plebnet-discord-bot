package utils

import "testing"

func TestParseArgs(t *testing.T) {
  tests := []struct {
    input     string
    expected  []string
  }{
    {`!echo hello world`, []string{"!echo", "hello", "world"}},
    {`!echo "hello world"`, []string{"!echo", "hello world"}},
    {`!cmd "one two" three`, []string{"!cmd", "one two", "three"}},
    {`!say`, []string{"!say"}},
    {``, []string{}},
  }

  for _, test := range tests {
    got := ParseArgs(test.input)
    if len(got) != len(test.expected) {
      t.Errorf("len(ParseArgs(%q)) = %d; want %d", test.input, len(got), len(test.expected))
      continue
    }

    for i := range got {
      if got[i] != test.expected[i] {
        t.Errorf("ParseArgs(%q)[%d] = %q; want %q", test.input, i, got[i], test.expected[i])
      }
    }
  }
}



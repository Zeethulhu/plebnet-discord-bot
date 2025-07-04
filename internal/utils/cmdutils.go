package utils


// ParseArgs splits on space, but preserves quoted phrases
// e.g. `!echo "hello world"` ->  ["echo", "hello world"]

func ParseArgs(input string) []string {
  parts := []string{}
  current := ""
  inQuotes := false
  
  for _, r := range input {
    switch {
    case r == '"':
      inQuotes = !inQuotes
      if !inQuotes {
        parts = append(parts, current)
        current = ""
      }
    case r == ' ' && !inQuotes:
      if current != "" {
        parts = append(parts, current)
        current = ""
      }
    default:
      current += string(r)
    }
  }
  if current != "" {
    parts = append(parts, current)
  }
  return parts
}

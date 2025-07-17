package lang

import (
  _ "embed"
  "encoding/json"
  "log"
)

//go:embed ../registry.yaml
var registryRaw []byte

type Lang struct {
  Code    string   `json:"code"`
  Name    string   `json:"name"`
  Regions []string `json:"regions"`
  Tier    string   `json:"tier"`
}

type Registry struct {
  Languages []Lang `json:"languages"`
}

var reg Registry

func init() {
  if err := json.Unmarshal(registryRaw, &reg); err != nil {
    log.Fatalf("lang registry parse failed: %v", err)
  }
}

func Codes() []string {
  out := make([]string, 0, len(reg.Languages))
  for _, l := range reg.Languages {
    out = append(out, l.Code)
  }
  return out
}

func IsSupported(code string) bool {
  for _, c := range Codes() {
    if c == code {
      return true
    }
  }
  return false
}

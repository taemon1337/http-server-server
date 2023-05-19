package tls

import (
  "testing"
)

var (
  OPTS = map[string]string{"foo": "bar"}
)

func TestSelfSign(t *testing.T) {
  ss, err := NewSelfSign(OPTS, OPTS)
  if err != nil {
    t.Fatalf("%s", err)
  }

  t.Logf("%s", ss)
}

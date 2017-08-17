package main

import "testing"

type testargs struct {
  args []string
  result uint64
}

func TestMaxBorder(t *testing.T) {
  var tests = []testargs{
    { []string{"name", "1"}, 1 },
    { []string{"name", "1", "2"}, 1 },
    { []string{"name", "123", "2"}, 123 },
  }

  for _, pair := range tests {
    v, _ := MaxBorder(pair.args)
    if v != pair.result {
      t.Error(
        "For", pair.args,
        "expected", pair.result,
        "got", v,
      )
    }
  }
}

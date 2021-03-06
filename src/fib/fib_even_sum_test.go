package main

import "testing"

type testpair struct {
  max uint64
  sum uint64
}

var tests = []testpair{
  { 0, 0 },
  { 1, 0 },
  { 2, 0 },
  { 3, 2 },
  { 4, 2 },
  { 5, 2 },
  { 7, 2 },
  { 8, 2 },
  { 9, 10 },
  { 13, 10 },
  { 21, 10 },
  { 34, 10 },
  { 35, 44 },
  { 55, 44 },
  { 89, 44 },
  { 144, 44 },
  { 145, 188 },
  { 75025, 60696 },
  { 121393, 60696 },
  { 196418, 60696 },
  { 317811, 257114 },
  { 514229, 257114 },
  { 832040, 257114 },
  { 1346269, 1089154 },
  { 2178309, 1089154 },
  { 3524578, 1089154 },
  { 4000000, 4613732 },
  { 5702887, 4613732 },
  { 9227465, 4613732 },
  { 14930352, 4613732 },
  { 24157817, 19544084 },
  { 39088169, 19544084 },
  { 63245986, 19544084 },
  { 102334155, 82790070 },
  { 165580141, 82790070 },
  { 267914296, 82790070 },
  { 433494437, 350704366 },
  { 701408733, 350704366 },
  { 1134903170, 350704366 },
  { 1836311903, 1485607536 },
  { 2971215073, 1485607536 },
  { 1100087778366101931, 889989708002357094 },
  { 4660046610375530309, 3770056902373173214 },
}

func TestFibEvenSum(t *testing.T) {
  for _, pair := range tests {
    v := FibEvenSum(pair.max)
    if v != pair.sum {
      t.Error(
        "For", pair.max,
        "expected", pair.sum,
        "got", v,
      )
    }
  }
}

package main

import "fmt"

func letterCombinations(digits string) []string {
  /* 
  https://leetcode.com/problems/letter-combinations-of-a-phone-number/submissions/
  */
  if digits == "" { return []string{} }
  keyboard := map[string]string{
      "2": "abc",
      "3": "def",
      "4": "ghi",
      "5": "jkl",
      "6": "mno",
      "7": "pqrs",
      "8": "tuv",
      "9": "wxyz",
  }
  ret := []string{""}
  for _, digit := range digits {
    // fmt.Printf("Digit: %s\n", string(digit))

    var tmp[]string
    for _, y := range ret{
      chars := keyboard[string(digit)]
      // fmt.Printf("Y: %s .. chars: %s\n", y, chars)
      
      for _, x := range chars{
        out := fmt.Sprintf("%s%s", string(y), string(x))
        // fmt.Printf("\tOut: %s ... x: %s\n", out, string(x))
        tmp = append(tmp, out)
      }
      
      ret = tmp
    }

  }
  return ret
}

func main() {
  // combinations := letterCombinations("237")
  combinations := letterCombinations("")
  for i, letters := range combinations {
    fmt.Printf("Combination %d: %s\n", i, string(letters))
  }
}

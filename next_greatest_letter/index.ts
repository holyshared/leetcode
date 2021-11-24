function nextGreatestLetter(letters: string[], target: string): string {
  let lo = 0, hi = letters.length
  while (lo < hi) {
    let mi = lo + Math.trunc((hi - lo) / 2)
    if (letters[mi] <= target) {
      lo = mi + 1
    } else {
      hi = mi
    }
  }
  return letters[lo % letters.length]
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts"

Deno.test("letters = [\"c\",\"f\",\"j\"], target = \"a\"", () => {
  assertEquals(nextGreatestLetter(["c","f","j"], "a"), "c")
})

Deno.test("letters = [\"c\",\"f\",\"j\"], target = \"c\"", () => {
  assertEquals(nextGreatestLetter(["c","f","j"], "c"), "f")
})

Deno.test("letters = [\"c\",\"f\",\"j\"], target = \"d\"", () => {
  assertEquals(nextGreatestLetter(["c","f","j"], "d"), "f")
})

Deno.test("letters = [\"c\",\"f\",\"j\"], target = \"g\"", () => {
  assertEquals(nextGreatestLetter(["c","f","j"], "g"), "j")
})

Deno.test("letters = [\"c\",\"f\",\"j\"], target = \"j\"", () => {
  assertEquals(nextGreatestLetter(["c","f","j"], "j"), "c")
})

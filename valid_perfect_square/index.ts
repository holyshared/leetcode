function isPerfectSquare(num: number): boolean {
  if (num < 2) {
    return true
  }

  let left = 2
  let right = Math.trunc(num/2)

  while (left <= right) {
    const x = left + Math.trunc((right - left) / 2)
    const guessSquared = x * x
    if (guessSquared === num) {
      return true
    }

    if (guessSquared > num) {
      right = x - 1
    } else {
      left = x + 1
    }
  }

  return false
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts"

Deno.test("num = 16", () => {
  assertEquals(isPerfectSquare(16), true)
})

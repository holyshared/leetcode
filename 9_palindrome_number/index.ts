function isPalindrome(x: number): boolean {
  let r = 0;
  let v = x;

  if (x < 0 || ((x % 10 === 0) && x > 0)) {
    return false;
  }

  while (v !== 0) {
    const p = v % 10;
    v = Math.trunc(v / 10);
    r = r * 10 + p;
  }
  return r === x;
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("input: 121", () => assert(isPalindrome(121) === true));
Deno.test("input: -121", () => assert(isPalindrome(-121) === false));
Deno.test("input: 10", () => assert(isPalindrome(10) === false));
Deno.test("input: -101", () => assert(isPalindrome(-101) === false));
Deno.test("input: 12321", () => assert(isPalindrome(12321) === true));

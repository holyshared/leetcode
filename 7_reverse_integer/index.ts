function reverse(x: number) : number {
  let v = x;
  let r = 0;
  while (v != 0) {
    const p = v % 10;
    v = Math.trunc(v / 10); 
    r = r * 10 + p;
    if (r < -2147483648 || r > 2147483647) {
      return 0;
    }
  }
  return r;
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("input: 123", () => assert(reverse(123) === 321));
Deno.test("input: -123", () => assert(reverse(-123) === -321));
Deno.test("input: 120", () => assert(reverse(120) === 21));
Deno.test("input: 0", () => assert(reverse(0) === 0));
Deno.test("input: 1534236469", () => assert(reverse(1534236469) === 0));

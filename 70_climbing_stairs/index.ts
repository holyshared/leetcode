function climbStairs(n: number): number {
  const _climbStairs = (cn: number, c: number): number => {
    if (c > cn) {
      return 0;
    }
    if (c === cn) {
      return 1;
    }
    return _climbStairs(cn, c + 1) + _climbStairs(cn, c + 2);
  };
  return _climbStairs(n, 0); 
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("n = 1", () => {
  assertEquals(climbStairs(1), 1);
});

Deno.test("n = 3", () => {
  assertEquals(climbStairs(3), 3);
});

Deno.test("n = 5", () => {
  assertEquals(climbStairs(5), 8);
});

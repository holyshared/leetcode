function climbStairs(n: number): number {
  const _climbStairs = (n: number, step: number, cache: number[]): number => {
    if (step > n) {
      return 0;
    }
    if (step === n) {
      return 1;
    }
    if (cache[step] > 0) {
      return cache[step];
    }
    cache[step] = _climbStairs(n, step + 1, cache) + _climbStairs(n, step + 2, cache);
    return cache[step];
  };
  return _climbStairs(n, 0, new Array(n + 1));
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

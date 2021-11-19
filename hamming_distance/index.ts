function hammingDistance(x: number, y: number): number {
  let xor = x ^ y;
  let distance = 0;
  while (xor != 0) {
    if (xor % 2 === 1) {
      distance += 1;
    }
    xor = xor >> 1;
  }
  return distance;
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("x = 1, y = 4", () => {
  assertEquals(hammingDistance(1, 4), 2);
});

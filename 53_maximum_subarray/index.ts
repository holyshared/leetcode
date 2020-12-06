function maxSubArray(nums: number[]): number {
  let curr = -Infinity;

  return nums.reduce((max, num) => {
    curr = Math.max(num, curr + num);
    return Math.max(max, curr);
  }, -Infinity);
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("nums = [-2,1,-3,4,-1,2,1,-5,4]", () => {
  assertEquals(maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]), 6);
});

Deno.test("nums = [1]", () => {
  assertEquals(maxSubArray([1]), 1);
});

Deno.test("nums = [0]", () => {
  assertEquals(maxSubArray([0]), 0);
});

Deno.test("nums = [-1]", () => {
  assertEquals(maxSubArray([-1]), -1);
});

Deno.test("nums = [-2147483647]", () => {
  assertEquals(maxSubArray([-2147483647]), -2147483647);
});


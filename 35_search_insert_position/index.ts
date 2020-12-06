function searchInsert(nums: number[], target: number): number {
  let i = 0;

  for (; i < nums.length; i++) {
    if (nums[i] === target) {
      return i;
    } else if (nums[i] > target) {
      return i;
    } else {
      continue;
    }
  }

  return i;
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("nums = [1,3,5,6], target = 5", () => {
  assert(searchInsert([1,3,5,6], 5) === 2);
});

Deno.test("nums = [1,3,5,6], target = 2", () => {
  const pos = searchInsert([1,3,5,6], 2);
  console.log(pos);
  assert(pos === 1);
});

Deno.test("nums = [1,3,5,6], target = 7", () => {
  assert(searchInsert([1,3,5,6], 7) === 4);
});

Deno.test("nums = [1,3,5,6], target = 0", () => {
  assert(searchInsert([1,3,5,6], 0) === 0);
});

Deno.test("nums = [1], target = 0", () => {
  assert(searchInsert([1], 0) === 0);
});

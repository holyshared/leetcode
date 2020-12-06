function searchInsert(nums: number[], target: number): number {
  let pivot, left = 0, right = nums.length - 1;

  while (left <= right) {
    pivot = left + Math.round((right - left) / 2);
    if (nums[pivot] === target) {
      return pivot;
    } else if (target < nums[pivot]) {
      right = pivot - 1;
    } else {
      left = pivot + 1;
    }
  }
  return left;
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("nums = [1,3,5,6], target = 5", () => {
  assertEquals(searchInsert([1,3,5,6], 5), 2);
});

Deno.test("nums = [1,3,5,6], target = 2", () => {
  assertEquals(searchInsert([1,3,5,6], 2), 1);
});

Deno.test("nums = [1,3,5,6], target = 7", () => {
  assertEquals(searchInsert([1,3,5,6], 7), 4);
});

Deno.test("nums = [1,3,5,6], target = 0", () => {
  assertEquals(searchInsert([1,3,5,6], 0), 0);
});

Deno.test("nums = [1], target = 0", () => {
  assertEquals(searchInsert([1], 0), 0);
});

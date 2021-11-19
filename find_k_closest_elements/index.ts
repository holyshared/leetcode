/*
Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]

Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]
*/
function findClosestElements(arr: number[], k: number, x: number): number[] {
  if (arr.length === k) {
    return [...arr];
  }
  const results = [];

  let left = 0;
  let right = arr.length;
  let mid = 0;

  while (left < right) {
    mid = Math.trunc((left + right) / 2);
    if (arr[mid] >= x) {
      right = mid;
    } else {
      left = mid + 1;
    }
  }

  left -= 1;
  right = left + 1;

  while (right - left - 1 < k) {
    if (left === -1) {
      right += 1;
      continue;
    }

    if (
      right === arr.length ||
      Math.abs(arr[left] - x) <= Math.abs(arr[right] - x)
    ) {
      left -= 1;
    } else {
      right += 1;
    }
  }

  for (let i = left + 1; i < right; i++) {
    results.push(arr[i]);
  }

  return results;
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("arr = [1,2,3,4,5], k = 4, x = 3", () => {
  assertEquals(findClosestElements([1, 2, 3, 4, 5], 4, 3), [1, 2, 3, 4]);
});

Deno.test("arr = [1,2,3,4,5], k = 4, x = -1", () => {
  assertEquals(findClosestElements([1, 2, 3, 4, 5], 4, -1), [1, 2, 3, 4]);
});

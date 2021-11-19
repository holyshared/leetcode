function findBound(nums: number[], target: number, isFirst: boolean): number {
  let begin = 0
  let end = nums.length - 1

  while (begin <= end) {
    let mid = Math.round((begin + end) / 2)
    if (nums[mid] === target) {
      if (isFirst) {
        if (mid === begin || nums[mid - 1] < target) {
          return mid
        }
        end = mid - 1
      } else {
        if (mid === end || nums[mid + 1] > target) {
          return mid
        }
        begin = mid + 1
      }
    } else if (nums[mid] > target) {
      end = mid - 1
    } else {
      begin = mid + 1
    }
  }
  return -1
}

function searchRange(nums: number[], target: number): number[] {
  const lowerBound = findBound(nums, target, true)
  if (lowerBound === -1) {
    return [-1, -1]
  }
  const upperBound = findBound(nums, target, false)

  return [lowerBound, upperBound]
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("nums = [5,7,7,8,8,10], n = 8", () => {
  assertEquals(searchRange([5,7,7,8,8,10],8), [3,4]);
});

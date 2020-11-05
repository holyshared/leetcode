function removeElement(nums: number[], val: number): number {
  let ip = 0;
  let len = nums.length;

  for (let i = 0; i < len; i++) {
    if (nums[i] !== val) {
      nums[ip] = nums[i];
      ip++;
    }
  }
  nums.splice(ip);

  return nums.length;
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("nums = [3,2,2,3], val = 3", () => {
  const nums = [3,2,2,3];
  const len = removeElement(nums, 3);
  assert(len === 2);
  assert(nums[0] === 2);
  assert(nums[1] === 2);
});

Deno.test("nums = [0,1,2,2,3,0,4,2], val = 2", () => {
  const nums = [0,1,2,2,3,0,4,2];
  const len = removeElement(nums, 2);
  assert(len === 5);
  assert(nums[0] === 0);
  assert(nums[1] === 1);
  assert(nums[2] === 3);
  assert(nums[3] === 0);
  assert(nums[4] === 4);
});

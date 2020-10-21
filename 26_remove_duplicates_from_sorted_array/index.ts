function removeDuplicates(nums: number[]): number {
  if (nums.length <= 1) {
    return nums.length;
  }

  const findEnd = (nums: number[], startAt: number, value: number) => {
    let i = startAt;
    while (nums[i] === value) {
      i++;
    }
    return i;
  };

  let i = 0;

  while (i < nums.length) {
    const end = findEnd(nums, i, nums[i]);
    if (i < (end - 1)) {
      nums.splice(i + 1, end - i - 1);
    }
    i++;
  }

  return nums.length;
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";


Deno.test("input: [1, 1, 2]", () => {
  const nums = [1, 1, 2];
  const ret = removeDuplicates(nums);
  assert(nums.length === ret);
  assert(nums[0] === 1);
  assert(nums[1] === 2);
});


Deno.test("input: [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]", () => {
  const nums = [0,0,1,1,1,2,2,3,3,4];
  const ret = removeDuplicates(nums);
  assert(nums.length === ret);
  assert(nums[0] === 0);
  assert(nums[1] === 1);
  assert(nums[2] === 2);
  assert(nums[3] === 3);
  assert(nums[4] === 4);
});

Deno.test("input: [0, 1, 1, 2, 2, 3, 3, 4, 4]", () => {
  const nums = [0, 1, 1, 2, 2, 3, 3, 4, 4];
  const ret = removeDuplicates(nums);
  assert(nums.length === ret);
  assert(nums[0] === 0);
  assert(nums[1] === 1);
  assert(nums[2] === 2);
  assert(nums[3] === 3);
  assert(nums[4] === 4);
});

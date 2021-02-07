function merge(nums1: number[], m: number, nums2: number[], n: number): void {
  for (let i = 0; i < n; i++) {
    nums1[i + m] = nums2[i];
  }
  nums1.sort((a, b) => a - b);
}


import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3", () => {
  const num1 = [1,2,3,0,0,0];
  const m = 3;
  const num2 = [2,5,6];
  const n = 3;
  merge(num1, m, num2, n);
  assertEquals(num1, [1,2,2,3,5,6]);
});

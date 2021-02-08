function merge(nums1: number[], m: number, nums2: number[], n: number): void {
  const nums1Clone = [...nums1]; 

  let p1 = 0; // num1 pointer
  let p2 = 0; // num2 pointer

  for (let p = 0; p < m + n; p++) {
    // 1. マージしなければならない項目がない
    // 2. 現在の項目より、追加しなければならない項目の方が大きい
    //    AND
    //    既存の項目がまだある
    // nums1 = [1,2,3,0,0,0]
    // nums2 = [2,5,6]
    //
    // [ 1, 2, 3, 0, 0, 0 ]
    // [ 1, 2, 2, 0, 0, 0 ]
    // [ 1, 2, 2, 3, 0, 0 ]
    // [ 1, 2, 2, 3, 5, 0 ]
    // [ 1, 2, 2, 3, 5, 6 ]
    //
    if (p2 >= n || (p1 < m && nums1Clone[p1] < nums2[p2])) {
      nums1[p] = nums1Clone[p1++];
    } else {
      nums1[p] = nums2[p2++];
    }
  }
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

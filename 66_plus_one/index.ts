function plusOne(digits: number[]): number[] {

  let i = digits.length - 1;

  while (i >= 0) {
    if (digits[i] + 1 >= 10) {
      digits[i] = 0;
      i--;
    } else {
      digits[i] = digits[i] + 1;
      return digits;
    }
  }
  if (digits[0] === 0) {
    return [1].concat(digits);
  }
  return digits;
}


import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("digits = [1,2,3]", () => {
  assertEquals(plusOne([1, 2, 3]), [1 , 2, 4]);
});

Deno.test("digits = [4,3,2,1]", () => {
  assertEquals(plusOne([4,3,2,1]), [4,3,2,2]);
});

Deno.test("digits = [0]", () => {
  assertEquals(plusOne([0]), [1]);
});

Deno.test("digits = [1, 9]", () => {
  assertEquals(plusOne([1, 9]), [2, 0]);
});

Deno.test("digits = [9, 9]", () => {
  assertEquals(plusOne([9, 9]), [1, 0, 0]);
});

Deno.test("digits = [9]", () => {
  assertEquals(plusOne([9]), [1, 0]);
});

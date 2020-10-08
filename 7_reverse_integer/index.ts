function reverse(x: number) : number {
  if (x === 0) {
    return 0;
  }
  const varray = x.toString().split('');
  const reverseValues = (values: string[]) : number => {
    let z = true;
    let y = 0;
    const rv = [];

    for (let i = values.length - 1; i >= 0; i--) {
      rv[y] = values[i];
      y++;
    }

    return parseInt(rv.reduce((acc: string[], v: string) => {
      if (v === "0" && z) {
        return acc;
      }
      z = false;
      acc.push(v);
      return acc;
    }, []).join(''), 10);
  };

  const c = (x < 0) ? -(reverseValues(varray.slice(1))) : reverseValues(varray);

  if (c >= -2147483648 && c <= 2147483647) {
    return c;
  }
  return 0;
};

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("input: 123", () => assert(reverse(123) === 321));
Deno.test("input: -123", () => assert(reverse(-123) === -321));
Deno.test("input: 120", () => assert(reverse(120) === 21));
Deno.test("input: 0", () => assert(reverse(0) === 0));
Deno.test("input: 1534236469", () => assert(reverse(1534236469) === 0));

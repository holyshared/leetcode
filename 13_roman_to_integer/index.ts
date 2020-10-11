/*
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

const symbols = {
  'I': 1,
  'V': 5,
  'X': 10,
  'L': 50,
  'C': 100,
  'D': 500,
  'M': 1000,
};
type Keys = keyof (typeof symbols);

function romanToInt(s: string): number {
  const chars = s.split('') as Keys[];

  let i = 0;
  let r = 0;
  while (i < chars.length) {
    const c = chars[i];

    if (c === "I" || c === "X" || c === "C") {
      if (i + 1 <= chars.length - 1) {
        const nc = chars[i + 1];

        if (c === "I") {
          if (nc === "X") {
            r = r + symbols.X - symbols.I;
            i = i + 2;
            continue;
          } else if (nc === "V") {
            r = r + symbols.V - symbols.I;
            i = i + 2;
            continue;
          }
        } else if (c === "X") {
          if (nc === "L") {
            r = r + symbols.L - symbols.X;
            i = i + 2;
            continue;
          } else if (nc === "C") {
            r = r + symbols.C - symbols.X;
            i = i + 2;
            continue;
          }
        } else if (c === "C") {
          if (nc === "D") {
            r = r + symbols.D - symbols.C;
            i = i + 2;
            continue;
          } else if (nc === "M") {
            r = r + symbols.M - symbols.C;
            i = i + 2;
            continue;
          }
        }
      }
    }
    r = r + symbols[c];
    i++;
  }

  return r;
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("input: III", () => assert(romanToInt("III") === 3));
Deno.test("input: IV", () => assert(romanToInt("IV") === 4));
Deno.test("input: IX", () => assert(romanToInt("IX") === 9));
Deno.test("input: XII", () => assert(romanToInt("XII") === 12));
Deno.test("input: XXVII", () => assert(romanToInt("XXVII") === 27));
Deno.test("input: LVIII", () => assert(romanToInt("LVIII") === 58));
Deno.test("input: MCMXCIV", () => assert(romanToInt("MCMXCIV") === 1994));

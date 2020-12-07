function lengthOfLastWord(s: string): number {
  if (s === " ") {
    return 0;
  }

  let i = s.length - 1;
  if (i === 0) {
    return 1;
  }
  for (; i > 0; i--) {
    if (s[i] !== " ") {
      break;
    }
  }
  if (i === 0 && s[i] !== " ") {
    return 1;
  }
  if (i === 0) {
    return 0;
  }

  let wc = 0;

  for (; i >= 0; i--) {
    if (s[i] === " ") {
      return wc;
    } else {
      wc++;
    }
  }

  return wc;
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("s = 'Hello World'", () => {
  assertEquals(lengthOfLastWord("Hello World"), 5);
});

Deno.test("s = 'Hello xWorld'", () => {
  assertEquals(lengthOfLastWord("Hello xWorld"), 6);
});

Deno.test("s = 'Hello xWorld '", () => {
  assertEquals(lengthOfLastWord("Hello xWorld "), 6);
});

Deno.test("s = ' '", () => {
  assertEquals(lengthOfLastWord(" "), 0);
});

Deno.test("s = 'a'", () => {
  assertEquals(lengthOfLastWord("a"), 1);
});

Deno.test("s = 'a '", () => {
  assertEquals(lengthOfLastWord("a "), 1);
});

Deno.test("s = '        '", () => {
  assertEquals(lengthOfLastWord("        "), 0);
});

Deno.test("s = 'day'", () => {
  assertEquals(lengthOfLastWord("day"), 3);
});


Deno.test("s = ' a'", () => {
  assertEquals(lengthOfLastWord(" a"), 1);
});

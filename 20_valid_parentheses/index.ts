const ss : { [key: string]: string }  = {
  "(": ")",
  "{": "}",
  "[": "]"
};
const es : { [key: string]: string } = {
  ")": "(",
  "}": "{",
  "]": "["
};

function isValid(s: string): boolean {
  let len = s.length;

  if (len < 2) {
    return false;
  }

  let i = 0;
  let opened: string[] = [];

  while (i < len) {
    const c = s.charAt(i);
    if (ss[c]) {
      opened.push(c);
    } else if (es[c]) {
      const last = opened.length - 1;
      if (opened[last] !== es[c]) {
        return false;
      }
      opened.splice(opened.length - 1, 1);
    }
    i++;
  }

  return opened.length <= 0;
};

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("input: ()", () => assert(isValid("()")));
Deno.test("input: ()[]{}", () => assert(isValid("()[]{}")));
Deno.test("input: (]", () => assert(!isValid("(]")));
Deno.test("input: ([)]", () => assert(!isValid("([)]")));
Deno.test("input: {[]}", () => assert(isValid("{[]}")));
Deno.test("input: { { } [ ] [ [ [ ] ] ] }", () => assert(isValid("{ { } [ ] [ [ [ ] ] ] }")));
Deno.test("input: [ [ [ ] ] ]", () => assert(isValid("[ [ [ ] ] ]")));

Deno.test("input: ((", () => assert(!isValid("((")));

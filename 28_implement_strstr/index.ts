function strStr(haystack: string, needle: string): number {
  if (haystack === needle) {
    return 0;
  }

  let i = 0;
  let j = 0; 
  let r = -1;
  while (i < haystack.length) {
    if (haystack.substr(i, 1) === needle.substr(j, 1)) {
      if (j >= needle.length - 1) {
        break;
      }
      if (r === -1) {
        r = i;
      }
      i++;
      j++;
    } else {
      i++;
      j = 0;
      r = -1;
    }
  }
  return r;
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("haystack = hello, needle = ll", () => {
  assert(strStr("hello", "ll") === 2);
});

Deno.test("haystack = aaaaa, needle = bba", () => {
  assert(strStr("aaaaa", "bba") === -1);
});

Deno.test("haystack = abcde, needle = de", () => {
  assert(strStr("abcde", "de") === 3);
});

Deno.test("haystack = abcde, needle = ab", () => {
  assert(strStr("abcde", "ab") === 0);
});

Deno.test("haystack = '', needle = ''", () => {
  assert(strStr("", "") === 0);
});

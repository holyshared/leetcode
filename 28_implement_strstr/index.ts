function strStr(haystack: string, needle: string): number {
  if (haystack === needle) {
    return 0;
  }

  if (needle === "") {
    return 0;
  }

  let j = 0; 
  let i = 0;
  let r = -1;

  const p = needle.substr(0, 1);

  const find = (s: string, p: number) => {
    while (p < haystack.length) {
      if (haystack.substr(p, 1) === s) {
        break;
      }
      p++;
    }
    return p;
  };

  while (i < haystack.length) {
    i = find(p, i);
    if (i >= haystack.length) {
      return -1;
    }
    r = i;

    while (j < needle.length) {
      if (haystack.substr(i, 1) === needle.substr(j, 1)) {
        j++;
        i++;
        if (j >= needle.length) {
          return r;
        }
      } else {
        i = r + 1;
        j = 0;
        break;
      }
    }
  }

  return -1;
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

Deno.test("haystack = ab, needle = a", () => {
  assert(strStr("ab", "a") === 0);
});

Deno.test("haystack = '', needle = ''", () => {
  assert(strStr("", "") === 0);
});

Deno.test("haystack = 'a', needle = ''", () => {
  assert(strStr("a", "") === 0);
});

Deno.test("haystack = 'mississippi', needle = 'issip'", () => {
  assert(strStr("mississippi", "issip") === 4);
});

function longestCommonPrefix(strs: string[]): string {
  if (strs.length <= 0) {
    return "";
  }

  if (strs.length <= 1) {
    return strs[0];
  }

  let i = 1;
  let prefix = ""; 
  let macthedPrefix = ""; 
  const base = strs[0]; 
  const remains = strs.slice(1); 

  while (i <= base.length) {
    prefix = base.substr(0, i);
    const matched = remains.every(r => r.startsWith(prefix));
    if (!matched) {
      break;
    }
    macthedPrefix = prefix;
    i++;
  }
  return macthedPrefix;
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("input: dog, racecar, car", () => assert(longestCommonPrefix(["dog", "racecar", "car"]) === ""));
Deno.test("input: flower, flow, flight", () => assert(longestCommonPrefix(["flower", "flow", "flight"]) === "fl"));
Deno.test("input: flower, flow", () => assert(longestCommonPrefix(["flower", "flow"]) === "flow"));
Deno.test("input: flow, flower", () => assert(longestCommonPrefix(["flow", "flower"]) === "flow"));
Deno.test("input: flow", () => assert(longestCommonPrefix(["flow"]) === "flow"));
Deno.test("input: a", () => assert(longestCommonPrefix(["a"]) === "a"));

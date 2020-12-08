function addBinary(a: string, b: string): string {
  const al = a.length;
  const bl = b.length;

  let av = a;
  let bv = b;
  let ml = al;
  let r = "";

  if (al > bl) {
//    return '';

    let diff = al - bl;
    bv = "0".repeat(diff) + b;
    av = a;
    ml = al;

  } else if (al < bl) {
    let diff = bl - al;
    av = "0".repeat(diff) + a;
    bv = b;
    ml = bl;

//    console.log(av);
   /// console.log(bv);

//    return r;
  }


  let c = false;
  for (let i = ml - 1; i >= 0; i--) {
    // 1 1

    // 11
    //  1 
    // 11
    // 11
    if (av[i] === "1" && bv[i] === "1" && c) {
      console.log(3);
      r = "1" + r;
      c = true;
    } else if (av[i] === "1" && bv[i] === "1" || (av[i] === "1" || bv[i] === "1") && c) {
      r = "0" + r;
      c = true;
    } else if (av[i] === "1" || bv[i] === "1" || c) {
      console.log('xxx');
      console.log(r);
      r = "1" + r;
      c = false;
//        return r;
    // 0 0
    } else {
      r = "0" + r;
      c = false;
    }
  }
  if (c) {
    r = "1" + r;
  }

  return r;
}

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("a = '1', b = '10'", () => {
  assertEquals(addBinary("1", "10"), "11");
});

Deno.test("a = '1', b = '11'", () => {
  assertEquals(addBinary("1", "11"), "100");
});

Deno.test("a = '11', b = '1'", () => {
  assertEquals(addBinary("11", "1"), "100");
});

Deno.test("a = '11', b = '10'", () => {
  assertEquals(addBinary("11", "10"), "101");
});

Deno.test("a = '1010', b = '1011'", () => {
  assertEquals(addBinary("1010", "1011"), "10101");
});

Deno.test("a = '1111', b = '1111'", () => {
  assertEquals(addBinary("1111", "1111"), "11110");
});

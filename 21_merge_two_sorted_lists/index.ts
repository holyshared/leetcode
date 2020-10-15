class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
  }
}

function mergeTwoLists(l1: ListNode | null, l2: ListNode | null): ListNode | null {
  const appendUntil = (lv: ListNode, pv: number, out: number[]): ListNode | null => {
    let c = lv as ListNode | null;
    while (c && pv > c.val) {
      out.push(c.val);
      c = c.next;
    }
    return c;
  };

  const _mergeTwoLists = (l1: ListNode | null, l2: ListNode | null): number[] => {
    if (!l1 && !l2) {
      return [];
    }

    if ((l1 !== undefined && l1 !== null) && !l2) {
      return toArray(l1);
    }
  
    if (!l1 && (l2 !== undefined && l2 !== null)) {
      return toArray(l2);
    }

    let l1c = l1;
    let l2c = l2;
    let out = [] as number[];

    if ((l1c as ListNode).val < (l2c as ListNode).val) {
      const end = (l2c as ListNode).val;
      l1c = appendUntil(l1c as ListNode, end, out);
      out = out.concat(_mergeTwoLists(l1c, l2c));
    } else if ((l1c as ListNode).val > (l2c as ListNode).val) {
      const end = (l1c as ListNode).val;
      l2c = appendUntil(l2c as ListNode, end, out);
      out = out.concat(_mergeTwoLists(l1c, l2c));
    } else {
      const end = (l1c as ListNode).val;
      l1c = appendUntil(l1c as ListNode, end + 1, out);
      l2c = appendUntil(l2c as ListNode, end + 1, out);
      out = out.concat(_mergeTwoLists(l1c, l2c));
    }
  
    return out;
  };

  const merged = _mergeTwoLists(l1, l2);
  const result = merged.reverse().reduce((acc: ListNode | null, v: number) => {
    return new ListNode(v, acc);
  }, null);

  return result;
}


import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

const toArray = (ln: ListNode | null) : number[] => {
  if (!ln) {
    return [];
  }
  let c = ln;

  if (!c.next) {
    if (c.val !== undefined) {
      return [c.val];
    } else {
      return [];
    }
  }

  const o = [];

  while (c.next) {
    o.push(c.val);
    c = c.next;
  }
  o.push(c.val);

  return o;
};
/*
Deno.test("input: l1 = [1,2,4], l2 = [1,3,4]", () => {
  const values = toArray(e1);
  assert(values[0] === 1);
  assert(values[1] === 2);
  assert(values[2] === 4);
});
*/

Deno.test("input: l1 = [1,2,4], l2 = [1,3,4]", () => {
  const e11 = new ListNode(1, new ListNode(2, new ListNode(4)));
  const e12 = new ListNode(1, new ListNode(3, new ListNode(4)));

  const values = toArray(mergeTwoLists(e11, e12) as ListNode);
  assert(values[0] === 1);
  assert(values[1] === 1);
  assert(values[2] === 2);
  assert(values[3] === 3);
  assert(values[4] === 4);
  assert(values[5] === 4);
});

Deno.test("input: l1 = [], l2 = []", () => {
  const e11 = null;
  const e12 = null;

  const values = toArray(mergeTwoLists(e11, e12) as ListNode);
  assert(values.length === 0);
});

Deno.test("input: l1 = [], l2 = [0]", () => {
  const e11 = null;
  const e12 = new ListNode(0);

  const values = toArray(mergeTwoLists(e11, e12) as ListNode);
  assert(values.length === 1);
  assert(values[0] === 0);
});


Deno.test("input: l1 = [0], l2 = []", () => {
  const e11 = new ListNode();
  const e12 = null;

  const values = toArray(mergeTwoLists(e11, e12) as ListNode);
  assert(values.length === 1);
  assert(values[0] === 0);
});

Deno.test("input: l1 = [1], l2 = [1]", () => {
  const e11 = new ListNode(1);
  const e12 = new ListNode(1);

  const values = toArray(mergeTwoLists(e11, e12) as ListNode);
  assert(values.length === 2);
  assert(values[0] === 1);
  assert(values[1] === 1);
});

Deno.test("input: l1 = [1, 2], l2 = [1]", () => {
  const e11 = new ListNode(1, new ListNode(2));
  const e12 = new ListNode(1);

  const values = toArray(mergeTwoLists(e11, e12) as ListNode);
  assert(values.length === 3);
  assert(values[0] === 1);
  assert(values[1] === 1);
  assert(values[2] === 2);
});

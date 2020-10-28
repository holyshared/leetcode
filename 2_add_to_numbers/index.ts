class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
  }
}


function toNumber(l: ListNode | null): number {
  if (l === null) {
    return 0;
  }
  let r = 0;
  let p = 0;
  let c = l;
  while (c.next) {
    r +=  p === 0 ? c.val : (10 ** p * c.val);
    c = c.next; 
    p++;
  }
  r +=  p === 0 ? c.val : (10 ** p * c.val);

  return r;
}

function reverseNumber(num: number): number {
  let r = 0;
  let dr = num;
  while (dr !== 0) {
    const d = dr % 10;
    dr = Math.trunc(dr / 10);
    r = (r * 10) + d;
  }
  return r;
}

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
  const sum = toNumber(l1) + toNumber(l2);

  if (sum <= 10) {
    return new ListNode(sum);
  }

  const rnum = reverseNumber(sum);

  const v = rnum.toString().split('');
  const r = new ListNode(Number(v[0]));
  let c = r;

  for (let i = 1; i < v.length; i++) {
    const n = new ListNode(Number(v[i]));
    c.next = n;
    c = n;
  }
  return r;
}

function sum(l: ListNode | null) {
  if (l === null) {
    return 0;
  }
  let c = l;
  let v = [];
  while (c.next) {
    v.push(c.val);
    c = c.next;
  }
  v.push(c.val);
  return Number(v.join(''));
}

import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("input: l1 = [2,4,3], l2 = [5,6,4]", () => {
  const l1 = new ListNode(2, new ListNode(4, new ListNode(3)));
  const l2 = new ListNode(5, new ListNode(6, new ListNode(4)));
  assert(sum(addTwoNumbers(l1, l2)) === 708);
});

Deno.test("input: l1 = [0], l2 = [0]", () => {
  const l1 = new ListNode(0);
  const l2 = new ListNode(0);
  assert(sum(addTwoNumbers(l1, l2)) === 0);
});

Deno.test("input: l1 = null, l2 = [0]", () => {
  const l1 = null;
  const l2 = new ListNode(0);
  assert(sum(addTwoNumbers(l1, l2)) === 0);
});

Deno.test("input: l1 = [0], l2 = null", () => {
  const l1 = new ListNode(0);
  const l2 = null;
  assert(sum(addTwoNumbers(l1, l2)) === 0);
});

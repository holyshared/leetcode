class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
  }
}


function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
  let sum = 0;
  let current = new ListNode(0);
  let result = current;

  while (l1 || l2) {
    if (l1) {
      sum += l1.val;
      l1 = l1.next;
    }

    if (l2) {
      sum += l2.val;
      l2 = l2.next;
    }

    current.next = new ListNode(sum % 10);
    current = current.next;

    sum = sum > 9 ? 1 : 0;
  }

  if (sum) {
    current.next = new ListNode(sum);
  }

  return result.next;
};


import { assert } from "https://deno.land/std@0.73.0/testing/asserts.ts";

function eq(l1: ListNode | null, l2: ListNode | null) {
  let c1 = l1;
  let c2 = l2;

  while (c1 || c2) {
    if (c1 === null || c2 === null) {
      return false;
    }
    if (c1.val !== c2.val) {
      return false;
    }
    c1 = c1.next;
    c2 = c2.next;
  }

  return true;
}


Deno.test("input: l1 = [2,4,3], l2 = [5,6,4]", () => {
  const l1 = new ListNode(2, new ListNode(4, new ListNode(3)));
  const l2 = new ListNode(5, new ListNode(6, new ListNode(4)));
  const expect = new ListNode(7, new ListNode(0, new ListNode(8)));
  assert(eq(addTwoNumbers(l1, l2), expect));
});


Deno.test("input: l1 = [101], l2 = [101]", () => {
  const l1 = new ListNode(1, new ListNode(0, new ListNode(1)));
  const l2 = new ListNode(1, new ListNode(0, new ListNode(1)));
  const expect = new ListNode(2, new ListNode(0, new ListNode(2)));
  assert(eq(addTwoNumbers(l1, l2), expect));
});


Deno.test("input: l1 = [0], l2 = [0]", () => {
  const l1 = new ListNode(0);
  const l2 = new ListNode(0);
  const expect = new ListNode(0);
  assert(eq(addTwoNumbers(l1, l2), expect));
});

Deno.test("input: l1 = null, l2 = [0]", () => {
  const l1 = null;
  const l2 = new ListNode(0);
  const expect = new ListNode(0);
  assert(eq(addTwoNumbers(l1, l2), expect));
});

Deno.test("input: l1 = [0], l2 = null", () => {
  const l1 = new ListNode(0);
  const l2 = null;
  const expect = new ListNode(0);
  assert(eq(addTwoNumbers(l1, l2), expect));
});

Deno.test("input: l1 = [9, 9], l2 = [1]", () => {
  const l1 = new ListNode(9, new ListNode(9));
  const l2 = new ListNode(1);
  const expect = new ListNode(0, new ListNode(0, new ListNode(1)));
  assert(eq(addTwoNumbers(l1, l2), expect));
});

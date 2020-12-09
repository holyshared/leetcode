/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */
class ListNode {
  val: number;
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
  }
}

function deleteDuplicates(head: ListNode | null): ListNode | null {
  if (!head) {
    return null;
  }

  let curr = head;
  let result = new ListNode(curr.val);
  let out = result;

  if (!curr.next) {
    return result;
  }
//  curr = curr.next;

  while (curr.next !== null) {
    if (curr.val === curr.next.val) {
      curr = curr.next;
      continue;
    } else {
      out.next = new ListNode(curr.next.val);
      out = out.next;
      curr = curr.next;
    }
  }

  return result;
};

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("1->1->2", () => {
  const input = new ListNode(1, new ListNode(1, new ListNode(2)));
  const expected = new ListNode(1, new ListNode(2));
  assertEquals(deleteDuplicates(input), expected);
});

Deno.test("1->1->2->3->3", () => {
  const input = new ListNode(1, new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(3)))));
  const expected = new ListNode(1, new ListNode(2, new ListNode(3)));
  assertEquals(deleteDuplicates(input), expected);
});

Deno.test("1->2->3", () => {
  const input = new ListNode(1, new ListNode(2, new ListNode(3)));
  const expected = new ListNode(1, new ListNode(2, new ListNode(3)));
  assertEquals(deleteDuplicates(input), expected);
});

Deno.test("1->1->2->2->3->3", () => {
  const input = new ListNode(1, new ListNode(1, new ListNode(2, new ListNode(2, new ListNode(3, new ListNode(3))))));
  const expected = new ListNode(1, new ListNode(2, new ListNode(3)));
  assertEquals(deleteDuplicates(input), expected);
});

Deno.test("1->2->3->3", () => {
  const input = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(3))));
  const expected = new ListNode(1, new ListNode(2, new ListNode(3)));
  assertEquals(deleteDuplicates(input), expected);
});

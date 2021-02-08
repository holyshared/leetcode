class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val===undefined ? 0 : val)
    this.left = (left===undefined ? null : left)
    this.right = (right===undefined ? null : right)
  }
}

function isSameTree(p: TreeNode | null, q: TreeNode | null): boolean {
  if (p === null && q === null) {
    return true;
  } else if (p !== null && q === null) {
    return false;
  } else if (p === null && q !== null) {
    return false;
  } else if (p!.val !== q!.val) {
    return false;
  }


  if (p?.left !== null && q?.left === null) {
    return false;
  }
  if (p?.left === null && q?.left !== null) {
    return false;
  }
  if (p?.right !== null && q?.right === null) {
    return false;
  }
  if (p?.right === null && q?.right !== null) {
    return false;
  }

  return isSameTree(p!.left, q!.left) && isSameTree(p!.right, q!.right);
};

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";

Deno.test("p = [1,2,3], q = [1,2,3]", () => {
  const left = new TreeNode(1, new TreeNode(2), new TreeNode(3));
  const right = new TreeNode(1, new TreeNode(2), new TreeNode(3));
  assertEquals(isSameTree(left, right), true);
});

Deno.test("p = [1,2], q = [1,null,2]", () => {
  const left = new TreeNode(1, new TreeNode(2), null);
  const right = new TreeNode(1, null, new TreeNode(2));
  assertEquals(isSameTree(left, right), false);
});

Deno.test("p = [1,2,1], q = [1,1,2]", () => {
  const left = new TreeNode(1, new TreeNode(2), new TreeNode(1));
  const right = new TreeNode(1, new TreeNode(1), new TreeNode(2));
  assertEquals(isSameTree(left, right), false);
});

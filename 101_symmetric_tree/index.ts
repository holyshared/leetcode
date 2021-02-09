class TreeNode {
  val: number;
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val===undefined ? 0 : val)
    this.left = (left===undefined ? null : left)
    this.right = (right===undefined ? null : right)
  }
}

function isMirror(left: TreeNode | null, right: TreeNode | null): boolean {
  if (left === null && right === null) {
    return true;
  } else if (left !== null && right === null || left === null && right !== null) {
    return false;
  } else {
    const l = (left as TreeNode);
    const r = (right as TreeNode);
    if (l.val !== r.val) {
      return false;
    }
    return isMirror(l.right, r.left) && isMirror(l.left, r.right);
  }
}

function isSymmetric(root: TreeNode | null): boolean {
  if (root === null) {
    return true;
  }
  return isMirror(root.left, root.right);
};

import { assertEquals } from "https://deno.land/std@0.73.0/testing/asserts.ts";


/*
    1
   / \
  2   2
 / \ / \
3  4 4  3
*/
Deno.test("[1,2,2,3,4,4,3]", () => {
  const root = new TreeNode(
    1,
    new TreeNode(2, new TreeNode(3), new TreeNode(4)),
    new TreeNode(2, new TreeNode(4), new TreeNode(3)),
  );
  assertEquals(isSymmetric(root), true);
});


/*
    1
   / \
  2   2
   \   \
   3    3
*/
Deno.test("[1,2,2,null,3,null,3]", () => {
  const root = new TreeNode(
    1,
    new TreeNode(2, null, new TreeNode(3)),
    new TreeNode(2, null, new TreeNode(3)),
  );
  assertEquals(isSymmetric(root), false);
});

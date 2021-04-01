function ListNode(val, next) {
  this.val = (val===undefined ? 0 : val);
  this.next = (next===undefined ? null : next);
}

var isPalindrome = function(head) {
  let slow = head;
  let fast = head;

  while (fast.next != null && fast.next.next != null) {
    fast = fast.next.next;
    slow = slow.next;
  }

  const reverse = (curr) => {
    let prev = null;
    while (curr != null) {
      next = curr.next;
      curr.next = prev;
      prev = curr;
      curr = next; 
    }
    return prev;
  };

  const rv = reverse(slow.next);

  let p1 = head;
  let p2 = rv;
  let ok = true;

  while (p2 != null) {
    if (p2.val != p1.val) {
      ok = false;
      break;
    }
    p2 = p2.next;
    p1 = p1.next;
  }

  slow.next = reverse(rv);

  return ok;
};


const r = new ListNode(1, new ListNode(2, new ListNode(2, new ListNode(1))));
const ret = isPalindrome(r);

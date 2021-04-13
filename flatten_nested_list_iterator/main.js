class NestedInteger {
  constructor(ni) {
    this.ni = ni;
  }
  isInteger() {
    return typeof this.ni === 'number'; 
  }
  getInteger() {
    if (this.isInteger()) {
      return this.ni;
    } else {
      return null;
    }
  }
  getList() {
    if (!this.isInteger()) {
      return this.ni;
    } else {
      return null;
    }
  }
}


function flatten(ni, outs) {
  return ni.reduce((acc, ni) => {
    if (ni.isInteger()) {
      outs.push(ni.getInteger());
      return outs;
    } else {
      outs = flatten(ni.getList(), outs);
      return outs;
    }
  }, outs)
}



/**
 * @constructor
 * @param {NestedInteger[]} nestedList
 */
var NestedIterator = function(nestedList) {
  this.items = flatten(nestedList, []);
  this.curr = 0;
};


/**
 * @this NestedIterator
 * @returns {boolean}
 */
NestedIterator.prototype.hasNext = function() {
  return this.curr <= this.items.length - 1;
};

/**
 * @this NestedIterator
 * @returns {integer}
 */
NestedIterator.prototype.next = function() {
  const v = this.items[this.curr];
  this.curr++;
  return v;
};

/**
 * Your NestedIterator will be called like this:
 * var i = new NestedIterator(nestedList), a = [];
 * while (i.hasNext()) a.push(i.next());
*/



function main() {
  const ni1 = [
    new NestedInteger([
      new NestedInteger(1),
      new NestedInteger(1)
    ]),
    new NestedInteger(2),
    new NestedInteger([
      new NestedInteger(1),
      new NestedInteger(1)
    ]),
  ];

  // [1,[4,[6]]]
  const ni2 = [
    new NestedInteger(1),
    new NestedInteger([
      new NestedInteger(4),
      new NestedInteger([
        new NestedInteger(6),
      ])
    ]),
  ];

  const outsa = [];
  const it1 = new NestedIterator(ni1);
  while (it1.hasNext()) {
    outsa.push(it1.next());
  }
  console.log(outsa);

  const outsb = [];
  const it2 = new NestedIterator(ni2);
  while (it2.hasNext()) {
    outsb.push(it2.next());
  }
  console.log(outsb);
}
main();

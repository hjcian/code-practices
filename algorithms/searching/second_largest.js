// Problem 2 - Find second largest within O(n) complexity.
const assert = require('assert');

function findSecondLargest(sequence) {
  // A return early pattern, it has better readability for production code.
  if (!Array.isArray(sequence) || sequence.length < 2) return null;
  if (sequence.length === 2) {
    if (sequence[0] === sequence[1]) {
        return null;
    }
    return Math.min(sequence[0], sequence[1]);
  }

  let largest = Number.NEGATIVE_INFINITY;
  let second = Number.NEGATIVE_INFINITY;

  for (const n of sequence) {
    if (n > largest) {
      second = largest;
      largest = n;
    } else if (n > second && n !== largest) {
      second = n;
    }
  }

  return second;
}

assert.strictEqual(findSecondLargest([3, 3, 2, 1]), 2);
assert.strictEqual(findSecondLargest([3, 3, 4, 2, 1]), 3);
assert.strictEqual(findSecondLargest([1, 5, 3, 9, 2]), 5);
assert.strictEqual(findSecondLargest([10, 5]), 5);
assert.strictEqual(findSecondLargest([7, 7, 3, 2]), 3);
assert.strictEqual(findSecondLargest([-1, -5, -3, -2]), -2);
assert.strictEqual(findSecondLargest([5]), null);
assert.strictEqual(findSecondLargest([5, 5]), null);
assert.strictEqual(findSecondLargest([]), null);

console.log('All tests passed');
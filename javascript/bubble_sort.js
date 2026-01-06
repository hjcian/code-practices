const assert = require('assert');

function bubbleSort(arr) {
	if (!Array.isArray(arr) || arr.length <= 1) return arr;
    for (let end = arr.length -1; end > 0; end --) {
        for (let i =0; i < end; i++) {
            if (arr[i] > arr[i+1]) {
                // swap
                const tmp = arr[i];
                arr[i] =  arr[i+1];
                arr[i+1] = tmp;
            }
        }
    }
	return arr;
}


const cases = [
    { name: 'given',
        input: [5, 1, 3, 2, 4], expected: [1, 2, 3, 4, 5] },
	{ name: 'empty slice',
        input: [], expected: [] },
	{ name: 'single element',
        input: [1], expected: [1] },
	{ name: 'already sorted',
        input: [1, 2, 3, 4, 5], expected: [1, 2, 3, 4, 5] },
	{ name: 'reverse sorted',
        input: [5, 4, 3, 2, 1], expected: [1, 2, 3, 4, 5] },
	{ name: 'unsorted',
        input: [3, 1, 4, 1, 5, 9, 2, 6], expected: [1, 1, 2, 3, 4, 5, 6, 9] },
	{ name: 'duplicates',
        input: [2, 2, 1, 1], expected: [1, 1, 2, 2] },
];

for (const { name, input, expected } of cases) {
	const result = bubbleSort([...input]); // use copy to avoid mutating the test case input
	assert.deepStrictEqual(result, expected, `${name}: got ${result} want ${expected}`);
}

console.log('All bubbleSort tests passed');



const { expect } = require('chai')
const composeRanges = require('./composeRanges')
const composeRangesPattern = require('./composeRanges.pattern')

describe('Pattern version', () => {
  const tests = [
    {
      name: 'Example',
      nums: [-1, 0, 1, 2, 6, 7, 9],
      want: ['-1->2', '6->7', '9']
    },
    {
      name: 'Test 2',
      nums: [],
      want: []
    },
    {
      name: 'Test 11',
      nums: [0, 5, 9],
      want: ['0', '5', '9']
    },
    {
      name: 'Test 13',
      nums: [7, 8, 9],
      want: ['7->9']
    },
    {
      name: 'Test 14',
      nums: [0, 1, 9],
      want: ['0->1', '9']
    }
  ]

  tests.forEach(test => {
    it(test.name, () => {
      const expectation = expect(composeRangesPattern(test.nums))
      expectation.to.deep.equal(test.want)
    })
  })
})

describe('Concise version', () => {
  const tests = [
    {
      name: 'Example',
      nums: [-1, 0, 1, 2, 6, 7, 9],
      want: ['-1->2', '6->7', '9']
    },
    {
      name: 'Test 2',
      nums: [],
      want: []
    },
    {
      name: 'Test 11',
      nums: [0, 5, 9],
      want: ['0', '5', '9']
    },
    {
      name: 'Test 13',
      nums: [7, 8, 9],
      want: ['7->9']
    },
    {
      name: 'Test 14',
      nums: [0, 1, 9],
      want: ['0->1', '9']
    }
  ]

  tests.forEach(test => {
    it(test.name, () => {
      const expectation = expect(composeRanges(test.nums))
      expectation.to.deep.equal(test.want)
    })
  })
})

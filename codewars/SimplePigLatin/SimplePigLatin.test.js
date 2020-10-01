const { expect } = require('chai')
const pigIt = require('./SimplePigLatin')

describe('Test', () => {
  const tests = [
    {
      name: 'Example 1',
      str: 'Pig latin is cool',
      want: 'igPay atinlay siay oolcay'
    },
    {
      name: 'Example 2',
      str: 'This is my string',
      want: 'hisTay siay ymay tringsay'
    },
    {
      name: 'Example 3',
      str: 'Hello world !',
      want: 'elloHay orldway !'
    }
  ]
  tests.forEach(test => {
    it(test.name, () => {
      const expectation = expect(pigIt(test.str))
      expectation.to.deep.equal(test.want)
    })
  })
})

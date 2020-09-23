function isBeautifulString (inputString) {
  const charCount = new Array(26).fill(0)
  for (const c of inputString) {
    charCount[c.charCodeAt(0) - 97] += 1
  }
  for (let i = 0; i < 25; i++) {
    if (charCount[i] < charCount[i + 1]) {
      return false
    }
  }
  return true
}

(() => {
  const testCases = [
    {
      name: 'Ex 1',
      inputString: 'bbbaacdafe',
      expect: true
    },
    {
      name: 'Ex 2',
      inputString: 'aabbb',
      expect: false
    },
    {
      name: 'Ex 3',
      inputString: 'bbc',
      expect: false
    },
    {
      name: 'Test 9',
      inputString: 'zaa',
      expect: false
    }
  ]
  for (const test of testCases) {
    const got = isBeautifulString(test.inputString)
    if (test.expect === got) {
      console.log(`[${test.name}] pass`)
    } else {
      console.log(`[${test.name}] failed, expect ${test.expect}, got ${got}`)
    }
  }
})()

function happyNumber(n: number): boolean {
    let memory = new Set()
    let current: number = n

    while (true) {
        console.log(`=== Iteration (current=${current})===`)
        if (current === 1) {
            return true
        } else if (memory.has(current)) {
            return false
        }

        const nStr:string = current.toString()
        let nextValue: number = 0
        for (let c of nStr) {
            let cNum: number = parseInt(c)
            nextValue+=cNum*cNum
            console.log(cNum, nextValue)
        }        
        memory.add(current)
        current = nextValue
    }
}

function main () {
  const testCases = [
    {
      name: 'Example',
      n: 19,
      expect: true
    }
  ]
  for (let test of testCases) {
    const got = happyNumber(test.n)
    if (got !== test.expect) {
        console.error(`[${test.name}] expect ${test.expect}, got ${got}`)
    }
  }
}

main()

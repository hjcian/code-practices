interface FactorialCache {
    [num: number]: number[]
}

interface PrimeCounter {
    [num: number]: number
}

function primeFactorsRecursiveCache(n: number, divisor: number, cache: FactorialCache): number[] {
    if (divisor > n) {
        return []
    }
    if (cache[n] !== undefined) {
        return cache[n]
    }
    if (n % divisor == 0) {
        return [divisor].concat(primeFactorsRecursiveCache(n / divisor, divisor, cache))
    } else {
        divisor += 1
        return primeFactorsRecursiveCache(n, divisor, cache)
    }
}

function primeFactorsRecursive(n: number, divisor: number): number[] {
    if (divisor > n) {
        return []
    }
    if (n % divisor == 0) {
        return [divisor].concat(primeFactorsRecursive(n / divisor, divisor))
    } else {
        divisor += 1
        return primeFactorsRecursive(n, divisor)
    }
}

function primeFactors(n: number): number[] {
    let factors: number[] = []
    let divisor: number = 2;

    while (n > 2) {
        if (n % divisor == 0) {
            factors.push(divisor);
            n = n / divisor;
        }
        else {
            divisor++;
        }
    }
    return factors;
}

class PrimeCountHandler {
    primeCounter: Map<number, number>
    constructor() {
        this.primeCounter = new Map<number, number>()
    }
    public update(primes: number[]) {
        primes.forEach(prime => {
            let count: number | undefined = this.primeCounter.get(prime)
            if (count === undefined) {
                this.primeCounter.set(prime, 1)
            } else {
                count++
                this.primeCounter.set(prime, count)
            }
        })
    }
    public format(): string {
        let ret: string = ''
        for (let [number, count] of this.primeCounter) {
            ret += `${number}(${count})*`
        }
        return ret.slice(0, -1)
    }
}

function factorial_fun(N: number): string {
    let factorialCache: FactorialCache = {}
    let primeCountHandler = new PrimeCountHandler()

    for (let i = 2; i <= N; i++) {
        const primes: number[] = primeFactorsRecursiveCache(i, 2, factorialCache)
        factorialCache[i] = primes
        primeCountHandler.update(primes)
    }

    return primeCountHandler.format()
}

(() => {

    const testCases = [
        {
            name: "Ex1",
            N: 5,
            expect: "2(3)*3(1)*5(1)"
        },
        {
            name: "Ex2",
            N: 6,
            expect: "2(4)*3(2)*5(1)"
        },
        // {
        //     name: "Ex3",
        //     N: 20,
        //     expect: "2(18)*3(8)*5(4)*7(2)*11(1)*13(1)*17(1)*19(1)"
        // },
        {
            name: "Test Cache",
            N: 8,
            expect: "2(7)*3(2)*5(1)*7(1)"
        },
    ]

    for (let test of testCases) {
        const got: string = factorial_fun(test.N)
        if (got !== test.expect) {
            console.log(`[${test.name}] expect '${test.expect}', got '${got}'`)
        }
    }
})()

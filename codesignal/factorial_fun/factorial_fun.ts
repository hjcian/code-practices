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

function factorial_fun(N: number): string {
    let factorial = 1
    for (let i = 2; i <= N; i++) {
        factorial *= i
    }

    const primes: number[] = primeFactorsRecursive(factorial, 2)
    console.log(N, primes)
    return ""
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
        {
            name: "Ex3",
            N: 20,
            expect: "2(18)*3(8)*5(4)*7(2)*11(1)*13(1)*17(1)*19(1)"
        }
    ]

    for (let test of testCases) {
        const got: string = factorial_fun(test.N)
        if (got !== test.expect) {
            console.log(`[${test.name}] expect '${test.expect}', got '${got}'`)
        }
    }
})()

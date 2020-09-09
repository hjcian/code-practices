function calcGCD(a: number, b: number): number {
    // Greatest Common Divisor
    return b ? calcGCD(b, a%b) : a
}

function reduce(numerator: number,denominator: number): number[] { 
    const gcd: number = calcGCD(numerator, denominator)
    return [ numerator / gcd, denominator / gcd ]
}  

function convetToSecond(dateStr: string): number {
    const hms: string[] = dateStr.split(":")
    return +hms[0] * 3600 + +hms[1] * 60 + +hms[2]
}

function videoPart(part: string, total: string): number[] {
    const partNum: number = convetToSecond(part)
    const totalNum: number = convetToSecond(total)    
    return reduce(partNum, totalNum)
}

class Test {
    name: string
    part: string
    total: string
    want: number[]
    constructor(name:string, part: string, total: string, want:number[]) {
        this.name = name
        this.part = part
        this.total = total
        this.want = want
    }
}

function main (): void {
    const tests: Test[] = [
        new Test(
            "Example", 
            "02:20:00", 
            "07:00:00", 
            [1, 3]
            ),        
        ]

    tests.forEach(test => {
        const got: number[] = videoPart(test.part, test.total)        

        if (got.length !== test.want.length || !got.every((value, index) => value === test.want[index])) {
            console.log(`[${test.name}] expect ${test.want}, got ${got}`)
        }
    })
}

main()
const calculatorMixin = Base => class extends Base {
  calc () {
    console.log('calc')
  }
}

// Same as:
// function calculatorMixin (Base) {
//   return class extends Base {
//     calc () {
//       console.log('calc')
//     }
//   }
// }

const randomizerMixin = Base => class extends Base {
  randomize () {
    console.log('randomize')
  }
}

class Foo {}
class Bar extends randomizerMixin(calculatorMixin(Foo)) { }

const demoMixins = () => {
  const b = new Bar()
  b.calc()
  b.randomize()
}

demoMixins()

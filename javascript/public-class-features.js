class PublicClass {
    static StaticField = 123
    static StaticMethod = () => {
      console.log('call StaticMethod()')
    }

    instanceField = 456
    instanceMethod = () => {
      console.log('call instanceMethod() and', this.instanceField)
    }

    constructor (name, age) {
      this.name = name
      this.age = age
    }

    show () {
      console.log('=============== show ===============')
      console.log(' class.StaticField:', PublicClass.StaticField)
      console.log('this.instanceField:', this.instanceField)
      console.log('         this.name:', this.name)
      console.log('          this.age:', this.age)
      console.log('=============== show ===============')
    }
}

//
// Class examples
//
console.log('Class.StaticField:', PublicClass.StaticField)
PublicClass.StaticMethod()

console.log('Class.instanceField:', PublicClass.instanceField) // undefined
// PublicClass.instanceMethod() // TypeError: PublicClass.instanceMethod is not a function

//
// instance examples
//
const instance = new PublicClass()
console.log('instance.StaticField:', instance.StaticField) // undefined
// instance.StaticMethod() // TypeError: instance.StaticMethod is not a function

console.log('instance.instanceField', instance.instanceField)
instance.instanceMethod()

//
// check these fields
//
const me = new PublicClass('max', 31)
me.show()
// =============== show ===============
//  class.StaticField: 123
// this.instanceField: 456
//          this.name: max
//           this.age: 31
// =============== show ===============

PublicClass.StaticField = 'new value'
me.instanceField = 'new instance value'
me.name = 'new max'
me.age = 310
me.show()
// =============== show ===============
//  class.StaticField: new value
// this.instanceField: new instance value
//          this.name: new max
//           this.age: 310
// =============== show ===============

const you = new PublicClass('cian', 99)
you.show()
// =============== show ===============
//  class.StaticField: new value
// this.instanceField: 456
//          this.name: cian
//           this.age: 99
// =============== show ===============

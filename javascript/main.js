class Polygon {
  constructor (height, width) {
    this.height = height
    this.width = width
  }

  // same as python's @property decorator
  get area () {
    return this.calcArea()
  }

  calcArea () {
    return this.height * this.width
  }
}

class Point {
  constructor (x, y) {
    this.x = x
    this.y = y
  }

  static distance (a, b) {
    const dx = a.x - b.x
    const dy = a.y - b.y
    return Math.sqrt(dx * dx + dy * dy)
  }
}

const demoGetter = () => {
  const polygon = new Polygon(123, 456)
  console.log(polygon.height)
  console.log(polygon.width)
  console.log(polygon.area) // = p.calcArea()
}

const demoStatic = () => {
  const a = new Point(123, 456)
  //   console.log(a.distance()) // TypeError: a.distance is not a function
  const b = new Point(789, 123)
  console.log(Point.distance(a, b)) // 靜態方法不需要實體化它所屬類別的實例就可以被呼叫，它也無法被已實體化的類別物件呼叫
}

const main = () => {
  console.log('hello world')
  demoGetter()
  demoStatic()
}

main()

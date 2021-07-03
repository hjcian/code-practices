class ClassWithPrivateAccessor {
    #message;

    get #decoratedMessage () {
      return `🎬 ${this.#message} 🛑`
    }

    set #decoratedMessage (msg) {
      this.#message = msg
    }

    constructor () {
      this.#decoratedMessage = 'hello world'
      console.log(this.#decoratedMessage)
    }

    foobar () {
      console.log('foobar')
      this.#foobar()
    }

    #foobar () {
      console.log('I am private foobar')
    }
}

const c = new ClassWithPrivateAccessor()
c.foobar()

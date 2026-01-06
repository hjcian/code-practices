function exampleScope() {
  // x is accessible outside the if block because it's function-scoped
  console.log(x); // Output: undefined

  if (true) {
    var x = 10;   // Function-scoped
    let y = 20;   // Block-scoped
    const z = 30; // Block-scoped
    console.log(x); // Output: 10
    console.log(y); // Output: 20
    console.log(z); // Output: 30
  }

  // x is accessible outside the if block because it's function-scoped
  console.log(x); // Output: 10

  // y and z are NOT accessible outside the if block (block-scoped)
  // console.log(y); // ReferenceError: y is not defined
  // console.log(z); // ReferenceError: z is not defined
}
exampleScope();

function exampleHositing() {
    console.log(a); // Output: undefined (due to hoisting)
    var a = 5;

    // console.log(b); // ReferenceError: Cannot access 'b' before initialization
    let b = 10;

    // console.log(c); // ReferenceError: Cannot access 'c' before initialization
    const c = 15;
}

exampleHositing();

function exampleReassignmentAndRedeclaration() {
    var a = 1;
    a = 2; // Reassignment is allowed
    console.log(a); // Output: 2

    var a = 3; // Redeclaration is allowed with var
    console.log(a); // Output: 3

    let b = 1;
    b = 2; // Reassignment is allowed
    console.log(b); // Output: 2

    // let b = 3; // Redeclaration is NOT allowed with let (uncommenting this line will cause an error: SyntaxError: Identifier 'b' has already been declared)
    const c = 1;
    // c = 2; // Reassignment is NOT allowed with const (uncommenting this line will cause an error: TypeError: Assignment to constant variable.)
    console.log(c); // Output: 1

    // const c = 3; // Redeclaration is NOT allowed with const (uncommenting this line will cause an error: SyntaxError: Identifier 'c' has already been declared)
}

exampleReassignmentAndRedeclaration();
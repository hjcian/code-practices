// Mutation: Changes the original object
const userA = { name: "Max", age: 35 };
const userAA = userA;
userA.age = 36;
console.log(userA); // { name: "Max", age: 36 }
console.log(userAA); // { name: "Max", age: 36 }, same reference as userA
console.log(userA === userAA); // true (Same reference)

// Immutability: Creates a new object (spread operator)
const userB = { name: "Max", age: 35 };
const userBB = { ...userB};
userB.age = 36;
console.log(userB); // { name: "Max", age: 35 }
console.log(userBB); // { name: "Max", age: 36 }, different reference
console.log(userB === userBB); // false (Different references)

let userC = { name: "Max", age: 35, address: { city: "New York", zip: "10001" } };
let userCC = userC; // Reference assignment
console.log(userC === userCC); // true (Same reference)
userC = { ...userC, age: 36 }; // Shallow copy, only top-level properties are copied
console.log(userC); // { name: "Max", age: 36, address: { city: "New York", zip: "10001" } }
console.log(userC === userCC); // false (Different references)
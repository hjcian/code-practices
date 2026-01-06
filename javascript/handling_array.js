// Problem 6 - Handling array
// Given below data, please write code to output all users under 40 years old in below format:
// 1. Mr. Daniel Deng (age 11)
// 2. Mrs. Maria Hanington (age 33)
// ...

const users = [
  {
    firstName: 'Freddie',
    lastName: 'Hong',
    gender: 'male',
    age: 32,
    married: true,
  },
  {
    firstName: 'Shaquille',
    lastName: 'Fang',
    gender: 'male',
    age: 3,
    married: false,
  },
  {
    firstName: 'Justin',
    lastName: 'Fan',
    gender: 'male',
    age: 42,
    married: true,
  },
  {
    firstName: 'Sophia',
    lastName: 'Liu',
    gender: 'female',
    age: 12,
    married: false,
  },
  {
    firstName: 'Maxwell',
    lastName: 'Jeng',
    gender: 'male',
    age: 43,
    married: false,
  },
];


users
    .filter(user => user.age < 40)
    .map(user => {
        const title = user.gender === 'male' ? 'Mr.' : (user.married ? 'Mrs.' : "Miss");
        return `${title} ${user.firstName} ${user.lastName} (age ${user.age})`;
    })
    .forEach((output, index) => {
        console.log(`${index + 1}. ${output}`);
    });

const NonAlphabetPattern = new RegExp('[^a-zA-Z]')

function transform (word) {
  if (word.length === 1 && NonAlphabetPattern.exec(word) !== null) {
    return word
  }
  return `${word.slice(1)}${word[0]}ay`
}

function pigIt (str) {
  const ret = str.split(' ').map(word => transform(word)).join(' ')

  return ret
}

module.exports = pigIt

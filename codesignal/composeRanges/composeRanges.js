function composeRanges (nums) {
  if (nums.length === 0) { return [] }

  const ret = []
  let start = nums[0]

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] - nums[i - 1] !== 1) {
      const end = nums[i - 1]
      ret.push(start === end ? `${end}` : `${start}->${end}`)

      start = nums[i]
    }
  }

  const end = nums[nums.length - 1]
  ret.push(start === end ? `${end}` : `${start}->${end}`)

  return ret
}

module.exports = composeRanges

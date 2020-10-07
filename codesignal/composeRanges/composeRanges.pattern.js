function Handler (start) {
  this.start = start
  this.ret = []
  return {
    isNotContinuous: (left, right) => (right - left) !== 1,
    push: (end) => {
      this.ret.push(this.start === end ? `${end}` : `${this.start}->${end}`)
    },
    refresh: (newStart) => {
      this.start = newStart
    },
    answer: () => this.ret
  }
}

function composeRangesPattern (nums) {
  if (nums.length === 0) { return [] }

  const handler = new Handler(nums[0])

  for (let i = 1; i < nums.length; i++) {
    if (handler.isNotContinuous(nums[i - 1], nums[i])) {
      handler.push(nums[i - 1])
      handler.refresh(nums[i])
    }
  }

  handler.push(nums[nums.length - 1])

  return handler.answer()
}

module.exports = composeRangesPattern

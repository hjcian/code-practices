function Handler (start) {
  this.start = start
  this.ret = []
  return {
    isNotContinuous: (left, right) => (right - left) !== 1,
    pushEnd: (end) => {
      this.ret.push(this.start === end ? `${end}` : `${this.start}->${end}`)
    },
    updateStart: (newStart) => {
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
      handler.pushEnd(nums[i - 1])
      handler.updateStart(nums[i])
    }
  }

  handler.pushEnd(nums[nums.length - 1])

  return handler.answer()
}

module.exports = composeRangesPattern

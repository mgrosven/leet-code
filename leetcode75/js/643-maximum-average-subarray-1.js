var findMaxAverage = function(nums, k) {
    let sum = 0
    for (let i = 0; i < k; i++) {
        sum += nums[i]
    }
    let maxAverage = sum / k
    for (let i = k; i < nums.length; i++) {
        sum = sum - nums[i - k] + nums[i]
        maxAverage = Math.max(maxAverage, sum / k)
    }
    return maxAverage    
}

var nums = [1,12,-5,-6,50,3], k = 4
console.log(findMaxAverage(nums, k)) // 12.75
var nums = [5], k = 1
console.log(findMaxAverage(nums, k)) // 5
var nums = [0,4,0,3,2], k = 1
console.log(findMaxAverage(nums, k)) // 4
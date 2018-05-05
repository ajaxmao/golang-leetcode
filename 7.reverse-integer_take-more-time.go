import "math"

/*
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (24.38%)
 * Total Accepted:    405.3K
 * Total Submissions: 1.7M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
 * this problem, assume that your function returns 0 when the reversed integer
 * overflows.
 *
 */
func reverse(x int) int {
	var ret int
	adjust := 1
	if x < 0 {
		x = -1 * x
		adjust = -1
	}
	var nums []int
	for {
		g := x % 10
		nums = append(nums, g)
		x = (x - g) / 10
		if x == 0 {
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		ret += nums[i] * int(math.Pow10(len(nums)-i-1))
	}

	if adjust > 0 && ret > 2147483647 {
		return 0
	}
	if adjust < 0 && ret > 2147483648 {
		return 0
	}
	return ret * adjust
}

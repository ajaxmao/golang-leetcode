import (
	"fmt"
	"strconv"
)

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
	adjust := 1
	if x < 0 {
		x = -1 * x
		adjust = -1
	}
	y := strconv.Itoa(x)
	bytes := []byte(y)
	for i, j := 0, len(y)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	y = string(bytes)
	ret, err := strconv.Atoi(y)
	if err != nil {
		fmt.Println(err.Error())
	}
	if x > 0 && ret > 2147483647 {
		return 0
	}
	if x < 0 && ret < 2147483648 {
		return 0
	}
	return ret * adjust
}

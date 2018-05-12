/*
 * [367] Valid Perfect Square
 *
 * https://leetcode.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (38.71%)
 * Total Accepted:    72.3K
 * Total Submissions: 186.8K
 * Testcase Example:  '16'
 *
 * Given a positive integer num, write a function which returns True if num is
 * a perfect square else False.
 *
 *
 * Note: Do not use any built-in library function such as sqrt.
 *
 *
 * Example 1:
 *
 * Input: 16
 * Returns: True
 *
 *
 *
 * Example 2:
 *
 * Input: 14
 * Returns: False
 *
 *
 *
 * Credits:Special thanks to @elmirap for adding this problem and creating all
 * test cases.
 */
func isPerfectSquare(num int) bool {
	//1.  虽然通过了，但是太low
	/* bot := 0*/
	//top := num
	//for bot+1 < top {
	//med := (top + bot) / 2
	//if (med * med) < num {
	//bot = med
	//} else if med*med > num {
	//top = med
	//} else {
	//return true
	//}
	//}
	//if num == 1 {
	//return true
	//}
	/*return false*/
	//2. 参考1＋3＋5＋7＋…＋（2n-1）＝n2*n, 连续的奇数之和
	/* for i := 1; num > 0; i += 2 {*/
	//num -= i
	/*}*/
	//return num == 0
	//3.
	/* for i := 1; i <= num/i; i++ {*/
	//if i*i == num {
	//return true
	//}
	//}
	/*return false*/
	r := num
	for r*r > num {
		r = (r + num/r) / 2
	}
	return r*r == num
}

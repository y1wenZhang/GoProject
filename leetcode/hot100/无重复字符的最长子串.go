package hot100

/*
示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	var (
		res   int
		right int
		tmp   map[byte]int
		n     int = len(s)
	)
	tmp = make(map[byte]int)
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(tmp, s[i-1])
		}

		for right < n && tmp[s[right]] == 0 {
			tmp[s[right]]++ // 滑动窗口
			right++
		}
		res = max(res, right-i)

	}

	return res
}

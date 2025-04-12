package hot100

/*
找到字符串中所有字母异位词

给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。


示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
*/

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	var (
		ns     = len(s)
		np     = len(p)
		tp, ts [26]int
	)

	if ns < np {
		return res
	}

	for _, m := range p {
		tp[m-'a']++
	}

	for i := 0; i < np; i++ {
		ts[s[i]-'a']++
	}

	if ts == tp {
		res = append(res, 0)
	}

	for i := np; i < ns; i++ {
		ts[s[i]-'a']++
		ts[s[i-np]-'a']--

		if ts == tp {
			res = append(res, i-np+1)
		}
	}

	return res
}

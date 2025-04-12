package hot100

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。



示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2
*/

/*
	解析：
	pre[i], 前i个元素之和
	前pre[i-1]元素之和 = pre[i] - 第i个元素
	分析：[i....j], 从i到j的元素之和等于pre[j]-pre[i]
*/

func subarraySum(nums []int, k int) int {
	tn := make(map[int]int)

	res, pre := 0, 0

	tn[0] = 1
	for _, num := range nums {
		pre += num
		if _, ok := tn[pre-k]; ok {
			res += tn[pre-k]
		}
		tn[pre]++
	}
	return res
}

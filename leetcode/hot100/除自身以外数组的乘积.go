package hot100

/*

给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]
*/

/*
分析：每个位置的乘积 =  左侧乘积 * 右侧乘积

注意： 起始位置左侧乘积设为1
      结束位置右侧乘积设为1
*/

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	// 左侧乘积
	rl := make([]int, n)
	rl[0] = 1

	for i := 1; i < n; i++ {
		rl[i] = rl[i-1] * nums[i-1]
	}

	// 右侧乘积
	rr := make([]int, n)
	rr[n-1] = 1
	for i := n - 1; i >= 0; i-- {
		res[i] = rl[i] * rr[i] // 左侧乘积 * 右侧乘积
		// 更新右侧乘积
		if i >= 1 {
			rr[i-1] = rr[i] * nums[i]
		}
	}
	return res
}

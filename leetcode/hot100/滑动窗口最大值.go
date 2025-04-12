package hot100

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]
*/

/*
分析：
单点队列
Pop: 如果弹出的元素为出口元素，则弹出，其他情况不处理
Push：如果push的元素大于入口元素，则弹出入口元素，直到push的元素小于或等于入口元素
*/

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewQueue()
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front())
	for i := k; i < len(nums); i++ {
		queue.Push(nums[i])
		queue.Pop(nums[i-k])
		res = append(res, queue.Front())
	}
	return res
}

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{data: []int{}}
}

func (q *Queue) Push(num int) {
	for !q.IsEmpty() && num > q.Back() {
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, num)
}

func (q *Queue) Pop(num int) {
	if !q.IsEmpty() && num == q.Front() {
		q.data = q.data[1:]
	}
}

func (q *Queue) Front() int {
	return q.data[0]
}

func (q *Queue) Back() int {
	return q.data[len(q.data)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

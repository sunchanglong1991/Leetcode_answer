//https://leetcode-cn.com/problems/3sum/
/*
    三数和为0  结果不重复
*/
func threeSum(nums []int) [][]int {
    n := len(nums)
    // 利用有序数组实现结果三元祖 不重复
    sort.Ints(nums)
    // 保留符合的结果
    result := make([][]int, 0)


    // nums:     [a,  ...b,  ...c]
    //             first second third
    /*
        限定first 指定元素
        双指针： second、third 分别指向first 后续数据的首尾位置
    */
    for first := 0; first < n; first++ {
        // 需要和上一次枚举的数不相同 去除结果重复 [-1, -1, 0, 1] --> [-1, 0, 1]
        if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
        // 对应的指针初始指向数组的最右端
        third := n - 1
        // 将指定值当成两数的和, first + second + third = 0 --> second + third = -first
        target := -1 * nums[first]
        // 枚举 b
        for second := first + 1; second < n; second++ {
            // 需要和上一次枚举的数不相同
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue
            }
            // 需要保证 b 的指针在 c 的指针的左侧
            for second < third && nums[second] + nums[third] > target {
                third--
            }
            // 如果指针重合，随着 b 后续的增加
            // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
            if second == third {
                break
            }
            if nums[second] + nums[third] == target {
                result = append(result, []int{nums[first], nums[second], nums[third]})
            }
        }
    }
    return result
}
        

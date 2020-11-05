/*
方法：
    当 nums[j] 与给定的值相等时，递增 j 以跳过该元素。只要 nums[j] !=val，我们就复制 nums[j]到 nums[i] 并同时递增两个索引。重复这一过程，直到 j 到达数组的末尾，该数组的新长度为 i
*/
func removeElement(int[] nums, int val) int {
    if len(nums) == 0 {
        return 0
    }

    i := 0
    // 遍历数组所有元素
    for j := 0; j < len(nums); j++ {
        // 不等就替换 计数 移位
        if(nums[j] != val) {
            nums[i] = nums[j]
            // i++ 一定要放在赋值后, 因为要从头开始 跟val 比较
            i++
        }
    }
        
    return i
}



// 优化： 当要删除的元素很少时，第一种方案会进行多余的复制操作；优化方案是 交换尾部 进行删除尾部操作，不同时后移
func removeElement(nums []int, val int) int {
    i, n := 0, len(nums)
    for (i < n) {
        if (nums[i] == val) {
            nums[i] = nums[n - 1]
            n--
        } else {
            i++
        }
    }
    return n
}

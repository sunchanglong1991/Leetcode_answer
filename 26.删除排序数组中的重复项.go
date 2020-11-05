/*
方法：
    首先数组为有序数组，不需要考虑后者出现小于的情况
    利用双指针，i, j  为最后输出的数组下标，j 为从 i +1 开始遍历源数组的下标
*/

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    i := 0
    // 遍历数组剩余元素
    for j := 1; j < len(nums); j++ {
        // 不等就替换 计数 移位
        if(nums[j] != nums[i]) {
            i++
            nums[i] = nums[j]
        }
    }
        
    return i+1
}

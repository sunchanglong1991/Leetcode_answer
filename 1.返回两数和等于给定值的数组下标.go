/*
    借助 map 是否存在指定值判断
*/
func twosum(nums []int, target int) []int  {
    if len(nums) == 0 {
        return nil
    }


    // result 留存 value 对应的数组下标 [value][bit]
    result := make(map[int]int)
    for i:= 0; i< len(nums); i++ {
        // 另一半
        another := target - nums[i]
        if _, ok := result[another]; ok {
            return []int{result[another], i}
        }
        // 保存当前value 的下标; 等到 遍历到另一半时取出下标
        result[nums[i]] = i
    }


    return nil
}

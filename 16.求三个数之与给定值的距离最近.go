func threeSumClosest_1(nums []int, target int) int {
   // 设置一个最大值用来比较出最小值 MaxInt32
   n, res, diff := len(nums), 0, math.MaxInt32
   // 确保数组数量满足条件
   if n > 2 {
      sort.Ints(nums)
      // 留最后俩个双指针 left right n-2
      for i := 0; i < n-2; i++ {
         // 保证和上一次枚举的元素不相等
         if i > 0 && nums[i] == nums[i-1] {
            continue
         }


         // 留最后一个指针 reft: i+1; right: n-1
         // j < k 限制遍历范围
         for j, k := i+1, n-1; j < k; {
            if j > (i+1) && nums[j] == nums[j-1] {
               continue
            }
            
            if k < (n-2) && nums[k] == nums[k+1] {
               continue
            }

            sum := nums[i] + nums[j] + nums[k]
            if sum == target {
               return res
            }

            if abs(sum-target) < diff {
               res, diff = sum, abs(sum-target)
            }

            if sum > target {
               k--
            } else {
               j++
            }
         }
      }
   }
   return res
}

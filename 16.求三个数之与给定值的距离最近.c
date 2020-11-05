/*
方法：
     最接近 即为差值的绝对值最小；排序后、初始时，b 指向位置 i+1，即左边界；c 指向位置 n-1，即右边界。在每一步枚举的过程中，我们用 a+b+c来更新答案，并且：
	* 
如果 a+b+c \geq \textit{target}a+b+c ≥ target，那么就将 c  向左移动一个位置; (如果 a+b+c≥target，并且我们知道 b到 c 这个范围内的所有数是按照升序排序的，那么如果 c 不变而 b 向右移动，那么 a+b+c 的值就会不断地增加，显然就不会成为最接近 target 的值)
	* 
如果 a+b+c < target，那么就将 b 向右移动一个位置.(如果 a+b+c<target，并且我们知道 b 到 c 这个范围内的所有数是按照升序排序的，那么如果 b 不变而 c 向左移动，那么 a+b+c 的值就会不断地减小，显然就不会成为最接近 target 的值)

*/

int comp(const void *a, const void *b) { return *(int *)a - *(int *)b; }
int threeSumClosest(int *nums, int numsSize, int target) {
    int n = numsSize;
    qsort(nums, n, sizeof(int), comp);
    int best = 1e7;


    // 根据差值的绝对值来更新答案


    // 枚举 a
    for (int i = 0; i < n; ++i) {
        // 保证和上一次枚举的元素不相等
        if (i > 0 && nums[i] == nums[i - 1]) {
            continue;
        }
        // 使用双指针枚举 b 和 c
        int j = i + 1, k = n - 1;
        while (j < k) {
            int sum = nums[i] + nums[j] + nums[k];
            // 如果和为 target 直接返回答案
            if (sum == target) {
                return target;
            }
            if (abs(sum - target) < abs(best - target)) {
                best = sum;
            }
            if (sum > target) {
                // 如果和大于 target，移动 c 对应的指针
                int k0 = k - 1;
                // 移动到下一个不相等的元素
                while (j < k0 && nums[k0] == nums[k]) {
                    --k0;
                }
                k = k0;
            } else {
                // 如果和小于 target，移动 b 对应的指针
                int j0 = j + 1;
                // 移动到下一个不相等的元素
                while (j0 < k && nums[j0] == nums[j]) {
                    ++j0;
                }
                j = j0;
            }
        }
    }
    return best;
}

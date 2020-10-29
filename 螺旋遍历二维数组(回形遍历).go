/*
    螺旋遍历二维数组; 回形遍历
*/
package main


import (
    "fmt"
    "os"
)


func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}


func spiralOrder(str [][]int) ([]int, error) {
    result := []int{}
    // 行
    m := len(str)
    // 列
    n := len(str[0])
    fmt.Println(m,"行", n, "列")
    if (str == nil || m == 0 || n == 0) {
        return result, fmt.Errorf("arr 行列不应为0")
    }


    // atom.Min(m,n)+1)/2 是求回形遍历的次数
    // 大循环，从外向内逐层遍历矩阵; 
    // i:可以理解为行下角标
    // j:可以理解为列下角标
    for i:= 0; i < (Min(m,n)+1)/2; i++ {
        //从左到右遍历“上边”
        for j := i; j<n-1; j++ {
            result = append(result, str[i][j])
        }
        //从上到下遍历“右边” 列不变，行逐步++
        for j := i+1; j < m-i; j++{
            /* 
                j=i+1 表示去除刚才遍历的行的下一行（从上到下 列不变但行++ 所以j++）
                n-1 表示 一共的列，(n-1) -i 表示每次遍历一圈i++ 列向内缩减一圈 所以-i
            */
            result = append(result, str[j][(n-1)-i])
        }
        //从右到左遍历“下边” 行不变 列逐步--
        for j := i+1; j < n-i; j++ {
            /*
                (m-1)-i 表示每遍历一圈 向内缩小一圈
                (n-1)-j 同理
            */
            result = append(result, str[(m-1)-i][(n-1)-j])
        }
        //从下到上遍历“左边”
        for j := i+1; j < m-1-i; j++ {
            /*
                i：列不变 就代表遍历的圈数
                (m-1)-j 表示逐步递减 
            */
            result = append(result, str[(m-1)-j][i])
        }
    }
    return result, nil
}


func main()  {
    arr := [][]int{
        { 1,  2,  3,  4,  5  },
        { 6,  7,  8,  9,  10 },
        { 11, 12, 13, 14, 15 },
        { 16, 17, 18, 19, 20 },
    }


    arr2 := [][]int{
        {1, 2, 3, 4, 5},
        {6, 7, 8, 9, 10},
        {11, 12, 13, 14, 15},
        {16, 17, 18, 19, 20},
        {21, 22, 23, 24, 25},
    }


    arr3 := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
        {10, 11, 12},
    }
    //fmt.Println(arr, arr2)


    str, err := spiralOrder(arr)
    if err != nil {
        fmt.Println("arr err ")
        os.Exit(2)
    }
    fmt.Println(str)


    str, err = spiralOrder(arr2)
    if err != nil {
        fmt.Println("arr2 err ")
        os.Exit(2)
    }
    fmt.Println(str)


    str, err = spiralOrder(arr3)
    if err != nil {
        fmt.Println("arr2 err ")
        os.Exit(2)
    }
    fmt.Println(str)
}

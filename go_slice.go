// 请给我一些数字！
numbers := []int{1, 2, 3, 4, 5}
 
log(numbers)         // 1. [1 2 3 4 5]
log(numbers[2:])     // 2. [3 4 5]
log(numbers[1:3])    // 3. [2 3]
 
// 有趣的是，你不能使用负数索引
//
// 来自 Python 的 numbers[:-1] 并不能正确工作，相反的是，
// 你必须这样做：
//
log(numbers[:len(numbers)-1])    // 4. [1 2 3 4]
 
// 可读性真实“太好了”，Pike 先生！干的漂亮！
//
// 现在，让我们在尾部插入一个6：
//
numbers = append(numbers, 6)
 
log(numbers) // 5. [1 2 3 4 5 6]
 
// 把3从numbers中移除 :
//
numbers = append(numbers[:2], numbers[3:]...)
 
log(numbers)    // 6. [1 2 4 5 6]
 
// 想要插入一些数？别急，这里是一个Go语言*通用*最佳实践
//
// 我特别喜欢。。。哈哈哈。
//
numbers = append(numbers[:2], append([]int{3}, numbers[2:]...)...)
 
log(numbers)    // 7. [1 2 3 4 5 6]
 
// 为了拷贝一份切片，你需要这样做：
//
copiedNumbers := make([]int, len(numbers))
copy(copiedNumbers, numbers)
 
log(copiedNumbers)    // 8. [1 2 3 4 5 6]
 
//还有一些其他操作。。。

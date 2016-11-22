package number

// ShellSort : 希尔排序
func ShellSort(data []int) {
	//一共有多少列: 4 -> 2 -> 1
	for col := len(data) / 2; col >= 1; col = col / 2 {
		//对每一列处理 0 -> col（即每一次循环，就是处理该列的排序）
		for i := 0; i < col; i++ {
			//从每一列的第二个数开始进行插入排序（因为只有 >= 2 个元素时，才有可比性）
			n := 1
			//i + col * n 表示该列的下一个元素的位置 , 它要 < len(data)
			for i+col*n < len(data) {
				//保存当前元素的位置（因为它要不断地与位置比它小的进行交换（即插入排序）
				k := i + col*n
				for k >= 0 && (k-col) >= 0 {
					//不断地与前一个元素进行比较（直到 前一个元素 >= 即可停止）这里是开始排序
					if data[k] < data[k-col] {
						data[k], data[k-col] = data[k-col], data[k]
						k = k - col
					} else {
						break
					}
				}
				//该列的下一个元素
				n++
			}
		}
	}
}

package number

// InsertSort : 插入排序
func InsertSort(data []int) {
	//从第二个元素开始排序，第一个为 i = 0
	for i := 1; i < len(data); i++ {
		//保存当前位置的一个副本，因为它要不断地减小，（因为要不断地缩小并比较之前的数据，以便进行交换）
		currentIndex := i
		//不断与前一个相比较
		for currentIndex-1 >= 0 {
			//符合条件则交换
			if data[currentIndex] < data[currentIndex-1] {
				data[currentIndex], data[currentIndex-1] = data[currentIndex-1], data[currentIndex]
				currentIndex--
			} else {
				break
			}
		}
	}
}

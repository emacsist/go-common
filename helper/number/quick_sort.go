package number

// QuickSort : 快速排序
func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}

	//假设第一个元素为 0，并且假设为"中值"
	middleValue := data[0]

	//这个是索引值了，不是表示大小了, 头尾指针
	head, tail := 0, len(data)-1

	//
	i := 1

	for head < tail {
		//如果下一个元素，比中值小，那从头开始交换位置
		if data[i] < middleValue {
			data[head], data[i] = data[i], data[head]

			//因为已经交换了，头指针向前移动一个
			head++
			//则进行下一个元素比较
			i++
		} else {
			//这个表示下一个元素比中值大，则从 尾开始交换
			data[tail], data[i] = data[i], data[tail]
			tail--
		}
	}

	data[head] = middleValue
	QuickSort(data[:head])

	QuickSort(data[head+1:])

}

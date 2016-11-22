package number

// SelectSort : 选择排序
func SelectSort(data []int) {
	swapIndex := 0
	for i := 0; i < len(data); i++ {
		index, _ := min(data, swapIndex)
		//fmt.Printf("min index = %v, i = %v\n", index, i)
		swap(data, swapIndex, index)
		swapIndex++
	}
}

func swap(data []int, src int, dest int) {
	data[src], data[dest] = data[dest], data[src]
	//fmt.Printf("after swap => %v\n", data)
}

func min(data []int, startIndex int) (index int, d int) {
	d = data[startIndex]
	index = startIndex
	for i := startIndex + 1; i < len(data); i++ {
		next := data[i]
		if next < d {
			d = next
			index = i
		}
	}
	return
}

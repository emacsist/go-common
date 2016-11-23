package number

// HeapSort : 堆排序
func HeapSort(data []int) {
	if len(data) <= 1 {
		return
	}
	maxIndex := len(data) - 1
	buildAllMaxHeap(data)
	for i := len(data) - 1; i >= 0; i-- {
		data[0], data[maxIndex] = data[maxIndex], data[0]
		localMaxHeap(data[:maxIndex], 0)
		maxIndex--
	}
}

//全局创建最大堆
func buildAllMaxHeap(data []int) {
	//从最后一颗树的根节点开始处理
	for i := (len(data) - 1) / 2; i >= 0; i-- {
		localMaxHeap(data, i)
	}
}

//每一颗小树中，要保持最大堆
func localMaxHeap(data []int, root int) {
	left := getLeft(root)
	right := getRight(root)

	biggest := root

	if left < len(data) && data[left] > data[biggest] {
		biggest = left
	}
	if right < len(data) && data[right] > data[biggest] {
		biggest = right
	}

	if biggest != root {
		data[root], data[biggest] = data[biggest], data[root]
		localMaxHeap(data, biggest)
	}
}

// 获取元素的左节点
func getLeft(index int) int {
	left := 2*index + 1
	return left
}

// 获取元素的右节点
func getRight(index int) int {
	right := 2 * (index + 1)
	return right
}

// 获取元素的父节点
func getParentIndex(index int) int {
	if index == 0 {
		return index
	}
	if index%2 == 0 {
		return index/2 - 1
	}
	return index / 2
}

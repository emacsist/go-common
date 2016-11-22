package number

import "math"

// HeapSort : 堆排序
func HeapSort(data []int) {
	if len(data) == 0 {
		return
	}
	startIndex := len(data) - 1
	for startIndex >= 0 {
		parentIndex := getParentIndex(startIndex)
		if data[parentIndex] < data[startIndex] {
			swapParentAndChild(data, parentIndex, startIndex)
		}
		startIndex--
	}
	//fmt.Printf("%v\n", data)
	HeapSort(data[1:])
}

//交换父子节点
func swapParentAndChild(data []int, parent, child int) {
	data[parent], data[child] = data[child], data[parent]
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

// 返回完全二叉树的层数
func getLevel(data []int) int {
	return int(math.Ceil(math.Sqrt(float64(len(data)))))
}

// n ： 层，n >= 1b
// 完全二叉树：最多结点： 2^n - 1

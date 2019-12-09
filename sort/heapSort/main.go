package main

import "fmt"

// 堆排序
// 1. 建堆
// 先找到最后一个非叶子节点的节点，调整该节点和其子树使其满足堆的条件，然后从右往左依次调整上一个节点，直到根节点也满足堆的条件，建堆完成。
// 整个建堆过程不需要引入额外的数组
// 2. 排序
// 排序也不需要引入额外的数组，因此是原地排序算法

func HeapSort(array []int) []int {
	// 为了简单，元素下标从 1 开始
	buildHeap(array, len(array)-1)
	fmt.Println(array)
	sort(array)
	return array
}

func buildHeap(array []int, n int) {
	// 由于完全二叉树的特性，最后一个非叶子节点的位置是 n/2
	lastNotLeafNode := n / 2
	for ; lastNotLeafNode > 0; lastNotLeafNode-- {
		heapify(array, lastNotLeafNode, n)
	}
}

// 自顶向下堆化过程
// 需要两个参数：
// startPostion 堆化的起始位置，一般是整个树的堆顶位置或者子树的堆顶位置
// lastNodePostion 堆中最后一个元素的位置，用来判断堆化的终止条件
func heapify(array []int, startPostion int, lastNodePostion int) {
	i := startPostion
	li := 2 * i
	ri := 2*i + 1
	for {
		maxIndex := i
		// 说明存在左子节点
		if li <= lastNodePostion && array[li] > array[maxIndex] {
			maxIndex = li
		}
		// 说明存在右子节点
		if ri <= lastNodePostion && array[ri] > array[maxIndex] {
			maxIndex = ri
		}
		if maxIndex == i {
			break
		}
		array[i], array[maxIndex] = array[maxIndex], array[i]
		i = maxIndex
		li = 2 * i
		ri = 2*i + 1
	}
}

// 注意排序也只排下标 1-n 的元素
func sort(array []int) {
	i := len(array) - 1
	// 排序的思路很巧妙，交换第一个元素和最后一个元素，于是便把最大的元素放在了数组的末端
	// 同时整个方式还恰好模拟了删除堆顶元素的过程，交换元素后，重新自顶向下堆化即可。
	for i > 1 {
		array[1], array[i] = array[i], array[1]
		i--
		heapify(array, 1, i)
	}
}

func main() {
	fmt.Println(HeapSort([]int{0, 8, 5, 3, 4, 1, 9}))
}

// 堆排序不如快速排序快的原因：
// 堆排序访问元素的方式不是连续的，是跳着访问的，这样对 cpu 的缓存不太友好
// 对于同样的数据，在排序过程中堆排序算法的数据交换次数要多于快速排序

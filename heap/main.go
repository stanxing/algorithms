package main

import "errors"

import "fmt"

// 堆
// 需要满足下列两个条件：
// 1.堆必须是一个完全二叉树
// 2.堆中每个节点都必须大于等于（小于等于）其子树中每个节点的值

// 最大堆：堆中每个节点的值大于等于其子树中每个节点的值
// 最小堆：堆中每个节点的值小于等于其子树中每个节点的值

// 堆的存储：
// 首选数组，因为内存利用率高，节点之间的关系也很清晰，单纯的通过下标就可以找到
// 假设数组从下标 1 开始存储堆中元素，若数组中节点的下标为 i ，其左子节点下标为 2*i，右子节点下标为 2*i+1，
// 其父节点的下标为 i/2
// 假设数组从下标 0 开始存储堆中元素，若数组中节点的下标为 i， 其左子节点下标为 2*i+1，右子节点下标为 2*i+1，
// 其父节点的下标为 (i-1)/2

type Heap struct {
	array   []int // 堆的底层数组
	maxSize int   // 堆中可以存储的最大元素个数
	count   int   // 堆中已经存储的元素数量
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		array:   make([]int, capacity+1),
		maxSize: capacity,
		count:   0,
	}
}

// 堆中插入数据的思路：
// 先把元素插到堆的末尾，然后从下往上堆化，修正不符合堆规则的元素
// 下面的例子以最大堆为例
func (h *Heap) Insert(value int) error {
	if h.count >= h.maxSize {
		return errors.New("full heap")
	}
	h.count++
	h.array[h.count] = value
	i := h.count
	// 堆顶元素的下标为 1
	// 如果插入到末尾的元素大于他的父节点，应该和父节点交换，一直循环下去，直到根节点
	pi := i / 2
	for pi > 0 && h.array[i] > h.array[pi] {
		h.array[i], h.array[pi] = h.array[pi], h.array[i]
		i = pi
		pi = i / 2
	}
	return nil
}

func (h *Heap) DeleteHeapTop() error {
	if h.count == 0 {
		return errors.New("empty heap")
	}
	// 删除堆顶元素的思路是把最后一个元素替换堆顶的位置
	// 然后从上到下堆化，修正堆中的元素
	// 堆顶的起始元素是 1
	h.array[1] = h.array[h.count]
	h.count--
	i := 1
	li := 2 * i
	ri := li + 1
	for li <= h.count {
		maxIndex := li
		// 说明存在右节点
		if ri <= h.count {
			if h.array[li] < h.array[ri] {
				maxIndex = ri
			}
		}
		if h.array[maxIndex] > h.array[i] {
			h.array[i], h.array[maxIndex] = h.array[maxIndex], h.array[i]
			i = maxIndex
			li = 2 * i
			ri = li + 1
		} else {
			break
		}
	}
	return nil
}

func main() {
	heap := NewHeap(8)
	heap.Insert(1)
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(6)
	heap.Insert(7)
	heap.Insert(4)
	heap.Insert(3)
	fmt.Println(heap)
	heap.DeleteHeapTop()
	fmt.Println(heap)
}

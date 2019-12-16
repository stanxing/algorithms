package main

import "fmt"

// 重建二叉树
// 输入二叉树的前序和中序遍历的结果，请重建二叉树，并输出它的头节点。
// 例如：前序遍历为 [1,2,4,7,3,5,6,8]，中序遍历为 [4,7,2,1,5,3,8,6]
// 保证输入的两个二叉树是有效的
type Node struct {
	value int
	left  *Node
	right *Node
}

// 解题思路：
// 前序遍历的第一个节点一定是根节点，中序遍历的根节点在中间，将数组分成了左右两个子树。因此通过前序和中序就可以唯一确定一棵二叉树。
// 递归的解决左右子树，即可构建出完整的二叉树。
func constructBinaryTree(preOrder, inOrder []int) *Node {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}
	return constructBinaryTreeRecursively(preOrder, inOrder, 0, len(preOrder)-1, 0, len(inOrder)-1)
}

// 分析这个递归程序需要准备的参数：
// 前序数组是主要用来读取元素、构建二叉树的数组，中序数组起辅助作用，主要用来确定左右子树的元素个数。因此每次递归需要传入构建该二叉树需要的元素位置。
// preOrder, startPreOrder, endPreOrder
// inOrder, startInOrder, endInOrder
func constructBinaryTreeRecursively(preOrder, inOrder []int, startPreOrder, endPreOrder, startInOrder, endInOrder int) *Node {
	if startPreOrder > endPreOrder || startInOrder > endInOrder {
		return nil
	}
	rootValue := preOrder[startPreOrder]
	root := &Node{
		value: rootValue,
	}
	// 在 inOrder 中寻找根节点的位置 i ，划分左右子树
	i := startInOrder
	for ; i <= endInOrder; i++ {
		if inOrder[i] == rootValue {
			break
		}
	}
	// 左子树有 leftLength 个元素，右子树有 rightLength 个元素
	leftLength := i - startInOrder
	rightLength := endInOrder - i

	// 构建左子树
	if leftLength > 0 {
		root.left = constructBinaryTreeRecursively(preOrder, inOrder, startPreOrder+1, startPreOrder+leftLength, startInOrder, i-1)
	}
	// 构建右子树
	if rightLength > 0 {
		root.right = constructBinaryTreeRecursively(preOrder, inOrder, startPreOrder+leftLength+1, endPreOrder, i+1, endInOrder)
	}
	return root
}

func main() {
	root := constructBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	fmt.Println(root)
	fmt.Println(root.right.right.left)
}

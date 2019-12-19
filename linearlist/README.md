# 线性结构

线性结构指的是数据元素之间存在着“一对一”的线性关系的数据结构，即除了第一个和最后一个数据元素之外，其它数据元素都是首尾相接的。

按照操作来分，可以分为

- 线性表
- 栈
- 队列

## 线性表

线性表是一种由同类数据类型的元素构成的有序序列的线性结构。包含以下概念：

- 长度：表中元素个数
- 线性表没有元素时称为空表
- 表的起始位置称为表头，结束位置称为表尾

按照存储结构划分：

- 顺序表
- 链表
  - 单向链表
  - 双向链表
  - 循环链表

### 线性表和数组的区别

线性表是逻辑层的概念，是一个数据集合更高级的抽象，用来管理和组织有序线性序列。
数组是用来分配一块连续内存存储数据的底层物理结构，是连续存储结构的一种实现。
线性表的顺序存储结构底层是用数组来实现的。
链表是另一个种非连续的存储结构，和数组对应。

- [数组和线性表的区别](https://www.jianshu.com/p/2008e29c39e2)

### 抽象数据类型

类型名称： List
数据集：n 个同类数据类型 Element 的有序序列

操作集： 定义变量和数据类型 List L; Element e;

- MakeEmptyList() List：创建一个空表 List
- FindElement(K int, L List) Element：查找位置为 k 的元素并返回它
- Find(e Element, L List) int：根据元素查找它在线性表中第一次出现的位置
- Insert(e Element, i int, L List)：在线性表的第 i 个元素前插入一个新元素 e
- Delete(i int, L List)：在线性表中删除第 i 个元素
- Length(L List) int：返回线性表的长度

## 顺序表

参考代码实现

## 链表

参考代码实现
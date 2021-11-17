# leetcode 解题报告

## 0001

两数相加：

* 使用 hash 表优化查找

## 0002

大数加法问题：

* 需要注意最后进位的处理
* 单链表操作，使用伪节点表示 head，能够简化处理

## 0003

最长无重复子串:

* 使用滑动窗口机制，遇到重复字符时，变更起点为重复字符的下一个字符
* 重复字符的判断，使用 hash 表优化查找
* 查找到的重复字符位置需要判断是否位于当前子串的区间

## 0005

最长回文子串：

* 需要注意对奇、偶字符串的判断
* 可以抽象出一个实用函数检查某一个特定起始的字符串是否为回文子串

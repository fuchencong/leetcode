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

## 0004

两个有序数组的中位数

* 二分搜索的极致用法
* 细节太多，好难

## 0005

最长回文子串：

* 需要注意对奇、偶字符串的判断
* 可以抽象出一个实用函数检查某一个特定起始的字符串是否为回文子串

## 0006

zigzag：

* 需要找到数据规律
* 按周期来计算

## 0007

反转数字：

* 需要注意，32 位有符号数的范围，正数的绝对值要比负数的绝对值少 1
* 统一转换成字符串进行比较或者逐位比较

## 0008

字符串转换为数字：

* 同样注意计算过程中的数字溢出问题，如果机器只支持 32 位数字，那么就需要考虑溢出问题，此时只能用字符串比较的方式
* 可以考虑使用 max/min 等 python 函数来限制范围

## 0009

反转数字：

* 直接对数进行反转，然后比较数的方法，需要考虑反转之后会不会溢出的问题，所以一种安全的方式是转换为字符串进行比较
* Python 解法：直接用 str() 和 str[::-1] 将数字转换为字符串，然后比较字符串即可。但是非 Python 解法可能需要自己构建字符列表
* 另外一种高效的方式：整数运算直接进行到中途即可（这种解法需要理解其数学规律）

## 0011

容器如何盛水最多：

* 贪心算法：关键需要找到一个数学规律
* i 和 j 分别是容器的最左和最右端，a[i] 和 a[j] 算出面积后，和当前最大值进行比较
* 接下来只有移动小的，才有可能找到更大值

## 0012

整数转化为罗马数字：

* 找准数学规律，核心就是判断数字里有多少个 1000，多少个 900，依次类推，多少个 1
* 使用一个列表，记录下数字与符号的映射关系。从大到小开始逐个整除，如果当前商为 0，则整除列表中的下一个
* 循环退出的调节：当前数为 0

## 0013

罗马数字转换为整数：

* 逐个遍历字符串，如果当前字符小于下一个字符，说明这两个字符需要连续读取
* 使用 hash 表保存字符与整数的映射关系

## 0014

最长公共前缀：

* Pythonic 代码：使用 zip 可以直接得到各个字符串相应位置字符的列表，
* 使用 set 判断该列表是否只有一个字符，从而判断出是否是公共字符


## 0105

二叉树还原：

* 根据前序遍历、中序遍历就可以找准根节点、左子树的范围、右子树的范围
* 递归还原左子树、右子树

## 0240

有序矩阵搜索：

* 每次比较矩阵 [0, columns-1] 位置的数，从而每次可以去掉一整行或一整列
* 注意退出的条件


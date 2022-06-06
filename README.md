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

## 0015

3 数之和

* 数排序之后，固定一个数，然后使用双指针的方法来遍历剩余的数. 遍历过程根据当前的值来调整指针
* 需要注意，过滤掉当前重复的数

## 0016

3 数之和最近的值

* 数排序之后，固定一个数，然后使用双指针的方法来遍历剩余的数. 遍历过程根据当前的值来调整指针
* 需要注意，过滤掉当前重复的数

## 0017

电话字符组合

* 使用 map 保存数字字符到英文字符的替换
* 使用递归方法保存所有的组合


## 0019

删除倒数第 k 个节点

* 双指针法
* 使用 dummy 头节点会简化一些，但是逻辑更复杂

## 0021

合并两个有序列表

* 任意一个链表达到尾部，需要注意
* 递归写法简单一些

## 0022

生成指定的括号

* 递归加回溯思想，递归时，每次根据当前左右括号的情况，选择括号
* 递归结束的条件：字符长度达到 2*n


## 0024

交换链表中的两个节点

* 一定要注意最后更新 currNode 时，要选择交换后的节点

## 0027

移除数组中的指定元素，不允许开辟新的空间：

* 直接遍历数组，通过一个索引来指示新写入的位置
* 只要当前元素不等于指定元素，就可以写入该位置

## 0028

strStr：

* 寻找子串，两次遍历
* 首先判断字符串的长度，可以简化后续边界的处理

## 0033

旋转数组搜索

* 找出有序的部分后，判断 target 是否在这个有序的部分
* 否则就在剩余部分

## 0034

有序数组 target 的第一个和最后一个元素：

* 仍然是采用二分搜索的思想
* 只不过搜索过程中，需要判断出什么时候往前搜索第一个元素，什么时候往后搜索最后一个元素

## 0035

有序数组插入位置

* 普通的二分搜索，没有什么特别的

## 0038

字符串计数

* 正序遍历，每次找到该字符的最后一个字符，得到统计个数

## 0039

数组的任意个数元素组合，等于指定 target

* 递归 + 回溯：选中当前数之后，判断是否等于 target
* 递归结束条件：当前数大于 target

## 0041

数组中缺失的最小正整数

* 核心思想是，最小正整数一定在 1-n 之间或者 n+1
* 将每个元素放置到对应的位置，用位置本身下标就可以作为 key
* 使用元素交换来保证数不丢失
* 交换时避免出现死循环

## 0043

字符串乘法

* 使用相应位置计算和的方式
* 注意倒序
* 去除前导 0

## 0045

跳跃游戏，调到指定位置所需要做少的步数

* 简单的动态规划，从后往前开始计算
* minSteps[i] 表示从位置 i 到 end 的最小步数
* 这样就可以算出所有的 minSteps[i-1]

## 0046

无重复数字的排列组合

* 本质上是深度优先递归搜索
* 需要保存当前已经使用的数字
* 找到了一个排列，本次递归结束，开始回溯

## 0047

有重复数字的排列组合

* 如何在迭代时去重，关键逻辑是，同一层，如果之前选过了这个数，就不用再选择这个数了
* 如果判断同一层选择过这个数，判断这个数是否和前一个数相等，且前一个数当前没有选中，则之前选择过这个数

## 0048

二维数组旋转

* 分别计算出旋转后的 4 个位置，进行替换
* 只需要遍历一般的数据，另一半会在循环内被处理

## 0049

字符串分组

* 对字符串进行排序
* 直接使用 hash table 保存所有的一组字符串

## 0050

power(x, n)

* 递归解决，每次计算 power(x, n/2)
* 注意奇、偶的的处理

## 0053

最大的子数组

* 每次看当前和，如果是一个小于的 0 的数，就可以忽略本次计算
* 每次计算，都和当前最大数进行比较
* 或者使用动态规划方法，f[i] 表示以 i 结尾的最大数，f[i+1] 等于 f[i-1] 或 nums[i]

## 0054

矩阵的蛇形遍历

* 考研边界条件的处理，最清晰的处理是，每次遍历完一个方向，显示增加值，进行下一个方向
* 这样就不会出现重复计数

## 0055

是否能跳跃到目的地

* 使用动态规划，从后往前计算
* 计算出 canJump[i] 后，canJump[i-1] 就可以计算出来，只要根据 nums[i-1] 所能覆盖的范围进行计算即可

## 0056

对 interval 进行合并

* 首先将所有的 inteval 按照左端点排序
* 排序之后，逐个比较：比较某一个 interval 的左端点比结果中最后一个 inteval 的右断点：
    * 如果其左端点更大，则表示是一个新的 interval
    * 否则出现了重合，在比较右端点，计算出新的右端点

## 0057

interval 的插入

* 一种比较简单的方法是，逐个遍历interval：如果当前 interval 在 新的 interval 之前，直接加入
* 之后一直找到和 newInterval 重叠的 inteval，进行合并
* 最后再把剩余的元素复制到结果中

## 0058

最后一个单词的长度

* 通过一个变量来记录当前遍历的字符，是否处于单词中

## 0061

旋转链表

* 找准旋转的规律，直接使用指针操作完成旋转

## 0062

唯一路径数：

* 简单的动态规划方法，抓住 a[i][j] = a[i+1][j] + a[i][j+1]
* 但是需要注意边界条件

## 0063

有障碍的唯一路径：

* 仍然是动态规划思想，只需要注意，从障碍物到目的节点的路径数为 0

## 0064

最小路径和

* 简单的动态规划方法，和 0062 类似
* 同样注意边界条件

## 0066

数组加一：

* 比较简单，唯一注意，需要使用各临时变量计算余数，不要直接赋值到数组中，否则影响后续计算

## 0067

二进制字符串相加

* 类似于大数加法
* 最后要反转结果

## 0069

数字的平方根

* 注意二分搜索的起始条件（从 4 开始二分搜索才是准确的）
* 二分搜索退出时，返回 left or right，需要仔细分析

## 0070

爬楼梯问题：

* 斐波那契额数列问题：要注意 n 为 1 的情况，不要越界访问

## 0073

对矩阵设置为 0:

* 如果想用常数空间，需要复用第一行和第一列作为保存当前行/当前列是否存在 0 值
* 此时还需要两个变量本身来设置第一行/第一列是否存在零值

## 0074

有序矩阵的搜索：

* 仍然是采用二分搜索方法，但是需要注意从 index 到矩阵 pos 的转换

## 0075

颜色排序：

* 采用计数排序的方式

## 0077

组合问题

* 使用回溯方法，思路更清晰，递归结束的条件，当前数组元素达到 k 个

## 0078

组合问题

* 同样选择深度优先搜索+回溯方法，每次第 i 个元素，可以使用，也可以不使用，之后再去选择第 i+1 个元素
* 与排列问题的区别是，排列需要使用一个数组来记录当前哪些元素已经使用

## 0079

矩阵单词搜索

* 递归 + 回溯：从矩阵中的每一个位置开始递归
* 只要当前位置的值和 target 相同，则继续递归其上下左右
* 需要使用一个数组记录是否使用过该位置
* 递归结束的条件，每次递归结束需要恢复相应的信息

## 0081

逆转数组的搜索

* 判断出那个区间是有序的，如果 target 在这个区间，则直接在这个区间搜索，否则就是去另一个区间搜索
* 当元素相等时，判断不出区间，这个时候直接减少一个元素

## 0082

移除链表所有重复元素

* 判断出何时需要删除元素，此时直接删除该元素及之后所有等于该值的元素

## 0086

划分链表

* 直接在原始链表中修改元素，操作比较复杂
* 可以创建出两个新链表，根据当前元素的值，将其加入到对应的链表中
* 最后将两个链表连接起来即可

## 0088

合并有序数组

* 需要注意的是，提前将元素移动到末尾位置，这样就不会在合并时被复制

## 0089

格雷码：

* 抓住这是一个递归问题，求出 n-1 个的 graycode 后，使用结果 result(n-1) | 1<<(n-1) 即为 n 的结果
* 但是需要注意，需要逆序使用 result(n-1) 的结果

## 0090

重复数字子集

* 选择时，需要将之前选过的数排除

## 0091

解码：

* 动态规划问题：需要推导出状态转义方程

## 0092

反转指定元素之间的链表

* 链表操作的巅峰之作
* 需要记录住这两个边界元素，执行对应的操作

## 0093

恢复 IP 地址

* 采用递归解法，每次判断当前的数是否是一个合法的 IP 字节
* 在满足要求后，继续递归判断剩余的字符能够拼接处指定个数的 IP 字节

## 0094

二叉树的中序遍历

* 非递归写法有点难度，需要使用一个栈来保存 left 节点
* 只有所有 left 节点都入栈后，才逐个出栈，开始遍历其右节点

## 0096

独特二叉树

* 使用递归解法，每次选择一个树作为根节点，则剩余的数已经能确定哪些是左子树，哪些是右子树

## 0098

验证二叉搜索数的有效性

* 按照中序遍历二叉树，如果当前元素的值始终大于前一个元素的值，则符合二叉搜索数
* 非递归遍历时，需要使用一个栈，保存当前的左子节点，只有左子节点遍历完了，才开始遍历其有子节点
* 遍历右子节点时，仍然需要保存该值

## 0105

二叉树还原：

* 根据前序遍历、中序遍历就可以找准根节点、左子树的范围、右子树的范围
* 递归还原左子树、右子树

## 0109

链表转换为二叉搜索树

* 首先遍历链表，得到元素数组
* 使用递归解法，将元素数组转换为 BST，这里使用到了二分搜索的思想

## 0114

flatten

* 将二叉树转换为链表，这里直接使用先序遍历二叉树，然后记录下上一个元素
* 设置其 right node

## 0154

可重复数字旋转之后，寻找最小数

* 旋转之后，数组分为两部分，左半部分的数均大于等于右半边的数
* 使用 mid 和右边边界进行比较，如果 mid 大，则 mid 肯定是在左半部分，则分隔点在 mid+1 之后
* 如果 mid 小，则 mid 肯定在右半边部分，则分隔点 mid 之前（可以等于 mid）
* 如果相等，则直接缩小右边界（-1）

## 0232

*两个栈实现队列：

* 用一个栈保存要入队的数据
* 用另一个栈保存要处对的数据
* 当出队的栈为空时，要将入队的栈数据转义到这个栈中


## 0240

有序矩阵搜索：

* 每次比较矩阵 [0, columns-1] 位置的数，从而每次可以去掉一整行或一整列
* 注意退出的条件

## 0509

斐波那契数列：

* 直接用递归解法即可

package leetcode

func merge(nums1 []int, m int, nums2 []int, n int)  {
    copy(nums1[n:], nums1[0:m])

    nums1Src := nums1[n:]
    currPos, i, j := 0,0,0

    for i < len(nums1Src) || j < len(nums2) {
        var target int

        if i >= len(nums1Src) {
            target = nums2[j]
            j++
        } else if j >= len(nums2) {
            target = nums1Src[i]
            i++
        } else {
            if nums1Src[i] < nums2[j] {
                target = nums1Src[i]
                i++
            } else {
                target = nums2[j]
                j++
            }
        }

        nums1[currPos] = target
        currPos++
    }
}

// func findLengthOfLCIS1(nums []int) int {
//     result := 1
//     for i := 0; i < len(nums); i++ {
//         currLength := 1
//         currNum := nums[i]
//         for j := i+1; j < len(nums); j++ {
//             if nums[j] > currNum {
//                 currLength++
//                 currNum = nums[j]
//             }else{
//                 break
//             }
//         }
//         result = max(result, currLength)
//     }
//     return result
// }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLengthOfLCIS(nums []int) int {
	left, right := 0, 0
	result := 1
	for right < len(nums)-1 {
		if nums[right+1] > nums[right] {
			right++
			result = max(result, right-left+1)
		} else {
			left = right + 1
			right++
		}
	}
	return result
}
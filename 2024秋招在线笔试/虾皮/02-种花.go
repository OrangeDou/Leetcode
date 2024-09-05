package main

import "fmt"

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 *
 * @param flowerbed int整型 一维数组
 * @param n int整型
 * @return bool布尔型
 */
func plantFlowers(flowerbed []int, n int) bool {
	// write code here
	count := 0
	if len(flowerbed) <= 0 {
		return false
	}
	for i := 0; i < len(flowerbed)-1; {
		if flowerbed[i] == 1 {
			// 处在中间地段三个0才能种
			if i+1 < len(flowerbed) && i+2 < len(flowerbed) && i+3 < len(flowerbed) {
				if flowerbed[i+1] == 0 && flowerbed[i+2] == 0 && flowerbed[i+3] == 0 {
					count++
					i = i + 3
				} else {
					i = i + 3
					continue
				}
			}
		}
		// 开头和结尾；两个0就可以种
		if flowerbed[i] == 0 && i == 0 && i+1 < len(flowerbed) && flowerbed[i+1] == 0 {
			count++ //第一格
			i = i + 1
		}
		if flowerbed[i] == 0 && i == len(flowerbed)-2 && i+1 < len(flowerbed) && flowerbed[i+1] == 0 {
			count++ //最后一格
			break
		}
		i++
	}
	if count >= n {
		return true
	}
	return false
}

func main() {
	num := []int{1, 0, 0, 0}
	fmt.Print(plantFlowers(num, 1))
}

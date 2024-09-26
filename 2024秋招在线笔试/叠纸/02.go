package main

import (
	"container/list"
)

type Point struct {
	X, Y int
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 需要计算经过的所有路点消耗值之和，终点不可达或找不到返回-1
 * @param map_width int整型 地图的宽度
 * @param map_height int整型 地图的高度
 * @param grids int整型二维数组 网格数据
 * @param start_pos Point类 起始位置
 * @param end_pos Point类 终点位置
 * @return int整型
 */
func FindPath(map_width int, map_height int, grids [][]int, start_pos *Point, end_pos *Point) int {
	// write code here
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	cost := make([][]int, map_height)
	for i := range cost {
		cost[i] = make([]int, map_width)
		for j := range cost[i] {
			cost[i][j] = 999
		}
	}

	queue := list.New()
	queue.PushBack(*start_pos)
	cost[start_pos.Y][start_pos.X] = 0
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(Point)
		if curr.X == end_pos.X && curr.Y == end_pos.Y {
			return cost[curr.Y][curr.X]
		}
		for _, dir := range directions {
			next := Point{curr.X + dir.X, curr.Y + dir.Y}
			if next.X >= 0 && next.X < map_width && next.Y >= 0 && next.Y < map_height {
				nextCost := cost[curr.Y][curr.X] + grids[next.Y][next.X]
				if grids[next.Y][next.X] == 0 || nextCost >= cost[next.Y][next.X] {
					continue
				}
				cost[next.Y][next.X] = nextCost
				queue.PushBack(next)
			}
		}
	}
	return -1
}

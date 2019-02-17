package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if (err != nil) {
		panic(err)
	}
	var row, col int
	// 读取几行几列
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type points struct {
	i, j int
}

// 走一步
func (p points) add(r points) points {
	return points{p.i + r.i, p.j + r.j}
}

// 判断当前脚步在哪里
func (p points) at(maze [][]int) (int, bool) {
	if p.i < 0 || p.j < 0  {
		return 0, false
	}
	if (p.i >= len(maze)) || p.j >= len(maze[p.i]) {
		return 0, false
	}
	return maze[p.i][p.j], true
}

var dirs = [4]points{
	// 向下走一步
	{1,0},
	// 向上走一步
	{-1,0},
	// 向右走一步
	{0,1},
	// 向左走一步
	{0,-1},
}

func walk(maze [][]int, start, end points) [][]int {
	// 定义脚步slice
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []points{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		// 向四周探索
		for _, dir := range dirs {
			// 下一个节点
			next := cur.add(dir)
			// 节点所处位置

			val, ok := next.at(maze)
			if !ok || val == 1 {
				// 走出界或者走到墙
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				// 走出界或者已经走过
				continue
			}
			if next == start {
				// 回掉原点
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			// 进入队列，进行下一次探索
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	// 获取地图放在slice里面
	maze := readMaze("maze/maze.in")

	steps := walk(maze, points{0,0}, points{len(maze) - 1, len(maze[0]) -1 })

	for _, i := range steps {
		for _, j := range i {
			fmt.Printf("%3d", j)
		}
		fmt.Println()
	}
}

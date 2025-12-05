package solution

type Wall struct {
	grid [][]string
}

const maxAccessibleRolls int = 4

func (w Wall) isAccessibleRollsFrom(x, y int) bool {
	count := 0
	// top left
	if y > 0 && x > 0 && w.grid[y-1][x-1] == "@" {
		count++
	}
	// top
	if y > 0 && w.grid[y-1][x] == "@" {
		count++
	}
	// top right
	if y > 0 && x < len(w.grid[0])-1 && w.grid[y-1][x+1] == "@" {
		count++
	}
	// left
	if x > 0 && w.grid[y][x-1] == "@" {
		count++
	}
	// right
	if x < len(w.grid[0])-1 && w.grid[y][x+1] == "@" {
		count++
	}
	// bottom left
	if y < len(w.grid)-1 && x > 0 && w.grid[y+1][x-1] == "@" {
		count++
	}
	// bottom
	if y < len(w.grid)-1 && w.grid[y+1][x] == "@" {
		count++
	}
	// bottom right
	if y < len(w.grid)-1 && x < len(w.grid[0])-1 && w.grid[y+1][x+1] == "@" {
		count++
	}
	return count < maxAccessibleRolls
}

func (w Wall) countAccessibleRolls() int {
	count := 0
	for y, row := range w.grid {
		for x, cell := range row {
			if cell == "@" && w.isAccessibleRollsFrom(x, y){
				count++
			}
		}
	}
	return count
}

func (w *Wall) removeAccessibleRolls() int {
	count := 0
	for y, row := range w.grid {
		for x, cell := range row {
			if cell == "@" && w.isAccessibleRollsFrom(x, y){
				w.grid[y][x] = "x"
				count++
			}
		}
	}
	return count
}

func (w *Wall) removeAllAccessibleRolls() int {
	total := 0
	lastCount := 1
	for lastCount > 0 {
		lastCount = w.removeAccessibleRolls()
		total += lastCount
	}
	return total
}
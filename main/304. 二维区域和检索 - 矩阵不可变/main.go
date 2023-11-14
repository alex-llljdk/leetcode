package main

type NumMatrix struct {
	prex [][]int
}

/*
一维前缀和
*/
func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix))
	for i, row := range matrix {
		sum[i] = make([]int, len(row)+1)
		for j, v := range row {
			sum[i][j+1] = sum[i][j] + v
		}
	}
	return NumMatrix{prex: sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.prex[i][col2+1] - this.prex[i][col1]
	}
	return sum
}

/*
二维前缀和
*/
func Constructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])
	prex := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		prex[i] = make([]int, col+1)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			prex[i][j] = prex[i-1][j] + prex[i][j-1] - prex[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{prex: prex}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.prex[row2+1][col2+1] - this.prex[row1][col2+1] - this.prex[row2+1][col1] + this.prex[row1][col1]
}

func main() {

}

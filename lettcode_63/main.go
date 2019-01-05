package lettcode_63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n == 0 {
		return 0
	}
	m := len(obstacleGrid[0])
	if m == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range obstacleGrid {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1 - obstacleGrid[0][0]

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if obstacleGrid[i][j] == 0 {
				if i > 0 {
					dp[i][j] += dp[i-1][j]
				}
				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}
	return dp[n-1][m-1]

}
func main() {}

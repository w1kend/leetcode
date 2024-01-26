package mindistance

func minDistance(word1 string, word2 string) int {
	n := max(len(word1), len(word2)) + 2
	dp := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
	}

	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 || j == 0 {
				//base case
				dp[i][j] = i + j
				continue
			}

			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j-1]+1, // replace
					dp[i][j-1]+1,   // insert
					dp[i-1][j]+1,   // delete
				)
			}

		}
	}

	return dp[len(word1)][len(word2)]
}

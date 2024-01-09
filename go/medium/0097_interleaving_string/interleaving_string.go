package interleavingstring

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, 0, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp = append(dp, make([]bool, len(s2)+1))
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = true
			case i == 0:
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			case j == 0:
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			default:
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1] || dp[i][j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

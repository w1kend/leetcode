package permitationinstring

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	// store "map" of our letters' frequency
	s1Cnts := [26]int{}
	sliceCnts := [26]int{}
	for i := range s1 {
		s1Cnts[s1[i]-'a']++
		sliceCnts[s2[i]-'a']++
	}

	l, r := 0, len(s1)

	for ; r < len(s2); r++ {
		// if letters and their frequency in the sliceCnts is the same as in s1Cnts
		// then we've found our substring
		if sliceCnts == s1Cnts {
			return true
		}

		// move the left border
		sliceCnts[s2[l]-'a']-- // reduce frequency of the letter at s2[l]
		l++
		sliceCnts[s2[r]-'a']++ // include a letter at the right border
	}

	if sliceCnts == s1Cnts {
		return true
	}

	return false
}

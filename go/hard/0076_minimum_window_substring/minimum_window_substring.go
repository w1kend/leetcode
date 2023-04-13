package minimumwindowsubstring

import (
	"fmt"
)

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	targetStrCharsMap := make(map[byte]int, 52)
	for _, v := range t {
		targetStrCharsMap[byte(v)]++
	}

	strCharsMap := make(map[byte]int, len(targetStrCharsMap))

	l, r := 0, 0 // window borders

	minLength := len(s)
	resL, resR := -1, -1

	haveAllChars := false
	cntChars := 0
	for l < len(s) {
		lChar := s[l]
		lCharCnt := strCharsMap[lChar]

		lCharCntTarget, existInTarget := targetStrCharsMap[lChar]
		if haveAllChars && (lCharCnt-1 >= lCharCntTarget || !existInTarget) {
			//move left border
			if existInTarget {
				lCharCnt--
				strCharsMap[lChar] = lCharCnt
			}
			l++

			wordLength := r - l
			if wordLength <= minLength {
				minLength = wordLength
				resL, resR = l, r
				if minLength == len(t) {
					break
				}
			}

			continue
		}

		if r == len(s) {
			break
		}
		// move right
		rChar := s[r]
		rCharCnt := strCharsMap[rChar]

		rCharCntTarget, rExistInTarget := targetStrCharsMap[rChar]

		if !rExistInTarget {
			r++
			continue
		}

		rCharCnt++
		strCharsMap[rChar] = rCharCnt

		// firstly we should fild a window which have all the required letters
		// then we'll move l and r to find a minimum window that satisfied the same condition
		if !haveAllChars && rCharCntTarget == rCharCnt {
			cntChars++
			if cntChars == len(targetStrCharsMap) {
				haveAllChars = true

				wordLength := r - l + 1
				if wordLength <= minLength {
					minLength = wordLength
					resL, resR = l, r+1
					if minLength == len(t) {
						break
					}
				}
			}
		}
		r++
	}

	if resL == -1 || resR == -1 {
		return ""
	}

	if resL == resR {
		return string(s[resL])
	}

	if resR > len(s) {
		resR = len(s)
	}

	return s[resL:resR]
}

func printM(m map[byte]int) string {
	if len(m) == 0 {
		return "[]"
	}
	s := ""
	for k, v := range m {
		s += fmt.Sprintf("%s:%d,", string(k), v)
	}
	return "[" + s[:len(s)-1] + "]"
}

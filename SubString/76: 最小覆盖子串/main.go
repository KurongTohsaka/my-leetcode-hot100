package main

func main() {

}

func isCovered(cntS, cntT []int) bool {
	for i := 'A'; i <= 'Z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}

func minWindow(s, t string) string {
	ansLeft, ansRight := -1, len(s)
	var cntS, cntT [128]int
	for _, c := range t {
		cntT[c]++
	}

	left := 0
	for right, c := range s { // 移动子串右端点
		cntS[c]++                         // 右端点字母移入子串
		for isCovered(cntS[:], cntT[:]) { // 涵盖
			if right-left < ansRight-ansLeft { // 找到更短的子串
				ansLeft, ansRight = left, right // 记录此时的左右端点
			}
			cntS[s[left]]-- // 左端点字母移出子串
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

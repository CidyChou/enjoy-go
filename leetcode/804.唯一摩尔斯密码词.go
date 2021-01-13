/*
 * @lc app=leetcode.cn id=804 lang=golang
 *
 * [804] 唯一摩尔斯密码词
 */
func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	var uniqueMorse []string
	var s string

	for _, word := range words {
		s = ""
		for _, l := range []rune(word) {
			s = s + morse[int(l)-97]
		}
		uniqueMorse = append(uniqueMorse, s)
	}

	return len(RemoveRepByLoop(uniqueMorse))
}

func RemoveRepByLoop(slc []string) []string {
	result := []string{}
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}
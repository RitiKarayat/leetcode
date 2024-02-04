// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	occ_s := make(map[string]int)

	for _, chr := range s {
		if _, ok := occ_s[string(chr)]; !ok {
			occ_s[string(chr)] = 1
			continue
		}
		occ_s[string(chr)] += 1
	}

	for _, chr := range t {
		if _, ok := occ_s[string(chr)]; !ok || occ_s[string(chr)] == 0 {
			return false
		}
		occ_s[string(chr)] -= 1
	}

	return true

}
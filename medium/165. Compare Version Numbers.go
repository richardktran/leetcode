package main

func compareVersion(version1 string, version2 string) int {
	v1 := 0
	v2 := 0
	i := 0
	j := 0

	for i < len(version1) || j < len(version2) {
		for i < len(version1) && version1[i] != '.' {
			v1 = v1*10 + int(version1[i]-'0')
			i++
		}

		for j < len(version2) && version2[j] != '.' {
			v2 = v2*10 + int(version2[j]-'0')
			j++
		}

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}

		v1 = 0
		v2 = 0
		i++
		j++
	}

	return 0
}

func compareVersionTest() {
	println(compareVersion("1.01", "1.001")) // 0
	println(compareVersion("1.0", "1.0.0"))  // 0
	println(compareVersion("0.1", "1.1"))    // -1
}

// func main() {
// 	compareVersionTest()
// }

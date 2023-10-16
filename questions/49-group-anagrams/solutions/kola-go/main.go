package main

func groupAnagrams(strs []string) [][]string {
    setMap := make(map[[26]int][]string)
    for _, str := range strs {
        var charCount [26]int
        for _, c := range str {
            charCount[c - 'a']++
			// sort 
        }
        setMap[charCount] = append(setMap[charCount], str)
    }
    result := make([][]string, len(setMap))
    idx := 0
    for _, r := range setMap {
        result[idx] = r
        idx++
    }
    return result
}
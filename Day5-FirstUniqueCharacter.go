// Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

// Examples:

// s = "leetcode"
// return 0.

// s = "loveleetcode",
// return 2.
// Note: You may assume the string contain only lowercase letters.


func firstUniqChar(s string) int {
    charCount := make(map[string]int)
    charIndex := make(map[string]int)
    for i, c := range s {
        cS := string(c)
        if _, ok := charCount[cS]; ok {
            charCount[cS] += 1
        }else {
            charIndex[cS] = i
            charCount[cS] += 1
        }
    }
    smallestIndex := len(s) - 1
    isThere := false
    for c, count := range charCount {
        fmt.Print(count)
        if count == 1 && charIndex[c] <= smallestIndex{
            isThere = true
            smallestIndex = charIndex[c]
        }
    }
    if !isThere {
        return -1
    }else {
        return smallestIndex
    }
    
}
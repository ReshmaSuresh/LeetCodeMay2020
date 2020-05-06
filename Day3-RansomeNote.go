// :Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

// Each letter in the magazine string can only be used once in your ransom note.

// Note:
// You may assume that both strings contain only lowercase letters.

// canConstruct("a", "b") -> false
// canConstruct("aa", "ab") -> false
// canConstruct("aa", "aab") -> true

func canConstruct(ransomNote string, magazine string) bool {
    magazineCharacters := make(map[string]int)
    for _, character := range magazine {
        magazineCharacters[string(character)] += 1
    }
    
    for _, ransomeChar:= range ransomNote {
        if val, ok := magazineCharacters[string(ransomeChar)]; ok {
            if val == 0 {
                return false
            }
            magazineCharacters[string(ransomeChar)] -= 1   
        } else {
            return false
        }
    }
    return true
}
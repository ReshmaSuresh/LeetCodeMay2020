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
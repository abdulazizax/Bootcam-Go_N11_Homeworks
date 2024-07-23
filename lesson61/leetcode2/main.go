package main

// url := https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
    isAlphanumeric := func(c byte) bool {
        return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
    }
    
    var cleaned []byte
    for i := 0; i < len(s); i++ {
        if isAlphanumeric(s[i]) {
            if s[i] >= 'A' && s[i] <= 'Z' {
                cleaned = append(cleaned, s[i]+'a'-'A')
            } else {
                cleaned = append(cleaned, s[i])
            }
        }
    }
    
    for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
        if cleaned[i] != cleaned[j] {
            return false
        }
    }
    
    return true
}

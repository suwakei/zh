package main

import (
    "github.com/suwakei/go-zhconv"
)


func Z2h(str string) string {
    var result string
    for _, char := range str {
        if idx := indexRune(ASCII_ZENKAKU_CHARS, char); idx != -1 {
            result += string(ASCII_HANKAKU_CHARS[idx])
        } else if idx := indexRune(KANA_ZENKAKU_CHARS, char); idx != -1 {
            result += string(KANA_HANKAKU_CHARS[idx])
        } else if idx := indexRune(DIGIT_ZENKAKU_CHARS, char); idx != -1 {
            result += string(DIGIT_HANKAKU_CHARS[idx])
        } else if newChar, ok := KANA_TEN_MAP[char]; ok {
            result += string(newChar) + "ﾞ"
        } else if newChar, ok := KANA_MARU_MAP[char]; ok {
            result += string(newChar) + "ﾟ"
        } else {
            result += string(char)
        }
    }
    return result
}

func indexRune(s string, r rune) int {
    for i, char := range s {
        if char == r {
            return i
        }
    }
    return -1
}
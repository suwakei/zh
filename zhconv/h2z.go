package main

import (
    "github.com/suwakei/go-zhconv/"
)

// strが空の場合や、無効な文字が含まれている場合にどう処理するか
// そういったエラーハンドリングを書く
// パフォーマンスの改善のためstringbuilder使う+=のところ

func  H2z(str string) string {
    var result string
    for i := 0; i < len(str); i++ {
        char := rune(str[i])
        if idx := indexRune(ASCII_HANKAKU_CHARS, char); idx != -1 {
            result += string(ASCII_ZENKAKU_CHARS[idx])
        } else if idx := indexRune(KANA_HANKAKU_CHARS, char); idx != -1 {
            result += string(KANA_ZENKAKU_CHARS[idx])
        } else if idx := indexRune(DIGIT_HANKAKU_CHARS, char); idx != -1 {
            result += string(DIGIT_ZENKAKU_CHARS[idx])
        } else if i+1 < len(str) && str[i+1] == 'ﾞ' && KANA_TEN_MAP[char] != 0 {
            result += string(KANA_TEN_MAP[char])
            i++ // Skip the next character
        } else if i+1 < len(str) && str[i+1] == 'ﾟ' && KANA_MARU_MAP[char] != 0 {
            result += string(KANA_MARU_MAP[char])
            i++ // Skip the next character
        } else {
            result += string(char)
        }
    }
    return result
}
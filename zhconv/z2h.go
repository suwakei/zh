package zhconv

import (
    "github.com/suwakei/go-zhconv/tables"
)

// strが空の場合や、無効な文字が含まれている場合にどう処理するか
// そういったエラーハンドリングを書く
// パフォーマンスの改善のためstringbuilder使う+=のところ

func Z2h(str string) string {
    if str == "" {
        return ""
    }
    var result string
    var t *tables.ConversionTables = tables.New()
    for _, char := range str {
        if idx := indexRune(string(t.ASCII_ZENKAKU_CHARS), char); idx != -1 {
            result += string(string(t.ASCII_HANKAKU_CHARS[idx]))
        } else if idx := indexRune(string(t.KANA_ZENKAKU_CHARS), char); idx != -1 {
            result += string(string(t.KANA_HANKAKU_CHARS[idx]))
        } else if idx := indexRune(string(t.DIGIT_ZENKAKU_CHARS), char); idx != -1 {
            result += string(string(t.DIGIT_HANKAKU_CHARS[idx]))
        } else if newChar, ok := t.KANA_TEN_MAP[char]; ok {
            result += string(newChar) + "ﾞ"
        } else if newChar, ok := t.KANA_MARU_MAP[char]; ok {
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
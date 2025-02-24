package zhconv

import (
    "github.com/suwakei/go-zhconv/tables"
)

// strが空の場合や、無効な文字が含まれている場合にどう処理するか
// そういったエラーハンドリングを書く
// パフォーマンスの改善のためstringbuilder使う+=のところ

func  H2z(str string) string {
    if str == "" {
        return ""
    }
    var result string
    var t *tables.ConversionTables = tables.New()
    for i := 0; i < len(str); i++ {
        char := rune(str[i])
        if idx := indexRune(string(t.ASCII_HANKAKU_CHARS), char); idx != -1 {
            result += string(t.ASCII_ZENKAKU_CHARS[idx])
        } else if idx := indexRune(string(t.KANA_HANKAKU_CHARS), char); idx != -1 {
            result += string(t.KANA_ZENKAKU_CHARS[idx])
        } else if idx := indexRune(string(t.DIGIT_HANKAKU_CHARS), char); idx != -1 {
            result += string(t.DIGIT_ZENKAKU_CHARS[idx])
        } else if i+1 < len(str) && str[i+1] == 'ﾞ' && string(t.KANA_TEN_MAP[char]) != 0 {
            result += string(t.KANA_TEN_MAP[char])
            i++ // Skip the next character
        } else if i+1 < len(str) && str[i+1] == 'ﾟ' && string(t.KANA_MARU_MAP[char]) != 0 {
            result += string(t.KANA_MARU_MAP[char])
            i++ // Skip the next character
        } else {
            result += string(char)
        }
    }
    return result
}
package util

import (
	"golang.org/x/text/language"
)

// 言語対応しており、受け付ける言語を列挙
func NewMatcher() language.Matcher {
	return language.NewMatcher([]language.Tag{
		language.Japanese, // デフォルトとしてマッチする。
		language.English,
	})
}

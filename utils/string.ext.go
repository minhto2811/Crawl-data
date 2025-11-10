package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func RemoveDiacritics(str string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, str)
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}


func ToSnakeCase(s string) string {
	// 1. Chuyển sang lowercase
	s = strings.ToLower(s)

	// 2. Thay dấu tiếng Việt
	s = removeVietnameseTones(s)

	// 3. Thay khoảng trắng và các ký tự không phải chữ/số bằng dấu _
	re := regexp.MustCompile(`[^\w]+`)
	s = re.ReplaceAllString(s, "_")

	// 4. Xóa các _ thừa đầu/cuối
	s = strings.Trim(s, "_")

	return s
}

// Hàm loại bỏ dấu tiếng Việt
func removeVietnameseTones(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case 'á', 'à', 'ả', 'ã', 'ạ', 'ă', 'ắ', 'ằ', 'ẳ', 'ẵ', 'ặ', 'â', 'ấ', 'ầ', 'ẩ', 'ẫ', 'ậ':
			b.WriteRune('a')
		case 'é', 'è', 'ẻ', 'ẽ', 'ẹ', 'ê', 'ế', 'ề', 'ể', 'ễ', 'ệ':
			b.WriteRune('e')
		case 'í', 'ì', 'ỉ', 'ĩ', 'ị':
			b.WriteRune('i')
		case 'ó', 'ò', 'ỏ', 'õ', 'ọ', 'ô', 'ố', 'ồ', 'ổ', 'ỗ', 'ộ', 'ơ', 'ớ', 'ờ', 'ở', 'ỡ', 'ợ':
			b.WriteRune('o')
		case 'ú', 'ù', 'ủ', 'ũ', 'ụ', 'ư', 'ứ', 'ừ', 'ử', 'ữ', 'ự':
			b.WriteRune('u')
		case 'ý', 'ỳ', 'ỷ', 'ỹ', 'ỵ':
			b.WriteRune('y')
		case 'đ':
			b.WriteRune('d')
		default:
			if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_' {
				b.WriteRune(r)
			} else {
				b.WriteRune(' ')
			}
		}
	}
	return b.String()
}
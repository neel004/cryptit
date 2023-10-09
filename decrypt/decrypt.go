package decrypt

func Luffy(enc_str string) string {
	dec_str := ""
	for _, char := range enc_str {
		ascci := int(char)
		dec_str += string(ascci - 3)
	}
	return dec_str
}

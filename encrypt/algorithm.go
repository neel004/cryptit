package encrypt

func Nimbus(str string) string {
	enc_string := ""

	for _, char := range str {
		ascci_code := int(char)
		charr := string(ascci_code + 3)
		enc_string += charr
	}

	return enc_string
}

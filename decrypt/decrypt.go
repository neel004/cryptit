// Decrypt Package provides utilities to decrypt strings.
package decrypt

// Oreva Luffy, Monkey D. Luffy. I Decrypt the strings.
func Luffy(enc_str string) string {
	dec_str := ""
	for _, char := range enc_str {
		ascci := int(char)
		dec_str += string(ascci - 3)
	}
	return dec_str
}

package iterations

func Repeated(alpha string, rep int) string {
	result := ""
	for i := 0; i < rep; i++ {
		result += alpha
	}
	return result
}

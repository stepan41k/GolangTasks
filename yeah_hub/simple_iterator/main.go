// Доступно без импорта: sort, fmt, encoding/json
package main
func CreateCharReader(s string) func() *string {
	i := 0
	
	return func() *string {
		if i == len(s) {
			return nil
		}
		
		x := string(s[i])
		i++
		return &x
	}
}

func main() {
	reader := CreateCharReader("hi")
	println(reader())
	println(reader())
	println(reader())
}
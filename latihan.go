package main
import "fmt"

func faktorialGanjil(n int) int{
	if n == 0{
		return 1
	}
	if n%2 != 0 {
		return n * faktorialGanjil(n-1)
	}
	return faktorialGanjil(n-1)
}

func main(){
	var masukan int
	fmt.Scan(&masukan)
	fmt.Print(faktorialGanjil(masukan))
}
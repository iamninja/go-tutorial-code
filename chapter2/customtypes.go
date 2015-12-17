package main

import "fmt"

// const (
// 	KB = 1024
// 	MB = KB * 1024
// 	GB = MB * 1024
// 	TB = GB * 1024
// 	PB = TB * 1024
// )
// Define constants in a better way using iota expression
// const (
// 	one int = iota + 1	// one = 1
// 	two 				// two = 2
// 	three 				// three = 3
// 	// ...
// )
const (
	KB ByteSize = 1<<(10*(iota+1))	// A<<B = (2*A)^B
	MB
	GB
	TB
	PB
)

type ByteSize float64

func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return "Very Big"
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b>= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%dB", b)
}

func main() {
	fmt.Println(ByteSize(2048))			// 2.00KB. We don't use string method. fmt package will use String method if it exists
	fmt.Println(ByteSize(3256879.55))	// 3.11MB
	b := ByteSize(345)
	fmt.Println(b.String())				// If we use it we do not get what is expected
}
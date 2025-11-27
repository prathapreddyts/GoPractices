package underStandingBasics

import "fmt"

func SquarePatterns(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func RightAngledTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func InvertedRightAngledTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func PyramidPattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func DiamondPattern(n int) {
	// Top half
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// Bottom half
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func FloydsTriangle(n int) {
	var sum int = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(sum, " ")
			sum++
		}
		fmt.Println()
	}
}
func PascalTriangle(n int) {

}
func NumberPyramid(n int) {
	s := 1
	for i := 0; i < n; i++ {
		for s := 0; s < n-i-1; s++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print(s, " ")
			s++
		}
		fmt.Println()
	}
}
func InvertedNumberPyramid(n int) {
	s := 2*n - 1
	for i := 0; i < n; i++ {
		for space := 0; space < i; space++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*(n-i)-1; j++ {
			fmt.Print(s, " ")
			s--
		}
		fmt.Println()
	}
}

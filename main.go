package main

import "fmt"

func main() {
	sol := Solve(puzzles["test-3"])
	fmt.Printf("The final solution:\n")
	PrintSolution(sol)
}

func PrintSolution(sol Solution) {
	for r := 0; r < sol.RowSize; r++ {
		for c := 0; c < sol.ColSize; c++ {
			switch sol.GetGrid(r, c) {
			case GridTypeEnum.Colored:
				fmt.Print("@ ")
			case GridTypeEnum.Blank:
				fmt.Print("X ")
			case GridTypeEnum.Unknown:
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
}

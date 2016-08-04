package test

import (
	"fmt"
	"go-chess/constants"
)

// BoardTesting tests board structure
func BoardTesting() {
	fmt.Println("Board structures:")
	for index := 0; index < constants.BrdSqNum; index++ {
		if index%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", constants.Sq120ToSq64[index])
	}
	fmt.Println()
	fmt.Println()
	for index := 0; index < 64; index++ {
		if index%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", constants.Sq64ToSq120[index])
	}
}

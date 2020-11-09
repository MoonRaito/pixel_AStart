package moon

import "fmt"

func InitMaps() [10][8]string {
	maps := [10][8]string{}

	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			maps[i][j] = "1"
			fmt.Print(maps[i][j] + " ")
		}
		fmt.Println()
	}

	return maps
}

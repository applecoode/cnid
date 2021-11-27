package cnid

import (
	"fmt"
	"log"
	"strconv"
)


var aW = []int{7,9,10,5,8,4,2,1,6,3,7,9,10,5,8,4,2}
var aA = []string{"1","0","X","9","8","7","6","5","4","3","2"}

func Id17to18(id string) string {
	sum := 0
	for i := 0; i < 17; i++ {
		num, err := strconv.Atoi(string(id[i]))
		if err != nil {
			log.Fatal("erro")
		}
		sum += aW[i]*num
	}
	tail := aA[sum%11]
	return fmt.Sprintf("%s%s",id,tail)
}

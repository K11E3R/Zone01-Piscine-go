package piscine

import (
	"strconv"
)

func BasicAtoi2(s string) int {
	res := 1
	pow := len(s)-1
	for i:=0; i<len(s);i++ {
		if(!('0'<=s[i]  && s[i] <= '9')){
			return 0
		}
		res += s[i]*(pow--)
	}

	return ;	
}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var result big.Int
	num := big.NewInt(30)
	resultptr := factorial(&result, num)
	fmt.Println(uint64((*resultptr).Int64()))
}

func factorial(z *big.Int, num *big.Int) (result *big.Int) {
	if (*num).Int64() > 0 {
		temp := big.NewInt((*num).Int64() - 1)
		z = z.Mul(num, factorial(z, temp))
		return z

	} else {
		return big.NewInt(1)
	}
}

//96821 65104 86229 8112

//92233 72036 85477 5807

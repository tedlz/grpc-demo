package main

import (
	"fmt"

	transfer "grpc-demo/99-other/00001-transfer"
	check "grpc-demo/99-other/00002-check"
)

func main() {
	idCard := "410303199912121212"
	fmt.Println(transfer.GetBornInfo(idCard))
	fmt.Println(check.IdCard(idCard))
}

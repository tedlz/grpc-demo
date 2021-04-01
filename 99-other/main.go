package main

import (
	"fmt"

	transfer "grpc-demo/99-other/00001-transfer"
)

func main() {
	idCard := "110305200104011036"
	fmt.Println(transfer.GetBornInfo(idCard))
}

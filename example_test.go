package depsy

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleExtract() {
	data, err := os.ReadFile("go.mod")

	if err != nil {
		panic(err.Error())
	}

	deps := Extract(data, true)

	for _, dep := range deps {
		fmt.Println(dep)
	}
}

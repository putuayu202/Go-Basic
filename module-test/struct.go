package main

import (
	"fmt"
)

type warga struct {
	depan    string
	belakang string
	age      int
	dokumen
}

type dokumen struct {
	sim bool
	ktp bool
}

func main() {
	w1 := warga{
		depan:    "dwiky",
		belakang: "danov",
		age:      18,
		dokumen: dokumen{
			sim: true,
			ktp: true,
		},
	}

	w2 := warga{
		depan:    "ehang",
		belakang: "coin",
		age:      18,
		dokumen: dokumen{
			ktp: true,
		},
	}

	fmt.Println(w1)
	fmt.Println(w2)
}

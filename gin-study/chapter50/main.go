package main

import "fmt"

type Goods struct {
	Name  string
	Price string
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	tv := TV{
		Goods: Goods{
			"电视机", "10000",
		},
		Brand: Brand{
			"xiapu", "shanghai",
		},
	}
	fmt.Println(tv)
	//

	mygood := &Goods{
		Name:  "YANG",
		Price: "10000",
	}
	fmt.Println(mygood)

	tv2 := TV2{
		Goods: &Goods{
			Name:  "YANG",
			Price: "10000",
		},
		Brand: nil,
	}

	fmt.Println(tv2.Goods)
	//fmt.Println(*tv2.Goods)
}

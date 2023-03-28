package main

type cat struct {
	name string
}

//}
//func functionDemo6() cat{
//	returnCat := &cat{}
//	return returnCat
//}

func swap(x, y *int) {
	*x, *y = *y, *x
}
func help() cat {
	var myCat cat
	myCat.name = "70pcice"
	return myCat
}
func t() {

}

package teststring

// CustomString struct
//
type CustomString struct {
	Ps string
	Pb string
}

// ShowString return CustomString
//
func ShowString() CustomString {
	s := "123"

	ps := &s

	b := []byte(*ps)

	pb := &b

	s += "4"

	*ps += "5"

	b[1] = '0'

	//fmt.Println(*ps)

	//fmt.Println(string(*pb))

	return CustomString{Ps: *ps, Pb: string(*pb)}
}

package pkg1

type inputLst = []int
type inputChan = chan int

func WorkAll[I interface{ []int | chan int }](
	_ I,
	_ chan bool,
) error {
	return nil
}

func WorkSync[I interface{ []int | chan int }](
	_ I,
) ([]bool, error) {
	return []bool{}, nil
}

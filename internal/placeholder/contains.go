package placeholder

func Contains(plc []Placeholder, elem [5]string) bool {
	for _, i := range plc {
		if elem == i {
			return true
		}
	}
	return false
}

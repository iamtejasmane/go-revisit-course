package conversion

import "strconv"

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, str := range strings {
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, val)
	}
	return floats, nil
}

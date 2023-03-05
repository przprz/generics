package functions

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func FloatMin(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

// won't compile:
// functions/functions.go:7:5: invalid operation: a < b (operator < not defined on interface)
//func InterfaceMin(a, b interface{}) (interface{}, error) {
//	switch a.(type) {
//	case int:
//		if a < b {
//			return a, nil
//		} else {
//			return b, nil
//		}
//	}
//	return nil, nil
//}

type comparable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func MinGeneric[T comparable](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

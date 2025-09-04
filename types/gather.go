// Package types @author: Violet-Eva @date  : 2025/9/3 @notes :
package types

type T interface {
	string | int | int8 | int16 | int32 | int64 | float32 | float64 | byte | Accesses
}

func Difference[t T](a []t, b ...t) []t {
	bMap := make(map[t]bool)
	for _, v := range b {
		bMap[v] = true
	}
	var diff []t
	for _, v := range a {
		if _, ok := bMap[v]; !ok {
			diff = append(diff, v)
		}
	}
	return diff
}

func Union[t T](a []t, b ...t) []t {
	unionMap := make(map[t]bool)

	for _, v := range a {
		unionMap[v] = true
	}

	for _, v := range b {
		unionMap[v] = true
	}

	union := make([]t, 0, len(unionMap))
	for v := range unionMap {
		union = append(union, v)
	}

	return union
}

func Complement[t T](a []t, b ...t) []t {
	setA := make(map[t]bool)
	setB := make(map[t]bool)

	for _, v := range a {
		setA[v] = true
	}

	for _, v := range b {
		setB[v] = true
	}

	var result []t

	for v := range setA {
		if !setB[v] {
			result = append(result, v)
		}
	}

	for v := range setB {
		if !setA[v] {
			result = append(result, v)
		}
	}

	return result
}

func Intersection[t T](a []t, b ...t) []t {
	aMap := make(map[t]bool)
	for _, v := range a {
		aMap[v] = true
	}

	resultMap := make(map[t]bool)
	for _, v := range b {
		if aMap[v] {
			resultMap[v] = true
		}
	}

	intersection := make([]t, 0, len(resultMap))
	for v := range resultMap {
		intersection = append(intersection, v)
	}

	return intersection
}

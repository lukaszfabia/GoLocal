package functools

func In[T any](toFind T, collection []T) bool {
	matched := false
	i := 0
	for !matched && i < len(collection) {
		matched = any(toFind) == any(collection[i])
		i++
	}

	return matched
}

func Filter[T any](predicate func(e T) bool, collection []T) []T {
	var lst = []T{}
	for _, e := range collection {
		if predicate(e) {
			lst = append(lst, e)
		}
	}

	return lst
}

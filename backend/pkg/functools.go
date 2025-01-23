package pkg

func In[T any](toFind T, collection []T) bool {
	matched := false
	i := 0
	for !matched && i < len(collection) {
		matched = any(toFind) == any(collection[i])
		i++
	}

	return matched
}

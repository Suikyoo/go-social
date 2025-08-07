package sliceutils

func Fill[T any] (s []T, i T) {
  for idx := range s {
    s[idx] = i
  }
}


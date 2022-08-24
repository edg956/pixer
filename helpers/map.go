package helpers

func CopyMap[K comparable, V comparable](source *map[K]V) *map[K]V {
	target := make(map[K]V)
	for k, v := range *source {
		target[k] = v
	}
	return &target
}

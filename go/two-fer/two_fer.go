package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {

	result := "One for "
	if name == "" {
		result = result + "you"
	} else {
		result = result + name
	}
	result = result + ", one for me."

	return result
}

package fs

// The function Exists checks to see if the file or directory exists.
func Exists(path string) bool {
	_, exists := vfs[path]
	return exists
}

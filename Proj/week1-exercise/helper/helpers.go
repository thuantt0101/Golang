package helper

/*
IsEmpty: return true if object is empty
*/
func IsEmpty(obj interface{}) bool {
	if obj == nil || obj == "" {
		return true
	}
	return false
}

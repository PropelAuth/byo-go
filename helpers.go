package byo

// Helper functions for working with pointers

// String returns a pointer to the given string
func String(s string) *string {
	return &s
}

// Int returns a pointer to the given int
func Int(i int) *int {
	return &i
}

// Int64 returns a pointer to the given int64
func Int64(i int64) *int64 {
	return &i
}

// Bool returns a pointer to the given bool
func Bool(b bool) *bool {
	return &b
}

// DerefString returns the value of the string pointer or the default value if nil
func DerefString(s *string, defaultVal string) string {
	if s == nil {
		return defaultVal
	}
	return *s
}

// DerefInt returns the value of the int pointer or the default value if nil
func DerefInt(i *int, defaultVal int) int {
	if i == nil {
		return defaultVal
	}
	return *i
}

// DerefInt64 returns the value of the int64 pointer or the default value if nil
func DerefInt64(i *int64, defaultVal int64) int64 {
	if i == nil {
		return defaultVal
	}
	return *i
}

// DerefBool returns the value of the bool pointer or the default value if nil
func DerefBool(b *bool, defaultVal bool) bool {
	if b == nil {
		return defaultVal
	}
	return *b
}

package boolz

// Ptr returns a pointer to the value.
func Ptr(v bool) *bool {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if false.
func PtrZeroToNil(v bool) *bool {
	if !v {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v bool, def bool) *bool {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to false if nil.
func Val(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *bool, def bool) bool {
	if v == nil {
		return def
	}
	return *v
}

// SafeIndex returns "s[i]" if possible, and false otherwise.
func SafeIndex(s []bool, i int) bool {
	if s == nil || i < 0 || i >= len(s) {
		return false
	}
	return s[i]
}

// SafeIndexDef returns "s[i]" if possible, and "def" otherwise.
func SafeIndexDef(s []bool, i int, def bool) bool {
	if s == nil || i < 0 || i >= len(s) {
		return def
	}
	return s[i]
}

// SafeIndexPtr returns "s[i]" if possible, and nil otherwise.
func SafeIndexPtr(s []bool, i int) *bool {
	if s == nil || i < 0 || i >= len(s) {
		return nil
	}
	return Ptr(s[i])
}

package enums

// Frequency describes the frequencies enum.
type Frequency string

// String implements the fmt.Stringer interface.
func (v Frequency) String() string {
	return string(v)
}

// Valid validates the enum.
func (v Frequency) Valid() bool {
	switch v {
	case
		FrequencySingle,
		FrequencyMulti:
		return true
	default:
		return false
	}
}

// Known enum values.
const (
	FrequencySingle Frequency = "single"
	FrequencyMulti  Frequency = "multi"
)

// Frequencies lists the known enum values.
var (
	Frequencies = []Frequency{
		FrequencySingle,
		FrequencyMulti,
	}
)

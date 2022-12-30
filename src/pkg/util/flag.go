package util

type (
	FlagVar struct {
		isBool bool
		isSet  bool
		value  string
	}
)

// Interface
func (f *FlagVar) Set(x string) error {
	f.value = x
	f.isSet = true
	return nil
}

func (f *FlagVar) String() string {
	return f.value
}

func (f *FlagVar) IsBoolFlag() bool {
	return f.isBool
}

// Helpers
func (f *FlagVar) IsSet() bool {
	return f.isSet
}

func (f *FlagVar) SetBoolFlag() {
	f.isBool = true
}

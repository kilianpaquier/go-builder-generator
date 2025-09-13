package testdata

type InvalidTag struct {
	Flag string `builder:'bla'"`
}

type InvalidOption struct {
	SimpleOption string `builder:"invalid_option"`
	EqualOption  string `builder:"equal_option=prop"`
}

type unexported struct {
	unexported bool
}

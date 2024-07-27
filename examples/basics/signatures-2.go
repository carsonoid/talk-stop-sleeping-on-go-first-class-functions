package main

// START OMIT
// func(
//    func() string,
//    func(string) error
// ) error
func getAndUse1(
	get func() string,
	use func(string) error,
) error {
	return use(get())
}

// func(func() string, func(string) error) error
func getAndUse2(get func() string, use func(string) error) error {
	return use(get())
}

// END OMIT

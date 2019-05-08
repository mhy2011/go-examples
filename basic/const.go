package main

// 常量示例
func main() {
	const ONE_THOUSAND int32 = 1000
	const (
		SUANDAY = iota
		MONDAY
		TUESDAY
		WEDENSDAY
		THURSDAY
		FRIDAY
		SATURDAY
	)

	println("monday =", MONDAY)
	println("thousand =", ONE_THOUSAND)
}

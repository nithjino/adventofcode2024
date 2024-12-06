package main

func IgnoreError[T any](val T, err error) T {
	return val
}

func main() {

}

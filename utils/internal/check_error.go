package internal

func CheckError(err error, message string) {
	if err != nil {
		panic(err)
	}
}

package crema

func HandleError(err error) {
	if err != nil {
		PrintfError(err.Error())
		panic(err)
	}
}

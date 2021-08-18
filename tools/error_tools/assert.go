package error_tools

func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

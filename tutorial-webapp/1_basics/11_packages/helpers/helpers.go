package helpers

type AStruct struct { // the name ahs to be capital camel to get sued externally
	AnInt   int
	AString string
}

func AddOne(num int) int { // captial cmale case will make it public
	return num + 10
}

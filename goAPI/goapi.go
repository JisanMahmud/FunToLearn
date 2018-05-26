package goapi

//SurroundWithQuote ..
func SurroundWithQuote(originTxt string) string {
	return "\"" + originTxt + "\""
}

//AddJSONFieldAndData ..
func AddJSONFieldAndData(Fieldname, Value string) string {
	return SurroundWithQuote(Fieldname) + ":" + SurroundWithQuote(Value)
}

//SurroundWithChars ..
func SurroundWithChars(LeftChar, RightChar, originTxt string) string {
	return LeftChar + originTxt + RightChar
}

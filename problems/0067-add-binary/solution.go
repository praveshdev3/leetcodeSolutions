func addBinary(a string, b string) string {
    aInt,bInt,sum:=new(big.Int),new(big.Int),new(big.Int)
	aInt.SetString(a,2)
	bInt.SetString(b,2)
	sum.Add(aInt,bInt)
	return sum.Text(2)
}

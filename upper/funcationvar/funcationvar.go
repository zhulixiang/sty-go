package funcationvar

func GetSum(n1 int ,n2 int) int  {
	return n1 + n2
}

func MyFun(funName func (int,int) int,n1 int,n2 int)int{
	return funName(n1,n2)
}

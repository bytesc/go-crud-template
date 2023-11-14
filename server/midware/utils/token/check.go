package token

func CheckHS(signature string) error {
	myClaims := UserClaims{}
	err := hs.Decode(signature, &myClaims)
	//fmt.Println(myClaims, err)
	return err
}

//signature, _ := token.IssueHS("hello")
//fmt.Println("签名内容",signature)
//err := token.CheckHS(signature)
//fmt.Println("验签",err)

func CheckRS(signature string) error {
	myClaims := UserClaims{}
	err := rs.Decode(signature, &myClaims)
	//fmt.Println(myClaims, err)
	return err
}

//signature,_ := token.IssueRS("helloword")
//fmt.Println("签名内容",signature)
//err := token.CheckRS(signature)
//fmt.Println("验签",err)

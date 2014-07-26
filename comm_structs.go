package helper

//import (
//	"log"
//)

type UrlToActionSM struct {
	/**
	UrlType
	// 0 common, 1 fixed , 2 prefix fixed and only myid
		3: node detail page url

	// if UrlType > 0 then NotUseActionTemplate is true
	//NotUseActionTemplate bool


	BizType:
		0, default
		1, node

	*/
	TmpltKey      string
	Action        string
	Myid          string
	RenderMapPath string
	Role          string
	ModuleDirPath string
	ErrorMsg      string
	IsDbTmplt     bool
	DbTName       string
	DbTTimestamp  string
	Prefix        string
	UrlType       int
	BizType       int
	RequestURI    string
	FilePath      string
}

type UniformField struct {
	EnName               string
	CnName               string
	HiddenEnFieldStr     string
	HiddenCnFieldStr     string
	DefaultInputFieldStr string
}

type CurdResult struct {
	Error   error
	IsOk    bool
	Message string
}

type ViewResult struct {
	/**
	  TransferObject
	*/
	IsOk      bool
	Error     error
	Type      string      // ->  array/ pointer /  pointer to array / map  / etc
	IMObjType string      // -> one stuct , or map  or   array  or pointer to obj , or obj array etc.
	IMObj     interface{} //->  the quesiton is : how we could convert IObj  to the real struct of where it come from?
	//by using  IObjType ....
}

package main 
import "fmt"
import "os"



func main(){
// 	demo1()
	// otherPackageImplictStruct()
	// otherEmbeddedFiled()
	// sameName()
	sameName2()
}
func sameName2(){
	type File struct{
		name string 
	}
	type Log struct{
		name string 
	}
	type Data struct{
		File 
		Log 
	}
	d := Data{}
	d.File.name = "file"
	d.Log.name = "log"
	fmt.Println(d)
}

func sameName(){
	type File struct{
		name []byte 
	}
	type Data struct{
		File 
		name string //和 File.name 重名
	}

	d := Data{
		name: "data",
		File: File{[]byte("file")},
	}
	d.name="data2"
	d.File.name = []byte("file2")
	fmt.Println(d.name, d.File.name)
}

func otherEmbeddedFiled(){
	type data struct{
		*int // int: *int 
		fmt.Stringer 
		// *fmt.Stringer //embedded field type cannot be a pointer to an interface
	}
	d := data{
		int: nil,
		Stringer: nil,
	}
	fmt.Println(d)
}

func otherPackageImplictStruct(){
	type data struct{
		os.File  //eq File: os.File
	}
	d := data{
		File: os.File{},
	}
	fmt.Printf("%#v\n", d)

}

func demo1(){
	type Attr struct{
		perm int 
	}
	
	type File struct{
		Attr //eq Attr: Attr
		name string 
	}
	f := File{
		name : "test.dat",
		Attr: Attr{
			perm: 0644,
		},
	}
	fmt.Printf("%s, %o\n", f.name, f.perm)
}
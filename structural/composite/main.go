// 组合模式 composite pattern.
// 用于表示树形的结构，这里以一个web静态目录为例
package main

type File struct {
	IsDir bool
	Name string
	ChildFile []*File
}

func (f *File)AddChild(file ...*File){
	f.ChildFile = append(f.ChildFile ,file...)
}

func checkFile(file *File) {
	println("into dir ->" ,file.Name)
	for _ ,v := range file.ChildFile{
		if v.IsDir{
			checkFile(v)
		}else{
			println("dir ->" ,file.Name ,".fileName ->" ,v.Name)
		}
	}
}

func main(){
	/*
	static|
		-js|
			-jquery.js
			-main|
				-index.js
				-login.js
		-css|
			-bootstrap.css
	 */
	static := &File{true ,"static" ,make([]*File ,0)}
	js := &File{true ,"js" ,make([]*File ,0)}
	css := &File{true ,"css" ,make([]*File ,0)}
	static.AddChild(js ,css)

	jquery := &File{false ,"jquery.js" ,nil}
	mjs := &File{true ,"main" ,make([]*File ,0)}
	js.AddChild(jquery ,mjs)

	injs := &File{false ,"index.js" ,nil}
	lojs := &File{false ,"login.js" ,nil}
	mjs.AddChild(injs ,lojs)


	btstrap := &File{false ,"bootstrap.css" ,nil}
	css.AddChild(btstrap)

	checkFile(static)
	/*
	output:
	into dir -> static
	into dir -> js
	dir -> js .fileName -> jquery.js
	into dir -> main
	dir -> main .fileName -> index.js
	dir -> main .fileName -> login.js
	into dir -> css
	dir -> css .fileName -> bootstrap.css
	 */
}



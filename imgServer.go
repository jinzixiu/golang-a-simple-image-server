package main
import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"fmt"
	"io"
	"crypto/md5"
	"strconv"
	"time"
	"github.com/satori/go.uuid"
	_"reflect"
	_"github.com/toolkits/file"
	"encoding/json"
)


const(
	adomin="localhost:9090"
	toolsdir="/static/tools/"
	devdir="/static/dev/"
	livedir="/static/live/"
	identify ="/static/identify/"
)


var realPath *string


type RelCode struct{
	Code int `json:"code"`
	Data string `json:"data"`
	Msg string `json:"msg"`
}



func main() {
	
	realPath = flag.String("path", ".", "static resource path")
	flag.Parse()
	
	
	http.HandleFunc("/", staticResource)
	http.HandleFunc("/tools/upload",vcatools_upload)
	http.HandleFunc("/dev/upload",dev_upload)
	http.HandleFunc("/live/upload",live_upload)
	http.HandleFunc("/help",help)
	
	
	makestaticDir(toolsdir)
	makestaticDir(devdir)
	makestaticDir(livedir)
	makestaticDir(identify)
	
	fmt.Println("将运行在："+adomin)
	err := http.ListenAndServe(adomin, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
		//return
	}
	
	
}



func help(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"sdfasdfasdfasdfasfasdf")
}

func vcatools_upload(w http.ResponseWriter, r *http.Request){
	
	// 1.表单中增加enctype="multipart/form-data"
	//fmt.Println("method:",r.Method)
	
	if r.Method == "GET" {
		
		crutime := time.Now().Unix()
		h :=md5.New()
		io.WriteString(h, strconv.FormatInt(crutime,10))
		token :=fmt.Sprintf("%x",h.Sum(nil))
		
		
		io.WriteString(w,token)
		//t,_ :=template.ParseFiles("upload.gtpl")
		//t.Execute(w,token)
	} else {
		
		// 2.服务端调用 r.ParseMultipartForm ,把上传的文件存储在内存和临时文件中 里面的参数表示 maxMemory ，
		r.ParseMultipartForm(32 << 20)
		//判断 key是否正确
		if r.Form.Get("publickey")=="vca"&&r.Form.Get("secretkey")=="vca666"{
			//fmt.Println("publickey:", r.Form.Get("publickey"))
			//fmt.Println("secretkey:", r.Form.Get("secretkey"))
		}else{
			result := RelCode{Code:408, Msg:"您填写的正确的key"}
			w.Header().Set("content-type", "application/json")
			b,err := json.Marshal(result)
			if err !=nil{
				log.Fatal(err)
			}
			fmt.Fprint(w,string(b))
			return
		}
		
		// 3. 使用 r.FormFile 获取文件句柄，然后对文件进行存储等处理。
		file,handler,err := r.FormFile("uploadfile")
		
		if err != nil {
			//fmt.Println("11111111")
			log.Fatal(err)
			w.Header().Set("content-type", "application/json")
			
			
			result := RelCode{Code:409, Msg:"您填写的表单参数错误"}
			
			b,err := json.Marshal(result)
			if err !=nil{
				log.Fatal(err)
			}
			fmt.Fprint(w,string(b))
			return
		}
		
		defer file.Close()
		
		u1 := uuid.NewV4()
		fmt.Printf("UUIDv4: %s\n", u1.String())
		extent_type := handler.Filename[strings.LastIndex(handler.Filename, "."):]
		
		//产生新的图片名称
		new_Filename := u1.String()+extent_type
		
		//println(new_Filename)
		
		f, err := os.OpenFile("."+toolsdir +new_Filename,os.O_WRONLY|os.O_CREATE,0666)
		
		if err !=nil{
			//fmt.Println("222222222222")
			log.Fatal(err)
			return
		}
		
		w.Header().Set("content-type", "application/json")
		result := RelCode{Code:200,Data:adomin+toolsdir+new_Filename, Msg:"success"}
		
		b,err := json.Marshal(result)
		if err !=nil{
			log.Fatal(err)
		}
		
		fmt.Fprint(w,string(b))
		
		defer f.Close()
		io.Copy(f,file)
	}
	
}

func dev_upload(w http.ResponseWriter, r *http.Request){
	
	// 1.表单中增加enctype="multipart/form-data"
	//fmt.Println("method:",r.Method)
	
	if r.Method == "GET" {
		
		crutime := time.Now().Unix()
		h :=md5.New()
		io.WriteString(h, strconv.FormatInt(crutime,10))
		token :=fmt.Sprintf("%x",h.Sum(nil))
		
		
		io.WriteString(w,token)
		//t,_ :=template.ParseFiles("upload.gtpl")
		//t.Execute(w,token)
	} else {
		
		// 2.服务端调用 r.ParseMultipartForm ,把上传的文件存储在内存和临时文件中 里面的参数表示 maxMemory ，
		r.ParseMultipartForm(32 << 20)
		//判断 key是否正确
		if r.Form.Get("publickey")=="vca"&&r.Form.Get("secretkey")=="vca666"{
			fmt.Println("publickey:", r.Form.Get("publickey"))
			fmt.Println("secretkey:", r.Form.Get("secretkey"))
		}else{
			result := RelCode{Code:408, Msg:"您填写的正确的key"}
			w.Header().Set("content-type", "application/json")
			b,err := json.Marshal(result)
			if err !=nil{
				log.Fatal(err)
			}
			fmt.Fprint(w,string(b))
			return
		}
		
		// 3. 使用 r.FormFile 获取文件句柄，然后对文件进行存储等处理。
		file,handler,err := r.FormFile("uploadfile")
		
		if err != nil {
			fmt.Println("11111111")
			log.Fatal(err)
			w.Header().Set("content-type", "application/json")
			
			
			result := RelCode{Code:409, Msg:"您填写的表单参数错误"}
			
			b,err := json.Marshal(result)
			if err !=nil{
				log.Fatal(err)
			}
			fmt.Fprint(w,string(b))
			return
		}
		
		defer file.Close()
		
		u1 := uuid.NewV4()
		fmt.Printf("UUIDv4: %s\n", u1.String())
		extent_type := handler.Filename[strings.LastIndex(handler.Filename, "."):]
		
		//产生新的图片名称
		new_Filename := u1.String()+extent_type
		
		//println(new_Filename)
		
		new_pathandFilename:=devdir+new_Filename
		
		f, err := os.OpenFile("."+new_pathandFilename,os.O_WRONLY|os.O_CREATE,0666)
		
		if err !=nil{
			//fmt.Println("222222222222")
			log.Fatal(err)
			return
		}
		
		w.Header().Set("content-type", "application/json")
		result := RelCode{Code:200,Data:adomin+new_pathandFilename, Msg:"success"}
		
		b,err := json.Marshal(result)
		if err !=nil{
			log.Fatal(err)
		}
		
		fmt.Fprint(w,string(b))
		
		defer f.Close()
		io.Copy(f,file)
	}
	
}

func live_upload(w http.ResponseWriter, r *http.Request){
	
	// 1.表单中增加enctype="multipart/form-data"
	fmt.Println("method:",r.Method)
	
	if r.Method == "GET" {
		
		crutime := time.Now().Unix()
		h :=md5.New()
		io.WriteString(h, strconv.FormatInt(crutime,10))
		token :=fmt.Sprintf("%x",h.Sum(nil))
		
		
		io.WriteString(w,token)
		//t,_ :=template.ParseFiles("upload.gtpl")
		//t.Execute(w,token)
	} else {
		
		// 2.服务端调用 r.ParseMultipartForm ,把上传的文件存储在内存和临时文件中 里面的参数表示 maxMemory ，
		r.ParseMultipartForm(32 << 20)
		//判断 key是否正确
		if r.Form.Get("publickey")=="vca"&&r.Form.Get("secretkey")=="vca666"{
			fmt.Println("publickey:", r.Form.Get("publickey"))
			fmt.Println("secretkey:", r.Form.Get("secretkey"))
		}else{
			result := RelCode{Code:408, Msg:"您填写的正确的key"}
			w.Header().Set("content-type", "application/json")
			b,err := json.Marshal(result)
			if err !=nil{
				log.Fatal(err)
			}
			fmt.Fprint(w,string(b))
			return
		}
		
		// 3. 使用 r.FormFile 获取文件句柄，然后对文件进行存储等处理。
		file,handler,err := r.FormFile("uploadfile")
		
		if err != nil {
			fmt.Println("11111111")
			log.Fatal(err)
			w.Header().Set("content-type", "application/json")
			
			
			result := RelCode{Code:409, Msg:"您填写的表单参数错误"}
			
			b,err := json.Marshal(result)
			if err !=nil{
				log.Fatal(err)
			}
			fmt.Fprint(w,string(b))
			return
		}
		
		defer file.Close()
		
		u1 := uuid.NewV4()
		fmt.Printf("UUIDv4: %s\n", u1.String())
		extent_type := handler.Filename[strings.LastIndex(handler.Filename, "."):]
		
		//产生新的图片名称
		new_Filename := u1.String()+extent_type
		
		
		
		new_pathandFilename:=livedir+new_Filename
		
		f, err := os.OpenFile("."+new_pathandFilename,os.O_WRONLY|os.O_CREATE,0666)
		
		if err !=nil{
			//fmt.Println("222222222222")
			log.Fatal(err)
			return
		}
		
		w.Header().Set("content-type", "application/json")
		result := RelCode{Code:200,Data:adomin+new_pathandFilename, Msg:"success"}
		
		b,err := json.Marshal(result)
		if err !=nil{
			log.Fatal(err)
		}
		
		fmt.Fprint(w,string(b))
		
		defer f.Close()
		io.Copy(f,file)
	}
	
}


func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}



func makestaticDir( dir string){
	//11111.
	isexits,exitsErr :=PathExists(*realPath + dir)
	if exitsErr != nil {
		log.Fatal("static PathExists:", exitsErr)
	}
	if isexits == false {
		err := os.MkdirAll(*realPath + dir, os.ModePerm)  //在当前目录下生成md目录
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("创建目录" + *realPath + dir + "成功")
	}
}

func staticResource(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	
	
	println(r)
	
	println(path)
	
	if path=="/favicon.ico"{
		return
	}
	
	haveExtend := strings.Index(path,".")
	
	if haveExtend ==-1{
		fmt.Fprint(w,"not have resource")
		return
	}
	
	request_type := path[strings.LastIndex(path, "."):]
	switch request_type {
	case ".css":
		w.Header().Set("content-type", "text/css")
	case ".js":
		w.Header().Set("content-type", "text/javascript")
	default:
	}
	
	isexits,exitsErr :=PathExists(*realPath + path)
	
	
	if exitsErr != nil {
		log.Fatal("static PathExists:", exitsErr)
	}
	
	//fmt.Printf("isexits %v\n",isexits)
	//fmt.Printf("exitsErr %v\n",exitsErr)
	
	if isexits {
		fmt.Printf("isexits %v\n",isexits)
		
		
		println(*realPath)
		
		fin, err := os.Open(*realPath + path)
		
		println(fin)
		
		defer fin.Close()
		
		if err != nil {
			//log.Fatal("static resource1:", err)
			fmt.Fprintf(w,"static resource err is:%v",err)
		}
		
		fd, error := ioutil.ReadAll(fin)
		
		if error != nil {
			log.Fatal("static resource2:", error)
			fmt.Fprintf(w,"static resource err is:%v",error)
		}
		
		w.Write(fd)
		
	} else {
		fmt.Printf("isexits %v\n",isexits)
		
		io.WriteString(w,"资源路径不存在")
	}
	
	
	
}
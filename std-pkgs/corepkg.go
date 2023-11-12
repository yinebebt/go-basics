package main

import (
	"container/list"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net" // for tcp
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"sort"
)

// opening a file
func openFIle() { //you can pass the file as input arg.  fileName string
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

// or inshort
func openFIle1() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

// creating a file
func fileCrt() {
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString("test")
}

// opening directory
func openDir() { //you can pass the path here, path stirng
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

// or inshort
func openDIr1() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
} //Walk is called for every file and folder in the root folder. (in this case .)

// pkg container - list and more collection of types
func empList() {
	var x list.List //List in go is doubly-linked list
	x.PushBack(1)   // to add items into the list
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

// pkg sort , *?Need more reading to get the idea
// Sort function in sort takes a sort.Interface and sorts it.sort.Interface requires 3 methods: Len, Less and Swap.
// sort by name
type Person struct {
	Name string
	Age  int
}
type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func toBe_main() { // take this to main
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}

//sort end

//pkg cryptography: cryptographic from pkg-crypto; and they are hard to revert. eg sha-1
// and non cryptographic- from pkg-hash eg. adler32,crc32, crc64 and fnv.

// examples: hashing may used to compare to datas
func empcrc32() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}

func empsha1() {
	h := sha1.New()
	h.Write([]byte("test")) //both crc32 and sha1 implement hash.Hash interface; that is why the codes are similar
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}

// pkg servers
// creating a TCP server
func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg) //Decode reads the next value from the input stream and stores it in
	//the data represented by the empty interface value.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}
func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999") //Dial connects to the address on the named network.
	if err != nil {
		fmt.Println(err)
		return
	}
	// send the message
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg) //Encode transmits the data item represented by the empty interface value,
	//guaranteeing that all necessary type information has been transmitted first.
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func toBemain() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}

// tcp server end

// /http example
func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type", // key
		"text/html",    // value
	)
	io.WriteString(
		res,
		`<doctype html>
	<html>
	<head>
	<title>Hello World</title>
	</head>
	<body>
	Hello World!
	</body>
	</html>`,
	)
}
func toBeser_main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
} //HandleFunc handles a URL route (/hello) by calling the given function.
//We can also handle static files by using FileServer:

//http end

// RPC example,*? Need more refactooring to get the idea
type Server struct{} //object to store methods those will invoked over ntw.
func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}
func rpcserver() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
func rpcclient() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999),
		&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =",
			result)
	}
}
func toBeRpcmain() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}

//rpc end

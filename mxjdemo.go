package main

import (
	"fmt"
	"os"

	"github.com/clbanning/mxj"
)

func main() {
	xml2json()
	json2xml()
}

func json2xml() {
	jsonValue, err := readFromFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========")
	fmt.Println("Json value:")
	fmt.Println("==========")
	fmt.Println(string(jsonValue))

	mv, err := mxj.NewMapJson(jsonValue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========")
	fmt.Println("Map value:")
	fmt.Println("==========")
	fmt.Println(mv)

	xmlValue, err := mv.XmlIndent("", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========")
	fmt.Println("XML value:")
	fmt.Println("==========")
	fmt.Println(string(xmlValue))

	jsonValue2, err := mv.JsonIndent("", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========")
	fmt.Println("JSON value:")
	fmt.Println("==========")
	fmt.Println(string(jsonValue2))

}

func xml2json() {
	xmlVlaue, err := readFromFile("data.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("==========")
	fmt.Println("xml value:")
	fmt.Println("==========")
	fmt.Println(string(xmlVlaue))

	mv, err := mxj.NewMapXml(xmlVlaue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========")
	fmt.Println("Map value:")
	fmt.Println("==========")
	fmt.Println(mv, "\n")

	jsonValue, err := mv.JsonIndent("", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("==========")
	fmt.Println("json value:")
	fmt.Println("==========")
	fmt.Println(string(jsonValue))

	xmlValue2, err := mv.XmlIndent("", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("============")
	fmt.Println("back to xml:")
	fmt.Println("============")
	fmt.Println(string(xmlValue2))
}

func readFromFile(fileName string) ([]byte, error) {
	fh, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	fs, _ := fh.Stat()
	fileData := make([]byte, fs.Size())
	n, err := fh.Read(fileData)
	if err != nil {
		return nil, err
	}
	if int64(n) != fs.Size() {
		fmt.Println("n:", n, "fs.Size():", fs.Size())
		return nil, err
	}
	fmt.Println("Read ", n, " bytes from ", fileName)
	return fileData, nil
}

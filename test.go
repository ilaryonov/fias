package main

import "fmt"

type DecodeNode func(Node) *interface{}

type Node struct {
	something string
}

func Import(pathName string, decodeNode DecodeNode, config Config) {
	node := parseFile()

	result := decodeNode(node)

	if result != nil {
		collection = append(collection, result)
	}

	BatchInsert(db, collection)

}



type House struct {
	bla string
}

func ImportHouse(...) {
	Import("abc", func(node Node) *interface{} {
		element = node.Name.Local
		if element == "House" {
			a := House{}
			decoder.DecodeElement(&a, &node)

			a.ID = 0

			return a
		}
		return nil
	})
}

type Address struct {
	bla2 string
}


func  ImportHouse(...) {
	Import("abc", func(node Node) *interface{} {
		element = node.Name.Local
		if element == "House" {

			decoder.DecodeElement(&a, &se)

			a.ID = 0

			return a
		}
		return nil
	})
}


func WalkDirectories() {
	if (isHouse(file)){
		House()
	}
}
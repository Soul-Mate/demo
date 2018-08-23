package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"github.com/Soul-Mate/todo"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "input error, you can use add or list\n")
		os.Exit(1)
	}
	var err error
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			err = fmt.Errorf("no input in add")
		} else {
			err = add(os.Args[2])
		}
	case "list":
		err = list()
	default:
		err = fmt.Errorf("not support cmd (%s)", os.Args[1])
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "execute error: %v\n", err)
		os.Exit(1)
	}
}

const dbPath = "./todo.db"

func add(text string) error {
	td := &Todo{
		Text: text,
		Done: false,
	}
	f, err := os.OpenFile(dbPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := proto.Marshal(td)
	if err != nil {
		return err
	}

	// 一个字节标识消息的大小
	if err = binary.Write(f, binary.LittleEndian, int64(len(data))); err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}
	return nil
}

func list() error {
	b, err := ioutil.ReadFile(dbPath)

	if err != nil {
		return fmt.Errorf("could not read %s: %v", dbPath, err)
	}
	for {
		if len(b) == 0 {
			return nil
		}

		if len(b) < 8 {
			return fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}

		var l int64
		// 读取消息的大小
		if err := binary.Read(bytes.NewReader(b[:8]), binary.LittleEndian, &l); err != nil {
			return fmt.Errorf("could not decode message length: %v", err)
		}

		b = b[8:]
		var td Todo
		// 解析消息
		if err := proto.Unmarshal(b[:l], &td); err != nil {
			return fmt.Errorf("could not read task: %v", err)
		}

		b = b[l:]

		if td.Done {
			fmt.Printf("👍")
		} else {
			fmt.Printf("😱")
		}
		fmt.Printf(" %s\n", td.Text)
	}
	return nil
}

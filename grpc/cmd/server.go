package main

import (
	"context"
	"github.com/Soul-Mate/demo/grpc/todo"
	"os"
	"fmt"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"bytes"
	"net"
	"log"
	"google.golang.org/grpc"
)

const (
	DB_FILE = "./todo.db"
)

func main() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcServer := grpc.NewServer()
	todo.RegisterTaskServiceServer(rpcServer, &TaskServer{})
	rpcServer.Serve(lis)
}

type TaskServer struct {
}

func (s *TaskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	var todoList = new(todo.TaskList)
	b, err := ioutil.ReadFile(DB_FILE)
	if err != nil {
		return todoList, fmt.Errorf("could not write task to file: %v", err)
	}
	for {
		if len(b) == 0 {
			return todoList, nil
		}

		if len(b) < 8 {
			return todoList, fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}
		var l int64
		if err = binary.Read(bytes.NewReader(b[:8]), binary.LittleEndian, &l); err != nil {
			return todoList, fmt.Errorf("could not decode message length: %v", err)
		}
		b = b[8:]
		todoTask := new(todo.Task)
		if err = proto.Unmarshal(b[:l], todoTask); err != nil {
			return todoList, fmt.Errorf("could not read task: %v", err)
		}
		var status string
		if todoTask.Done {
			status = "👍"
		} else {
			status = "😱"
		}
		todoTask.Text = fmt.Sprintf("%s %s", status, todoTask.Text)
		todoList.Tasks = append(todoList.Tasks, todoTask)
		b = b[l:]
	}
	return todoList, nil
}

func (s *TaskServer) Add(ctx context.Context, task *todo.Task) (*todo.Void, error) {
	void := &todo.Void{}
	f, err := os.OpenFile(DB_FILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return void, fmt.Errorf("can't create file: %v", err)
	}
	b, err := proto.Marshal(task)
	if err != nil {
		return void, fmt.Errorf("parse task message error: %v", err)
	}
	if err = binary.Write(f, binary.LittleEndian, int64(len(b))); err != nil {
		return void, fmt.Errorf("write task message length error: %v", err)
	}

	if _, err = f.Write(b); err != nil {
		return void, fmt.Errorf("write task message error: %v", err)
	}

	if err = f.Close(); err != nil {
		return void, fmt.Errorf("close io stream error: %v", err)
	}
	return void, nil
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
	ch   chan string
}

type Server struct {
	clients    map[net.Conn]*Client
	joining    chan *Client
	leaving    chan *Client
	messages   chan string
	mutex      sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		clients:  make(map[net.Conn]*Client),
		joining:  make(chan *Client),
		leaving:  make(chan *Client),
		messages: make(chan string),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server started on :8080")
	
	go s.run()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go s.handleClient(conn)
	}
}

func (s *Server) run() {
	for {
		select {
		case client := <-s.joining:
			s.mutex.Lock()
			s.clients[client.conn] = client
			s.mutex.Unlock()
			
			go func() {
				client.ch <- fmt.Sprintf("Welcome %s! There are %d users online.", client.name, len(s.clients))
			}()
			
			s.broadcast(fmt.Sprintf("%s joined the chat", client.name), client)

		case client := <-s.leaving:
			s.mutex.Lock()
			delete(s.clients, client.conn)
			close(client.ch)
			s.mutex.Unlock()
			
			s.broadcast(fmt.Sprintf("%s left the chat", client.name), nil)

		case message := <-s.messages:
			s.broadcast(message, nil)
		}
	}
}

func (s *Server) broadcast(message string, sender *Client) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	for _, client := range s.clients {
		if client != sender {
			select {
			case client.ch <- message:
			default:
				close(client.ch)
				delete(s.clients, client.conn)
			}
		}
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()
	
	conn.Write([]byte("Enter your name: "))
	scanner := bufio.NewScanner(conn)
	
	if !scanner.Scan() {
		return
	}
	
	name := strings.TrimSpace(scanner.Text())
	if name == "" {
		name = "Anonymous"
	}

	client := &Client{
		conn: conn,
		name: name,
		ch:   make(chan string),
	}

	s.joining <- client

	go func() {
		for message := range client.ch {
			conn.Write([]byte(message + "\n"))
		}
	}()

	for scanner.Scan() {
		message := strings.TrimSpace(scanner.Text())
		if message == "" {
			continue
		}
		
		if message == "/quit" {
			break
		}
		
		if strings.HasPrefix(message, "/users") {
			s.mutex.RLock()
			userList := make([]string, 0, len(s.clients))
			for _, c := range s.clients {
				userList = append(userList, c.name)
			}
			s.mutex.RUnlock()
			
			client.ch <- fmt.Sprintf("Online users: %s", strings.Join(userList, ", "))
			continue
		}

		s.messages <- fmt.Sprintf("%s: %s", client.name, message)
	}

	s.leaving <- client
}

func main() {
	server := NewServer()
	server.Start()
}

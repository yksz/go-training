package ftp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func ListenAndServe(port int) {
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	conn := newControlConn(c)
	conn.reply(220)
	for {
		line, err := conn.reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print(err)
			break
		}
		conn.interpret(line)
	}
}

type controlConn struct {
	conn    net.Conn
	reader  *bufio.Reader
	writer  *bufio.Writer
	session *session
}

type session struct {
	user         string
	dataConnAddr string
	dataConnPort int
}

func newControlConn(c net.Conn) *controlConn {
	return &controlConn{
		conn:   c,
		reader: bufio.NewReader(c),
		writer: bufio.NewWriter(c),
	}
}

func (c *controlConn) reply(code int) {
	text := fmt.Sprintf("%d %s\n", code, replies[code])
	log.Printf("[REPLY] " + text)
	c.writer.WriteString(text)
	c.writer.Flush()
}

func (c *controlConn) interpret(line string) {
	log.Print("[COMMAND] " + line)
	cmd, params, ok := c.parse(line)
	if ok {
		c.execute(cmd, params)
	}
}

func (c *controlConn) parse(line string) (command, []string, bool) {
	line = strings.TrimSpace(line)
	s := strings.Split(line, " ")
	cmd, ok := commands[s[0]]
	if !ok {
		log.Printf("Unknown command: %s\n", s[0])
		c.reply(502)
		return Unknown, nil, false
	}
	return cmd, s[1:], ok
}

func (c *controlConn) execute(cmd command, params []string) {
	switch cmd {
	case User:
		if len(params) != 1 {
			c.reply(501)
			return
		}
		c.session = &session{user: params[0]}
		c.reply(230)
	case Quit:
		c.session = nil
		c.reply(221)
	case Port:
		addr, port, ok := parsePort(params)
		if !ok {
			c.reply(501)
			return
		}
		c.session.dataConnAddr = addr
		c.session.dataConnPort = port
		c.reply(200)
	case Type:
		c.reply(502)
	case Retr:
		c.reply(502)
	case Stor:
		c.reply(502)
	}
}

func parsePort(params []string) (string, int, bool) {
	if len(params) != 1 {
		return "", 0, false
	}
	s := strings.Split(params[0], ",")
	if len(s) != 6 {
		return "", 0, false
	}
	fields := make([]int, 6)
	for i, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			return "", 0, false
		}
		fields[i] = n
	}
	addr := strings.Join(s[:4], ".")
	port := fields[4]*256 + fields[5]
	return addr, port, true
}

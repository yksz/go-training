package ftp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
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
	workingDir   string
	dataConnAddr string
}

func newControlConn(c net.Conn) *controlConn {
	return &controlConn{
		conn:   c,
		reader: bufio.NewReader(c),
		writer: bufio.NewWriter(c),
	}
}

func (c *controlConn) reply(code int) error {
	text := fmt.Sprintf("%d %s\n", code, replies[code])
	log.Printf("[REPLY] " + text)
	if _, err := c.writer.WriteString(text); err != nil {
		return err
	}
	if err := c.writer.Flush(); err != nil {
		return err
	}
	return nil
}

func (c *controlConn) interpret(line string) {
	log.Print("[COMMAND] " + line)
	cmd, args, ok := c.parse(line)
	if ok {
		c.execute(cmd, args)
	}
}

func (c *controlConn) parse(line string) (command, []string, bool) {
	line = strings.TrimSpace(line)
	s := strings.Split(line, " ")
	cmd, ok := commands[s[0]]
	if !ok {
		log.Printf("unknown command: %s\n", s[0])
		c.reply(502)
		return Unknown, nil, false
	}
	return cmd, s[1:], ok
}

func (c *controlConn) execute(cmd command, args []string) {
	switch cmd {
	case User:
		user, ok := paramAsString(args)
		if !ok {
			c.reply(501)
			return
		}
		if !authenticate(user) {
			c.reply(530)
			return
		}
		c.session = &session{
			user:       user,
			workingDir: "/home/" + user}
		c.reply(230)
	case Quit:
		c.session = nil
		c.reply(221)
	case Cwd:
		targetDir, ok := paramAsString(args)
		if !ok {
			c.reply(501)
			return
		}
		dir, ok := changeDir(c.session.workingDir, targetDir)
		if !ok {
			c.reply(550)
			return
		}
		log.Printf("pwd: %s\n", dir)
		c.session.workingDir = dir
		c.reply(200)
	case Port:
		ipAddr, port, ok := paramsPort(args)
		if !ok {
			c.reply(501)
			return
		}
		c.session.dataConnAddr = ipAddr + ":" + strconv.Itoa(port)
		c.reply(200)
	case Type:
		c.reply(502)
	case Retr:
		filename, ok := paramAsString(args)
		if !ok {
			c.reply(501)
			return
		}
		filepath := c.session.workingDir + string(os.PathSeparator) + filename
		if !exists(filepath) {
			log.Printf("file not found: %s\n", filename)
			c.reply(550)
			return
		}
		file, err := os.Stat(filepath)
		if err != nil {
			log.Print(err)
			c.reply(550)
			return
		}
		if file.IsDir() {
			log.Printf("file is a directory: %s\n", filename)
			c.reply(550)
			return
		}
		bytes, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Print(err)
			c.reply(550)
			return
		}
		c.sendData(bytes)
	case Stor:
		c.reply(502)
	case List:
		list, err := fileList(c.session.workingDir)
		if err != nil {
			log.Print(err)
			c.reply(550)
			return
		}
		c.sendData([]byte(list))
	}
}

func (c *controlConn) sendData(data []byte) {
	dataConn, err := net.Dial("tcp", c.session.dataConnAddr)
	if err != nil {
		log.Print(err)
		c.reply(425)
		return
	}
	c.reply(150)
	defer dataConn.Close()
	if _, err := dataConn.Write(data); err != nil {
		log.Print(err)
		c.reply(426)
		return
	}
	c.reply(226)
}

func paramAsString(args []string) (string, bool) {
	if len(args) != 1 {
		return "", false
	}
	return args[0], true
}

func authenticate(user string) bool {
	return exists("/home/" + user)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func changeDir(base, target string) (string, bool) {
	dir := ""
	if strings.HasPrefix(target, "/") {
		dir = filepath.Clean(target)
	} else {
		dir = filepath.Clean(base + string(os.PathSeparator) + target)
	}
	if exists(dir) {
		return dir, true
	}
	return "", false
}

func paramsPort(args []string) (string, int, bool) {
	if len(args) != 1 {
		return "", 0, false
	}
	s := strings.Split(args[0], ",")
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
	ipAddr := strings.Join(s[:4], ".")
	port := fields[4]*256 + fields[5]
	return ipAddr, port, true
}

func fileList(dir string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	for _, file := range files {
		if file.IsDir() {
			fmt.Fprintf(&buf, "%s/\r\n", file.Name())
		} else {
			fmt.Fprintf(&buf, "%s\r\n", file.Name())
		}
	}
	return buf.String(), nil
}

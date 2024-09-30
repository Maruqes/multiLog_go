package multiLog

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var messageChan = make(chan string)
var replyChan = make(chan string)
var conn net.Conn

func execProgram(path string, port string) {
	to_exec := path
	cmd := exec.Command(to_exec, port)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
	}
}

func checkIdentifier(identifier string) error {
	if strings.Contains(identifier, " ") {
		return fmt.Errorf("identifier cannot contain spaces")
	}
	return nil
}

func add_tab(identifier string, content string) {
	err := checkIdentifier(identifier)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	contentb64 := base64.StdEncoding.EncodeToString([]byte(content))

	final_content := "add_tab " + identifier + " " + contentb64

	messageChan <- final_content
}

func add_content(identifier string, content string) {
	err := checkIdentifier(identifier)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	contentb64 := base64.StdEncoding.EncodeToString([]byte(content))

	final_content := "add_content " + identifier + " " + contentb64

	messageChan <- final_content
}

func remove_tab(identifier string) {
	err := checkIdentifier(identifier)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	final_content := "remove_tab " + identifier

	messageChan <- final_content
}

func createFolder(name string) string {
	homeDir, _ := os.UserHomeDir()
	newDir := filepath.Join(homeDir, name)

	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		// Create the folder
		os.Mkdir(newDir, 0755)
	}

	return newDir
}

func chmodFile(path string) {
	fmt.Println("Changing file permissions to executable...")
	err := os.Chmod(path, 0755)
	if err != nil {
		fmt.Println("Error changing file permissions:", err)
	}
}

func downloadMultiLog(path string) {
	fmt.Println("Downloading multilog executable...")
	defer func() {
		fmt.Println("Download complete")
	}()
	url := "https://github.com/Maruqes/multiLog/releases/download/v0.1.0/multilog_0.1.0_amd64.AppImage"
	output := filepath.Join(path, "multilog_0.1.0_amd64.AppImage")

	// Create the file
	out, err := os.Create(output)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Bad status:", resp.Status)
		return
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Downloaded multilog executable successfully")
}

func check_if_file_exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func isPortInUse(port int) bool {
	// Try to create a listener on the given port
	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		// If we get an error, the port is in use or unavailable
		return true
	}
	// Close the listener as we are just checking
	listener.Close()
	return false
}

func generate_random_port() int {
	port := rand.Intn(65535-1024) + 1024
	if port < 1024 {
		return generate_random_port()
	}
	if isPortInUse(port) {
		return generate_random_port()
	}
	return port
}

func Init_multiLog() {
	port := generate_random_port()

	path := createFolder("multi_logs")
	exists := check_if_file_exists(path + "/multilog_0.1.0_amd64.AppImage")
	if !exists {
		downloadMultiLog(path)
	}
	chmodFile(filepath.Join(path, "multilog_0.1.0_amd64.AppImage"))
	execProgram(path+"/multilog_0.1.0_amd64.AppImage", fmt.Sprint(port))

	var err error
	for i := 0; i < 20; i++ {
		conn, err = net.Dial("tcp", "127.0.0.1:"+fmt.Sprint(port))
		if err == nil {
			break
		}
		fmt.Println("Waiting for server to start...")
		time.Sleep(500 * time.Millisecond)
	}
	can_it_continue := false

	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	fmt.Println("Connected to server at 127.0.0.1:" + fmt.Sprint(port))

	go func() {
		for {
			reply, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from server:", err)
				close(replyChan)
				return
			}
			replyChan <- reply
		}
	}()

	go func() {
		for {
			select {
			case message := <-messageChan:
				_, err := fmt.Fprintf(conn, message+"\n")
				if err != nil {
					fmt.Println("Error writing to server:", err)
					return
				}
			case reply, ok := <-replyChan:
				if !ok {
					return
				}
				if strings.Contains(reply, "continue_execution") {
					can_it_continue = true
				}
				fmt.Println("Server reply:", reply)
			}
		}
	}()

	for !can_it_continue {
		time.Sleep(100 * time.Millisecond)
	}
}

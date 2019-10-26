package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	// "strconv"
	// "strings"
	"time"
	"encoding/hex"
)

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide a port number!")
                return
        }

        PORT := ":" + arguments[1]
        l, err := net.Listen("tcp4", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()
        rand.Seed(time.Now().Unix())

        for {
                c, err := l.Accept()
                if err != nil {
                        fmt.Println(err)
                        return
                }
                go handleConnection(c)
        }
}

func handleConnection(c net.Conn) {
        fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	scanner := bufio.NewScanner(c)

	// Call Split to specify that we want to Scan each individual byte.
	scanner.Split(bufio.ScanBytes)

	// Use For-loop.
	for scanner.Scan() {
		// Get Bytes and display the byte.
		b := scanner.Bytes()
		dst := make([]byte, hex.EncodedLen(len(b)))
		hex.Encode(dst, b)
		fmt.Printf("%v = %s = %v\n", b, dst, string(b))
	}

	// buf := make([]byte, 1024)
	// // Read the incoming connection into the buffer.
	// _, err := c.Read(buf)
	// if err != nil {
	// 	fmt.Println("Error reading:", err.Error())
	// }
	// fmt.Printf(hex.EncodeToString(buf))
	// Send a response back to person contacting us.
	// c.Write([]byte("Message received."))
	// Close the connection when you're done with it.

	// netData, err := bufio.NewReader(c).ReadString('\n')
	// if err != nil {
	//         fmt.Println(err)
	//         return
	// }

	// temp := strings.TrimSpace(string(netData))
	// if temp == "STOP" {
	//         break
	// }

	// result := strconv.Itoa(random()) + "\n"
	// c.Write([]byte(string(result)))
        c.Close()
}

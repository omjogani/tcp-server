package server

import (
	"flag"
	"net"

	"github.com/fatih/color"
	"github.com/omjogani/tcp-server/protocol"
)

const (
	TCP  string = "tcp"
	TCP4 string = "tcp4"
	TCP6 string = "tcp6"
	UNIX string = "unix"
)

func isSupportedProtocol(network string) bool {
	switch network {
	case TCP, TCP4, TCP6, UNIX:
		return true
	default:
		return false
	}
}

func networkListener(network string, addr string) net.Listener {
	listener, err := net.Listen(network, addr)
	handleError(err, "failed to create listener")

	color.Blue("Service started: (%s) %s\n", network, addr)
	return listener
}

func handleIncomingRequestContinuously(listener net.Listener) {
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			color.Red("TCP_SERVER_ERROR: failed to accept incoming connection request: ", err)
			closeConnection(conn)
			continue
		}

		color.Green("Connected to ", conn.RemoteAddr(), "\n")
		go maintainConnection(conn)
	}
}

func closeConnection(conn net.Conn) {
	err := conn.Close()
	handleError(err, "failed to close connection")
}

func showGuideOfUsage(conn net.Conn) {
	_, err := conn.Write([]byte("Usage: GET /path <Type: JSON, NORMAL>\n"))
	handleError(err, "failed to write usage guide")
}

func maintainConnection(conn net.Conn) {
	defer closeConnection(conn)
	showGuideOfUsage(conn)

	for {
		fullCommand := make([]byte, (1024 * 4))
		commandLength, err := conn.Read(fullCommand)
		handleError(err, "failed to read command")
		if commandLength == 0 {
			color.Red("TCP_SERVER_ERROR: failed to read command | ", err)
			return
		}

		command, path, bodyType := protocol.ParseCommand(string(fullCommand[0:commandLength]))
		if command == "" {
			_, err := conn.Write([]byte("Invalid Command!\n"))
			handleError(err, "failed to acknowledge user about Invalid Command!")
			continue
		} else if command == "GET" {
			GET(conn)
		} else if command == "POST" {
			POST(conn, bodyType)
		}

		color.Blue(path)
		color.Blue(bodyType)
	}

}

func TCPServer() {
	var addr string
	var network string
	flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
	flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
	flag.Parse()

	if !isSupportedProtocol(network) {
		color.Red("TCP_SERVER_ERROR: Unsupported network protocol: ", network)
	}

	listener := networkListener(network, addr)
	handleIncomingRequestContinuously(listener)
}

func handleError(err error, errorMessage string) {
	if err != nil {
		color.Red("TCP_SERVER_ERROR: ", errorMessage, " | ", err)
	}
}

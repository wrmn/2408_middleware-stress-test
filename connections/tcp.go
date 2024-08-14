package connections

import (
	"2408_middleware-stress-test/global"
	"fmt"
	"net"
	"time"
)

func SendTCPRequest(message []byte) (string, error) {
	host := fmt.Sprintf("%s:%s", *global.HOST, *global.PORT)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return "", fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		return "", fmt.Errorf("failed to send message: %w", err)
	}

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	response := string(buffer[:n])
	return response, nil
}

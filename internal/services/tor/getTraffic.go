package tor

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Seicrypto/torcontroller/internal/services/logger"
)

func GetTorTrafficMetrics() (string, string, error) {
	logger := logger.GetLogger()

	conn, err := net.Dial("tcp", "127.0.0.1:9051")
	if err != nil {
		return "", "", fmt.Errorf("failed to connect to Tor control port: %v", err)
	}
	defer conn.Close()

	// 讀取 control.authcookie 並檢查長度
	cookie, err := os.ReadFile("/var/lib/tor/control.authcookie")
	if err != nil {
		return "", "", fmt.Errorf("failed to read control.authcookie: %v", err)
	}
	if len(cookie) != 32 {
		return "", "", fmt.Errorf("invalid control.authcookie length: expected 32 bytes, got %d", len(cookie))
	}

	// 發送 AUTHENTICATE 命令
	authCommand := fmt.Sprintf("AUTHENTICATE %s\r\n", hex.EncodeToString(cookie))
	_, err = conn.Write([]byte(authCommand))
	if err != nil {
		return "", "", fmt.Errorf("failed to send authenticate command: %v", err)
	}

	// 讀取 AUTHENTICATE 響應
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("failed to read authenticate response: %v", err)
	}
	if !strings.HasPrefix(line, "250 OK") {
		return "", "", fmt.Errorf("authentication failed: %s", line)
	}

	// 獲取 traffic/read 和 traffic/written
	metrics := make(map[string]string)
	queries := []string{"traffic/read", "traffic/written"}
	for _, query := range queries {
		_, err = conn.Write([]byte(fmt.Sprintf("GETINFO %s\r\n", query)))
		if err != nil {
			return "", "", fmt.Errorf("failed to send GETINFO command for %s: %v", query, err)
		}

		// 收集回應行
		for {
			line, err = reader.ReadString('\n')
			if err != nil {
				return "", "", fmt.Errorf("failed to read response for %s: %v", query, err)
			}

			logger.Info(fmt.Sprintf("Received line for %s: %s", query, line)) // 日誌輸出

			if strings.HasPrefix(line, "250-") {
				// 提取數據並存入 map
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					metrics[query] = strings.TrimSpace(parts[1])
				}
			} else if strings.HasPrefix(line, "250 OK") {
				break // 碰到結束符退出
			}
		}
	}

	readTraffic, okRead := metrics["traffic/read"]
	writtenTraffic, okWritten := metrics["traffic/written"]

	if !okRead || !okWritten {
		return "", "", fmt.Errorf("failed to parse traffic metrics: %+v", metrics)
	}

	return readTraffic, writtenTraffic, nil
}

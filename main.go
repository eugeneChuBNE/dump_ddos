package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Helper function to generate random IP addresses
func randomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// Helper function to simulate random excess rate
func randomExcess() float64 {
	return 50.0 + rand.Float64()*0.2
}

// Helper function to generate a random user agent
func randomUserAgent() string {
	agents := []string{
		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 5.0) AppleWebKit/537.36 (KHTML, like Gecko) Mobile Safari/537.36 (compatible; Bytespider; spider-feedback@bytedance.com)",
		"Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) Chrome/116.0.1938.76 Safari/537.36",
		"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/25.0 Chrome/121.0.0.0 Mobile Safari/537.36",
	}
	return agents[rand.Intn(len(agents))]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Open or create both log files
	elogFile, err := os.OpenFile("ddos.elog", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening elog file:", err)
		return
	}
	defer elogFile.Close()

	logFile, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	for {
		currentTime := time.Now()
		timestamp := currentTime.Format("2006/01/02 15:04:05")
		ip := randomIP()
		excess := randomExcess()

		// Write to error log (.elog)
		elogEntry := fmt.Sprintf("%s [error] 925#925: *13782709 limiting requests, excess: %.3f by zone \"zoneC\", client: %s, server: replicacollect.com, request: \"POST /my-account/ HTTP/2.0\", host: \"replicacollect.com\", referrer: \"https://replicacollect.com/\"\n", timestamp, excess, ip)
		if _, err := elogFile.WriteString(elogEntry); err != nil {
			fmt.Println("Error writing to elog file:", err)
			return
		}

		// Write to access log (.log)
		logTimestamp := currentTime.Format("02/Jan/2006:15:04:05 -0700")
		userAgent := randomUserAgent()
		logEntry := fmt.Sprintf("replicacollect.com - %s - DE [%s] \"GET / HTTP/2.0\" 307 168 \"-\" \"%s\" \"-\" | \"\"\n", ip, logTimestamp, userAgent)
		if _, err := logFile.WriteString(logEntry); err != nil {
			fmt.Println("Error writing to log file:", err)
			return
		}

		time.Sleep(100 * time.Millisecond) // Adjust this value to control the frequency of log entries
	}
}

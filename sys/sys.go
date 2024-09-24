package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	epfd, err := unix.EpollCreate1(0)
	if err != nil {
		log.Fatalf("failed to create epoll instance: %v", err)
	}
	defer unix.Close(epfd)

	// Open a file descriptor to monitor (stdin)
	fd := int(os.Stdin.Fd())

	// Define the event to monitor (input available to read)
	event := unix.EpollEvent{
		Events: unix.EPOLLIN,
		Fd:     int32(fd),
	}

	// Register the file descriptor with epoll
	err = unix.EpollCtl(epfd, unix.EPOLL_CTL_ADD, fd, &event)
	if err != nil {
		log.Fatalf("failed to add fd to epoll: %v", err)
	}

	events := make([]unix.EpollEvent, 10)

	fmt.Println("waiting for events...")

	for {
		n, err := unix.EpollWait(epfd, events, 1000) // 1s timeout
		if err != nil {
			log.Fatalf("epoll wait failed: %v", err)
		}

		for i := 0; i < n; i++ {
			if events[i].Events&unix.EPOLLIN != 0 {
				buf := make([]byte, 1024)
				n, err := os.Stdin.Read(buf)
				if err != nil {
					log.Fatalf("failed to read from stdin: %v", err)
				}
				fmt.Printf("read %d bytes: %s\n", n, string(buf[:n]))
			}
		}
	}
}

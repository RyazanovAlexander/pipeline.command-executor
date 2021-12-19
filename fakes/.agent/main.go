/*
MIT License

Copyright The command-executor Authors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	pb "fake_go_module/server"
)

const EXECUTOR_TARGET = "PIPELINE_AGENT__EXECUTORS__0__TARGET"

func main() {
	log.SetOutput(os.Stdout)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	errCh := make(chan error)
	go func() {
		errCh <- runLoop(ctx)
	}()

	select {
	case <-sigCh:
		fmt.Println("Interrupt signal received. Finishing the application...")
		cancel()
		return
	case err := <-errCh:
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func runLoop(interruptCtx context.Context) error {
	address := os.Getenv(EXECUTOR_TARGET)
	log.Printf("Agent address: %s", address)

	target := flag.String("addr", address, "the address to connect to")
	conn, err := grpc.Dial(*target, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewExecServiceClient(conn)
	commandCtx, cancel := context.WithCancel(interruptCtx)
	defer cancel()

	select {
	case <-time.After(time.Millisecond):
		if err := execCommands(client, commandCtx); err != nil {
			fmt.Println(err)
		}
	case <-interruptCtx.Done():
		return nil
	}

	return nil
}

func execCommands(client pb.ExecServiceClient, commandCtx context.Context) error {
	for i := 0; ; i++ {
		commands := pb.ExecCommands{
			Commands: []string{
				fmt.Sprintf("echo 'first %d'", i),
				fmt.Sprintf("echo 'second %d'", i),
				fmt.Sprintf("echo 'third %d'", i),
			},
		}

		result, err := client.ExecuteCommands(commandCtx, &commands)
		if err != nil {
			log.Print(err)
			time.Sleep(time.Second)
		} else {
			log.Printf("Result: %s", result.Output)
		}
	}
}

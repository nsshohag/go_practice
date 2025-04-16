// controlling timeouts
// cancelling go routines
// and passing metadata across go application -uthentication tokens, user IDs) across function calls without modifying function signatures.
// Graceful Shutdown of Services

/*
The context package in Go provides a way to manage t
imeouts, deadlines, cancellations, and metadata propagation across API calls
and goroutines. It is primarily used in concurrent programming to control long-running operations
and prevent resource leaks.
*/

/*
1. context.Background()
- Root Context
- Auto-Cancellation?	❌ No
- When to Use?	In main(), database connections, long-running tasks
- Use Background context to gracefully shut down the server

2. req.Context()
- Tied to HTTP Request? ✅ Yes
- Auto-Cancellation? ✅ Yes (when client cancels request)
- When to Use?	Inside HTTP handlers to manage request lifetimes
-
*/

// when to use contexts
/*
✅ When handling HTTP requests to cancel processing when the client disconnects.
✅ When executing database queries to set a timeout for long-running queries.
✅ When calling external APIs to avoid waiting indefinitely.
✅ When handling background tasks that should be cancelled if no longer needed.
*/

/*
1. ctx.Done() - Returns a channel that gets closed when the context is canceled.
2. ctx.Err() - Returns an error (context.Canceled or context.DeadlineExceeded).
3. ctx.Value(key) - Retrieves metadata stored in the context.
4. ctx.Deadline() - Returns the time when the operation should be canceled.
5. cancel() - is a function returned by context.WithTimeout,
and calling it manually cancels the context earlier if needed.
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	fmt.Println("context error before timeout =", ctx.Err())

	/*  // we can also cancel the context then
	// ctx.Err().Error() = context canceled
	go func() {
		time.Sleep(time.Second * 1)
		cancel()
	}()
	*/
	select {
	case <-time.After(7 * time.Second):
		fmt.Fprintf(w, "Within Context Time Done")

	case <-ctx.Done(): // Timeout Reached // if not done sends nill and if timeout sends a channel the closes when the contexts is done
		fmt.Println("context error after timeout =", ctx.Err())
		http.Error(w, "Request cancelled: "+ctx.Err().Error(), http.StatusRequestTimeout)

	}
}

func callThirdPartyAPI(ctx context.Context, userID int) (bool, error) {

	time.Sleep(200 * time.Millisecond)

	if ctx.Err() == context.DeadlineExceeded {
		return false, errors.New("context timeout exceeded")
	}
	fmt.Println(userID, "is done working within context time.")
	return true, nil
}

func main() {

	// 1.
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	id := 1123
	isUserSubbed, err := callThirdPartyAPI(ctx, id)

	if err != nil {
		log.Fatalf("error fetching user status for :%d error = %s\n", id, err)

	}

	if isUserSubbed {
		fmt.Printf("User %d is finished\n", id)
	}

	// 2.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe("localhost:8099", nil)

}

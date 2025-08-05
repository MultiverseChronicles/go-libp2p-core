
import (
)

// Introspector is the interface to be satisfied by components that are capable
// of spelunking the state of the system, and representing in accordance with
//
// It's very rare to build a custom implementation of this interface;
// it exists mostly for mocking. In most cases, you'll end up using the
// default introspector.
//
// to serve the data to clients, but they can also be used separately for
// embedding or testing.
//
// Experimental.

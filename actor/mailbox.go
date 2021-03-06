package actor

import "github.com/AsynkronIT/protoactor-go/mailbox"

var (
	defaultDispatcher mailbox.Dispatcher = mailbox.NewDefaultDispatcher(300)
)

var defaultMailboxProducer = mailbox.Unbounded()

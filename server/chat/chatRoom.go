package chat

import "sync"

var rooms sync.Map = sync.Map{}

type ChatRoom interface {
}

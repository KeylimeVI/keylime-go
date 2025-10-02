package essentials

import (
	"github.com/KeylimeVI/keylime-go/list"
	"github.com/KeylimeVI/keylime-go/pair"
	"github.com/KeylimeVI/keylime-go/set"
)

type List[T any] = kl.List[T]

var ListOf = kl.ListOf

type Set[T comparable] = ks.Set[T]

var SetOf = ks.SetOf

type Pair[A any, B any] = kp.Pair[A, B]

var PairOf = kp.PairOf

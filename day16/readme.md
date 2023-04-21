## --- Day 16: Packet Decoder ---

There is nothing complicated about this puzzle just usual bit stream parsing which we probably already implemented in our lives at least once elsewhere. I know I did. The implementation felt a bit tedious, but that's how these things go. Each packet is either literal or operator, and has a version and a value. This gives us the packet interface on the golang level:

```go
type packet interface {
	isLiteral() bool
	getVersion() int
	getValue() int64
}
```

Both literal and operator include a header, and the header implements `isLiteral` and `getVersion`. The implementation of `getValue` is defined for literal and operator separately. I chose to convert the hexadecimal input into a binary string. This way, when need to read `n` bits, I can use golang `x[:n]` facility and then convert the result with `strconv.ParseInt(s, 2, 32)` to get the final integer. This operation is implemented in the `getBits` method.

The rest of the solution both for parts 1 and 2 is even more straight-forward.

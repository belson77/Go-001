package comet

type Proto struct {
	PackLen   int32
	HeaderLen int16
	Ver       int16
	Seq       int32
	Body      []byte
}

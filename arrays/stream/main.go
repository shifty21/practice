package main

type OrderedStream struct {
	stream map[int]string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make(map[int]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	return nil
}

func main() {

}

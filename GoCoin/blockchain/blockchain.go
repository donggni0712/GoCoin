package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
	/*
		B1
			B1Hash = (data +  'X')
		B2
			B2hash = (data + B1Hash)
		B3
			B3Hash = (data + B2Hash)
		==> 이전 블록을 변경하면 이 블록이 무의미해짐 ==> 변경 불가능
	*/
}

// singleton pattern => 변수를 직접 드러내지 않고 함수를 통해 드러내는 것

type blockchain struct {
	blocks []block
}

var b *blockchain

//블록체인을 호출할 때, 처음 호출한다면 b를 만들어줌.
func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}

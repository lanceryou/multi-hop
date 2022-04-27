package multi_hop

//		{
//			"token0_symbol": "FRAX",
//			"token0_address": "0x853d955acef822db058eb8505911ed77f175b99e",
//			"token1_symbol": "USDC",
//			"token1_address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
//			"pair": "FRAX/USDC"
//		},
//		{
//			"token0_symbol": "FNK",
//			"token0_address": "0xb5fe099475d3030dde498c3bb6f3854f762a48ad",
//			"token1_symbol": "USDT",
//			"token1_address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
//			"pair": "FNK/USDT"
//		},
type TokenPair struct {
	Token0Symbol  string `json:"token0_symbol"`
	Token0Address string `json:"token0_address"`
	Token1Symbol  string `json:"token1_symbol"`
	Token1Address string `json:"token1_address"`
	Pair          string `json:"pair"`
}

type MultiHop interface {
	// MultiHopSwap 中间桥接token只选择以下四个稳定币: DAI，USDC，USDT，WETH。
	MultiHopSwap(paris []TokenPair, src string, dst string, maxStep uint32) [][]string
}

type MultiHopFunc func(paris []TokenPair, src string, dst string, maxStep uint32) [][]string

func (m MultiHopFunc) MultiHopSwap(paris []TokenPair, src string, dst string, maxStep uint32) [][]string {
	return m(paris, src, dst, maxStep)
}

func BacktraceMultiHop(paris []TokenPair, src string, dst string, step uint32) [][]string {
	pairMap := make(map[string][]TokenPair)
	for _, pair := range paris {
		pairMap[pair.Token0Symbol] = append(pairMap[pair.Token0Symbol], pair)
	}

	set := make(map[string]struct{})
	for _, mid := range []string{"DAI", "USDC", "USDT", "WETH"} {
		set[mid] = struct{}{}
	}

	var result [][]string
	var bt func(pairMap map[string][]TokenPair, curToken string, step uint32, mid []string)
	bt = func(pairMap map[string][]TokenPair, curToken string, step uint32, mid []string) {
		tokens := pairMap[curToken]
		for _, token := range tokens {
			if step == 0 {
				if dst == token.Token1Symbol {
					ret := make([]string, len(mid)+1)
					copy(ret, append(mid, dst))
					result = append(result, ret)
					return
				}
				continue
			}

			// mid-value must be in set
			if _, ok := set[token.Token1Symbol]; !ok && step != 1 {
				continue
			}

			mid = append(mid, token.Token1Symbol)
			bt(pairMap, token.Token1Symbol, step-1, mid)
			mid = mid[:len(mid)-1]
		}
		return
	}

	bt(pairMap, src, step, []string{src})
	return result
}

package multi_hop

import (
	"encoding/json"
	"testing"
)

func TestBacktraceMultiHop(t *testing.T) {
	ts := []struct {
		pairs  []TokenPair
		src    string
		dst    string
		step   uint32
		expect [][]string
	}{
		{},
		// test pairs empty
		{
			src: "123",
			dst: "456",
		},
		// test pairs has one
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
			},
			src:  "FRAX",
			dst:  "USD",
			step: 1,
		},
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/MOCK",
				},
			},
			src:  "FRAX",
			dst:  "WETH",
			step: 1,
		},
		// find one
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/WETH",
				},
			},
			src:  "FRAX",
			dst:  "WETH",
			step: 1,
			expect: [][]string{
				{"FRAX", "USDC", "WETH"},
			},
		},
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDT",
				},
				{
					Token0Symbol:  "USDT",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDT/WETH",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/WETH",
				},
			},
			src:  "FRAX",
			dst:  "WETH",
			step: 1,
			expect: [][]string{
				{"FRAX", "USDC", "WETH"},
				{"FRAX", "USDT", "WETH"},
			},
		},
		// test mid-state has other state
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDT",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/MOCK",
				},
				{
					Token0Symbol:  "USDT",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDT/WETH",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/WETH",
				},
			},
			src:  "FRAX",
			dst:  "WETH",
			step: 1,
			expect: [][]string{
				{"FRAX", "USDC", "WETH"},
				{"FRAX", "USDT", "WETH"},
			},
		},
		// test final can not reach
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDT",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/MOCK",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "DAI",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/DAI",
				},
				{
					Token0Symbol:  "USDT",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDT/WETH",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/WETH",
				},
				{
					Token0Symbol:  "DAI",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "DAI/MOCK",
				},
			},
			src:  "FRAX",
			dst:  "WETH",
			step: 1,
			expect: [][]string{
				{"FRAX", "USDC", "WETH"},
				{"FRAX", "USDT", "WETH"},
			},
		},
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDT",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/MOCK",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "DAI",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/DAI",
				},
				{
					Token0Symbol:  "USDT",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "RAY",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDT/RAY",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "RAY",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/RAY",
				},
				{
					Token0Symbol:  "DAI",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "DAI/MOCK",
				},
			},
			src:  "FRAX",
			dst:  "RAY",
			step: 1,
			expect: [][]string{
				{"FRAX", "USDC", "RAY"},
				{"FRAX", "USDT", "RAY"},
			},
		},
		// step 2 can not find final
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDT",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/MOCK",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "DAI",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/DAI",
				},
				{
					Token0Symbol:  "USDT",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "RAY",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDT/RAY",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "RAY",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/RAY",
				},
				{
					Token0Symbol:  "DAI",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "DAI/MOCK",
				},
			},
			src:  "FRAX",
			dst:  "RAY",
			step: 2,
		},
		// step 2 can find final,the last value is in mid-value
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDT",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/MOCK",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "DAI",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/DAI",
				},
				{
					Token0Symbol:  "USDT",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDT/WETH",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/USDT",
				},
				{
					Token0Symbol:  "DAI",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "MOCK",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "DAI/MOCK",
				},
			},
			src:  "FRAX",
			dst:  "WETH",
			step: 2,
			expect: [][]string{
				{"FRAX", "USDC", "USDT", "WETH"},
			},
		},
		// step 2 can find final, the last value is in mid-value
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDT",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "DAI",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/DAI",
				},
				{
					Token0Symbol:  "USDT",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDT/WETH",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/USDT",
				},
				{
					Token0Symbol:  "DAI",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "DAI/USDT",
				},
			},
			src:  "FRAX",
			dst:  "WETH",
			step: 2,
			expect: [][]string{
				{"FRAX", "USDC", "USDT", "WETH"},
				{"FRAX", "DAI", "USDT", "WETH"},
			},
		},
		{
			pairs: []TokenPair{
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDC",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDC",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/USDT",
				},
				{
					Token0Symbol:  "FRAX",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "DAI",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "FRAX/DAI",
				},
				{
					Token0Symbol:  "USDT",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDT/WETH",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/USDT",
				},
				{
					Token0Symbol:  "USDC",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "DAI",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "USDC/DAI",
				},
				{
					Token0Symbol:  "DAI",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "WETH",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "DAI/WETH",
				},
				{
					Token0Symbol:  "DAI",
					Token0Address: "0x853d955acef822db058eb8505911ed77f175b99e",
					Token1Symbol:  "USDT",
					Token1Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Pair:          "DAI/USDT",
				},
			},
			src:  "FRAX",
			dst:  "WETH",
			step: 2,
			expect: [][]string{
				{"FRAX", "USDC", "USDT", "WETH"},
				{"FRAX", "USDC", "DAI", "WETH"},
				{"FRAX", "DAI", "USDT", "WETH"},
			},
		},
	}

	for _, d := range ts {
		result := BacktraceMultiHop(d.pairs, d.src, d.dst, d.step)
		if mustJsonMarshal(result) != mustJsonMarshal(d.expect) {
			t.Errorf("expect %v, but %v", d.expect, result)
		}
	}
}

func mustJsonMarshal(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(data)
}

// {
//			"token0_symbol": "FRAX",
//			"token0_address": "0x853d955acef822db058eb8505911ed77f175b99e",
//			"token1_symbol": "USDC",
//			"token1_address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
//			"pair": "FRAX/USDC"
//		},

package main

import (
	"fmt"
	"sort"
)

type packet struct {
	version int
	typeID  int

	value int

	lengthTypeID int
	length       int
	packets      []packet
}

var m = map[byte][]int{
	'0': {0, 0, 0, 0},
	'1': {0, 0, 0, 1},
	'2': {0, 0, 1, 0},
	'3': {0, 0, 1, 1},
	'4': {0, 1, 0, 0},
	'5': {0, 1, 0, 1},
	'6': {0, 1, 1, 0},
	'7': {0, 1, 1, 1},
	'8': {1, 0, 0, 0},
	'9': {1, 0, 0, 1},
	'A': {1, 0, 1, 0},
	'B': {1, 0, 1, 1},
	'C': {1, 1, 0, 0},
	'D': {1, 1, 0, 1},
	'E': {1, 1, 1, 0},
	'F': {1, 1, 1, 1},
}

func main() {
	buf := make([]int, 0, len(input)*4)
	for _, b := range input {
		buf = append(buf, m[b]...)
	}
	pkt, _ := process(buf)

	fmt.Println(sumVersion(pkt))
	fmt.Println(calc(pkt))
}

func calc(pkt packet) int {
	switch pkt.typeID {
	case 0:
		var sum int
		for _, p := range pkt.packets {
			sum += calc(p)
		}
		return sum
	case 1:
		sum := 1
		for _, p := range pkt.packets {
			sum *= calc(p)
		}
		return sum
	case 2:
		vals := make([]int, 0)
		for _, p := range pkt.packets {
			vals = append(vals, calc(p))
		}
		sort.Ints(vals)
		return vals[0]
	case 3:
		vals := make([]int, 0)
		for _, p := range pkt.packets {
			vals = append(vals, calc(p))
		}
		sort.Ints(vals)
		return vals[len(vals)-1]
	case 4:
		return pkt.value
	case 5:
		a, b := calc(pkt.packets[0]), calc(pkt.packets[1])
		if a > b {
			return 1
		}
		return 0
	case 6:
		a, b := calc(pkt.packets[0]), calc(pkt.packets[1])
		if a < b {
			return 1
		}
		return 0

	case 7:
		a, b := calc(pkt.packets[0]), calc(pkt.packets[1])
		if a == b {
			return 1
		}
		return 0
	}
	panic(pkt)
}

func sumVersion(pkt packet) int {
	sum := pkt.version
	for _, p := range pkt.packets {
		sum += sumVersion(p)
	}
	return sum
}

func process(b []int) (packet, int) {
	pkt := packet{
		version: b2i(b[:3]),
		typeID:  b2i(b[3:6]),
	}
	if pkt.typeID == 4 {
		offset := 6
		var rawVal []int
		for {
			rawVal = append(rawVal, b[offset+1:offset+5]...)
			if b[offset] == 0 {
				offset += 5
				break
			}
			offset += 5
		}
		pkt.value = b2i(rawVal)
		return pkt, offset
	}
	pkt.lengthTypeID = b[6]
	var offset int
	if pkt.lengthTypeID == 0 {
		bits := b2i(b[7:22])
		offset = 22
		for sum := bits; sum > 0; {
			subPkt, cnt := process(b[offset:])
			pkt.packets = append(pkt.packets, subPkt)
			sum -= cnt
			offset += cnt
		}
		offset = 22 + bits
	} else {
		pktCount := b2i(b[7:18])
		offset = 18
		for i := 0; i < pktCount; i++ {
			subPkt, cnt := process(b[offset:])
			pkt.packets = append(pkt.packets, subPkt)
			offset += cnt
		}
	}
	return pkt, offset
}

func b2i(b []int) int {
	var sum int
	mul := 1
	for i := len(b) - 1; i >= 0; i-- {
		sum += mul * b[i]
		mul *= 2
	}
	return sum
}

var (
	input0 = []byte(`D2FE28`)
	input1 = []byte(`38006F45291200`)
	input  = []byte(`E20D41802B2984BD00540010F82D09E35880350D61A41D3004E5611E585F40159ED7AD7C90CF6BD6BE49C802DEB00525272CC1927752698693DA7C70029C0081002140096028C5400F6023C9C00D601ED88070070030005C2201448400E400F40400C400A50801E20004C1000809D14700B67676EE661137ADC64FF2BBAD745B3F2D69026335E92A0053533D78932A9DFE23AC7858C028920A973785338832CFA200F47C81D2BBBC7F9A9E1802FE00ACBA44F4D1E775DDC19C8054D93B7E72DBE7006AA200C41A8510980010D8731720CB80132918319804738AB3A8D3E773C4A4015A498E680292B1852E753E2B29D97F0DE6008CB3D4D031802D2853400D24DEAE0137AB8210051D24EB600844B95C56781B3004F002B99D8F635379EDE273AF26972D4A5610BA51004C12D1E25D802F32313239377B37100105343327E8031802B801AA00021D07231C2F10076184668693AC6600BCD83E8025231D752E5ADE311008A4EA092754596C6789727F069F99A4645008247D2579388DCF53558AE4B76B257200AAB80107947E94789FE76E36402868803F0D62743F00043A1646288800084C3F8971308032996A2BD8023292DF8BE467BB3790047F2572EF004A699E6164C013A007C62848DE91CC6DB459B6B40087E530AB31EE633BD23180393CBF36333038E011CBCE73C6FB098F4956112C98864EA1C2801D2D0F319802D60088002190620E479100622E4358952D84510074C0188CF0923410021F1CE1146E3006E3FC578EE600A4B6C4B002449C97E92449C97E92459796EB4FF874400A9A16100A26CEA6D0E5E5EC8841C9B8FE37109C99818023A00A4FD8BA531586BB8B1DC9AE080293B6972B7FA444285CC00AE492BC910C1697B5BDD8425409700562F471201186C0120004322B42489A200D4138A71AA796D00374978FE07B2314E99BFB6E909678A0`)
)

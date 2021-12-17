package main

import (
	"AdventOfCode/helpers"
	"fmt"
	"math"
	"strconv"
)

func convertHexToBinary(input string) (output string) {
	for i := 0; i < len(input); i++ {
		value, err := strconv.ParseInt(string(input[i]), 16, 64)
		if err == nil {
			bin := strconv.FormatInt(value, 2)
			for len(bin) < 4 {
				bin = "0" + bin
			}

			output += bin
		}
	}
	return
}

func parseVersion(binary string) (version int64, err error) {
	if len(binary) > 3 {
		version, err = strconv.ParseInt(binary[0:3], 2, 10)
		if err != nil {
			err = fmt.Errorf("Could not parse version from: " +
				binary[0:3] + ": " + err.Error())
		}
	}
	return
}

func parseTypeID(binary string) (typeid int64, err error) {
	if len(binary) > 3 {
		typeid, err = strconv.ParseInt(binary[0:3], 2, 10)
		if err != nil {
			err = fmt.Errorf("Could not parse type from: " +
				binary[0:3] + ": " + err.Error())
		}
	}
	return
}

func partOne(input string) (result int64) {
	binary := convertHexToBinary(input)
	for i := 0; i < len(binary); {
		version, err := parseVersion(binary[i:])
		if err == nil {
			result += version
			i += 3
		} else {
			fmt.Println(err.Error())
		}

		typeid, err := parseTypeID(binary[i:])
		if err == nil {
			i += 3
		} else {
			fmt.Println(err.Error())
		}

		if typeid == 4 {
			for i < len(binary) && string(binary[i]) == "1" {
				i += 5
			}
			i += 5
		} else {
			if i < len(binary) {
				length := string(binary[i])
				if length == "1" {
					i += 12
				} else {
					i += 16
				}
			}
		}

		if len(binary)-i < 4 {
			i = len(binary) + 1
		}
	}
	return
}

func handlePacket(packet string) (result int64, bitsUsed int) {
	typeid := packet[3:6]

	if typeid == "100" {
		number := ""
		i := 6
		for ; string(packet[i]) == "1"; i += 5 {
			number += packet[i+1 : i+5]
		}
		number += packet[i+1 : i+5]
		result, _ = strconv.ParseInt(number, 2, 64)
		bitsUsed = i + 5
		return
	}
	numBits := ""
	if string(packet[6]) == "1" {
		numBits = packet[7:18]
		bitsUsed = 18
	} else {
		numBits = packet[7:22]
		bitsUsed = 22
	}
	subPacketSize, _ := strconv.ParseInt(numBits, 2, 64)

	switch typeid {
	case "000": // sum
		if string(packet[6]) == "1" {
			for i := 0; i < int(subPacketSize); i++ {
				r, b := handlePacket(packet[bitsUsed:])
				result += r
				bitsUsed += b
			}
		} else {
			for i := 0; i < int(subPacketSize); {
				r, b := handlePacket(packet[bitsUsed+i:])
				result += r
				i += b
			}
			bitsUsed += int(subPacketSize)
		}
	case "001": // product
		result = 1
		if string(packet[6]) == "1" {
			for i := 0; i < int(subPacketSize); i++ {
				r, b := handlePacket(packet[bitsUsed:])
				result *= r
				bitsUsed += b
			}
		} else {
			for i := 0; i < int(subPacketSize); {
				r, b := handlePacket(packet[bitsUsed+i:])
				result *= r
				i += b
			}
			bitsUsed += int(subPacketSize)
		}
	case "010": // minimum
		result = int64(math.MaxInt64)
		if string(packet[6]) == "1" {
			for i := 0; i < int(subPacketSize); i++ {
				r, b := handlePacket(packet[bitsUsed:])
				if r < result {
					result = r
				}
				bitsUsed += b
			}
		} else {
			for i := 0; i < int(subPacketSize); {
				r, b := handlePacket(packet[bitsUsed+i:])
				if r < result {
					result = r
				}
				i += b
			}
			bitsUsed += int(subPacketSize)
		}
	case "011": // maximum
		result = -1
		if string(packet[6]) == "1" {
			for i := 0; i < int(subPacketSize); i++ {
				r, b := handlePacket(packet[bitsUsed:])
				if r > result {
					result = r
				}
				bitsUsed += b
			}
		} else {
			for i := 0; i < int(subPacketSize); {
				r, b := handlePacket(packet[bitsUsed+i:])
				if r > result {
					result = r
				}
				i += b
			}
			bitsUsed += int(subPacketSize)
		}
	case "101": // greater than
		r1, b := handlePacket(packet[bitsUsed:])
		bitsUsed += b
		r2, b := handlePacket(packet[bitsUsed:])
		bitsUsed += b
		if r1 > r2 {
			result = 1
		}
	case "110": // less than
		r1, b := handlePacket(packet[bitsUsed:])
		bitsUsed += b
		r2, b := handlePacket(packet[bitsUsed:])
		bitsUsed += b
		if r1 < r2 {
			result = 1
		}
	case "111": // equal
		r1, b := handlePacket(packet[bitsUsed:])
		bitsUsed += b
		r2, b := handlePacket(packet[bitsUsed:])
		bitsUsed += b
		if r1 == r2 {
			result = 1
		}
	}
	return
}

func partTwo(input string) (result int64) {
	binary := convertHexToBinary(input)
	result, _ = handlePacket(binary)
	return
}

func main() {
	//fmt.Println(partOne("8A004A801A8002F478") == 16)
	//fmt.Println(partOne("620080001611562C8802118E34") == 12)
	//fmt.Println(partOne("C0015000016115A2E0802F182340") == 23)
	//fmt.Println(partOne("A0016C880162017C3686B18A3D4780") == 31)
	//input, _ := helpers.ParseInputFile("input.txt")
	//fmt.Println(partOne(input[0]))
	//fmt.Println(partTwo("C200B40A82") == 3)
	//fmt.Println(partTwo("04005AC33890") == 54)
	//fmt.Println(partTwo("880086C3E88112") == 7)
	//fmt.Println(partTwo("CE00C43D881120") == 9)
	//fmt.Println(partTwo("D8005AC2A8F0") == 1)
	//fmt.Println(partTwo("F600BC2D8F") == 0)
	//fmt.Println(partTwo("9C005AC2F8F0") == 0)
	//fmt.Println(partTwo("9C0141080250320F1802104A08") == 1)
	input, _ := helpers.ParseInputFile("input.txt")
	fmt.Println(partTwo(input[0]))
}

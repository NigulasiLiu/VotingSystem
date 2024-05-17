package utils

import (
	"crypto/sha512"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const lambda = 256

func PG(bitArray []byte) []byte {
	extendedBitArray := make([]byte, len(bitArray)*2+2) // 创建扩展后的比特数组

	seed := int(bitArray[0])         // 种子值初始化为第一个比特位
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机数种子

	extendedIndex := 0 // 扩展比特数组的索引

	// 循环生成比特串
	for i := 0; i < len(bitArray); i++ {
		// 复制当前比特位
		extendedBitArray[extendedIndex] = bitArray[i]
		extendedIndex++

		// 更新种子值
		seed = (seed*9301 + 49297) % 233280

		// 生成一个随机比特位
		rnd := float64(seed) / 233280.0
		var randomBit byte
		if rnd >= 0.5 {
			randomBit = 1
		} else {
			randomBit = 0
		}

		// 将随机生成的比特位追加到扩展比特数组中
		extendedBitArray[extendedIndex] = randomBit
		extendedIndex++
		if i < 2 {
			if rnd >= 0.5 {
				randomBit = 0
			} else {
				randomBit = 1
			}
			extendedBitArray[extendedIndex] = randomBit
			extendedIndex++
		}
	}

	return extendedBitArray
}

func Getbit(bitArray []byte, pos int) byte {
	// 确保位置在比特数组范围内
	if pos < 0 || pos >= len(bitArray) {
		return 0
	}

	// 返回指定位置的比特值
	return bitArray[pos]
}

func Hash(bitArray []byte) []byte {
	targetLength := 4 * lambda
	if len(bitArray) >= targetLength {
		return bitArray[:targetLength] // 如果原始比特数组长度已经大于等于目标长度，直接返回原始比特数组的前 targetLength 位
	}

	repeatedBits := (targetLength + len(bitArray) - 1) / len(bitArray) // 计算需要重复几次原始比特数组才能达到目标长度
	extendedBitArray := make([]byte, repeatedBits*len(bitArray))       // 创建一个足够大的数组来存储重复的比特数组
	for i := 0; i < repeatedBits; i++ {
		copy(extendedBitArray[i*len(bitArray):], bitArray) // 将原始比特数组复制到扩展数组中
	}

	return extendedBitArray[:targetLength] // 截取目标长度的比特数组
}

func Convert(bitArray []byte) byte {
	return bitArray[0]
}

func toComplement(integer int, targetLength int) []byte {
	// 获取整数的绝对值
	absInteger := int(math.Abs(float64(integer)))

	// 将整数转换为二进制字符串
	binaryString := fmt.Sprintf("%b", absInteger)

	// 如果目标长度大于二进制字符串的长度，则在前面添加零
	for len(binaryString) < targetLength {
		binaryString = "0" + binaryString
	}

	// 如果目标长度小于二进制字符串的长度，则截取最后 targetLength 位
	if len(binaryString) > targetLength {
		binaryString = binaryString[len(binaryString)-targetLength:]
	}

	// 如果原始整数为负数，则计算其补码
	if integer < 0 {
		// 计算反码（将所有位取反）
		onesComplement := ""
		for _, bit := range binaryString {
			if bit == '0' {
				onesComplement += "1"
			} else {
				onesComplement += "0"
			}
		}

		// 将反码转换为数字
		onesComplementNumber := 0
		for _, bit := range onesComplement {
			onesComplementNumber = onesComplementNumber*2 + int(bit-'0')
		}

		// 计算补码（反码加1）
		twosComplementNumber := onesComplementNumber + 1

		// 将补码转换为二进制字符串
		binaryString = fmt.Sprintf("%b", twosComplementNumber)

		// 如果二进制字符串的长度超过目标长度，则截取最后 targetLength 位
		if len(binaryString) > targetLength {
			binaryString = binaryString[len(binaryString)-targetLength:]
		}
	}

	// 将二进制字符串转换为 []byte
	result := make([]byte, len(binaryString))
	for i, bit := range binaryString {
		result[i] = byte(bit - '0')
	}

	return result
}

func xor(bitArray1, bitArray2 []byte) []byte {
	// 获取两个比特数组的最大长度
	len1 := len(bitArray1)
	len2 := len(bitArray2)
	maxLength := len1
	if maxLength < len2 {
		maxLength = len2
	}

	// 给较短的数组补前导0
	if len1 < maxLength {
		bitArray1 = append(make([]byte, maxLength-len1), bitArray1...)
	}
	if len2 < maxLength {
		bitArray2 = append(make([]byte, maxLength-len2), bitArray2...)
	}

	// 创建结果数组
	result := make([]byte, maxLength)

	// 执行异或操作
	for i := 0; i < maxLength; i++ {
		result[i] = bitArray1[i] ^ bitArray2[i]
	}

	return result
}

func and(bit byte, bitArray []byte) byte {
	l := len(bitArray)
	if bit == byte(0) {
		return byte(0)
	}
	return bitArray[l-1]
}

func add(bitArray1, bitArray2 []byte) []byte {
	// 获取两个比特数组的最大长度
	len1 := len(bitArray1)
	len2 := len(bitArray2)
	maxLength := len1
	if maxLength < len2 {
		maxLength = len2
	}

	// 给较短的数组补前导0
	if len1 < maxLength {
		bitArray1 = append(make([]byte, maxLength-len1), bitArray1...)
	}
	if len2 < maxLength {
		bitArray2 = append(make([]byte, maxLength-len2), bitArray2...)
	}

	// 创建结果数组
	result := make([]byte, maxLength)

	// 定义进位标志
	carry := byte(0)

	// 从最低位开始逐位相加
	for i := maxLength - 1; i >= 0; i-- {
		sum := bitArray1[i] + bitArray2[i] + carry
		result[i] = sum % 2 // 当前位的和
		carry = sum / 2     // 计算进位
	}

	return result
}

func concat(arr1, arr2 []byte) []byte {
	// 创建结果数组，长度为两个输入数组长度之和
	result := make([]byte, len(arr1)+len(arr2))

	// 将 arr1 拷贝到结果数组的前半部分
	copy(result[:len(arr1)], arr1)

	// 将 arr2 拷贝到结果数组的后半部分
	copy(result[len(arr1):], arr2)

	return result
}

func toInt(bits []byte) int {
	result := 0
	for i := 0; i < len(bits); i++ {
		result = result << 1
		result += int(bits[i])
	}
	return result
}

func HashPrime(input []byte) []byte {
	// Compute the SHA-512 hash
	hash := sha512.Sum512(input)

	// Convert the hash to a bit array (512 bits)
	bitArray := make([]byte, 512)
	for i, b := range hash {
		for j := 0; j < 8; j++ {
			bitArray[i*8+j] = (b >> (7 - j)) & 1
		}
	}

	return bitArray
}

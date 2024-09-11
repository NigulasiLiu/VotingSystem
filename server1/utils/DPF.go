package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"server1/common"
	"server1/model"
)

func Eval(ID int, voteid int, k []byte) {
	// 获取数据库连接
	db := common.GetDB()

	// 计算 eta
	var eta int64
	db.Model(&model.Participate{}).Where("vote_id = ?", voteid).Count(&eta)
	if eta < 1 {
		log.Printf("无效的 eta 值：%d", eta)
		return
	}

	n := int(math.Ceil(math.Log2(float64(eta))))
	if n < 0 {
		log.Printf("计算出不合理的 n 值：%d", n)
		return
	}

	// 提取 s、t、CW、cs、pos
	const_s := k[:lambda]
	const_t := k[lambda]
	i := lambda + 1
	CW := make([][]byte, n+2, lambda+2)
	for j := 1; j <= n+1; j++ {
		CW[j] = k[i : i+lambda+2]
		i += lambda + 2
	}
	fmt.Println(CW[n+1])
	cs := k[i : i+4*lambda]
	i += 4 * lambda
	pos := k[i : i+8]

	// 初始化 outputs
	outputs := make([]*big.Int, int(eta))
	pi := make([]byte, 2*lambda)
	val := big.NewInt(0)
	VW := make([]byte, 4*lambda)

	for w := 0; w < int(eta); w++ {
		s := const_s
		t := const_t
		x := toComplement(big.NewInt(int64(w)), n)
		scw := make([]byte, lambda)
		tcw := make([]byte, 2)
		for j := 1; j <= n; j++ {
			scw = CW[j][:lambda]
			tcw[0] = CW[j][lambda]
			tcw[1] = CW[j][lambda+1]

			tmp := PG(s)
			S := make([][]byte, 2, lambda)
			T := make([]byte, 2)
			S[0] = tmp[:lambda]
			T[0] = tmp[lambda]
			S[1] = tmp[lambda+1 : 2*lambda+1]
			T[1] = tmp[2*lambda+1]

			if x[j-1] == 0 {
				if t == 1 {
					s = xor(S[0], scw)
				} else {
					s = S[0]
				}
				t = T[0] ^ (t * tcw[0])
			} else {
				if t == 1 {
					s = xor(S[1], scw)
				} else {
					s = S[1]
				}
				t = T[1] ^ (t * tcw[1])
			}
		}
		fmt.Println(s)
		y := Convert(s)
		fmt.Println(y)
		ycg := toInt(CW[n+1])
		fmt.Println(ycg)
		if t == 1 {
			y.Add(y, ycg)
		}
		y.Mul(y, big.NewInt(-1))
		fmt.Println(y)
		Pi := Hash(concat(x, s))
		t = Getbit(s, int(toInt(pos).Int64()))
		ver := HashPrime(cs)
		if t == 0 {
			ver = xor(ver, HashPrime(Pi))
		} else {
			ver = xor(ver, HashPrime(xor(Pi, cs)))
		}
		outputs[w] = y
		val.Add(val, y)
		pi = xor(pi, ver)
		VW = concat(pi, toComplement(val, 2*lambda))
	}

	var argument model.Argument

	// 查找并加载要更新的记录
	if err := db.Where("id = ?", ID).First(&argument).Error; err != nil {
		fmt.Println("Record not found:", err)
		return
	}

	var outputStrings []string
	for _, output := range outputs {
		outputStrings = append(outputStrings, output.String())
	}

	// Marshal the string array to JSON
	outputsJSON, err := json.Marshal(outputStrings)
	if err != nil {
		log.Fatal(err)
	}

	// 更新记录的特定字段
	argument.VW1 = VW
	argument.Outputs1 = string(outputsJSON)

	// 保存更改
	if err := db.Save(&argument).Error; err != nil {
		fmt.Println("Failed to update record:", err)
		return
	}

	fmt.Println("Record updated successfully")

}

func Tallying(voteid int, N int, eta int) []*big.Int {
	db := common.GetDB()

	// 查找与 voteid 匹配的记录
	var arguments []model.Argument
	if err := db.Where("vote_id = ?", voteid).Find(&arguments).Error; err != nil {
		fmt.Println("Failed to find arguments:", err)
		return nil
	}

	// 初始化 outs 数组
	outs := make([]*big.Int, eta)
	for i := range outs {
		outs[i] = big.NewInt(0)
	}

	// 遍历记录并进行累加
	for w := 0; w < N; w++ {
		var outputStrings []string
		if err := json.Unmarshal([]byte(arguments[w].Outputs1), &outputStrings); err != nil {
			fmt.Println("Failed to parse Outputs1 JSON:", err)
			return nil
		}

		// 调试输出解析后的字符串数组
		// fmt.Println("Parsed output strings:", outputStrings)

		// 转换为 big.Int 数组并累加
		for i, outputStr := range outputStrings {
			output := new(big.Int)
			if _, ok := output.SetString(outputStr, 10); !ok {
				fmt.Println("Failed to convert string to big.Int:", outputStr)
				return nil
			}
			outs[i].Add(outs[i], output)
		}
	}

	// 调试输出累加后的 outs 数组
	// fmt.Println("Tallying result:", outs)
	return outs
}

func Verify(voteid int, N int, outs []*big.Int) []*big.Int {
	db := common.GetDB()

	// 查找与 voteid 匹配的记录
	var arguments []model.Argument
	if err := db.Where("vote_id = ?", voteid).Find(&arguments).Error; err != nil {
		fmt.Println("Failed to find arguments:", err)
		return nil
	}

	// 实现减法
	sub := func(a []*big.Int, b []*big.Int) []*big.Int {
		result := make([]*big.Int, len(a))
		for i := range a {
			result[i] = new(big.Int).Sub(a[i], b[i])
		}
		return result
	}

	for w := 0; w < N; w++ {
		fmt.Printf("Processing argument %d\n", w)

		// 解析 Outputs1 的 JSON 字符串为字符串数组
		var outputStrings []string
		if err := json.Unmarshal([]byte(arguments[w].Outputs1), &outputStrings); err != nil {
			fmt.Println("Failed to parse Outputs1 JSON:", err)
			return nil
		}
		// fmt.Printf("Parsed Outputs1 JSON for argument %d: %v\n", w, outputStrings)

		// 转换为 big.Int 数组
		var outputs1 []*big.Int
		for _, outputStr := range outputStrings {
			output := new(big.Int)
			if _, ok := output.SetString(outputStr, 10); !ok {
				fmt.Println("Failed to convert string to big.Int:", outputStr)
				return nil
			}
			outputs1 = append(outputs1, output)
		}
		// fmt.Printf("Converted Outputs1 to big.Int array for argument %d: %v\n", w, outputs1)

		pi0 := arguments[w].VW0[:2*lambda]
		val0 := arguments[w].VW0[2*lambda]
		pi1 := arguments[w].VW1[:2*lambda]
		val1 := arguments[w].VW1[2*lambda]

		if !bytes.Equal(pi0, pi1) || (val0+val1 < 0) || (val0+val1 > 1) {
			outs = sub(outs, outputs1)
			fmt.Printf("Updated outs after subtraction for argument %d: %v\n", w, outs)
		}
	}

	return outs
}

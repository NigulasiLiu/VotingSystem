package utils

import (
	"bytes"
	"fmt"
	"math"
	"server0/common"
	"server0/model"
)

func Eval(ID uint, voteid int, k []byte) {
	// 获取数据库连接
	db := common.GetDB()

	// 计算 eta
	var eta int64
	db.Model(&model.Participate{}).Where("vote_id = ?", voteid).Count(&eta)

	// 计算 n
	n := int(math.Ceil(math.Log2(float64(eta))))

	// 提取 s、t、CW、cs、pos
	const_s := k[:lambda]
	const_t := k[lambda]
	i := lambda + 1
	CW := make([][]byte, n+2, lambda+2)
	for j := 1; j <= n+1; j++ {
		CW[j] = k[i : i+lambda+2]
		i += lambda + 2
	}
	cs := k[i : i+4*lambda]
	i += 4 * lambda
	pos := k[i : i+8]

	// 初始化 outputs
	outputs := make([]byte, int(eta))
	pi := make([]byte, 2*lambda)
	val := byte(0)
	VW := make([]byte, 2*lambda+1)

	for w := 0; w < int(eta); w++ {
		s := const_s
		t := const_t
		x := toComplement(w, n)
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
				s = xor(S[0], []byte{and(t, scw)})
				t = T[0] ^ (t & tcw[0])
			} else {
				s = xor(S[0], []byte{and(t, scw)})
				t = T[1] ^ (t & tcw[1])
			}
		}

		y := Convert(s) + and(t, CW[n+1])
		if y == byte(2) {
			y = byte(0)
		}
		Pi := Hash(concat(x, s))
		t = Getbit(s, toInt(pos))
		ver := HashPrime(cs)
		if t == 0 {
			ver = xor(ver, HashPrime(Pi))
		} else {
			ver = xor(ver, HashPrime(xor(Pi, cs)))
		}
		outputs[w] = y
		val += y
		pi = xor(pi, ver)
		VW = concat(pi, []byte{val})
	}

	var argument model.Argument

	// 查找并加载要更新的记录
	if err := db.Where("id = ?", ID).First(&argument).Error; err != nil {
		fmt.Println("Record not found:", err)
		return
	}

	// 更新记录的特定字段
	argument.VW0 = VW
	argument.Outputs0 = outputs

	// 保存更改
	if err := db.Save(&argument).Error; err != nil {
		fmt.Println("Failed to update record:", err)
		return
	}

	fmt.Println("Record updated successfully")

}

func Tallying(voteid int, N int, eta int) []int {
	db := common.GetDB()

	// 查找与 voteid 匹配的记录
	var arguments []model.Argument
	if err := db.Where("vote_id = ?", voteid).Find(&arguments).Error; err != nil {
		fmt.Println("Failed to find arguments:", err)
		return nil
	}

	// 初始化 outs 数组
	outs := make([]int, eta)

	// 遍历记录并进行累加
	for w := 0; w < N; w++ {
		for i := 0; i < eta; i++ {
			outs[i] += int(arguments[w].Outputs0[i])
		}
	}

	return outs
}

func Verify(voteid int, N int, outs []int) []int {
	db := common.GetDB()

	// 查找与 voteid 匹配的记录
	var arguments []model.Argument
	if err := db.Where("vote_id = ?", voteid).Find(&arguments).Error; err != nil {
		fmt.Println("Failed to find arguments:", err)
		return nil
	}

	// 实现减法
	sub := func(a []int, b []byte) []int {
		result := make([]int, len(a))
		for i := range a {
			result[i] = a[i] - int(b[i])
		}
		return result
	}

	for w := 0; w < N; w++ {
		pi0 := arguments[w].VW0[:2*lambda]
		val0 := arguments[w].VW0[2*lambda]
		pi1 := arguments[w].VW1[:2*lambda]
		val1 := arguments[w].VW1[2*lambda]

		if !bytes.Equal(pi0, pi1) || (val0+val1 < 0) || (val0+val1 > 1) {
			outs = sub(outs, arguments[w].Outputs0)
		}
	}

	return outs
}

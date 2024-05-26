/* eslint-disable no-undef */
/* eslint-disable no-param-reassign */
/* eslint-disable no-bitwise */
/* eslint-disable no-plusplus */
const lambda = 128;

function PG(bitArray) {
  const extendedBitArray = new Uint8Array(bitArray.length * 2 + 2); // 扩展后的比特数组

  let seed = bitArray[0]; // 种子值初始化为第一个比特位
  let extendedIndex = 0; // 扩展比特数组的索引

  // 循环生成比特串
  for (let i = 0; i < bitArray.length; i++) {
    // 复制当前比特位
    extendedBitArray[extendedIndex++] = bitArray[i];

    // 更新种子值
    seed = (seed * 9301 + 49297) % 233280;

    // 生成一个随机比特位
    const rnd = seed / 233280.0;
    let randomBit = rnd >= 0.5 ? 1 : 0;

    // 将随机生成的比特位追加到扩展比特数组中
    extendedBitArray[extendedIndex++] = randomBit;
    if (i < 2) {
      randomBit = rnd >= 0.5 ? 0 : 1;
      extendedBitArray[extendedIndex++] = randomBit;
    }
  }

  return extendedBitArray;
}

function Getbit(bitArray, pos) {
  // 确保位置在比特数组范围内
  if (pos < 0 || pos >= bitArray.length) {
    return false;
  }

  // 返回指定位置的比特值
  return bitArray[pos] === 1;
}

function Hash(bitArray) {
  const targetLength = 4 * lambda;
  if (bitArray.length >= targetLength) {
    return bitArray.slice(0, targetLength); // 如果原始比特数组长度已经大于等于目标长度，直接返回原始比特数组的前 targetLength 位
  }
  const repeatedBits = Math.ceil(targetLength / bitArray.length); // 计算需要重复几次原始比特数组才能达到目标长度
  const extendedBitArray = new Uint8Array(bitArray.length * repeatedBits); // 重复原始比特数组
  for (let i = 0; i < repeatedBits; i++) {
    extendedBitArray.set(bitArray, i * bitArray.length);
  }
  return extendedBitArray.slice(0, targetLength); // 截取目标长度的比特数组
}

function toInt(bitArray) {
  // 确保输入的数组不为空
  if (!bitArray || bitArray.length === 0) {
    return 0n;
  }

  let result = 0n;
  for (let i = 0; i < bitArray.length; i++) {
    // 将比特位转换为整数并累加到结果中
    result = (result << 1n) | BigInt(bitArray[i]);
  }
  return result;
}

function toComplement(integer, targetLength) {
  // 获取整数的绝对值
  const absInteger = integer >= 0n ? integer : 0n - integer;

  // 将整数转换为二进制字符串
  let binaryString = absInteger.toString(2);

  // 如果目标长度大于二进制字符串的长度，则在前面添加零
  if (targetLength > binaryString.length) {
    binaryString = binaryString.padStart(targetLength, '0');
  }

  // 如果目标长度小于二进制字符串的长度，则截取最后 targetLength 位
  if (targetLength < binaryString.length) {
    binaryString = binaryString.slice(-targetLength);
  }

  // 如果原始整数为负数，则计算其补码
  if (integer < 0n) {
    // 计算反码（将所有位取反）
    let onesComplement = '';
    for (let i = 0; i < binaryString.length; i++) {
      onesComplement += binaryString[i] === '0' ? '1' : '0';
    }

    // 将反码转换为数字
    const onesComplementNumber = BigInt(`0b${onesComplement}`);

    // 计算补码（反码加1）
    const twosComplementNumber = onesComplementNumber + 1n;

    // 将补码转换为二进制字符串
    binaryString = twosComplementNumber.toString(2);

    // 如果二进制字符串的长度超过目标长度，则截取最后 targetLength 位
    if (binaryString.length > targetLength) {
      binaryString = binaryString.slice(-targetLength);
    }
  }

  return binaryString.split('').map(Number);
}

function ConvertCW(beta, t1, s0, s1) {
  if (t1 === 0) {
    const result = toComplement(BigInt(beta) - toInt(s0) + toInt(s1), lambda + 2);
    return result;
  }
  const result = toComplement((-1n) * (BigInt(beta) - toInt(s0) + toInt(s1)), lambda + 2);
  return result;
}

function generateRandomBitArray(targetLength) {
  const bitArray = new Uint8Array(targetLength);
  for (let i = 0; i < targetLength; i++) {
    // 生成随机的 0 或 1
    const randomBit = Math.random() < 0.5 ? 0 : 1;
    bitArray[i] = randomBit;
  }
  return bitArray;
}

function xor(bitArray1, bitArray2) {
  // 获取两个比特数组的长度
  const maxLength = Math.max(bitArray1.length, bitArray2.length);

  // 创建结果数组
  const result = new Uint8Array(maxLength);

  // 创建填充后的比特数组
  const paddedBitArray1 = new Uint8Array(maxLength);
  const paddedBitArray2 = new Uint8Array(maxLength);

  // 前导零填充
  paddedBitArray1.set(bitArray1, maxLength - bitArray1.length);
  paddedBitArray2.set(bitArray2, maxLength - bitArray2.length);

  // 执行异或操作
  for (let i = 0; i < maxLength; i++) {
    result[i] = paddedBitArray1[i] ^ paddedBitArray2[i];
  }

  return result;
}

function and(bitArray1, bitArray2) {
  const result = new Uint8Array(1);
  if (bitArray1 === 0) {
    result[0] = 0;
    return result;
  }
  result[0] = bitArray2[bitArray2.length - 1];
  return result;
}

function concat(arr1, arr2) {
  const result = new Uint8Array(arr1.length + arr2.length);
  result.set(arr1, 0);
  result.set(arr2, arr1.length);
  return result;
}

function sub(bitArray1, bitArray2) {
  // 确保输入的比特数组不为空
  if (!bitArray1 || !bitArray2 || bitArray1.length === 0 || bitArray2.length === 0) {
    throw new Error('Input bit arrays must be defined and non-empty');
  }

  // 获取比特数组的长度
  const length1 = bitArray1.length;
  const length2 = bitArray2.length;

  // 补前导0使两个比特数组的长度相同
  const maxLength = Math.max(length1, length2);
  const extendedBitArray1 = new Uint8Array(maxLength);
  const extendedBitArray2 = new Uint8Array(maxLength);

  extendedBitArray1.set(bitArray1, maxLength - length1);
  extendedBitArray2.set(bitArray2, maxLength - length2);

  // 创建结果数组
  const result = new Int8Array(maxLength);

  // 初始化借位为0
  let borrow = 0;

  // 从最低位开始逐位相减
  for (let i = maxLength - 1; i >= 0; i--) {
    let diff = extendedBitArray1[i] - extendedBitArray2[i] - borrow; // 当前位的差
    if (diff < 0) {
      diff += 2; // 如果差小于0，则加上2
      borrow = 1; // 更新借位
    } else {
      borrow = 0; // 如果差大于等于0，则借位清零
    }
    result[i] = diff; // 将差作为当前位的值
  }

  return result;
}

export default {
  PG,
  Getbit,
  Hash,
  ConvertCW,
  toComplement,
  generateRandomBitArray,
  xor,
  and,
  concat,
  sub,
  toInt,
};

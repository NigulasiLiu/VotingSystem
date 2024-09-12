/* eslint-disable no-bitwise */
/* eslint-disable no-plusplus */
import votedService from '@/service/votedService';
import share from './share';

const lambda = 128;

function genRandom(n) {
  return share.generateRandomBitArray(lambda + 1 + (n + 1) * (lambda + 2) + 4 * lambda + 8);
}

function dpfGen(a, beta, eta, id) {
  return new Promise((resolve) => {
    let k0 = new Uint8Array(lambda);
    let k1 = new Uint8Array(lambda);
    const n = Math.ceil(Math.log2(eta));
    if (beta !== 0) {
      const alpha = share.toComplement(a, n);
      let s0 = share.generateRandomBitArray(lambda);
      let s1 = share.generateRandomBitArray(lambda);
      let t0 = Math.random() < 0.5 ? 0 : 1;
      let t1 = t0 === 1 ? 0 : 1;

      k0 = share.concat(s0, share.toComplement(t0, 1));
      k1 = share.concat(s1, share.toComplement(t1, 1));

      for (let j = 0; j < n; j++) {
        const tmp0 = share.PG(s0);
        const S0 = [tmp0.slice(0, lambda), tmp0.slice(lambda + 1, 2 * lambda + 1)];
        const T0 = [tmp0[lambda], tmp0[2 * lambda + 1]];
        const tmp1 = share.PG(s1);
        const S1 = [tmp1.slice(0, lambda), tmp1.slice(lambda + 1, 2 * lambda + 1)];
        const T1 = [tmp1[lambda], tmp1[2 * lambda + 1]];
        const keep = alpha[j];
        const lose = keep === 1 ? 0 : 1;
        const scw = share.xor(S0[lose], S1[lose]);
        const tcw = [T0[0] ^ T1[0] ^ alpha[j] ^ 1, T0[1] ^ T1[1] ^ alpha[j]];

        k0 = share.concat(k0, share.concat(scw, share.toComplement(tcw[0], 1)));
        k1 = share.concat(k1, share.concat(scw, share.toComplement(tcw[1], 1)));

        s0 = t0 === 0 ? S0[keep] : share.xor(S0[keep], scw);
        t0 = T0[keep] ^ (t0 * tcw[keep]);
        s1 = t1 === 0 ? S1[keep] : share.xor(S1[keep], scw);
        t1 = T1[keep] ^ (t1 * tcw[keep]);
      }

      const CW = share.ConvertCW(beta, t1, s0, s1);
      k0 = share.concat(k0, CW);
      k1 = share.concat(k1, CW);
      const pi = [share.Hash(share.concat(alpha, s0)), share.Hash(share.concat(alpha, s1))];
      const cs = share.xor(pi[0], pi[1]);
      k0 = share.concat(k0, cs);
      k1 = share.concat(k1, cs);
      let pos = Math.floor(Math.random() * n);
      while (share.Getbit(s0, pos) === share.Getbit(s1, pos)) pos = Math.floor(Math.random() * lambda);
      k0 = share.concat(k0, share.toComplement(pos, 8));
      k1 = share.concat(k1, share.toComplement(pos, 8));

      console.log('Sending vote0 with id:', id, 'and k0:', k0);
      votedService.vote0({ id, k: k0 })
        .then((res0) => {
          console.log('Sending vote1 with id:', id, 'and k1:', k1, 'and nid:', res0.data.data.id);
          return votedService.vote1({ id, k: k1, nid: res0.data.data.id });
        })
        .then(() => resolve(true))
        .catch((error) => {
          console.error('Error during voting:', error);
          resolve(false);
        });
    } else {
      // 使用 genRandom 生成随机键，如果beta为0
      k0 = genRandom(n);
      k1 = k0;
      resolve(true);
    }
  });
}

export default {
  dpfGen,
};

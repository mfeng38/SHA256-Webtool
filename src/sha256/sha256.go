//Psudocode from https://en.wikipedia.org/wiki/SHA-2 used as a reference.
package sha256

import(
  //"fmt"
  //"encoding/hex"
  "encoding/binary"
)

var h0 uint32 = 0x6a09e667
var h1 uint32 = 0xbb67ae85
var h2 uint32 = 0x3c6ef372
var h3 uint32 = 0xa54ff53a
var h4 uint32 = 0x510e527f
var h5 uint32 = 0x9b05688c
var h6 uint32 = 0x1f83d9ab
var h7 uint32 = 0x5be0cd19

var roundConstant = [...]uint32{0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
   0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
   0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
   0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
   0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
   0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
   0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
   0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}

func CalcPadding(b []byte) []byte{
  k := 512 - (((len(b)*8) + 1 + 64) % 512)
  var temp []byte
  temp = append(temp, 0x80)
  for i := 0; i < k-7; i+=8 {
    temp = append(temp, 0x00)
  }
  //fmt.Println(hex.EncodeToString(temp))
  l := make([]byte, 8)
  binary.BigEndian.PutUint64(l, uint64(len(b)*8))
  temp = append(temp, l ...)
  //fmt.Println(hex.EncodeToString(temp))
  b = append(b, temp...)
  //fmt.Println(len(b)*8)
  return b
}

func CalcDigest(b []byte) []byte {
  b = CalcPadding(b)
  var chunks [][]byte
  for i := 0; i < len(b); i += 64 {
    chunks = append(chunks, b[i:i+64])
  }
  //fmt.Printf("%x\n", chunks)
  for j := 0; j < len(chunks); j++ {
    w := make([]uint32, 64)
    for m := 0; m < 64; m += 4 {
      w[m/4] = binary.BigEndian.Uint32(chunks[j][m:(m+4)])
    }
    var s0, s1 uint32
    for o := 16; o < len(w); o++ {
      s0 = RightRotate(w[o-15], 7) ^ RightRotate(w[o-15], 18) ^ (w[o-15] >> 3)
      //fmt.Printf("%x", s0)
      s1 = RightRotate(w[o-2], 17) ^ RightRotate(w[o-2], 19) ^ (w[o-2] >> 10)
      w[o] = w[o-16] + s0 + w[o-7] + s1
    }
    var a, b, c, d, e, f, g, h uint32 = h0, h1, h2, h3, h4, h5, h6, h7
    //fmt.Printf("%x %x %x %x %x %x %x %x\n", a,b,c,d,e,f,g,h)
    var S0, S1, ch, maj, temp1, temp2 uint32
    for q := 0; q < 64; q++ {
      S1 = RightRotate(e, 6) ^ RightRotate(e, 11) ^ RightRotate(e, 25)
      ch = (e & f) ^ ((^e) & g)
      temp1 = h + S1 + ch + roundConstant[q] + w[q]
      S0 = RightRotate(a, 2) ^ RightRotate(a, 13) ^ RightRotate(a, 22)
      maj = (a & b) ^ (a & c) ^ (b & c)
      temp2 = S0 + maj

      h = g
      g = f
      f = e
      e = d + temp1
      d = c
      c = b
      b = a
      a = temp1 + temp2
    }
    //fmt.Printf("\na: %x\n", a)
    h0 = h0 + a
    h1 = h1 + b
    h2 = h2 + c
    h3 = h3 + d
    h4 = h4 + e
    h5 = h5 + f
    h6 = h6 + g
    h7 = h7 + h
  }
  tempDigest := []uint32{h0, h1, h2, h3, h4, h5, h6, h7}
  digest := make([]byte, 32)
  for i := 0; i < len(tempDigest); i++ {
    binary.BigEndian.PutUint32(digest[i*4:], tempDigest[i])
  }
  //fmt.Printf("%08x\n", digest)
  return digest
}

func RightRotate(b uint32, n uint32) uint32{
  //fmt.Printf("%032b\n", (b >> n)|(b << (32-n)))
  return (b >> n)|(b << (32-n))
}

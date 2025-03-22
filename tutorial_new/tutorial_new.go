package main

import (
	//"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"example.com/util"
)

var DBSize uint64
var ChunkSize uint64
var ChunkNum uint64
var DB []uint64

type LocalHint struct {
	key             util.PrfKey
	parity          uint64
	programmedPoint uint64
	isProgrammed    bool
}

// Elem returns the element in the chunkID-th chunk of the hint. It takes care of the case when the hint is programmed.
// 这里可以理解为一个PRP。
func Elem(hint *LocalHint, chunkId uint64) uint64 {

	//fmt.Println(hint.isProgrammed, chunkId, hint.programmedPoint, ChunkSize, hint.isProgrammed && chunkId == hint.programmedPoint/ChunkSize)

	if hint.isProgrammed && chunkId == hint.programmedPoint/ChunkSize {
		//这个语句块是什么意思呢？就是我们将backupHint添加进来的时候，我们直接返回x，注意x是index，DB[x]才是值。
		return hint.programmedPoint
	} else {
		return util.PRFEval(&hint.key, chunkId)%ChunkSize + chunkId*ChunkSize
	}
}

func PIR() {
	// Suppose there's a public DB.
	DB = make([]uint64, DBSize)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < int(DBSize); i++ {
		DB[i] = rng.Uint64() //uint64(i) //
		//log.Printf("DB[%d]: %d", i, DB[i])
	}

	// setup the parameters
	ChunkSize = uint64(math.Sqrt(float64(DBSize)))
	ChunkNum = uint64(math.Ceil(float64(DBSize) / float64(ChunkSize)))
	log.Printf("DBSize: %d, ChunkSize: %d, ChunkNum: %d", DBSize, ChunkSize, ChunkNum)

	// The following is the client side algorithm.
	Q := uint64(math.Sqrt(float64(DBSize)) * math.Log(float64(DBSize)))    //\sqrt{DBsize} * log_e^DBsize=4*2.9, Q代表查询次数
	M1 := 4 * uint64(math.Sqrt(float64(DBSize))*math.Log(float64(DBSize))) //4 * \sqrt{DBsize} * log_e^DBsize=4*4*2.9
	M2 := 4 * uint64(math.Log(float64(DBSize)))                            //4 * log_e^DBsize
	log.Printf("Q: %d, M1: %d, M2: %d", Q, M1, M2)

	//Setup Phase
	//The client first samples the hints
	primaryHints := make([]LocalHint, M1)             //primaryHints 有M1个
	replacementIndices := make([]uint64, M2*ChunkNum) //replacement 每一个ChunkNum有M2个
	replacementValues := make([]uint64, M2*ChunkNum)  //backupHints 每一个ChunkNum有M2个
	backupHints := make([]LocalHint, M2*ChunkNum)

	//为每一个primaryHints生成prfkey
	for i := uint64(0); i < M1; i++ {
		primaryHints[i] = LocalHint{util.RandKey(rng), 0, 0, false}
	}

	//为每一个backupHints生成prfkey
	for i := uint64(0); i < M2*ChunkNum; i++ {
		backupHints[i] = LocalHint{util.RandKey(rng), 0, 0, false}
	}

	//The client streamingly downloads the chunks from the server
	for i := uint64(0); i < ChunkNum; i++ {
		// suppose the client receives the i-th chunk, DB[i*ChunkSize:(i+1)*ChunkSize]
		for j := uint64(0); j < M1; j++ {

			//temp := primaryHints[j].parity
			primaryHints[j].parity ^= DB[Elem(&primaryHints[j], i)] // XOR操作
			//fmt.Printf("i= %d, j= %d : %d, %d, %d\n", i, j, temp, DB[Elem(&primaryHints[j], i)], primaryHints[j].parity)
		}

		//这里多了一个步骤就是第i个去掉了。
		for j := uint64(0); j < M2*ChunkNum; j++ {
			if j/M2 != i {
				backupHints[j].parity ^= DB[Elem(&backupHints[j], i)]
			}
		}

		//随机替换了
		for j := i * M2; j < (i+1)*M2; j++ {
			ind := rng.Uint64()%ChunkSize + i*ChunkSize
			replacementIndices[j] = ind
			replacementValues[j] = DB[ind]
		}
	}

	//Online Query Phase
	localCache := make(map[uint64]uint64)
	consumedReplacementNum := make([]uint64, ChunkNum)
	consumedHintNum := make([]uint64, ChunkNum)
	for q := uint64(0); q < Q; q++ {
		// just do random query for now
		x := rng.Uint64() % DBSize

		// make sure x is not in the local cache
		for {
			if _, ok := localCache[x]; ok == false {
				break
			}
			x = rng.Uint64() % DBSize
		}
		//fmt.Println("x: ", x)

		//找到x所在的chunk
		chunkId := x / ChunkSize

		//M1是primarytable的数量，从primarytable中找到hintID
		//只需要再这个chunk中判断就好了。
		hitId := uint64(999999999)
		for i := uint64(0); i < M1; i++ {
			if Elem(&primaryHints[i], chunkId) == x {
				hitId = i
				break
			}
		}
		if hitId == uint64(999999999) {
			log.Fatalf("Error: cannot find the hitId")
		}

		expandedSet := make([]uint64, ChunkNum)
		for i := uint64(0); i < ChunkNum; i++ {
			expandedSet[i] = Elem(&primaryHints[hitId], i)
			//fmt.Print(expandedSet[i], ", ")
		}
		//fmt.Println()

		// edit the expandedSet
		replacementInd := uint64(0)
		replacementVal := uint64(0)
		//某一个chunk的replacement<M2
		if consumedReplacementNum[chunkId] < M2 {
			// fetch the next unconsumed replacement pair
			tmp := consumedReplacementNum[chunkId] + chunkId*M2
			replacementInd = replacementIndices[tmp] //这里调用了我们的算法。
			replacementVal = replacementValues[tmp]
			consumedReplacementNum[chunkId]++

			expandedSet[chunkId] = replacementInd //放到了expandset中
		} else {
			log.Fatalf("Not enough replacement values")
		}

		/********** This is the server side algorithm **********/
		parity := uint64(0)
		for _, index := range expandedSet {
			parity ^= DB[index]
		}
		/********** The server side algorithm is done. The client receives parity. **********/

		// Upon receiving the parity, the client computes the answer
		answer := parity ^ primaryHints[hitId].parity ^ replacementVal

		// This verification only happens in this demo experiment.
		if answer != DB[x] {
			log.Fatalf("Error: answer is not correct")
		}

		// update the local cache
		localCache[x] = answer

		//将backup中更新到primaryHints中
		// refresh the hint
		if consumedHintNum[chunkId] < M2 {
			primaryHints[hitId] = backupHints[chunkId*M2+consumedHintNum[chunkId]]
			primaryHints[hitId].isProgrammed = true
			primaryHints[hitId].programmedPoint = x
			primaryHints[hitId].parity ^= answer
			consumedHintNum[chunkId]++
		} else {
			log.Fatalf("Not enough backup hints")
		}
	}
	log.Printf("PIR finished successfully")
}

func main() {
	DBSize = 256 // please make sure DBSize is a perfect square
	PIR()
}

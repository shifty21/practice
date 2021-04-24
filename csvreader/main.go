package main

import (
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"unsafe"

	"github.com/gocarina/gocsv"
)

type ROW struct {
	Time float32 `csv:"Time"` // .csv column headers
	GX   float64 `csv:"AccX"`
	GY   float64 `csv:"AccY"`
	GZ   float64 `csv:"AccZ"` // .csv column headers
	AX   float64 `csv:"P"`
	AY   float64 `csv:"Q"`
	AZ   float64 `csv:"R"`
}

type DATA struct {
	Time float64 `csv:"Time"` // .csv column headers
	GX   int32   `csv:"AccX"`
	GY   int32   `csv:"AccY"`
	GZ   int32   `csv:"AccZ"` // .csv column headers
	AX   int32   `csv:"P"`
	AY   int32   `csv:"Q"`
	AZ   int32   `csv:"R"`
}

type Header struct {
	szHeader           []rune
	bIsIntelOrMotorola int8
	dVersionNumber     float64
	bDeltaTheta        int32
	bDeltaVelocity     int32
	dDataRateHz        float64
	dGyroScaleFactor   float64
	dAccelScaleFactor  float64
	iUtcOrGpsTime      int32
	iRcvTimeOrCorrTime int32
	dTimeTagBias       float64
	szImuName          []rune
	reserved1          []uint8
	szProgramName      []rune
	tCreate            []int8
	bLeverArmValid     bool
	lXoffset           int32
	lYoffset           int32
	lZoffset           int32
	Reserved           []int8
}

func main() {
	in, err := os.Open("testIMU.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	rows := []*ROW{}
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use pipe as delimiter
	})
	if err := gocsv.UnmarshalFile(in, &rows); err != nil {
		panic(err)
	}
	writeFile(rows)
}

func writeFile(rows []*ROW) {
	file, err := os.Create("test.bin")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	szImuName := make([]rune, 32)
	text := "ADIS16488"
	for pos, char := range text {
		szImuName[pos] = char
	}
	szProgramName := make([]rune, 32)
	text1 := "Decoder"
	for pos, char := range text1 {
		szProgramName[pos] = char
	}
	header := &Header{
		szHeader:           []rune("$IMURAW\000"),
		bIsIntelOrMotorola: 1,
		dVersionNumber:     8.7,
		bDeltaTheta:        2,
		bDeltaVelocity:     3,
		dDataRateHz:        100,
		dGyroScaleFactor:   0.00000001,
		dAccelScaleFactor:  0.00000001,
		iUtcOrGpsTime:      1,
		iRcvTimeOrCorrTime: 1,
		dTimeTagBias:       1,
		szImuName:          szImuName,
		reserved1:          make([]uint8, 1),
		szProgramName:      szProgramName,
		tCreate:            make([]int8, 12),
		bLeverArmValid:     false,
		lXoffset:           0,
		lYoffset:           0,
		lZoffset:           0,
		Reserved:           make([]int8, 354),
	}

	var bin_buf bytes.Buffer

	log.Printf("Header %v", header)
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(header); err != nil {
		fmt.Println(err)
	}
	// dec := gob.NewDecoder(buf)
	// var s2 Header
	// if err := dec.Decode(&s2); err != nil {
	// 	fmt.Println(err)
	// }
	var test_buf bytes.Buffer
	err = binary.Write(&test_buf, binary.LittleEndian, header.bIsIntelOrMotorola)
	if err != nil {
		log.Printf("Error while converting dVersionNumber into bytes ")
	}
	fmt.Printf("%+v\n", buf.Cap())
	fmt.Printf("version=%v\n", test_buf.Cap())
	fmt.Printf("testbuf=%v\n", unsafe.Sizeof(make([]rune, 245)))
	writeNextBytes(file, buf.Bytes())

	for _, v := range rows {
		data := &DATA{
			Time: float64(v.Time),
			GX:   int32(v.GX * float64(header.dGyroScaleFactor)),
			GY:   int32(v.GY * float64(header.dGyroScaleFactor)),
			GZ:   int32(v.GZ * float64(header.dGyroScaleFactor)),
			AX:   int32(v.AX * float64(header.dGyroScaleFactor)),
			AY:   int32(v.AY * float64(header.dGyroScaleFactor)),
			AZ:   int32(v.AZ * float64(header.dGyroScaleFactor)),
		}
		log.Println(data)
		log.Println(v)
		err := binary.Write(&bin_buf, binary.BigEndian, data)
		if err != nil {
			log.Printf("Error while writing row %v", err)
		}
		log.Printf("Size of buf %v", bin_buf.Len())
		writeNextBytes(file, bin_buf.Bytes())
		bin_buf.Reset()
	}
}
func writeNextBytes(file *os.File, bytes []byte) {

	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

}

package dashbase

import (
	"bytes"
	"encoding/binary"

	"github.com/linkedin/goavro"
)

// See http://avro.apache.org/docs/1.8.2/spec.html#schema_fingerprints
// fingerprint -1959126995677700088L

const (
	emptyCRC64 uint64 = 0xc15d213aa4d7a795
	avroSchema string = `{"name":"io.dashbase.avro.DashbaseEvent","type":"record","fields":[{"name":"timeInMillis","type":"long"},{"name":"metaColumns","type":{"type":"map","values":"string"}},{"name":"numberColumns","type":{"type":"map","values":"double"}},{"name":"textColumns","type":{"type":"map","values":"string"}},{"name":"idColumns","type":{"type":"map","values":"string"}},{"name":"omitPayload","type":"boolean"},{"name": "raw", "type":["null", "string"],"default":"null"}]}`
)

var (
	fingerprint int64 = -1959126995677700088;
)

type Avro struct {
	schemaChecksum uint64
	codec          *goavro.Codec
}

type Event struct {
	TimeInMillis  int64
	MetaColumns   map[string]string
	TextColumns   map[string]string
	NumberColumns map[string]float64
	IdColumns     map[string]string
	Raw           string
	OmitPayload   bool
}

func makeAvroCodec() *goavro.Codec {
	codec, err := goavro.NewCodec(avroSchema)
	if err != nil {
		panic(err)
	}
	return codec
}

func getAvroCRC64(buf []byte) uint64 {
	var table []uint64
	table = make([]uint64, 256)
	for i := 0; i < 256; i++ {
		fp := uint64(i)
		for j := 0; j < 8; j++ {
			fp = (fp >> 1) ^ (emptyCRC64 & -(fp & 1))
		}
		table[i] = fp
	}
	fp := emptyCRC64
	for _, val := range buf {
		fp = (fp >> 8) ^ table[int(fp^uint64(val))&0xff]
	}
	return fp
}

func (a *Avro) Encode(event Event) ([]byte, error) {
	rawEvent := make(map[string]interface{})
	rawEvent["timeInMillis"] = event.TimeInMillis
	rawEvent["metaColumns"] = event.MetaColumns
	rawEvent["textColumns"] = event.TextColumns
	rawEvent["numberColumns"] = event.NumberColumns
	rawEvent["idColumns"] = event.IdColumns
	rawEvent["omitPayload"] = event.OmitPayload
	rawEvent["raw"] = event.Raw
	if event.Raw == "" {
		rawEvent["raw"] = goavro.Union("null", nil)
	}
	body, err := a.codec.BinaryFromNative(nil, rawEvent)
	if err != nil {
		return nil, err
	}

	message := new(bytes.Buffer)
	message.Write([]byte{0xC3, 0x01})
	binary.Write(message, binary.LittleEndian, uint64(fingerprint))
	message.Write(body)

	return message.Bytes(), nil
}

func NewAvro() *Avro {
	bs := make([]byte, 8)
	var num int64 = -1959126995677700088;
	binary.LittleEndian.PutUint64(bs, uint64(num))

	return &Avro{
		codec: makeAvroCodec(),
		//schemaChecksum: getAvroCRC64([]byte(avroSchema)),
	}
}

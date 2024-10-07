package raw

import (
	"encoding/binary"
	"io"
	"time"

	"github.com/SladeThe/checked-go/must"
	"github.com/SladeThe/yav"
)

func handleErrors[V yav.Validatable](r *Remote, v V) (V, error) {
	if err := r.Error(); err != nil {
		return v, err
	}

	return v, v.Validate()
}

func (r *Remote) Error() error {
	if err := r.err.Load(); err != nil {
		return *err
	}

	return nil
}

func (r *Remote) setError(err error) {
	if err == nil {
		r.err.Store(nil)
	}

	r.err.Store(&err)
}

func (r *Remote) flush() error {
	return r.Error()
}

func (r *Remote) setDeadline(deadline time.Duration) {
	r.setError(r.conn.SetDeadline(time.Now().Add(deadline)))
}

func (r *Remote) readByte() byte {
	if r.Error() != nil {
		return 0
	}

	v := []byte{0}

	if _, err := io.ReadFull(r.reader, v); err != nil {
		r.setError(err)
	}

	return v[0]
}

func (r *Remote) writeByte(v byte) {
	if r.Error() != nil {
		return
	}

	if _, err := r.writer.Write([]byte{v}); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readBool() bool {
	return r.readByte() != 0
}

func (r *Remote) writeBool(v bool) {
	if v {
		r.writeByte(1)
	} else {
		r.writeByte(0)
	}
}

func (r *Remote) readInt16() int16 {
	if r.Error() != nil {
		return 0
	}

	var v int16

	if err := binary.Read(r.reader, byteOrder, &v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeInt16(v int16) {
	if err := binary.Write(r.writer, byteOrder, v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readUint16() uint16 {
	if r.Error() != nil {
		return 0
	}

	var v uint16

	if err := binary.Read(r.reader, byteOrder, &v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeUint16(v uint16) {
	if err := binary.Write(r.writer, byteOrder, v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readInt32() int32 {
	if r.Error() != nil {
		return 0
	}

	var v int32

	if err := binary.Read(r.reader, byteOrder, &v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeInt32(v int32) {
	if err := binary.Write(r.writer, byteOrder, v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readUint32() uint32 {
	if r.Error() != nil {
		return 0
	}

	var v uint32

	if err := binary.Read(r.reader, byteOrder, &v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeUint32(v uint32) {
	if err := binary.Write(r.writer, byteOrder, v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readInt64() int64 {
	if r.Error() != nil {
		return 0
	}

	var v int64

	if err := binary.Read(r.reader, byteOrder, &v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeInt64(v int64) {
	if err := binary.Write(r.writer, byteOrder, v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readUint64() uint64 {
	if r.Error() != nil {
		return 0
	}

	var v uint64

	if err := binary.Read(r.reader, byteOrder, &v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeUint64(v uint64) {
	if err := binary.Write(r.writer, byteOrder, v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readFloat32() float32 {
	if r.Error() != nil {
		return 0
	}

	var v float32

	if err := binary.Read(r.reader, byteOrder, &v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeFloat32(v float32) {
	if err := binary.Write(r.writer, byteOrder, v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readFloat64() float64 {
	if r.Error() != nil {
		return 0
	}

	var v float64

	if err := binary.Read(r.reader, byteOrder, &v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeFloat64(v float64) {
	if err := binary.Write(r.writer, byteOrder, v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readBytes() []byte {
	length := r.readInt32()

	if length <= 0 || r.Error() != nil {
		return nil
	}

	if length > maxRemoteBytesLength {
		r.setError(ErrRemoteDataExceedsLimit)
		return nil
	}

	v := make([]byte, length)

	if _, err := io.ReadFull(r.reader, v); err != nil {
		r.setError(err)
	}

	return v
}

func (r *Remote) writeBytes(v []byte) {
	length := len(v)

	r.writeInt32(must.IntToInt32(length))

	if length <= 0 || r.Error() != nil {
		return
	}

	if _, err := r.writer.Write(v); err != nil {
		r.setError(err)
	}
}

func (r *Remote) readString() string {
	return string(r.readBytes())
}

func (r *Remote) writeString(v string) {
	r.writeBytes([]byte(v))
}

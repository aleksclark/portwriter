package portwriter

import "github.com/aleksclark/sharedi2c"

type SharedI2CWriter struct {
	Port byte
	SubPort byte
	writer *sharedi2c.SharedWriter
}

func NewSharedI2CWriter(port, subPort byte) *SharedI2CWriter {
	w := new(SharedI2CWriter)
	w.Port = port
	w.SubPort = subPort
	w.writer = sharedi2c.NewSharedWriter(port)
	return w
}

func (w *SharedI2CWriter) WriteByte(data byte) {
	msg := sharedi2c.I2CMsg{Addr: w.SubPort, Value: data}
	w.writer.SendMsg(msg)
}

package writerhook

import (
	"io"

	"github.com/sirupsen/logrus"
)

type WriterHook struct {
	writer    io.Writer
	levels    []logrus.Level
	formatter logrus.Formatter
}

// NewWriterHook creates a new writer hook logger
func NewWriterHook(writer io.Writer, level logrus.Level, formatter logrus.Formatter) (logrus.Hook, error) {
	hook := WriterHook{
		writer:    writer,
		levels:    logrus.AllLevels[:level+1],
		formatter: formatter,
	}
	return &hook, nil
}

// Levels determines what this hook will support
func (hook *WriterHook) Levels() []logrus.Level {
	return hook.levels
}

// Fire sends the entry
func (hook *WriterHook) Fire(entry *logrus.Entry) (err error) {
	b, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}
	hook.writer.Write(b)
	return nil
}

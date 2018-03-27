# WriterHook

This is a simple hook for logrus to write log files to any Writer

```
import (
    "github.com/snowzach/writerhook"
    "github.com/sirupsen/logrus"
)

writerHook, err := writerhook.NewWriterHook(os.Stderr, logrus.LevelDebug, &logrus.TextFormatter{ForceColors: true})
```

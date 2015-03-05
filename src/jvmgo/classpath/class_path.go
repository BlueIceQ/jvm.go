package classpath

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

var classNotFoundErr = errors.New("class not found!")

type ClassPath struct {
	compoundEntry CompoundClassPathEntry
}

func ParseClassPath(cpOption string) *ClassPath {
	if cpOption == "" {
		return &ClassPath{}
	}

	cpOptionSplitted := strings.Split(cpOption, pathListSeparator)
	cpEntries := make([]ClassPathEntry, len(cpOptionSplitted))

	for i, p := range cpOptionSplitted {
		absPath, err := filepath.Abs(p)
		if err == nil {
			cpEntries[i] = parseClassPathEntry(absPath)
		} else {
			// todo
		}
	}

	compoundEntry := CompoundClassPathEntry{cpEntries}
	return &ClassPath{compoundEntry}
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClassData(className string) (ClassPathEntry, []byte, error) {
	className = className + ".class"
	return self.compoundEntry.readClassData(className)
}

package ut

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/quexer/utee"
)

type FileUt struct {
}

func (p *FileUt) ReadAsset(filepath string) string {
	b, _ := ioutil.ReadAll(Assets.Files[filepath])
	return string(b)
}

func (p *FileUt) Exist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func (p *FileUt) Remove(filename string) error {
	return os.Remove(filename)
}

func (p *FileUt) MkDir(dir string) {
	if !p.Exist(dir) {
		err := os.MkdirAll(dir, 0777)
		utee.Chk(err)
	}
}

func (p *FileUt) MakeFile(dir, fileName string) (*os.File, error) {
	var f *os.File
	var err error
	if !p.Exist(dir) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			log.Println("loc 01")
			return nil, err
		}
	}
	fileFullName := fmt.Sprint(dir, fileName)
	if p.Exist(fileFullName) { // 如果文件存在
		f, err = os.OpenFile(fileFullName, os.O_RDWR, 0666) // 打开文件
		log.Println("@fileFullPath:", fileFullName, " is exist")
	} else {
		f, err = os.Create(fileFullName) // 创建文件
		log.Println("loc 02 @path:", fileFullName)
	}
	return f, err
}

func (p *FileUt) ReadLinesChannel(filePath string) (<-chan string, error) {
	c := make(chan string)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	go func() {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()
	return c, nil
}

func (p *FileUt) ReadLinesSlice(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (p *FileUt) NewFilePath(s string) FilePath {
	return FilePath(s)
}

func (p *FileUt) WriteLinesSlice(lines []string, path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		file, err = p.NewFilePath(path).Open()
		if err != nil {
			return err
		}
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func (p *FileUt) ReadAll(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer fi.Close()
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks), nil
}

func (p *FileUt) FlushWrite(path, content string) int {
	_, err := os.OpenFile(path, os.O_TRUNC, 0666)
	if err != nil {
		_, err = os.Create(path)
		utee.Chk(err)
	}
	err = ioutil.WriteFile(path, []byte(content), 0666)
	utee.Chk(err)
	return len([]byte(content))
}

func (p *FileUt) FlushWriteBytes(path string, b []byte) int {
	_, err := os.OpenFile(path, os.O_TRUNC, 0666)
	if err != nil {
		_, err = os.Create(path)
		utee.Chk(err)
	}
	err = ioutil.WriteFile(path, b, 0666)
	utee.Chk(err)
	return len(b)
}

type FilePath string

func (p FilePath) FullPath() (string, error) {
	f, err := p.Open()
	if err != nil {
		return "", err
	}
	return filepath.Abs(f.Name())
}

func (p FilePath) UnixPath() string {
	return strings.Replace(p.String(), "\\", "/", -1)
}

func (p FilePath) WindowsPath() string {
	return strings.Replace(p.String(), "/", "\\", -1)
}

func (p FilePath) FileName() string {
	toks := strings.Split(p.UnixPath(), "/")
	return toks[len(toks)-1]
}

func (p FilePath) FolderName() string {
	return strings.Replace(p.String(), p.FileName(), "", -1)
}

func (p FilePath) Exist() bool {
	var exist = true
	if _, err := os.Stat(p.String()); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func (p FilePath) String() string {
	return string(p)
}

func (p FilePath) Writer() (*bufio.Writer, error) {
	if f, err := p.Open(); err != nil {
		return nil, err
	} else {
		return bufio.NewWriter(f), nil
	}
}

func (p FilePath) Open() (*os.File, error) {
	var file *os.File
	var err error
	if p.Exist() {
		file, err = os.OpenFile(p.String(), os.O_RDWR, 0666)
	} else {
		if err := os.MkdirAll(p.FolderName(), 0777); err != nil {
			return nil, err
		}
		file, err = os.Create(p.String()) // 创建文件
	}
	return file, err
}

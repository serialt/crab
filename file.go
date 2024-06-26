package crab

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取项目路径
func GetRootPath() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		print(err.Error())
	}

	RootPath := strings.Replace(dir, "\\", "/", -1)
	return RootPath
}

// IsExist checks if a file or directory exists.
// Play: https://go.dev/play/p/nKKXt8ZQbmh
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// CreateFile create a file in path.
// Play: https://go.dev/play/p/lDt8PEsTNKI
func CreateFile(path string) bool {
	file, err := os.Create(path)
	if err != nil {
		return false
	}

	defer file.Close()
	return true
}

// CreateDir create directory in absolute path. param `absPath` like /a/, /a/b/.
// Play: https://go.dev/play/p/qUuCe1OGQnM
func CreateDir(absPath string) error {
	// return os.MkdirAll(path.Dir(absPath), os.ModePerm)
	return os.MkdirAll(absPath, os.ModePerm)
}

// MkDir 创建文件夹,支持x/a/a  多层级
func MkDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			//文件夹不存在，创建
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

// IsDir checks if the path is directory or not.
// Play: https://go.dev/play/p/WkVwEKqtOWk
func IsDir(path string) bool {
	if len(path) == 0 {
		return false
	}
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}

// IsFile reports whether the named file or directory exists. 是否是文件
func IsFile(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return !fi.IsDir()
	}
	return false
}

// IsAbsPath is abs path. 是否是绝对路径
func IsAbsPath(aPath string) bool {
	return path.IsAbs(aPath)
}

// Dir get dir path, without last name. 获取路径的目录
func PathDir(fpath string) string {
	return filepath.Dir(fpath)
}

// Name get file/dir name 获取路径的文件名
func Name(fpath string) string {
	// return path.Base(fpath)
	return filepath.Base(fpath)
}

// RemoveFile remove the path file.
// Play: https://go.dev/play/p/P2y0XW8a1SH
func RemoveFile(path string) error {
	return os.Remove(path)
}

// CopyFile copy src file to dest file.
// Play: https://go.dev/play/p/Jg9AMJMLrJi
func CopyFile(srcFilePath string, dstFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	distFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer distFile.Close()

	var tmp = make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(tmp)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		_, err = distFile.Write(tmp[:n])
		if err != nil {
			return err
		}
	}
}

// ClearFile write empty string to path file.
// Play: https://go.dev/play/p/NRZ0ZT-G94H
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("")
	return err
}

// ReadFileToString return string of file content.
// Play: https://go.dev/play/p/cmfwp_5SQTp
func ReadFileToString(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadFileByLine read file line by line.
// Play: https://go.dev/play/p/svJP_7ZrBrD
func ReadFileByLine(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := string(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		result = append(result, l)
	}

	return result, nil
}

// FileReadPointLine 从文件中读取指定行并返回字符串数组
func FileReadPointLine(filePath string, line int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	finReader := bufio.NewReader(file)
	lineCount := 1
	for {
		inputString, err := finReader.ReadString('\n')
		//fmt.Println(inputString)
		if err == io.EOF {
			if lineCount == line {
				return inputString, nil
			}
			return "", errors.New("index out of line count")
		}
		if lineCount == line {
			return inputString, nil
		}
		lineCount++
	}
}

// ListFileNames return all file names in the path.
// Play: https://go.dev/play/p/Tjd7Y07rejl
func ListFileNames(path string) ([]string, error) {
	if !IsExist(path) {
		return []string{}, nil
	}

	fs, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	result := []string{}
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			result = append(result, fs[i].Name())
		}
	}

	return result, nil
}

// IsZipFile checks if file is zip or not.
// Play: https://go.dev/play/p/9M0g2j_uF_e
func IsZipFile(filepath string) bool {
	f, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// Zip create zip file, fpath could be a single file or a directory.
// Play: https://go.dev/play/p/j-3sWBp8ik_P
func Zip(path string, destPath string) error {
	if IsDir(path) {
		return zipFolder(path, destPath)
	}

	return zipFile(path, destPath)
}

func zipFile(filePath string, destPath string) error {
	zipFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	return addFileToArchive1(filePath, archive)
}

func zipFolder(folderPath string, destPath string) error {
	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)

	err = addFileToArchive2(w, folderPath, "")
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func addFileToArchive1(fpath string, archive *zip.Writer) error {
	err := filepath.Walk(fpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(fpath)+"/")

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			if _, err := io.Copy(writer, file); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func addFileToArchive2(w *zip.Writer, basePath, baseInZip string) error {
	files, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(basePath, "/") {
		basePath = basePath + "/"
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := os.ReadFile(basePath + file.Name())
			if err != nil {
				return err
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			addFileToArchive2(w, newBase, baseInZip+file.Name()+"/")
		}
	}

	return nil
}

// UnZip unzip the file and save it to destPath.
// Play: https://go.dev/play/p/g0w34kS7B8m
func UnZip(zipFile string, destPath string) error {

	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		//issue#62: fix ZipSlip bug
		path, err := safeFilepathJoin(destPath, f.Name)
		if err != nil {
			return err
		}

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func safeFilepathJoin(path1, path2 string) (string, error) {
	relPath, err := filepath.Rel(".", path2)
	if err != nil || strings.HasPrefix(relPath, "..") {
		return "", fmt.Errorf("(zipslip) filepath is unsafe %q: %v", path2, err)
	}
	if path1 == "" {
		path1 = "."
	}
	return filepath.Join(path1, filepath.Join("/", relPath)), nil
}

// IsLink checks if a file is symbol link or not.
// Play: https://go.dev/play/p/TL-b-Kzvf44
func IsLink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

// FileMode return file's mode and permission.
// Play: https://go.dev/play/p/2l2hI42fA3p
func FileMode(path string) (fs.FileMode, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

// Suffix get filename ext. alias of path.Ext() 获取文件的后缀, main.go 获取的后缀是.go
func FileExt(fpath string) string {
	return filepath.Ext(fpath)
}

// Suffix get filename ext. alias of path.Ext() 获取文件的后缀, main.go 获取的后缀是.go
func Suffix(fpath string) string {
	return filepath.Ext(fpath)
}

// Prefix 获取文件名前缀, /tmp/main.go 获取的文件前缀是main
func Prefix(fpath string) string {
	return strings.TrimSuffix(filepath.Base(fpath), filepath.Ext(fpath))
}

// MiMeType return file mime type
// param `file` should be string(file path) or *os.File.
// Play: https://go.dev/play/p/bd5sevSUZNu
func MiMeType(file any) string {
	var mediatype string

	readBuffer := func(f *os.File) ([]byte, error) {
		buffer := make([]byte, 512)
		_, err := f.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}

	if filePath, ok := file.(string); ok {
		f, err := os.Open(filePath)
		if err != nil {
			return mediatype
		}
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}

	if f, ok := file.(*os.File); ok {
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}
	return mediatype
}

// CurrentPath return current absolute path.
// Play: todo
func CurrentPath() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		absPath = path.Dir(filename)
	}

	return absPath
}

// FileParentPath 文件父路径
func FileParentPath(filePath string) string {
	return filePath[0:strings.LastIndex(filePath, "/")]
}

// WriteFile 写文件
func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

// WriteStringToFile write string to file
func WriteStringToFile(filename, data string, mode os.FileMode) (err error) {
	return os.WriteFile(filename, []byte(data), mode)
}

// WriteJsonToFile write json to file
func WriteJsonToFile(filename string, data any, mode os.FileMode) (err error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return
	}
	os.WriteFile(filename, byteData, mode)
	return
}

// FilePathExists 判断路径是否存在
func FilePathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileAppend 追加内容到文件中
func FileAppend(filePath string, data []byte, force bool) (int, error) {
	var (
		file *os.File
		n    int
		err  error
	)
	exist := FilePathExists(filePath)
	if exist {
		if force {
			// 创建文件，如果文件已存在，会将文件清空
			if file, err = os.Create(filePath); err != nil {
				return 0, err
			}
		} else {
			if file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644); nil != err {
				return 0, err
			}
		}
	} else {
		parentPath := FileParentPath(filePath)
		if err = os.MkdirAll(parentPath, os.ModePerm); nil != err {
			return 0, err
		}
		if file, err = os.Create(filePath); err != nil {
			return 0, err
		}
	}
	defer func() {
		_ = file.Close()
	}()
	// 将数据写入文件中
	//file.WriteString(string(data)) //写入字符串
	if n, err = file.Write(data); nil != err { // 写入byte的slice数据
		return 0, err
	}
	return n, nil
}

// RecreateDir recreate dir
func RecreateDir(dir string) error {
	mode, err := FileMode(dir)
	if err != nil {
		return err
	}
	_ = os.RemoveAll(dir)
	return os.MkdirAll(dir, mode)
}

// GetFilepaths get all filepaths in a directory tree
func GetFilepaths(dir string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	return paths, err
}

// FileLoopDirs 遍历目录下的所有子目录，即返回pathname下面的所有目录，目录为绝对路径
func FileLoopDirs(pathname string) ([]string, error) {
	var s []string

	rd, err := os.ReadDir(pathname)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

// FileLoopOneDirs 遍历目录下的所有子目录，即返回pathname下面的所有目录，目录为相对路径
func FileLoopOneDirs(pathname string) ([]string, error) {
	var s []string
	rd, err := os.ReadDir(pathname)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			s = append(s, fi.Name())
		}
	}
	return s, nil
}

// FileLoopFiles 遍历文件夹及子文件夹下的所有文件，即返回pathname目录下所有的文件，文件名为绝对路径
func FileLoopFiles(pathname string) ([]string, error) {
	var s []string
	rd, err := os.ReadDir(pathname)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := path.Join(pathname, fi.Name())
			sNew, err := FileLoopFiles(fullDir)
			if err != nil {
				return s, err
			}
			s = append(s, sNew...)
		} else {
			fullName := filepath.Join(pathname, fi.Name())
			s = append(s, fullName)
		}
	}
	return s, nil
}

// FileLoopFileNames 遍历文件夹及子文件夹下的所有文件名，即返回pathname目录下所有的文件，文件名为相对路径
func FileLoopFileNames(pathname string) ([]string, error) {
	var s []string
	rd, err := os.ReadDir(pathname)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := path.Join(pathname, fi.Name())
			sNew, err := FileLoopFileNames(fullDir)
			if err != nil {
				return s, err
			}
			s = append(s, sNew...)
		} else {
			s = append(s, fi.Name())
		}
	}
	return s, nil
}

// FileMove 移动文件
func FileMove(src string, dst string) (err error) {
	if dst == "" {
		return nil
	}
	src, err = filepath.Abs(src)
	if err != nil {
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}
	var revoke = false
	dir := filepath.Dir(dst)
Redirect:
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		if !revoke {
			revoke = true
			goto Redirect
		}
	}
	return os.Rename(src, dst)
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// ReadCsvFile read file content into slice.
// Play: https://go.dev/play/p/OExTkhGEd3_u
func ReadCsvFile(filepath string, delimiter ...rune) ([][]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	if len(delimiter) > 0 {
		reader.Comma = delimiter[0]
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// WriteCsvFile write content to target csv file.
// append: append to existing csv file
// delimiter: specifies csv delimiter
// Play: https://go.dev/play/p/dAXm58Q5U1o
func WriteCsvFile(filepath string, records [][]string, append bool, delimiter ...rune) error {
	flag := os.O_RDWR | os.O_CREATE

	if append {
		flag = flag | os.O_APPEND
	}

	f, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	// 设置默认分隔符为逗号，除非另外指定
	if len(delimiter) > 0 {
		writer.Comma = delimiter[0]
	} else {
		writer.Comma = ','
	}

	// 遍历所有记录并处理包含分隔符或双引号的单元格
	for i := range records {
		for j := range records[i] {
			records[i][j] = escapeCSVField(records[i][j], writer.Comma)
		}
	}

	return writer.WriteAll(records)
}

// escapeCSVField 处理单元格内容，如果包含分隔符，则用双引号包裹
func escapeCSVField(field string, delimiter rune) string {
	// 替换所有的双引号为两个双引号
	escapedField := strings.ReplaceAll(field, "\"", "\"\"")

	// 如果字段包含分隔符、双引号或换行符，用双引号包裹整个字段
	if strings.ContainsAny(escapedField, string(delimiter)+"\"\n") {
		escapedField = fmt.Sprintf("\"%s\"", escapedField)
	}

	return escapedField
}

// WriteMapsToCsv write slice of map to csv file.
// Play: todo
func WriteMapsToCsv(filepath string, records []map[string]string, append_to_existing_file bool, delimiter ...rune) error {
	var datas_to_write [][]string
	// 标题（列名）
	var headers []string
	if len(records) > 0 {
		for key := range records[0] {
			headers = append(headers, key)
		}
	}
	// 追加模式不重复写字段名
	if !append_to_existing_file {
		datas_to_write = append(datas_to_write, headers)
	}
	// 写入数据行
	for _, record := range records {
		var row []string
		for _, header := range headers {
			row = append(row, record[header])
		}
		datas_to_write = append(datas_to_write, row)
	}
	// 提取自定义分隔符
	var sep rune
	if len(delimiter) > 0 {
		sep = delimiter[0]
	} else {
		sep = ','
	}
	return WriteCsvFile(filepath, datas_to_write, append_to_existing_file, sep)
}

package crab

import (
	"os"
	"testing"

	"github.com/serialt/crab/internal"
)

func TestIsExist(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsExist")

	cases := []string{"./", "./file.go", "./a.txt"}
	expected := []bool{true, true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsExist(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestCreateFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestCreateFile")

	f := "./text.txt"
	if CreateFile(f) {
		file, err := os.Open(f)
		assert.IsNil(err)
		assert.Equal(f, file.Name())

		defer file.Close()
	} else {
		t.FailNow()
	}
	os.Remove(f)
}

func TestCreateDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestCreateDir")

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	dirPath := pwd + "/a/"
	err = CreateDir(dirPath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(true, IsExist(dirPath))
	os.Remove(dirPath)
	assert.Equal(false, IsExist(dirPath))
}

func TestIsDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsDir")

	cases := []string{"./", "./a.txt"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsDir(cases[i])
		assert.Equal(expected[i], actual)
	}
}
func TestMkDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestMkDir")
	cases := "testdata/hello"

	actual := MkDir(cases)
	if actual == nil {
		if IsDir(cases) {
			assert.Equal("true", "true")
			RemoveFile(cases)
		}
	} else {
		assert.Equal("true", "false")
	}

}
func TestIsFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsFile")

	cases := []string{"./go.mod", "./testccc"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsFile(cases[i])
		assert.Equal(expected[i], actual)
	}
}
func TestIsAbsPath(t *testing.T) {
	assert := internal.NewAssert(t, "IsAbsPath")

	cases := []string{"/tmp", "./go.mod"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsAbsPath(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestRemoveFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestRemoveFile")
	f := "./text.txt"
	if !IsExist(f) {
		CreateFile(f)
		err := RemoveFile(f)
		assert.IsNil(err)
	}
}

func TestCopyFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestCopyFile")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	destFile := "./text_copy.txt"

	err := CopyFile(srcFile, destFile)
	if err != nil {
		file, err := os.Open(destFile)
		assert.IsNil(err)
		assert.Equal(destFile, file.Name())
	}
	os.Remove(srcFile)
	os.Remove(destFile)
}

func TestReadFileToString(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadFileToString")

	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	_, err := f.WriteString("hello world")
	if err != nil {
		t.Log(err)
	}

	content, _ := ReadFileToString(path)
	assert.Equal("hello world", content)

	os.Remove(path)
}

func TestClearFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestClearFile")

	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	_, err := f.WriteString("hello world")
	if err != nil {
		t.Log(err)
	}

	err = ClearFile(path)
	assert.IsNil(err)

	content, _ := ReadFileToString(path)
	assert.Equal("", content)

	os.Remove(path)
}

func TestReadFileByLine(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadFileByLine")

	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)

	defer f.Close()

	_, err := f.WriteString("hello\nworld")
	if err != nil {
		t.Log(err)
	}

	expected := []string{"hello", "world"}
	actual, _ := ReadFileByLine(path)
	assert.Equal(expected, actual)

	os.Remove(path)
}

func TestZipAndUnZip(t *testing.T) {
	assert := internal.NewAssert(t, "TestZipAndUnZip")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	file, _ := os.OpenFile(srcFile, os.O_WRONLY|os.O_TRUNC, 0777)
	defer file.Close()

	_, err := file.WriteString("hello\nworld")
	if err != nil {
		t.Fail()
	}

	zipFile := "./text.zip"
	err = Zip(srcFile, zipFile)
	assert.IsNil(err)

	unZipPath := "./unzip"
	err = UnZip(zipFile, unZipPath)
	assert.IsNil(err)

	unZipFile := "./unzip/text.txt"
	assert.Equal(true, IsExist(unZipFile))

	os.Remove(srcFile)
	os.Remove(zipFile)
	os.RemoveAll(unZipPath)
}

func TestFileMode(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileMode")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	mode, err := FileMode(srcFile)
	assert.IsNil(err)

	t.Log(mode)

	os.Remove(srcFile)
}

func TestIsLink(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsLink")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	linkFile := "./text.link"
	if !IsExist(linkFile) {
		_ = os.Symlink(srcFile, linkFile)
	}
	assert.Equal(true, IsLink(linkFile))

	assert.Equal(false, IsLink("./file.go"))

	os.Remove(srcFile)
	os.Remove(linkFile)
}

func TestMiMeType(t *testing.T) {
	assert := internal.NewAssert(t, "TestMiMeType")

	f, _ := os.Open("./file.go")
	defer f.Close()
	assert.Equal("text/plain; charset=utf-8", MiMeType(f))
	assert.Equal("text/plain; charset=utf-8", MiMeType("./file.go"))
}

func TestListFileNames(t *testing.T) {
	assert := internal.NewAssert(t, "TestListFileNames")

	filesInPath, err := ListFileNames("./internal/")
	assert.IsNil(err)

	expected := []string{"assert.go", "assert_test.go", "error_join.go"}
	assert.Equal(expected, filesInPath)
}

func TestCurrentPath(t *testing.T) {
	absPath := CurrentPath()
	t.Log(absPath)
}

func TestWriteFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteFile")
	filename := "./testdata/writeFile.txt"
	txtData := "hello,world"
	result := WriteFile(filename, []byte(txtData))
	assert.IsNil(result)

	data, err := os.ReadFile(filename)
	assert.IsNil(err)
	assert.Equal(string(data), txtData)
	RemoveFile(filename)

}

func TestWriteStringToFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteStringToFile")
	filename := "./testdata/writeStringToFile.txt"
	txtData := "hello,world"
	CreateFile(filename)
	result := WriteStringToFile(filename, txtData, 0644)
	t.Logf("TestWriteStringToFile: %v", result)
	assert.IsNil(result)

	data, err := os.ReadFile(filename)
	assert.IsNil(err)

	assert.Equal(string(data), txtData)
	RemoveFile(filename)

}

func TestWriteJsonToFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteJsonToFile")
	filename := "./testdata/writeJsonToFile.txt"
	txtData := `{"hello": "world"}`
	CreateFile(filename)
	result := WriteStringToFile(filename, txtData, 0644)
	t.Logf("TestWriteStringToFile: %v", result)
	assert.IsNil(result)

	data, err := os.ReadFile(filename)
	assert.IsNil(err)

	assert.Equal(string(data), txtData)
	RemoveFile(filename)

}

func TestReadCsvFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadCsvFile")

	content, err := ReadCsvFile("./testdata/demo.csv")

	assert.IsNil(err)

	assert.Equal(3, len(content))
	assert.Equal(3, len(content[0]))
	assert.Equal("Bob", content[0][0])
}

func TestWriteCsvFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteCsvFile")

	csvFilePath := "./testdata/test1.csv"
	content := [][]string{
		{"Lili", "22", "female"},
		{"Jim", "21", "male"},
	}

	err := WriteCsvFile(csvFilePath, content, false)
	assert.IsNil(err)

	readContent, err := ReadCsvFile(csvFilePath)

	assert.IsNil(err)

	assert.Equal(2, len(readContent))
	assert.Equal(3, len(readContent[0]))
	assert.Equal("Lili", readContent[0][0])

	// RemoveFile(csvFilePath)
}

func TestWriteMapsToCsv(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteMapsToCSV")

	csvFilePath := "./testdata/test4.csv"
	records := []map[string]string{
		{"Name": "Lili", "Age": "22", "gender": "female"},
		{"Name": "Jim", "Age": "21", "gender": "male"},
	}

	err := WriteMapsToCsv(csvFilePath, records, false, ';')

	assert.IsNil(err)

	content, err := ReadCsvFile(csvFilePath, ';')

	assert.IsNil(err)

	assert.Equal(3, len(content))
	assert.Equal(3, len(content[0]))
	// assert.Equal("Lili", content[1][0])
}

package beater

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/elastic/elastic-agent-libs/paths"
)

func getRegistry() []byte {
	fmt.Println("================================================== getRegistry")
	buf := bytes.Buffer{}
	dataPath := paths.Resolve(paths.Data, "")
	registryPath := filepath.Join(dataPath, "registry")
	f, err := os.CreateTemp("", "filebeat-registry-*.tar")
	if err != nil {
		panic(err)
	}
	f.Close()

	defer func() {
		if err := os.Remove(f.Name()); err != nil {
			panic(err)
		}
	}()

	tarFolder(registryPath, f.Name())
	// tarFolder("/home/tiago/devel/beats/x-pack/filebeat/data/registry", "/tmp/registry.tar")
	gzipFile(f.Name(), &buf)

	return buf.Bytes()
}

func gzipFile(src string, dst io.Writer) error {
	reader, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("cannot open '%s': '%w'", src, err)
	}
	defer reader.Close()

	writer := gzip.NewWriter(dst)
	defer writer.Close()
	writer.Name = filepath.Base(src)

	if _, err := io.Copy(writer, reader); err != nil {
		if err != nil {
			fmt.Errorf("cannot gzip file '%s': '%w'", src, err)
		}
	}

	return nil
}

// tarFolder creates a tar archive from the folder src and stores it at dst.
//
// dst must be the full path with extension, e.g: /tmp/foo.tar
// If src is not a folder an error is retruned
func tarFolder(src, dst string) error {
	fullPath, err := filepath.Abs(src)

	fmt.Println("============================== src:", src)
	fmt.Println("============================== fullPath:", fullPath)
	fmt.Println("============================== dst:", dst)
	if err != nil {
		fmt.Errorf("cannot get full path from '%s': '%w'", src, err)
	}
	tarFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("cannot create tar file '%s': '%w'", dst, err)
	}
	defer tarFile.Close()

	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	info, err := os.Stat(fullPath)
	if err != nil {
		return fmt.Errorf("cannot stat '%s': '%w'", fullPath, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("'%s' is not a directory", fullPath)
	}
	baseDir := filepath.Base(src)

	return filepath.Walk(fullPath, func(path string, info fs.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(info, info.Name())
		header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, src))

		if err := tarWriter.WriteHeader(header); err != nil {
			return fmt.Errorf("cannot write tar header for '%s': '%w'", path, err)
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("cannot open '%s' for reading: '%w", path, err)
		}
		defer file.Close()

		if _, err := io.Copy(tarWriter, file); err != nil {
			return fmt.Errorf("cannot read '%s': '%w'", path, err)
		}

		return nil
	})
}

// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package mage

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"crypto/sha512"
	"debug/elf"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
	"unicode"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

// Expand expands the given Go text/template string.
func Expand(in string, args ...map[string]interface{}) (string, error) {
	return expandTemplate("inline", in, FuncMap, EnvMap(args...))
}

// MustExpand expands the given Go text/template string. It panics if there is
// an error.
func MustExpand(in string, args ...map[string]interface{}) string {
	out, err := Expand(in, args...)
	if err != nil {
		panic(err)
	}
	return out
}

// ExpandFile expands the Go text/template read from src and writes the output
// to dst.
func ExpandFile(src, dst string, args ...map[string]interface{}) error {
	return expandFile(src, dst, EnvMap(args...))
}

// MustExpandFile expands the Go text/template read from src and writes the
// output to dst. It panics if there is an error.
func MustExpandFile(src, dst string, args ...map[string]interface{}) {
	if err := ExpandFile(src, dst, args...); err != nil {
		panic(err)
	}
}

func expandTemplate(name, tmpl string, funcs template.FuncMap, args ...map[string]interface{}) (string, error) {
	t := template.New(name).Option("missingkey=error")
	if len(funcs) > 0 {
		t = t.Funcs(funcs)
	}

	t, err := t.Parse(tmpl)
	if err != nil {
		if name == "inline" {
			return "", fmt.Errorf("failed to parse template '%v': %w", tmpl, err)
		}
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, joinMaps(args...)); err != nil {
		if name == "inline" {
			return "", fmt.Errorf("failed to expand template '%v': %w", tmpl, err)
		}
		return "", fmt.Errorf("failed to expand template: %w", err)
	}

	return buf.String(), nil
}

func joinMaps(args ...map[string]interface{}) map[string]interface{} {
	switch len(args) {
	case 0:
		return nil
	case 1:
		return args[0]
	}

	out := map[string]interface{}{}
	for _, m := range args {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}

func expandFile(src, dst string, args ...map[string]interface{}) error {
	tmplData, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed reading from template %v: %w", src, err)
	}

	output, err := expandTemplate(src, string(tmplData), FuncMap, args...)
	if err != nil {
		return err
	}

	dst, err = expandTemplate("inline", dst, FuncMap, args...)
	if err != nil {
		return err
	}

	if err = os.WriteFile(createDir(dst), []byte(output), 0644); err != nil {
		return fmt.Errorf("failed to write rendered template: %w", err)
	}

	return nil
}

// CWD return the current working directory.
func CWD(elem ...string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get the CWD: %w", err))
	}
	return filepath.Join(append([]string{wd}, elem...)...)
}

// EnvOr returns the value of the specified environment variable if it is
// non-empty. Otherwise it return def.
func EnvOr(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		return def
	}
	return s
}

var (
	dockerInfoValue *DockerInfo
	dockerInfoErr   error
	dockerInfoOnce  sync.Once
)

// DockerInfo contains information about the docker daemon.
type DockerInfo struct {
	OperatingSystem string   `json:"OperatingSystem"`
	Labels          []string `json:"Labels"`
	NCPU            int      `json:"NCPU"`
	MemTotal        int      `json:"MemTotal"`
}

// IsBoot2Docker returns true if the Docker OS is boot2docker.
func (info *DockerInfo) IsBoot2Docker() bool {
	return strings.Contains(strings.ToLower(info.OperatingSystem), "boot2docker")
}

// HaveDocker returns an error if docker is unavailable.
func HaveDocker() error {
	if _, err := GetDockerInfo(); err != nil {
		return fmt.Errorf("docker is not available: %w", err)
	}
	return nil
}

// GetDockerInfo returns data from the docker info command.
func GetDockerInfo() (*DockerInfo, error) {
	dockerInfoOnce.Do(func() {
		dockerInfoValue, dockerInfoErr = dockerInfo()
	})

	return dockerInfoValue, dockerInfoErr
}

func dockerInfo() (*DockerInfo, error) {
	data, err := sh.Output("docker", "info", "-f", "{{ json .}}")
	if err != nil {
		return nil, err
	}

	var info DockerInfo
	if err = json.Unmarshal([]byte(data), &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// HaveDockerCompose returns an error if docker-compose is not found on the
// PATH.
func HaveDockerCompose() error {
	_, err := exec.LookPath("docker-compose")
	if err != nil {
		return fmt.Errorf("docker-compose is not available")
	}
	return nil
}

// HaveKubectl returns an error if kind is not found on the PATH.
func HaveKubectl() error {
	_, err := exec.LookPath("kubectl")
	if err != nil {
		return fmt.Errorf("kubectl is not available")
	}
	return nil
}

// IsDarwinUniversal indicates whether ot not the darwin/universal should be
// assembled. If both platforms darwin/adm64 and darwin/arm64 are listed, then
// IsDarwinUniversal returns true.
// Note: Platforms might be edited at different moments, therefore it's necessary
// to perform this check on the fly.
func IsDarwinUniversal() bool {
	var darwinAMD64, darwinARM64 bool
	for _, p := range Platforms {
		if p.Name == "darwin/arm64" {
			darwinARM64 = true
		}
		if p.Name == "darwin/amd64" {
			darwinAMD64 = true
		}
	}

	return darwinAMD64 && darwinARM64
}

// FindReplace reads a file, performs a find/replace operation, then writes the
// output to the same file path.
func FindReplace(file string, re *regexp.Regexp, repl string) error {
	info, err := os.Stat(file)
	if err != nil {
		return err
	}

	contents, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	out := re.ReplaceAllString(string(contents), repl)
	return os.WriteFile(file, []byte(out), info.Mode().Perm())
}

// MustFindReplace invokes FindReplace and panics if an error occurs.
func MustFindReplace(file string, re *regexp.Regexp, repl string) {
	if err := FindReplace(file, re, repl); err != nil {
		panic(fmt.Errorf("failed to find and replace: %w", err))
	}
}

// DownloadFile downloads the given URL and writes the file to destinationDir.
// The path to the file is returned.
func DownloadFile(url, destinationDir string) (string, error) {
	fmt.Println("Downloading", url)

	resp, err := http.Get(url) //nolint:gosec // we trust the url
	if err != nil {
		return "", fmt.Errorf("http get failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download failed with http status: %v", resp.StatusCode)
	}

	name := filepath.Join(destinationDir, filepath.Base(url))
	f, err := os.Create(createDir(name))
	if err != nil {
		return "", fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	if _, err = io.Copy(f, resp.Body); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return name, f.Close()
}

// Extract extracts .zip, .tar.gz, or .tgz files to destinationDir.
func Extract(sourceFile, destinationDir string) error {
	ext := filepath.Ext(sourceFile)
	switch {
	case strings.HasSuffix(sourceFile, ".tar.gz"), ext == ".tgz":
		return untar(sourceFile, destinationDir)
	case ext == ".zip":
		return unzip(sourceFile, destinationDir)
	default:
		return fmt.Errorf("failed to extract %v, unhandled file extension", sourceFile)
	}
}

func unzip(sourceFile, destinationDir string) error {
	r, err := zip.OpenReader(sourceFile)
	if err != nil {
		return err
	}
	defer r.Close()

	if err = os.MkdirAll(destinationDir, 0755); err != nil {
		return err
	}

	extractAndWriteFile := func(f *zip.File) error {
		innerFile, err := f.Open()
		if err != nil {
			return err
		}
		defer innerFile.Close()

		path := filepath.Join(destinationDir, f.Name)
		if !strings.HasPrefix(path, destinationDir) {
			return fmt.Errorf("illegal file path in zip: %v", f.Name)
		}

		if f.FileInfo().IsDir() {
			return os.MkdirAll(path, f.Mode())
		}

		if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}

		out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err = io.Copy(out, innerFile); err != nil {
			return err
		}

		return out.Close()
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

// Tar compress a directory using tar + gzip algorithms but without adding
// the directory
func TarWithOptions(src string, targetFile string, trimSource bool) error {
	fmt.Printf(">> creating TAR file from directory: %s, target: %s\n", src, targetFile)

	f, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("error creating tar file: %w", err)
	}
	defer f.Close()

	// tar > gzip > file
	zr := gzip.NewWriter(f)
	tw := tar.NewWriter(zr)

	// walk through every file in the folder
	filepath.Walk(src, func(file string, fi os.FileInfo, errFn error) error {
		if errFn != nil {
			return fmt.Errorf("error traversing the file system: %w", errFn)
		}

		// if a symlink, skip file
		if fi.Mode().Type() == os.ModeSymlink {
			fmt.Printf(">> skipping symlink: %s\n", file)
			return nil
		}

		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return fmt.Errorf("error getting file info header: %w", err)
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)
		// Replace the source folder in the files to be compressed
		if trimSource {
			header.Name = strings.ReplaceAll(filepath.ToSlash(file), filepath.ToSlash(src), "")
			header.Name = strings.TrimPrefix(header.Name, "/")
			if header.Name == "" {
				fmt.Print(">> skipping root directory\n")
				return nil
			}
		}

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return fmt.Errorf("error writing header: %w", err)
		}

		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return fmt.Errorf("error opening file: %w", err)
			}
			defer data.Close()
			if _, err := io.Copy(tw, data); err != nil {
				return fmt.Errorf("error compressing file: %w", err)
			}
		}
		return nil
	})

	// produce tar
	if err := tw.Close(); err != nil {
		return fmt.Errorf("error closing tar file: %w", err)
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return fmt.Errorf("error closing gzip file: %w", err)
	}

	return nil
}

// Tar compress a directory using tar + gzip algorithms
func Tar(src string, targetFile string) error {
	return TarWithOptions(src, targetFile, false)
}

func untar(sourceFile, destinationDir string) error {
	file, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer file.Close()

	var fileReader io.ReadCloser = file

	if strings.HasSuffix(sourceFile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			return fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer fileReader.Close()
	}

	tarReader := tar.NewReader(fileReader)

	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading tar: %w", err)
		}

		path := filepath.Join(destinationDir, header.Name)
		if !strings.HasPrefix(path, filepath.Clean(destinationDir)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path in tar: %v", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err = os.MkdirAll(path, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg:
			writer, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			_, err = io.Copy(writer, tarReader)
			writer.Close()
			if err != nil {
				return fmt.Errorf("failed to write file contents: %w", err)
			}
		default:
			return fmt.Errorf("unsupported tar entry type: %c for file: %s", header.Typeflag, path)
		}
	}

	return nil
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || r == ',' || r == ';'
}

// RunCmds runs the given commands and stops upon the first error.
func RunCmds(cmds ...[]string) error {
	for _, cmd := range cmds {
		if err := sh.Run(cmd[0], cmd[1:]...); err != nil {
			return err
		}
	}
	return nil
}

var parallelJobsSemaphore chan struct{}

// parallelJobs returns a semaphore channel to limit the number of parallel jobs.
func parallelJobs() chan struct{} {
	if parallelJobsSemaphore == nil {
		max := numParallel()
		parallelJobsSemaphore = make(chan struct{}, max)
		fmt.Printf("Max parallel jobs: %d\n", max)
	}

	return parallelJobsSemaphore
}

// numParallel determines the maximum number of parallel jobs to run.
// It considers the MAX_PARALLEL environment variable, the number of CPUs on the host,
// and the number of CPUs reported by Docker.
func numParallel() int {
	if maxParallel, err := strconv.Atoi(os.Getenv("MAX_PARALLEL")); err == nil && maxParallel > 0 {
		return maxParallel
	}

	maxParallel := runtime.NumCPU()

	// Adjust based on Docker-reported CPUs if available
	info, err := GetDockerInfo()
	// Check that info.NCPU != 0 since docker info doesn't return with an
	// error status if communication with the daemon fails.
	if err == nil && info.NCPU != 0 {
		maxParallel = int(math.Min(float64(maxParallel), float64(info.NCPU)))
	}

	// Parallelize conservatively to avoid overloading the host.
	if maxParallel >= 2 {
		return maxParallel / 2
	}
	return maxParallel
}

// ParallelCtx runs the given functions in parallel with an upper limit set
// based on GOMAXPROCS. The provided ctx is passed to the functions (if they
// accept it as a param).
func ParallelCtx(ctx context.Context, fns ...interface{}) {
	var fnWrappers []func(context.Context) error
	for _, f := range fns {
		fnWrapper := funcTypeWrap(f)
		if fnWrapper == nil {
			panic(fmt.Sprintf("unsupported function type: %T", f))
		}
		fnWrappers = append(fnWrappers, fnWrapper)
	}

	errChan := make(chan error, len(fnWrappers))
	var wg sync.WaitGroup

	for _, fw := range fnWrappers {
		wg.Add(1)
		go func(fw func(context.Context) error) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					errChan <- fmt.Errorf("panic: %v", r)
				}
				<-parallelJobs()
			}()

			select {
			case parallelJobs() <- struct{}{}:
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			}

			waitStart := time.Now()
			fmt.Printf("Parallel job waited %v before starting.\n", time.Since(waitStart))

			if err := fw(ctx); err != nil {
				errChan <- err
			}
		}(fw)
	}

	wg.Wait()
	close(errChan)

	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		panic(errors.Join(errs...))
	}
}

// Parallel runs the given functions in parallel with an upper limit set based
// on GOMAXPROCS.
func Parallel(fns ...interface{}) {
	ParallelCtx(context.TODO(), fns...)
}

// funcTypeWrap wraps a valid FuncType to FuncContextError
func funcTypeWrap(fn interface{}) func(context.Context) error {
	switch f := fn.(type) {
	case func():
		return func(context.Context) error { f(); return nil }
	case func() error:
		return func(context.Context) error { return f() }
	case func(context.Context):
		return func(ctx context.Context) error { f(ctx); return nil }
	case func(context.Context) error:
		return f
	default:
		return nil
	}
}

// FindFiles return a list of file matching the given glob patterns.
func FindFiles(globs ...string) ([]string, error) {
	var configFiles []string
	for _, glob := range globs {
		files, err := filepath.Glob(glob)
		if err != nil {
			return nil, fmt.Errorf("failed on glob %v: %w", glob, err)
		}
		configFiles = append(configFiles, files...)
	}
	return configFiles, nil
}

// FindFilesRecursive recursively traverses from the CWD and invokes the given
// match function on each regular file to determine if the given path should be
// returned as a match. It ignores files in .git directories.
func FindFilesRecursive(match func(path string, info os.FileInfo) bool) ([]string, error) {
	var matches []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Don't look for files in git directories
		if info.Mode().IsDir() && filepath.Base(path) == ".git" {
			return filepath.SkipDir
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		if match(filepath.ToSlash(path), info) {
			matches = append(matches, path)
		}
		return nil
	})
	return matches, err
}

// FileConcat concatenates files and writes the output to out.
func FileConcat(out string, perm os.FileMode, files ...string) error {
	f, err := os.OpenFile(createDir(out), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	for _, file := range files {
		in, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open input file %s: %w", file, err)
		}
		defer in.Close()

		if _, err := io.Copy(w, in); err != nil {
			return fmt.Errorf("failed to copy from %s: %w", file, err)
		}
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	return nil
}

// MustFileConcat invokes FileConcat and panics if an error occurs.
func MustFileConcat(out string, perm os.FileMode, files ...string) {
	if err := FileConcat(out, perm, files...); err != nil {
		panic(err)
	}
}

// VerifySHA256 reads a file and verifies that its SHA256 sum matches the
// specified hash.
func VerifySHA256(file string, hash string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file for sha256 verification: %w", err)
	}
	defer f.Close()

	sum := sha256.New()
	if _, err := io.Copy(sum, f); err != nil {
		return fmt.Errorf("failed reading from input file: %w", err)
	}

	computedHash := hex.EncodeToString(sum.Sum(nil))
	expectedHash := strings.TrimSpace(hash)

	if computedHash != expectedHash {
		return fmt.Errorf("SHA256 verification of %v failed. Expected=%v, computed=%v", f.Name(), expectedHash, computedHash)
	}

	fmt.Printf("SHA256 OK: %s\n", f.Name())
	return nil
}

// CreateSHA512File computes the sha512 sum of the specified file the writes
// a sidecar file containing the hash and filename.
func CreateSHA512File(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file for sha512 summing: %w", err)
	}
	defer f.Close()

	sum := sha512.New()
	if _, err := io.Copy(sum, f); err != nil {
		return fmt.Errorf("failed reading from input file: %w", err)
	}

	computedHash := hex.EncodeToString(sum.Sum(nil))
	out := fmt.Sprintf("%v  %v", computedHash, filepath.Base(file))

	return os.WriteFile(file+".sha512", []byte(out), 0644)
}

// Mage executes mage targets in the specified directory.
func Mage(dir string, targets ...string) error {
	c := exec.Command("mage", targets...)
	c.Dir = dir
	if mg.Verbose() {
		c.Env = append(os.Environ(), "MAGEFILE_VERBOSE=1")
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	log.Println("exec:", strings.Join(c.Args, " "))
	err := c.Run()
	return err
}

// IsUpToDate returns true iff dst exists and is older based on modtime than all
// of the sources.
func IsUpToDate(dst string, sources ...string) bool {
	if len(sources) == 0 {
		panic("No sources passed to IsUpToDate")
	}

	var files []string
	for _, s := range sources {
		filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
				return err
			}

			if info.Mode().IsRegular() {
				files = append(files, path)
			}

			return nil
		})
	}

	execute, err := target.Path(dst, files...)
	return err == nil && !execute
}

// OSSBeatDir returns the OSS beat directory. You can pass paths and they will
// be joined and appended to the OSS beat dir.
func OSSBeatDir(path ...string) string {
	ossDir := CWD()

	// Check if we need to correct ossDir because it's in x-pack.
	if parentDir := filepath.Base(filepath.Dir(ossDir)); parentDir == "x-pack" {
		// If the OSS version of the beat exists.
		tmp := filepath.Join(ossDir, "../..", BeatName)
		if _, err := os.Stat(tmp); !os.IsNotExist(err) {
			ossDir = tmp
		}
	}

	return filepath.Join(append([]string{ossDir}, path...)...)
}

// XPackBeatDir returns the X-Pack beat directory. You can pass paths and they
// will be joined and appended to the X-Pack beat dir.
func XPackBeatDir(path ...string) string {
	// Check if we have an X-Pack only beats
	cur := CWD()

	if parentDir := filepath.Base(filepath.Dir(cur)); parentDir == "x-pack" {
		tmp := filepath.Join(filepath.Dir(cur), BeatName)
		return filepath.Join(append([]string{tmp}, path...)...)
	}

	return OSSBeatDir(append([]string{XPackDir, BeatName}, path...)...)
}

// LibbeatDir returns the libbeat directory. You can pass paths and
// they will be joined and appended to the libbeat dir.
func LibbeatDir(path ...string) string {
	esBeatsDir, err := ElasticBeatsDir()
	if err != nil {
		panic(fmt.Errorf("failed determine libbeat dir location: %w", err))
	}

	return filepath.Join(append([]string{esBeatsDir, "libbeat"}, path...)...)
}

// createDir creates the parent directory for the given file.
// Deprecated: Use CreateDir.
func createDir(file string) string {
	return CreateDir(file)
}

// CreateDir creates the parent directory for the given file.
func CreateDir(file string) string {
	// Create the output directory.
	if dir := filepath.Dir(file); dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(fmt.Errorf("failed to create parent dir for %v: %w", file, err))
		}
	}
	return file
}

// binaryExtension returns the appropriate file extension based on GOOS.
func binaryExtension(goos string) string {
	if goos == "windows" {
		return ".exe"
	}
	return ""
}

var parseVersionRegex = regexp.MustCompile(`(?m)^[^\d]*(?P<major>\d+)\.(?P<minor>\d+)(?:\.(?P<patch>\d+).*)?$`)

// ParseVersion extracts the major, minor, and optional patch number from a
// version string.
func ParseVersion(version string) (major, minor, patch int, err error) {
	names := parseVersionRegex.SubexpNames()
	matches := parseVersionRegex.FindStringSubmatch(version)
	if len(matches) == 0 {
		err = fmt.Errorf("failed to parse version '%v'", version)
		return
	}

	data := map[string]string{}
	for i, match := range matches {
		data[names[i]] = match
	}
	major, _ = strconv.Atoi(data["major"])
	minor, _ = strconv.Atoi(data["minor"])
	patch, _ = strconv.Atoi(data["patch"])
	return
}

// ListMatchingEnvVars returns all of the environment variables names that begin
// with prefix.
func ListMatchingEnvVars(prefixes ...string) []string {
	var vars []string
	for _, v := range os.Environ() {
		for _, prefix := range prefixes {
			if strings.HasPrefix(v, prefix) {
				eqIdx := strings.Index(v, "=")
				if eqIdx != -1 {
					vars = append(vars, v[:eqIdx])
				}
				break
			}
		}
	}
	return vars
}

// IntegrationTestEnvVars returns the names of environment variables needed to configure
// connections to integration test environments.
func IntegrationTestEnvVars() []string {
	// Environment variables that can be configured with paths to files
	// with authentication information.
	vars := []string{
		"AWS_SHARED_CREDENTIAL_FILE",
		"AZURE_AUTH_LOCATION",
		"GOOGLE_APPLICATION_CREDENTIALS",
	}
	// Environment variables with authentication information.
	prefixes := []string{
		"AWS_",
		"AZURE_",
		"GCP_",

		// Accepted by terraform, but not by many clients, including Beats
		"GOOGLE_",
		"GCLOUD_",
	}
	for _, prefix := range prefixes {
		vars = append(vars, ListMatchingEnvVars(prefix)...)
	}
	return vars
}

// ReadGLIBCRequirement returns the required glibc version for a dynamically
// linked ELF binary. The target machine must have a version equal to or
// greater than (newer) the returned value.
func ReadGLIBCRequirement(elfFile string) (*SemanticVersion, error) {
	e, err := elf.Open(elfFile)
	if err != nil {
		return nil, err
	}

	symbols, err := e.DynamicSymbols()
	if err != nil {
		return nil, err
	}

	versionSet := map[SemanticVersion]struct{}{}
	for _, sym := range symbols {
		if strings.HasPrefix(sym.Version, "GLIBC_") {
			semver, err := NewSemanticVersion(strings.TrimPrefix(sym.Version, "GLIBC_"))
			if err != nil {
				continue
			}

			versionSet[*semver] = struct{}{}
		}
	}

	if len(versionSet) == 0 {
		return nil, errors.New("no GLIBC symbols found in binary (is this a static binary?)")
	}

	var versions []SemanticVersion
	for ver := range versionSet {
		versions = append(versions, ver)
	}

	sort.Slice(versions, func(i, j int) bool {
		a := versions[i]
		b := versions[j]
		return a.LessThan(&b)
	})

	max := versions[len(versions)-1]
	return &max, nil
}

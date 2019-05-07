// Code generated by vfsgen; DO NOT EDIT.

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 5, 5, 20, 58, 9, 867694403, time.UTC),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2019, 5, 7, 3, 17, 11, 846928687, time.UTC),
			uncompressedSize: 4271,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x58\x5d\x8f\xd3\x38\x17\xbe\x9f\x5f\x71\xc8\x55\x2a\xaa\xa4\x2f\xbc\x2b\xa1\xb6\x41\x5a\x58\xb4\xb0\x62\x61\x45\x47\xb0\xa3\x51\x2f\x3c\xc9\x69\x62\x48\xed\xc8\x76\xdb\x19\xa1\xf9\xef\x2b\xc7\x6e\x62\x3b\x09\x94\xab\x36\x3e\x9f\x7e\xce\x73\xec\x93\xac\x2b\xb5\xaf\x5f\x5e\x01\xac\x2b\x24\x85\xfe\x03\xb0\x96\xb9\xa0\x8d\x02\xf5\xd0\x60\x16\x29\xbc\x57\xe9\x57\x72\x24\x66\x35\x32\x3a\x00\x47\x22\x40\xf2\xfc\x1b\xaa\xd5\x95\x5d\xda\x1d\x58\xae\x28\x67\x20\xb0\xa4\x52\xa1\x78\x73\x44\xa6\xde\x12\x56\xd4\x28\x62\xc6\x0b\x9c\xb7\x4e\xe7\x40\x8b\x19\x7c\xb7\x56\x00\x5a\x92\x90\xa2\x68\xd5\xdf\x6b\x4b\x86\x22\x36\x9a\x67\x9f\x31\x1e\x95\x6b\x03\x50\xa3\x02\xd4\x16\x7f\x10\x45\x20\x83\x28\x5a\x39\x52\x79\xa2\x2a\xaf\xa0\xf5\xe2\xdb\x01\xe4\x44\x22\x44\x79\x45\x58\x89\xd1\xd2\x13\x81\xe7\xf2\x3b\x7c\x26\xf5\x01\x97\x80\x47\x95\x28\x22\x4a\x54\xc9\x51\xaf\xc0\xe3\x2a\x30\xbb\x13\x48\xbe\xad\xc6\xc2\xd4\x34\xff\x16\x44\x31\x92\x3d\x3f\x48\x2c\xf8\x89\x4d\x4b\x0f\xcd\xb4\x0c\x99\x42\x31\x2d\xae\x91\x1c\xc3\xdd\x39\xe2\x3d\x1f\x48\x83\xbd\x07\x32\x80\x57\x07\xa5\x38\x93\x06\x8d\x3b\xf3\x30\x1f\x68\xbd\xae\x29\x32\xf5\xaf\xd1\xca\xcd\xc3\x94\xd6\x8d\xab\x75\x33\xd4\xfa\xb8\xdb\x49\x3c\xfb\xe2\xe6\x61\x4a\xeb\xc6\xd5\xba\x09\x94\x7e\x5e\xae\xc7\xab\x80\x5a\x05\xdf\xb7\x7c\x1c\x40\x71\xfd\xd0\xe0\xd2\xf0\xd8\x5b\x7f\x53\xe3\x1e\x99\x7a\x57\x2c\x81\x16\xbe\x48\x43\xba\xec\xd1\xbd\x9a\x48\x2c\xe7\x4c\xf2\x1a\x93\x9a\x97\xf1\x39\xfc\xcc\x23\x75\xdb\x6f\x89\x44\x56\xc4\x7f\x6d\x3e\x7e\x48\xa4\x12\x94\x95\x74\xf7\xd0\xeb\x3b\x06\x8f\xdd\xff\xc7\x41\x8b\xe6\x02\x89\x42\x9b\xf3\x27\xcc\x0f\x42\xd2\x23\xc6\xf2\x58\x7e\x20\x7b\x94\x0d\xc9\x71\x0e\x68\xc4\x6e\xff\x68\x68\x74\xbf\x42\x06\xec\x50\xd7\x7d\x30\xba\x03\xcf\xda\x6f\x3a\x6b\x52\xf0\xfc\xa0\x3d\x26\x5e\xf8\x0f\x9b\xd8\x83\x2b\xaa\x94\x6a\x96\x69\x7a\x3a\x9d\x92\xd3\xf3\x84\x8b\x32\x7d\xb6\x58\x2c\x52\x79\x2c\x23\x1f\x58\x9b\x60\xa2\xa3\x3a\x02\x17\x03\xc0\x5a\xe2\xe5\xb9\xc4\xae\x4b\xd7\x4f\xcf\x8f\x1d\x17\x10\xeb\xe3\x8f\x91\x3d\x02\x65\x5d\x16\xbf\x2b\x25\xe4\x70\xdf\x89\x44\xa5\x45\xf4\xee\xa0\x30\xd6\x46\x73\xdf\xe4\x56\xaf\x6d\x7f\x12\x8c\x42\x06\x8b\x15\x50\x58\x77\xc6\xaf\x2b\x5a\x17\x02\x59\x52\x23\x2b\x55\xb5\x02\xfa\xf4\xa9\x1f\x5e\x17\x25\xd4\xbe\xa5\xdb\x44\x73\x18\xb2\x0c\x9e\x85\x47\x63\x9b\x2f\x65\x0c\xc5\x35\xde\x6b\xee\x8f\x59\x9b\xa4\x23\xad\x11\x6d\xbd\x26\x1a\x82\x0d\x90\x6b\x4b\xc8\x7e\x89\x71\x6e\xb8\xd9\x6a\x98\x21\x69\x1a\x64\x45\xab\x14\xb7\xfe\x67\x7e\x2f\xff\x22\x90\xee\x3d\x25\xa7\xd1\x1c\xbd\xd6\x06\xc9\x8d\x33\xd4\x0b\x71\x2e\xc1\xb8\xaa\x85\x97\x16\xd1\x76\x82\xd3\xfd\xae\x04\xaa\x83\x60\x6d\xdc\xe9\x5e\x2f\x51\x59\xd8\x5f\x3d\xfc\x43\x54\x15\x37\x44\x55\x61\x53\xdb\xe8\x5d\x5f\x77\xc2\x34\xd5\x0e\x40\x70\xde\x2b\xd1\x1d\x50\x05\x78\x4f\xa5\x92\x73\xc0\xfb\xbc\x3e\x14\x94\x95\x40\xd8\x03\x50\xf6\x15\x73\x85\x05\xd8\x31\xc2\xda\xc8\x1f\x56\xa2\x6b\xc6\x3b\x5e\x3c\x24\xf9\x45\xc4\x7e\x12\x8f\x5b\xdd\xd2\x2d\x50\x26\x15\x61\x39\xf2\x1d\xbc\xbd\xfe\xfb\xfd\xa6\x4d\xc5\x82\x30\x0b\x49\xdf\x6f\x7d\xd2\xa1\xcf\xc1\x91\xdb\xe3\x42\xc6\x69\xe4\xa7\x77\xd5\x27\x72\xe6\x42\x97\x82\x36\xbc\xa5\xdb\xad\xcb\x82\x80\x04\xd6\x66\x9a\x07\xa4\x69\x6a\x4d\x80\xbc\xd2\x43\x96\x66\x81\xf9\xe7\x26\xa1\xef\x20\x05\xef\x98\x44\xa1\xf3\x58\xb8\x14\xfd\xd4\x4e\x0d\x90\xc1\xff\xfc\xd5\xa6\x26\xb9\x5e\x7e\xe6\x2e\x6b\x16\x6f\x50\xfb\x78\x1e\x2e\x77\x7e\xfe\xef\x4a\xf4\x71\x62\x0c\x7e\x5b\xfd\x14\x44\x93\x78\xd2\x6e\x66\x1a\x4f\xcd\xeb\x56\x19\xb2\xd0\xa8\xad\xe9\x40\x55\x8c\xd0\x60\x15\x68\x99\x39\x10\xb2\xd1\xb6\xca\x2b\xed\xbe\xf2\x8e\x23\xcd\x55\x6b\xf4\xc4\x74\x57\x48\xc0\x2e\xb0\x9d\x31\xcd\xf3\x1b\xbf\x9e\x7e\xcd\xfb\x09\xd7\x44\xbd\x9e\x9a\x73\x6d\x79\xc2\x51\xcf\x6f\xf9\x89\xb3\x79\x30\x6a\xf5\x20\x6e\x3e\xff\xd9\x1f\xdb\xe3\x7a\x89\xf5\x17\x48\x67\x1e\xee\x3e\x40\xe1\x0e\x7a\x70\x12\x61\xf6\x61\xce\x7c\x9b\xf9\xdc\x02\x36\x0b\xc7\xbb\xd1\x8b\xc8\x71\xe6\xde\x1f\xe7\x21\x67\xe0\xe3\xc2\x09\xdf\x12\x3d\x04\xd8\x96\xd2\xbb\xfb\x0d\x2c\xfa\xb9\x1d\x2f\xe6\xe0\x2c\xb4\x6f\x1a\x83\x24\x7e\x14\xd4\xb4\xd1\x44\x5c\xd1\x0a\x27\x43\x5f\x1c\xc8\x76\xe5\x44\x14\x77\x52\x08\xf7\x72\xc1\xc8\x1d\xfe\xeb\x8e\xac\x93\x84\x0c\x28\xa3\xea\xcb\x26\x9e\x0d\xdf\x2f\xcf\x12\xa7\xc0\x66\x34\xd6\xb7\x17\x9e\xe0\x0b\xde\x6d\xda\xe7\x38\x3a\xc9\x65\x9a\xd6\x3c\x27\x75\xc5\xa5\x5a\xbe\x58\xbc\x58\xa4\xe6\x7d\x23\x72\x30\xb0\x83\x35\x67\xbc\x41\x06\x59\xff\xd6\xe9\x13\xd2\x1d\xd0\x23\x1b\x51\x5b\xb8\xae\x1e\x47\xbc\xee\x51\x4a\x52\xa2\xeb\x38\x68\xd6\xae\xb1\x20\x83\x76\xb2\x6f\x88\x90\x18\x63\x52\x10\x45\x66\x53\x2f\x09\x51\x67\xb6\x8c\xe6\xbd\x0f\x4f\x7f\xf2\xd0\xff\x71\xca\x79\xcd\x25\xfe\x1a\x12\xad\xc9\x04\x14\xf6\x8a\x3a\x7f\x30\x70\xcb\xbe\x4e\xcd\xa0\xd0\x7e\x85\x48\xcf\x9f\x21\xd6\xfa\xe8\x7d\xb9\x4e\xdb\x9f\xab\x75\x6a\xbe\x53\xfc\x17\x00\x00\xff\xff\x41\x01\xc8\x67\xaf\x10\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/index.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}

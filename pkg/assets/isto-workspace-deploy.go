// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package assets generated by go-bindata.// sources:
// deploy/cluster_role_binding.yaml
// deploy/cluster_role.yaml
// deploy/crds/maistra.io_sessions.yaml
// deploy/operator.tpl.yaml
// deploy/operator.yaml
// deploy/role_binding.yaml
// deploy/role.yaml
// deploy/service_account.yaml
// template/strategies/_basic-remove.tpl
// template/strategies/_basic-version.tpl
// template/strategies/prepared-image.tpl
// template/strategies/prepared-image.var
// template/strategies/telepresence.tpl
// template/strategies/telepresence.var
package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _deployCluster_role_bindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x41\x6f\xdb\x30\x0c\x85\xef\xfe\x15\x44\xb0\x6b\x3c\xec\x36\xe8\x96\x05\xc1\x4e\x1b\x86\x24\xd8\x9d\x96\xe8\x98\x8b\x2d\x6a\x14\x95\x00\x2d\xfa\xdf\x8b\x28\x4e\x5b\xb4\x40\xaa\x9b\xf8\x88\xc7\xef\xbd\x23\xc7\xe0\x60\x4f\x53\x1a\xd1\xa8\xc1\xc4\x7f\x49\x33\x4b\x74\x60\xf3\xb0\x95\x44\x31\x0f\xdc\x5b\xcb\xf2\xf5\xf4\xad\x49\xa8\x38\x91\x91\x66\xd7\x00\x2c\x21\xe2\x44\x0e\x7e\xaf\x7e\x6d\x76\x7f\x56\xeb\x4d\x03\x00\x10\x28\x7b\xe5\x64\xd5\x69\xb1\x47\x3d\x90\xd5\xc5\x9c\xd0\x13\xf4\xa2\x70\x1e\xd8\x0f\xa0\x32\x12\x74\x1c\x03\xc7\x03\xe4\x41\xca\x18\xa0\x23\x08\xd4\x73\xa4\xb0\xa8\x66\x4a\xff\x0b\x2b\x05\x07\xa6\x85\xea\xe8\x84\x63\x21\x07\x9c\x8d\x65\x79\x16\x3d\x56\xdf\xa5\x24\x52\x34\xd1\x46\xba\x7f\xe4\x6d\x06\xbc\x86\x5c\x8f\x25\x1b\xe9\x56\x46\xfa\x71\xbd\x57\x9d\xde\x46\xd6\x0e\x7d\x8b\xc5\x06\x51\x7e\xc0\x0b\x7c\x7b\xfc\x9e\xe7\xd8\x97\xe5\x89\x0c\x03\x1a\xba\xfa\x83\x39\xfa\x3b\x8a\xaa\xe5\xf2\x4a\x70\x79\x37\x8a\x1d\xe9\x89\x3d\xad\xbc\x97\x12\x6d\x16\xef\x19\xdd\xd4\x3a\x71\xf0\xe5\xf1\xa5\xe8\xa7\x6b\x39\x32\xd2\x96\xfa\xdb\x9d\x0f\x59\x3f\x25\xad\x15\xfc\x54\x29\xe9\x4e\x01\xcd\x73\x00\x00\x00\xff\xff\x5a\x3b\x0c\xba\x29\x02\x00\x00")

func deployCluster_role_bindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployCluster_role_bindingYaml,
		"deploy/cluster_role_binding.yaml",
	)
}

func deployCluster_role_bindingYaml() (*asset, error) {
	bytes, err := deployCluster_role_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/cluster_role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployCluster_roleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xb1\x8e\xd4\x40\x0c\xed\xf3\x15\xa3\x6b\x4e\x42\x4a\x4e\x74\x28\x2d\x05\x1d\x05\x42\xf4\xbe\x89\xb3\x67\xed\xcc\x78\x64\x7b\x82\xb8\xaf\x47\x93\xcd\x42\x16\x92\x45\x41\xb7\xd5\x8e\x1d\xbf\xe7\x67\x3f\x43\xa6\x6f\x28\x4a\x9c\x7a\x27\xcf\xe0\x3b\x28\xf6\xc2\x42\xaf\x60\xc4\xa9\x3b\x7f\xd0\x8e\xf8\x69\x7a\xdf\x9c\x29\x0d\xbd\xfb\x18\x8a\x1a\xca\x17\x0e\xd8\x44\x34\x18\xc0\xa0\x6f\x9c\xf3\x82\x73\xc1\x57\x8a\xa8\x06\x31\xf7\x2e\x95\x10\x1a\xe7\x12\x44\xec\x1d\xa9\x11\xb7\xdf\x59\xce\x9a\xc1\x63\x23\x25\xa0\xd6\xc2\xd6\x41\xa6\x4f\xc2\x25\xcf\xcf\xfa\x6b\xdd\xc3\xc3\xfc\x57\x50\xb9\x88\xc7\x55\xa6\xa2\xcd\x08\x3a\x87\x26\x94\xe7\x55\xf6\x84\x76\x1c\x32\xf3\xa0\xbf\x1e\x8a\x32\xd1\x15\xbd\x06\x30\x0d\x99\x29\xd9\xef\x48\xae\xe3\x52\xc3\x64\x13\x87\x12\xd1\x07\xa0\xb8\x2a\x98\x70\xfd\xb5\xe7\x34\xd2\x29\x42\x5e\x73\x78\x41\xdb\x14\xf0\xf8\xee\x71\x4f\x00\xe4\x05\x62\x43\xc2\x80\x39\xf0\x8f\x78\x43\x3c\x00\x46\x4e\x8a\xab\x90\x60\x0e\xe4\xe1\x26\xa6\x06\x86\x63\x09\xfa\x7f\x1d\x75\x9c\x31\xe9\x0b\x8d\xd6\x11\xff\xbb\xbd\xcb\x34\x8e\x12\x45\x4e\x64\x2c\x94\x4e\x9d\x67\x41\xd6\xce\x73\xdc\x23\x5b\x36\xb8\xd4\xdc\xb1\xc9\xb2\x9f\x6a\x5c\xdc\x63\x9e\x6d\xbb\xd2\x78\x87\xf7\xd2\xff\x11\x59\x09\xad\x1e\x44\x95\x75\xe1\xd9\x9f\xe0\x71\xf0\x08\xa4\x26\xf0\x66\x98\x1b\x06\xfc\x5c\x6f\xf1\x9a\xfd\xf3\xc0\x37\x48\x6f\x7c\xfa\x34\x52\x82\x40\xaf\xf8\xf7\x8a\x5a\x57\xf2\x50\x97\xf2\x33\x00\x00\xff\xff\x5b\x05\xbf\xda\x9c\x04\x00\x00")

func deployCluster_roleYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployCluster_roleYaml,
		"deploy/cluster_role.yaml",
	)
}

func deployCluster_roleYaml() (*asset, error) {
	bytes, err := deployCluster_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/cluster_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployCrdsMaistraIo_sessionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x0c\xd2\x43\x2e\xb5\x9c\x45\x2e\x85\x6e\xc1\xb6\x87\xa0\x69\x10\xec\x06\xb9\x14\x3d\x8c\xc5\x91\x3c\x5d\x8a\x54\xc9\xa1\xd3\x6d\xd1\xff\x5e\x90\x94\x6c\x59\x96\xbd\x1f\x41\x80\x22\x9d\x9b\x9e\x86\x33\xe3\xf7\x66\x48\x5a\xc5\x6a\xb5\x2a\xb0\xe7\x4f\xe4\x3c\x5b\x53\x01\xf6\x4c\x7f\x0a\x99\xf8\xe4\xcb\xbb\x1f\x7c\xc9\x76\xbd\xbb\xda\x90\xe0\x55\x71\xc7\x46\x55\x70\x1d\xbc\xd8\xee\x86\xbc\x0d\xae\xa6\x1f\xa9\x61\xc3\xc2\xd6\x14\x1d\x09\x2a\x14\xac\x0a\x00\x34\xc6\x0a\x46\xd8\xc7\x47\x80\xda\x1a\x71\x56\x6b\x72\xab\x96\x4c\x79\x17\x36\xb4\x09\xac\x15\xb9\x94\x61\xcc\xbf\x7b\x55\xbe\x2e\x5f\x15\x00\xb5\xa3\xb4\xfc\x23\x77\xe4\x05\xbb\xbe\x02\x13\xb4\x2e\x00\x0c\x76\x54\x81\x27\x9f\x6b\xec\x90\xbd\x38\x2c\xd9\x16\xbe\xa7\x3a\x66\x6b\x9d\x0d\x7d\x05\x93\x37\x79\xd5\x50\x4a\xfe\x19\xb7\x39\x40\x42\x34\x7b\xf9\x79\x8a\xbe\x63\x2f\xe9\x4d\xaf\x83\x43\x7d\x48\x97\x40\xcf\xa6\x0d\x1a\xdd\x1e\x2e\x00\x7c\x6d\x7b\xaa\xe0\x7d\x4c\xd3\x63\x4d\x2a\x62\x61\xe3\x06\x9a\x86\xd4\x5e\x50\x82\xaf\xe0\xef\x7f\x0a\x80\x1d\x6a\x56\xe9\x47\xe6\x97\xb6\x27\xf3\xe6\xc3\xdb\x4f\xaf\x6f\xeb\x2d\x75\x98\x41\x00\x45\xbe\x76\xdc\x27\xbf\xb1\x3e\x60\x0f\xb2\x25\xc8\x9e\xd0\x58\x97\x1e\xc7\x2a\xe1\xcd\x87\xb7\xe5\xb0\xbc\x77\xb6\x27\x27\x3c\x96\x10\x6d\xa2\xf8\x1e\x9b\x25\x7a\x19\x2b\xc9\x3e\xa0\xa2\xc6\x94\x33\x0e\x4a\x91\x02\x9f\x73\xdb\x06\x64\xcb\x1e\x1c\xf5\x8e\x3c\x99\xac\xfa\x24\x2c\x44\x17\x34\x60\x37\xbf\x53\x2d\x25\xdc\x92\x8b\x41\xc0\x6f\x6d\xd0\x2a\x36\xc6\x8e\x9c\x80\xa3\xda\xb6\x86\xff\xda\x47\xf6\x20\x36\xa5\xd4\x28\x34\xc8\x31\x1a\x1b\x21\x67\x50\x47\x0e\x03\x7d\x0f\x68\x14\x74\x78\x0f\x8e\x62\x0e\x08\x66\x12\x2d\xb9\xf8\x12\x7e\xb1\x8e\x80\x4d\x63\x2b\xd8\x8a\xf4\xbe\x5a\xaf\x5b\x96\xb1\xc7\x6b\xdb\x75\xc1\xb0\xdc\xaf\x53\xa7\xf2\x26\x88\x75\x7e\xad\x68\x47\x7a\xed\xb9\x5d\xa1\xab\xb7\x2c\x54\x4b\x70\xb4\xc6\x9e\x57\xa9\x70\x23\xb9\x09\xd5\x77\x7b\xa5\x5f\x4e\x2a\x95\xfb\xd8\x14\x5e\x1c\x9b\x76\x0f\xa7\xfe\x3b\xcb\x7b\xec\xc3\x28\x2f\x0e\xcb\x72\xfd\x07\x7a\x23\x14\x59\xb9\xf9\xe9\xf6\x23\x8c\x49\x93\x04\xc7\x9c\x27\xb6\x0f\xcb\xfc\x81\xf8\x48\x14\x9b\x86\x5c\x16\xae\x71\xb6\x4b\x11\xc9\xa8\xde\xb2\x91\xf4\x50\x6b\x26\x73\x4c\xba\x0f\x9b\x8e\x25\x2a\xfd\x47\x20\x2f\x51\x9f\x12\xae\xd3\xa4\xc3\x86\x20\xf4\x0a\x85\x54\x09\x6f\x0d\x5c\x63\x47\xfa\x1a\x3d\x7d\x75\xda\x23\xc3\x7e\x15\x29\x7d\x98\xf8\xe9\x06\x75\xec\x98\xd9\xda\xc3\xe3\x46\xb2\xa8\xd0\x30\x82\xb7\x3d\xd5\x47\x93\xa1\xc8\xb3\x8b\xdd\x2b\x28\x14\x7b\x7e\x70\x2c\x27\x81\x96\x86\x31\x9a\xa3\xe6\x18\x00\x60\xa1\xce\xcf\xc1\x59\x29\x37\xd4\x5c\x28\x21\xee\x0a\x98\x36\x2b\x1d\xfb\xa7\x21\x47\xa6\xa6\x93\x88\x00\x9f\x59\xb6\x6c\xf2\x86\x72\x5a\xf3\xe5\xca\xb3\xa1\x6b\x17\x71\x00\x54\x2a\x9d\x0e\xa8\x3f\x5c\x8c\x00\xe7\x44\x3b\x75\x98\x89\x75\xb0\x74\x32\x5c\x58\x78\x36\x72\x3c\x26\x84\xda\xfb\x67\x2c\xbe\x50\x52\x7e\x85\xce\xe1\xfd\xb1\xd8\x36\xc8\x49\x9d\xc7\xba\x46\x8f\x23\x65\xc7\x0a\x93\xa8\x5b\xfb\x39\x81\xe2\xb0\x69\xb8\x8e\xbb\x45\x8a\xa9\x4e\xab\xcb\x3b\xe8\x0d\x35\x73\x45\x2f\xe9\x79\x8e\xc7\x87\x89\x78\xea\xa2\xb4\xb7\x3d\x71\xd5\x19\xc6\x97\x07\x39\x9f\xb7\x0f\x8d\x72\xf2\x3a\xe2\xdb\x6e\x7c\xdc\x2d\x9f\x3d\xcd\x27\xac\x3e\x72\x9c\x1f\x55\x4a\x44\xd9\x28\xde\xb1\x0a\xa8\x17\xd8\x5b\xd0\xfb\x1b\x9f\xe0\xd9\x3d\x6b\x6e\x67\xd8\xcf\x36\xd7\x60\xbc\xda\x2e\xab\xb0\xcf\x04\x5d\x88\x82\xa8\x75\xba\xaa\x2e\x0c\xdf\x68\xe8\xa1\x47\x27\xa3\x72\x8b\xe2\x64\xbb\x2c\xd1\x10\xad\x96\xd9\xbd\x6d\x6e\x0f\x8a\x00\x0b\x97\x90\x67\x05\x39\x2f\xd7\x13\x82\xc4\x1f\x7d\x29\xc8\x53\xfa\xef\x09\x69\x1f\xec\xc5\x47\xb9\x9c\xdb\xe3\xb3\x7d\xc1\xb9\x02\x20\xe8\x5a\x92\x2f\x6e\xe8\x77\xb8\x21\x4d\x6a\xda\xd7\xe9\x76\x39\x05\xe2\x05\x20\x3b\xfa\x6f\xa5\x35\x75\xfa\x35\xff\xd5\xbe\xfa\x1f\x8c\x4e\xfc\x97\x10\x2f\xa3\xe7\xf2\xad\x06\x89\xbe\xc6\xe4\x3d\xe3\x52\x96\x0e\xd7\x79\xad\x67\xd8\x58\x08\x3f\x83\x0e\xdf\x33\xae\x50\xf7\x5b\xbc\x3a\x60\x49\x80\xd5\xf0\x15\x63\xf2\x1a\x20\x1f\x30\x15\x88\x0b\x34\x7c\x2d\xb0\x0e\x5b\x1a\x90\xc3\x5d\x06\xeb\x9a\x7a\x21\xf5\x7e\xfe\x4d\xe3\xc5\x8b\xa3\x8f\x16\xe9\xb1\xb6\x26\x37\x81\xaf\xe0\xd7\xdf\x8a\x1c\x95\xd4\xa7\xb1\x9a\x08\xfe\x1b\x00\x00\xff\xff\xd6\x67\xbf\xd4\x07\x12\x00\x00")

func deployCrdsMaistraIo_sessionsYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployCrdsMaistraIo_sessionsYaml,
		"deploy/crds/maistra.io_sessions.yaml",
	)
}

func deployCrdsMaistraIo_sessionsYaml() (*asset, error) {
	bytes, err := deployCrdsMaistraIo_sessionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/crds/maistra.io_sessions.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployOperatorTplYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\x51\x6b\xe3\x48\x0c\xc7\xdf\xf3\x29\x44\xd8\xd7\xa6\xdb\xc2\xc1\x32\x6f\xa6\xc9\xe5\xc2\xee\x36\x21\x31\x7b\xdc\x53\x51\xc6\x4a\xad\xcb\x78\x66\x76\x66\xec\xe0\x5b\xf2\xdd\x0f\xbb\x76\x9b\x38\x4e\x5d\x38\xb8\x79\xb3\x24\xf4\xff\x49\x96\xb4\x67\x9d\x08\x88\x29\xb3\x0a\x03\x8d\xd0\xf2\x0f\x72\x9e\x8d\x16\x10\x1a\xe3\xc4\x58\xd2\x3e\xe5\x5d\x98\xb0\xb9\x2d\xee\x46\x16\x1d\x66\x14\xc8\x79\x31\x02\xb8\x01\x8d\x19\x09\x58\x7c\x9d\x3d\x4d\x97\x0f\x5f\x67\xeb\xa7\xf5\x6c\xbe\xd8\xc4\xeb\xbf\x46\x00\x00\x09\x79\xe9\xd8\x86\x3a\xe7\x78\x6a\xe4\x9e\x1c\x38\x7a\x66\x1f\x5c\x09\x87\x94\x1c\x41\x42\x56\x99\x92\x12\xe0\x0c\x9f\x09\xd8\x03\x16\xc8\x0a\xb7\x8a\xc6\x75\x12\x47\x3f\x73\x76\x94\x08\x08\x2e\xa7\xda\x54\xa0\xca\x49\xc0\xcf\x1c\xcb\x09\x9b\x6b\x20\xab\xe5\x66\x11\x2f\x7b\x51\xd6\x64\x8d\xe7\x60\x5c\x09\xac\xe1\x90\xb2\x4c\x21\xa4\xd4\x30\x48\xd4\xb0\x25\xd8\x99\x5c\x27\x43\x0c\x19\x56\xc5\x60\x87\x61\xf1\x3d\x9a\xcf\x9e\x1e\xa3\xef\xb3\x1e\xf1\x38\xa5\x3a\x14\xcc\xee\x44\xf4\xc0\x21\x05\xde\x13\x6c\x59\xa3\x2b\x87\x74\xd9\x07\x36\x37\x07\xe3\xf6\xde\xa2\xa4\x5e\xfd\x38\x9a\x5f\x91\x0f\xf8\x7c\xae\x1e\x4c\x55\x71\xee\x69\xb0\xe0\x6a\x2c\x7c\xe8\xe8\xfd\x98\xad\x37\x8b\xe5\xe3\x15\xb5\xe2\x65\xac\x5a\xc5\x8f\x15\x78\xa1\xf3\x67\x14\x3f\xfc\x51\xf7\x74\xb3\x8a\x1e\xde\x6b\x6c\xdd\x91\xaa\xa4\x03\x06\x99\x4e\xe0\x1b\x61\x41\x40\x99\x0d\x25\xec\x8c\x03\xa9\x72\x1f\xc8\xc1\x81\x13\x6a\x62\xba\x34\x3b\x54\xfe\x0c\x67\x3c\x1e\x99\xed\xdf\x24\x43\x33\xf8\x2f\xcb\x33\xad\x87\x37\x23\x1d\xea\xd8\xd3\x15\x42\x6b\x7d\xb5\x31\x95\x3d\xa3\x80\x09\x06\x14\xf5\x17\x34\x05\x5d\xfe\xc2\xea\x29\xdc\x92\xf2\x6d\x64\x95\xd3\x5e\x0b\x85\xb6\xb3\x02\x3e\xfd\x3a\xf9\x0d\xc7\xda\xef\x2d\xc9\x36\x8d\x23\xab\x58\xa2\x17\x70\xd7\x58\x3c\x29\x92\xc1\xb8\x37\xa1\xac\x6a\xc4\xb7\x8e\xfa\x80\xfe\xfb\x04\xf0\x7a\x47\x4e\x54\x3a\xad\x18\x6a\x47\x5f\x4b\x3e\x80\x35\x04\x76\xde\x9e\x97\x86\xb8\x82\x25\x45\x52\x9a\x5c\x87\xc7\x01\x24\x69\x74\x40\xd6\xcd\x1d\x7c\x7b\x37\x83\xc5\x54\xaf\xde\xb9\x96\xac\x73\x38\x8f\xb7\x1d\x73\x7b\xc6\x5a\xc7\xdb\x6d\x39\x8a\x53\x4b\x1c\xcd\x8f\x1d\x1d\x69\xb2\x0c\x75\x22\x3a\xe6\x0a\x93\xf7\x5d\x28\x74\xcf\xbe\x2f\xb2\x6a\x4c\x6f\x01\xab\x5c\xa9\x95\x51\x2c\x4b\x01\x91\x3a\x60\xe9\x3b\x51\xa4\x8b\xbe\x84\xd7\xd7\xf9\xfc\xb5\x9b\xf7\xe9\x57\x27\xf6\x38\xbe\x9a\x75\xb5\x9c\xbe\x9d\xdd\x9e\x74\xbf\x3b\x93\x5d\x32\x55\x6f\xc7\xa4\x92\x35\xed\xfa\xbd\x8d\x7f\x85\x21\x15\xaf\x33\x3c\xa9\x34\xaf\xa2\x2c\x57\xb3\x75\x14\x2f\xd7\xef\xf2\x08\x18\x77\x46\xa5\x5b\x9b\x23\x4c\x58\x93\xf7\x2b\x67\xb6\x74\x49\x97\x86\x60\xe7\x14\xfa\xb0\x6d\x4d\x7b\x5b\x65\x28\xff\xe9\xf3\x1b\x17\x04\x7c\xb9\xff\x72\x7f\xe1\x64\xcd\x81\x51\x4d\x49\x61\xb9\x21\x69\x74\x72\x72\x3b\x4e\x32\x90\x63\x93\xbc\x06\xdc\x7f\xee\x44\x28\x2e\xe8\x3f\xb1\xa7\x84\x2a\xa4\xff\x0b\xfc\x6f\x17\x01\x3b\x64\x95\x3b\x8a\x53\x47\x3e\x35\x2a\x11\x70\xf7\x79\xf4\x6f\x00\x00\x00\xff\xff\x3c\x46\x43\xf9\x36\x09\x00\x00")

func deployOperatorTplYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployOperatorTplYaml,
		"deploy/operator.tpl.yaml",
	)
}

func deployOperatorTplYaml() (*asset, error) {
	bytes, err := deployOperatorTplYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/operator.tpl.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployOperatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\xcb\x6e\xe2\x4c\x10\x85\xf7\x7e\x8a\x52\x58\xdb\x5c\xf2\xff\xb3\xe8\x9d\xc5\x45\x41\x0a\xd8\x02\x6b\xa2\x59\x8d\x6a\xda\x85\x69\xd1\xb7\xe9\x6e\x3b\xf2\xdb\x8f\x4c\x08\x31\xc1\x49\xad\xa0\xce\xa9\xf3\xb9\xba\x46\x23\x48\xeb\x60\x2a\xd2\xe4\x30\x50\x09\x07\x67\x14\x18\xdb\xfd\x33\x2e\x09\x56\x26\x2d\x2a\x09\x18\xe0\x85\x4a\x98\xfc\x0f\x69\x5d\xc1\x6c\x32\x9b\xc0\x74\xc6\xa6\x3f\xd8\x7f\x8f\x90\x6f\x60\xbe\xdc\x17\xd1\x68\x04\x8b\x0c\xb6\x59\x01\x9b\x6c\xb1\x5e\xfd\x82\xe2\x69\xbd\x87\xd5\xfa\x79\x99\x40\x2e\x09\x3d\x01\x3f\xa2\xae\x68\x20\x5f\x68\x1f\x08\xcb\x24\x8a\xd0\x8a\x9f\xe4\xbc\x30\x9a\x01\x5a\xeb\xc7\xcd\x34\x3a\x09\x5d\x32\x58\x90\x95\xa6\x55\xa4\x43\xa4\x28\x60\x89\x01\x59\x04\x20\xf1\x0f\x49\xdf\xfd\x82\x6e\x80\x81\xf0\x41\x98\xf8\xd5\xb8\x93\xb7\xc8\xe9\x2c\x34\xef\x91\xcd\x24\x99\x24\x8f\x11\x80\x46\x45\xf7\x5e\x6f\x89\x77\x51\x8e\xac\x14\x1c\x3d\x83\x69\x04\xe0\x49\x12\x0f\xc6\xbd\x41\x14\x06\x7e\x7c\xee\x51\xbf\xe1\x0e\x91\x03\x29\x2b\x31\xd0\x25\xad\xb7\x4a\x57\xf2\x26\xf8\xdb\xe8\xa1\xf0\xae\x86\x57\xeb\x94\xf7\xf5\xba\xe2\x46\x07\x14\x9a\xdc\x15\x16\x03\xba\xaa\x87\x8e\xc1\x93\x6b\x3e\x68\xdc\x28\x85\xba\xec\x1b\xc4\xe9\x43\x26\xdd\xf4\xa5\xb7\xaf\x78\x49\x8b\xf9\xd3\xef\x6d\xba\x59\xee\xf3\x74\xbe\xbc\xea\x00\x0d\xca\x9a\x18\x3c\x3c\xdc\xcd\xe4\xd9\xe2\x3c\xf1\xd9\xbc\x72\x46\xb1\x5e\x13\xe0\x20\x48\x96\x3b\x3a\xdc\x76\x2f\xfd\x1c\xc3\x91\x5d\xdf\x37\xe9\xb2\xef\x50\x59\xbe\xdc\xa5\x45\xb6\x1b\xe4\x7d\xfd\xee\x42\x61\x45\x0c\xfe\xd6\xd8\x26\xc2\x8c\x15\x0a\x1f\x1c\x8e\x3f\xd9\xd9\xcd\x55\x2e\x53\x79\x2d\x65\x6e\xa4\xe0\x2d\x83\x54\xbe\x62\xeb\xaf\xfa\xd7\x77\x83\xf3\x25\x04\xa7\x94\x73\x53\xeb\xb0\x1d\x74\xc6\x71\x1c\xfd\x0b\x00\x00\xff\xff\x00\xae\x8f\x15\xd1\x03\x00\x00")

func deployOperatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployOperatorYaml,
		"deploy/operator.yaml",
	)
}

func deployOperatorYaml() (*asset, error) {
	bytes, err := deployOperatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/operator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployRole_bindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x31\x6f\x1b\x31\x0c\x85\xf7\xfb\x15\x84\xd1\xd5\x57\x74\x2b\xb4\xb9\x85\xd1\xa9\x45\x61\x1b\xd9\x79\x12\xcf\xc7\xf8\x4e\x54\x28\xca\x06\x12\xe4\xbf\x07\x96\xcf\x71\x90\xc1\xd1\x26\x3e\xe2\xbd\xf7\xf1\xc0\x31\x38\xd8\xd1\x94\x46\x34\x6a\x30\xf1\x03\x69\x66\x89\x0e\x6c\x1e\xb6\x92\x28\xe6\x81\x7b\x6b\x59\xbe\x1f\x7f\x34\x09\x15\x27\x32\xd2\xec\x1a\x80\x25\x44\x9c\xc8\xc1\xbf\xd5\xdf\xf5\xf6\xff\xea\xf7\xba\x01\x00\x08\x94\xbd\x72\xb2\xea\xb4\xd8\xa1\xee\xc9\xea\x62\x4e\xe8\x09\x7a\x51\x38\x0d\xec\x07\x50\x19\x09\x3a\x8e\x81\xe3\x1e\xf2\x20\x65\x0c\xd0\x11\x04\xea\x39\x52\x58\x54\x33\xa5\xa7\xc2\x4a\xc1\x81\x69\xa1\x3a\x3a\xe2\x58\xc8\x01\x67\x63\x59\x9e\x44\x0f\xd5\x77\x29\x89\x14\x4d\xb4\x91\xee\x91\xbc\xcd\x05\x2f\x90\x1b\x19\xe9\xd7\x25\xa8\x5a\x7c\x64\xd5\x0e\x7d\x8b\xc5\x06\x51\x7e\xc6\x73\xeb\xf6\xf0\x33\xcf\xbc\xe7\xe5\x89\x0c\x03\x1a\xba\xfa\x83\x99\xf9\x53\x7c\xd5\x72\xb9\x45\x9f\xdf\x35\x7e\x4b\x7a\x64\x4f\x2b\xef\xa5\x44\x9b\xc5\x7b\x46\x57\xb5\x4e\x1c\x7c\x7b\x79\xbf\xf0\xeb\xe5\x2a\x32\xd2\x86\xfa\x6b\xce\x0d\xf2\xcb\x8a\x95\xfd\x8f\x4a\x49\x77\xc8\x9b\xb7\x00\x00\x00\xff\xff\x0b\xb5\x61\x6a\x1b\x02\x00\x00")

func deployRole_bindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployRole_bindingYaml,
		"deploy/role_binding.yaml",
	)
}

func deployRole_bindingYaml() (*asset, error) {
	bytes, err := deployRole_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xc1\x8e\xd4\x30\x0c\xbd\xf7\x2b\xa2\xbd\xac\x84\xd4\xae\xb8\xa1\xfe\x00\x37\x0e\x08\x71\xf7\xa6\xee\xac\x35\x49\x1c\xd9\x4e\x11\xf3\xf5\x28\x9d\xc2\x74\xa0\x1d\x54\x69\x7b\x6a\x1c\xdb\xcf\xef\xf9\x05\x32\x7d\x47\x51\xe2\xd4\x3b\x79\x05\xdf\x41\xb1\x37\x16\xba\x80\x11\xa7\xee\xfc\x49\x3b\xe2\x97\xe9\x63\x73\xa6\x34\xf4\xee\x2b\x07\x6c\x22\x1a\x0c\x60\xd0\x37\xce\x79\xc1\x39\xf3\x1b\x45\x54\x83\x98\x7b\x97\x4a\x08\x8d\x73\x09\x22\xf6\x8e\xd4\x88\xdb\x1f\x2c\x67\xcd\xe0\xb1\x91\x12\x50\x6b\x61\xeb\x20\xd3\x67\xe1\x92\xe7\x63\xfd\x5a\xf7\xf4\x34\xff\x0a\x2a\x17\xf1\xb8\xba\xc9\x3c\xe8\x9f\x83\xa2\x4c\xe4\xf1\x16\xc0\x34\x64\xa6\x64\xb7\x48\xae\xa4\xd4\x30\xd9\xc4\xa1\x44\xf4\x01\x28\xae\x0a\x26\x5c\x67\x7b\x4e\x23\x9d\x22\xe4\x35\x86\x17\x5c\x52\x26\x94\xd7\xd5\x2c\xcf\x1f\x9e\x8f\x13\xa8\x72\xcc\x12\x6c\xb6\x3c\xa1\xed\xb5\x84\xbc\x4c\xb5\xd1\x74\xc0\x1c\xf8\x67\xbc\xe3\x32\x00\x46\x4e\x8a\xab\x90\x60\x0e\xe4\xe1\x2e\xa6\x06\x86\x63\x09\x7a\x9c\x64\x9d\xa8\xe3\x8c\x49\xdf\x68\xb4\x8e\xf8\xff\xe3\x5d\x05\x3e\x0a\x14\x39\x91\xb1\x50\x3a\x75\x9e\x05\x59\x3b\xcf\x71\x0f\x6c\x31\xc5\x52\xf3\x40\xe5\x65\xe5\xd5\xb8\xb8\x87\x3c\xdb\x76\xc5\xf1\x01\xee\x75\xfe\x23\xb4\x12\x5a\x7d\x10\x95\xd6\x15\x67\x5f\xc1\xe3\xcd\x23\x90\x9a\xc0\xbb\xf5\xdc\x30\xe0\x97\x6a\xe5\xdf\xb7\x7f\x3f\xf0\x0d\xd0\x3b\x9f\xbe\x8c\x94\x20\xd0\x05\xff\x5d\x51\xeb\x4a\x1e\xea\x52\x7e\x05\x00\x00\xff\xff\x2e\xd1\xd4\xd7\x95\x04\x00\x00")

func deployRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployRoleYaml,
		"deploy/role.yaml",
	)
}

func deployRoleYaml() (*asset, error) {
	bytes, err := deployRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployService_accountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x31\x0e\x80\x30\x08\x00\xc0\x9d\x57\xf0\x01\x07\x57\x36\xdf\x60\xe2\x4e\x28\x03\x69\x0a\x4d\xc1\xfa\x7d\x8f\xa7\x3d\xba\xd2\xc2\x09\xf7\x09\xdd\xbc\x11\xde\xba\xb6\x89\x5e\x22\xf1\x7a\xc1\xd0\xe2\xc6\xc5\x04\x88\xce\x43\x09\x2d\xcb\xe2\xf8\x62\xf5\x9c\x2c\x0a\x7f\x00\x00\x00\xff\xff\x94\xa0\xb7\x3f\x46\x00\x00\x00")

func deployService_accountYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployService_accountYaml,
		"deploy/service_account.yaml",
	)
}

func deployService_accountYaml() (*asset, error) {
	bytes, err := deployService_accountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/service_account.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateStrategies_basicRemoveTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8e\xbd\x4e\xc4\x30\x10\x06\xfb\x7b\x0a\xcb\xf5\xe9\x4c\x4d\x4d\x41\x49\x81\xe8\x17\xfb\x03\x2c\xc5\x3f\xda\xdd\xa4\xb1\xfc\xee\x28\xa1\x89\x40\x28\x8e\x74\xa5\xad\xfd\x66\xa6\x35\x13\x3f\xcc\xed\x89\x94\x6e\xcf\x24\xc6\x3a\xa9\xf0\x4e\x91\xea\x44\x8a\x9f\x97\x2f\x59\x29\x66\xb0\xb8\x07\x37\xc5\x05\x19\x22\x2f\x5c\xde\x61\x4d\xef\x97\x66\x4b\xb5\x8f\xc6\x32\x52\x59\x60\xaf\xc6\x56\xd2\xaf\xf5\xe7\x24\xac\x5f\x2f\xad\x19\xe4\xb0\x51\x4f\x97\x31\x28\xc4\xbb\xa5\xfd\xa2\x1d\xb4\x25\x28\x05\x52\x72\x0c\x29\x33\x7b\xbc\x81\x25\x96\x7c\x54\xf1\xef\x6e\xd4\xf7\x89\x0c\x26\x3d\xa3\xda\x4d\x46\x2d\x73\x0c\xc3\xf8\xf5\x76\x94\xeb\x19\x5b\xc8\x6b\x4c\x10\xa5\x54\x87\x2d\x7f\x97\x7d\xa7\xfc\x0e\x00\x00\xff\xff\x77\xe1\xaa\x7c\xd7\x02\x00\x00")

func templateStrategies_basicRemoveTplBytes() ([]byte, error) {
	return bindataRead(
		_templateStrategies_basicRemoveTpl,
		"template/strategies/_basic-remove.tpl",
	)
}

func templateStrategies_basicRemoveTpl() (*asset, error) {
	bytes, err := templateStrategies_basicRemoveTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/strategies/_basic-remove.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateStrategies_basicVersionTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x4d\x6f\xf2\x30\x10\x84\xef\xfc\x8a\xd5\x9e\xde\x57\x02\x7c\xe7\x5a\x2a\xf5\x50\xf5\xc8\x7d\xeb\x2c\x25\xaa\xbf\x1a\x1b\x2a\x64\xf9\xbf\x57\x09\x1f\x05\x0a\xc5\x31\xf4\x98\xc4\x3b\xf3\x64\x76\xe4\x18\xa1\x9e\x83\xb1\x01\xfe\x8d\xa7\x14\x68\xfc\x44\x1e\x50\x78\xc7\x52\x04\xd6\x4e\x51\x60\xa1\x39\x50\x45\x81\xf0\x3f\xa4\x34\x88\x68\x1d\x4e\x00\xa9\xaa\x70\x08\xe8\x28\x2c\xda\xc7\x4b\x33\x43\xc0\x15\xa9\x25\xe3\x04\x62\x4a\xc3\x41\x8c\xc0\xa6\xea\x84\xfa\x78\x0b\x45\xaf\xac\x7c\x09\xc2\x6e\xf4\x1a\x49\x36\x84\x58\x71\xe3\x6b\x6b\xf0\x90\x45\x5a\xb7\x6e\x2d\xe6\x8d\xd5\xd7\x61\xf6\x12\xd9\xf8\xbb\x89\x91\xb7\xcb\x46\x32\xb6\x3f\xb0\xb5\x6e\xd8\x29\x92\xdc\x5f\xeb\x30\x12\x8c\x71\xfc\xc2\x9f\xb3\xcd\x97\x94\xf0\xe6\x5d\xed\x5d\xca\x77\x76\x5f\x50\xcf\x8a\x65\xb0\x4d\x06\xd0\xfe\x68\x5e\x69\x1e\x3f\x96\xa4\x00\xc5\x7b\x6d\x2a\x04\x9c\xb2\x53\x76\xad\xd9\x84\xae\x22\x00\x39\x54\x42\x53\x90\x8b\xe7\x83\x9a\x03\x64\x31\x1e\x0d\x9e\xf2\x76\xde\x5b\xe2\x1d\x47\x16\xc2\x51\xc9\xbf\x49\x2e\x76\xed\x57\x8d\x2b\xeb\x3b\x07\x99\x1f\xd6\x71\xcf\x0a\x42\x2b\xc1\xec\xd9\x83\x07\x6b\xe6\xf5\x1b\xe6\x6d\xa1\x34\xf9\xbf\x4b\xbb\x28\xe1\xbb\xa5\xba\x21\xca\xb9\x84\xcf\x65\x74\xfb\x95\x92\x23\x6e\x48\xf3\xa9\x64\x07\x3f\x6b\x5f\xfc\x38\x99\xd2\xe8\x8c\xe7\x57\x00\x00\x00\xff\xff\x82\xd4\x0f\xab\x8e\x07\x00\x00")

func templateStrategies_basicVersionTplBytes() ([]byte, error) {
	return bindataRead(
		_templateStrategies_basicVersionTpl,
		"template/strategies/_basic-version.tpl",
	)
}

func templateStrategies_basicVersionTpl() (*asset, error) {
	bytes, err := templateStrategies_basicVersionTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/strategies/_basic-version.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateStrategiesPreparedImageTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x31\x6b\xc3\x30\x10\x85\x77\xff\x8a\xc7\x4d\x2d\xb8\x52\xbb\x7a\xee\xd0\x5f\xd0\xa5\x94\x72\x95\xaf\x8d\xc0\x96\x84\xa4\x78\x11\xfa\xef\x41\xb1\x33\x78\x48\x20\x19\x8f\x7b\xdf\xf7\xe0\x7d\x75\x1d\x50\x0a\xb2\xcc\x61\xe2\x2c\xa0\x9f\x5f\x4e\xd6\xbc\x2c\x12\x93\xf5\x8e\xa0\x50\xeb\x16\xb2\x7f\x70\x3e\xe3\x49\xbd\x73\x66\xf5\xc1\x09\xa4\x53\x10\xa3\x2f\xf4\x7a\x45\x09\x93\x35\x9c\xe8\xb9\xa1\x40\x21\x1f\x68\x00\xf1\x38\x52\x0f\x0a\x9c\x0f\xed\xbc\x89\xf6\xa0\x85\xa7\xa3\xd0\x80\x52\x6b\xbf\xf6\x8b\x1b\xf7\xc6\x16\x67\x23\x8f\x58\xe9\x8d\x56\xed\x7d\x2a\xe3\x5d\x66\xeb\x24\x26\xfd\xaa\xed\xcc\xff\xb2\x93\x96\xa2\x3e\x39\x26\x75\xfe\xd4\xda\x2a\xae\xec\x1b\x65\xf6\x8b\x6c\xf3\x7e\x77\xa7\x00\x00\x00\xff\xff\x0a\x21\x5d\x3a\x88\x01\x00\x00")

func templateStrategiesPreparedImageTplBytes() ([]byte, error) {
	return bindataRead(
		_templateStrategiesPreparedImageTpl,
		"template/strategies/prepared-image.tpl",
	)
}

func templateStrategiesPreparedImageTpl() (*asset, error) {
	bytes, err := templateStrategiesPreparedImageTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/strategies/prepared-image.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateStrategiesPreparedImageVar = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xcc\x4d\x4c\x4f\xb5\xe5\x02\x04\x00\x00\xff\xff\xe0\xf3\xaa\xf1\x07\x00\x00\x00")

func templateStrategiesPreparedImageVarBytes() ([]byte, error) {
	return bindataRead(
		_templateStrategiesPreparedImageVar,
		"template/strategies/prepared-image.var",
	)
}

func templateStrategiesPreparedImageVar() (*asset, error) {
	bytes, err := templateStrategiesPreparedImageVarBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/strategies/prepared-image.var", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateStrategiesTelepresenceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\x5f\x6b\xdb\x30\x14\xc5\xdf\xfd\x29\x2e\xf7\x69\x83\xc4\x5e\xdf\x86\xdf\x42\xaa\xb1\xc2\xe6\x85\xb4\xe4\xa5\x94\x70\x63\x5f\x77\x62\xb2\x64\x24\xcd\x1b\x08\x7f\xf7\xe1\x7f\x69\xdd\xd2\x91\xa4\x0f\x09\x28\x9c\xfb\x3b\x07\x9d\x1b\x85\x00\x25\x49\x75\x53\xee\xc8\x4a\x3a\x28\xbe\x36\xec\x32\xe3\xc5\x5f\xe9\x3c\xc4\x3b\xb2\x0e\xb0\x61\xeb\xa4\xd1\x08\xcb\xb6\x8d\xa2\xfb\x08\x20\x04\xf0\x5c\xd5\x8a\x3c\x03\xee\x0f\xe4\x64\xbe\x3c\xaa\x62\xe8\x64\xbd\x48\x96\xa0\x8d\x87\x0f\xf1\x35\x79\x8a\xbf\x92\x03\x4c\x5c\xcd\x79\x32\x4d\x0f\x27\xcb\xb5\x92\x39\x39\xfc\xd8\x8d\x02\x04\x34\x35\xa6\x80\x54\x14\xb8\x00\xac\xc9\xff\xec\x8e\xff\x1d\x5d\x00\x36\xa4\x7e\x33\xa6\x10\xda\x76\x31\xf8\xb3\x2e\xe6\xc4\x4e\x4e\x39\x5f\x42\xc5\x2b\x1c\xb0\x27\x85\xab\xd8\x53\x41\x9e\x12\x45\x07\x56\x2e\xf1\xac\xb8\xb6\xec\x58\x0f\xee\x47\xaa\x67\xe7\xe7\xe0\x13\x33\xe6\x46\x7b\x92\x9a\xad\x4b\x3e\x25\xb2\xa2\xc7\x39\xb7\x33\xff\x23\x2d\xcf\x9c\x97\xbf\x3e\xbb\x34\x84\xbe\xd7\x78\x2c\xac\x6d\x71\xba\xae\x13\xeb\x9a\x39\xb3\x6e\x2e\xa8\xed\x15\xe2\x59\xf4\xfb\x87\x37\xeb\xbb\x80\x9c\x2c\x67\xab\x11\x01\x00\xa0\xa6\xaa\xbf\xa4\x3b\xf1\x4d\x6c\xb6\xe2\x56\x64\x6b\xb1\x5f\xff\xc8\xee\x56\x37\x99\xd8\xee\xb3\xd5\x77\x71\xbb\x59\xad\x05\x2e\x06\x7d\x3f\xfe\xc5\x9a\xea\x88\x00\xc0\x52\xb2\x2a\xb6\x5c\x3e\xfb\x0d\x00\xa9\x96\xbb\xf1\x9f\x90\x02\x36\x57\x23\xe2\x69\x62\x33\x86\x9f\x36\x24\xee\xc2\xb8\xba\x6b\x7c\x54\xb6\xd1\xf4\xdd\x7f\x9e\xca\x39\xa7\x17\xb2\x8f\x0e\x5f\x2e\x7f\x65\x9a\xf3\xf6\xaa\xa7\xbc\xae\xe3\xec\x34\xb9\xa9\x2a\xd2\xc5\xfb\x03\x4d\xa0\x17\x99\xde\x78\x93\x46\x83\xe1\x49\x7a\x88\xfe\x05\x00\x00\xff\xff\xe0\x41\x6e\xb4\xee\x04\x00\x00")

func templateStrategiesTelepresenceTplBytes() ([]byte, error) {
	return bindataRead(
		_templateStrategiesTelepresenceTpl,
		"template/strategies/telepresence.tpl",
	)
}

func templateStrategiesTelepresenceTpl() (*asset, error) {
	bytes, err := templateStrategiesTelepresenceTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/strategies/telepresence.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateStrategiesTelepresenceVar = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4b\x2d\x2a\xce\xcc\xcf\xb3\xe5\x02\x04\x00\x00\xff\xff\xa0\x0b\xe9\x34\x09\x00\x00\x00")

func templateStrategiesTelepresenceVarBytes() ([]byte, error) {
	return bindataRead(
		_templateStrategiesTelepresenceVar,
		"template/strategies/telepresence.var",
	)
}

func templateStrategiesTelepresenceVar() (*asset, error) {
	bytes, err := templateStrategiesTelepresenceVarBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/strategies/telepresence.var", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"deploy/cluster_role_binding.yaml":       deployCluster_role_bindingYaml,
	"deploy/cluster_role.yaml":               deployCluster_roleYaml,
	"deploy/crds/maistra.io_sessions.yaml":   deployCrdsMaistraIo_sessionsYaml,
	"deploy/operator.tpl.yaml":               deployOperatorTplYaml,
	"deploy/operator.yaml":                   deployOperatorYaml,
	"deploy/role_binding.yaml":               deployRole_bindingYaml,
	"deploy/role.yaml":                       deployRoleYaml,
	"deploy/service_account.yaml":            deployService_accountYaml,
	"template/strategies/_basic-remove.tpl":  templateStrategies_basicRemoveTpl,
	"template/strategies/_basic-version.tpl": templateStrategies_basicVersionTpl,
	"template/strategies/prepared-image.tpl": templateStrategiesPreparedImageTpl,
	"template/strategies/prepared-image.var": templateStrategiesPreparedImageVar,
	"template/strategies/telepresence.tpl":   templateStrategiesTelepresenceTpl,
	"template/strategies/telepresence.var":   templateStrategiesTelepresenceVar,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"deploy": &bintree{nil, map[string]*bintree{
		"cluster_role.yaml":         &bintree{deployCluster_roleYaml, map[string]*bintree{}},
		"cluster_role_binding.yaml": &bintree{deployCluster_role_bindingYaml, map[string]*bintree{}},
		"crds": &bintree{nil, map[string]*bintree{
			"maistra.io_sessions.yaml": &bintree{deployCrdsMaistraIo_sessionsYaml, map[string]*bintree{}},
		}},
		"operator.tpl.yaml":    &bintree{deployOperatorTplYaml, map[string]*bintree{}},
		"operator.yaml":        &bintree{deployOperatorYaml, map[string]*bintree{}},
		"role.yaml":            &bintree{deployRoleYaml, map[string]*bintree{}},
		"role_binding.yaml":    &bintree{deployRole_bindingYaml, map[string]*bintree{}},
		"service_account.yaml": &bintree{deployService_accountYaml, map[string]*bintree{}},
	}},
	"template": &bintree{nil, map[string]*bintree{
		"strategies": &bintree{nil, map[string]*bintree{
			"_basic-remove.tpl":  &bintree{templateStrategies_basicRemoveTpl, map[string]*bintree{}},
			"_basic-version.tpl": &bintree{templateStrategies_basicVersionTpl, map[string]*bintree{}},
			"prepared-image.tpl": &bintree{templateStrategiesPreparedImageTpl, map[string]*bintree{}},
			"prepared-image.var": &bintree{templateStrategiesPreparedImageVar, map[string]*bintree{}},
			"telepresence.tpl":   &bintree{templateStrategiesTelepresenceTpl, map[string]*bintree{}},
			"telepresence.var":   &bintree{templateStrategiesTelepresenceVar, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

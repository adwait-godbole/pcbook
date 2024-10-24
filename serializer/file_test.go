package serializer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adwait-godbole/pcbook/pb"
	"github.com/adwait-godbole/pcbook/sample"
	"github.com/adwait-godbole/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	// tmpDir := t.TempDir()
	tmpDir := "../tmp"
	err := os.MkdirAll(tmpDir, 0755)
	require.NoError(t, err)

	// t.Cleanup(func() {
	// 	os.RemoveAll(tmpDir)
	// })

	binaryFile := filepath.Join(tmpDir, "laptop.bin")
	jsonFile := filepath.Join(tmpDir, "laptop.json")

	laptop1 := sample.NewLaptop()
	err = serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}

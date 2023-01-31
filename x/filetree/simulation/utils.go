package simulation

import (
	"crypto/sha256"
	"encoding/binary"
	"io"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	folderNames = []string{"filetree", "storage", "jackal", "canine"}
	fileNames   = []string{"rns.txt", "cat.png", "dog.jpeg", "meow.mp3", "wolf.mp4"}
)

const maxDepth = 5

// Randomly generate a file directory using account address and returns paths of files
// ie. for path: s/home/document/meow.md
// returns {"s","s/home", "s/home/document", "s/home/document/meow.md"}
func GenerateDirectory(address sdk.AccAddress) (paths []string, err error) {
	// Generate seed with account address
	h := sha256.New()
	_, err = io.WriteString(h, address.String())
	if err != nil {
		return nil, err
	}
	seed := binary.BigEndian.Uint64(h.Sum(nil))
	r := rand.New(rand.NewSource(int64(seed)))

	dir := make([]string, r.Intn(maxDepth)+3)
	for i := 0; i < len(dir); i++ {
		f := folderNames[r.Intn(len(folderNames))]
		dir = append(dir, f)
	}
	// complete directory by adding s/home at the front and random file at the end
	dir = append([]string{"s", "home"}, dir...)
	dir = append(dir, fileNames[r.Intn(len(fileNames))])

	paths[0] = dir[0]
	for i := 1; i < len(dir); i++ {
		paths = append(paths, paths[i-1]+"/"+dir[i])
	}

	return
}

// Returns seven predefined paths
// ie. "s", "s/home", "s/home/filetree", "s/home/filetree/meow.mp3"
func GetDirectory() (paths []string) {
	return []string{
		"s/home",
		"s/home/filetree",
		"s/home/filetree/meow.mp3",
		"s/home/jackal",
		"s/home/jackal/meow.mp3",
		"s/home/jackal/wolf.mp4",
	}
}

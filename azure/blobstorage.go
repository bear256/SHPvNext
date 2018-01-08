package azure

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/storage"
)

type Blob struct {
	Layer  string
	Path   string
	Blocks []*Block
}

type Block struct {
	Id   string
	Recs []*Record
}

type Record struct {
	Id      uint32
	Content []byte
}

func NewBlob(layer, path string) *Blob {
	blocks := make([]*Block, 0)
	return &Blob{layer, path, blocks}
}

func NewBlock(id string) *Block {
	recs := make([]*Record, 0)
	return &Block{id, recs}
}

func NewRecord(idStr string, content []byte) *Record {
	id, _ := strconv.Atoi(idStr)
	return &Record{uint32(id), content}
}

func (block Block) Marshal() []byte {
	var buf bytes.Buffer
	for _, rec := range block.Recs {
		buf.Write(rec.Marshal())
	}
	return buf.Bytes()
}

func (block *Block) Add(rec *Record) {
	block.Recs = append(block.Recs, rec)
}

func (block *Block) Update(rec Record) {
	for _, item := range block.Recs {
		if item.Id == rec.Id {
			item.Content = rec.Content
			break
		}
	}
}

func (rec Record) Marshal() []byte {
	var buf bytes.Buffer
	size := uint32(4 + len(rec.Content))
	bSize := make([]byte, 4)
	binary.LittleEndian.PutUint32(bSize, size)
	buf.Write(bSize)
	bId := make([]byte, 4)
	binary.LittleEndian.PutUint32(bId, rec.Id)
	buf.Write(bId)
	buf.Write(rec.Content)
	return buf.Bytes()
}

type Store struct {
	Client storage.Client
}

func (store Store) BlobStream(layer, path string) (chan []byte, error) {
	client := store.Client
	service := client.GetBlobService()
	container := service.GetContainerReference(layer)
	container.CreateIfNotExists(nil)
	blob := container.GetBlobReference(path)
	rd, err := blob.Get(nil)
	if err != nil {
		return nil, err
	} else {
		ch := make(chan []byte, 100)
		go func() {
			for {
				bSize := make([]byte, 4)
				if _, err := io.ReadAtLeast(rd, bSize, 4); err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println("Read bSize error:", err)
				}
				size := binary.LittleEndian.Uint32(bSize)
				fmt.Println(size)
				bId := make([]byte, 4)
				if _, err := io.ReadAtLeast(rd, bId, 4); err != nil {
					fmt.Println("Read bId error:", err)
				}
				id := binary.LittleEndian.Uint32(bId)
				fmt.Println(id)
				content := make([]byte, size-4)
				if _, err := io.ReadAtLeast(rd, content, int(size-4)); err != nil {
					fmt.Println("Read content error:", err)
				}
				ch <- content
			}
		}()
		return ch, nil
	}

}

func (store Store) BlockExists(layer, path, blockId string) (bool, uint64, uint64) {
	client := store.Client
	service := client.GetBlobService()
	container := service.GetContainerReference(layer)
	container.CreateIfNotExists(nil)
	blob := container.GetBlobReference(path)
	resp, _ := blob.GetBlockList(storage.BlockListTypeCommitted, nil)
	flag, start, end := false, uint64(0), uint64(0)
	for _, committed := range resp.CommittedBlocks {
		id, _ := base64.StdEncoding.DecodeString(committed.Name)
		if string(id) == blockId {
			flag = true
			end = start + uint64(committed.Size)
			break
		} else {
			start += uint64(committed.Size)
		}
	}
	return flag, start, end
}

func (store Store) BlockStream(layer, path, blockId string) (chan []byte, error) {
	flag, start, end := store.BlockExists(layer, path, blockId)
	if flag == false {
		return nil, fmt.Errorf("Block not exists")
	}
	client := store.Client
	service := client.GetBlobService()
	container := service.GetContainerReference(layer)
	blob := container.GetBlobReference(path)
	bRange := &storage.BlobRange{start, end - 1}
	rd, err := blob.GetRange(&storage.GetBlobRangeOptions{bRange, false, nil})
	if err != nil {
		return nil, err
	} else {
		ch := make(chan []byte, 100)
		go func() {
			for {
				bSize := make([]byte, 4)
				if _, err := io.ReadAtLeast(rd, bSize, 4); err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println("Read bSize error:", err)
				}
				size := binary.LittleEndian.Uint32(bSize)
				fmt.Println(size)
				bId := make([]byte, 4)
				if _, err := io.ReadAtLeast(rd, bId, 4); err != nil {
					fmt.Println("Read bId error:", err)
				}
				id := binary.LittleEndian.Uint32(bId)
				fmt.Println(id)
				content := make([]byte, size-4)
				if _, err := io.ReadAtLeast(rd, content, int(size-4)); err != nil {
					fmt.Println("Read content error:", err)
				}
				ch <- content
			}
		}()
		return ch, nil
	}
}

func (store Store) BlockRead(layer, path, blockId string) Block {
	flag, start, end := store.BlockExists(layer, path, blockId)
	if flag == false {
		return nil
	}
	client := store.Client
	service := client.GetBlobService()
	container := service.GetContainerReference(layer)
	blob := container.GetBlobReference(path)
	bRange := &storage.BlobRange{start, end - 1}
	rd, err := blob.GetRange(&storage.GetBlobRangeOptions{bRange, false, nil})
	if err != nil {
		return nil, err
	} else {
		block := NewBlock(blockId)
		for {
			bSize := make([]byte, 4)
			if _, err := io.ReadAtLeast(rd, bSize, 4); err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Read bSize error:", err)
			}
			size := binary.LittleEndian.Uint32(bSize)
			fmt.Println(size)
			bId := make([]byte, 4)
			if _, err := io.ReadAtLeast(rd, bId, 4); err != nil {
				fmt.Println("Read bId error:", err)
			}
			id := binary.LittleEndian.Uint32(bId)
			fmt.Println(id)
			content := make([]byte, size-4)
			if _, err := io.ReadAtLeast(rd, content, int(size-4)); err != nil {
				fmt.Println("Read content error:", err)
			}
			block.Add(&Record{id, content})
		}
		return block
	}
}

func UpdateRecord(blob *storage.Blob, blockId string) {

}

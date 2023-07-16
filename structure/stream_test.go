package structure

import (
	"github.com/ByteStorage/FlyDB/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func initStreamDB() (*StreamStructure, *config.Options) {
	opts := config.DefaultOptions
	dir, _ := os.MkdirTemp("", "TestStreamStructure_Get")
	opts.DirPath = dir
	stc, _ := NewStreamStructure(opts)
	return stc, &opts
}

func TestStreamStructure_XRead(t *testing.T) {
	stc, _ := initStreamDB()
	defer stc.db.Clean()

	ok1, err := stc.XAdd("test", "1", map[string]interface{}{"name1": "flydb1"})
	assert.Nil(t, err)
	assert.True(t, ok1)

	ok2, err := stc.XAdd("test", "2", map[string]interface{}{"name2": "flydb2"})
	assert.Nil(t, err)
	assert.True(t, ok2)

	ok3, err := stc.XAdd("test", "3", map[string]interface{}{"name3": "flydb3"})
	assert.Nil(t, err)
	assert.True(t, ok3)

	ok4, err := stc.XAdd("test1", "1", map[string]interface{}{"name11": "flydb11"})
	assert.Nil(t, err)
	assert.True(t, ok4)

	item, err := stc.XRead("test", 1)
	assert.NotNil(t, item)
	assert.Nil(t, err)

	item1, err := stc.XRead("test1", 1)
	assert.NotNil(t, item1)
	assert.Nil(t, err)

}

func TestStreamStructure_XDel(t *testing.T) {
	stc, _ := initStreamDB()
	defer stc.db.Clean()

	ok1, err := stc.XAdd("test", "1", map[string]interface{}{"name1": "flydb1"})
	assert.Nil(t, err)
	assert.True(t, ok1)

	ok2, err := stc.XAdd("test", "2", map[string]interface{}{"name2": "flydb2"})
	assert.Nil(t, err)
	assert.True(t, ok2)

	ok3, err := stc.XAdd("test", "3", map[string]interface{}{"name3": "flydb3"})
	assert.Nil(t, err)
	assert.True(t, ok3)

	ok4, err := stc.XAdd("test1", "1", map[string]interface{}{"name11": "flydb11"})
	assert.Nil(t, err)
	assert.True(t, ok4)

	item, err := stc.XRead("test", 1)
	assert.Nil(t, err)
	assert.NotNil(t, item)

	item1, err := stc.XRead("test1", 1)
	assert.Nil(t, err)
	assert.NotNil(t, item1)

	ok1, length1, err := stc.XDel("test", "1")
	assert.True(t, ok1)
	assert.Nil(t, err)
	assert.Equal(t, 2, length1)

	ok2, length2, err := stc.XDel("test", "2")
	assert.True(t, ok2)
	assert.Nil(t, err)
	assert.Equal(t, 1, length2)

	ok3, length3, err := stc.XDel("test", "3")
	assert.True(t, ok3)
	assert.Nil(t, err)
	assert.Equal(t, 0, length3)

	ok4, length4, err := stc.XDel("test1", "1")
	assert.True(t, ok4)
	assert.Nil(t, err)
	assert.Equal(t, 0, length4)

	item, err = stc.XRead("test", 1)
	assert.NotNil(t, err)
	assert.Nil(t, item)

	item1, err = stc.XRead("test1", 1)
	assert.NotNil(t, err)
	assert.Nil(t, item1)

}

func TestStreamStructure_XLen(t *testing.T) {
	stc, _ := initStreamDB()
	defer stc.db.Clean()

	ok1, err := stc.XAdd("test", "1", map[string]interface{}{"name1": "flydb1"})
	assert.Nil(t, err)
	assert.True(t, ok1)

	ok2, err := stc.XAdd("test", "2", map[string]interface{}{"name2": "flydb2"})
	assert.Nil(t, err)
	assert.True(t, ok2)

	ok3, err := stc.XAdd("test", "3", map[string]interface{}{"name3": "flydb3"})
	assert.Nil(t, err)
	assert.True(t, ok3)

	ok4, err := stc.XAdd("test1", "1", map[string]interface{}{"name11": "flydb11"})
	assert.Nil(t, err)
	assert.True(t, ok4)

	length, err := stc.XLen("test")
	assert.Nil(t, err)
	assert.Equal(t, 3, length)

	length, err = stc.XLen("test1")
	assert.Nil(t, err)
	assert.Equal(t, 1, length)

}

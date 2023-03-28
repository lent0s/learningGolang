// works with db file
package database

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type config struct {
	dbDir    string
	attempts int
}

func readConfig() *config {

	cfg, err := os.ReadFile("./cmd/config.ini")
	if err != nil {
		log.Fatalf("access to config file denied (#rC): %v", err)
	}
	configStr := strings.Split(string(cfg), "\n")

	attempts, err := strconv.Atoi(configStr[1][22:])
	if err != nil {
		log.Fatalf("invalid data in config file "+
			"(attempts to start db) (#rC): %v", err)
	}

	return &config{
		dbDir:    configStr[0][21 : len(configStr[0])-1],
		attempts: attempts,
	}
}

type Storage struct {
	users   map[int]*User
	fileDB  *os.File
	changed time.Time
}

func NewDatabase() *Storage {

	return &Storage{
		users:   make(map[int]*User),
		fileDB:  openDb(),
		changed: time.Unix(0, 0),
	}
}

func openDb() (f *os.File) {

	cfg := readConfig()
	problem := true
	var err error

	for i := 0; i < cfg.attempts; i++ {
		filename := fmt.Sprintf("%sstorage_%d.db", cfg.dbDir, i)
		f, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			continue
		}

		fileInfo, err := f.Stat()
		if err != nil {
			continue
		}
		if fileInfo.Size() == 0 {
			if _, err = f.Write([]byte("[\n    {\n    }\n]")); err != nil {
				continue
			}
			problem = false
			break
		}

		buf, err := io.ReadAll(f)
		if err != nil {
			continue
		}
		if _, err = readData(buf); err != nil {
			continue
		}
		problem = false
		break
	}
	if problem {
		log.Fatalf("access to db denied (#o): %v", err)
	}
	defer f.Close()
	return f
}

func fileInfo(f *os.File) (fs.FileInfo, error) {

	file, err := os.Open(f.Name())
	if err != nil {
		return nil, fmt.Errorf("access to db denied (#fI),"+
			" try again later: %s", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("access to db denied (#fI),"+
			" try again later: %s", err)
	}

	return fileInfo, nil
}

func dbNotChanged(s *Storage) (bool, error) {

	info, err := fileInfo(s.fileDB)
	if err != nil {
		return false, err
	}
	if info.ModTime().UnixNano() == s.changed.UnixNano() {
		return true, nil
	}
	return false, nil
}

func readDb(s *Storage) error {

	notChanged, err := dbNotChanged(s)
	if err != nil {
		return err
	}
	if notChanged {
		return nil
	}

	buf, err := os.ReadFile(s.fileDB.Name())
	if err != nil {
		return fmt.Errorf("access to db denied (#rD): %s", err)
	}
	users, err := readData(buf)
	if err != nil {
		return err
	}

	for id := 1; id < len(users); id++ {
		if users[id].Name == "" {
			s.users[id] = nil
			continue
		}
		s.users[id] = users[id]
	}

	info, err := fileInfo(s.fileDB)
	if err != nil {
		return err
	}
	s.changed = info.ModTime()
	return nil
}

func appendDb(f *os.File, data []byte) error {

	dbFile, err := os.OpenFile(f.Name(), os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("access to db denied (#aD): %s", err)
	}
	defer dbFile.Close()

	reData := []byte(",\n    " + string(data) + "\n]")
	info, err := fileInfo(dbFile)
	if err != nil {
		return err
	}
	if _, err := dbFile.WriteAt(reData, info.Size()-2); err != nil {
		return fmt.Errorf("access to db denied (#aD): %s", err)
	}
	return nil
}

func updateDb(f *os.File, data []byte, id int) error {

	dbFile, err := os.OpenFile(f.Name(), os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("access to db denied (#uD): %s", err)
	}
	defer dbFile.Close()
	buf, err := io.ReadAll(dbFile)
	if err != nil {
		return fmt.Errorf("access to db denied (#uD): %s", err)
	}

	buf = updateData(buf, data, id)
	if _, err = dbFile.WriteAt(buf, 0); err != nil {
		return fmt.Errorf("access to db denied (#uD): %s", err)
	}
	if err = dbFile.Truncate(int64(len(buf))); err != nil {
		return fmt.Errorf("access to db denied (#uD): %s", err)
	}
	return nil
}

package machines

import (
	"database/sql"
	"fmt"
)

type Machine struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Space int64  `json:"space"`
}

type Disk struct {
	Id    int64
	Space int64
}

type Connect struct {
	DiskId      int64
	MachineName string
}

type Storage struct {
	Db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{Db: db}
}

func (s *Storage) ListMachines() ([]*Machine, error) {
	rows, err := s.Db.Query(`
		SELECT m.id as id, m.name as name, COALESCE(SUM(d.space), 0) as space
		FROM machines as m
		LEFT JOIN disks as d ON d.machineid = m.id GROUP BY m.id, m.name
		LIMIT 200;`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res = make([]*Machine, 0)
	for rows.Next() {
		var c Machine
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	return res, nil
}

func (s *Storage) Connect(connect Connect) error {
	if len(connect.MachineName) < 0 {
		return fmt.Errorf("Machine name is not provided")
	}
	_, err := s.Db.Exec(`
		UPDATE disks
		SET machineid = (SELECT id FROM machines WHERE name = $1)
		WHERE id = $2`, connect.MachineName, connect.DiskId)
	return err
}

package snowflake

import (
	"sync"
	"time"
)

type Snowflake struct {
	mutex         sync.Mutex
	sequence      int64
	lastTimestamp int64
	machineID     int64
}

func NewSnowflake(machineID int64) *Snowflake {
	return &Snowflake{
		machineID: machineID & 0x3FF, // マシンIDは10bitを使用
	}
}

func (s *Snowflake) Generate() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 現在のタイムスタンプをミリ秒単位で取得
	ts := time.Now().UnixNano() / 1000000

	// システム時計が巻き戻された場合はエラーを返す
	if ts < s.lastTimestamp {
		panic("system clock was rewound")
	}

	// 同じミリ秒内の場合、シーケンス番号をインクリメント
	if ts == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & 0xFFF // シーケンス番号は12bitを使用
		if s.sequence == 0 {
			// シーケンス番号が枯渇した場合は次のミリ秒まで待機
			ts = s.waitNextMillis(ts)
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = ts

	// IDの生成
	// - 41bit: タイムスタンプ
	// - 10bit: マシンID
	// - 12bit: シーケンス番号
	id := (ts << 22) | (s.machineID << 12) | s.sequence
	return id
}

func (s *Snowflake) waitNextMillis(ts int64) int64 {
	for ts == s.lastTimestamp {
		ts = time.Now().UnixNano() / 1000000
	}
	return ts
}

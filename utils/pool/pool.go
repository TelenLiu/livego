package pool

type Pool struct {
	pos int
	buf []byte
}

//const maxpoolsize = 500 * 1024
const maxpoolsize = 20 * 1024 * 1024

func (pool *Pool) Get(size int) []byte {
	if maxpoolsize-pool.pos < size {
		pool.pos = 0
		pool.buf = make([]byte, maxpoolsize)
	}
	b := pool.buf[pool.pos : pool.pos+size]
	pool.pos += size
	return b
}

func NewPool() *Pool {
	return &Pool{
		buf: make([]byte, maxpoolsize),
	}
}

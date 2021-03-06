// This is a very basic example of a program that implements rdb.decoder and
// outputs a human readable diffable dump of the rdb file.
package main

import (
	"fmt"
	"os"

	"github.com/titanous/rdb"
)

type decoder struct {
	db int
	i  int
}

func (p *decoder) StartRDB() {
}

func (p *decoder) StartDatabase(n int) {
	p.db = n
}

func (p *decoder) EndDatabase(n int) {
}

func (p *decoder) EndRDB() {
}

func (p *decoder) Set(key, value []byte, expiry int64) {
	fmt.Printf("db=%d %q -> %q\n", p.db, key, value)
}

func (p *decoder) StartHash(key []byte, length, expiry int64) {
}

func (p *decoder) Hset(key, field, value []byte) {
	fmt.Printf("db=%d %q . %q -> %q\n", p.db, key, field, value)
}

func (p *decoder) EndHash(key []byte) {
}

func (p *decoder) StartSet(key []byte, cardinality, expiry int64) {
}

func (p *decoder) Sadd(key, member []byte) {
	fmt.Printf("db=%d %q { %q }\n", p.db, key, member)
}

func (p *decoder) EndSet(key []byte) {
}

func (p *decoder) StartList(key []byte, length, expiry int64) {
	p.i = 0
}

func (p *decoder) Rpush(key, value []byte) {
	fmt.Printf("db=%d %q[%d] -> %q\n", p.db, key, p.i, value)
	p.i++
}

func (p *decoder) EndList(key []byte) {
}

func (p *decoder) StartZSet(key []byte, cardinality, expiry int64) {
	p.i = 0
}

func (p *decoder) Zadd(key []byte, score float64, member []byte) {
	fmt.Printf("db=%d %q[%d] -> {%q, score=%g}\n", p.db, key, p.i, member, score)
	p.i++
}

func (p *decoder) EndZSet(key []byte) {
}

func maybeFatal(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	maybeFatal(err)
	err = rdb.Decode(f, &decoder{})
	maybeFatal(err)
}

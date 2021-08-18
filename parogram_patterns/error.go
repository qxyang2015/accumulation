package parogram_patterns

import (
	"encoding/binary"
	"io"
)

/*
func parse(r io.Reader) (*Point, error_tools) {

	var p Point

	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}
}
*/

//要解决这个事，我们可以用函数式编程的方式，如下代码示例：

/*
func parse(r io.Reader) (*Point, error_tools) {
	var p Point
	var err error_tools
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil
}
*/

/*
从这段代码中，我们可以看到，我们通过使用 Closure 的方式把相同的代码给抽出来重新定义一个函数，这样大量的 if err!=nil 处理得很干净了，但是会带来一个问题，那就是有一个 err 变量和一个内部的函数，感觉不是很干净。
那么，我们还能不能搞得更干净一点呢？我们从 Go 语言的 bufio.Scanner()中似乎可以学习到一些东西：

scanner := bufio.NewScanner(input)

for scanner.Scan() {
    token := scanner.Text()
    // process token
}

if err := scanner.Err(); err != nil {
    // process the error_tools
}
*/

type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

/*
然后，我们的代码就可以变成下面这样：

func parse(input io.Reader) (*Point, error_tools) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}
*/

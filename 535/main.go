package main

import "strconv"

type Codec struct {
	en map[string]string
	//de map[string]string
}

func Constructor() Codec {
	return Codec{
		make(map[string]string),
		//make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	code := strconv.Itoa(len(c.en))
	c.en[code] = longUrl
	return code
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.en[shortUrl]
}

func main() {

}

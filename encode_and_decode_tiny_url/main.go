package main

import (
"strings"
"strconv"	
)

var shortUrlPrefix = "http://tinyurl.com/"

type Codec struct {
	index int
	dict map[int]string
}

func Constructor() Codec {
	return Codec{index:0, dict: map[int]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) Encode(longUrl string) string {
	this.index++
	this.dict[this.index] = longUrl
	return shortUrlPrefix + strconv.Itoa(this.index)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) Decode(shortUrl string) string {
	index := strings.Replace(shortUrl, shortUrlPrefix, "", 1)
	i, _ := strconv.Atoi(index)
	v, _ := this.dict[i]
	return v
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

type Codec struct {
    str string 
    code2url map[string]string
    url2code map[string]string
}


func Constructor() Codec {
    return Codec{
        str: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
        code2url: make(map[string]string),
        url2code: make(map[string]string),
    }
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if code, ok := this.url2code[longUrl]; ok {
        return "http://tinyurl.com/"+code
    }
    arr := make([]string,6)
    for i:=0;i<6;i++{
        arr[i] = string(this.str[rand.Intn(len(this.str))])
    }
    code := strings.Join(arr,"")
    this.url2code[longUrl] = code
    this.code2url[code] = longUrl
    return "http://tinyurl.com/"+code
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    code := shortUrl[len(shortUrl)-6:]
    if url,ok := this.code2url[code];ok{
        return url
    }
    return ""
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */


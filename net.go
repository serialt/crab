package crab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const ipv4_regex = `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`

// SubNetMaskToLen ipv4 子网掩码长度换算
// 如 255.255.255.0 对应的网络位长度为 24
func SubNetMaskToLen(netmask string) (int, error) {
	ipSplitArr := strings.Split(netmask, ".")
	if len(ipSplitArr) != 4 {
		return 0, fmt.Errorf("netmask:%v is not valid, pattern should like: 255.255.255.0", netmask)
	}
	ipv4MaskArr := make([]byte, 4)
	for i, value := range ipSplitArr {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return 0, fmt.Errorf("ipMaskToInt call strconv.Atoi error:[%v] string value is: [%s]", err, value)
		}
		if intValue > 255 {
			return 0, fmt.Errorf("netmask cannot greater than 255, current value is: [%s]", value)
		}
		ipv4MaskArr[i] = byte(intValue)
	}

	ones, _ := net.IPv4Mask(ipv4MaskArr[0], ipv4MaskArr[1], ipv4MaskArr[2], ipv4MaskArr[3]).Size()
	return ones, nil

}

// LenToSubNetMask ipv4 网络位长度转换为子网掩码地址
// 如 24 对应的子网掩码地址为 255.255.255.0
func LenToSubNetMask(subnet int) string {
	var buff bytes.Buffer
	for i := 0; i < subnet; i++ {
		buff.WriteString("1")
	}
	for i := subnet; i < 32; i++ {
		buff.WriteString("0")
	}
	masker := buff.String()
	a, _ := strconv.ParseUint(masker[:8], 2, 64)
	b, _ := strconv.ParseUint(masker[8:16], 2, 64)
	c, _ := strconv.ParseUint(masker[16:24], 2, 64)
	d, _ := strconv.ParseUint(masker[24:32], 2, 64)
	resultMask := fmt.Sprintf("%v.%v.%v.%v", a, b, c, d)
	return resultMask

}

// IsPublicIPv4 ipv4 判断是否是公网ip
func IsPublicIPv4(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// 获取所在区域的公网ip的网站
var urlList = []string{
	"https://ip.tool.lu",
	"http://cip.cc",
}

func SetPubIpUrl(uri ...string) {
	urlList = uri
}

// IPGet 返回客户端 IP
func IPGet(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// 获取公网ip, 如果两个ip不同，则访问ip.tool.lu 和false, ip都相同则返回true
func GetPubIP() (ip string, all bool) {
	var ipList []string
	for _, url := range urlList {
		client := &http.Client{}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
		request.Header.Add("Connection", "keep-alive")
		request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36")
		resp, err := client.Do(request)
		if resp.StatusCode != 200 && err != nil {
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		reg := regexp.MustCompile(ipv4_regex)
		ipList = reg.FindAllString(string(body), -1)

	}
	if len(ipList) > 0 {
		// fmt.Printf("my public ip is: %s\n", ipList[0])
		ip = ipList[0]
		if ipList[0] == ipList[1] {
			all = true

		} else {
			all = false
		}

	}
	return
}

type Location struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

var MyIPLocationURL = "https://ipinfo.io"

func SetIPLocationURL(ipURL string) {
	MyIPLocationURL = ipURL
}
func GetMyIPLocation() (lo *Location, err error) {
	lo = &Location{}
	resp, err := http.Get(MyIPLocationURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, lo)

	return
}

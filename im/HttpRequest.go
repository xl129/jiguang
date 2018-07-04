package im

import (
	"net/http"
	"io"
	"net/url"
	"bytes"
	"time"
	"net"
	"io/ioutil"
	"os"
	"mime/multipart"
	"path/filepath"
)

type Response struct {
	Err        error  // 请求错误
	Data       []byte // 请求返回的字节
	StatusCode int
	Status     string
}

type HttpRequest struct {
	connectTimeout   time.Duration
	readWriteTimeout time.Duration
}

// GET请求
func (h *HttpRequest) Get(uri string, query map[string]string) Response {
	val := url.Values{}
	for k, v := range query {
		val.Add(k, v)
	}
	str := val.Encode()
	if len(str) > 0 {
		uri = uri + "?" + str
	}

	return h.request(uri, "GET", nil, nil)
}

// POST 请求   外部传json的字节
func (h *HttpRequest) Post(uri string, data []byte) Response {
	body := bytes.NewReader(data)
	return h.request(uri, "POST", body, nil)
}

// PUT 请求
func (h *HttpRequest) Put(uri string, data []byte) Response {
	body := bytes.NewReader(data)
	return h.request(uri, "PUT", body, nil)
}

// DELETE 请求
func (h *HttpRequest) Delete(uri string, data []byte) Response {
	body := bytes.NewReader(data)
	return h.request(uri, "DELETE", body, nil)
}

// 上传文件  fileType = "image", "file", "voice"
func (h *HttpRequest) Upload(uri, path string) Response {
	Response := Response{}
	headers := map[string]string{}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		Response.Err = err
		return Response
	}
	body := bytes.NewBufferString("")
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("filename", filepath.Base(path))
	if err != nil {
		Response.Err = err
		return Response
	}
	//这里的io.Copy实现, 会把file文件都读取到内存里面，然后当做一个buffer传给NewRequest.对于大文件来说会占用很多内存
	_, err = io.Copy(part, file)
	// 修改header
	headers["Content-Type"] = writer.FormDataContentType()
	err = writer.Close()
	if err != nil {
		Response.Err = err
		return Response
	}

	return h.request(uri, "POST", body, headers)
}

//发起网路请求
func (h *HttpRequest) request(uri, method string, body io.Reader, headers map[string]string) Response {
	Response := Response{}
	//生成client 参数为默认
	trans := &http.Transport{
		TLSClientConfig: nil,
		Proxy:           nil,
		Dial:            TimeoutDialer(h.connectTimeout, h.readWriteTimeout),
	}
	client := &http.Client{Transport: trans}
	req, err := http.NewRequest(method, uri, body)

	if err != nil {
		Response.Err = err
		return Response
	}
	default_headers := map[string]string{
		"Content-Type": "application/json",
	}
	if headers == nil {
		headers = default_headers
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Charset", "utf-8")
	req.Header.Add("Connection", "Keep-Alive")
	//添加认证
	req.Header.Add("Authorization", getAuthorization())
	//处理返回结果
	response, err := client.Do(req)
	defer response.Body.Close()
	//解析
	if err != nil {
		Response.Err = err
		return Response
	}
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Response.Err = err
		return Response
	}
	Response.Status = response.Status
	Response.StatusCode = response.StatusCode
	Response.Data = data
	Response.Err = nil
	return Response
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

var httpRequestInstance *HttpRequest

func GetHttpRequestInstance() *HttpRequest {
	if httpRequestInstance != nil {
		return httpRequestInstance
	}
	httpRequestInstance = &HttpRequest{
		connectTimeout:   60 * time.Second,
		readWriteTimeout: 60 * time.Second,
	}

	return httpRequestInstance
}

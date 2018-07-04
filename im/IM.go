package im

type IM struct {
	uri    string
	config *Config
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (i *IM) setUri(uri string) {
	i.uri = uri
}

func (i *IM) post(data []byte) Response {
	return GetHttpRequestInstance(i.config).Post(i.uri, data)
}

func (i *IM) put(data []byte) Response {
	return GetHttpRequestInstance(i.config).Put(i.uri, data)
}

func (i *IM) upload(path string) Response {
	return GetHttpRequestInstance(i.config).Upload(i.uri, path)
}

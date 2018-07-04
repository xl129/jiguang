package im

type IM struct {
	uri string
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (i *IM) setUri(uri string) {
	i.uri = uri
}

func (i *IM) post(data []byte) Response {
	return GetHttpRequestInstance().Post(i.uri, data)
}

func (i *IM) put(data []byte) Response {
	return GetHttpRequestInstance().Put(i.uri, data)
}

func (i *IM) upload(path string) Response {
	return GetHttpRequestInstance().Upload(i.uri, path)
}

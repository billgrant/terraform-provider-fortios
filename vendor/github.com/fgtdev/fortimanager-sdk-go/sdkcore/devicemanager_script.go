package fmgclient

import (
	"fmt"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

type JSONDVMScript struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Content     string `json:"content"`
	Target      string `json:"target"`
	Type        string `json:"type"`
}

// Create and Update function
func (c *FmgSDKClient) CreateUpdateDVMScript(params *JSONDVMScript, method string) (err error) {
	defer c.Trace("CreateUpdateDVMScript")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/dvmdb/adom/root/script",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateDVMScript failed: %s", err)
		return
	}

	return
}

func (c *FmgSDKClient) ReadDVMScript(id string) (out *JSONDVMScript, err error) {
	defer c.Trace("ReadDVMScript")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/root/script/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadDVMScript failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONDVMScript{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["desc"] != nil {
		out.Description = data["desc"].(string)
	}
	if data["content"] != nil {
		out.Content = data["content"].(string)
	}
	if data["target"] != nil {
		out.Target = util.ScriptTarget2Str(int(data["target"].(float64)))
	}

	return
}

func (c *FmgSDKClient) DeleteDVMScript(id string) (err error) {
	defer c.Trace("DeleteDVMScript")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/root/script/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteDVMScript failed :%s", err)
		return
	}

	return
}

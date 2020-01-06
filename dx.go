package gotag

import (
	"context"
	"fmt"
	"net"

	mxtag_pb "github.com/MOXA-ISD/gotag/protobuf"
	"github.com/golang/protobuf/proto"
)

const (
	sockBufferSize = 1024
	domainSockPath = "/var/tag/taghubd.sock"
)

func NewDXApi() *DXApi {
	dx := DXApi{}
	return &dx
}

type DXApi struct {
	ctx    context.Context
	cancel context.CancelFunc
	conn   *net.UnixConn
}

func (d *DXApi) sendRequest(data []byte) ([]byte, error) {
	// send request
	_, err := d.conn.Write(data)
	if err != nil {
		fmt.Printf("DIAL: Write Error: %v\n", err)
		return nil, err
	}
	fmt.Println("cccccccccccccccc")
	// wait for response
	buf := make([]byte, sockBufferSize)
	n, err := d.conn.Read(buf)
	fmt.Println("ddddddddddddd")
	if err != nil && n == 0 {
		fmt.Printf("read error: %v, %d\n", err.Error(), n)
		return nil, fmt.Errorf("DIAL: Read Error: %v", err.Error())
	}
	payload := make([]byte, n)
	copy(payload, buf)
	return payload, nil
}

func (d *DXApi) GetTagValue(module, device, tag string) *Tag {
	var (
		dxCmdGetTagValue string = "dx_get_tag_value"
		cmd     mxtag_pb.Request
		reqData mxtag_pb.RequestData
		data    []byte
		err     error
		result  []Tag
	)
	reqData = mxtag_pb.RequestData{
		Module: &module,
		Device: &device,
		Name:   &tag,
	}
	cmd = mxtag_pb.Request{
		Method: &dxCmdGetTagValue,
		List: []*mxtag_pb.RequestData{
			&reqData,
		},
	}
	if data, err = proto.Marshal(&cmd); err != nil {
		return nil
	}

	resp := mxtag_pb.Response{}
	payload, err := d.sendRequest(data)
	if err := proto.Unmarshal(payload, &resp); err != nil {
		return nil
	}

	result = ProtobufToTag(resp.List)
	/*for _, tag := range resp.List {
		result = append(result, Tag{
			sourceName: tag.GetEquipment(),
			tagName:    tag.GetTag(),
		})
	}*/
	if len(result) == 0 {
		fmt.Println("result len is 0")
		return nil
	}
	return &result[0]
}

func (d *DXApi) GetTagList() {

}

func (d *DXApi) Dail() error {
	addr, err := net.ResolveUnixAddr("unix", domainSockPath)
	if err != nil {
		fmt.Printf("Failed to resolve: %v\n", err)
		return err
	}
	d.conn, err = net.DialUnix("unix", nil, addr)
	if err != nil {
		fmt.Printf("Failed to dial: %v\n", err)
		return err
	}
	return nil
}

func (d *DXApi) Close() {
	d.conn.Close()
}

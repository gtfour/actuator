package run

import "flag"

var Props = GetProps()

func GetProps()(props map[string]string){

    props=make(map[string]string,0)

    // serve
    // flag.Lookup

    ip_version_ptr       := flag.String("ip_version","v4","ip version")
    ip_port_ptr          := flag.String("port","80","port number")
    ip_addr_ptr          := flag.String("ip_addr","0.0.0.0","ip addr")

    flag.Parse()
    //out_file_ptr       := flag.String("outfile","out.txt","Out file")

    ip_version       := *ip_version_ptr
    ip_port          := *ip_port_ptr
    ip_addr          := *ip_addr_ptr

    props["ip_version"]  = ip_version
    props["ip_port"]     = ip_port
    props["ip_addr"]     = ip_addr

    props["server_addr"] = props["ip_addr"]+":"+props["ip_port"]

    return

}

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/lixiangzhong/dnsutil"
)

func DigMain( params []string ) {
	var dig dnsutil.Dig

	dns := "8.8.8.8"
	dtype := "A"
	d := 0

	if( len(params)==0 ) {
		fmt.Println( "Usage: dig [[@DNS] type] Question" )
		os.Exit(0)
	} else if ( len(params)==2 ) {
		dtype = strings.ToUpper(params[0])
		d = 1
	} else if ( len(params)==3 ) {
		dns = strings.TrimPrefix(params[0],"@")
		dtype = strings.ToUpper(params[1])
		d = 2
	}

	dig.SetDNS(dns)
	dig.SetTimeOut(3*time.Second)

	for i:=d ; i<len(params) ; i++ {
		switch dtype {
		    case "A": if a, err := dig.A( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "AAAA": if a, err := dig.AAAA( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "CAA": if a, err := dig.CAA( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "CNAME": if a, err := dig.CNAME( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "MX": if a, err := dig.MX( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "NS": if a, err := dig.NS( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "PTR": if a, err := dig.PTR( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "SPF": if a, err := dig.SPF( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "SRV": if a, err := dig.SRV( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
			case "TXT": if a, err := dig.TXT( params[i] ); err==nil { if(len(a)>0 ) { for j:=0;j<len(a);j++ { fmt.Println(a[j]) } } }
		default: fmt.Println("Unknown DNS command")
		}
	}
}

func main() {
	DigMain(os.Args[1:])
}
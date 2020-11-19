package services

var SERVICE_MAP = ServiceMap{
	"1ci-smcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "1ci-smcs",
			Port:  3091,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "1ci-smcs",
			Port:  3091,
		},
	},
	"2ping": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "2ping",
			Port:  15998,
		},
	},
	"3Com-nsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3Com-nsd",
			Port:  1742,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3Com-nsd",
			Port:  1742,
		},
	},
	"3com-amp3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3com-amp3",
			Port:  629,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3com-amp3",
			Port:  629,
		},
	},
	"3com-net-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3com-net-mgmt",
			Port:  2391,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3com-net-mgmt",
			Port:  2391,
		},
	},
	"3com-njack-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3com-njack-1",
			Port:  5264,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3com-njack-1",
			Port:  5264,
		},
	},
	"3com-njack-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3com-njack-2",
			Port:  5265,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3com-njack-2",
			Port:  5265,
		},
	},
	"3com-webview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3com-webview",
			Port:  2339,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3com-webview",
			Port:  2339,
		},
	},
	"3comfaxrpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3comfaxrpc",
			Port:  3446,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3comfaxrpc",
			Port:  3446,
		},
	},
	"3comnetman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3comnetman",
			Port:  1181,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3comnetman",
			Port:  1181,
		},
	},
	"3d-nfsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3d-nfsd",
			Port:  2323,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3d-nfsd",
			Port:  2323,
		},
	},
	"3ds-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3ds-lm",
			Port:  1538,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3ds-lm",
			Port:  1538,
		},
	},
	"3exmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3exmp",
			Port:  5221,
		},
	},
	"3gpp-cbsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3gpp-cbsp",
			Port:  48049,
		},
	},
	"3l-l1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3l-l1",
			Port:  1511,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3l-l1",
			Port:  1511,
		},
	},
	"3link": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3link",
			Port:  15363,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3link",
			Port:  15363,
		},
	},
	"3m-image-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3m-image-lm",
			Port:  1550,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3m-image-lm",
			Port:  1550,
		},
	},
	"3par-evts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3par-evts",
			Port:  5781,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3par-evts",
			Port:  5781,
		},
	},
	"3par-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3par-mgmt",
			Port:  5782,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3par-mgmt",
			Port:  5782,
		},
	},
	"3par-mgmt-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3par-mgmt-ssl",
			Port:  5783,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3par-mgmt-ssl",
			Port:  5783,
		},
	},
	"3par-rcopy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "3par-rcopy",
			Port:  5785,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "3par-rcopy",
			Port:  5785,
		},
	},
	"4-tieropmcli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "4-tieropmcli",
			Port:  2934,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "4-tieropmcli",
			Port:  2934,
		},
	},
	"4-tieropmgw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "4-tieropmgw",
			Port:  2933,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "4-tieropmgw",
			Port:  2933,
		},
	},
	"4talk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "4talk",
			Port:  3284,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "4talk",
			Port:  3284,
		},
	},
	"6a44": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "6a44",
			Port:  1027,
		},
	},
	"802-11-iapp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "802-11-iapp",
			Port:  3517,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "802-11-iapp",
			Port:  3517,
		},
	},
	"914c/g": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "914c/g",
			Port:  211,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "914c/g",
			Port:  211,
		},
	},
	"9pfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "9pfs",
			Port:  564,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "9pfs",
			Port:  564,
		},
	},
	"BESApi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "BESApi",
			Port:  3408,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "BESApi",
			Port:  3408,
		},
	},
	"CAIlic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "CAIlic",
			Port:  216,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "CAIlic",
			Port:  216,
		},
	},
	"CodeMeter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "CodeMeter",
			Port:  22350,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "CodeMeter",
			Port:  22350,
		},
	},
	"DMExpress": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "DMExpress",
			Port:  32636,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "DMExpress",
			Port:  32636,
		},
	},
	"EtherNet/IP-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "EtherNet/IP-1",
			Port:  2222,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "EtherNet/IP-1",
			Port:  2222,
		},
	},
	"EtherNet/IP-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "EtherNet/IP-2",
			Port:  44818,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "EtherNet/IP-2",
			Port:  44818,
		},
	},
	"LiebDevMgmt_A": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "LiebDevMgmt_A",
			Port:  3029,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "LiebDevMgmt_A",
			Port:  3029,
		},
	},
	"LiebDevMgmt_C": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "LiebDevMgmt_C",
			Port:  3027,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "LiebDevMgmt_C",
			Port:  3027,
		},
	},
	"LiebDevMgmt_DM": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "LiebDevMgmt_DM",
			Port:  3028,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "LiebDevMgmt_DM",
			Port:  3028,
		},
	},
	"MOS-aux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "MOS-aux",
			Port:  10542,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "MOS-aux",
			Port:  10542,
		},
	},
	"MOS-lower": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "MOS-lower",
			Port:  10540,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "MOS-lower",
			Port:  10540,
		},
	},
	"MOS-soap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "MOS-soap",
			Port:  10543,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "MOS-soap",
			Port:  10543,
		},
	},
	"MOS-soap-opt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "MOS-soap-opt",
			Port:  10544,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "MOS-soap-opt",
			Port:  10544,
		},
	},
	"MOS-upper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "MOS-upper",
			Port:  10541,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "MOS-upper",
			Port:  10541,
		},
	},
	"MaxumSP": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "MaxumSP",
			Port:  4179,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "MaxumSP",
			Port:  4179,
		},
	},
	"MobilitySrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "MobilitySrv",
			Port:  6997,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "MobilitySrv",
			Port:  6997,
		},
	},
	"Mon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "Mon",
			Port:  9255,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "Mon",
			Port:  9255,
		},
	},
	"PowerAlert-nsa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "PowerAlert-nsa",
			Port:  4150,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "PowerAlert-nsa",
			Port:  4150,
		},
	},
	"SunVTS-RMI": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "SunVTS-RMI",
			Port:  6483,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "SunVTS-RMI",
			Port:  6483,
		},
	},
	"WibuKey": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "WibuKey",
			Port:  22347,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "WibuKey",
			Port:  22347,
		},
	},
	"XSIP-network": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "XSIP-network",
			Port:  1354,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "XSIP-network",
			Port:  1354,
		},
	},
	"XmlIpcRegSvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "XmlIpcRegSvc",
			Port:  9092,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "XmlIpcRegSvc",
			Port:  9092,
		},
	},
	"a1-bs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a1-bs",
			Port:  5603,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a1-bs",
			Port:  5603,
		},
	},
	"a1-msc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a1-msc",
			Port:  5602,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a1-msc",
			Port:  5602,
		},
	},
	"a13-an": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a13-an",
			Port:  3125,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a13-an",
			Port:  3125,
		},
	},
	"a14": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a14",
			Port:  3597,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a14",
			Port:  3597,
		},
	},
	"a15": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a15",
			Port:  3598,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a15",
			Port:  3598,
		},
	},
	"a16-an-an": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a16-an-an",
			Port:  4598,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a16-an-an",
			Port:  4598,
		},
	},
	"a17-an-an": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a17-an-an",
			Port:  4599,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a17-an-an",
			Port:  4599,
		},
	},
	"a21-an-1xbs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a21-an-1xbs",
			Port:  4597,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a21-an-1xbs",
			Port:  4597,
		},
	},
	"a25-fap-fgw": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "a25-fap-fgw",
			Port:  4502,
		},
	},
	"a26-fap-fgw": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "a26-fap-fgw",
			Port:  4726,
		},
	},
	"a27-ran-ran": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "a27-ran-ran",
			Port:  28119,
		},
	},
	"a3-sdunode": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a3-sdunode",
			Port:  5604,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a3-sdunode",
			Port:  5604,
		},
	},
	"a4-sdunode": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "a4-sdunode",
			Port:  5605,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "a4-sdunode",
			Port:  5605,
		},
	},
	"aairnet-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aairnet-1",
			Port:  3618,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aairnet-1",
			Port:  3618,
		},
	},
	"aairnet-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aairnet-2",
			Port:  3619,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aairnet-2",
			Port:  3619,
		},
	},
	"aal-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aal-lm",
			Port:  1469,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aal-lm",
			Port:  1469,
		},
	},
	"aamp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aamp",
			Port:  3939,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aamp",
			Port:  3939,
		},
	},
	"aap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aap",
			Port:  2878,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aap",
			Port:  2878,
		},
	},
	"aas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aas",
			Port:  1601,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aas",
			Port:  1601,
		},
	},
	"abacus-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abacus-remote",
			Port:  2894,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abacus-remote",
			Port:  2894,
		},
	},
	"abarsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abarsd",
			Port:  8402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abarsd",
			Port:  8402,
		},
	},
	"abatemgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abatemgr",
			Port:  3655,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abatemgr",
			Port:  3655,
		},
	},
	"abatjss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abatjss",
			Port:  3656,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abatjss",
			Port:  3656,
		},
	},
	"abb-escp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abb-escp",
			Port:  6316,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abb-escp",
			Port:  6316,
		},
	},
	"abbaccuray": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abbaccuray",
			Port:  1546,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abbaccuray",
			Port:  1546,
		},
	},
	"abbs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abbs",
			Port:  4885,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abbs",
			Port:  4885,
		},
	},
	"abcsoftware": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abcsoftware",
			Port:  3996,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abcsoftware",
			Port:  3996,
		},
	},
	"abcvoice-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abcvoice-port",
			Port:  3781,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abcvoice-port",
			Port:  3781,
		},
	},
	"about": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "about",
			Port:  2019,
		},
	},
	"abr-api": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abr-api",
			Port:  1954,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abr-api",
			Port:  1954,
		},
	},
	"abr-secure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "abr-secure",
			Port:  1955,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "abr-secure",
			Port:  1955,
		},
	},
	"ac-cluster": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ac-cluster",
			Port:  18463,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ac-cluster",
			Port:  18463,
		},
	},
	"ac-tech": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ac-tech",
			Port:  2796,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ac-tech",
			Port:  2796,
		},
	},
	"acap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acap",
			Port:  674,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acap",
			Port:  674,
		},
	},
	"acas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acas",
			Port:  62,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acas",
			Port:  62,
		},
	},
	"acc-raid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acc-raid",
			Port:  2800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acc-raid",
			Port:  2800,
		},
	},
	"accel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "accel",
			Port:  4108,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "accel",
			Port:  4108,
		},
	},
	"accelenet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "accelenet",
			Port:  1182,
		},
	},
	"accelenet-data": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "accelenet-data",
			Port:  1182,
		},
	},
	"accessnetwork": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "accessnetwork",
			Port:  699,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "accessnetwork",
			Port:  699,
		},
	},
	"accord-mgc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "accord-mgc",
			Port:  1205,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "accord-mgc",
			Port:  1205,
		},
	},
	"acctopus-cc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acctopus-cc",
			Port:  6868,
		},
	},
	"acctopus-st": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "acctopus-st",
			Port:  6868,
		},
	},
	"accu-lmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "accu-lmgr",
			Port:  7781,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "accu-lmgr",
			Port:  7781,
		},
	},
	"accuracer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "accuracer",
			Port:  12007,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "accuracer",
			Port:  12007,
		},
	},
	"accuracer-dbms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "accuracer-dbms",
			Port:  12008,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "accuracer-dbms",
			Port:  12008,
		},
	},
	"acd-pm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acd-pm",
			Port:  8793,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acd-pm",
			Port:  8793,
		},
	},
	"ace-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ace-client",
			Port:  2334,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ace-client",
			Port:  2334,
		},
	},
	"ace-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ace-proxy",
			Port:  2335,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ace-proxy",
			Port:  2335,
		},
	},
	"ace-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ace-server",
			Port:  2475,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ace-server",
			Port:  2475,
		},
	},
	"ace-svr-prop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ace-svr-prop",
			Port:  2476,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ace-svr-prop",
			Port:  2476,
		},
	},
	"aci": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aci",
			Port:  187,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aci",
			Port:  187,
		},
	},
	"acis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acis",
			Port:  9953,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acis",
			Port:  9953,
		},
	},
	"acl-manager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acl-manager",
			Port:  4013,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acl-manager",
			Port:  4013,
		},
	},
	"acmaint_dbd": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "acmaint_dbd",
			Port:  774,
		},
	},
	"acmaint_transd": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "acmaint_transd",
			Port:  775,
		},
	},
	"acme": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acme",
			Port:  9216,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acme",
			Port:  9216,
		},
	},
	"acms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acms",
			Port:  3980,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acms",
			Port:  3980,
		},
	},
	"acmsoda": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acmsoda",
			Port:  6969,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acmsoda",
			Port:  6969,
		},
	},
	"acnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acnet",
			Port:  6801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acnet",
			Port:  6801,
		},
	},
	"acp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acp",
			Port:  599,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acp",
			Port:  599,
		},
	},
	"acp-conduit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acp-conduit",
			Port:  3823,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acp-conduit",
			Port:  3823,
		},
	},
	"acp-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acp-discovery",
			Port:  3822,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acp-discovery",
			Port:  3822,
		},
	},
	"acp-policy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acp-policy",
			Port:  3824,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acp-policy",
			Port:  3824,
		},
	},
	"acp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acp-port",
			Port:  2071,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acp-port",
			Port:  2071,
		},
	},
	"acp-proto": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acp-proto",
			Port:  4046,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acp-proto",
			Port:  4046,
		},
	},
	"acplt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acplt",
			Port:  7509,
		},
	},
	"acptsys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acptsys",
			Port:  2149,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acptsys",
			Port:  2149,
		},
	},
	"acr-nema": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acr-nema",
			Port:  104,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acr-nema",
			Port:  104,
		},
	},
	"acter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "acter",
			Port:  4671,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "acter",
			Port:  4671,
		},
	},
	"actifio-c2c": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "actifio-c2c",
			Port:  5103,
		},
	},
	"activememory": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "activememory",
			Port:  2859,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "activememory",
			Port:  2859,
		},
	},
	"activesync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "activesync",
			Port:  1034,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "activesync",
			Port:  1034,
		},
	},
	"actnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "actnet",
			Port:  5411,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "actnet",
			Port:  5411,
		},
	},
	"ada-cip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ada-cip",
			Port:  2085,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ada-cip",
			Port:  2085,
		},
	},
	"adap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adap",
			Port:  6350,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adap",
			Port:  6350,
		},
	},
	"adapt-sna": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adapt-sna",
			Port:  1365,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adapt-sna",
			Port:  1365,
		},
	},
	"adaptecmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adaptecmgr",
			Port:  2521,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adaptecmgr",
			Port:  2521,
		},
	},
	"adcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adcp",
			Port:  7508,
		},
	},
	"adi-gxp-srvprt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adi-gxp-srvprt",
			Port:  6769,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adi-gxp-srvprt",
			Port:  6769,
		},
	},
	"admind": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "admind",
			Port:  3279,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "admind",
			Port:  3279,
		},
	},
	"admind2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "admind2",
			Port:  8403,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "admind2",
			Port:  8403,
		},
	},
	"admins-lms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "admins-lms",
			Port:  2692,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "admins-lms",
			Port:  2692,
		},
	},
	"adobeserver-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adobeserver-1",
			Port:  1102,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adobeserver-1",
			Port:  1102,
		},
	},
	"adobeserver-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adobeserver-2",
			Port:  1103,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adobeserver-2",
			Port:  1103,
		},
	},
	"adobeserver-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adobeserver-3",
			Port:  3703,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adobeserver-3",
			Port:  3703,
		},
	},
	"adobeserver-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adobeserver-4",
			Port:  3704,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adobeserver-4",
			Port:  3704,
		},
	},
	"adobeserver-5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adobeserver-5",
			Port:  3705,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adobeserver-5",
			Port:  3705,
		},
	},
	"adrep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adrep",
			Port:  3954,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adrep",
			Port:  3954,
		},
	},
	"ads": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ads",
			Port:  2550,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ads",
			Port:  2550,
		},
	},
	"ads-c": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "ads-c",
			Port:  5913,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "ads-c",
			Port:  5913,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ads-c",
			Port:  5913,
		},
	},
	"adtech-test": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adtech-test",
			Port:  3357,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adtech-test",
			Port:  3357,
		},
	},
	"adtempusclient": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adtempusclient",
			Port:  3760,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "adtempusclient",
			Port:  3760,
		},
	},
	"advant-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "advant-lm",
			Port:  2295,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "advant-lm",
			Port:  2295,
		},
	},
	"adws": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "adws",
			Port:  9389,
		},
	},
	"aed-512": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aed-512",
			Port:  149,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aed-512",
			Port:  149,
		},
	},
	"aegate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aegate",
			Port:  4549,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aegate",
			Port:  4549,
		},
	},
	"aequus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aequus",
			Port:  23456,
		},
	},
	"aequus-alt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aequus-alt",
			Port:  23457,
		},
	},
	"aero": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "aero",
			Port:  8060,
		},
	},
	"aeroflight-ads": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aeroflight-ads",
			Port:  1218,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aeroflight-ads",
			Port:  1218,
		},
	},
	"aeroflight-ret": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aeroflight-ret",
			Port:  1219,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aeroflight-ret",
			Port:  1219,
		},
	},
	"aes-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aes-discovery",
			Port:  3224,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aes-discovery",
			Port:  3224,
		},
	},
	"aes-x170": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "aes-x170",
			Port:  7107,
		},
	},
	"aesop": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "aesop",
			Port:  8202,
		},
	},
	"af": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "af",
			Port:  1411,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "af",
			Port:  1411,
		},
	},
	"afbackup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afbackup",
			Port:  2988,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afbackup",
			Port:  2988,
		},
	},
	"afesc-mc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afesc-mc",
			Port:  6628,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afesc-mc",
			Port:  6628,
		},
	},
	"affiliate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "affiliate",
			Port:  6579,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "affiliate",
			Port:  6579,
		},
	},
	"afore-vdp-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "afore-vdp-disc",
			Port:  4362,
		},
	},
	"afpovertcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afpovertcp",
			Port:  548,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afpovertcp",
			Port:  548,
		},
	},
	"afrog": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afrog",
			Port:  1042,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afrog",
			Port:  1042,
		},
	},
	"afs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs",
			Port:  1483,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs",
			Port:  1483,
		},
	},
	"afs3-bos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-bos",
			Port:  7007,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-bos",
			Port:  7007,
		},
	},
	"afs3-callback": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-callback",
			Port:  7001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-callback",
			Port:  7001,
		},
	},
	"afs3-errors": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-errors",
			Port:  7006,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-errors",
			Port:  7006,
		},
	},
	"afs3-fileserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-fileserver",
			Port:  7000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-fileserver",
			Port:  7000,
		},
	},
	"afs3-kaserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-kaserver",
			Port:  7004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-kaserver",
			Port:  7004,
		},
	},
	"afs3-prserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-prserver",
			Port:  7002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-prserver",
			Port:  7002,
		},
	},
	"afs3-rmtsys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-rmtsys",
			Port:  7009,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-rmtsys",
			Port:  7009,
		},
	},
	"afs3-update": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-update",
			Port:  7008,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-update",
			Port:  7008,
		},
	},
	"afs3-vlserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-vlserver",
			Port:  7003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-vlserver",
			Port:  7003,
		},
	},
	"afs3-volser": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "afs3-volser",
			Port:  7005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "afs3-volser",
			Port:  7005,
		},
	},
	"aftmux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aftmux",
			Port:  3917,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aftmux",
			Port:  3917,
		},
	},
	"agcat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agcat",
			Port:  3915,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agcat",
			Port:  3915,
		},
	},
	"agentsease-db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agentsease-db",
			Port:  3997,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agentsease-db",
			Port:  3997,
		},
	},
	"agentview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agentview",
			Port:  2331,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agentview",
			Port:  2331,
		},
	},
	"agentx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agentx",
			Port:  705,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agentx",
			Port:  705,
		},
	},
	"agpolicy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agpolicy",
			Port:  38203,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agpolicy",
			Port:  38203,
		},
	},
	"agps-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agps-port",
			Port:  3425,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agps-port",
			Port:  3425,
		},
	},
	"agri-gateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agri-gateway",
			Port:  3026,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agri-gateway",
			Port:  3026,
		},
	},
	"agriserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agriserver",
			Port:  3021,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agriserver",
			Port:  3021,
		},
	},
	"agslb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "agslb",
			Port:  4149,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "agslb",
			Port:  4149,
		},
	},
	"ah-esp-encap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ah-esp-encap",
			Port:  2070,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ah-esp-encap",
			Port:  2070,
		},
	},
	"aiagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aiagent",
			Port:  7738,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aiagent",
			Port:  7738,
		},
	},
	"aibkup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aibkup",
			Port:  4071,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aibkup",
			Port:  4071,
		},
	},
	"aic-np": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aic-np",
			Port:  2785,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aic-np",
			Port:  2785,
		},
	},
	"aic-oncrpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aic-oncrpc",
			Port:  2786,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aic-oncrpc",
			Port:  2786,
		},
	},
	"aicc-cmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aicc-cmi",
			Port:  3316,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aicc-cmi",
			Port:  3316,
		},
	},
	"aimpp-hello": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aimpp-hello",
			Port:  2846,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aimpp-hello",
			Port:  2846,
		},
	},
	"aimpp-port-req": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aimpp-port-req",
			Port:  2847,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aimpp-port-req",
			Port:  2847,
		},
	},
	"aipn-auth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aipn-auth",
			Port:  3833,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aipn-auth",
			Port:  3833,
		},
	},
	"aipn-reg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aipn-reg",
			Port:  4113,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aipn-reg",
			Port:  4113,
		},
	},
	"aironetddp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aironetddp",
			Port:  2887,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aironetddp",
			Port:  2887,
		},
	},
	"airs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "airs",
			Port:  1481,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "airs",
			Port:  1481,
		},
	},
	"airshot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "airshot",
			Port:  3975,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "airshot",
			Port:  3975,
		},
	},
	"airsync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "airsync",
			Port:  2175,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "airsync",
			Port:  2175,
		},
	},
	"aises": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aises",
			Port:  2783,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aises",
			Port:  2783,
		},
	},
	"aja-ntv4-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "aja-ntv4-disc",
			Port:  4804,
		},
	},
	"aker-cdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aker-cdp",
			Port:  2473,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aker-cdp",
			Port:  2473,
		},
	},
	"alaris-disc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alaris-disc",
			Port:  3613,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alaris-disc",
			Port:  3613,
		},
	},
	"alarm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alarm",
			Port:  2740,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alarm",
			Port:  2740,
		},
	},
	"alarm-clock-c": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alarm-clock-c",
			Port:  2668,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alarm-clock-c",
			Port:  2668,
		},
	},
	"alarm-clock-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alarm-clock-s",
			Port:  2667,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alarm-clock-s",
			Port:  2667,
		},
	},
	"alchemy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alchemy",
			Port:  3234,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alchemy",
			Port:  3234,
		},
	},
	"alesquery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alesquery",
			Port:  5074,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alesquery",
			Port:  5074,
		},
	},
	"alfin": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "alfin",
			Port:  16003,
		},
	},
	"alias": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alias",
			Port:  1187,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alias",
			Port:  1187,
		},
	},
	"alljoyn": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "alljoyn",
			Port:  9956,
		},
	},
	"alljoyn-mcm": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "alljoyn-mcm",
			Port:  9955,
		},
	},
	"alljoyn-stm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alljoyn-stm",
			Port:  9955,
		},
	},
	"allpeers": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "allpeers",
			Port:  36001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "allpeers",
			Port:  36001,
		},
	},
	"allstorcns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "allstorcns",
			Port:  2901,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "allstorcns",
			Port:  2901,
		},
	},
	"almobile-system": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "almobile-system",
			Port:  9209,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "almobile-system",
			Port:  9209,
		},
	},
	"alpes": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alpes",
			Port:  463,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alpes",
			Port:  463,
		},
	},
	"alpha-sms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alpha-sms",
			Port:  1849,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alpha-sms",
			Port:  1849,
		},
	},
	"alphatech-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alphatech-lm",
			Port:  1653,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alphatech-lm",
			Port:  1653,
		},
	},
	"alta-ana-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "alta-ana-lm",
			Port:  1346,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "alta-ana-lm",
			Port:  1346,
		},
	},
	"altalink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "altalink",
			Port:  1845,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "altalink",
			Port:  1845,
		},
	},
	"altav-remmgt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "altav-remmgt",
			Port:  2456,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "altav-remmgt",
			Port:  2456,
		},
	},
	"altav-tunnel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "altav-tunnel",
			Port:  3265,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "altav-tunnel",
			Port:  3265,
		},
	},
	"altbsdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "altbsdp",
			Port:  7799,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "altbsdp",
			Port:  7799,
		},
	},
	"altcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "altcp",
			Port:  4165,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "altcp",
			Port:  4165,
		},
	},
	"altova-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "altova-lm",
			Port:  35355,
		},
	},
	"altova-lm-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "altova-lm-disc",
			Port:  35355,
		},
	},
	"altovacentral": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "altovacentral",
			Port:  4689,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "altovacentral",
			Port:  4689,
		},
	},
	"altserviceboot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "altserviceboot",
			Port:  4011,
		},
	},
	"amanda": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amanda",
			Port:  10080,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amanda",
			Port:  10080,
		},
	},
	"amandaidx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amandaidx",
			Port:  10082,
		},
	},
	"amberon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amberon",
			Port:  8301,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amberon",
			Port:  8301,
		},
	},
	"ambit-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ambit-lm",
			Port:  6831,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ambit-lm",
			Port:  6831,
		},
	},
	"amc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amc",
			Port:  5506,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amc",
			Port:  5506,
		},
	},
	"amcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amcs",
			Port:  8766,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amcs",
			Port:  8766,
		},
	},
	"amdsched": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amdsched",
			Port:  1931,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amdsched",
			Port:  1931,
		},
	},
	"amidxtape": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amidxtape",
			Port:  10083,
		},
	},
	"amiganetfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amiganetfs",
			Port:  2100,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amiganetfs",
			Port:  2100,
		},
	},
	"aminet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aminet",
			Port:  2639,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aminet",
			Port:  2639,
		},
	},
	"amp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amp",
			Port:  3811,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amp",
			Port:  3811,
		},
	},
	"ampify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ampify",
			Port:  8040,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ampify",
			Port:  8040,
		},
	},
	"ampl-lic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ampl-lic",
			Port:  5195,
		},
	},
	"ampl-tableproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ampl-tableproxy",
			Port:  5196,
		},
	},
	"ampr-info": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ampr-info",
			Port:  1535,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ampr-info",
			Port:  1535,
		},
	},
	"ampr-inter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ampr-inter",
			Port:  1536,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ampr-inter",
			Port:  1536,
		},
	},
	"ampr-rcmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ampr-rcmd",
			Port:  459,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ampr-rcmd",
			Port:  459,
		},
	},
	"amqp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "amqp",
			Port:  5672,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "amqp",
			Port:  5672,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amqp",
			Port:  5672,
		},
	},
	"amqps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amqps",
			Port:  5671,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amqps",
			Port:  5671,
		},
	},
	"ams": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ams",
			Port:  1037,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ams",
			Port:  1037,
		},
	},
	"amt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amt",
			Port:  2268,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amt",
			Port:  2268,
		},
	},
	"amt-blc-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amt-blc-port",
			Port:  2848,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amt-blc-port",
			Port:  2848,
		},
	},
	"amt-cnf-prot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amt-cnf-prot",
			Port:  3054,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amt-cnf-prot",
			Port:  3054,
		},
	},
	"amt-esd-prot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amt-esd-prot",
			Port:  1082,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amt-esd-prot",
			Port:  1082,
		},
	},
	"amt-redir-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amt-redir-tcp",
			Port:  16994,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amt-redir-tcp",
			Port:  16994,
		},
	},
	"amt-redir-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amt-redir-tls",
			Port:  16995,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amt-redir-tls",
			Port:  16995,
		},
	},
	"amt-soap-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amt-soap-http",
			Port:  16992,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amt-soap-http",
			Port:  16992,
		},
	},
	"amt-soap-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amt-soap-https",
			Port:  16993,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amt-soap-https",
			Port:  16993,
		},
	},
	"amx-axbnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amx-axbnet",
			Port:  1320,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amx-axbnet",
			Port:  1320,
		},
	},
	"amx-icsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amx-icsp",
			Port:  1319,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amx-icsp",
			Port:  1319,
		},
	},
	"amx-rms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amx-rms",
			Port:  3839,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amx-rms",
			Port:  3839,
		},
	},
	"amx-webadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amx-webadmin",
			Port:  2929,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amx-webadmin",
			Port:  2929,
		},
	},
	"amx-weblinx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "amx-weblinx",
			Port:  2930,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "amx-weblinx",
			Port:  2930,
		},
	},
	"an-pcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "an-pcp",
			Port:  3846,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "an-pcp",
			Port:  3846,
		},
	},
	"and-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "and-lm",
			Port:  2646,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "and-lm",
			Port:  2646,
		},
	},
	"anet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "anet",
			Port:  212,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "anet",
			Port:  212,
		},
	},
	"anet-b": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "anet-b",
			Port:  3338,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "anet-b",
			Port:  3338,
		},
	},
	"anet-h": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "anet-h",
			Port:  3341,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "anet-h",
			Port:  3341,
		},
	},
	"anet-l": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "anet-l",
			Port:  3339,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "anet-l",
			Port:  3339,
		},
	},
	"anet-m": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "anet-m",
			Port:  3340,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "anet-m",
			Port:  3340,
		},
	},
	"anoto-rendezv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "anoto-rendezv",
			Port:  3715,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "anoto-rendezv",
			Port:  3715,
		},
	},
	"ans-console": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ans-console",
			Port:  3440,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ans-console",
			Port:  3440,
		},
	},
	"ansanotify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ansanotify",
			Port:  116,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ansanotify",
			Port:  116,
		},
	},
	"ansatrader": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ansatrader",
			Port:  124,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ansatrader",
			Port:  124,
		},
	},
	"ansoft-lm-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ansoft-lm-1",
			Port:  1083,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ansoft-lm-1",
			Port:  1083,
		},
	},
	"ansoft-lm-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ansoft-lm-2",
			Port:  1084,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ansoft-lm-2",
			Port:  1084,
		},
	},
	"answersoft-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "answersoft-lm",
			Port:  1781,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "answersoft-lm",
			Port:  1781,
		},
	},
	"ansys-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ansys-lm",
			Port:  1800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ansys-lm",
			Port:  1800,
		},
	},
	"ansysli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ansysli",
			Port:  2325,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ansysli",
			Port:  2325,
		},
	},
	"ansyslmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ansyslmd",
			Port:  1055,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ansyslmd",
			Port:  1055,
		},
	},
	"anthony-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "anthony-data",
			Port:  1206,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "anthony-data",
			Port:  1206,
		},
	},
	"antidotemgrsvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "antidotemgrsvr",
			Port:  2247,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "antidotemgrsvr",
			Port:  2247,
		},
	},
	"aocp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aocp",
			Port:  2712,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aocp",
			Port:  2712,
		},
	},
	"aodv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aodv",
			Port:  654,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aodv",
			Port:  654,
		},
	},
	"aol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aol",
			Port:  5190,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aol",
			Port:  5190,
		},
	},
	"aol-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aol-1",
			Port:  5191,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aol-1",
			Port:  5191,
		},
	},
	"aol-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aol-2",
			Port:  5192,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aol-2",
			Port:  5192,
		},
	},
	"aol-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aol-3",
			Port:  5193,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aol-3",
			Port:  5193,
		},
	},
	"ap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ap",
			Port:  47806,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ap",
			Port:  47806,
		},
	},
	"apani1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apani1",
			Port:  9160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apani1",
			Port:  9160,
		},
	},
	"apani2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apani2",
			Port:  9161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apani2",
			Port:  9161,
		},
	},
	"apani3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apani3",
			Port:  9162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apani3",
			Port:  9162,
		},
	},
	"apani4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apani4",
			Port:  9163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apani4",
			Port:  9163,
		},
	},
	"apani5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apani5",
			Port:  9164,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apani5",
			Port:  9164,
		},
	},
	"apc-2160": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-2160",
			Port:  2160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-2160",
			Port:  2160,
		},
	},
	"apc-2161": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-2161",
			Port:  2161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-2161",
			Port:  2161,
		},
	},
	"apc-2260": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-2260",
			Port:  2260,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-2260",
			Port:  2260,
		},
	},
	"apc-3052": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-3052",
			Port:  3052,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-3052",
			Port:  3052,
		},
	},
	"apc-3506": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-3506",
			Port:  3506,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-3506",
			Port:  3506,
		},
	},
	"apc-5454": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-5454",
			Port:  5454,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-5454",
			Port:  5454,
		},
	},
	"apc-5455": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-5455",
			Port:  5455,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-5455",
			Port:  5455,
		},
	},
	"apc-5456": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-5456",
			Port:  5456,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-5456",
			Port:  5456,
		},
	},
	"apc-6547": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-6547",
			Port:  6547,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-6547",
			Port:  6547,
		},
	},
	"apc-6548": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-6548",
			Port:  6548,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-6548",
			Port:  6548,
		},
	},
	"apc-6549": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-6549",
			Port:  6549,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-6549",
			Port:  6549,
		},
	},
	"apc-7845": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-7845",
			Port:  7845,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-7845",
			Port:  7845,
		},
	},
	"apc-7846": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-7846",
			Port:  7846,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-7846",
			Port:  7846,
		},
	},
	"apc-9950": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-9950",
			Port:  9950,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-9950",
			Port:  9950,
		},
	},
	"apc-9951": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-9951",
			Port:  9951,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-9951",
			Port:  9951,
		},
	},
	"apc-9952": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-9952",
			Port:  9952,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-9952",
			Port:  9952,
		},
	},
	"apc-necmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apc-necmp",
			Port:  18888,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apc-necmp",
			Port:  18888,
		},
	},
	"apcupsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apcupsd",
			Port:  3551,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apcupsd",
			Port:  3551,
		},
	},
	"apdap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apdap",
			Port:  3948,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apdap",
			Port:  3948,
		},
	},
	"apertus-ldp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apertus-ldp",
			Port:  539,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apertus-ldp",
			Port:  539,
		},
	},
	"apex-edge": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apex-edge",
			Port:  913,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apex-edge",
			Port:  913,
		},
	},
	"apex-mesh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apex-mesh",
			Port:  912,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apex-mesh",
			Port:  912,
		},
	},
	"aplx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aplx",
			Port:  1134,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aplx",
			Port:  1134,
		},
	},
	"apm-link": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apm-link",
			Port:  32483,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apm-link",
			Port:  32483,
		},
	},
	"apocd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apocd",
			Port:  3809,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apocd",
			Port:  3809,
		},
	},
	"apogeex-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apogeex-port",
			Port:  3184,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apogeex-port",
			Port:  3184,
		},
	},
	"apollo-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apollo-admin",
			Port:  8122,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apollo-admin",
			Port:  8122,
		},
	},
	"apollo-cc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apollo-cc",
			Port:  2754,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apollo-cc",
			Port:  2754,
		},
	},
	"apollo-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apollo-data",
			Port:  8121,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apollo-data",
			Port:  8121,
		},
	},
	"apollo-gms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apollo-gms",
			Port:  2759,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apollo-gms",
			Port:  2759,
		},
	},
	"apollo-relay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apollo-relay",
			Port:  10252,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apollo-relay",
			Port:  10252,
		},
	},
	"apollo-status": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apollo-status",
			Port:  2758,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apollo-status",
			Port:  2758,
		},
	},
	"apparenet-as": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apparenet-as",
			Port:  3238,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apparenet-as",
			Port:  3238,
		},
	},
	"apparenet-tps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apparenet-tps",
			Port:  3237,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apparenet-tps",
			Port:  3237,
		},
	},
	"apparenet-ts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apparenet-ts",
			Port:  3236,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apparenet-ts",
			Port:  3236,
		},
	},
	"apparenet-ui": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apparenet-ui",
			Port:  3239,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apparenet-ui",
			Port:  3239,
		},
	},
	"appiq-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appiq-mgmt",
			Port:  4674,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appiq-mgmt",
			Port:  4674,
		},
	},
	"apple-licman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apple-licman",
			Port:  1381,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apple-licman",
			Port:  1381,
		},
	},
	"apple-sasl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apple-sasl",
			Port:  3659,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apple-sasl",
			Port:  3659,
		},
	},
	"apple-vpns-rp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apple-vpns-rp",
			Port:  4112,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apple-vpns-rp",
			Port:  4112,
		},
	},
	"appleqtc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appleqtc",
			Port:  458,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appleqtc",
			Port:  458,
		},
	},
	"appleqtcsrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appleqtcsrvr",
			Port:  545,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appleqtcsrvr",
			Port:  545,
		},
	},
	"appleugcontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appleugcontrol",
			Port:  2336,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appleugcontrol",
			Port:  2336,
		},
	},
	"appliance-cfg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appliance-cfg",
			Port:  2898,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appliance-cfg",
			Port:  2898,
		},
	},
	"applix": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "applix",
			Port:  999,
		},
	},
	"applus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "applus",
			Port:  2037,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "applus",
			Port:  2037,
		},
	},
	"applusservice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "applusservice",
			Port:  4087,
		},
	},
	"appman-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appman-server",
			Port:  3312,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appman-server",
			Port:  3312,
		},
	},
	"appserv-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appserv-http",
			Port:  4848,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appserv-http",
			Port:  4848,
		},
	},
	"appserv-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appserv-https",
			Port:  4849,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appserv-https",
			Port:  4849,
		},
	},
	"appss-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appss-lm",
			Port:  3879,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appss-lm",
			Port:  3879,
		},
	},
	"appswitch-emp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appswitch-emp",
			Port:  2616,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appswitch-emp",
			Port:  2616,
		},
	},
	"appworxsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "appworxsrv",
			Port:  2136,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "appworxsrv",
			Port:  2136,
		},
	},
	"apri-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apri-lm",
			Port:  1447,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apri-lm",
			Port:  1447,
		},
	},
	"aprigo-cs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aprigo-cs",
			Port:  5086,
		},
	},
	"apw-registry": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apw-registry",
			Port:  3758,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apw-registry",
			Port:  3758,
		},
	},
	"apwi-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "apwi-disc",
			Port:  4394,
		},
	},
	"apwi-imserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apwi-imserver",
			Port:  4391,
		},
	},
	"apwi-rxserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apwi-rxserver",
			Port:  4392,
		},
	},
	"apwi-rxspooler": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apwi-rxspooler",
			Port:  4393,
		},
	},
	"apx500api-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apx500api-1",
			Port:  2264,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apx500api-1",
			Port:  2264,
		},
	},
	"apx500api-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "apx500api-2",
			Port:  2265,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "apx500api-2",
			Port:  2265,
		},
	},
	"arbortext-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arbortext-lm",
			Port:  1557,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arbortext-lm",
			Port:  1557,
		},
	},
	"arcisdms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arcisdms",
			Port:  262,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arcisdms",
			Port:  262,
		},
	},
	"arcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arcp",
			Port:  7070,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arcp",
			Port:  7070,
		},
	},
	"arcpd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arcpd",
			Port:  3513,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arcpd",
			Port:  3513,
		},
	},
	"ardt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ardt",
			Port:  1826,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ardt",
			Port:  1826,
		},
	},
	"ardus-cntl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ardus-cntl",
			Port:  1116,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ardus-cntl",
			Port:  1116,
		},
	},
	"ardus-mtrns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ardus-mtrns",
			Port:  1117,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ardus-mtrns",
			Port:  1117,
		},
	},
	"ardus-trns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ardus-trns",
			Port:  1115,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ardus-trns",
			Port:  1115,
		},
	},
	"ardusmul": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ardusmul",
			Port:  1835,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ardusmul",
			Port:  1835,
		},
	},
	"ardusuni": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ardusuni",
			Port:  1834,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ardusuni",
			Port:  1834,
		},
	},
	"areaguard-neo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "areaguard-neo",
			Port:  23546,
		},
	},
	"arena-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arena-server",
			Port:  11321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arena-server",
			Port:  11321,
		},
	},
	"arepa-cas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arepa-cas",
			Port:  3030,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arepa-cas",
			Port:  3030,
		},
	},
	"arepa-raft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arepa-raft",
			Port:  3025,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arepa-raft",
			Port:  3025,
		},
	},
	"argis-ds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "argis-ds",
			Port:  2582,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "argis-ds",
			Port:  2582,
		},
	},
	"argis-te": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "argis-te",
			Port:  2581,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "argis-te",
			Port:  2581,
		},
	},
	"aria": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aria",
			Port:  2624,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aria",
			Port:  2624,
		},
	},
	"ariel1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ariel1",
			Port:  419,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ariel1",
			Port:  419,
		},
	},
	"ariel2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ariel2",
			Port:  421,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ariel2",
			Port:  421,
		},
	},
	"ariel3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ariel3",
			Port:  422,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ariel3",
			Port:  422,
		},
	},
	"aries-kfinder": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aries-kfinder",
			Port:  7570,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aries-kfinder",
			Port:  7570,
		},
	},
	"ariliamulti": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ariliamulti",
			Port:  3140,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ariliamulti",
			Port:  3140,
		},
	},
	"arkivio": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arkivio",
			Port:  3426,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arkivio",
			Port:  3426,
		},
	},
	"armadp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "armadp",
			Port:  1913,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "armadp",
			Port:  1913,
		},
	},
	"armagetronad": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "armagetronad",
			Port:  4534,
		},
	},
	"armcenterhttp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "armcenterhttp",
			Port:  9294,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "armcenterhttp",
			Port:  9294,
		},
	},
	"armcenterhttps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "armcenterhttps",
			Port:  9295,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "armcenterhttps",
			Port:  9295,
		},
	},
	"armi-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "armi-server",
			Port:  3174,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "armi-server",
			Port:  3174,
		},
	},
	"armtechdaemon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "armtechdaemon",
			Port:  9292,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "armtechdaemon",
			Port:  9292,
		},
	},
	"arns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "arns",
			Port:  384,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "arns",
			Port:  384,
		},
	},
	"array-manager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "array-manager",
			Port:  3726,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "array-manager",
			Port:  3726,
		},
	},
	"ars-master": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ars-master",
			Port:  3176,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ars-master",
			Port:  3176,
		},
	},
	"ars-vista": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ars-vista",
			Port:  27782,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ars-vista",
			Port:  27782,
		},
	},
	"artifact-msg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "artifact-msg",
			Port:  3518,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "artifact-msg",
			Port:  3518,
		},
	},
	"aruba-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aruba-server",
			Port:  7166,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aruba-server",
			Port:  7166,
		},
	},
	"as-debug": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "as-debug",
			Port:  4026,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "as-debug",
			Port:  4026,
		},
	},
	"as-servermap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "as-servermap",
			Port:  449,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "as-servermap",
			Port:  449,
		},
	},
	"asa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asa",
			Port:  386,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asa",
			Port:  386,
		},
	},
	"asa-appl-proto": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asa-appl-proto",
			Port:  502,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asa-appl-proto",
			Port:  502,
		},
	},
	"asam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asam",
			Port:  3451,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asam",
			Port:  3451,
		},
	},
	"asap-sctp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "asap-sctp",
			Port:  3863,
		},
	},
	"asap-sctp-tls": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "asap-sctp-tls",
			Port:  3864,
		},
	},
	"asap-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asap-tcp",
			Port:  3863,
		},
	},
	"asap-tcp-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asap-tcp-tls",
			Port:  3864,
		},
	},
	"asap-udp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "asap-udp",
			Port:  3863,
		},
	},
	"asc-slmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asc-slmd",
			Port:  4448,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asc-slmd",
			Port:  4448,
		},
	},
	"asci-val": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asci-val",
			Port:  1560,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asci-val",
			Port:  1560,
		},
	},
	"ascomalarm": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ascomalarm",
			Port:  4077,
		},
	},
	"asctrl-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asctrl-agent",
			Port:  5155,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asctrl-agent",
			Port:  5155,
		},
	},
	"asdis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asdis",
			Port:  2192,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asdis",
			Port:  2192,
		},
	},
	"asf-rmcp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "asf-rmcp",
			Port:  623,
		},
	},
	"asf-secure-rmcp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "asf-secure-rmcp",
			Port:  664,
		},
	},
	"asgcypresstcps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asgcypresstcps",
			Port:  11489,
		},
	},
	"asgenf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asgenf",
			Port:  5727,
		},
	},
	"asi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asi",
			Port:  1827,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asi",
			Port:  1827,
		},
	},
	"asia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asia",
			Port:  626,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asia",
			Port:  626,
		},
	},
	"asip-webadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asip-webadmin",
			Port:  311,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asip-webadmin",
			Port:  311,
		},
	},
	"asipregistry": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asipregistry",
			Port:  687,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asipregistry",
			Port:  687,
		},
	},
	"asmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asmp",
			Port:  45000,
		},
	},
	"asmp-mon": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "asmp-mon",
			Port:  45000,
		},
	},
	"asmps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asmps",
			Port:  45001,
		},
	},
	"asnaacceler8db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asnaacceler8db",
			Port:  5042,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asnaacceler8db",
			Port:  5042,
		},
	},
	"asoki-sma": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asoki-sma",
			Port:  3087,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asoki-sma",
			Port:  3087,
		},
	},
	"asp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asp",
			Port:  27374,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asp",
			Port:  27374,
		},
	},
	"aspeclmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aspeclmd",
			Port:  1544,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aspeclmd",
			Port:  1544,
		},
	},
	"aspen-services": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aspen-services",
			Port:  1749,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aspen-services",
			Port:  1749,
		},
	},
	"aspentec-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aspentec-lm",
			Port:  6142,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aspentec-lm",
			Port:  6142,
		},
	},
	"asprovatalk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asprovatalk",
			Port:  1079,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asprovatalk",
			Port:  1079,
		},
	},
	"asr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asr",
			Port:  7800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asr",
			Port:  7800,
		},
	},
	"assoc-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "assoc-disc",
			Port:  24850,
		},
	},
	"assuria-ins": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "assuria-ins",
			Port:  4704,
		},
	},
	"assuria-slm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "assuria-slm",
			Port:  4119,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "assuria-slm",
			Port:  4119,
		},
	},
	"assyst-dr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "assyst-dr",
			Port:  4485,
		},
	},
	"astergate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "astergate",
			Port:  9106,
		},
	},
	"astergate-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "astergate-disc",
			Port:  9106,
		},
	},
	"astergatefax": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "astergatefax",
			Port:  9107,
		},
	},
	"asterix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "asterix",
			Port:  8600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "asterix",
			Port:  8600,
		},
	},
	"astrolink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "astrolink",
			Port:  27876,
		},
	},
	"astromed-main": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "astromed-main",
			Port:  2864,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "astromed-main",
			Port:  2864,
		},
	},
	"at-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "at-3",
			Port:  203,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "at-3",
			Port:  203,
		},
	},
	"at-5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "at-5",
			Port:  205,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "at-5",
			Port:  205,
		},
	},
	"at-7": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "at-7",
			Port:  207,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "at-7",
			Port:  207,
		},
	},
	"at-8": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "at-8",
			Port:  208,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "at-8",
			Port:  208,
		},
	},
	"at-echo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "at-echo",
			Port:  204,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "at-echo",
			Port:  204,
		},
	},
	"at-nbp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "at-nbp",
			Port:  202,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "at-nbp",
			Port:  202,
		},
	},
	"at-rtmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "at-rtmp",
			Port:  201,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "at-rtmp",
			Port:  201,
		},
	},
	"at-zis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "at-zis",
			Port:  206,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "at-zis",
			Port:  206,
		},
	},
	"atc-appserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atc-appserver",
			Port:  1171,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atc-appserver",
			Port:  1171,
		},
	},
	"atc-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atc-lm",
			Port:  1170,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atc-lm",
			Port:  1170,
		},
	},
	"atex_elmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atex_elmd",
			Port:  1385,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atex_elmd",
			Port:  1385,
		},
	},
	"athand-mmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "athand-mmp",
			Port:  20999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "athand-mmp",
			Port:  20999,
		},
	},
	"ati-ip-to-ncpe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ati-ip-to-ncpe",
			Port:  3965,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ati-ip-to-ncpe",
			Port:  3965,
		},
	},
	"atlinks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atlinks",
			Port:  4154,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atlinks",
			Port:  4154,
		},
	},
	"atm-uhas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atm-uhas",
			Port:  11367,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atm-uhas",
			Port:  11367,
		},
	},
	"atm-zip-office": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atm-zip-office",
			Port:  1520,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atm-zip-office",
			Port:  1520,
		},
	},
	"atmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atmp",
			Port:  5150,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atmp",
			Port:  5150,
		},
	},
	"atmtcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atmtcp",
			Port:  2812,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atmtcp",
			Port:  2812,
		},
	},
	"ats": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ats",
			Port:  2201,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ats",
			Port:  2201,
		},
	},
	"atsc-mh-ssc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "atsc-mh-ssc",
			Port:  4937,
		},
	},
	"attachmate-g32": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "attachmate-g32",
			Port:  2317,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "attachmate-g32",
			Port:  2317,
		},
	},
	"attachmate-s2s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "attachmate-s2s",
			Port:  2419,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "attachmate-s2s",
			Port:  2419,
		},
	},
	"attachmate-uts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "attachmate-uts",
			Port:  2304,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "attachmate-uts",
			Port:  2304,
		},
	},
	"atul": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "atul",
			Port:  7543,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "atul",
			Port:  7543,
		},
	},
	"audio-activmail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "audio-activmail",
			Port:  1397,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "audio-activmail",
			Port:  1397,
		},
	},
	"audiojuggler": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "audiojuggler",
			Port:  3643,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "audiojuggler",
			Port:  3643,
		},
	},
	"audit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "audit",
			Port:  182,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "audit",
			Port:  182,
		},
	},
	"audit-transfer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "audit-transfer",
			Port:  1146,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "audit-transfer",
			Port:  1146,
		},
	},
	"auditd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "auditd",
			Port:  48,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "auditd",
			Port:  48,
		},
	},
	"aura": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aura",
			Port:  2066,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aura",
			Port:  2066,
		},
	},
	"auriga-router": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "auriga-router",
			Port:  5680,
		},
	},
	"auris": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "auris",
			Port:  2772,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "auris",
			Port:  2772,
		},
	},
	"aurora": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "aurora",
			Port:  9084,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "aurora",
			Port:  9084,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aurora",
			Port:  9084,
		},
	},
	"aurora-balaena": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aurora-balaena",
			Port:  33123,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aurora-balaena",
			Port:  33123,
		},
	},
	"aurora-cmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aurora-cmgr",
			Port:  364,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aurora-cmgr",
			Port:  364,
		},
	},
	"aurp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aurp",
			Port:  387,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aurp",
			Port:  387,
		},
	},
	"auth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "auth",
			Port:  113,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "auth",
			Port:  113,
		},
	},
	"authentx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "authentx",
			Port:  5067,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "authentx",
			Port:  5067,
		},
	},
	"autobuild": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autobuild",
			Port:  5115,
		},
	},
	"autocueds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autocueds",
			Port:  3437,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "autocueds",
			Port:  3437,
		},
	},
	"autocuelog": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autocuelog",
			Port:  3104,
		},
	},
	"autocuesmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autocuesmi",
			Port:  3103,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "autocuesmi",
			Port:  3103,
		},
	},
	"autocuetime": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "autocuetime",
			Port:  3104,
		},
	},
	"autodesk-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autodesk-lm",
			Port:  1422,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "autodesk-lm",
			Port:  1422,
		},
	},
	"autodesk-nlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autodesk-nlm",
			Port:  2080,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "autodesk-nlm",
			Port:  2080,
		},
	},
	"autonoc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autonoc",
			Port:  1140,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "autonoc",
			Port:  1140,
		},
	},
	"autopac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autopac",
			Port:  4685,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "autopac",
			Port:  4685,
		},
	},
	"autotrac-acp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "autotrac-acp",
			Port:  31020,
		},
	},
	"av-emb-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "av-emb-config",
			Port:  2050,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "av-emb-config",
			Port:  2050,
		},
	},
	"availant-mgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "availant-mgr",
			Port:  1122,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "availant-mgr",
			Port:  1122,
		},
	},
	"avantageb2b": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avantageb2b",
			Port:  2131,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avantageb2b",
			Port:  2131,
		},
	},
	"avanti_cdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avanti_cdp",
			Port:  4065,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avanti_cdp",
			Port:  4065,
		},
	},
	"avauthsrvprtcl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avauthsrvprtcl",
			Port:  2068,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avauthsrvprtcl",
			Port:  2068,
		},
	},
	"avdecc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avdecc",
			Port:  17221,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avdecc",
			Port:  17221,
		},
	},
	"avenue": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avenue",
			Port:  2134,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avenue",
			Port:  2134,
		},
	},
	"avenyo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avenyo",
			Port:  2992,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avenyo",
			Port:  2992,
		},
	},
	"avian": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avian",
			Port:  486,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avian",
			Port:  486,
		},
	},
	"avinstalldisc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avinstalldisc",
			Port:  3502,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avinstalldisc",
			Port:  3502,
		},
	},
	"aviva-sna": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aviva-sna",
			Port:  2238,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aviva-sna",
			Port:  2238,
		},
	},
	"avocent-adsap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avocent-adsap",
			Port:  3871,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avocent-adsap",
			Port:  3871,
		},
	},
	"avocent-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avocent-proxy",
			Port:  1078,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avocent-proxy",
			Port:  1078,
		},
	},
	"avsecuremgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "avsecuremgmt",
			Port:  3211,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avsecuremgmt",
			Port:  3211,
		},
	},
	"avt-profile-1": map[string]Service{
		"dccp": Service{
			Proto: "dccp",
			Name:  "avt-profile-1",
			Port:  5004,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "avt-profile-1",
			Port:  5004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avt-profile-1",
			Port:  5004,
		},
	},
	"avt-profile-2": map[string]Service{
		"dccp": Service{
			Proto: "dccp",
			Name:  "avt-profile-2",
			Port:  5005,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "avt-profile-2",
			Port:  5005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "avt-profile-2",
			Port:  5005,
		},
	},
	"awacs-ice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "awacs-ice",
			Port:  4488,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "awacs-ice",
			Port:  4488,
		},
	},
	"awg-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "awg-proxy",
			Port:  3277,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "awg-proxy",
			Port:  3277,
		},
	},
	"aws-brf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aws-brf",
			Port:  22800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aws-brf",
			Port:  22800,
		},
	},
	"axis-wimp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "axis-wimp-port",
			Port:  10260,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "axis-wimp-port",
			Port:  10260,
		},
	},
	"axon-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "axon-lm",
			Port:  1548,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "axon-lm",
			Port:  1548,
		},
	},
	"ayiya": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ayiya",
			Port:  5072,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ayiya",
			Port:  5072,
		},
	},
	"azeti": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "azeti",
			Port:  4192,
		},
	},
	"azeti-bd": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "azeti-bd",
			Port:  4192,
		},
	},
	"aztec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "aztec",
			Port:  3512,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "aztec",
			Port:  3512,
		},
	},
	"b-novative-ls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "b-novative-ls",
			Port:  1896,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "b-novative-ls",
			Port:  1896,
		},
	},
	"b2-license": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "b2-license",
			Port:  2204,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "b2-license",
			Port:  2204,
		},
	},
	"b2-runtime": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "b2-runtime",
			Port:  2203,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "b2-runtime",
			Port:  2203,
		},
	},
	"b2n": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "b2n",
			Port:  1179,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "b2n",
			Port:  1179,
		},
	},
	"babel": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "babel",
			Port:  6696,
		},
	},
	"backburner": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "backburner",
			Port:  2635,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "backburner",
			Port:  2635,
		},
	},
	"backroomnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "backroomnet",
			Port:  3387,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "backroomnet",
			Port:  3387,
		},
	},
	"backup-express": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "backup-express",
			Port:  6123,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "backup-express",
			Port:  6123,
		},
	},
	"backupedge": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "backupedge",
			Port:  3946,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "backupedge",
			Port:  3946,
		},
	},
	"bacnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bacnet",
			Port:  47808,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bacnet",
			Port:  47808,
		},
	},
	"bacula-dir": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bacula-dir",
			Port:  9101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bacula-dir",
			Port:  9101,
		},
	},
	"bacula-fd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bacula-fd",
			Port:  9102,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bacula-fd",
			Port:  9102,
		},
	},
	"bacula-sd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bacula-sd",
			Port:  9103,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bacula-sd",
			Port:  9103,
		},
	},
	"badm_priv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "badm_priv",
			Port:  6505,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "badm_priv",
			Port:  6505,
		},
	},
	"badm_pub": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "badm_pub",
			Port:  6506,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "badm_pub",
			Port:  6506,
		},
	},
	"balour": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "balour",
			Port:  4324,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "balour",
			Port:  4324,
		},
	},
	"bandwiz-system": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bandwiz-system",
			Port:  1929,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bandwiz-system",
			Port:  1929,
		},
	},
	"banyan-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "banyan-net",
			Port:  2708,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "banyan-net",
			Port:  2708,
		},
	},
	"banyan-rpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "banyan-rpc",
			Port:  567,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "banyan-rpc",
			Port:  567,
		},
	},
	"banyan-vip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "banyan-vip",
			Port:  573,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "banyan-vip",
			Port:  573,
		},
	},
	"barracuda-bbs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "barracuda-bbs",
			Port:  5120,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "barracuda-bbs",
			Port:  5120,
		},
	},
	"base": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "base",
			Port:  5429,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "base",
			Port:  5429,
		},
	},
	"batman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "batman",
			Port:  4305,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "batman",
			Port:  4305,
		},
	},
	"bb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bb",
			Port:  1984,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bb",
			Port:  1984,
		},
	},
	"bbars": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bbars",
			Port:  3327,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bbars",
			Port:  3327,
		},
	},
	"bbn-mmc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bbn-mmc",
			Port:  1347,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bbn-mmc",
			Port:  1347,
		},
	},
	"bbn-mmx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bbn-mmx",
			Port:  1348,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bbn-mmx",
			Port:  1348,
		},
	},
	"bccp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bccp",
			Port:  4175,
		},
	},
	"bcinameservice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bcinameservice",
			Port:  3415,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bcinameservice",
			Port:  3415,
		},
	},
	"bcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bcs",
			Port:  4677,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bcs",
			Port:  4677,
		},
	},
	"bcs-broker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bcs-broker",
			Port:  1704,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bcs-broker",
			Port:  1704,
		},
	},
	"bcs-lmserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bcs-lmserver",
			Port:  1951,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bcs-lmserver",
			Port:  1951,
		},
	},
	"bcslogc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bcslogc",
			Port:  13216,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bcslogc",
			Port:  13216,
		},
	},
	"bctp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bctp",
			Port:  8999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bctp",
			Port:  8999,
		},
	},
	"bctp-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bctp-server",
			Port:  10107,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bctp-server",
			Port:  10107,
		},
	},
	"bdir_priv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bdir_priv",
			Port:  6507,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bdir_priv",
			Port:  6507,
		},
	},
	"bdir_pub": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bdir_pub",
			Port:  6508,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bdir_pub",
			Port:  6508,
		},
	},
	"bdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bdp",
			Port:  581,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bdp",
			Port:  581,
		},
	},
	"beacon-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "beacon-port",
			Port:  3124,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "beacon-port",
			Port:  3124,
		},
	},
	"beacon-port-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "beacon-port-2",
			Port:  4426,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "beacon-port-2",
			Port:  4426,
		},
	},
	"bears-01": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bears-01",
			Port:  2852,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bears-01",
			Port:  2852,
		},
	},
	"bears-02": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bears-02",
			Port:  3146,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bears-02",
			Port:  3146,
		},
	},
	"beeyond": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "beeyond",
			Port:  2414,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "beeyond",
			Port:  2414,
		},
	},
	"beeyond-media": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "beeyond-media",
			Port:  1943,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "beeyond-media",
			Port:  1943,
		},
	},
	"beorl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "beorl",
			Port:  5633,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "beorl",
			Port:  5633,
		},
	},
	"berknet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "berknet",
			Port:  2005,
		},
	},
	"beserver-msg-q": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "beserver-msg-q",
			Port:  3527,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "beserver-msg-q",
			Port:  3527,
		},
	},
	"bess": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bess",
			Port:  3960,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bess",
			Port:  3960,
		},
	},
	"bex-webadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bex-webadmin",
			Port:  6122,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bex-webadmin",
			Port:  6122,
		},
	},
	"bex-xr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bex-xr",
			Port:  15660,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bex-xr",
			Port:  15660,
		},
	},
	"beyond-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "beyond-remote",
			Port:  5424,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "beyond-remote",
			Port:  5424,
		},
	},
	"bf-game": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "bf-game",
			Port:  25954,
		},
	},
	"bf-master": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "bf-master",
			Port:  25955,
		},
	},
	"bfd-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bfd-control",
			Port:  3784,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bfd-control",
			Port:  3784,
		},
	},
	"bfd-echo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bfd-echo",
			Port:  3785,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bfd-echo",
			Port:  3785,
		},
	},
	"bfd-lag": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "bfd-lag",
			Port:  6784,
		},
	},
	"bfd-multi-ctl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bfd-multi-ctl",
			Port:  4784,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bfd-multi-ctl",
			Port:  4784,
		},
	},
	"bflckmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bflckmgr",
			Port:  3966,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bflckmgr",
			Port:  3966,
		},
	},
	"bftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bftp",
			Port:  152,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bftp",
			Port:  152,
		},
	},
	"bgmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bgmp",
			Port:  264,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bgmp",
			Port:  264,
		},
	},
	"bgp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "bgp",
			Port:  179,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "bgp",
			Port:  179,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bgp",
			Port:  179,
		},
	},
	"bgs-nsi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bgs-nsi",
			Port:  482,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bgs-nsi",
			Port:  482,
		},
	},
	"bh611": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bh611",
			Port:  354,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bh611",
			Port:  354,
		},
	},
	"bhevent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bhevent",
			Port:  357,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bhevent",
			Port:  357,
		},
	},
	"bhfhs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bhfhs",
			Port:  248,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bhfhs",
			Port:  248,
		},
	},
	"bhmds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bhmds",
			Port:  310,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bhmds",
			Port:  310,
		},
	},
	"biap-mp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "biap-mp",
			Port:  1962,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "biap-mp",
			Port:  1962,
		},
	},
	"biff": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "biff",
			Port:  512,
		},
	},
	"biimenu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "biimenu",
			Port:  18000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "biimenu",
			Port:  18000,
		},
	},
	"bim-pem": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bim-pem",
			Port:  3783,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bim-pem",
			Port:  3783,
		},
	},
	"binderysupport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "binderysupport",
			Port:  2302,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "binderysupport",
			Port:  2302,
		},
	},
	"bingbang": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bingbang",
			Port:  29999,
		},
	},
	"binkp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "binkp",
			Port:  24554,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "binkp",
			Port:  24554,
		},
	},
	"bintec-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bintec-admin",
			Port:  2107,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bintec-admin",
			Port:  2107,
		},
	},
	"bintec-capi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bintec-capi",
			Port:  2662,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bintec-capi",
			Port:  2662,
		},
	},
	"bintec-tapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bintec-tapi",
			Port:  2663,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bintec-tapi",
			Port:  2663,
		},
	},
	"biolink-auth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "biolink-auth",
			Port:  3411,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "biolink-auth",
			Port:  3411,
		},
	},
	"bioserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bioserver",
			Port:  6946,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bioserver",
			Port:  6946,
		},
	},
	"bip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bip",
			Port:  4376,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bip",
			Port:  4376,
		},
	},
	"bis-sync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bis-sync",
			Port:  5585,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bis-sync",
			Port:  5585,
		},
	},
	"bis-web": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bis-web",
			Port:  5584,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bis-web",
			Port:  5584,
		},
	},
	"bitforestsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bitforestsrv",
			Port:  5068,
		},
	},
	"bitspeer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bitspeer",
			Port:  2178,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bitspeer",
			Port:  2178,
		},
	},
	"bl-idm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bl-idm",
			Port:  142,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bl-idm",
			Port:  142,
		},
	},
	"blackboard": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blackboard",
			Port:  2032,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blackboard",
			Port:  2032,
		},
	},
	"blackjack": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blackjack",
			Port:  1025,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blackjack",
			Port:  1025,
		},
	},
	"blaze": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blaze",
			Port:  1150,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blaze",
			Port:  1150,
		},
	},
	"blizwow": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blizwow",
			Port:  3724,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blizwow",
			Port:  3724,
		},
	},
	"blockade": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blockade",
			Port:  2911,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blockade",
			Port:  2911,
		},
	},
	"blockade-bpsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blockade-bpsp",
			Port:  2574,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blockade-bpsp",
			Port:  2574,
		},
	},
	"blocks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blocks",
			Port:  10288,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blocks",
			Port:  10288,
		},
	},
	"blp1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blp1",
			Port:  8194,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blp1",
			Port:  8194,
		},
	},
	"blp2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blp2",
			Port:  8195,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blp2",
			Port:  8195,
		},
	},
	"blp3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blp3",
			Port:  8292,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blp3",
			Port:  8292,
		},
	},
	"blp4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blp4",
			Port:  8294,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blp4",
			Port:  8294,
		},
	},
	"blp5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blp5",
			Port:  48129,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blp5",
			Port:  48129,
		},
	},
	"blueberry-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blueberry-lm",
			Port:  1432,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blueberry-lm",
			Port:  1432,
		},
	},
	"bluectrlproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bluectrlproxy",
			Port:  2277,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bluectrlproxy",
			Port:  2277,
		},
	},
	"bluelance": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bluelance",
			Port:  2877,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bluelance",
			Port:  2877,
		},
	},
	"blwnkl-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "blwnkl-port",
			Port:  2625,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "blwnkl-port",
			Port:  2625,
		},
	},
	"bmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmap",
			Port:  3421,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmap",
			Port:  3421,
		},
	},
	"bmc-ar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-ar",
			Port:  2494,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-ar",
			Port:  2494,
		},
	},
	"bmc-data-coll": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-data-coll",
			Port:  3695,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-data-coll",
			Port:  3695,
		},
	},
	"bmc-ea": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-ea",
			Port:  3683,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-ea",
			Port:  3683,
		},
	},
	"bmc-gms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-gms",
			Port:  10129,
		},
	},
	"bmc-grx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-grx",
			Port:  6300,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-grx",
			Port:  6300,
		},
	},
	"bmc-jmx-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-jmx-port",
			Port:  3604,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-jmx-port",
			Port:  3604,
		},
	},
	"bmc-messaging": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-messaging",
			Port:  2059,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-messaging",
			Port:  2059,
		},
	},
	"bmc-net-adm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-net-adm",
			Port:  1769,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-net-adm",
			Port:  1769,
		},
	},
	"bmc-net-svc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-net-svc",
			Port:  1770,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-net-svc",
			Port:  1770,
		},
	},
	"bmc-onekey": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-onekey",
			Port:  3561,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-onekey",
			Port:  3561,
		},
	},
	"bmc-perf-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-perf-agent",
			Port:  6767,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-perf-agent",
			Port:  6767,
		},
	},
	"bmc-perf-mgrd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-perf-mgrd",
			Port:  6768,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-perf-mgrd",
			Port:  6768,
		},
	},
	"bmc-perf-sd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-perf-sd",
			Port:  10128,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-perf-sd",
			Port:  10128,
		},
	},
	"bmc-reporting": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc-reporting",
			Port:  4568,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc-reporting",
			Port:  4568,
		},
	},
	"bmc_ctd_ldap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmc_ctd_ldap",
			Port:  6301,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmc_ctd_ldap",
			Port:  6301,
		},
	},
	"bmc_patroldb": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "bmc_patroldb",
			Port:  1313,
		},
	},
	"bmcpatrolagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmcpatrolagent",
			Port:  3181,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmcpatrolagent",
			Port:  3181,
		},
	},
	"bmcpatrolrnvu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmcpatrolrnvu",
			Port:  3182,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmcpatrolrnvu",
			Port:  3182,
		},
	},
	"bmdss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmdss",
			Port:  13823,
		},
	},
	"bmpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bmpp",
			Port:  632,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bmpp",
			Port:  632,
		},
	},
	"bnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bnet",
			Port:  415,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bnet",
			Port:  415,
		},
	},
	"bnetfile": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bnetfile",
			Port:  1120,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bnetfile",
			Port:  1120,
		},
	},
	"bnetgame": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bnetgame",
			Port:  1119,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bnetgame",
			Port:  1119,
		},
	},
	"bnt-manager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bnt-manager",
			Port:  3344,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bnt-manager",
			Port:  3344,
		},
	},
	"board-roar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "board-roar",
			Port:  9700,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "board-roar",
			Port:  9700,
		},
	},
	"board-voip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "board-voip",
			Port:  9750,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "board-voip",
			Port:  9750,
		},
	},
	"boe-cachesvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-cachesvr",
			Port:  6403,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-cachesvr",
			Port:  6403,
		},
	},
	"boe-cms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-cms",
			Port:  6400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-cms",
			Port:  6400,
		},
	},
	"boe-eventsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-eventsrv",
			Port:  6402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-eventsrv",
			Port:  6402,
		},
	},
	"boe-filesvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-filesvr",
			Port:  6404,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-filesvr",
			Port:  6404,
		},
	},
	"boe-pagesvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-pagesvr",
			Port:  6405,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-pagesvr",
			Port:  6405,
		},
	},
	"boe-processsvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-processsvr",
			Port:  6406,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-processsvr",
			Port:  6406,
		},
	},
	"boe-resssvr1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-resssvr1",
			Port:  6407,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-resssvr1",
			Port:  6407,
		},
	},
	"boe-resssvr2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-resssvr2",
			Port:  6408,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-resssvr2",
			Port:  6408,
		},
	},
	"boe-resssvr3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-resssvr3",
			Port:  6409,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-resssvr3",
			Port:  6409,
		},
	},
	"boe-resssvr4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-resssvr4",
			Port:  6410,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-resssvr4",
			Port:  6410,
		},
	},
	"boe-was": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boe-was",
			Port:  6401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boe-was",
			Port:  6401,
		},
	},
	"boinc-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boinc-client",
			Port:  1043,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boinc-client",
			Port:  1043,
		},
	},
	"boks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boks",
			Port:  6500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boks",
			Port:  6500,
		},
	},
	"boks_clntd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boks_clntd",
			Port:  6503,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boks_clntd",
			Port:  6503,
		},
	},
	"boks_servc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boks_servc",
			Port:  6501,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boks_servc",
			Port:  6501,
		},
	},
	"boks_servm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boks_servm",
			Port:  6502,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boks_servm",
			Port:  6502,
		},
	},
	"boldsoft-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boldsoft-lm",
			Port:  2961,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boldsoft-lm",
			Port:  2961,
		},
	},
	"bones": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bones",
			Port:  4914,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bones",
			Port:  4914,
		},
	},
	"boomerang": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boomerang",
			Port:  1304,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boomerang",
			Port:  1304,
		},
	},
	"boosterware": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boosterware",
			Port:  2913,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boosterware",
			Port:  2913,
		},
	},
	"bootclient": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "bootclient",
			Port:  2017,
		},
	},
	"bootpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bootpc",
			Port:  68,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bootpc",
			Port:  68,
		},
	},
	"bootps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bootps",
			Port:  67,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bootps",
			Port:  67,
		},
	},
	"bootserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bootserver",
			Port:  2016,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bootserver",
			Port:  2016,
		},
	},
	"borland-dsj": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "borland-dsj",
			Port:  707,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "borland-dsj",
			Port:  707,
		},
	},
	"boscap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boscap",
			Port:  2990,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boscap",
			Port:  2990,
		},
	},
	"bounzza": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bounzza",
			Port:  2218,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bounzza",
			Port:  2218,
		},
	},
	"boxbackupstore": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boxbackupstore",
			Port:  4186,
		},
	},
	"boxp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "boxp",
			Port:  9380,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "boxp",
			Port:  9380,
		},
	},
	"bpcd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bpcd",
			Port:  13782,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bpcd",
			Port:  13782,
		},
	},
	"bpcp-poll": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bpcp-poll",
			Port:  2844,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bpcp-poll",
			Port:  2844,
		},
	},
	"bpcp-trap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bpcp-trap",
			Port:  2845,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bpcp-trap",
			Port:  2845,
		},
	},
	"bpdbm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bpdbm",
			Port:  13721,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bpdbm",
			Port:  13721,
		},
	},
	"bpjava-msvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bpjava-msvc",
			Port:  13722,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bpjava-msvc",
			Port:  13722,
		},
	},
	"bpmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bpmd",
			Port:  3593,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bpmd",
			Port:  3593,
		},
	},
	"bprd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bprd",
			Port:  13720,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bprd",
			Port:  13720,
		},
	},
	"br-channel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "br-channel",
			Port:  5425,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "br-channel",
			Port:  5425,
		},
	},
	"brain": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brain",
			Port:  2169,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brain",
			Port:  2169,
		},
	},
	"brcd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brcd",
			Port:  1323,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brcd",
			Port:  1323,
		},
	},
	"brcm-comm-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brcm-comm-port",
			Port:  3188,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brcm-comm-port",
			Port:  3188,
		},
	},
	"brdptc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brdptc",
			Port:  2155,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brdptc",
			Port:  2155,
		},
	},
	"bre": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bre",
			Port:  4096,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bre",
			Port:  4096,
		},
	},
	"brf-gw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brf-gw",
			Port:  22951,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brf-gw",
			Port:  22951,
		},
	},
	"bridgecontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bridgecontrol",
			Port:  1073,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bridgecontrol",
			Port:  1073,
		},
	},
	"brightcore": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "brightcore",
			Port:  5682,
		},
	},
	"brlp-0": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brlp-0",
			Port:  4101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brlp-0",
			Port:  4101,
		},
	},
	"brlp-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brlp-1",
			Port:  4102,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brlp-1",
			Port:  4102,
		},
	},
	"brlp-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brlp-2",
			Port:  4103,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brlp-2",
			Port:  4103,
		},
	},
	"brlp-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brlp-3",
			Port:  4104,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brlp-3",
			Port:  4104,
		},
	},
	"broker_service": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "broker_service",
			Port:  3014,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "broker_service",
			Port:  3014,
		},
	},
	"brp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brp",
			Port:  3043,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brp",
			Port:  3043,
		},
	},
	"bruce": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bruce",
			Port:  2619,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bruce",
			Port:  2619,
		},
	},
	"brutus": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "brutus",
			Port:  2003,
		},
	},
	"brvread": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "brvread",
			Port:  1054,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "brvread",
			Port:  1054,
		},
	},
	"bsfserver-zn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bsfserver-zn",
			Port:  5320,
		},
	},
	"bsfsvr-zn-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bsfsvr-zn-ssl",
			Port:  5321,
		},
	},
	"bspne-pcc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bspne-pcc",
			Port:  1252,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bspne-pcc",
			Port:  1252,
		},
	},
	"bsquare-voip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bsquare-voip",
			Port:  1071,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bsquare-voip",
			Port:  1071,
		},
	},
	"btpp2audctr1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "btpp2audctr1",
			Port:  2536,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "btpp2audctr1",
			Port:  2536,
		},
	},
	"btpp2sectrans": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "btpp2sectrans",
			Port:  2444,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "btpp2sectrans",
			Port:  2444,
		},
	},
	"btprjctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "btprjctrl",
			Port:  2803,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "btprjctrl",
			Port:  2803,
		},
	},
	"btrieve": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "btrieve",
			Port:  3351,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "btrieve",
			Port:  3351,
		},
	},
	"bts-appserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bts-appserver",
			Port:  1961,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bts-appserver",
			Port:  1961,
		},
	},
	"bts-x73": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bts-x73",
			Port:  3681,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bts-x73",
			Port:  3681,
		},
	},
	"buddy-draw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "buddy-draw",
			Port:  1854,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "buddy-draw",
			Port:  1854,
		},
	},
	"bues_service": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bues_service",
			Port:  2446,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bues_service",
			Port:  2446,
		},
	},
	"bullant-rap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bullant-rap",
			Port:  2965,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bullant-rap",
			Port:  2965,
		},
	},
	"bullant-srap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bullant-srap",
			Port:  2964,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bullant-srap",
			Port:  2964,
		},
	},
	"busboy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "busboy",
			Port:  998,
		},
	},
	"buschtrommel": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "buschtrommel",
			Port:  4747,
		},
	},
	"business": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "business",
			Port:  3107,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "business",
			Port:  3107,
		},
	},
	"busycal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "busycal",
			Port:  4990,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "busycal",
			Port:  4990,
		},
	},
	"bv-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bv-agent",
			Port:  3993,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bv-agent",
			Port:  3993,
		},
	},
	"bv-ds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bv-ds",
			Port:  3992,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bv-ds",
			Port:  3992,
		},
	},
	"bv-is": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bv-is",
			Port:  3990,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bv-is",
			Port:  3990,
		},
	},
	"bv-queryengine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bv-queryengine",
			Port:  3989,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bv-queryengine",
			Port:  3989,
		},
	},
	"bv-smcsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bv-smcsrv",
			Port:  3991,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bv-smcsrv",
			Port:  3991,
		},
	},
	"bvcdaemon-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bvcdaemon-port",
			Port:  3626,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bvcdaemon-port",
			Port:  3626,
		},
	},
	"bvcontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bvcontrol",
			Port:  1236,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bvcontrol",
			Port:  1236,
		},
	},
	"bveapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bveapi",
			Port:  10880,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bveapi",
			Port:  10880,
		},
	},
	"bvtsonar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bvtsonar",
			Port:  1149,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bvtsonar",
			Port:  1149,
		},
	},
	"bxp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bxp",
			Port:  4027,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bxp",
			Port:  4027,
		},
	},
	"bytex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bytex",
			Port:  1375,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bytex",
			Port:  1375,
		},
	},
	"bzflag": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bzflag",
			Port:  5154,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bzflag",
			Port:  5154,
		},
	},
	"bzr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "bzr",
			Port:  4155,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "bzr",
			Port:  4155,
		},
	},
	"c-h-it-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "c-h-it-port",
			Port:  3778,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "c-h-it-port",
			Port:  3778,
		},
	},
	"c1222-acse": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "c1222-acse",
			Port:  1153,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "c1222-acse",
			Port:  1153,
		},
	},
	"c3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "c3",
			Port:  2472,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "c3",
			Port:  2472,
		},
	},
	"ca-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ca-1",
			Port:  5064,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ca-1",
			Port:  5064,
		},
	},
	"ca-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ca-2",
			Port:  5065,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ca-2",
			Port:  5065,
		},
	},
	"ca-audit-da": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ca-audit-da",
			Port:  8025,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ca-audit-da",
			Port:  8025,
		},
	},
	"ca-audit-ds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ca-audit-ds",
			Port:  8026,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ca-audit-ds",
			Port:  8026,
		},
	},
	"ca-idms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ca-idms",
			Port:  3709,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ca-idms",
			Port:  3709,
		},
	},
	"ca-web-update": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ca-web-update",
			Port:  14414,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ca-web-update",
			Port:  14414,
		},
	},
	"caaclang2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caaclang2",
			Port:  5249,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caaclang2",
			Port:  5249,
		},
	},
	"caacws": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caacws",
			Port:  5248,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caacws",
			Port:  5248,
		},
	},
	"cab-protocol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cab-protocol",
			Port:  595,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cab-protocol",
			Port:  595,
		},
	},
	"cableport-ax": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cableport-ax",
			Port:  282,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cableport-ax",
			Port:  282,
		},
	},
	"cabsm-comm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cabsm-comm",
			Port:  7161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cabsm-comm",
			Port:  7161,
		},
	},
	"caci-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caci-lm",
			Port:  1554,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caci-lm",
			Port:  1554,
		},
	},
	"cacsambroker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cacsambroker",
			Port:  7163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cacsambroker",
			Port:  7163,
		},
	},
	"cadabra-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadabra-lm",
			Port:  1563,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadabra-lm",
			Port:  1563,
		},
	},
	"cadencecontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadencecontrol",
			Port:  2318,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadencecontrol",
			Port:  2318,
		},
	},
	"cadis-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadis-1",
			Port:  1441,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadis-1",
			Port:  1441,
		},
	},
	"cadis-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadis-2",
			Port:  1442,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadis-2",
			Port:  1442,
		},
	},
	"cadkey-licman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadkey-licman",
			Port:  1399,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadkey-licman",
			Port:  1399,
		},
	},
	"cadkey-tablet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadkey-tablet",
			Port:  1400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadkey-tablet",
			Port:  1400,
		},
	},
	"cadlock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadlock",
			Port:  770,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadlock",
			Port:  770,
		},
	},
	"cadlock2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadlock2",
			Port:  1000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadlock2",
			Port:  1000,
		},
	},
	"cadsi-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadsi-lm",
			Port:  1387,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadsi-lm",
			Port:  1387,
		},
	},
	"cadview-3d": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cadview-3d",
			Port:  649,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cadview-3d",
			Port:  649,
		},
	},
	"caerpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caerpc",
			Port:  42510,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caerpc",
			Port:  42510,
		},
	},
	"caevms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caevms",
			Port:  5251,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caevms",
			Port:  5251,
		},
	},
	"caicci": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caicci",
			Port:  1721,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caicci",
			Port:  1721,
		},
	},
	"caiccipc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caiccipc",
			Port:  1202,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caiccipc",
			Port:  1202,
		},
	},
	"caids-sensor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caids-sensor",
			Port:  1192,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caids-sensor",
			Port:  1192,
		},
	},
	"caistoragemgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caistoragemgr",
			Port:  7162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caistoragemgr",
			Port:  7162,
		},
	},
	"cajo-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cajo-discovery",
			Port:  1198,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cajo-discovery",
			Port:  1198,
		},
	},
	"cal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cal",
			Port:  588,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cal",
			Port:  588,
		},
	},
	"caldsoft-backup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caldsoft-backup",
			Port:  22537,
		},
	},
	"call-logging": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "call-logging",
			Port:  2552,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "call-logging",
			Port:  2552,
		},
	},
	"call-sig-trans": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "call-sig-trans",
			Port:  2517,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "call-sig-trans",
			Port:  2517,
		},
	},
	"caller9": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caller9",
			Port:  2906,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caller9",
			Port:  2906,
		},
	},
	"calltrax": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "calltrax",
			Port:  3675,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "calltrax",
			Port:  3675,
		},
	},
	"callwaveiam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "callwaveiam",
			Port:  9283,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "callwaveiam",
			Port:  9283,
		},
	},
	"camac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "camac",
			Port:  3545,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "camac",
			Port:  3545,
		},
	},
	"cambertx-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cambertx-lm",
			Port:  1734,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cambertx-lm",
			Port:  1734,
		},
	},
	"camp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "camp",
			Port:  4450,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "camp",
			Port:  4450,
		},
	},
	"can-dch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "can-dch",
			Port:  1919,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "can-dch",
			Port:  1919,
		},
	},
	"can-ferret": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "can-ferret",
			Port:  1920,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "can-ferret",
			Port:  1920,
		},
	},
	"can-ferret-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "can-ferret-ssl",
			Port:  3661,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "can-ferret-ssl",
			Port:  3661,
		},
	},
	"can-nds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "can-nds",
			Port:  1918,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "can-nds",
			Port:  1918,
		},
	},
	"can-nds-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "can-nds-ssl",
			Port:  3660,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "can-nds-ssl",
			Port:  3660,
		},
	},
	"canditv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canditv",
			Port:  24676,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canditv",
			Port:  24676,
		},
	},
	"candp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "candp",
			Port:  42508,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "candp",
			Port:  42508,
		},
	},
	"candrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "candrp",
			Port:  42509,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "candrp",
			Port:  42509,
		},
	},
	"canex-watch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canex-watch",
			Port:  3583,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canex-watch",
			Port:  3583,
		},
	},
	"canit_store": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canit_store",
			Port:  6568,
		},
	},
	"canna": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canna",
			Port:  5680,
		},
	},
	"canocentral0": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canocentral0",
			Port:  1871,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canocentral0",
			Port:  1871,
		},
	},
	"canocentral1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canocentral1",
			Port:  1872,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canocentral1",
			Port:  1872,
		},
	},
	"canon-bjnp1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canon-bjnp1",
			Port:  8611,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canon-bjnp1",
			Port:  8611,
		},
	},
	"canon-bjnp2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canon-bjnp2",
			Port:  8612,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canon-bjnp2",
			Port:  8612,
		},
	},
	"canon-bjnp3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canon-bjnp3",
			Port:  8613,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canon-bjnp3",
			Port:  8613,
		},
	},
	"canon-bjnp4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canon-bjnp4",
			Port:  8614,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canon-bjnp4",
			Port:  8614,
		},
	},
	"canon-capt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canon-capt",
			Port:  3756,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canon-capt",
			Port:  3756,
		},
	},
	"canon-cpp-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "canon-cpp-disc",
			Port:  8609,
		},
	},
	"canon-mfnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "canon-mfnp",
			Port:  8610,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "canon-mfnp",
			Port:  8610,
		},
	},
	"cap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cap",
			Port:  1026,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cap",
			Port:  1026,
		},
	},
	"capfast-lmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "capfast-lmd",
			Port:  1756,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "capfast-lmd",
			Port:  1756,
		},
	},
	"capioverlan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "capioverlan",
			Port:  1147,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "capioverlan",
			Port:  1147,
		},
	},
	"capmux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "capmux",
			Port:  4728,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "capmux",
			Port:  4728,
		},
	},
	"caps-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caps-lm",
			Port:  3290,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caps-lm",
			Port:  3290,
		},
	},
	"capwap-control": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "capwap-control",
			Port:  5246,
		},
	},
	"capwap-data": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "capwap-data",
			Port:  5247,
		},
	},
	"car": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "car",
			Port:  5090,
		},
	},
	"cardax": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cardax",
			Port:  1072,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cardax",
			Port:  1072,
		},
	},
	"cardbox": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cardbox",
			Port:  3105,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cardbox",
			Port:  3105,
		},
	},
	"cardbox-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cardbox-http",
			Port:  3106,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cardbox-http",
			Port:  3106,
		},
	},
	"carrius-rshell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "carrius-rshell",
			Port:  1197,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "carrius-rshell",
			Port:  1197,
		},
	},
	"cart-o-rama": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cart-o-rama",
			Port:  3292,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cart-o-rama",
			Port:  3292,
		},
	},
	"cartographerxmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cartographerxmp",
			Port:  5270,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cartographerxmp",
			Port:  5270,
		},
	},
	"cas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cas",
			Port:  2418,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cas",
			Port:  2418,
		},
	},
	"cas-mapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cas-mapi",
			Port:  3682,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cas-mapi",
			Port:  3682,
		},
	},
	"casanswmgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "casanswmgmt",
			Port:  3669,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "casanswmgmt",
			Port:  3669,
		},
	},
	"casp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "casp",
			Port:  1130,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "casp",
			Port:  1130,
		},
	},
	"caspssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caspssl",
			Port:  1131,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caspssl",
			Port:  1131,
		},
	},
	"casrmagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "casrmagent",
			Port:  7167,
		},
	},
	"castorproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "castorproxy",
			Port:  3450,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "castorproxy",
			Port:  3450,
		},
	},
	"catalyst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "catalyst",
			Port:  2836,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "catalyst",
			Port:  2836,
		},
	},
	"catchpole": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "catchpole",
			Port:  1185,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "catchpole",
			Port:  1185,
		},
	},
	"caupc-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "caupc-remote",
			Port:  2122,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "caupc-remote",
			Port:  2122,
		},
	},
	"cautcpd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cautcpd",
			Port:  3061,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cautcpd",
			Port:  3061,
		},
	},
	"cawas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cawas",
			Port:  12168,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cawas",
			Port:  12168,
		},
	},
	"cba8": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cba8",
			Port:  9593,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cba8",
			Port:  9593,
		},
	},
	"cbos-ip-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cbos-ip-port",
			Port:  3750,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cbos-ip-port",
			Port:  3750,
		},
	},
	"cbserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cbserver",
			Port:  3388,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cbserver",
			Port:  3388,
		},
	},
	"cbt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cbt",
			Port:  7777,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cbt",
			Port:  7777,
		},
	},
	"cc-tracking": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cc-tracking",
			Port:  4870,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cc-tracking",
			Port:  4870,
		},
	},
	"ccag-pib": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccag-pib",
			Port:  7169,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccag-pib",
			Port:  7169,
		},
	},
	"ccm-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccm-port",
			Port:  3575,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccm-port",
			Port:  3575,
		},
	},
	"ccmad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccmad",
			Port:  3114,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccmad",
			Port:  3114,
		},
	},
	"ccmail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccmail",
			Port:  3264,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccmail",
			Port:  3264,
		},
	},
	"ccmcomm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccmcomm",
			Port:  3505,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccmcomm",
			Port:  3505,
		},
	},
	"ccmrmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccmrmi",
			Port:  3154,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccmrmi",
			Port:  3154,
		},
	},
	"ccnx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccnx",
			Port:  9695,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccnx",
			Port:  9695,
		},
	},
	"ccowcmr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccowcmr",
			Port:  2116,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccowcmr",
			Port:  2116,
		},
	},
	"ccp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccp",
			Port:  3947,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccp",
			Port:  3947,
		},
	},
	"ccs-software": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccs-software",
			Port:  2734,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccs-software",
			Port:  2734,
		},
	},
	"ccss-qmm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccss-qmm",
			Port:  4969,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccss-qmm",
			Port:  4969,
		},
	},
	"ccss-qsm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccss-qsm",
			Port:  4970,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccss-qsm",
			Port:  4970,
		},
	},
	"cctv-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cctv-port",
			Port:  3559,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cctv-port",
			Port:  3559,
		},
	},
	"ccu-comm-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccu-comm-1",
			Port:  4053,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccu-comm-1",
			Port:  4053,
		},
	},
	"ccu-comm-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccu-comm-2",
			Port:  4054,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccu-comm-2",
			Port:  4054,
		},
	},
	"ccu-comm-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ccu-comm-3",
			Port:  4055,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ccu-comm-3",
			Port:  4055,
		},
	},
	"cd3o-protocol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cd3o-protocol",
			Port:  3616,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cd3o-protocol",
			Port:  3616,
		},
	},
	"cdbroker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cdbroker",
			Port:  3376,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cdbroker",
			Port:  3376,
		},
	},
	"cdc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cdc",
			Port:  223,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cdc",
			Port:  223,
		},
	},
	"cddbp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cddbp",
			Port:  888,
		},
	},
	"cddbp-alt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cddbp-alt",
			Port:  8880,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cddbp-alt",
			Port:  8880,
		},
	},
	"cdfunc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cdfunc",
			Port:  2045,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cdfunc",
			Port:  2045,
		},
	},
	"cdid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cdid",
			Port:  3315,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cdid",
			Port:  3315,
		},
	},
	"cdl-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cdl-server",
			Port:  3056,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cdl-server",
			Port:  3056,
		},
	},
	"cdn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cdn",
			Port:  2412,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cdn",
			Port:  2412,
		},
	},
	"cds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cds",
			Port:  4115,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cds",
			Port:  4115,
		},
	},
	"cecsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cecsvc",
			Port:  2571,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cecsvc",
			Port:  2571,
		},
	},
	"cedros_fds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cedros_fds",
			Port:  4140,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cedros_fds",
			Port:  4140,
		},
	},
	"celatalk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "celatalk",
			Port:  3485,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "celatalk",
			Port:  3485,
		},
	},
	"centerline": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "centerline",
			Port:  3987,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "centerline",
			Port:  3987,
		},
	},
	"centra": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "centra",
			Port:  1709,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "centra",
			Port:  1709,
		},
	},
	"cequint-cityid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cequint-cityid",
			Port:  4074,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cequint-cityid",
			Port:  4074,
		},
	},
	"cera-bcm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cera-bcm",
			Port:  1794,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cera-bcm",
			Port:  1794,
		},
	},
	"cernsysmgmtagt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cernsysmgmtagt",
			Port:  3830,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cernsysmgmtagt",
			Port:  3830,
		},
	},
	"cert-initiator": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cert-initiator",
			Port:  1639,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cert-initiator",
			Port:  1639,
		},
	},
	"cert-responder": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cert-responder",
			Port:  1640,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cert-responder",
			Port:  1640,
		},
	},
	"cesdcdman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cesdcdman",
			Port:  2921,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cesdcdman",
			Port:  2921,
		},
	},
	"cesdcdtrn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cesdcdtrn",
			Port:  2922,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cesdcdtrn",
			Port:  2922,
		},
	},
	"cesdinv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cesdinv",
			Port:  2856,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cesdinv",
			Port:  2856,
		},
	},
	"cfdptkt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cfdptkt",
			Port:  120,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cfdptkt",
			Port:  120,
		},
	},
	"cfengine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cfengine",
			Port:  5308,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cfengine",
			Port:  5308,
		},
	},
	"cfinger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cfinger",
			Port:  2003,
		},
	},
	"cfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cfs",
			Port:  7546,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cfs",
			Port:  7546,
		},
	},
	"cft-0": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cft-0",
			Port:  1761,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cft-0",
			Port:  1761,
		},
	},
	"cft-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cft-1",
			Port:  1762,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cft-1",
			Port:  1762,
		},
	},
	"cft-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cft-2",
			Port:  1763,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cft-2",
			Port:  1763,
		},
	},
	"cft-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cft-3",
			Port:  1764,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cft-3",
			Port:  1764,
		},
	},
	"cft-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cft-4",
			Port:  1765,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cft-4",
			Port:  1765,
		},
	},
	"cft-5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cft-5",
			Port:  1766,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cft-5",
			Port:  1766,
		},
	},
	"cft-6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cft-6",
			Port:  1767,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cft-6",
			Port:  1767,
		},
	},
	"cft-7": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cft-7",
			Port:  1768,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cft-7",
			Port:  1768,
		},
	},
	"cfw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cfw",
			Port:  7563,
		},
	},
	"cgi-starapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cgi-starapi",
			Port:  3893,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cgi-starapi",
			Port:  3893,
		},
	},
	"cgms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cgms",
			Port:  3003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cgms",
			Port:  3003,
		},
	},
	"cgn-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cgn-config",
			Port:  2183,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cgn-config",
			Port:  2183,
		},
	},
	"cgn-stat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cgn-stat",
			Port:  2182,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cgn-stat",
			Port:  2182,
		},
	},
	"chargen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "chargen",
			Port:  19,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "chargen",
			Port:  19,
		},
	},
	"charsetmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "charsetmgr",
			Port:  3903,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "charsetmgr",
			Port:  3903,
		},
	},
	"checkoutdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "checkoutdb",
			Port:  5505,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "checkoutdb",
			Port:  5505,
		},
	},
	"checkpoint-rtm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "checkpoint-rtm",
			Port:  18241,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "checkpoint-rtm",
			Port:  18241,
		},
	},
	"checksum": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "checksum",
			Port:  1386,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "checksum",
			Port:  1386,
		},
	},
	"chevinservices": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "chevinservices",
			Port:  3349,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "chevinservices",
			Port:  3349,
		},
	},
	"childkey-ctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "childkey-ctrl",
			Port:  1892,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "childkey-ctrl",
			Port:  1892,
		},
	},
	"childkey-notif": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "childkey-notif",
			Port:  1891,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "childkey-notif",
			Port:  1891,
		},
	},
	"chimera-hwm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "chimera-hwm",
			Port:  4009,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "chimera-hwm",
			Port:  4009,
		},
	},
	"chip-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "chip-lm",
			Port:  1572,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "chip-lm",
			Port:  1572,
		},
	},
	"chipper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "chipper",
			Port:  17219,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "chipper",
			Port:  17219,
		},
	},
	"chmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "chmd",
			Port:  3099,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "chmd",
			Port:  3099,
		},
	},
	"choiceview-agt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "choiceview-agt",
			Port:  4314,
		},
	},
	"choiceview-clt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "choiceview-clt",
			Port:  4316,
		},
	},
	"chromagrafx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "chromagrafx",
			Port:  1373,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "chromagrafx",
			Port:  1373,
		},
	},
	"chshell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "chshell",
			Port:  562,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "chshell",
			Port:  562,
		},
	},
	"ci3-software-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ci3-software-1",
			Port:  1301,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ci3-software-1",
			Port:  1301,
		},
	},
	"ci3-software-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ci3-software-2",
			Port:  1302,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ci3-software-2",
			Port:  1302,
		},
	},
	"cichild-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cichild-lm",
			Port:  1523,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cichild-lm",
			Port:  1523,
		},
	},
	"cichlid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cichlid",
			Port:  1377,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cichlid",
			Port:  1377,
		},
	},
	"cifs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cifs",
			Port:  3020,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cifs",
			Port:  3020,
		},
	},
	"cimplex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cimplex",
			Port:  673,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cimplex",
			Port:  673,
		},
	},
	"cimtrak": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cimtrak",
			Port:  3749,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cimtrak",
			Port:  3749,
		},
	},
	"cindycollab": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cindycollab",
			Port:  3770,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cindycollab",
			Port:  3770,
		},
	},
	"cinegrfx-elmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cinegrfx-elmd",
			Port:  2891,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cinegrfx-elmd",
			Port:  2891,
		},
	},
	"cinegrfx-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cinegrfx-lm",
			Port:  1743,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cinegrfx-lm",
			Port:  1743,
		},
	},
	"ciphire-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ciphire-data",
			Port:  3887,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ciphire-data",
			Port:  3887,
		},
	},
	"ciphire-serv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ciphire-serv",
			Port:  3888,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ciphire-serv",
			Port:  3888,
		},
	},
	"circle-x": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "circle-x",
			Port:  2931,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "circle-x",
			Port:  2931,
		},
	},
	"cis": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cis",
			Port:  22305,
		},
	},
	"cis-secure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cis-secure",
			Port:  22343,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cis-secure",
			Port:  22343,
		},
	},
	"cisco-avp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-avp",
			Port:  8470,
		},
	},
	"cisco-fna": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-fna",
			Port:  130,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-fna",
			Port:  130,
		},
	},
	"cisco-ipsla": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "cisco-ipsla",
			Port:  1167,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-ipsla",
			Port:  1167,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-ipsla",
			Port:  1167,
		},
	},
	"cisco-net-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-net-mgmt",
			Port:  1741,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-net-mgmt",
			Port:  1741,
		},
	},
	"cisco-redu": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-redu",
			Port:  5786,
		},
	},
	"cisco-snat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-snat",
			Port:  15555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-snat",
			Port:  15555,
		},
	},
	"cisco-sys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-sys",
			Port:  132,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-sys",
			Port:  132,
		},
	},
	"cisco-tdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-tdp",
			Port:  711,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-tdp",
			Port:  711,
		},
	},
	"cisco-tna": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-tna",
			Port:  131,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-tna",
			Port:  131,
		},
	},
	"cisco-vpath-tun": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-vpath-tun",
			Port:  6633,
		},
	},
	"cisco-wafs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cisco-wafs",
			Port:  4050,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cisco-wafs",
			Port:  4050,
		},
	},
	"ciscocsdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ciscocsdb",
			Port:  43441,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ciscocsdb",
			Port:  43441,
		},
	},
	"citadel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citadel",
			Port:  504,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "citadel",
			Port:  504,
		},
	},
	"citrix-rtmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citrix-rtmp",
			Port:  2897,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "citrix-rtmp",
			Port:  2897,
		},
	},
	"citrixadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citrixadmin",
			Port:  2513,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "citrixadmin",
			Port:  2513,
		},
	},
	"citrixima": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citrixima",
			Port:  2512,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "citrixima",
			Port:  2512,
		},
	},
	"citriximaclient": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citriximaclient",
			Port:  2598,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "citriximaclient",
			Port:  2598,
		},
	},
	"citrixupp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citrixupp",
			Port:  7228,
		},
	},
	"citrixuppg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citrixuppg",
			Port:  7229,
		},
	},
	"citynl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citynl",
			Port:  1729,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "citynl",
			Port:  1729,
		},
	},
	"citysearch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "citysearch",
			Port:  3974,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "citysearch",
			Port:  3974,
		},
	},
	"cl-db-attach": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cl-db-attach",
			Port:  4135,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cl-db-attach",
			Port:  4135,
		},
	},
	"cl-db-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cl-db-remote",
			Port:  4137,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cl-db-remote",
			Port:  4137,
		},
	},
	"cl-db-request": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cl-db-request",
			Port:  4136,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cl-db-request",
			Port:  4136,
		},
	},
	"cl/1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cl/1",
			Port:  172,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cl/1",
			Port:  172,
		},
	},
	"clariion-evr01": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clariion-evr01",
			Port:  6389,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "clariion-evr01",
			Port:  6389,
		},
	},
	"classic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "classic",
			Port:  9087,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "classic",
			Port:  9087,
		},
	},
	"cleanerliverc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cleanerliverc",
			Port:  3481,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cleanerliverc",
			Port:  3481,
		},
	},
	"clearcase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clearcase",
			Port:  371,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "clearcase",
			Port:  371,
		},
	},
	"clearvisn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clearvisn",
			Port:  2052,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "clearvisn",
			Port:  2052,
		},
	},
	"clever-ctrace": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clever-ctrace",
			Port:  6687,
		},
	},
	"clever-tcpip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clever-tcpip",
			Port:  6688,
		},
	},
	"client-ctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "client-ctrl",
			Port:  3730,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "client-ctrl",
			Port:  3730,
		},
	},
	"client-wakeup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "client-wakeup",
			Port:  9694,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "client-wakeup",
			Port:  9694,
		},
	},
	"cloanto-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cloanto-lm",
			Port:  3397,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cloanto-lm",
			Port:  3397,
		},
	},
	"cloanto-net-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cloanto-net-1",
			Port:  356,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cloanto-net-1",
			Port:  356,
		},
	},
	"close-combat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "close-combat",
			Port:  1944,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "close-combat",
			Port:  1944,
		},
	},
	"cloudsignaling": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cloudsignaling",
			Port:  7550,
		},
	},
	"clp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clp",
			Port:  2567,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "clp",
			Port:  2567,
		},
	},
	"cluster-disc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cluster-disc",
			Port:  3374,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cluster-disc",
			Port:  3374,
		},
	},
	"clusterxl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clusterxl",
			Port:  18243,
		},
	},
	"clutild": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clutild",
			Port:  7174,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "clutild",
			Port:  7174,
		},
	},
	"clvm-cfg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "clvm-cfg",
			Port:  1476,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "clvm-cfg",
			Port:  1476,
		},
	},
	"cm": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "cm",
			Port:  5910,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "cm",
			Port:  5910,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cm",
			Port:  5910,
		},
	},
	"cma": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cma",
			Port:  1050,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cma",
			Port:  1050,
		},
	},
	"cmadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cmadmin",
			Port:  2617,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cmadmin",
			Port:  2617,
		},
	},
	"cmc-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cmc-port",
			Port:  3576,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cmc-port",
			Port:  3576,
		},
	},
	"cmip-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cmip-agent",
			Port:  164,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cmip-agent",
			Port:  164,
		},
	},
	"cmip-man": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cmip-man",
			Port:  163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cmip-man",
			Port:  163,
		},
	},
	"cmmdriver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cmmdriver",
			Port:  1294,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cmmdriver",
			Port:  1294,
		},
	},
	"cmtp-av": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cmtp-av",
			Port:  8501,
		},
	},
	"cmtp-mgt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cmtp-mgt",
			Port:  8501,
		},
	},
	"cnap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cnap",
			Port:  7262,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cnap",
			Port:  7262,
		},
	},
	"cnckadserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cnckadserver",
			Port:  7168,
		},
	},
	"cncp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cncp",
			Port:  4785,
		},
	},
	"cnhrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cnhrp",
			Port:  1757,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cnhrp",
			Port:  1757,
		},
	},
	"cnrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cnrp",
			Port:  2757,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cnrp",
			Port:  2757,
		},
	},
	"cnrprotocol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cnrprotocol",
			Port:  1096,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cnrprotocol",
			Port:  1096,
		},
	},
	"cns-srv-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cns-srv-port",
			Port:  2976,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cns-srv-port",
			Port:  2976,
		},
	},
	"coap": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "coap",
			Port:  5683,
		},
	},
	"coauthor": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "coauthor",
			Port:  1529,
		},
	},
	"codaauth2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "codaauth2",
			Port:  370,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "codaauth2",
			Port:  370,
		},
	},
	"codasrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "codasrv",
			Port:  2432,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "codasrv",
			Port:  2432,
		},
	},
	"codasrv-se": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "codasrv-se",
			Port:  2433,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "codasrv-se",
			Port:  2433,
		},
	},
	"codima-rtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "codima-rtp",
			Port:  2415,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "codima-rtp",
			Port:  2415,
		},
	},
	"cogitate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cogitate",
			Port:  3039,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cogitate",
			Port:  3039,
		},
	},
	"cognex-dataman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cognex-dataman",
			Port:  44444,
		},
	},
	"cognex-insight": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cognex-insight",
			Port:  1069,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cognex-insight",
			Port:  1069,
		},
	},
	"cognima": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cognima",
			Port:  3779,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cognima",
			Port:  3779,
		},
	},
	"cogsys-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cogsys-lm",
			Port:  3377,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cogsys-lm",
			Port:  3377,
		},
	},
	"collaber": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "collaber",
			Port:  7689,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "collaber",
			Port:  7689,
		},
	},
	"collaborator": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "collaborator",
			Port:  622,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "collaborator",
			Port:  622,
		},
	},
	"colubris": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "colubris",
			Port:  3490,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "colubris",
			Port:  3490,
		},
	},
	"com-bardac-dw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "com-bardac-dw",
			Port:  48556,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "com-bardac-dw",
			Port:  48556,
		},
	},
	"combox-web-acc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "combox-web-acc",
			Port:  2534,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "combox-web-acc",
			Port:  2534,
		},
	},
	"comcam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "comcam",
			Port:  2108,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "comcam",
			Port:  2108,
		},
	},
	"comcam-io": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "comcam-io",
			Port:  3605,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "comcam-io",
			Port:  3605,
		},
	},
	"commandport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "commandport",
			Port:  3416,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "commandport",
			Port:  3416,
		},
	},
	"commerce": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "commerce",
			Port:  542,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "commerce",
			Port:  542,
		},
	},
	"commlinx-avl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "commlinx-avl",
			Port:  1190,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "commlinx-avl",
			Port:  1190,
		},
	},
	"commonspace": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "commonspace",
			Port:  1592,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "commonspace",
			Port:  1592,
		},
	},
	"commplex-link": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "commplex-link",
			Port:  5001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "commplex-link",
			Port:  5001,
		},
	},
	"commplex-main": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "commplex-main",
			Port:  5000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "commplex-main",
			Port:  5000,
		},
	},
	"commtact-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "commtact-http",
			Port:  20002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "commtact-http",
			Port:  20002,
		},
	},
	"commtact-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "commtact-https",
			Port:  20003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "commtact-https",
			Port:  20003,
		},
	},
	"community": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "community",
			Port:  2459,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "community",
			Port:  2459,
		},
	},
	"comotionback": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "comotionback",
			Port:  2262,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "comotionback",
			Port:  2262,
		},
	},
	"comotionmaster": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "comotionmaster",
			Port:  2261,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "comotionmaster",
			Port:  2261,
		},
	},
	"compaq-evm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "compaq-evm",
			Port:  619,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "compaq-evm",
			Port:  619,
		},
	},
	"compaq-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "compaq-https",
			Port:  2381,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "compaq-https",
			Port:  2381,
		},
	},
	"compaq-scp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "compaq-scp",
			Port:  2766,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "compaq-scp",
			Port:  2766,
		},
	},
	"compaq-wcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "compaq-wcp",
			Port:  2555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "compaq-wcp",
			Port:  2555,
		},
	},
	"composit-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "composit-server",
			Port:  2417,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "composit-server",
			Port:  2417,
		},
	},
	"compressnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "compressnet",
			Port:  2,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "compressnet",
			Port:  2,
		},
	},
	"compx-lockview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "compx-lockview",
			Port:  4308,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "compx-lockview",
			Port:  4308,
		},
	},
	"comscm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "comscm",
			Port:  437,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "comscm",
			Port:  437,
		},
	},
	"con": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "con",
			Port:  759,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "con",
			Port:  759,
		},
	},
	"conclave-cpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "conclave-cpp",
			Port:  2491,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "conclave-cpp",
			Port:  2491,
		},
	},
	"concomp1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "concomp1",
			Port:  1802,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "concomp1",
			Port:  1802,
		},
	},
	"concurrent-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "concurrent-lm",
			Port:  1648,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "concurrent-lm",
			Port:  1648,
		},
	},
	"condor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "condor",
			Port:  9618,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "condor",
			Port:  9618,
		},
	},
	"conf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "conf",
			Port:  2008,
		},
	},
	"conference": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "conference",
			Port:  531,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "conference",
			Port:  531,
		},
	},
	"conferencetalk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "conferencetalk",
			Port:  1713,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "conferencetalk",
			Port:  1713,
		},
	},
	"config-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "config-port",
			Port:  3577,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "config-port",
			Port:  3577,
		},
	},
	"confluent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "confluent",
			Port:  1484,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "confluent",
			Port:  1484,
		},
	},
	"connect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "connect",
			Port:  2137,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "connect",
			Port:  2137,
		},
	},
	"connect-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "connect-client",
			Port:  3441,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "connect-client",
			Port:  3441,
		},
	},
	"connect-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "connect-server",
			Port:  3442,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "connect-server",
			Port:  3442,
		},
	},
	"connected": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "connected",
			Port:  16384,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "connected",
			Port:  16384,
		},
	},
	"connection": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "connection",
			Port:  2607,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "connection",
			Port:  2607,
		},
	},
	"connendp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "connendp",
			Port:  693,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "connendp",
			Port:  693,
		},
	},
	"connlcli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "connlcli",
			Port:  1358,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "connlcli",
			Port:  1358,
		},
	},
	"conspiracy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "conspiracy",
			Port:  4692,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "conspiracy",
			Port:  4692,
		},
	},
	"consul-insight": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "consul-insight",
			Port:  5992,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "consul-insight",
			Port:  5992,
		},
	},
	"contamac_icm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "contamac_icm",
			Port:  4846,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "contamac_icm",
			Port:  4846,
		},
	},
	"contclientms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "contclientms",
			Port:  4665,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "contclientms",
			Port:  4665,
		},
	},
	"contentserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "contentserver",
			Port:  454,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "contentserver",
			Port:  454,
		},
	},
	"continuus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "continuus",
			Port:  5412,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "continuus",
			Port:  5412,
		},
	},
	"coord-svr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "coord-svr",
			Port:  2565,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "coord-svr",
			Port:  2565,
		},
	},
	"cops": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cops",
			Port:  3288,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cops",
			Port:  3288,
		},
	},
	"cops-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cops-tls",
			Port:  3183,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cops-tls",
			Port:  3183,
		},
	},
	"copy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "copy",
			Port:  8445,
		},
	},
	"copy-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "copy-disc",
			Port:  8445,
		},
	},
	"copycat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "copycat",
			Port:  9093,
		},
	},
	"corba-iiop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "corba-iiop",
			Port:  683,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "corba-iiop",
			Port:  683,
		},
	},
	"corba-iiop-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "corba-iiop-ssl",
			Port:  684,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "corba-iiop-ssl",
			Port:  684,
		},
	},
	"corbaloc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "corbaloc",
			Port:  2809,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "corbaloc",
			Port:  2809,
		},
	},
	"corel_vncadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "corel_vncadmin",
			Port:  2654,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "corel_vncadmin",
			Port:  2654,
		},
	},
	"corelccam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "corelccam",
			Port:  4300,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "corelccam",
			Port:  4300,
		},
	},
	"corelvideo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "corelvideo",
			Port:  1566,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "corelvideo",
			Port:  1566,
		},
	},
	"corerjd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "corerjd",
			Port:  284,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "corerjd",
			Port:  284,
		},
	},
	"cosir": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cosir",
			Port:  10321,
		},
	},
	"cosmocall": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cosmocall",
			Port:  2324,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cosmocall",
			Port:  2324,
		},
	},
	"couchdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "couchdb",
			Port:  5984,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "couchdb",
			Port:  5984,
		},
	},
	"courier": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "courier",
			Port:  530,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "courier",
			Port:  530,
		},
	},
	"covia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "covia",
			Port:  64,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "covia",
			Port:  64,
		},
	},
	"cp-cluster": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cp-cluster",
			Port:  8116,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cp-cluster",
			Port:  8116,
		},
	},
	"cp-spxdpy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cp-spxdpy",
			Port:  4378,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cp-spxdpy",
			Port:  4378,
		},
	},
	"cp-spxrpts": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cp-spxrpts",
			Port:  5079,
		},
	},
	"cp-spxsvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cp-spxsvr",
			Port:  4377,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cp-spxsvr",
			Port:  4377,
		},
	},
	"cpdi-pidas-cm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cpdi-pidas-cm",
			Port:  3609,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cpdi-pidas-cm",
			Port:  3609,
		},
	},
	"cpdlc": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "cpdlc",
			Port:  5911,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "cpdlc",
			Port:  5911,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cpdlc",
			Port:  5911,
		},
	},
	"cplscrambler-al": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cplscrambler-al",
			Port:  1088,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cplscrambler-al",
			Port:  1088,
		},
	},
	"cplscrambler-in": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cplscrambler-in",
			Port:  1087,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cplscrambler-in",
			Port:  1087,
		},
	},
	"cplscrambler-lg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cplscrambler-lg",
			Port:  1086,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cplscrambler-lg",
			Port:  1086,
		},
	},
	"cppdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cppdp",
			Port:  4051,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cppdp",
			Port:  4051,
		},
	},
	"cpq-tasksmart": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cpq-tasksmart",
			Port:  3201,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cpq-tasksmart",
			Port:  3201,
		},
	},
	"cpq-wbem": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cpq-wbem",
			Port:  2301,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cpq-wbem",
			Port:  2301,
		},
	},
	"cpqrpm-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cpqrpm-agent",
			Port:  3256,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cpqrpm-agent",
			Port:  3256,
		},
	},
	"cpqrpm-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cpqrpm-server",
			Port:  3257,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cpqrpm-server",
			Port:  3257,
		},
	},
	"cps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cps",
			Port:  14250,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cps",
			Port:  14250,
		},
	},
	"cpscomm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cpscomm",
			Port:  5194,
		},
	},
	"cpsp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cpsp",
			Port:  17222,
		},
	},
	"cpudpencap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cpudpencap",
			Port:  2746,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cpudpencap",
			Port:  2746,
		},
	},
	"cqg-netlan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cqg-netlan",
			Port:  2823,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cqg-netlan",
			Port:  2823,
		},
	},
	"cqg-netlan-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cqg-netlan-1",
			Port:  2824,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cqg-netlan-1",
			Port:  2824,
		},
	},
	"cr-websystems": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cr-websystems",
			Port:  2314,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cr-websystems",
			Port:  2314,
		},
	},
	"creativepartnr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "creativepartnr",
			Port:  455,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "creativepartnr",
			Port:  455,
		},
	},
	"creativeserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "creativeserver",
			Port:  453,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "creativeserver",
			Port:  453,
		},
	},
	"crestron-cip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "crestron-cip",
			Port:  41794,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "crestron-cip",
			Port:  41794,
		},
	},
	"crestron-cips": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "crestron-cips",
			Port:  41796,
		},
	},
	"crestron-ctp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "crestron-ctp",
			Port:  41795,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "crestron-ctp",
			Port:  41795,
		},
	},
	"crestron-ctps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "crestron-ctps",
			Port:  41797,
		},
	},
	"crinis-hb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "crinis-hb",
			Port:  3818,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "crinis-hb",
			Port:  3818,
		},
	},
	"crip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "crip",
			Port:  6253,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "crip",
			Port:  6253,
		},
	},
	"crmsbits": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "crmsbits",
			Port:  2422,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "crmsbits",
			Port:  2422,
		},
	},
	"crs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "crs",
			Port:  507,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "crs",
			Port:  507,
		},
	},
	"cruise-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cruise-config",
			Port:  8378,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cruise-config",
			Port:  8378,
		},
	},
	"cruise-diags": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cruise-diags",
			Port:  8379,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cruise-diags",
			Port:  8379,
		},
	},
	"cruise-enum": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cruise-enum",
			Port:  8376,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cruise-enum",
			Port:  8376,
		},
	},
	"cruise-swroute": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cruise-swroute",
			Port:  8377,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cruise-swroute",
			Port:  8377,
		},
	},
	"cruise-update": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cruise-update",
			Port:  8380,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cruise-update",
			Port:  8380,
		},
	},
	"cryptoadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cryptoadmin",
			Port:  624,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cryptoadmin",
			Port:  624,
		},
	},
	"cs-auth-svr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cs-auth-svr",
			Port:  3113,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cs-auth-svr",
			Port:  3113,
		},
	},
	"cs-live": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cs-live",
			Port:  2129,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cs-live",
			Port:  2129,
		},
	},
	"cs-remote-db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cs-remote-db",
			Port:  3630,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cs-remote-db",
			Port:  3630,
		},
	},
	"cs-services": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cs-services",
			Port:  3631,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cs-services",
			Port:  3631,
		},
	},
	"csbphonemaster": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csbphonemaster",
			Port:  1724,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csbphonemaster",
			Port:  1724,
		},
	},
	"csc_proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csc_proxy",
			Port:  4187,
		},
	},
	"csccfirewall": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csccfirewall",
			Port:  40843,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csccfirewall",
			Port:  40843,
		},
	},
	"csccredir": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csccredir",
			Port:  40842,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csccredir",
			Port:  40842,
		},
	},
	"cscp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cscp",
			Port:  40841,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cscp",
			Port:  40841,
		},
	},
	"csd-mgmt-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csd-mgmt-port",
			Port:  3071,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csd-mgmt-port",
			Port:  3071,
		},
	},
	"csd-monitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csd-monitor",
			Port:  3072,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csd-monitor",
			Port:  3072,
		},
	},
	"csdm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csdm",
			Port:  1468,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csdm",
			Port:  1468,
		},
	},
	"csdmbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csdmbase",
			Port:  1467,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csdmbase",
			Port:  1467,
		},
	},
	"csi-lfap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csi-lfap",
			Port:  3145,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csi-lfap",
			Port:  3145,
		},
	},
	"csi-sgwp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csi-sgwp",
			Port:  348,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csi-sgwp",
			Port:  348,
		},
	},
	"cslg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cslg",
			Port:  24754,
		},
	},
	"cslistener": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cslistener",
			Port:  9000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cslistener",
			Port:  9000,
		},
	},
	"csms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csms",
			Port:  3399,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csms",
			Port:  3399,
		},
	},
	"csms2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csms2",
			Port:  3400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csms2",
			Port:  3400,
		},
	},
	"csnet-ns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csnet-ns",
			Port:  105,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csnet-ns",
			Port:  105,
		},
	},
	"csnotify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csnotify",
			Port:  2955,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csnotify",
			Port:  2955,
		},
	},
	"csoft-plusclnt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csoft-plusclnt",
			Port:  2699,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csoft-plusclnt",
			Port:  2699,
		},
	},
	"csoft-prev": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csoft-prev",
			Port:  3271,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csoft-prev",
			Port:  3271,
		},
	},
	"csoft1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csoft1",
			Port:  1837,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csoft1",
			Port:  1837,
		},
	},
	"csoftragent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csoftragent",
			Port:  3004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csoftragent",
			Port:  3004,
		},
	},
	"cspclmulti": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cspclmulti",
			Port:  2890,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cspclmulti",
			Port:  2890,
		},
	},
	"cspmlockmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cspmlockmgr",
			Port:  1272,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cspmlockmgr",
			Port:  1272,
		},
	},
	"cspmulti": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cspmulti",
			Port:  2807,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cspmulti",
			Port:  2807,
		},
	},
	"cspuni": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cspuni",
			Port:  2806,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cspuni",
			Port:  2806,
		},
	},
	"csregagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csregagent",
			Port:  3022,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csregagent",
			Port:  3022,
		},
	},
	"csrpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csrpc",
			Port:  5063,
		},
	},
	"cssc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cssc",
			Port:  5637,
		},
	},
	"cssp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cssp",
			Port:  4078,
		},
	},
	"cst-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cst-port",
			Port:  3742,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cst-port",
			Port:  3742,
		},
	},
	"csvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csvr",
			Port:  3417,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csvr",
			Port:  3417,
		},
	},
	"csvr-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csvr-proxy",
			Port:  3190,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csvr-proxy",
			Port:  3190,
		},
	},
	"csvr-sslproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "csvr-sslproxy",
			Port:  3191,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "csvr-sslproxy",
			Port:  3191,
		},
	},
	"ct2nmcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ct2nmcs",
			Port:  7023,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ct2nmcs",
			Port:  7023,
		},
	},
	"ctcd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctcd",
			Port:  1851,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctcd",
			Port:  1851,
		},
	},
	"ctdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctdb",
			Port:  4379,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctdb",
			Port:  4379,
		},
	},
	"ctdhercules": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctdhercules",
			Port:  3773,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctdhercules",
			Port:  3773,
		},
	},
	"ctdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctdp",
			Port:  7022,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctdp",
			Port:  7022,
		},
	},
	"ctechlicensing": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctechlicensing",
			Port:  9346,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctechlicensing",
			Port:  9346,
		},
	},
	"ctf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctf",
			Port:  84,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctf",
			Port:  84,
		},
	},
	"cti-redwood": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cti-redwood",
			Port:  2563,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cti-redwood",
			Port:  2563,
		},
	},
	"ctiprogramload": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctiprogramload",
			Port:  4452,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctiprogramload",
			Port:  4452,
		},
	},
	"ctisystemmsg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctisystemmsg",
			Port:  4451,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctisystemmsg",
			Port:  4451,
		},
	},
	"ctlptc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctlptc",
			Port:  2153,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctlptc",
			Port:  2153,
		},
	},
	"ctp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctp",
			Port:  3772,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctp",
			Port:  3772,
		},
	},
	"ctp-state": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctp-state",
			Port:  4047,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctp-state",
			Port:  4047,
		},
	},
	"ctsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctsd",
			Port:  5137,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctsd",
			Port:  5137,
		},
	},
	"ctt-broker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctt-broker",
			Port:  1932,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctt-broker",
			Port:  1932,
		},
	},
	"ctx-bridge": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctx-bridge",
			Port:  3127,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctx-bridge",
			Port:  3127,
		},
	},
	"ctxlic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ctxlic",
			Port:  7279,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ctxlic",
			Port:  7279,
		},
	},
	"cuelink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cuelink",
			Port:  5271,
		},
	},
	"cuelink-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "cuelink-disc",
			Port:  5271,
		},
	},
	"cuillamartin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cuillamartin",
			Port:  1356,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cuillamartin",
			Port:  1356,
		},
	},
	"cumulus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cumulus",
			Port:  9287,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cumulus",
			Port:  9287,
		},
	},
	"cumulus-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cumulus-admin",
			Port:  8954,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cumulus-admin",
			Port:  8954,
		},
	},
	"cuseeme": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cuseeme",
			Port:  7648,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cuseeme",
			Port:  7648,
		},
	},
	"custix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "custix",
			Port:  528,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "custix",
			Port:  528,
		},
	},
	"cvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cvc",
			Port:  1495,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cvc",
			Port:  1495,
		},
	},
	"cvc_hostd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cvc_hostd",
			Port:  442,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cvc_hostd",
			Port:  442,
		},
	},
	"cvd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cvd",
			Port:  8400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cvd",
			Port:  8400,
		},
	},
	"cvmmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cvmmon",
			Port:  2300,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cvmmon",
			Port:  2300,
		},
	},
	"cvmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cvmon",
			Port:  1686,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cvmon",
			Port:  1686,
		},
	},
	"cvspserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cvspserver",
			Port:  2401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cvspserver",
			Port:  2401,
		},
	},
	"cvsup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cvsup",
			Port:  5999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cvsup",
			Port:  5999,
		},
	},
	"cwmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cwmp",
			Port:  7547,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cwmp",
			Port:  7547,
		},
	},
	"cxtp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "cxtp",
			Port:  5091,
		},
	},
	"cxws": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cxws",
			Port:  4673,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cxws",
			Port:  4673,
		},
	},
	"cyaserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cyaserv",
			Port:  2584,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cyaserv",
			Port:  2584,
		},
	},
	"cybercash": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cybercash",
			Port:  551,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cybercash",
			Port:  551,
		},
	},
	"cyborg-systems": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cyborg-systems",
			Port:  9888,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cyborg-systems",
			Port:  9888,
		},
	},
	"cybro-a-bus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cybro-a-bus",
			Port:  8442,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cybro-a-bus",
			Port:  8442,
		},
	},
	"cyc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cyc",
			Port:  3645,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cyc",
			Port:  3645,
		},
	},
	"cycleserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cycleserv",
			Port:  763,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cycleserv",
			Port:  763,
		},
	},
	"cycleserv2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cycleserv2",
			Port:  772,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cycleserv2",
			Port:  772,
		},
	},
	"cylink-c": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cylink-c",
			Port:  5420,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cylink-c",
			Port:  5420,
		},
	},
	"cymtec-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cymtec-port",
			Port:  1898,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cymtec-port",
			Port:  1898,
		},
	},
	"cypress": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cypress",
			Port:  2015,
		},
	},
	"cypress-stat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cypress-stat",
			Port:  2017,
		},
	},
	"cytel-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "cytel-lm",
			Port:  3297,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "cytel-lm",
			Port:  3297,
		},
	},
	"d-cinema-csp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d-cinema-csp",
			Port:  4170,
		},
	},
	"d-cinema-rrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d-cinema-rrp",
			Port:  1173,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d-cinema-rrp",
			Port:  1173,
		},
	},
	"d-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d-data",
			Port:  4301,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d-data",
			Port:  4301,
		},
	},
	"d-data-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d-data-control",
			Port:  4302,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d-data-control",
			Port:  4302,
		},
	},
	"d-fence": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d-fence",
			Port:  8555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d-fence",
			Port:  8555,
		},
	},
	"d-s-n": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d-s-n",
			Port:  8086,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d-s-n",
			Port:  8086,
		},
	},
	"d2000kernel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d2000kernel",
			Port:  3119,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d2000kernel",
			Port:  3119,
		},
	},
	"d2000webserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d2000webserver",
			Port:  3120,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d2000webserver",
			Port:  3120,
		},
	},
	"d2dconfig": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d2dconfig",
			Port:  9387,
		},
	},
	"d2ddatatrans": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d2ddatatrans",
			Port:  9388,
		},
	},
	"d2k-datamover1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d2k-datamover1",
			Port:  2297,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d2k-datamover1",
			Port:  2297,
		},
	},
	"d2k-datamover2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d2k-datamover2",
			Port:  2298,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d2k-datamover2",
			Port:  2298,
		},
	},
	"d2k-tapestry1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d2k-tapestry1",
			Port:  3393,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d2k-tapestry1",
			Port:  3393,
		},
	},
	"d2k-tapestry2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d2k-tapestry2",
			Port:  3394,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d2k-tapestry2",
			Port:  3394,
		},
	},
	"d3winosfi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "d3winosfi",
			Port:  3458,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "d3winosfi",
			Port:  3458,
		},
	},
	"daap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "daap",
			Port:  3689,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "daap",
			Port:  3689,
		},
	},
	"dab-sti-c": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dab-sti-c",
			Port:  1076,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dab-sti-c",
			Port:  1076,
		},
	},
	"dai-shell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dai-shell",
			Port:  45824,
		},
	},
	"daishi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "daishi",
			Port:  2870,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "daishi",
			Port:  2870,
		},
	},
	"dali-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dali-port",
			Port:  5777,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dali-port",
			Port:  5777,
		},
	},
	"dandv-tester": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dandv-tester",
			Port:  3889,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dandv-tester",
			Port:  3889,
		},
	},
	"danf-ak2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "danf-ak2",
			Port:  1041,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "danf-ak2",
			Port:  1041,
		},
	},
	"daqstream": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "daqstream",
			Port:  7411,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "daqstream",
			Port:  7411,
		},
	},
	"darcorp-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "darcorp-lm",
			Port:  1679,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "darcorp-lm",
			Port:  1679,
		},
	},
	"dashpas-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dashpas-port",
			Port:  3498,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dashpas-port",
			Port:  3498,
		},
	},
	"dasp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dasp",
			Port:  439,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dasp",
			Port:  439,
		},
	},
	"data-insurance": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "data-insurance",
			Port:  2764,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "data-insurance",
			Port:  2764,
		},
	},
	"data-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "data-port",
			Port:  3578,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "data-port",
			Port:  3578,
		},
	},
	"datacaptor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datacaptor",
			Port:  1857,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "datacaptor",
			Port:  1857,
		},
	},
	"datalens": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datalens",
			Port:  2229,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "datalens",
			Port:  2229,
		},
	},
	"datametrics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datametrics",
			Port:  1645,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "datametrics",
			Port:  1645,
		},
	},
	"datascaler-ctl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datascaler-ctl",
			Port:  6625,
		},
	},
	"datascaler-db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datascaler-db",
			Port:  6624,
		},
	},
	"datasurfsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datasurfsrv",
			Port:  461,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "datasurfsrv",
			Port:  461,
		},
	},
	"datasurfsrvsec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datasurfsrvsec",
			Port:  462,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "datasurfsrvsec",
			Port:  462,
		},
	},
	"datex-asn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datex-asn",
			Port:  355,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "datex-asn",
			Port:  355,
		},
	},
	"datusorb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "datusorb",
			Port:  3282,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "datusorb",
			Port:  3282,
		},
	},
	"davsrc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "davsrc",
			Port:  9800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "davsrc",
			Port:  9800,
		},
	},
	"davsrcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "davsrcs",
			Port:  9802,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "davsrcs",
			Port:  9802,
		},
	},
	"dawn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dawn",
			Port:  1908,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dawn",
			Port:  1908,
		},
	},
	"dayliteserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dayliteserver",
			Port:  6113,
		},
	},
	"daylitetouch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "daylitetouch",
			Port:  6117,
		},
	},
	"daytime": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "daytime",
			Port:  13,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "daytime",
			Port:  13,
		},
	},
	"db-lsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "db-lsp",
			Port:  17500,
		},
	},
	"db-lsp-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "db-lsp-disc",
			Port:  17500,
		},
	},
	"dbabble": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbabble",
			Port:  8132,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbabble",
			Port:  8132,
		},
	},
	"dbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbase",
			Port:  217,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbase",
			Port:  217,
		},
	},
	"dbbrowse": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbbrowse",
			Port:  47557,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbbrowse",
			Port:  47557,
		},
	},
	"dbcontrol-oms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbcontrol-oms",
			Port:  1158,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbcontrol-oms",
			Port:  1158,
		},
	},
	"dbcontrol_agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbcontrol_agent",
			Port:  3938,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbcontrol_agent",
			Port:  3938,
		},
	},
	"dbdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbdb",
			Port:  6104,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbdb",
			Port:  6104,
		},
	},
	"dberegister": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dberegister",
			Port:  1479,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dberegister",
			Port:  1479,
		},
	},
	"dbisamserver1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbisamserver1",
			Port:  12005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbisamserver1",
			Port:  12005,
		},
	},
	"dbisamserver2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbisamserver2",
			Port:  12006,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbisamserver2",
			Port:  12006,
		},
	},
	"dbm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbm",
			Port:  2345,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbm",
			Port:  2345,
		},
	},
	"dbref": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbref",
			Port:  2365,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbref",
			Port:  2365,
		},
	},
	"dbreporter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbreporter",
			Port:  1379,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbreporter",
			Port:  1379,
		},
	},
	"dbsa-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbsa-lm",
			Port:  1407,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbsa-lm",
			Port:  1407,
		},
	},
	"dbstar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbstar",
			Port:  1415,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dbstar",
			Port:  1415,
		},
	},
	"dbsyncarbiter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dbsyncarbiter",
			Port:  4953,
		},
	},
	"dc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dc",
			Port:  2001,
		},
	},
	"dca": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dca",
			Port:  1456,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dca",
			Port:  1456,
		},
	},
	"dcap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dcap",
			Port:  22125,
		},
	},
	"dccm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dccm",
			Port:  5679,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dccm",
			Port:  5679,
		},
	},
	"dccp-udp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "dccp-udp",
			Port:  6511,
		},
	},
	"dcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dcp",
			Port:  93,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dcp",
			Port:  93,
		},
	},
	"dcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dcs",
			Port:  1367,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dcs",
			Port:  1367,
		},
	},
	"dcs-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dcs-config",
			Port:  3988,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dcs-config",
			Port:  3988,
		},
	},
	"dcsl-backup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dcsl-backup",
			Port:  11202,
		},
	},
	"dcsoftware": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dcsoftware",
			Port:  3793,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dcsoftware",
			Port:  3793,
		},
	},
	"dctp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dctp",
			Port:  675,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dctp",
			Port:  675,
		},
	},
	"dcutility": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dcutility",
			Port:  1044,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dcutility",
			Port:  1044,
		},
	},
	"dddp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dddp",
			Port:  9131,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dddp",
			Port:  9131,
		},
	},
	"ddgn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddgn",
			Port:  4167,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ddgn",
			Port:  4167,
		},
	},
	"ddi-tcp-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddi-tcp-1",
			Port:  8888,
		},
	},
	"ddi-tcp-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddi-tcp-2",
			Port:  8889,
		},
	},
	"ddi-tcp-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddi-tcp-3",
			Port:  8890,
		},
	},
	"ddi-tcp-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddi-tcp-4",
			Port:  8891,
		},
	},
	"ddi-tcp-5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddi-tcp-5",
			Port:  8892,
		},
	},
	"ddi-tcp-6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddi-tcp-6",
			Port:  8893,
		},
	},
	"ddi-tcp-7": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddi-tcp-7",
			Port:  8894,
		},
	},
	"ddi-udp-1": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ddi-udp-1",
			Port:  8888,
		},
	},
	"ddi-udp-2": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ddi-udp-2",
			Port:  8889,
		},
	},
	"ddi-udp-3": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ddi-udp-3",
			Port:  8890,
		},
	},
	"ddi-udp-4": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ddi-udp-4",
			Port:  8891,
		},
	},
	"ddi-udp-5": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ddi-udp-5",
			Port:  8892,
		},
	},
	"ddi-udp-6": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ddi-udp-6",
			Port:  8893,
		},
	},
	"ddi-udp-7": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ddi-udp-7",
			Port:  8894,
		},
	},
	"ddm-dfm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddm-dfm",
			Port:  447,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ddm-dfm",
			Port:  447,
		},
	},
	"ddm-rdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddm-rdb",
			Port:  446,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ddm-rdb",
			Port:  446,
		},
	},
	"ddm-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddm-ssl",
			Port:  448,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ddm-ssl",
			Port:  448,
		},
	},
	"ddns-v3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddns-v3",
			Port:  2164,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ddns-v3",
			Port:  2164,
		},
	},
	"ddrepl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddrepl",
			Port:  4126,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ddrepl",
			Port:  4126,
		},
	},
	"ddt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ddt",
			Port:  1052,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ddt",
			Port:  1052,
		},
	},
	"de-cache-query": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "de-cache-query",
			Port:  1255,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "de-cache-query",
			Port:  1255,
		},
	},
	"de-noc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "de-noc",
			Port:  1254,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "de-noc",
			Port:  1254,
		},
	},
	"de-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "de-server",
			Port:  1256,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "de-server",
			Port:  1256,
		},
	},
	"de-spot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "de-spot",
			Port:  2753,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "de-spot",
			Port:  2753,
		},
	},
	"dec-mbadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dec-mbadmin",
			Port:  1655,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dec-mbadmin",
			Port:  1655,
		},
	},
	"dec-mbadmin-h": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dec-mbadmin-h",
			Port:  1656,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dec-mbadmin-h",
			Port:  1656,
		},
	},
	"dec-notes": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dec-notes",
			Port:  3333,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dec-notes",
			Port:  3333,
		},
	},
	"dec_dlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dec_dlm",
			Port:  625,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dec_dlm",
			Port:  625,
		},
	},
	"decap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "decap",
			Port:  403,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "decap",
			Port:  403,
		},
	},
	"decauth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "decauth",
			Port:  316,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "decauth",
			Port:  316,
		},
	},
	"decbsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "decbsrv",
			Port:  579,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "decbsrv",
			Port:  579,
		},
	},
	"decladebug": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "decladebug",
			Port:  410,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "decladebug",
			Port:  410,
		},
	},
	"dectalk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dectalk",
			Port:  2007,
		},
	},
	"decvms-sysmgt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "decvms-sysmgt",
			Port:  441,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "decvms-sysmgt",
			Port:  441,
		},
	},
	"dei-icda": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dei-icda",
			Port:  618,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dei-icda",
			Port:  618,
		},
	},
	"delibo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "delibo",
			Port:  2562,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "delibo",
			Port:  2562,
		},
	},
	"dell-eql-asm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dell-eql-asm",
			Port:  7569,
		},
	},
	"dell-rm-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dell-rm-port",
			Port:  3668,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dell-rm-port",
			Port:  3668,
		},
	},
	"dellpwrappks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dellpwrappks",
			Port:  1266,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dellpwrappks",
			Port:  1266,
		},
	},
	"dellwebadmin-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dellwebadmin-1",
			Port:  1278,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dellwebadmin-1",
			Port:  1278,
		},
	},
	"dellwebadmin-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dellwebadmin-2",
			Port:  1279,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dellwebadmin-2",
			Port:  1279,
		},
	},
	"delos-dms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "delos-dms",
			Port:  3714,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "delos-dms",
			Port:  3714,
		},
	},
	"delta-mcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "delta-mcp",
			Port:  1324,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "delta-mcp",
			Port:  1324,
		},
	},
	"denali-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "denali-server",
			Port:  3444,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "denali-server",
			Port:  3444,
		},
	},
	"deos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "deos",
			Port:  76,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "deos",
			Port:  76,
		},
	},
	"derby-repli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "derby-repli",
			Port:  4851,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "derby-repli",
			Port:  4851,
		},
	},
	"descent3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "descent3",
			Port:  2092,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "descent3",
			Port:  2092,
		},
	},
	"deskshare": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "deskshare",
			Port:  1702,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "deskshare",
			Port:  1702,
		},
	},
	"desktop-dna": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "desktop-dna",
			Port:  2763,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "desktop-dna",
			Port:  2763,
		},
	},
	"deskview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "deskview",
			Port:  3298,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "deskview",
			Port:  3298,
		},
	},
	"devbasic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "devbasic",
			Port:  5426,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "devbasic",
			Port:  5426,
		},
	},
	"device": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "device",
			Port:  801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "device",
			Port:  801,
		},
	},
	"device2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "device2",
			Port:  2030,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "device2",
			Port:  2030,
		},
	},
	"devshr-nts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "devshr-nts",
			Port:  552,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "devshr-nts",
			Port:  552,
		},
	},
	"dey-keyneg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dey-keyneg",
			Port:  8750,
		},
	},
	"dey-sapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dey-sapi",
			Port:  4330,
		},
	},
	"dfn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dfn",
			Port:  1133,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dfn",
			Port:  1133,
		},
	},
	"dfoxserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dfoxserver",
			Port:  2960,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dfoxserver",
			Port:  2960,
		},
	},
	"dfserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dfserver",
			Port:  21554,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dfserver",
			Port:  21554,
		},
	},
	"dgi-serv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dgi-serv",
			Port:  33333,
		},
	},
	"dgpf-exchg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dgpf-exchg",
			Port:  6785,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dgpf-exchg",
			Port:  6785,
		},
	},
	"dhanalakshmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dhanalakshmi",
			Port:  34567,
		},
	},
	"dhcp-failover": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dhcp-failover",
			Port:  647,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dhcp-failover",
			Port:  647,
		},
	},
	"dhcp-failover2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dhcp-failover2",
			Port:  847,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dhcp-failover2",
			Port:  847,
		},
	},
	"dhcpv6-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dhcpv6-client",
			Port:  546,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dhcpv6-client",
			Port:  546,
		},
	},
	"dhcpv6-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dhcpv6-server",
			Port:  547,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dhcpv6-server",
			Port:  547,
		},
	},
	"dhct-alerts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dhct-alerts",
			Port:  4676,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dhct-alerts",
			Port:  4676,
		},
	},
	"dhct-status": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dhct-status",
			Port:  4675,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dhct-status",
			Port:  4675,
		},
	},
	"dhe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dhe",
			Port:  3252,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dhe",
			Port:  3252,
		},
	},
	"di-ase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "di-ase",
			Port:  3046,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "di-ase",
			Port:  3046,
		},
	},
	"di-drm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "di-drm",
			Port:  2226,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "di-drm",
			Port:  2226,
		},
	},
	"di-msg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "di-msg",
			Port:  2227,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "di-msg",
			Port:  2227,
		},
	},
	"di-traceware": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "di-traceware",
			Port:  3041,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "di-traceware",
			Port:  3041,
		},
	},
	"diagmond": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "diagmond",
			Port:  1508,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "diagmond",
			Port:  1508,
		},
	},
	"diagnose-proc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "diagnose-proc",
			Port:  6072,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "diagnose-proc",
			Port:  6072,
		},
	},
	"dialog-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dialog-port",
			Port:  2098,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dialog-port",
			Port:  2098,
		},
	},
	"dialogic-elmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dialogic-elmd",
			Port:  1945,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dialogic-elmd",
			Port:  1945,
		},
	},
	"dialpad-voice1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dialpad-voice1",
			Port:  2860,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dialpad-voice1",
			Port:  2860,
		},
	},
	"dialpad-voice2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dialpad-voice2",
			Port:  2861,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dialpad-voice2",
			Port:  2861,
		},
	},
	"diameter": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "diameter",
			Port:  3868,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "diameter",
			Port:  3868,
		},
	},
	"diameters": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "diameters",
			Port:  5868,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "diameters",
			Port:  5868,
		},
	},
	"diamondport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "diamondport",
			Port:  33331,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "diamondport",
			Port:  33331,
		},
	},
	"dic-aida": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dic-aida",
			Port:  1941,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dic-aida",
			Port:  1941,
		},
	},
	"dicom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dicom",
			Port:  11112,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dicom",
			Port:  11112,
		},
	},
	"dicom-iscl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dicom-iscl",
			Port:  2761,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dicom-iscl",
			Port:  2761,
		},
	},
	"dicom-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dicom-tls",
			Port:  2762,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dicom-tls",
			Port:  2762,
		},
	},
	"dict": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dict",
			Port:  2628,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dict",
			Port:  2628,
		},
	},
	"dict-lookup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dict-lookup",
			Port:  2289,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dict-lookup",
			Port:  2289,
		},
	},
	"dif-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dif-port",
			Port:  2251,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dif-port",
			Port:  2251,
		},
	},
	"digiman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "digiman",
			Port:  2362,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "digiman",
			Port:  2362,
		},
	},
	"digital-notary": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "digital-notary",
			Port:  1335,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "digital-notary",
			Port:  1335,
		},
	},
	"digital-vrc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "digital-vrc",
			Port:  466,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "digital-vrc",
			Port:  466,
		},
	},
	"digivote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "digivote",
			Port:  3223,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "digivote",
			Port:  3223,
		},
	},
	"direcpc-dll": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "direcpc-dll",
			Port:  1844,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "direcpc-dll",
			Port:  1844,
		},
	},
	"direcpc-si": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "direcpc-si",
			Port:  2464,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "direcpc-si",
			Port:  2464,
		},
	},
	"direcpc-video": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "direcpc-video",
			Port:  1825,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "direcpc-video",
			Port:  1825,
		},
	},
	"direct": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "direct",
			Port:  242,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "direct",
			Port:  242,
		},
	},
	"directnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directnet",
			Port:  3447,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directnet",
			Port:  3447,
		},
	},
	"directplay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directplay",
			Port:  2234,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directplay",
			Port:  2234,
		},
	},
	"directplay8": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directplay8",
			Port:  6073,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directplay8",
			Port:  6073,
		},
	},
	"directplaysrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directplaysrvr",
			Port:  47624,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directplaysrvr",
			Port:  47624,
		},
	},
	"directv-catlg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directv-catlg",
			Port:  3337,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directv-catlg",
			Port:  3337,
		},
	},
	"directv-soft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directv-soft",
			Port:  3335,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directv-soft",
			Port:  3335,
		},
	},
	"directv-tick": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directv-tick",
			Port:  3336,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directv-tick",
			Port:  3336,
		},
	},
	"directv-web": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directv-web",
			Port:  3334,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directv-web",
			Port:  3334,
		},
	},
	"directvdata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "directvdata",
			Port:  3287,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "directvdata",
			Port:  3287,
		},
	},
	"dirgis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dirgis",
			Port:  2496,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dirgis",
			Port:  2496,
		},
	},
	"discard": map[string]Service{
		"dccp": Service{
			Proto: "dccp",
			Name:  "discard",
			Port:  9,
		},
		"sctp": Service{
			Proto: "sctp",
			Name:  "discard",
			Port:  9,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "discard",
			Port:  9,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "discard",
			Port:  9,
		},
	},
	"disclose": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "disclose",
			Port:  667,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "disclose",
			Port:  667,
		},
	},
	"discovery-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "discovery-port",
			Port:  1925,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "discovery-port",
			Port:  1925,
		},
	},
	"discp-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "discp-client",
			Port:  2601,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "discp-client",
			Port:  2601,
		},
	},
	"discp-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "discp-server",
			Port:  2602,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "discp-server",
			Port:  2602,
		},
	},
	"display": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "display",
			Port:  7236,
		},
	},
	"dist-upgrade": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dist-upgrade",
			Port:  3624,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dist-upgrade",
			Port:  3624,
		},
	},
	"distcc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "distcc",
			Port:  3632,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "distcc",
			Port:  3632,
		},
	},
	"distinct": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "distinct",
			Port:  9999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "distinct",
			Port:  9999,
		},
	},
	"distinct32": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "distinct32",
			Port:  9998,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "distinct32",
			Port:  9998,
		},
	},
	"dixie": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dixie",
			Port:  96,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dixie",
			Port:  96,
		},
	},
	"dj-ice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dj-ice",
			Port:  5419,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dj-ice",
			Port:  5419,
		},
	},
	"dj-ilm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dj-ilm",
			Port:  3362,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dj-ilm",
			Port:  3362,
		},
	},
	"dka": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dka",
			Port:  1263,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dka",
			Port:  1263,
		},
	},
	"dkmessenger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dkmessenger",
			Port:  1177,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dkmessenger",
			Port:  1177,
		},
	},
	"dl_agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dl_agent",
			Port:  3876,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dl_agent",
			Port:  3876,
		},
	},
	"dlip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dlip",
			Port:  7201,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dlip",
			Port:  7201,
		},
	},
	"dlms-cosem": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dlms-cosem",
			Port:  4059,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dlms-cosem",
			Port:  4059,
		},
	},
	"dlpx-sp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dlpx-sp",
			Port:  8415,
		},
	},
	"dls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dls",
			Port:  197,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dls",
			Port:  197,
		},
	},
	"dls-mon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dls-mon",
			Port:  198,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dls-mon",
			Port:  198,
		},
	},
	"dls-monitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dls-monitor",
			Port:  2048,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dls-monitor",
			Port:  2048,
		},
	},
	"dlsrap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dlsrap",
			Port:  1973,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dlsrap",
			Port:  1973,
		},
	},
	"dlsrpn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dlsrpn",
			Port:  2065,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dlsrpn",
			Port:  2065,
		},
	},
	"dlswpn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dlswpn",
			Port:  2067,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dlswpn",
			Port:  2067,
		},
	},
	"dmaf-caster": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "dmaf-caster",
			Port:  3574,
		},
	},
	"dmaf-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dmaf-server",
			Port:  3574,
		},
	},
	"dmdocbroker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dmdocbroker",
			Port:  1489,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dmdocbroker",
			Port:  1489,
		},
	},
	"dmidi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dmidi",
			Port:  1199,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dmidi",
			Port:  1199,
		},
	},
	"dmod-workspace": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dmod-workspace",
			Port:  3199,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dmod-workspace",
			Port:  3199,
		},
	},
	"dmp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "dmp",
			Port:  5031,
		},
	},
	"dn6-nlm-aud": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dn6-nlm-aud",
			Port:  195,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dn6-nlm-aud",
			Port:  195,
		},
	},
	"dn6-smm-red": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dn6-smm-red",
			Port:  196,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dn6-smm-red",
			Port:  196,
		},
	},
	"dna": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dna",
			Port:  2287,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dna",
			Port:  2287,
		},
	},
	"dna-cml": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dna-cml",
			Port:  436,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dna-cml",
			Port:  436,
		},
	},
	"dnap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dnap",
			Port:  1172,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dnap",
			Port:  1172,
		},
	},
	"dnc-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dnc-port",
			Port:  3448,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dnc-port",
			Port:  3448,
		},
	},
	"dnox": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dnox",
			Port:  4022,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dnox",
			Port:  4022,
		},
	},
	"dnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dnp",
			Port:  20000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dnp",
			Port:  20000,
		},
	},
	"dnp-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dnp-sec",
			Port:  19999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dnp-sec",
			Port:  19999,
		},
	},
	"dns-llq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dns-llq",
			Port:  5352,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dns-llq",
			Port:  5352,
		},
	},
	"dns2go": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dns2go",
			Port:  1227,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dns2go",
			Port:  1227,
		},
	},
	"dnsix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dnsix",
			Port:  90,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dnsix",
			Port:  90,
		},
	},
	"dnx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dnx",
			Port:  3998,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dnx",
			Port:  3998,
		},
	},
	"doc-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "doc-server",
			Port:  7165,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "doc-server",
			Port:  7165,
		},
	},
	"doc1lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "doc1lm",
			Port:  3161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "doc1lm",
			Port:  3161,
		},
	},
	"docent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "docent",
			Port:  2151,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "docent",
			Port:  2151,
		},
	},
	"doceri-ctl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "doceri-ctl",
			Port:  7019,
		},
	},
	"doceri-view": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "doceri-view",
			Port:  7019,
		},
	},
	"docstor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "docstor",
			Port:  1488,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "docstor",
			Port:  1488,
		},
	},
	"documentum": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "documentum",
			Port:  10002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "documentum",
			Port:  10002,
		},
	},
	"documentum_s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "documentum_s",
			Port:  10003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "documentum_s",
			Port:  10003,
		},
	},
	"doglms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "doglms",
			Port:  6088,
		},
	},
	"doglms-notify": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "doglms-notify",
			Port:  6088,
		},
	},
	"doip-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "doip-data",
			Port:  13400,
		},
	},
	"doip-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "doip-disc",
			Port:  13400,
		},
	},
	"domain": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "domain",
			Port:  53,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "domain",
			Port:  53,
		},
	},
	"domaintime": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "domaintime",
			Port:  9909,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "domaintime",
			Port:  9909,
		},
	},
	"domiq": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "domiq",
			Port:  44544,
		},
	},
	"donnyworld": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "donnyworld",
			Port:  1821,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "donnyworld",
			Port:  1821,
		},
	},
	"dossier": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dossier",
			Port:  1175,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dossier",
			Port:  1175,
		},
	},
	"down": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "down",
			Port:  2022,
		},
	},
	"downtools": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "downtools",
			Port:  5245,
		},
	},
	"downtools-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "downtools-disc",
			Port:  5245,
		},
	},
	"dpap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpap",
			Port:  8770,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpap",
			Port:  8770,
		},
	},
	"dpcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpcp",
			Port:  4099,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpcp",
			Port:  4099,
		},
	},
	"dpi-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpi-proxy",
			Port:  1795,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpi-proxy",
			Port:  1795,
		},
	},
	"dpkeyserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpkeyserv",
			Port:  1780,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpkeyserv",
			Port:  1780,
		},
	},
	"dpm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpm",
			Port:  5718,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpm",
			Port:  5718,
		},
	},
	"dpm-acm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpm-acm",
			Port:  6075,
		},
	},
	"dpm-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpm-agent",
			Port:  5719,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpm-agent",
			Port:  5719,
		},
	},
	"dproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dproxy",
			Port:  1296,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dproxy",
			Port:  1296,
		},
	},
	"dpserve": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpserve",
			Port:  7020,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpserve",
			Port:  7020,
		},
	},
	"dpserveadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpserveadmin",
			Port:  7021,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpserveadmin",
			Port:  7021,
		},
	},
	"dpsi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dpsi",
			Port:  315,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dpsi",
			Port:  315,
		},
	},
	"dragonfly": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dragonfly",
			Port:  8913,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dragonfly",
			Port:  8913,
		},
	},
	"drip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "drip",
			Port:  3949,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "drip",
			Port:  3949,
		},
	},
	"driveappserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "driveappserver",
			Port:  1930,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "driveappserver",
			Port:  1930,
		},
	},
	"drizzle": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "drizzle",
			Port:  4427,
		},
	},
	"drm-production": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "drm-production",
			Port:  7171,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "drm-production",
			Port:  7171,
		},
	},
	"drmsfsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "drmsfsd",
			Port:  4098,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "drmsfsd",
			Port:  4098,
		},
	},
	"drmsmc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "drmsmc",
			Port:  1878,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "drmsmc",
			Port:  1878,
		},
	},
	"drp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "drp",
			Port:  1974,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "drp",
			Port:  1974,
		},
	},
	"drwcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "drwcs",
			Port:  2193,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "drwcs",
			Port:  2193,
		},
	},
	"ds-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ds-admin",
			Port:  4404,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ds-admin",
			Port:  4404,
		},
	},
	"ds-clnt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ds-clnt",
			Port:  4402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ds-clnt",
			Port:  4402,
		},
	},
	"ds-mail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ds-mail",
			Port:  4405,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ds-mail",
			Port:  4405,
		},
	},
	"ds-slp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ds-slp",
			Port:  4406,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ds-slp",
			Port:  4406,
		},
	},
	"ds-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ds-srv",
			Port:  4400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ds-srv",
			Port:  4400,
		},
	},
	"ds-srvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ds-srvr",
			Port:  4401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ds-srvr",
			Port:  4401,
		},
	},
	"ds-user": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ds-user",
			Port:  4403,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ds-user",
			Port:  4403,
		},
	},
	"dsETOS": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsETOS",
			Port:  378,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsETOS",
			Port:  378,
		},
	},
	"dsatp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsatp",
			Port:  2111,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsatp",
			Port:  2111,
		},
	},
	"dsc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsc",
			Port:  3390,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsc",
			Port:  3390,
		},
	},
	"dsdn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsdn",
			Port:  1292,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsdn",
			Port:  1292,
		},
	},
	"dserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dserver",
			Port:  4309,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dserver",
			Port:  4309,
		},
	},
	"dsf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsf",
			Port:  555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsf",
			Port:  555,
		},
	},
	"dsfgw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsfgw",
			Port:  438,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsfgw",
			Port:  438,
		},
	},
	"dslremote-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dslremote-mgmt",
			Port:  2420,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dslremote-mgmt",
			Port:  2420,
		},
	},
	"dsm-scm-target": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsm-scm-target",
			Port:  9987,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsm-scm-target",
			Port:  9987,
		},
	},
	"dsmcc-ccp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsmcc-ccp",
			Port:  13822,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsmcc-ccp",
			Port:  13822,
		},
	},
	"dsmcc-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsmcc-config",
			Port:  13818,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsmcc-config",
			Port:  13818,
		},
	},
	"dsmcc-download": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsmcc-download",
			Port:  13821,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsmcc-download",
			Port:  13821,
		},
	},
	"dsmcc-passthru": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsmcc-passthru",
			Port:  13820,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsmcc-passthru",
			Port:  13820,
		},
	},
	"dsmcc-session": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsmcc-session",
			Port:  13819,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsmcc-session",
			Port:  13819,
		},
	},
	"dsmeter_iatc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsmeter_iatc",
			Port:  4060,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsmeter_iatc",
			Port:  4060,
		},
	},
	"dsmipv6": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "dsmipv6",
			Port:  4191,
		},
	},
	"dsom-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsom-server",
			Port:  3053,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsom-server",
			Port:  3053,
		},
	},
	"dsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsp",
			Port:  33,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsp",
			Port:  33,
		},
	},
	"dsp3270": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsp3270",
			Port:  246,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsp3270",
			Port:  246,
		},
	},
	"dssiapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dssiapi",
			Port:  1265,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dssiapi",
			Port:  1265,
		},
	},
	"dsx-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsx-agent",
			Port:  3685,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dsx-agent",
			Port:  3685,
		},
	},
	"dsx_monitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dsx_monitor",
			Port:  31685,
		},
	},
	"dt-mgmtsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dt-mgmtsvc",
			Port:  6325,
		},
	},
	"dt-vra": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dt-vra",
			Port:  6326,
		},
	},
	"dta-systems": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dta-systems",
			Port:  13929,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dta-systems",
			Port:  13929,
		},
	},
	"dtag-ste-sb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtag-ste-sb",
			Port:  352,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtag-ste-sb",
			Port:  352,
		},
	},
	"dtk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtk",
			Port:  365,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtk",
			Port:  365,
		},
	},
	"dtn-bundle-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtn-bundle-tcp",
			Port:  4556,
		},
	},
	"dtn-bundle-udp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "dtn-bundle-udp",
			Port:  4556,
		},
	},
	"dtn1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtn1",
			Port:  2445,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtn1",
			Port:  2445,
		},
	},
	"dtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtp",
			Port:  3663,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtp",
			Port:  3663,
		},
	},
	"dtp-dia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtp-dia",
			Port:  3489,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtp-dia",
			Port:  3489,
		},
	},
	"dtp-net": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "dtp-net",
			Port:  8732,
		},
	},
	"dtpt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtpt",
			Port:  5721,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtpt",
			Port:  5721,
		},
	},
	"dts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dts",
			Port:  2594,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dts",
			Port:  2594,
		},
	},
	"dtserver-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtserver-port",
			Port:  4028,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtserver-port",
			Port:  4028,
		},
	},
	"dtspcd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtspcd",
			Port:  6112,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtspcd",
			Port:  6112,
		},
	},
	"dtv-chan-req": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dtv-chan-req",
			Port:  2253,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dtv-chan-req",
			Port:  2253,
		},
	},
	"dvapps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dvapps",
			Port:  3831,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dvapps",
			Port:  3831,
		},
	},
	"dvbservdsc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dvbservdsc",
			Port:  3937,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dvbservdsc",
			Port:  3937,
		},
	},
	"dvcprov-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dvcprov-port",
			Port:  3776,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dvcprov-port",
			Port:  3776,
		},
	},
	"dvl-activemail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dvl-activemail",
			Port:  1396,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dvl-activemail",
			Port:  1396,
		},
	},
	"dvr-esm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dvr-esm",
			Port:  2804,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dvr-esm",
			Port:  2804,
		},
	},
	"dvt-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dvt-data",
			Port:  3247,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dvt-data",
			Port:  3247,
		},
	},
	"dvt-system": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dvt-system",
			Port:  3246,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dvt-system",
			Port:  3246,
		},
	},
	"dwf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dwf",
			Port:  1450,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dwf",
			Port:  1450,
		},
	},
	"dwmsgserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dwmsgserver",
			Port:  3228,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dwmsgserver",
			Port:  3228,
		},
	},
	"dwnmshttp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dwnmshttp",
			Port:  3227,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dwnmshttp",
			Port:  3227,
		},
	},
	"dwr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dwr",
			Port:  644,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dwr",
			Port:  644,
		},
	},
	"dx-instrument": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dx-instrument",
			Port:  1325,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dx-instrument",
			Port:  1325,
		},
	},
	"dxadmind": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dxadmind",
			Port:  1958,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dxadmind",
			Port:  1958,
		},
	},
	"dxmessagebase1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dxmessagebase1",
			Port:  2874,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dxmessagebase1",
			Port:  2874,
		},
	},
	"dxmessagebase2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dxmessagebase2",
			Port:  2875,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dxmessagebase2",
			Port:  2875,
		},
	},
	"dxspider": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dxspider",
			Port:  8873,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dxspider",
			Port:  8873,
		},
	},
	"dyn-site": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dyn-site",
			Port:  3932,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dyn-site",
			Port:  3932,
		},
	},
	"dyna-access": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dyna-access",
			Port:  3310,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dyna-access",
			Port:  3310,
		},
	},
	"dyna-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dyna-lm",
			Port:  3395,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dyna-lm",
			Port:  3395,
		},
	},
	"dynamid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dynamid",
			Port:  9002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dynamid",
			Port:  9002,
		},
	},
	"dyniplookup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dyniplookup",
			Port:  3295,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dyniplookup",
			Port:  3295,
		},
	},
	"dzdaemon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dzdaemon",
			Port:  3866,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dzdaemon",
			Port:  3866,
		},
	},
	"dzoglserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "dzoglserver",
			Port:  3867,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "dzoglserver",
			Port:  3867,
		},
	},
	"e-builder": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "e-builder",
			Port:  4121,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "e-builder",
			Port:  4121,
		},
	},
	"e-design-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "e-design-net",
			Port:  6702,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "e-design-net",
			Port:  6702,
		},
	},
	"e-design-web": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "e-design-web",
			Port:  6703,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "e-design-web",
			Port:  6703,
		},
	},
	"e-dpnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "e-dpnet",
			Port:  2036,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "e-dpnet",
			Port:  2036,
		},
	},
	"e-mdu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "e-mdu",
			Port:  3727,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "e-mdu",
			Port:  3727,
		},
	},
	"e-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "e-net",
			Port:  3286,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "e-net",
			Port:  3286,
		},
	},
	"e-woa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "e-woa",
			Port:  3728,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "e-woa",
			Port:  3728,
		},
	},
	"e3consultants": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "e3consultants",
			Port:  3157,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "e3consultants",
			Port:  3157,
		},
	},
	"ea": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ea",
			Port:  17729,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ea",
			Port:  17729,
		},
	},
	"ea1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ea1",
			Port:  1791,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ea1",
			Port:  1791,
		},
	},
	"eapsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eapsp",
			Port:  2291,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eapsp",
			Port:  2291,
		},
	},
	"easy-soft-mux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "easy-soft-mux",
			Port:  2168,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "easy-soft-mux",
			Port:  2168,
		},
	},
	"eba": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eba",
			Port:  45678,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eba",
			Port:  45678,
		},
	},
	"ebinsite": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ebinsite",
			Port:  2651,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ebinsite",
			Port:  2651,
		},
	},
	"echo": map[string]Service{
		"ddp": Service{
			Proto: "ddp",
			Name:  "echo",
			Port:  4,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "echo",
			Port:  7,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "echo",
			Port:  7,
		},
	},
	"echonet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "echonet",
			Port:  3610,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "echonet",
			Port:  3610,
		},
	},
	"ecmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecmp",
			Port:  6160,
		},
	},
	"ecmp-data": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ecmp-data",
			Port:  6160,
		},
	},
	"ecmport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecmport",
			Port:  3524,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ecmport",
			Port:  3524,
		},
	},
	"ecnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecnp",
			Port:  2858,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ecnp",
			Port:  2858,
		},
	},
	"ecolor-imager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecolor-imager",
			Port:  3263,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ecolor-imager",
			Port:  3263,
		},
	},
	"ecomm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecomm",
			Port:  3477,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ecomm",
			Port:  3477,
		},
	},
	"ecovisiong6-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecovisiong6-1",
			Port:  2896,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ecovisiong6-1",
			Port:  2896,
		},
	},
	"ecp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecp",
			Port:  3134,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ecp",
			Port:  3134,
		},
	},
	"ecsqdmn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecsqdmn",
			Port:  1882,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ecsqdmn",
			Port:  1882,
		},
	},
	"ecwcfg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ecwcfg",
			Port:  2263,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ecwcfg",
			Port:  2263,
		},
	},
	"edb-server1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edb-server1",
			Port:  1635,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edb-server1",
			Port:  1635,
		},
	},
	"edb-server2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edb-server2",
			Port:  3711,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edb-server2",
			Port:  3711,
		},
	},
	"edbsrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edbsrvr",
			Port:  12010,
		},
	},
	"editbench": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "editbench",
			Port:  1350,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "editbench",
			Port:  1350,
		},
	},
	"edix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edix",
			Port:  3123,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edix",
			Port:  3123,
		},
	},
	"edm-adm-notify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edm-adm-notify",
			Port:  3463,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edm-adm-notify",
			Port:  3463,
		},
	},
	"edm-manager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edm-manager",
			Port:  3460,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edm-manager",
			Port:  3460,
		},
	},
	"edm-mgr-cntrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edm-mgr-cntrl",
			Port:  3465,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edm-mgr-cntrl",
			Port:  3465,
		},
	},
	"edm-mgr-sync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edm-mgr-sync",
			Port:  3464,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edm-mgr-sync",
			Port:  3464,
		},
	},
	"edm-stager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edm-stager",
			Port:  3461,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edm-stager",
			Port:  3461,
		},
	},
	"edm-std-notify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edm-std-notify",
			Port:  3462,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edm-std-notify",
			Port:  3462,
		},
	},
	"edtools": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "edtools",
			Port:  1142,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "edtools",
			Port:  1142,
		},
	},
	"eenet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eenet",
			Port:  5234,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eenet",
			Port:  5234,
		},
	},
	"efb-aci": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "efb-aci",
			Port:  6159,
		},
	},
	"efcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "efcp",
			Port:  3671,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "efcp",
			Port:  3671,
		},
	},
	"efi-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "efi-lm",
			Port:  3392,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "efi-lm",
			Port:  3392,
		},
	},
	"efi-mg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "efi-mg",
			Port:  2224,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "efi-mg",
			Port:  2224,
		},
	},
	"efidiningport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "efidiningport",
			Port:  2553,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "efidiningport",
			Port:  2553,
		},
	},
	"eforward": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eforward",
			Port:  2181,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eforward",
			Port:  2181,
		},
	},
	"efs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "efs",
			Port:  520,
		},
	},
	"egptlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "egptlm",
			Port:  3328,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "egptlm",
			Port:  3328,
		},
	},
	"egs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "egs",
			Port:  1926,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "egs",
			Port:  1926,
		},
	},
	"ehome-ms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ehome-ms",
			Port:  2228,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ehome-ms",
			Port:  2228,
		},
	},
	"ehp-backup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ehp-backup",
			Port:  3638,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ehp-backup",
			Port:  3638,
		},
	},
	"ehs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ehs",
			Port:  4535,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ehs",
			Port:  4535,
		},
	},
	"ehs-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ehs-ssl",
			Port:  4536,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ehs-ssl",
			Port:  4536,
		},
	},
	"ehtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ehtp",
			Port:  1295,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ehtp",
			Port:  1295,
		},
	},
	"eicon-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eicon-server",
			Port:  1438,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eicon-server",
			Port:  1438,
		},
	},
	"eicon-slp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eicon-slp",
			Port:  1440,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eicon-slp",
			Port:  1440,
		},
	},
	"eicon-x25": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eicon-x25",
			Port:  1439,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eicon-x25",
			Port:  1439,
		},
	},
	"eims-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eims-admin",
			Port:  4199,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eims-admin",
			Port:  4199,
		},
	},
	"eis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eis",
			Port:  3982,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eis",
			Port:  3982,
		},
	},
	"eisp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eisp",
			Port:  3983,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eisp",
			Port:  3983,
		},
	},
	"eisport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eisport",
			Port:  3525,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eisport",
			Port:  3525,
		},
	},
	"eklogin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eklogin",
			Port:  2105,
		},
	},
	"elad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elad",
			Port:  1893,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elad",
			Port:  1893,
		},
	},
	"elan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elan",
			Port:  1378,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elan",
			Port:  1378,
		},
	},
	"elanlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elanlm",
			Port:  4346,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elanlm",
			Port:  4346,
		},
	},
	"elatelink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elatelink",
			Port:  2124,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elatelink",
			Port:  2124,
		},
	},
	"elcn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elcn",
			Port:  7101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elcn",
			Port:  7101,
		},
	},
	"elcsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elcsd",
			Port:  704,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elcsd",
			Port:  704,
		},
	},
	"elektron-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elektron-admin",
			Port:  5398,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elektron-admin",
			Port:  5398,
		},
	},
	"elfiq-repl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elfiq-repl",
			Port:  1148,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elfiq-repl",
			Port:  1148,
		},
	},
	"eli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eli",
			Port:  2087,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eli",
			Port:  2087,
		},
	},
	"elipse-rec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elipse-rec",
			Port:  6515,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elipse-rec",
			Port:  6515,
		},
	},
	"ellpack": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ellpack",
			Port:  2025,
		},
	},
	"elm-momentum": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elm-momentum",
			Port:  1914,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elm-momentum",
			Port:  1914,
		},
	},
	"elpro_tunnel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elpro_tunnel",
			Port:  4370,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elpro_tunnel",
			Port:  4370,
		},
	},
	"els": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "els",
			Port:  1315,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "els",
			Port:  1315,
		},
	},
	"elvin_client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elvin_client",
			Port:  2917,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elvin_client",
			Port:  2917,
		},
	},
	"elvin_server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elvin_server",
			Port:  2916,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elvin_server",
			Port:  2916,
		},
	},
	"elxmgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "elxmgmt",
			Port:  23333,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "elxmgmt",
			Port:  23333,
		},
	},
	"em7-secom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "em7-secom",
			Port:  7700,
		},
	},
	"ema-sent-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ema-sent-lm",
			Port:  2526,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ema-sent-lm",
			Port:  2526,
		},
	},
	"emb-proj-cmd": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "emb-proj-cmd",
			Port:  5116,
		},
	},
	"embl-ndt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "embl-ndt",
			Port:  394,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "embl-ndt",
			Port:  394,
		},
	},
	"embrace-dp-c": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "embrace-dp-c",
			Port:  3198,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "embrace-dp-c",
			Port:  3198,
		},
	},
	"embrace-dp-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "embrace-dp-s",
			Port:  3197,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "embrace-dp-s",
			Port:  3197,
		},
	},
	"emc-gateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emc-gateway",
			Port:  1273,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emc-gateway",
			Port:  1273,
		},
	},
	"emc-pp-mgmtsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emc-pp-mgmtsvc",
			Port:  9083,
		},
	},
	"emc-vcas-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emc-vcas-tcp",
			Port:  13218,
		},
	},
	"emc-vcas-udp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "emc-vcas-udp",
			Port:  13218,
		},
	},
	"emcads": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emcads",
			Port:  3945,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emcads",
			Port:  3945,
		},
	},
	"emce": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "emce",
			Port:  2004,
		},
	},
	"emcrmirccd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emcrmirccd",
			Port:  10004,
		},
	},
	"emcrmird": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emcrmird",
			Port:  10005,
		},
	},
	"emcsymapiport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emcsymapiport",
			Port:  2707,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emcsymapiport",
			Port:  2707,
		},
	},
	"emfis-cntl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emfis-cntl",
			Port:  141,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emfis-cntl",
			Port:  141,
		},
	},
	"emfis-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emfis-data",
			Port:  140,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emfis-data",
			Port:  140,
		},
	},
	"emgmsg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emgmsg",
			Port:  6656,
		},
	},
	"emp-server1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emp-server1",
			Port:  6321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emp-server1",
			Port:  6321,
		},
	},
	"emp-server2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emp-server2",
			Port:  6322,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emp-server2",
			Port:  6322,
		},
	},
	"emperion": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emperion",
			Port:  1282,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emperion",
			Port:  1282,
		},
	},
	"empire-empuma": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "empire-empuma",
			Port:  1691,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "empire-empuma",
			Port:  1691,
		},
	},
	"empowerid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "empowerid",
			Port:  7080,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "empowerid",
			Port:  7080,
		},
	},
	"emprise-lls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emprise-lls",
			Port:  3585,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emprise-lls",
			Port:  3585,
		},
	},
	"emprise-lsc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emprise-lsc",
			Port:  3586,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emprise-lsc",
			Port:  3586,
		},
	},
	"ems": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ems",
			Port:  4664,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ems",
			Port:  4664,
		},
	},
	"emsd-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emsd-port",
			Port:  1928,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emsd-port",
			Port:  1928,
		},
	},
	"emwavemsg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emwavemsg",
			Port:  20480,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emwavemsg",
			Port:  20480,
		},
	},
	"emwin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "emwin",
			Port:  2211,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "emwin",
			Port:  2211,
		},
	},
	"enc-eps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enc-eps",
			Port:  3567,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "enc-eps",
			Port:  3567,
		},
	},
	"enc-eps-mc-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enc-eps-mc-sec",
			Port:  5567,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "enc-eps-mc-sec",
			Port:  5567,
		},
	},
	"enc-tunel-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enc-tunel-sec",
			Port:  3568,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "enc-tunel-sec",
			Port:  3568,
		},
	},
	"enc-tunnel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enc-tunnel",
			Port:  8567,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "enc-tunnel",
			Port:  8567,
		},
	},
	"encore": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "encore",
			Port:  1740,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "encore",
			Port:  1740,
		},
	},
	"encrypted-llrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "encrypted-llrp",
			Port:  5085,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "encrypted-llrp",
			Port:  5085,
		},
	},
	"encrypted_admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "encrypted_admin",
			Port:  1138,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "encrypted_admin",
			Port:  1138,
		},
	},
	"enfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enfs",
			Port:  5233,
		},
	},
	"enl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enl",
			Port:  1804,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "enl",
			Port:  1804,
		},
	},
	"enl-name": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enl-name",
			Port:  1805,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "enl-name",
			Port:  1805,
		},
	},
	"enpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enpc",
			Port:  3289,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "enpc",
			Port:  3289,
		},
	},
	"enpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "enpp",
			Port:  2968,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "enpp",
			Port:  2968,
		},
	},
	"enrp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "enrp",
			Port:  9901,
		},
	},
	"enrp-sctp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "enrp-sctp",
			Port:  9901,
		},
	},
	"enrp-sctp-tls": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "enrp-sctp-tls",
			Port:  9902,
		},
	},
	"ent-engine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ent-engine",
			Port:  3665,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ent-engine",
			Port:  3665,
		},
	},
	"entexthigh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entexthigh",
			Port:  12002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entexthigh",
			Port:  12002,
		},
	},
	"entextlow": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entextlow",
			Port:  12004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entextlow",
			Port:  12004,
		},
	},
	"entextmed": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entextmed",
			Port:  12003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entextmed",
			Port:  12003,
		},
	},
	"entextnetwk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entextnetwk",
			Port:  12001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entextnetwk",
			Port:  12001,
		},
	},
	"entextxid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entextxid",
			Port:  12000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entextxid",
			Port:  12000,
		},
	},
	"entomb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entomb",
			Port:  775,
		},
	},
	"entp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entp",
			Port:  1865,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entp",
			Port:  1865,
		},
	},
	"entrust-aaas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entrust-aaas",
			Port:  680,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entrust-aaas",
			Port:  680,
		},
	},
	"entrust-aams": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entrust-aams",
			Port:  681,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entrust-aams",
			Port:  681,
		},
	},
	"entrust-ash": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entrust-ash",
			Port:  710,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entrust-ash",
			Port:  710,
		},
	},
	"entrust-kmsh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entrust-kmsh",
			Port:  709,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entrust-kmsh",
			Port:  709,
		},
	},
	"entrust-sps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entrust-sps",
			Port:  640,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entrust-sps",
			Port:  640,
		},
	},
	"entrusttime": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "entrusttime",
			Port:  309,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "entrusttime",
			Port:  309,
		},
	},
	"eor-game": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "eor-game",
			Port:  8149,
		},
	},
	"eoss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eoss",
			Port:  1210,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eoss",
			Port:  1210,
		},
	},
	"ep-nsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ep-nsp",
			Port:  3621,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ep-nsp",
			Port:  3621,
		},
	},
	"ep-pcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ep-pcp",
			Port:  3620,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ep-pcp",
			Port:  3620,
		},
	},
	"epc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "epc",
			Port:  1267,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "epc",
			Port:  1267,
		},
	},
	"epicon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "epicon",
			Port:  2912,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "epicon",
			Port:  2912,
		},
	},
	"epl-slp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "epl-slp",
			Port:  3819,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "epl-slp",
			Port:  3819,
		},
	},
	"epmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "epmap",
			Port:  135,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "epmap",
			Port:  135,
		},
	},
	"epmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "epmd",
			Port:  4369,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "epmd",
			Port:  4369,
		},
	},
	"epncdp2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "epncdp2",
			Port:  3259,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "epncdp2",
			Port:  3259,
		},
	},
	"epnsdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "epnsdp",
			Port:  2051,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "epnsdp",
			Port:  2051,
		},
	},
	"eportcomm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eportcomm",
			Port:  4666,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eportcomm",
			Port:  4666,
		},
	},
	"eportcommdata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eportcommdata",
			Port:  4669,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eportcommdata",
			Port:  4669,
		},
	},
	"epp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "epp",
			Port:  700,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "epp",
			Port:  700,
		},
	},
	"eppc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eppc",
			Port:  3031,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eppc",
			Port:  3031,
		},
	},
	"ept-machine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ept-machine",
			Port:  3628,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ept-machine",
			Port:  3628,
		},
	},
	"eq-office-4940": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eq-office-4940",
			Port:  4940,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eq-office-4940",
			Port:  4940,
		},
	},
	"eq-office-4941": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eq-office-4941",
			Port:  4941,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eq-office-4941",
			Port:  4941,
		},
	},
	"eq-office-4942": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eq-office-4942",
			Port:  4942,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eq-office-4942",
			Port:  4942,
		},
	},
	"eq3-config": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "eq3-config",
			Port:  43439,
		},
	},
	"eq3-update": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eq3-update",
			Port:  43439,
		},
	},
	"equationbuilder": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "equationbuilder",
			Port:  1351,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "equationbuilder",
			Port:  1351,
		},
	},
	"ergolight": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ergolight",
			Port:  2109,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ergolight",
			Port:  2109,
		},
	},
	"eristwoguns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eristwoguns",
			Port:  2650,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eristwoguns",
			Port:  2650,
		},
	},
	"erp-scale": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "erp-scale",
			Port:  5135,
		},
	},
	"erpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "erpc",
			Port:  121,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "erpc",
			Port:  121,
		},
	},
	"erunbook_agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "erunbook_agent",
			Port:  9616,
		},
	},
	"erunbook_server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "erunbook_server",
			Port:  9617,
		},
	},
	"es-elmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "es-elmd",
			Port:  1822,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "es-elmd",
			Port:  1822,
		},
	},
	"esbroker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esbroker",
			Port:  1342,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esbroker",
			Port:  1342,
		},
	},
	"escp-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "escp-ip",
			Port:  621,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "escp-ip",
			Port:  621,
		},
	},
	"escvpnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "escvpnet",
			Port:  3629,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "escvpnet",
			Port:  3629,
		},
	},
	"eserver-pap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eserver-pap",
			Port:  3666,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eserver-pap",
			Port:  3666,
		},
	},
	"esimport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esimport",
			Port:  3564,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esimport",
			Port:  3564,
		},
	},
	"esinstall": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esinstall",
			Port:  5599,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esinstall",
			Port:  5599,
		},
	},
	"esip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esip",
			Port:  2950,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esip",
			Port:  2950,
		},
	},
	"esl-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esl-lm",
			Port:  1455,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esl-lm",
			Port:  1455,
		},
	},
	"esmagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esmagent",
			Port:  5601,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esmagent",
			Port:  5601,
		},
	},
	"esmmanager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esmmanager",
			Port:  5600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esmmanager",
			Port:  5600,
		},
	},
	"esnm-zoning": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esnm-zoning",
			Port:  4023,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esnm-zoning",
			Port:  4023,
		},
	},
	"esp-encap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esp-encap",
			Port:  2797,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esp-encap",
			Port:  2797,
		},
	},
	"esp-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esp-lm",
			Port:  3383,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esp-lm",
			Port:  3383,
		},
	},
	"espeech": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "espeech",
			Port:  8416,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "espeech",
			Port:  8416,
		},
	},
	"espeech-rtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "espeech-rtp",
			Port:  8417,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "espeech-rtp",
			Port:  8417,
		},
	},
	"esps-portal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esps-portal",
			Port:  2867,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esps-portal",
			Port:  2867,
		},
	},
	"esri_sde": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esri_sde",
			Port:  5151,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esri_sde",
			Port:  5151,
		},
	},
	"esro-emsdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esro-emsdp",
			Port:  642,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esro-emsdp",
			Port:  642,
		},
	},
	"esro-gen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "esro-gen",
			Port:  259,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "esro-gen",
			Port:  259,
		},
	},
	"essbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "essbase",
			Port:  1423,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "essbase",
			Port:  1423,
		},
	},
	"essp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "essp",
			Port:  2969,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "essp",
			Port:  2969,
		},
	},
	"essweb-gw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "essweb-gw",
			Port:  1772,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "essweb-gw",
			Port:  1772,
		},
	},
	"estamp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "estamp",
			Port:  1982,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "estamp",
			Port:  1982,
		},
	},
	"etb4j": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "etb4j",
			Port:  16309,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "etb4j",
			Port:  16309,
		},
	},
	"etc-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "etc-control",
			Port:  6107,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "etc-control",
			Port:  6107,
		},
	},
	"etebac5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "etebac5",
			Port:  1216,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "etebac5",
			Port:  1216,
		},
	},
	"etftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "etftp",
			Port:  1818,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "etftp",
			Port:  1818,
		},
	},
	"ethercat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ethercat",
			Port:  34980,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ethercat",
			Port:  34980,
		},
	},
	"ethoscan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ethoscan",
			Port:  6935,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ethoscan",
			Port:  6935,
		},
	},
	"etlservicemgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "etlservicemgr",
			Port:  9001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "etlservicemgr",
			Port:  9001,
		},
	},
	"etp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "etp",
			Port:  1798,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "etp",
			Port:  1798,
		},
	},
	"ets": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ets",
			Port:  1569,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ets",
			Port:  1569,
		},
	},
	"eudora-set": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eudora-set",
			Port:  592,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eudora-set",
			Port:  592,
		},
	},
	"ev-services": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ev-services",
			Port:  5114,
		},
	},
	"evb-elm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "evb-elm",
			Port:  1504,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "evb-elm",
			Port:  1504,
		},
	},
	"event-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "event-port",
			Port:  2069,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "event-port",
			Port:  2069,
		},
	},
	"event_listener": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "event_listener",
			Port:  3017,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "event_listener",
			Port:  3017,
		},
	},
	"everydayrc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "everydayrc",
			Port:  2782,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "everydayrc",
			Port:  2782,
		},
	},
	"evm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "evm",
			Port:  1139,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "evm",
			Port:  1139,
		},
	},
	"evtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "evtp",
			Port:  2834,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "evtp",
			Port:  2834,
		},
	},
	"evtp-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "evtp-data",
			Port:  2835,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "evtp-data",
			Port:  2835,
		},
	},
	"ew-disc-cmd": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ew-disc-cmd",
			Port:  43440,
		},
	},
	"ew-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ew-mgmt",
			Port:  43440,
		},
	},
	"ewall": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ewall",
			Port:  1328,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ewall",
			Port:  1328,
		},
	},
	"ewcappsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ewcappsrv",
			Port:  1876,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ewcappsrv",
			Port:  1876,
		},
	},
	"ewctsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ewctsp",
			Port:  6066,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ewctsp",
			Port:  6066,
		},
	},
	"ewdgs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ewdgs",
			Port:  4092,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ewdgs",
			Port:  4092,
		},
	},
	"ewinstaller": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ewinstaller",
			Port:  4091,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ewinstaller",
			Port:  4091,
		},
	},
	"ewnn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ewnn",
			Port:  2674,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ewnn",
			Port:  2674,
		},
	},
	"exapt-lmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exapt-lmgr",
			Port:  3759,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exapt-lmgr",
			Port:  3759,
		},
	},
	"exasoftport1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exasoftport1",
			Port:  3920,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exasoftport1",
			Port:  3920,
		},
	},
	"exbit-escp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exbit-escp",
			Port:  1316,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exbit-escp",
			Port:  1316,
		},
	},
	"exce": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exce",
			Port:  2769,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exce",
			Port:  2769,
		},
	},
	"excerpt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "excerpt",
			Port:  5400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "excerpt",
			Port:  5400,
		},
	},
	"excerpts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "excerpts",
			Port:  5401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "excerpts",
			Port:  5401,
		},
	},
	"excw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "excw",
			Port:  1271,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "excw",
			Port:  1271,
		},
	},
	"exec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exec",
			Port:  512,
		},
	},
	"exlm-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exlm-agent",
			Port:  3002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exlm-agent",
			Port:  3002,
		},
	},
	"exoconfig": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exoconfig",
			Port:  26487,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exoconfig",
			Port:  26487,
		},
	},
	"exoline-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exoline-tcp",
			Port:  26486,
		},
	},
	"exoline-udp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "exoline-udp",
			Port:  26486,
		},
	},
	"exonet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "exonet",
			Port:  26489,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exonet",
			Port:  26489,
		},
	},
	"exp1": map[string]Service{
		"dccp": Service{
			Proto: "dccp",
			Name:  "exp1",
			Port:  1021,
		},
		"sctp": Service{
			Proto: "sctp",
			Name:  "exp1",
			Port:  1021,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "exp1",
			Port:  1021,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exp1",
			Port:  1021,
		},
	},
	"exp2": map[string]Service{
		"dccp": Service{
			Proto: "dccp",
			Name:  "exp2",
			Port:  1022,
		},
		"sctp": Service{
			Proto: "sctp",
			Name:  "exp2",
			Port:  1022,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "exp2",
			Port:  1022,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "exp2",
			Port:  1022,
		},
	},
	"expresspay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "expresspay",
			Port:  2755,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "expresspay",
			Port:  2755,
		},
	},
	"extensis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "extensis",
			Port:  2666,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "extensis",
			Port:  2666,
		},
	},
	"eye2eye": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eye2eye",
			Port:  1948,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eye2eye",
			Port:  1948,
		},
	},
	"eyelink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eyelink",
			Port:  589,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eyelink",
			Port:  589,
		},
	},
	"eyetv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "eyetv",
			Port:  2170,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "eyetv",
			Port:  2170,
		},
	},
	"ezmeeting": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ezmeeting",
			Port:  26261,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ezmeeting",
			Port:  26261,
		},
	},
	"ezmeeting-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ezmeeting-2",
			Port:  10101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ezmeeting-2",
			Port:  10101,
		},
	},
	"ezmessagesrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ezmessagesrv",
			Port:  4085,
		},
	},
	"ezproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ezproxy",
			Port:  26260,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ezproxy",
			Port:  26260,
		},
	},
	"ezproxy-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ezproxy-2",
			Port:  10102,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ezproxy-2",
			Port:  10102,
		},
	},
	"ezrelay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ezrelay",
			Port:  10103,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ezrelay",
			Port:  10103,
		},
	},
	"f5-globalsite": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "f5-globalsite",
			Port:  2792,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "f5-globalsite",
			Port:  2792,
		},
	},
	"f5-iquery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "f5-iquery",
			Port:  4353,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "f5-iquery",
			Port:  4353,
		},
	},
	"fac-restore": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fac-restore",
			Port:  5582,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fac-restore",
			Port:  5582,
		},
	},
	"facelink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "facelink",
			Port:  1915,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "facelink",
			Port:  1915,
		},
	},
	"facilityview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "facilityview",
			Port:  1561,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "facilityview",
			Port:  1561,
		},
	},
	"facsys-ntp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "facsys-ntp",
			Port:  2514,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "facsys-ntp",
			Port:  2514,
		},
	},
	"facsys-router": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "facsys-router",
			Port:  2515,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "facsys-router",
			Port:  2515,
		},
	},
	"fagordnc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fagordnc",
			Port:  3873,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fagordnc",
			Port:  3873,
		},
	},
	"fairview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fairview",
			Port:  38202,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fairview",
			Port:  38202,
		},
	},
	"farenet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "farenet",
			Port:  5557,
		},
	},
	"fast-rem-serv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fast-rem-serv",
			Port:  2495,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fast-rem-serv",
			Port:  2495,
		},
	},
	"fastlynx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fastlynx",
			Port:  2689,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fastlynx",
			Port:  2689,
		},
	},
	"fatpipe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fatpipe",
			Port:  3353,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fatpipe",
			Port:  3353,
		},
	},
	"fatserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fatserv",
			Port:  347,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fatserv",
			Port:  347,
		},
	},
	"fax": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fax",
			Port:  4557,
		},
	},
	"faxcomservice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "faxcomservice",
			Port:  6417,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "faxcomservice",
			Port:  6417,
		},
	},
	"faximum": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "faximum",
			Port:  7437,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "faximum",
			Port:  7437,
		},
	},
	"faxportwinport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "faxportwinport",
			Port:  1620,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "faxportwinport",
			Port:  1620,
		},
	},
	"faxstfx-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "faxstfx-port",
			Port:  3684,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "faxstfx-port",
			Port:  3684,
		},
	},
	"fazzt-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fazzt-admin",
			Port:  4039,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fazzt-admin",
			Port:  4039,
		},
	},
	"fazzt-ptp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fazzt-ptp",
			Port:  4038,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fazzt-ptp",
			Port:  4038,
		},
	},
	"fc-cli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fc-cli",
			Port:  1371,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fc-cli",
			Port:  1371,
		},
	},
	"fc-faultnotify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fc-faultnotify",
			Port:  2819,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fc-faultnotify",
			Port:  2819,
		},
	},
	"fc-ser": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fc-ser",
			Port:  1372,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fc-ser",
			Port:  1372,
		},
	},
	"fcip-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcip-port",
			Port:  3225,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcip-port",
			Port:  3225,
		},
	},
	"fcis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcis",
			Port:  4727,
		},
	},
	"fcis-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "fcis-disc",
			Port:  4727,
		},
	},
	"fcmsys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcmsys",
			Port:  2344,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcmsys",
			Port:  2344,
		},
	},
	"fcopy-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcopy-server",
			Port:  5745,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcopy-server",
			Port:  5745,
		},
	},
	"fcopys-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcopys-server",
			Port:  5746,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcopys-server",
			Port:  5746,
		},
	},
	"fcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcp",
			Port:  510,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcp",
			Port:  510,
		},
	},
	"fcp-addr-srvr1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcp-addr-srvr1",
			Port:  5500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcp-addr-srvr1",
			Port:  5500,
		},
	},
	"fcp-addr-srvr2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcp-addr-srvr2",
			Port:  5501,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcp-addr-srvr2",
			Port:  5501,
		},
	},
	"fcp-cics-gw1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcp-cics-gw1",
			Port:  5504,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcp-cics-gw1",
			Port:  5504,
		},
	},
	"fcp-srvr-inst1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcp-srvr-inst1",
			Port:  5502,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcp-srvr-inst1",
			Port:  5502,
		},
	},
	"fcp-srvr-inst2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcp-srvr-inst2",
			Port:  5503,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcp-srvr-inst2",
			Port:  5503,
		},
	},
	"fcp-udp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fcp-udp",
			Port:  810,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fcp-udp",
			Port:  810,
		},
	},
	"fdt-rcatp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fdt-rcatp",
			Port:  4320,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fdt-rcatp",
			Port:  4320,
		},
	},
	"fdtracks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fdtracks",
			Port:  5579,
		},
	},
	"febooti-aw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "febooti-aw",
			Port:  36524,
		},
	},
	"feitianrockey": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "feitianrockey",
			Port:  3152,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "feitianrockey",
			Port:  3152,
		},
	},
	"femis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "femis",
			Port:  1776,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "femis",
			Port:  1776,
		},
	},
	"ferrari-foam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ferrari-foam",
			Port:  3216,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ferrari-foam",
			Port:  3216,
		},
	},
	"ff-annunc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ff-annunc",
			Port:  1089,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ff-annunc",
			Port:  1089,
		},
	},
	"ff-fms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ff-fms",
			Port:  1090,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ff-fms",
			Port:  1090,
		},
	},
	"ff-lr-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ff-lr-port",
			Port:  3622,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ff-lr-port",
			Port:  3622,
		},
	},
	"ff-sm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ff-sm",
			Port:  1091,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ff-sm",
			Port:  1091,
		},
	},
	"ffserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ffserver",
			Port:  3825,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ffserver",
			Port:  3825,
		},
	},
	"fg-fps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fg-fps",
			Port:  3293,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fg-fps",
			Port:  3293,
		},
	},
	"fg-gip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fg-gip",
			Port:  3294,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fg-gip",
			Port:  3294,
		},
	},
	"fg-sysupdate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fg-sysupdate",
			Port:  6550,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fg-sysupdate",
			Port:  6550,
		},
	},
	"fhc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fhc",
			Port:  1499,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fhc",
			Port:  1499,
		},
	},
	"fhsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fhsp",
			Port:  1807,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fhsp",
			Port:  1807,
		},
	},
	"fibotrader-com": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fibotrader-com",
			Port:  6715,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fibotrader-com",
			Port:  6715,
		},
	},
	"fido": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fido",
			Port:  60179,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fido",
			Port:  60179,
		},
	},
	"filecast": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filecast",
			Port:  3401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filecast",
			Port:  3401,
		},
	},
	"filemq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filemq",
			Port:  5670,
		},
	},
	"filenet-cm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-cm",
			Port:  32773,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-cm",
			Port:  32773,
		},
	},
	"filenet-nch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-nch",
			Port:  32770,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-nch",
			Port:  32770,
		},
	},
	"filenet-obrok": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-obrok",
			Port:  32777,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-obrok",
			Port:  32777,
		},
	},
	"filenet-pa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-pa",
			Port:  32772,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-pa",
			Port:  32772,
		},
	},
	"filenet-pch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-pch",
			Port:  32775,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-pch",
			Port:  32775,
		},
	},
	"filenet-peior": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-peior",
			Port:  32776,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-peior",
			Port:  32776,
		},
	},
	"filenet-powsrm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-powsrm",
			Port:  32767,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-powsrm",
			Port:  32767,
		},
	},
	"filenet-re": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-re",
			Port:  32774,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-re",
			Port:  32774,
		},
	},
	"filenet-rmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-rmi",
			Port:  32771,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-rmi",
			Port:  32771,
		},
	},
	"filenet-rpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-rpc",
			Port:  32769,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-rpc",
			Port:  32769,
		},
	},
	"filenet-tms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filenet-tms",
			Port:  32768,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filenet-tms",
			Port:  32768,
		},
	},
	"filesphere": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filesphere",
			Port:  24242,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filesphere",
			Port:  24242,
		},
	},
	"filex-lport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "filex-lport",
			Port:  1887,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "filex-lport",
			Port:  1887,
		},
	},
	"find": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "find",
			Port:  24922,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "find",
			Port:  24922,
		},
	},
	"findviatv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "findviatv",
			Port:  3350,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "findviatv",
			Port:  3350,
		},
	},
	"finger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "finger",
			Port:  79,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "finger",
			Port:  79,
		},
	},
	"finisar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "finisar",
			Port:  4682,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "finisar",
			Port:  4682,
		},
	},
	"finle-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "finle-lm",
			Port:  1784,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "finle-lm",
			Port:  1784,
		},
	},
	"fintrx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fintrx",
			Port:  3787,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fintrx",
			Port:  3787,
		},
	},
	"fio-cmgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fio-cmgmt",
			Port:  9051,
		},
	},
	"fiorano-msgsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fiorano-msgsvc",
			Port:  1856,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fiorano-msgsvc",
			Port:  1856,
		},
	},
	"fiorano-rtrsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fiorano-rtrsvc",
			Port:  1855,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fiorano-rtrsvc",
			Port:  1855,
		},
	},
	"firefox": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "firefox",
			Port:  1689,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "firefox",
			Port:  1689,
		},
	},
	"firemonrcc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "firemonrcc",
			Port:  3192,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "firemonrcc",
			Port:  3192,
		},
	},
	"firepower": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "firepower",
			Port:  2615,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "firepower",
			Port:  2615,
		},
	},
	"first-defense": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "first-defense",
			Port:  1232,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "first-defense",
			Port:  1232,
		},
	},
	"firstcall42": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "firstcall42",
			Port:  2673,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "firstcall42",
			Port:  2673,
		},
	},
	"fis": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "fis",
			Port:  5912,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "fis",
			Port:  5912,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fis",
			Port:  5912,
		},
	},
	"fisa-svc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fisa-svc",
			Port:  7018,
		},
	},
	"fiveacross": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fiveacross",
			Port:  1193,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fiveacross",
			Port:  1193,
		},
	},
	"fj-hdnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fj-hdnet",
			Port:  1717,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fj-hdnet",
			Port:  1717,
		},
	},
	"fjappmgrbulk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjappmgrbulk",
			Port:  2510,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjappmgrbulk",
			Port:  2510,
		},
	},
	"fjcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjcp",
			Port:  3648,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjcp",
			Port:  3648,
		},
	},
	"fjdmimgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjdmimgr",
			Port:  9374,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjdmimgr",
			Port:  9374,
		},
	},
	"fjdocdist": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjdocdist",
			Port:  1848,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjdocdist",
			Port:  1848,
		},
	},
	"fjhpjp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjhpjp",
			Port:  3067,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjhpjp",
			Port:  3067,
		},
	},
	"fjicl-tep-a": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjicl-tep-a",
			Port:  1901,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjicl-tep-a",
			Port:  1901,
		},
	},
	"fjicl-tep-b": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjicl-tep-b",
			Port:  1902,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjicl-tep-b",
			Port:  1902,
		},
	},
	"fjicl-tep-c": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjicl-tep-c",
			Port:  1904,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjicl-tep-c",
			Port:  1904,
		},
	},
	"fjinvmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjinvmgr",
			Port:  9396,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjinvmgr",
			Port:  9396,
		},
	},
	"fjippol-cnsl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjippol-cnsl",
			Port:  2749,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjippol-cnsl",
			Port:  2749,
		},
	},
	"fjippol-polsvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjippol-polsvr",
			Port:  2748,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjippol-polsvr",
			Port:  2748,
		},
	},
	"fjippol-port1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjippol-port1",
			Port:  2750,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjippol-port1",
			Port:  2750,
		},
	},
	"fjippol-port2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjippol-port2",
			Port:  2751,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjippol-port2",
			Port:  2751,
		},
	},
	"fjippol-swrly": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjippol-swrly",
			Port:  2747,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjippol-swrly",
			Port:  2747,
		},
	},
	"fjitsuappmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjitsuappmgr",
			Port:  2425,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjitsuappmgr",
			Port:  2425,
		},
	},
	"fjmpcm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjmpcm",
			Port:  2975,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjmpcm",
			Port:  2975,
		},
	},
	"fjmpjps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjmpjps",
			Port:  1873,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjmpjps",
			Port:  1873,
		},
	},
	"fjmpss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjmpss",
			Port:  2509,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjmpss",
			Port:  2509,
		},
	},
	"fjsv-gssagt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjsv-gssagt",
			Port:  3035,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjsv-gssagt",
			Port:  3035,
		},
	},
	"fjsvmpor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjsvmpor",
			Port:  2946,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjsvmpor",
			Port:  2946,
		},
	},
	"fjswapsnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fjswapsnp",
			Port:  1874,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fjswapsnp",
			Port:  1874,
		},
	},
	"fksp-audit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fksp-audit",
			Port:  3729,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fksp-audit",
			Port:  3729,
		},
	},
	"flamenco-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "flamenco-proxy",
			Port:  3210,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "flamenco-proxy",
			Port:  3210,
		},
	},
	"flashfiler": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "flashfiler",
			Port:  24677,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "flashfiler",
			Port:  24677,
		},
	},
	"flashmsg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "flashmsg",
			Port:  2884,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "flashmsg",
			Port:  2884,
		},
	},
	"flcrs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "flcrs",
			Port:  5638,
		},
	},
	"flexlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "flexlm",
			Port:  744,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "flexlm",
			Port:  744,
		},
	},
	"flirtmitmir": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "flirtmitmir",
			Port:  3840,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "flirtmitmir",
			Port:  3840,
		},
	},
	"fln-spx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fln-spx",
			Port:  221,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fln-spx",
			Port:  221,
		},
	},
	"florence": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "florence",
			Port:  1228,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "florence",
			Port:  1228,
		},
	},
	"flr_agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "flr_agent",
			Port:  4901,
		},
	},
	"flukeserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "flukeserver",
			Port:  2359,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "flukeserver",
			Port:  2359,
		},
	},
	"fly": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fly",
			Port:  4396,
		},
	},
	"fmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fmp",
			Port:  4745,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fmp",
			Port:  4745,
		},
	},
	"fmpro-fdal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fmpro-fdal",
			Port:  2399,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fmpro-fdal",
			Port:  2399,
		},
	},
	"fmpro-internal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fmpro-internal",
			Port:  5003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fmpro-internal",
			Port:  5003,
		},
	},
	"fmpro-v6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fmpro-v6",
			Port:  5013,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fmpro-v6",
			Port:  5013,
		},
	},
	"fmsas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fmsas",
			Port:  16000,
		},
	},
	"fmsascon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fmsascon",
			Port:  16001,
		},
	},
	"fmtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fmtp",
			Port:  8500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fmtp",
			Port:  8500,
		},
	},
	"fmwp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fmwp",
			Port:  5015,
		},
	},
	"fnet-remote-ui": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fnet-remote-ui",
			Port:  1174,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fnet-remote-ui",
			Port:  1174,
		},
	},
	"fodms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fodms",
			Port:  7200,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fodms",
			Port:  7200,
		},
	},
	"foliocorp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "foliocorp",
			Port:  2242,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "foliocorp",
			Port:  2242,
		},
	},
	"font-service": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "font-service",
			Port:  7100,
		},
	},
	"foresyte-clear": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "foresyte-clear",
			Port:  5407,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "foresyte-clear",
			Port:  5407,
		},
	},
	"foresyte-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "foresyte-sec",
			Port:  5408,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "foresyte-sec",
			Port:  5408,
		},
	},
	"fortisphere-vm": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "fortisphere-vm",
			Port:  4084,
		},
	},
	"fotogcad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fotogcad",
			Port:  3878,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fotogcad",
			Port:  3878,
		},
	},
	"found": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "found",
			Port:  4411,
		},
	},
	"fpitp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fpitp",
			Port:  1045,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fpitp",
			Port:  1045,
		},
	},
	"fpo-fns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fpo-fns",
			Port:  1066,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fpo-fns",
			Port:  1066,
		},
	},
	"fprams": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fprams",
			Port:  4122,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fprams",
			Port:  4122,
		},
	},
	"frc-hp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "frc-hp",
			Port:  6704,
		},
	},
	"frc-lp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "frc-lp",
			Port:  6706,
		},
	},
	"frc-mp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "frc-mp",
			Port:  6705,
		},
	},
	"frcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "frcs",
			Port:  4915,
		},
	},
	"freeciv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "freeciv",
			Port:  5556,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "freeciv",
			Port:  5556,
		},
	},
	"freezexservice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "freezexservice",
			Port:  7726,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "freezexservice",
			Port:  7726,
		},
	},
	"fronet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fronet",
			Port:  4130,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fronet",
			Port:  4130,
		},
	},
	"fryeserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fryeserv",
			Port:  2788,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fryeserv",
			Port:  2788,
		},
	},
	"fs-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fs-agent",
			Port:  8042,
		},
	},
	"fs-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fs-mgmt",
			Port:  8044,
		},
	},
	"fs-qos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fs-qos",
			Port:  41111,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fs-qos",
			Port:  41111,
		},
	},
	"fs-rh-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fs-rh-srv",
			Port:  3488,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fs-rh-srv",
			Port:  3488,
		},
	},
	"fs-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fs-server",
			Port:  8043,
		},
	},
	"fsc-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fsc-port",
			Port:  9217,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fsc-port",
			Port:  9217,
		},
	},
	"fse": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fse",
			Port:  7394,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fse",
			Port:  7394,
		},
	},
	"fsportmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fsportmap",
			Port:  4349,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fsportmap",
			Port:  4349,
		},
	},
	"fsr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fsr",
			Port:  7164,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fsr",
			Port:  7164,
		},
	},
	"ft-role": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ft-role",
			Port:  2429,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ft-role",
			Port:  2429,
		},
	},
	"ftp": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "ftp",
			Port:  21,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftp",
			Port:  21,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftp",
			Port:  21,
		},
	},
	"ftp-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftp-agent",
			Port:  574,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftp-agent",
			Port:  574,
		},
	},
	"ftp-data": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "ftp-data",
			Port:  20,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftp-data",
			Port:  20,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftp-data",
			Port:  20,
		},
	},
	"ftps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftps",
			Port:  990,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftps",
			Port:  990,
		},
	},
	"ftps-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftps-data",
			Port:  989,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftps-data",
			Port:  989,
		},
	},
	"ftranhc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftranhc",
			Port:  1105,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftranhc",
			Port:  1105,
		},
	},
	"ftrapid-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftrapid-1",
			Port:  1746,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftrapid-1",
			Port:  1746,
		},
	},
	"ftrapid-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftrapid-2",
			Port:  1747,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftrapid-2",
			Port:  1747,
		},
	},
	"ftsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ftsrv",
			Port:  1359,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ftsrv",
			Port:  1359,
		},
	},
	"ftsync": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ftsync",
			Port:  4086,
		},
	},
	"fud": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "fud",
			Port:  4201,
		},
	},
	"fujitsu-dev": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fujitsu-dev",
			Port:  747,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fujitsu-dev",
			Port:  747,
		},
	},
	"fujitsu-dtc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fujitsu-dtc",
			Port:  1513,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fujitsu-dtc",
			Port:  1513,
		},
	},
	"fujitsu-dtcns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fujitsu-dtcns",
			Port:  1514,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fujitsu-dtcns",
			Port:  1514,
		},
	},
	"fujitsu-mmpdc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fujitsu-mmpdc",
			Port:  1657,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fujitsu-mmpdc",
			Port:  1657,
		},
	},
	"fujitsu-neat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fujitsu-neat",
			Port:  3382,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fujitsu-neat",
			Port:  3382,
		},
	},
	"funk-dialout": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "funk-dialout",
			Port:  2909,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "funk-dialout",
			Port:  2909,
		},
	},
	"funk-license": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "funk-license",
			Port:  1787,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "funk-license",
			Port:  1787,
		},
	},
	"funk-logger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "funk-logger",
			Port:  1786,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "funk-logger",
			Port:  1786,
		},
	},
	"funkproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "funkproxy",
			Port:  1505,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "funkproxy",
			Port:  1505,
		},
	},
	"fuscript": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fuscript",
			Port:  1144,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fuscript",
			Port:  1144,
		},
	},
	"futrix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "futrix",
			Port:  2358,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "futrix",
			Port:  2358,
		},
	},
	"fxaengine-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fxaengine-net",
			Port:  3402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fxaengine-net",
			Port:  3402,
		},
	},
	"fxp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fxp",
			Port:  286,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fxp",
			Port:  286,
		},
	},
	"fxuptp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fxuptp",
			Port:  19539,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fxuptp",
			Port:  19539,
		},
	},
	"fyre-messanger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "fyre-messanger",
			Port:  2731,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "fyre-messanger",
			Port:  2731,
		},
	},
	"g-talk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "g-talk",
			Port:  2421,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "g-talk",
			Port:  2421,
		},
	},
	"g2tag": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "g2tag",
			Port:  4110,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "g2tag",
			Port:  4110,
		},
	},
	"g5m": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "g5m",
			Port:  2732,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "g5m",
			Port:  2732,
		},
	},
	"gacp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gacp",
			Port:  190,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gacp",
			Port:  190,
		},
	},
	"gadgetgate1way": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gadgetgate1way",
			Port:  2677,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gadgetgate1way",
			Port:  2677,
		},
	},
	"gadgetgate2way": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gadgetgate2way",
			Port:  2678,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gadgetgate2way",
			Port:  2678,
		},
	},
	"gadugadu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gadugadu",
			Port:  8074,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gadugadu",
			Port:  8074,
		},
	},
	"gaia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gaia",
			Port:  4340,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gaia",
			Port:  4340,
		},
	},
	"galaxy-network": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "galaxy-network",
			Port:  5235,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "galaxy-network",
			Port:  5235,
		},
	},
	"galaxy-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "galaxy-server",
			Port:  3051,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "galaxy-server",
			Port:  3051,
		},
	},
	"galaxy4d": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "galaxy4d",
			Port:  8881,
		},
	},
	"galaxy7-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "galaxy7-data",
			Port:  38201,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "galaxy7-data",
			Port:  38201,
		},
	},
	"galileo": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "galileo",
			Port:  3519,
		},
	},
	"galileolog": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "galileolog",
			Port:  3520,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "galileolog",
			Port:  3520,
		},
	},
	"gamegen1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gamegen1",
			Port:  1738,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gamegen1",
			Port:  1738,
		},
	},
	"gamelobby": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gamelobby",
			Port:  2914,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gamelobby",
			Port:  2914,
		},
	},
	"gamesmith-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gamesmith-port",
			Port:  31765,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gamesmith-port",
			Port:  31765,
		},
	},
	"gammafetchsvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gammafetchsvr",
			Port:  1859,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gammafetchsvr",
			Port:  1859,
		},
	},
	"gandalf-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gandalf-lm",
			Port:  1421,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gandalf-lm",
			Port:  1421,
		},
	},
	"gap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gap",
			Port:  10800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gap",
			Port:  10800,
		},
	},
	"garcon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "garcon",
			Port:  999,
		},
	},
	"gat-lmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gat-lmd",
			Port:  1708,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gat-lmd",
			Port:  1708,
		},
	},
	"gbjd816": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gbjd816",
			Port:  2626,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gbjd816",
			Port:  2626,
		},
	},
	"gbmt-stars": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gbmt-stars",
			Port:  3912,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gbmt-stars",
			Port:  3912,
		},
	},
	"gbs-smp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gbs-smp",
			Port:  3762,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gbs-smp",
			Port:  3762,
		},
	},
	"gbs-stp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gbs-stp",
			Port:  3484,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gbs-stp",
			Port:  3484,
		},
	},
	"gc-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gc-config",
			Port:  3436,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gc-config",
			Port:  3436,
		},
	},
	"gcm-app": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gcm-app",
			Port:  14145,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gcm-app",
			Port:  14145,
		},
	},
	"gcmonitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gcmonitor",
			Port:  2660,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gcmonitor",
			Port:  2660,
		},
	},
	"gcsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gcsp",
			Port:  3429,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gcsp",
			Port:  3429,
		},
	},
	"gdbremote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gdbremote",
			Port:  2159,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gdbremote",
			Port:  2159,
		},
	},
	"gdoi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gdoi",
			Port:  848,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gdoi",
			Port:  848,
		},
	},
	"gdomap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gdomap",
			Port:  538,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gdomap",
			Port:  538,
		},
	},
	"gdp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gdp-port",
			Port:  1997,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gdp-port",
			Port:  1997,
		},
	},
	"gdrive-sync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gdrive-sync",
			Port:  37483,
		},
	},
	"gds-adppiw-db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gds-adppiw-db",
			Port:  4550,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gds-adppiw-db",
			Port:  4550,
		},
	},
	"gds_db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gds_db",
			Port:  3050,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gds_db",
			Port:  3050,
		},
	},
	"gearman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gearman",
			Port:  4730,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gearman",
			Port:  4730,
		},
	},
	"gemini-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gemini-lm",
			Port:  1590,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gemini-lm",
			Port:  1590,
		},
	},
	"geneous": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "geneous",
			Port:  3381,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "geneous",
			Port:  3381,
		},
	},
	"genie": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "genie",
			Port:  402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "genie",
			Port:  402,
		},
	},
	"genie-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "genie-lm",
			Port:  1453,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "genie-lm",
			Port:  1453,
		},
	},
	"genisar-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "genisar-port",
			Port:  3475,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "genisar-port",
			Port:  3475,
		},
	},
	"geniuslm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "geniuslm",
			Port:  3005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "geniuslm",
			Port:  3005,
		},
	},
	"genrad-mux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "genrad-mux",
			Port:  176,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "genrad-mux",
			Port:  176,
		},
	},
	"genstat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "genstat",
			Port:  7283,
		},
	},
	"geognosis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "geognosis",
			Port:  4326,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "geognosis",
			Port:  4326,
		},
	},
	"geognosisman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "geognosisman",
			Port:  4325,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "geognosisman",
			Port:  4325,
		},
	},
	"geolocate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "geolocate",
			Port:  3108,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "geolocate",
			Port:  3108,
		},
	},
	"gerhcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gerhcs",
			Port:  4985,
		},
	},
	"gf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gf",
			Port:  3530,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gf",
			Port:  3530,
		},
	},
	"ggf-ncp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ggf-ncp",
			Port:  678,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ggf-ncp",
			Port:  678,
		},
	},
	"ggz": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ggz",
			Port:  5688,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ggz",
			Port:  5688,
		},
	},
	"ghvpn": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ghvpn",
			Port:  12009,
		},
	},
	"giga-pocket": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "giga-pocket",
			Port:  3862,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "giga-pocket",
			Port:  3862,
		},
	},
	"gilatskysurfer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gilatskysurfer",
			Port:  3013,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gilatskysurfer",
			Port:  3013,
		},
	},
	"ginad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ginad",
			Port:  634,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ginad",
			Port:  634,
		},
	},
	"giop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "giop",
			Port:  2481,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "giop",
			Port:  2481,
		},
	},
	"giop-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "giop-ssl",
			Port:  2482,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "giop-ssl",
			Port:  2482,
		},
	},
	"gist": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "gist",
			Port:  270,
		},
	},
	"git": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "git",
			Port:  9418,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "git",
			Port:  9418,
		},
	},
	"glbp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "glbp",
			Port:  3222,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "glbp",
			Port:  3222,
		},
	},
	"gld": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gld",
			Port:  6267,
		},
	},
	"glishd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "glishd",
			Port:  2833,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "glishd",
			Port:  2833,
		},
	},
	"global-cd-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "global-cd-port",
			Port:  3229,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "global-cd-port",
			Port:  3229,
		},
	},
	"global-dtserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "global-dtserv",
			Port:  1774,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "global-dtserv",
			Port:  1774,
		},
	},
	"global-wlink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "global-wlink",
			Port:  1909,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "global-wlink",
			Port:  1909,
		},
	},
	"globe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "globe",
			Port:  2002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "globe",
			Port:  2002,
		},
	},
	"globecast-id": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "globecast-id",
			Port:  6109,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "globecast-id",
			Port:  6109,
		},
	},
	"globmsgsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "globmsgsvc",
			Port:  2519,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "globmsgsvc",
			Port:  2519,
		},
	},
	"glogger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "glogger",
			Port:  2033,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "glogger",
			Port:  2033,
		},
	},
	"glrpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "glrpc",
			Port:  9080,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "glrpc",
			Port:  9080,
		},
	},
	"gmmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gmmp",
			Port:  4183,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gmmp",
			Port:  4183,
		},
	},
	"gmrupdateserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gmrupdateserv",
			Port:  1070,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gmrupdateserv",
			Port:  1070,
		},
	},
	"gntp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gntp",
			Port:  23053,
		},
	},
	"gnunet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gnunet",
			Port:  2086,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gnunet",
			Port:  2086,
		},
	},
	"gnutella-rtr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gnutella-rtr",
			Port:  6347,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gnutella-rtr",
			Port:  6347,
		},
	},
	"gnutella-svc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gnutella-svc",
			Port:  6346,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gnutella-svc",
			Port:  6346,
		},
	},
	"go-login": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "go-login",
			Port:  491,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "go-login",
			Port:  491,
		},
	},
	"goahead-fldup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "goahead-fldup",
			Port:  3057,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "goahead-fldup",
			Port:  3057,
		},
	},
	"goldleaf-licman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "goldleaf-licman",
			Port:  1401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "goldleaf-licman",
			Port:  1401,
		},
	},
	"gopher": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gopher",
			Port:  70,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gopher",
			Port:  70,
		},
	},
	"gotodevice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gotodevice",
			Port:  2217,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gotodevice",
			Port:  2217,
		},
	},
	"gpfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gpfs",
			Port:  1191,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gpfs",
			Port:  1191,
		},
	},
	"gppitnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gppitnp",
			Port:  103,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gppitnp",
			Port:  103,
		},
	},
	"gprs-cube": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gprs-cube",
			Port:  3751,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gprs-cube",
			Port:  3751,
		},
	},
	"gprs-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gprs-data",
			Port:  3386,
		},
	},
	"gprs-sig": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "gprs-sig",
			Port:  3386,
		},
	},
	"gpsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gpsd",
			Port:  2947,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gpsd",
			Port:  2947,
		},
	},
	"gradecam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gradecam",
			Port:  5117,
		},
	},
	"graphics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "graphics",
			Port:  41,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "graphics",
			Port:  41,
		},
	},
	"grcmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "grcmp",
			Port:  9122,
		},
	},
	"grcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "grcp",
			Port:  9123,
		},
	},
	"grf-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "grf-port",
			Port:  3757,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "grf-port",
			Port:  3757,
		},
	},
	"grid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "grid",
			Port:  6268,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "grid",
			Port:  6268,
		},
	},
	"grid-alt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "grid-alt",
			Port:  6269,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "grid-alt",
			Port:  6269,
		},
	},
	"gridgen-elmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gridgen-elmd",
			Port:  1542,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gridgen-elmd",
			Port:  1542,
		},
	},
	"griffin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "griffin",
			Port:  2458,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "griffin",
			Port:  2458,
		},
	},
	"gris": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gris",
			Port:  2135,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gris",
			Port:  2135,
		},
	},
	"groove": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "groove",
			Port:  2492,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "groove",
			Port:  2492,
		},
	},
	"groove-dpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "groove-dpp",
			Port:  1211,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "groove-dpp",
			Port:  1211,
		},
	},
	"groupwise": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "groupwise",
			Port:  1677,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "groupwise",
			Port:  1677,
		},
	},
	"grubd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "grubd",
			Port:  3136,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "grubd",
			Port:  3136,
		},
	},
	"gsakmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gsakmp",
			Port:  3761,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gsakmp",
			Port:  3761,
		},
	},
	"gsi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gsi",
			Port:  1850,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gsi",
			Port:  1850,
		},
	},
	"gsidcap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gsidcap",
			Port:  22128,
		},
	},
	"gsiftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gsiftp",
			Port:  2811,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gsiftp",
			Port:  2811,
		},
	},
	"gsigatekeeper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gsigatekeeper",
			Port:  2119,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gsigatekeeper",
			Port:  2119,
		},
	},
	"gsmp-ancp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gsmp-ancp",
			Port:  6068,
		},
	},
	"gsms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gsms",
			Port:  16002,
		},
	},
	"gsmtap": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "gsmtap",
			Port:  4729,
		},
	},
	"gss-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gss-http",
			Port:  488,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gss-http",
			Port:  488,
		},
	},
	"gss-xlicen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gss-xlicen",
			Port:  128,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gss-xlicen",
			Port:  128,
		},
	},
	"gt-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gt-proxy",
			Port:  9889,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gt-proxy",
			Port:  9889,
		},
	},
	"gtaua": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gtaua",
			Port:  2186,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gtaua",
			Port:  2186,
		},
	},
	"gte-samp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gte-samp",
			Port:  2643,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gte-samp",
			Port:  2643,
		},
	},
	"gtegsc-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gtegsc-lm",
			Port:  1452,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gtegsc-lm",
			Port:  1452,
		},
	},
	"gtp-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gtp-control",
			Port:  2123,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gtp-control",
			Port:  2123,
		},
	},
	"gtp-user": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gtp-user",
			Port:  2152,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gtp-user",
			Port:  2152,
		},
	},
	"gtrack-ne": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gtrack-ne",
			Port:  3592,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gtrack-ne",
			Port:  3592,
		},
	},
	"gtrack-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gtrack-server",
			Port:  3591,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gtrack-server",
			Port:  3591,
		},
	},
	"guibase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "guibase",
			Port:  9321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "guibase",
			Port:  9321,
		},
	},
	"guttersnex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "guttersnex",
			Port:  35356,
		},
	},
	"gv-pf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gv-pf",
			Port:  18262,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gv-pf",
			Port:  18262,
		},
	},
	"gv-us": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gv-us",
			Port:  1369,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gv-us",
			Port:  1369,
		},
	},
	"gvcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gvcp",
			Port:  3956,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gvcp",
			Port:  3956,
		},
	},
	"gw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gw",
			Port:  3010,
		},
	},
	"gw-asv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gw-asv",
			Port:  4842,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gw-asv",
			Port:  4842,
		},
	},
	"gw-call-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gw-call-port",
			Port:  3745,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gw-call-port",
			Port:  3745,
		},
	},
	"gw-log": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gw-log",
			Port:  4844,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gw-log",
			Port:  4844,
		},
	},
	"gwen-sonya": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gwen-sonya",
			Port:  2778,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gwen-sonya",
			Port:  2778,
		},
	},
	"gwha": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gwha",
			Port:  1383,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gwha",
			Port:  1383,
		},
	},
	"gxs-data-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gxs-data-port",
			Port:  2073,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gxs-data-port",
			Port:  2073,
		},
	},
	"gxtelmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "gxtelmd",
			Port:  2356,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "gxtelmd",
			Port:  2356,
		},
	},
	"h2250-annex-g": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "h2250-annex-g",
			Port:  2099,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h2250-annex-g",
			Port:  2099,
		},
	},
	"h248-binary": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "h248-binary",
			Port:  2945,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "h248-binary",
			Port:  2945,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h248-binary",
			Port:  2945,
		},
	},
	"h263-video": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "h263-video",
			Port:  2979,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h263-video",
			Port:  2979,
		},
	},
	"h2gf-w-2m": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "h2gf-w-2m",
			Port:  3179,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h2gf-w-2m",
			Port:  3179,
		},
	},
	"h323callsigalt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "h323callsigalt",
			Port:  11720,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h323callsigalt",
			Port:  11720,
		},
	},
	"h323gatedisc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "h323gatedisc",
			Port:  1718,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h323gatedisc",
			Port:  1718,
		},
	},
	"h323gatestat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "h323gatestat",
			Port:  1719,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h323gatestat",
			Port:  1719,
		},
	},
	"h323hostcall": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "h323hostcall",
			Port:  1720,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h323hostcall",
			Port:  1720,
		},
	},
	"h323hostcallsc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "h323hostcallsc",
			Port:  1300,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "h323hostcallsc",
			Port:  1300,
		},
	},
	"ha-cluster": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ha-cluster",
			Port:  694,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ha-cluster",
			Port:  694,
		},
	},
	"hacl-cfg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-cfg",
			Port:  5302,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-cfg",
			Port:  5302,
		},
	},
	"hacl-gs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-gs",
			Port:  5301,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-gs",
			Port:  5301,
		},
	},
	"hacl-hb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-hb",
			Port:  5300,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-hb",
			Port:  5300,
		},
	},
	"hacl-local": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-local",
			Port:  5304,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-local",
			Port:  5304,
		},
	},
	"hacl-monitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-monitor",
			Port:  3542,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-monitor",
			Port:  3542,
		},
	},
	"hacl-poll": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-poll",
			Port:  5315,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-poll",
			Port:  5315,
		},
	},
	"hacl-probe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-probe",
			Port:  5303,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-probe",
			Port:  5303,
		},
	},
	"hacl-qs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-qs",
			Port:  1238,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-qs",
			Port:  1238,
		},
	},
	"hacl-test": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hacl-test",
			Port:  5305,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hacl-test",
			Port:  5305,
		},
	},
	"hagel-dump": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hagel-dump",
			Port:  3036,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hagel-dump",
			Port:  3036,
		},
	},
	"haipe-discover": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "haipe-discover",
			Port:  3623,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "haipe-discover",
			Port:  3623,
		},
	},
	"haipe-otnk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "haipe-otnk",
			Port:  3769,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "haipe-otnk",
			Port:  3769,
		},
	},
	"hao": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hao",
			Port:  2245,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hao",
			Port:  2245,
		},
	},
	"hap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hap",
			Port:  661,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hap",
			Port:  661,
		},
	},
	"harp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "harp",
			Port:  1816,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "harp",
			Port:  1816,
		},
	},
	"hart-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hart-ip",
			Port:  5094,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hart-ip",
			Port:  5094,
		},
	},
	"hassle": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hassle",
			Port:  375,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hassle",
			Port:  375,
		},
	},
	"hawk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hawk",
			Port:  7630,
		},
	},
	"hb-engine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hb-engine",
			Port:  1703,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hb-engine",
			Port:  1703,
		},
	},
	"hbci": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hbci",
			Port:  3000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hbci",
			Port:  3000,
		},
	},
	"hcp-wismar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hcp-wismar",
			Port:  686,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hcp-wismar",
			Port:  686,
		},
	},
	"hdap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hdap",
			Port:  263,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hdap",
			Port:  263,
		},
	},
	"hde-lcesrvr-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hde-lcesrvr-1",
			Port:  14936,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hde-lcesrvr-1",
			Port:  14936,
		},
	},
	"hde-lcesrvr-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hde-lcesrvr-2",
			Port:  14937,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hde-lcesrvr-2",
			Port:  14937,
		},
	},
	"hdl-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hdl-srv",
			Port:  2641,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hdl-srv",
			Port:  2641,
		},
	},
	"health-polling": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "health-polling",
			Port:  1161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "health-polling",
			Port:  1161,
		},
	},
	"health-trap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "health-trap",
			Port:  1162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "health-trap",
			Port:  1162,
		},
	},
	"healthd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "healthd",
			Port:  1281,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "healthd",
			Port:  1281,
		},
	},
	"heartbeat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "heartbeat",
			Port:  3740,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "heartbeat",
			Port:  3740,
		},
	},
	"heathview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "heathview",
			Port:  35000,
		},
	},
	"hecmtl-db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hecmtl-db",
			Port:  1551,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hecmtl-db",
			Port:  1551,
		},
	},
	"helix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "helix",
			Port:  10860,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "helix",
			Port:  10860,
		},
	},
	"hello": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hello",
			Port:  1789,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hello",
			Port:  1789,
		},
	},
	"hello-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hello-port",
			Port:  652,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hello-port",
			Port:  652,
		},
	},
	"hems": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hems",
			Port:  151,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hems",
			Port:  151,
		},
	},
	"here-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "here-lm",
			Port:  1409,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "here-lm",
			Port:  1409,
		},
	},
	"hermes": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hermes",
			Port:  1248,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hermes",
			Port:  1248,
		},
	},
	"herodotus-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "herodotus-net",
			Port:  3921,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "herodotus-net",
			Port:  3921,
		},
	},
	"hexarc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hexarc",
			Port:  7397,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hexarc",
			Port:  7397,
		},
	},
	"hfcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hfcs",
			Port:  4900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hfcs",
			Port:  4900,
		},
	},
	"hfcs-manager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hfcs-manager",
			Port:  4999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hfcs-manager",
			Port:  4999,
		},
	},
	"hhb-gateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hhb-gateway",
			Port:  1136,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hhb-gateway",
			Port:  1136,
		},
	},
	"hhb-handheld": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hhb-handheld",
			Port:  4148,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hhb-handheld",
			Port:  4148,
		},
	},
	"hicp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hicp",
			Port:  3250,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hicp",
			Port:  3250,
		},
	},
	"hid": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "hid",
			Port:  24322,
		},
	},
	"high-criteria": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "high-criteria",
			Port:  2467,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "high-criteria",
			Port:  2467,
		},
	},
	"hillrserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hillrserv",
			Port:  4117,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hillrserv",
			Port:  4117,
		},
	},
	"hinp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hinp",
			Port:  9954,
		},
	},
	"hip-nat-t": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "hip-nat-t",
			Port:  10500,
		},
	},
	"hiperscan-id": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hiperscan-id",
			Port:  8293,
		},
	},
	"hiq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hiq",
			Port:  1410,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hiq",
			Port:  1410,
		},
	},
	"hislip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hislip",
			Port:  4880,
		},
	},
	"hivep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hivep",
			Port:  12172,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hivep",
			Port:  12172,
		},
	},
	"hivestor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hivestor",
			Port:  4884,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hivestor",
			Port:  4884,
		},
	},
	"hks-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hks-lm",
			Port:  1722,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hks-lm",
			Port:  1722,
		},
	},
	"hl7": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hl7",
			Port:  2575,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hl7",
			Port:  2575,
		},
	},
	"hlibmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hlibmgr",
			Port:  3634,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hlibmgr",
			Port:  3634,
		},
	},
	"hlserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hlserver",
			Port:  3047,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hlserver",
			Port:  3047,
		},
	},
	"hmmp-ind": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hmmp-ind",
			Port:  612,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hmmp-ind",
			Port:  612,
		},
	},
	"hmmp-op": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hmmp-op",
			Port:  613,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hmmp-op",
			Port:  613,
		},
	},
	"hnm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hnm",
			Port:  6791,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hnm",
			Port:  6791,
		},
	},
	"hnmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hnmp",
			Port:  6790,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hnmp",
			Port:  6790,
		},
	},
	"homeportal-web": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "homeportal-web",
			Port:  3941,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "homeportal-web",
			Port:  3941,
		},
	},
	"homesteadglory": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "homesteadglory",
			Port:  2597,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "homesteadglory",
			Port:  2597,
		},
	},
	"honyaku": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "honyaku",
			Port:  2744,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "honyaku",
			Port:  2744,
		},
	},
	"hostmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hostmon",
			Port:  5355,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hostmon",
			Port:  5355,
		},
	},
	"hostname": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hostname",
			Port:  101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hostname",
			Port:  101,
		},
	},
	"hotu-chat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hotu-chat",
			Port:  3449,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hotu-chat",
			Port:  3449,
		},
	},
	"houdini-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "houdini-lm",
			Port:  1715,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "houdini-lm",
			Port:  1715,
		},
	},
	"houston": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "houston",
			Port:  4041,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "houston",
			Port:  4041,
		},
	},
	"hp-3000-telnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-3000-telnet",
			Port:  2564,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-3000-telnet",
			Port:  2564,
		},
	},
	"hp-alarm-mgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-alarm-mgr",
			Port:  383,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-alarm-mgr",
			Port:  383,
		},
	},
	"hp-clic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-clic",
			Port:  3384,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-clic",
			Port:  3384,
		},
	},
	"hp-collector": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-collector",
			Port:  381,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-collector",
			Port:  381,
		},
	},
	"hp-dataprotect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-dataprotect",
			Port:  3612,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-dataprotect",
			Port:  3612,
		},
	},
	"hp-device-disc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-device-disc",
			Port:  3329,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-device-disc",
			Port:  3329,
		},
	},
	"hp-hcip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-hcip",
			Port:  1782,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-hcip",
			Port:  1782,
		},
	},
	"hp-hcip-gwy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-hcip-gwy",
			Port:  1803,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-hcip-gwy",
			Port:  1803,
		},
	},
	"hp-managed-node": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-managed-node",
			Port:  382,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-managed-node",
			Port:  382,
		},
	},
	"hp-nnm-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-nnm-data",
			Port:  2690,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-nnm-data",
			Port:  2690,
		},
	},
	"hp-pdl-datastr": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "hp-pdl-datastr",
			Port:  9100,
		},
	},
	"hp-pxpib": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-pxpib",
			Port:  3101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-pxpib",
			Port:  3101,
		},
	},
	"hp-san-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-san-mgmt",
			Port:  3037,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-san-mgmt",
			Port:  3037,
		},
	},
	"hp-sca": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-sca",
			Port:  19411,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-sca",
			Port:  19411,
		},
	},
	"hp-sci": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-sci",
			Port:  1299,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-sci",
			Port:  1299,
		},
	},
	"hp-sco": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-sco",
			Port:  19410,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-sco",
			Port:  19410,
		},
	},
	"hp-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-server",
			Port:  5225,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-server",
			Port:  5225,
		},
	},
	"hp-sessmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-sessmon",
			Port:  19412,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-sessmon",
			Port:  19412,
		},
	},
	"hp-status": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-status",
			Port:  5226,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-status",
			Port:  5226,
		},
	},
	"hp-webadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-webadmin",
			Port:  1188,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-webadmin",
			Port:  1188,
		},
	},
	"hp-webqosdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hp-webqosdb",
			Port:  1877,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hp-webqosdb",
			Port:  1877,
		},
	},
	"hpbladems": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpbladems",
			Port:  5316,
		},
	},
	"hpdevms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpdevms",
			Port:  5317,
		},
	},
	"hpidsadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpidsadmin",
			Port:  2984,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpidsadmin",
			Port:  2984,
		},
	},
	"hpidsagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpidsagent",
			Port:  2985,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpidsagent",
			Port:  2985,
		},
	},
	"hpiod": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpiod",
			Port:  2208,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpiod",
			Port:  2208,
		},
	},
	"hpocbus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpocbus",
			Port:  2206,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpocbus",
			Port:  2206,
		},
	},
	"hpoms-ci-lstn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpoms-ci-lstn",
			Port:  5403,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpoms-ci-lstn",
			Port:  5403,
		},
	},
	"hpoms-dps-lstn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpoms-dps-lstn",
			Port:  5404,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpoms-dps-lstn",
			Port:  5404,
		},
	},
	"hpppssvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpppssvr",
			Port:  2448,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpppssvr",
			Port:  2448,
		},
	},
	"hppronetman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hppronetman",
			Port:  3908,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hppronetman",
			Port:  3908,
		},
	},
	"hpss-ndapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpss-ndapi",
			Port:  1217,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpss-ndapi",
			Port:  1217,
		},
	},
	"hpssd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpssd",
			Port:  2207,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpssd",
			Port:  2207,
		},
	},
	"hpssmgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpssmgmt",
			Port:  4484,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpssmgmt",
			Port:  4484,
		},
	},
	"hpstgmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpstgmgr",
			Port:  2600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpstgmgr",
			Port:  2600,
		},
	},
	"hpstgmgr2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpstgmgr2",
			Port:  2715,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpstgmgr2",
			Port:  2715,
		},
	},
	"hpvirtctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpvirtctrl",
			Port:  5224,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpvirtctrl",
			Port:  5224,
		},
	},
	"hpvirtgrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpvirtgrp",
			Port:  5223,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpvirtgrp",
			Port:  5223,
		},
	},
	"hpvmmagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpvmmagent",
			Port:  1125,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpvmmagent",
			Port:  1125,
		},
	},
	"hpvmmcontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpvmmcontrol",
			Port:  1124,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpvmmcontrol",
			Port:  1124,
		},
	},
	"hpvmmdata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpvmmdata",
			Port:  1126,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hpvmmdata",
			Port:  1126,
		},
	},
	"hpvroom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hpvroom",
			Port:  5228,
		},
	},
	"hrd-ncs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hrd-ncs",
			Port:  6324,
		},
	},
	"hrd-ns-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "hrd-ns-disc",
			Port:  6324,
		},
	},
	"hri-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hri-port",
			Port:  3439,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hri-port",
			Port:  3439,
		},
	},
	"hrpd-ith-at-an": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "hrpd-ith-at-an",
			Port:  4592,
		},
	},
	"hs-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hs-port",
			Port:  2570,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hs-port",
			Port:  2570,
		},
	},
	"hsl-storm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hsl-storm",
			Port:  2113,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hsl-storm",
			Port:  2113,
		},
	},
	"hsrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hsrp",
			Port:  1985,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hsrp",
			Port:  1985,
		},
	},
	"hsrpv6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hsrpv6",
			Port:  2029,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hsrpv6",
			Port:  2029,
		},
	},
	"htcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "htcp",
			Port:  4827,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "htcp",
			Port:  4827,
		},
	},
	"htrust": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "htrust",
			Port:  5628,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "htrust",
			Port:  5628,
		},
	},
	"http": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "http",
			Port:  80,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "http",
			Port:  80,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "http",
			Port:  80,
		},
	},
	"http-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "http-mgmt",
			Port:  280,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "http-mgmt",
			Port:  280,
		},
	},
	"http-rpc-epmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "http-rpc-epmap",
			Port:  593,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "http-rpc-epmap",
			Port:  593,
		},
	},
	"http-wmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "http-wmap",
			Port:  8990,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "http-wmap",
			Port:  8990,
		},
	},
	"https": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "https",
			Port:  443,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "https",
			Port:  443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "https",
			Port:  443,
		},
	},
	"https-wmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "https-wmap",
			Port:  8991,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "https-wmap",
			Port:  8991,
		},
	},
	"httpx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "httpx",
			Port:  4180,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "httpx",
			Port:  4180,
		},
	},
	"htuilsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "htuilsrv",
			Port:  5023,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "htuilsrv",
			Port:  5023,
		},
	},
	"hub-open-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hub-open-net",
			Port:  8313,
		},
	},
	"hughes-ap": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "hughes-ap",
			Port:  5105,
		},
	},
	"husky": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "husky",
			Port:  1310,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "husky",
			Port:  1310,
		},
	},
	"hybrid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hybrid",
			Port:  1424,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hybrid",
			Port:  1424,
		},
	},
	"hybrid-pop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hybrid-pop",
			Port:  473,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hybrid-pop",
			Port:  473,
		},
	},
	"hydap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hydap",
			Port:  15000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hydap",
			Port:  15000,
		},
	},
	"hydra": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hydra",
			Port:  2374,
		},
	},
	"hylafax": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hylafax",
			Port:  4559,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hylafax",
			Port:  4559,
		},
	},
	"hyper-g": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hyper-g",
			Port:  418,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hyper-g",
			Port:  418,
		},
	},
	"hypercube-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hypercube-lm",
			Port:  1577,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hypercube-lm",
			Port:  1577,
		},
	},
	"hyperip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hyperip",
			Port:  3919,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hyperip",
			Port:  3919,
		},
	},
	"hyperscsi-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hyperscsi-port",
			Port:  5674,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hyperscsi-port",
			Port:  5674,
		},
	},
	"hyperwave-isp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "hyperwave-isp",
			Port:  692,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "hyperwave-isp",
			Port:  692,
		},
	},
	"i-net-2000-npr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "i-net-2000-npr",
			Port:  5069,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "i-net-2000-npr",
			Port:  5069,
		},
	},
	"i-zipqd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "i-zipqd",
			Port:  13160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "i-zipqd",
			Port:  13160,
		},
	},
	"i3-sessionmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "i3-sessionmgr",
			Port:  3952,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "i3-sessionmgr",
			Port:  3952,
		},
	},
	"iRAPP": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iRAPP",
			Port:  4073,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iRAPP",
			Port:  4073,
		},
	},
	"iad1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iad1",
			Port:  1030,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iad1",
			Port:  1030,
		},
	},
	"iad2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iad2",
			Port:  1031,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iad2",
			Port:  1031,
		},
	},
	"iad3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iad3",
			Port:  1032,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iad3",
			Port:  1032,
		},
	},
	"iadt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iadt",
			Port:  4169,
		},
	},
	"iadt-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "iadt-disc",
			Port:  4169,
		},
	},
	"iadt-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iadt-tls",
			Port:  9614,
		},
	},
	"iafdbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iafdbase",
			Port:  480,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iafdbase",
			Port:  480,
		},
	},
	"iafserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iafserver",
			Port:  479,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iafserver",
			Port:  479,
		},
	},
	"ianywhere-dbns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ianywhere-dbns",
			Port:  3968,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ianywhere-dbns",
			Port:  3968,
		},
	},
	"iapp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iapp",
			Port:  2313,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iapp",
			Port:  2313,
		},
	},
	"ias-admind": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ias-admind",
			Port:  2141,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ias-admind",
			Port:  2141,
		},
	},
	"ias-auth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ias-auth",
			Port:  2139,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ias-auth",
			Port:  2139,
		},
	},
	"ias-neighbor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ias-neighbor",
			Port:  4596,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ias-neighbor",
			Port:  4596,
		},
	},
	"ias-paging": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ias-paging",
			Port:  4595,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ias-paging",
			Port:  4595,
		},
	},
	"ias-reg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ias-reg",
			Port:  2140,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ias-reg",
			Port:  2140,
		},
	},
	"ias-session": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ias-session",
			Port:  4594,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ias-session",
			Port:  4594,
		},
	},
	"iascontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iascontrol",
			Port:  1157,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iascontrol",
			Port:  1157,
		},
	},
	"iascontrol-oms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iascontrol-oms",
			Port:  1156,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iascontrol-oms",
			Port:  1156,
		},
	},
	"iasd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iasd",
			Port:  432,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iasd",
			Port:  432,
		},
	},
	"iatp-highpri": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iatp-highpri",
			Port:  6998,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iatp-highpri",
			Port:  6998,
		},
	},
	"iatp-normalpri": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iatp-normalpri",
			Port:  6999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iatp-normalpri",
			Port:  6999,
		},
	},
	"iax": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iax",
			Port:  4569,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iax",
			Port:  4569,
		},
	},
	"ibar": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ibar",
			Port:  5784,
		},
	},
	"iberiagames": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iberiagames",
			Port:  1726,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iberiagames",
			Port:  1726,
		},
	},
	"ibm-abtact": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-abtact",
			Port:  1586,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-abtact",
			Port:  1586,
		},
	},
	"ibm-app": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-app",
			Port:  385,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-app",
			Port:  385,
		},
	},
	"ibm-cics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-cics",
			Port:  1435,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-cics",
			Port:  1435,
		},
	},
	"ibm-db2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-db2",
			Port:  523,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-db2",
			Port:  523,
		},
	},
	"ibm-dial-out": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-dial-out",
			Port:  3267,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-dial-out",
			Port:  3267,
		},
	},
	"ibm-diradm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-diradm",
			Port:  3538,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-diradm",
			Port:  3538,
		},
	},
	"ibm-diradm-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-diradm-ssl",
			Port:  3539,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-diradm-ssl",
			Port:  3539,
		},
	},
	"ibm-dt-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-dt-2",
			Port:  1792,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-dt-2",
			Port:  1792,
		},
	},
	"ibm-mgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-mgr",
			Port:  3801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-mgr",
			Port:  3801,
		},
	},
	"ibm-mqisdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-mqisdp",
			Port:  1883,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-mqisdp",
			Port:  1883,
		},
	},
	"ibm-mqseries": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-mqseries",
			Port:  1414,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-mqseries",
			Port:  1414,
		},
	},
	"ibm-mqseries2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-mqseries2",
			Port:  1881,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-mqseries2",
			Port:  1881,
		},
	},
	"ibm-pps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-pps",
			Port:  1376,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-pps",
			Port:  1376,
		},
	},
	"ibm-res": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-res",
			Port:  1405,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-res",
			Port:  1405,
		},
	},
	"ibm-rsyscon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-rsyscon",
			Port:  9085,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-rsyscon",
			Port:  9085,
		},
	},
	"ibm-ssd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm-ssd",
			Port:  1260,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm-ssd",
			Port:  1260,
		},
	},
	"ibm3494": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm3494",
			Port:  3494,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm3494",
			Port:  3494,
		},
	},
	"ibm_wrless_lan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibm_wrless_lan",
			Port:  1461,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibm_wrless_lan",
			Port:  1461,
		},
	},
	"ibp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibp",
			Port:  2572,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibp",
			Port:  2572,
		},
	},
	"ibprotocol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibprotocol",
			Port:  6714,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibprotocol",
			Port:  6714,
		},
	},
	"ibridge-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibridge-data",
			Port:  2275,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibridge-data",
			Port:  2275,
		},
	},
	"ibridge-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibridge-mgmt",
			Port:  2276,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibridge-mgmt",
			Port:  2276,
		},
	},
	"ibus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ibus",
			Port:  8733,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ibus",
			Port:  8733,
		},
	},
	"ica": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ica",
			Port:  1494,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ica",
			Port:  1494,
		},
	},
	"icabrowser": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icabrowser",
			Port:  1604,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icabrowser",
			Port:  1604,
		},
	},
	"icad-el": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icad-el",
			Port:  425,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icad-el",
			Port:  425,
		},
	},
	"icap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icap",
			Port:  1344,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icap",
			Port:  1344,
		},
	},
	"iccrushmore": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iccrushmore",
			Port:  6850,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iccrushmore",
			Port:  6850,
		},
	},
	"ice-location": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ice-location",
			Port:  4061,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ice-location",
			Port:  4061,
		},
	},
	"ice-router": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ice-router",
			Port:  4063,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ice-router",
			Port:  4063,
		},
	},
	"ice-slocation": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ice-slocation",
			Port:  4062,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ice-slocation",
			Port:  4062,
		},
	},
	"ice-srouter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ice-srouter",
			Port:  4064,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ice-srouter",
			Port:  4064,
		},
	},
	"iceedcp_rx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iceedcp_rx",
			Port:  31949,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iceedcp_rx",
			Port:  31949,
		},
	},
	"iceedcp_tx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iceedcp_tx",
			Port:  31948,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iceedcp_tx",
			Port:  31948,
		},
	},
	"icg-bridge": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icg-bridge",
			Port:  2063,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icg-bridge",
			Port:  2063,
		},
	},
	"icg-iprelay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icg-iprelay",
			Port:  2064,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icg-iprelay",
			Port:  2064,
		},
	},
	"icg-swp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icg-swp",
			Port:  2062,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icg-swp",
			Port:  2062,
		},
	},
	"ici": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ici",
			Port:  2200,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ici",
			Port:  2200,
		},
	},
	"icl-twobase1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase1",
			Port:  25000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase1",
			Port:  25000,
		},
	},
	"icl-twobase10": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase10",
			Port:  25009,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase10",
			Port:  25009,
		},
	},
	"icl-twobase2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase2",
			Port:  25001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase2",
			Port:  25001,
		},
	},
	"icl-twobase3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase3",
			Port:  25002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase3",
			Port:  25002,
		},
	},
	"icl-twobase4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase4",
			Port:  25003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase4",
			Port:  25003,
		},
	},
	"icl-twobase5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase5",
			Port:  25004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase5",
			Port:  25004,
		},
	},
	"icl-twobase6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase6",
			Port:  25005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase6",
			Port:  25005,
		},
	},
	"icl-twobase7": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase7",
			Port:  25006,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase7",
			Port:  25006,
		},
	},
	"icl-twobase8": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase8",
			Port:  25007,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase8",
			Port:  25007,
		},
	},
	"icl-twobase9": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icl-twobase9",
			Port:  25008,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icl-twobase9",
			Port:  25008,
		},
	},
	"iclcnet-locate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclcnet-locate",
			Port:  886,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclcnet-locate",
			Port:  886,
		},
	},
	"iclcnet_svinfo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclcnet_svinfo",
			Port:  887,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclcnet_svinfo",
			Port:  887,
		},
	},
	"iclid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclid",
			Port:  18242,
		},
	},
	"iclpv-dm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclpv-dm",
			Port:  1389,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclpv-dm",
			Port:  1389,
		},
	},
	"iclpv-nlc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclpv-nlc",
			Port:  1394,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclpv-nlc",
			Port:  1394,
		},
	},
	"iclpv-nls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclpv-nls",
			Port:  1393,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclpv-nls",
			Port:  1393,
		},
	},
	"iclpv-pm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclpv-pm",
			Port:  1392,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclpv-pm",
			Port:  1392,
		},
	},
	"iclpv-sas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclpv-sas",
			Port:  1391,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclpv-sas",
			Port:  1391,
		},
	},
	"iclpv-sc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclpv-sc",
			Port:  1390,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclpv-sc",
			Port:  1390,
		},
	},
	"iclpv-wsm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iclpv-wsm",
			Port:  1395,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iclpv-wsm",
			Port:  1395,
		},
	},
	"icmpd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icmpd",
			Port:  5813,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icmpd",
			Port:  5813,
		},
	},
	"icms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icms",
			Port:  4486,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icms",
			Port:  4486,
		},
	},
	"icon-discover": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icon-discover",
			Port:  2799,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icon-discover",
			Port:  2799,
		},
	},
	"iconp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iconp",
			Port:  3972,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iconp",
			Port:  3972,
		},
	},
	"iconstructsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iconstructsrv",
			Port:  6077,
		},
	},
	"icp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icp",
			Port:  1112,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icp",
			Port:  1112,
		},
	},
	"icpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icpp",
			Port:  14142,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icpp",
			Port:  14142,
		},
	},
	"icpv2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icpv2",
			Port:  3130,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icpv2",
			Port:  3130,
		},
	},
	"ics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ics",
			Port:  5639,
		},
	},
	"icshostsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icshostsvc",
			Port:  4553,
		},
	},
	"icslap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "icslap",
			Port:  2869,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "icslap",
			Port:  2869,
		},
	},
	"ida-discover1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ida-discover1",
			Port:  5741,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ida-discover1",
			Port:  5741,
		},
	},
	"ida-discover2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ida-discover2",
			Port:  5742,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ida-discover2",
			Port:  5742,
		},
	},
	"idac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idac",
			Port:  3881,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idac",
			Port:  3881,
		},
	},
	"idcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idcp",
			Port:  2326,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idcp",
			Port:  2326,
		},
	},
	"ideafarm-door": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ideafarm-door",
			Port:  902,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ideafarm-door",
			Port:  902,
		},
	},
	"ideafarm-panic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ideafarm-panic",
			Port:  903,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ideafarm-panic",
			Port:  903,
		},
	},
	"ideesrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ideesrv",
			Port:  2337,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ideesrv",
			Port:  2337,
		},
	},
	"iden-ralp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iden-ralp",
			Port:  1725,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iden-ralp",
			Port:  1725,
		},
	},
	"identify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "identify",
			Port:  2987,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "identify",
			Port:  2987,
		},
	},
	"idfp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idfp",
			Port:  549,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idfp",
			Port:  549,
		},
	},
	"idig_mux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idig_mux",
			Port:  4152,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idig_mux",
			Port:  4152,
		},
	},
	"idmaps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idmaps",
			Port:  1884,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idmaps",
			Port:  1884,
		},
	},
	"idmgratm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idmgratm",
			Port:  32896,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idmgratm",
			Port:  32896,
		},
	},
	"idonix-metanet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idonix-metanet",
			Port:  2112,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idonix-metanet",
			Port:  2112,
		},
	},
	"idotdist": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idotdist",
			Port:  2590,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idotdist",
			Port:  2590,
		},
	},
	"idp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idp",
			Port:  4067,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idp",
			Port:  4067,
		},
	},
	"idp-infotrieve": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idp-infotrieve",
			Port:  2966,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idp-infotrieve",
			Port:  2966,
		},
	},
	"idps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idps",
			Port:  3797,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idps",
			Port:  3797,
		},
	},
	"idrs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idrs",
			Port:  2995,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idrs",
			Port:  2995,
		},
	},
	"idtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idtp",
			Port:  25604,
		},
	},
	"idware-router": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idware-router",
			Port:  2079,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idware-router",
			Port:  2079,
		},
	},
	"idxp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "idxp",
			Port:  603,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "idxp",
			Port:  603,
		},
	},
	"iec-104": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iec-104",
			Port:  2404,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iec-104",
			Port:  2404,
		},
	},
	"iec-104-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iec-104-sec",
			Port:  19998,
		},
	},
	"iee-qfx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iee-qfx",
			Port:  1284,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iee-qfx",
			Port:  1284,
		},
	},
	"ieee-mih": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ieee-mih",
			Port:  4551,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ieee-mih",
			Port:  4551,
		},
	},
	"ieee-mms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ieee-mms",
			Port:  651,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ieee-mms",
			Port:  651,
		},
	},
	"ieee-mms-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ieee-mms-ssl",
			Port:  695,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ieee-mms-ssl",
			Port:  695,
		},
	},
	"ies-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ies-lm",
			Port:  1443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ies-lm",
			Port:  1443,
		},
	},
	"ifcp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ifcp-port",
			Port:  3420,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ifcp-port",
			Port:  3420,
		},
	},
	"ife_icorp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ife_icorp",
			Port:  5165,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ife_icorp",
			Port:  5165,
		},
	},
	"ifor-protocol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ifor-protocol",
			Port:  1515,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ifor-protocol",
			Port:  1515,
		},
	},
	"ifsf-hb-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ifsf-hb-port",
			Port:  3486,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ifsf-hb-port",
			Port:  3486,
		},
	},
	"ifsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ifsp",
			Port:  4744,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ifsp",
			Port:  4744,
		},
	},
	"igcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "igcp",
			Port:  2801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "igcp",
			Port:  2801,
		},
	},
	"igi-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "igi-lm",
			Port:  1404,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "igi-lm",
			Port:  1404,
		},
	},
	"igmpv3lite": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "igmpv3lite",
			Port:  465,
		},
	},
	"igo-incognito": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "igo-incognito",
			Port:  4100,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "igo-incognito",
			Port:  4100,
		},
	},
	"igrid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "igrid",
			Port:  19000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "igrid",
			Port:  19000,
		},
	},
	"igrs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "igrs",
			Port:  3880,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "igrs",
			Port:  3880,
		},
	},
	"ii-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ii-admin",
			Port:  3006,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ii-admin",
			Port:  3006,
		},
	},
	"iims": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iims",
			Port:  4800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iims",
			Port:  4800,
		},
	},
	"iiop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iiop",
			Port:  535,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iiop",
			Port:  535,
		},
	},
	"iiw-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iiw-port",
			Port:  3186,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iiw-port",
			Port:  3186,
		},
	},
	"ild": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ild",
			Port:  24321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ild",
			Port:  24321,
		},
	},
	"ill": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ill",
			Port:  1611,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ill",
			Port:  1611,
		},
	},
	"ilss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ilss",
			Port:  4802,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ilss",
			Port:  4802,
		},
	},
	"imagepump": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imagepump",
			Port:  27345,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imagepump",
			Port:  27345,
		},
	},
	"imagequery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imagequery",
			Port:  2239,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imagequery",
			Port:  2239,
		},
	},
	"imap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imap",
			Port:  143,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imap",
			Port:  143,
		},
	},
	"imap3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imap3",
			Port:  220,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imap3",
			Port:  220,
		},
	},
	"imaps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imaps",
			Port:  993,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imaps",
			Port:  993,
		},
	},
	"imdocsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imdocsvc",
			Port:  2637,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imdocsvc",
			Port:  2637,
		},
	},
	"imgames": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imgames",
			Port:  1077,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imgames",
			Port:  1077,
		},
	},
	"imink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imink",
			Port:  8615,
		},
	},
	"imip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imip",
			Port:  11319,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imip",
			Port:  11319,
		},
	},
	"imip-channels": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imip-channels",
			Port:  11320,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imip-channels",
			Port:  11320,
		},
	},
	"immedianet-bcn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "immedianet-bcn",
			Port:  3657,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "immedianet-bcn",
			Port:  3657,
		},
	},
	"imoguia-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imoguia-port",
			Port:  3907,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imoguia-port",
			Port:  3907,
		},
	},
	"impera": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "impera",
			Port:  1710,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "impera",
			Port:  1710,
		},
	},
	"imprs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imprs",
			Port:  3164,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imprs",
			Port:  3164,
		},
	},
	"imqbrokerd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imqbrokerd",
			Port:  7676,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imqbrokerd",
			Port:  7676,
		},
	},
	"imqstomp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imqstomp",
			Port:  7672,
		},
	},
	"imqstomps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imqstomps",
			Port:  7673,
		},
	},
	"imqtunnel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imqtunnel",
			Port:  7675,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imqtunnel",
			Port:  7675,
		},
	},
	"imqtunnels": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imqtunnels",
			Port:  7674,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imqtunnels",
			Port:  7674,
		},
	},
	"imsldoc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imsldoc",
			Port:  2035,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imsldoc",
			Port:  2035,
		},
	},
	"imsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imsp",
			Port:  406,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imsp",
			Port:  406,
		},
	},
	"imtc-map": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imtc-map",
			Port:  2202,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imtc-map",
			Port:  2202,
		},
	},
	"imtc-mcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imtc-mcs",
			Port:  1503,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imtc-mcs",
			Port:  1503,
		},
	},
	"imyx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "imyx",
			Port:  1143,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "imyx",
			Port:  1143,
		},
	},
	"inbusiness": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inbusiness",
			Port:  244,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inbusiness",
			Port:  244,
		},
	},
	"incognitorv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "incognitorv",
			Port:  3139,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "incognitorv",
			Port:  3139,
		},
	},
	"incp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "incp",
			Port:  2932,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "incp",
			Port:  2932,
		},
	},
	"index-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "index-net",
			Port:  2970,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "index-net",
			Port:  2970,
		},
	},
	"index-pc-wb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "index-pc-wb",
			Port:  2127,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "index-pc-wb",
			Port:  2127,
		},
	},
	"indi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "indi",
			Port:  7624,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "indi",
			Port:  7624,
		},
	},
	"indigo-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "indigo-server",
			Port:  1176,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "indigo-server",
			Port:  1176,
		},
	},
	"indigo-vbcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "indigo-vbcp",
			Port:  8131,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "indigo-vbcp",
			Port:  8131,
		},
	},
	"indigo-vrmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "indigo-vrmi",
			Port:  8130,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "indigo-vrmi",
			Port:  8130,
		},
	},
	"indura": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "indura",
			Port:  3156,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "indura",
			Port:  3156,
		},
	},
	"indx-dds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "indx-dds",
			Port:  2454,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "indx-dds",
			Port:  2454,
		},
	},
	"indy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "indy",
			Port:  5963,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "indy",
			Port:  5963,
		},
	},
	"infiniswitchcl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infiniswitchcl",
			Port:  3602,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infiniswitchcl",
			Port:  3602,
		},
	},
	"influence": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "influence",
			Port:  3345,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "influence",
			Port:  3345,
		},
	},
	"infobright": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infobright",
			Port:  5029,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infobright",
			Port:  5029,
		},
	},
	"infocrypt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infocrypt",
			Port:  2233,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infocrypt",
			Port:  2233,
		},
	},
	"infoexch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infoexch",
			Port:  3667,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infoexch",
			Port:  3667,
		},
	},
	"infolibria": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infolibria",
			Port:  2319,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infolibria",
			Port:  2319,
		},
	},
	"infoman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infoman",
			Port:  1451,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infoman",
			Port:  1451,
		},
	},
	"infomover": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infomover",
			Port:  2854,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infomover",
			Port:  2854,
		},
	},
	"informatik-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "informatik-lm",
			Port:  1428,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "informatik-lm",
			Port:  1428,
		},
	},
	"informer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "informer",
			Port:  3856,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "informer",
			Port:  3856,
		},
	},
	"infoseek": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infoseek",
			Port:  414,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infoseek",
			Port:  414,
		},
	},
	"infotos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infotos",
			Port:  18881,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infotos",
			Port:  18881,
		},
	},
	"infowave": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "infowave",
			Port:  2082,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "infowave",
			Port:  2082,
		},
	},
	"ingres-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ingres-net",
			Port:  134,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ingres-net",
			Port:  134,
		},
	},
	"ingreslock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ingreslock",
			Port:  1524,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ingreslock",
			Port:  1524,
		},
	},
	"ininmessaging": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ininmessaging",
			Port:  5597,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ininmessaging",
			Port:  5597,
		},
	},
	"iniserve-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iniserve-port",
			Port:  3560,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iniserve-port",
			Port:  3560,
		},
	},
	"initlsmsad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "initlsmsad",
			Port:  2793,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "initlsmsad",
			Port:  2793,
		},
	},
	"innosys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "innosys",
			Port:  1412,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "innosys",
			Port:  1412,
		},
	},
	"innosys-acl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "innosys-acl",
			Port:  1413,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "innosys-acl",
			Port:  1413,
		},
	},
	"inova-ip-disco": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inova-ip-disco",
			Port:  2716,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inova-ip-disco",
			Port:  2716,
		},
	},
	"inovaport1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inovaport1",
			Port:  23000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inovaport1",
			Port:  23000,
		},
	},
	"inovaport2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inovaport2",
			Port:  23001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inovaport2",
			Port:  23001,
		},
	},
	"inovaport3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inovaport3",
			Port:  23002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inovaport3",
			Port:  23002,
		},
	},
	"inovaport4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inovaport4",
			Port:  23003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inovaport4",
			Port:  23003,
		},
	},
	"inovaport5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inovaport5",
			Port:  23004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inovaport5",
			Port:  23004,
		},
	},
	"inovaport6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inovaport6",
			Port:  23005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inovaport6",
			Port:  23005,
		},
	},
	"insis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "insis",
			Port:  9215,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "insis",
			Port:  9215,
		},
	},
	"insitu-conf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "insitu-conf",
			Port:  1490,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "insitu-conf",
			Port:  1490,
		},
	},
	"inspect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "inspect",
			Port:  1602,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "inspect",
			Port:  1602,
		},
	},
	"inst-discovery": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "inst-discovery",
			Port:  4878,
		},
	},
	"instantia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "instantia",
			Port:  1240,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "instantia",
			Port:  1240,
		},
	},
	"instl_bootc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "instl_bootc",
			Port:  1068,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "instl_bootc",
			Port:  1068,
		},
	},
	"instl_boots": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "instl_boots",
			Port:  1067,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "instl_boots",
			Port:  1067,
		},
	},
	"int-rcv-cntrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "int-rcv-cntrl",
			Port:  3603,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "int-rcv-cntrl",
			Port:  3603,
		},
	},
	"intecom-ps1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intecom-ps1",
			Port:  5056,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intecom-ps1",
			Port:  5056,
		},
	},
	"intecom-ps2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intecom-ps2",
			Port:  5057,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intecom-ps2",
			Port:  5057,
		},
	},
	"intecourier": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intecourier",
			Port:  495,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intecourier",
			Port:  495,
		},
	},
	"integra-sme": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "integra-sme",
			Port:  484,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "integra-sme",
			Port:  484,
		},
	},
	"integral": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "integral",
			Port:  3459,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "integral",
			Port:  3459,
		},
	},
	"integrius-stp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "integrius-stp",
			Port:  17234,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "integrius-stp",
			Port:  17234,
		},
	},
	"intel-rci-mp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intel-rci-mp",
			Port:  16991,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intel-rci-mp",
			Port:  16991,
		},
	},
	"intel_rci": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intel_rci",
			Port:  24386,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intel_rci",
			Port:  24386,
		},
	},
	"intellistor-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intellistor-lm",
			Port:  1539,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intellistor-lm",
			Port:  1539,
		},
	},
	"intelsync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intelsync",
			Port:  3692,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intelsync",
			Port:  3692,
		},
	},
	"interact": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interact",
			Port:  4052,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interact",
			Port:  4052,
		},
	},
	"interactionweb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interactionweb",
			Port:  3508,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interactionweb",
			Port:  3508,
		},
	},
	"interbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interbase",
			Port:  2041,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interbase",
			Port:  2041,
		},
	},
	"interhdl_elmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interhdl_elmd",
			Port:  1454,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interhdl_elmd",
			Port:  1454,
		},
	},
	"interintelli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interintelli",
			Port:  2633,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interintelli",
			Port:  2633,
		},
	},
	"intermapper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intermapper",
			Port:  8181,
		},
	},
	"interpathpanel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interpathpanel",
			Port:  2652,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interpathpanel",
			Port:  2652,
		},
	},
	"intersan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intersan",
			Port:  1331,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intersan",
			Port:  1331,
		},
	},
	"interserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interserver",
			Port:  3060,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interserver",
			Port:  3060,
		},
	},
	"intersys-cache": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intersys-cache",
			Port:  1972,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intersys-cache",
			Port:  1972,
		},
	},
	"interwise": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interwise",
			Port:  7778,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interwise",
			Port:  7778,
		},
	},
	"interworld": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "interworld",
			Port:  3548,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "interworld",
			Port:  3548,
		},
	},
	"intraintra": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intraintra",
			Port:  3202,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intraintra",
			Port:  3202,
		},
	},
	"intrastar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intrastar",
			Port:  1907,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intrastar",
			Port:  1907,
		},
	},
	"intrepid-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intrepid-ssl",
			Port:  11751,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intrepid-ssl",
			Port:  11751,
		},
	},
	"intrinsa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intrinsa",
			Port:  503,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intrinsa",
			Port:  503,
		},
	},
	"intu-ec-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intu-ec-client",
			Port:  8021,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intu-ec-client",
			Port:  8021,
		},
	},
	"intu-ec-svcdisc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intu-ec-svcdisc",
			Port:  8020,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intu-ec-svcdisc",
			Port:  8020,
		},
	},
	"intuitive-edge": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intuitive-edge",
			Port:  1355,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intuitive-edge",
			Port:  1355,
		},
	},
	"intv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "intv",
			Port:  1585,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "intv",
			Port:  1585,
		},
	},
	"invision": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "invision",
			Port:  1641,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "invision",
			Port:  1641,
		},
	},
	"invision-ag": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "invision-ag",
			Port:  45054,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "invision-ag",
			Port:  45054,
		},
	},
	"invokator": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "invokator",
			Port:  2006,
		},
	},
	"io-dist-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "io-dist-data",
			Port:  5728,
		},
	},
	"io-dist-group": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "io-dist-group",
			Port:  5728,
		},
	},
	"ioc-sea-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ioc-sea-lm",
			Port:  1579,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ioc-sea-lm",
			Port:  1579,
		},
	},
	"ionixnetmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ionixnetmon",
			Port:  7410,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ionixnetmon",
			Port:  7410,
		},
	},
	"iop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iop",
			Port:  2055,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iop",
			Port:  2055,
		},
	},
	"ip-blf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ip-blf",
			Port:  2088,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ip-blf",
			Port:  2088,
		},
	},
	"ip-provision": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ip-provision",
			Port:  43190,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ip-provision",
			Port:  43190,
		},
	},
	"ip-qsig": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ip-qsig",
			Port:  4029,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ip-qsig",
			Port:  4029,
		},
	},
	"ipass": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipass",
			Port:  2549,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipass",
			Port:  2549,
		},
	},
	"ipcd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipcd",
			Port:  576,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipcd",
			Port:  576,
		},
	},
	"ipcd3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipcd3",
			Port:  1209,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipcd3",
			Port:  1209,
		},
	},
	"ipcore": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipcore",
			Port:  2215,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipcore",
			Port:  2215,
		},
	},
	"ipcs-command": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipcs-command",
			Port:  3743,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipcs-command",
			Port:  3743,
		},
	},
	"ipcserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipcserver",
			Port:  600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipcserver",
			Port:  600,
		},
	},
	"ipdcesgbs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipdcesgbs",
			Port:  9214,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipdcesgbs",
			Port:  9214,
		},
	},
	"ipdd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipdd",
			Port:  578,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipdd",
			Port:  578,
		},
	},
	"ipdr-sp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipdr-sp",
			Port:  4737,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipdr-sp",
			Port:  4737,
		},
	},
	"ipdtp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipdtp-port",
			Port:  20202,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipdtp-port",
			Port:  20202,
		},
	},
	"ipether232port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipether232port",
			Port:  3497,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipether232port",
			Port:  3497,
		},
	},
	"ipfix": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "ipfix",
			Port:  4739,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipfix",
			Port:  4739,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipfix",
			Port:  4739,
		},
	},
	"ipfixs": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "ipfixs",
			Port:  4740,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipfixs",
			Port:  4740,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipfixs",
			Port:  4740,
		},
	},
	"ipfltbcst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipfltbcst",
			Port:  4068,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipfltbcst",
			Port:  4068,
		},
	},
	"iph-policy-adm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iph-policy-adm",
			Port:  2963,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iph-policy-adm",
			Port:  2963,
		},
	},
	"iph-policy-cli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iph-policy-cli",
			Port:  2962,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iph-policy-cli",
			Port:  2962,
		},
	},
	"iposplanet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iposplanet",
			Port:  7031,
		},
	},
	"ipp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipp",
			Port:  631,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipp",
			Port:  631,
		},
	},
	"ipr-dglt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipr-dglt",
			Port:  3678,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipr-dglt",
			Port:  3678,
		},
	},
	"ipsec-nat-t": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipsec-nat-t",
			Port:  4500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipsec-nat-t",
			Port:  4500,
		},
	},
	"ipt-anri-anri": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipt-anri-anri",
			Port:  4593,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipt-anri-anri",
			Port:  4593,
		},
	},
	"ipulse-ics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipulse-ics",
			Port:  20222,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipulse-ics",
			Port:  20222,
		},
	},
	"ipx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ipx",
			Port:  213,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ipx",
			Port:  213,
		},
	},
	"iqnet-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iqnet-port",
			Port:  3804,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iqnet-port",
			Port:  3804,
		},
	},
	"iqobject": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iqobject",
			Port:  48619,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iqobject",
			Port:  48619,
		},
	},
	"iqrm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iqrm",
			Port:  10117,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iqrm",
			Port:  10117,
		},
	},
	"iqserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iqserver",
			Port:  2527,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iqserver",
			Port:  2527,
		},
	},
	"ique": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ique",
			Port:  18769,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ique",
			Port:  18769,
		},
	},
	"iracinghelper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iracinghelper",
			Port:  32034,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iracinghelper",
			Port:  32034,
		},
	},
	"irc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "irc",
			Port:  194,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "irc",
			Port:  194,
		},
	},
	"irc-serv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "irc-serv",
			Port:  529,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "irc-serv",
			Port:  529,
		},
	},
	"ircu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ircu",
			Port:  6665,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ircu",
			Port:  6665,
		},
	},
	"ircu-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ircu-2",
			Port:  6666,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ircu-2",
			Port:  6666,
		},
	},
	"ircu-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ircu-3",
			Port:  6667,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ircu-3",
			Port:  6667,
		},
	},
	"ircu-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ircu-4",
			Port:  6668,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ircu-4",
			Port:  6668,
		},
	},
	"ircu-5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ircu-5",
			Port:  6669,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ircu-5",
			Port:  6669,
		},
	},
	"irdg-post": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "irdg-post",
			Port:  2632,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "irdg-post",
			Port:  2632,
		},
	},
	"irdmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "irdmi",
			Port:  8000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "irdmi",
			Port:  8000,
		},
	},
	"irdmi2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "irdmi2",
			Port:  7999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "irdmi2",
			Port:  7999,
		},
	},
	"iris-beep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iris-beep",
			Port:  702,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iris-beep",
			Port:  702,
		},
	},
	"iris-lwz": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iris-lwz",
			Port:  715,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iris-lwz",
			Port:  715,
		},
	},
	"iris-xpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iris-xpc",
			Port:  713,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iris-xpc",
			Port:  713,
		},
	},
	"iris-xpcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iris-xpcs",
			Port:  714,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iris-xpcs",
			Port:  714,
		},
	},
	"irisa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "irisa",
			Port:  11000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "irisa",
			Port:  11000,
		},
	},
	"ironmail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ironmail",
			Port:  3206,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ironmail",
			Port:  3206,
		},
	},
	"ironstorm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ironstorm",
			Port:  3504,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ironstorm",
			Port:  3504,
		},
	},
	"irtrans": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "irtrans",
			Port:  21000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "irtrans",
			Port:  21000,
		},
	},
	"is99c": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "is99c",
			Port:  379,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "is99c",
			Port:  379,
		},
	},
	"is99s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "is99s",
			Port:  380,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "is99s",
			Port:  380,
		},
	},
	"isakmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isakmp",
			Port:  500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isakmp",
			Port:  500,
		},
	},
	"isbconference1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isbconference1",
			Port:  1244,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isbconference1",
			Port:  1244,
		},
	},
	"isbconference2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isbconference2",
			Port:  1245,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isbconference2",
			Port:  1245,
		},
	},
	"iscape": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "iscape",
			Port:  5047,
		},
	},
	"ischat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ischat",
			Port:  1336,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ischat",
			Port:  1336,
		},
	},
	"iscsi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iscsi",
			Port:  860,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iscsi",
			Port:  860,
		},
	},
	"iscsi-target": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iscsi-target",
			Port:  3260,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iscsi-target",
			Port:  3260,
		},
	},
	"isdc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isdc",
			Port:  1636,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isdc",
			Port:  1636,
		},
	},
	"isdd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isdd",
			Port:  8148,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isdd",
			Port:  8148,
		},
	},
	"isdnlog": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isdnlog",
			Port:  20011,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isdnlog",
			Port:  20011,
		},
	},
	"isg-uda-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isg-uda-server",
			Port:  2551,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isg-uda-server",
			Port:  2551,
		},
	},
	"isi-gl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isi-gl",
			Port:  55,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isi-gl",
			Port:  55,
		},
	},
	"isi-irp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isi-irp",
			Port:  3226,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isi-irp",
			Port:  3226,
		},
	},
	"isis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isis",
			Port:  2042,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isis",
			Port:  2042,
		},
	},
	"isis-am": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isis-am",
			Port:  1642,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isis-am",
			Port:  1642,
		},
	},
	"isis-ambc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isis-ambc",
			Port:  1643,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isis-ambc",
			Port:  1643,
		},
	},
	"isis-bcast": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isis-bcast",
			Port:  2043,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isis-bcast",
			Port:  2043,
		},
	},
	"islc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "islc",
			Port:  1637,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "islc",
			Port:  1637,
		},
	},
	"ismaeasdaqlive": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ismaeasdaqlive",
			Port:  1949,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ismaeasdaqlive",
			Port:  1949,
		},
	},
	"ismaeasdaqtest": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ismaeasdaqtest",
			Port:  1950,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ismaeasdaqtest",
			Port:  1950,
		},
	},
	"ismc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ismc",
			Port:  1638,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ismc",
			Port:  1638,
		},
	},
	"ismserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ismserver",
			Port:  9500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ismserver",
			Port:  9500,
		},
	},
	"isnetserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isnetserv",
			Port:  48128,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isnetserv",
			Port:  48128,
		},
	},
	"isns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isns",
			Port:  3205,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isns",
			Port:  3205,
		},
	},
	"iso-ill": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iso-ill",
			Port:  499,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iso-ill",
			Port:  499,
		},
	},
	"iso-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iso-ip",
			Port:  147,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iso-ip",
			Port:  147,
		},
	},
	"iso-tp0": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iso-tp0",
			Port:  146,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iso-tp0",
			Port:  146,
		},
	},
	"iso-tp0s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iso-tp0s",
			Port:  3782,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iso-tp0s",
			Port:  3782,
		},
	},
	"iso-tsap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iso-tsap",
			Port:  102,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iso-tsap",
			Port:  102,
		},
	},
	"iso-tsap-c2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iso-tsap-c2",
			Port:  399,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iso-tsap-c2",
			Port:  399,
		},
	},
	"isode-dua": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isode-dua",
			Port:  17007,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isode-dua",
			Port:  17007,
		},
	},
	"isoft-p2p": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isoft-p2p",
			Port:  3501,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isoft-p2p",
			Port:  3501,
		},
	},
	"isoipsigport-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isoipsigport-1",
			Port:  1106,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isoipsigport-1",
			Port:  1106,
		},
	},
	"isoipsigport-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isoipsigport-2",
			Port:  1107,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isoipsigport-2",
			Port:  1107,
		},
	},
	"isomair": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isomair",
			Port:  3589,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isomair",
			Port:  3589,
		},
	},
	"ispipes": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ispipes",
			Port:  2853,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ispipes",
			Port:  2853,
		},
	},
	"ispmmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ispmmgr",
			Port:  3775,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ispmmgr",
			Port:  3775,
		},
	},
	"isrp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isrp-port",
			Port:  3788,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isrp-port",
			Port:  3788,
		},
	},
	"iss-mgmt-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iss-mgmt-ssl",
			Port:  3995,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iss-mgmt-ssl",
			Port:  3995,
		},
	},
	"issd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "issd",
			Port:  1600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "issd",
			Port:  1600,
		},
	},
	"isysg-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "isysg-lm",
			Port:  1609,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "isysg-lm",
			Port:  1609,
		},
	},
	"ita-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ita-agent",
			Port:  5051,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ita-agent",
			Port:  5051,
		},
	},
	"ita-manager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ita-manager",
			Port:  5052,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ita-manager",
			Port:  5052,
		},
	},
	"itach": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itach",
			Port:  8184,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itach",
			Port:  8184,
		},
	},
	"itactionserver1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itactionserver1",
			Port:  7280,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itactionserver1",
			Port:  7280,
		},
	},
	"itactionserver2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itactionserver2",
			Port:  7281,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itactionserver2",
			Port:  7281,
		},
	},
	"italk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "italk",
			Port:  12345,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "italk",
			Port:  12345,
		},
	},
	"itap-ddtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itap-ddtp",
			Port:  10100,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itap-ddtp",
			Port:  10100,
		},
	},
	"itelserverport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itelserverport",
			Port:  3719,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itelserverport",
			Port:  3719,
		},
	},
	"item": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "item",
			Port:  3848,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "item",
			Port:  3848,
		},
	},
	"itinternet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itinternet",
			Port:  2691,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itinternet",
			Port:  2691,
		},
	},
	"itm-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itm-lm",
			Port:  2828,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itm-lm",
			Port:  2828,
		},
	},
	"itm-mccs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itm-mccs",
			Port:  3084,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itm-mccs",
			Port:  3084,
		},
	},
	"itm-mcell-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itm-mcell-s",
			Port:  828,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itm-mcell-s",
			Port:  828,
		},
	},
	"itm-mcell-u": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itm-mcell-u",
			Port:  1828,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itm-mcell-u",
			Port:  1828,
		},
	},
	"ito-e-gui": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ito-e-gui",
			Port:  2531,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ito-e-gui",
			Port:  2531,
		},
	},
	"itose": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itose",
			Port:  4348,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itose",
			Port:  4348,
		},
	},
	"itscomm-ns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itscomm-ns",
			Port:  1573,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itscomm-ns",
			Port:  1573,
		},
	},
	"itu-bicc-stc": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "itu-bicc-stc",
			Port:  3097,
		},
	},
	"itv-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itv-control",
			Port:  3899,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "itv-control",
			Port:  3899,
		},
	},
	"itwo-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "itwo-server",
			Port:  4410,
		},
	},
	"iua": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "iua",
			Port:  9900,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "iua",
			Port:  9900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iua",
			Port:  9900,
		},
	},
	"iuhsctpassoc": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "iuhsctpassoc",
			Port:  29169,
		},
	},
	"ivcollector": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ivcollector",
			Port:  1275,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ivcollector",
			Port:  1275,
		},
	},
	"ivecon-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ivecon-port",
			Port:  3258,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ivecon-port",
			Port:  3258,
		},
	},
	"ivmanager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ivmanager",
			Port:  1276,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ivmanager",
			Port:  1276,
		},
	},
	"ivocalize": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ivocalize",
			Port:  5049,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ivocalize",
			Port:  5049,
		},
	},
	"ivs-video": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ivs-video",
			Port:  2232,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ivs-video",
			Port:  2232,
		},
	},
	"ivsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ivsd",
			Port:  2241,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ivsd",
			Port:  2241,
		},
	},
	"iw-mmogame": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iw-mmogame",
			Port:  3596,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iw-mmogame",
			Port:  3596,
		},
	},
	"iwb-whiteboard": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iwb-whiteboard",
			Port:  2982,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iwb-whiteboard",
			Port:  2982,
		},
	},
	"iwec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iwec",
			Port:  4801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iwec",
			Port:  4801,
		},
	},
	"iwg1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iwg1",
			Port:  7071,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iwg1",
			Port:  7071,
		},
	},
	"iwlistener": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iwlistener",
			Port:  2866,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iwlistener",
			Port:  2866,
		},
	},
	"iwserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "iwserver",
			Port:  2166,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "iwserver",
			Port:  2166,
		},
	},
	"izm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "izm",
			Port:  4109,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "izm",
			Port:  4109,
		},
	},
	"j-ac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "j-ac",
			Port:  4107,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "j-ac",
			Port:  4107,
		},
	},
	"j-lan-p": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "j-lan-p",
			Port:  2808,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "j-lan-p",
			Port:  2808,
		},
	},
	"j-link": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "j-link",
			Port:  19020,
		},
	},
	"jacobus-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jacobus-lm",
			Port:  1578,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jacobus-lm",
			Port:  1578,
		},
	},
	"jaleosnd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jaleosnd",
			Port:  1623,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jaleosnd",
			Port:  1623,
		},
	},
	"jamlink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jamlink",
			Port:  8091,
		},
	},
	"jamserverport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jamserverport",
			Port:  3627,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jamserverport",
			Port:  3627,
		},
	},
	"jargon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jargon",
			Port:  148,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jargon",
			Port:  148,
		},
	},
	"jaugsremotec-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jaugsremotec-1",
			Port:  3472,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jaugsremotec-1",
			Port:  3472,
		},
	},
	"jaugsremotec-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jaugsremotec-2",
			Port:  3473,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jaugsremotec-2",
			Port:  3473,
		},
	},
	"jaus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jaus",
			Port:  3794,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jaus",
			Port:  3794,
		},
	},
	"jaxer-manager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jaxer-manager",
			Port:  4328,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jaxer-manager",
			Port:  4328,
		},
	},
	"jaxer-web": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jaxer-web",
			Port:  4327,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jaxer-web",
			Port:  4327,
		},
	},
	"jboss-iiop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jboss-iiop",
			Port:  3528,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jboss-iiop",
			Port:  3528,
		},
	},
	"jboss-iiop-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jboss-iiop-ssl",
			Port:  3529,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jboss-iiop-ssl",
			Port:  3529,
		},
	},
	"jbroker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jbroker",
			Port:  2506,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jbroker",
			Port:  2506,
		},
	},
	"jcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jcp",
			Port:  19541,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jcp",
			Port:  19541,
		},
	},
	"jdatastore": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jdatastore",
			Port:  2508,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jdatastore",
			Port:  2508,
		},
	},
	"jdl-dbkitchen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jdl-dbkitchen",
			Port:  3086,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jdl-dbkitchen",
			Port:  3086,
		},
	},
	"jdmn-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jdmn-port",
			Port:  4030,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jdmn-port",
			Port:  4030,
		},
	},
	"jdp-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "jdp-disc",
			Port:  7095,
		},
	},
	"jediserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jediserver",
			Port:  2406,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jediserver",
			Port:  2406,
		},
	},
	"jeol-nsddp-1": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "jeol-nsddp-1",
			Port:  6241,
		},
	},
	"jeol-nsddp-2": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "jeol-nsddp-2",
			Port:  6242,
		},
	},
	"jeol-nsddp-3": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "jeol-nsddp-3",
			Port:  6243,
		},
	},
	"jeol-nsddp-4": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "jeol-nsddp-4",
			Port:  6244,
		},
	},
	"jeol-nsdtp-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jeol-nsdtp-1",
			Port:  6241,
		},
	},
	"jeol-nsdtp-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jeol-nsdtp-2",
			Port:  6242,
		},
	},
	"jeol-nsdtp-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jeol-nsdtp-3",
			Port:  6243,
		},
	},
	"jeol-nsdtp-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jeol-nsdtp-4",
			Port:  6244,
		},
	},
	"jerand-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jerand-lm",
			Port:  1810,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jerand-lm",
			Port:  1810,
		},
	},
	"jesmsjc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jesmsjc",
			Port:  27442,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jesmsjc",
			Port:  27442,
		},
	},
	"jetcmeserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jetcmeserver",
			Port:  1936,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jetcmeserver",
			Port:  1936,
		},
	},
	"jetdirect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jetdirect",
			Port:  9100,
		},
	},
	"jetform": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jetform",
			Port:  1706,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jetform",
			Port:  1706,
		},
	},
	"jetformpreview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jetformpreview",
			Port:  2097,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jetformpreview",
			Port:  2097,
		},
	},
	"jetstream": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jetstream",
			Port:  6901,
		},
	},
	"jibe-eb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jibe-eb",
			Port:  3777,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jibe-eb",
			Port:  3777,
		},
	},
	"jini-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jini-discovery",
			Port:  4160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jini-discovery",
			Port:  4160,
		},
	},
	"jlicelmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jlicelmd",
			Port:  1567,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jlicelmd",
			Port:  1567,
		},
	},
	"jmact3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jmact3",
			Port:  6961,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jmact3",
			Port:  6961,
		},
	},
	"jmact5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jmact5",
			Port:  2957,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jmact5",
			Port:  2957,
		},
	},
	"jmact6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jmact6",
			Port:  2958,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jmact6",
			Port:  2958,
		},
	},
	"jmb-cds1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jmb-cds1",
			Port:  8900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jmb-cds1",
			Port:  8900,
		},
	},
	"jmb-cds2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jmb-cds2",
			Port:  8901,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jmb-cds2",
			Port:  8901,
		},
	},
	"jmevt2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jmevt2",
			Port:  6962,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jmevt2",
			Port:  6962,
		},
	},
	"jmq-daemon-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jmq-daemon-1",
			Port:  3214,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jmq-daemon-1",
			Port:  3214,
		},
	},
	"jmq-daemon-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jmq-daemon-2",
			Port:  3215,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jmq-daemon-2",
			Port:  3215,
		},
	},
	"jms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jms",
			Port:  5673,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jms",
			Port:  5673,
		},
	},
	"joaJewelSuite": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "joaJewelSuite",
			Port:  6583,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "joaJewelSuite",
			Port:  6583,
		},
	},
	"joltid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "joltid",
			Port:  3531,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "joltid",
			Port:  3531,
		},
	},
	"jomamqmonitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jomamqmonitor",
			Port:  4114,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jomamqmonitor",
			Port:  4114,
		},
	},
	"joost": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "joost",
			Port:  4166,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "joost",
			Port:  4166,
		},
	},
	"journee": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "journee",
			Port:  3042,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "journee",
			Port:  3042,
		},
	},
	"jpegmpeg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jpegmpeg",
			Port:  3155,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jpegmpeg",
			Port:  3155,
		},
	},
	"jprinter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jprinter",
			Port:  5309,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jprinter",
			Port:  5309,
		},
	},
	"jps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jps",
			Port:  2205,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jps",
			Port:  2205,
		},
	},
	"jstel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jstel",
			Port:  1064,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jstel",
			Port:  1064,
		},
	},
	"jt400": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jt400",
			Port:  3470,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jt400",
			Port:  3470,
		},
	},
	"jt400-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jt400-ssl",
			Port:  3471,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jt400-ssl",
			Port:  3471,
		},
	},
	"jtag-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jtag-server",
			Port:  1309,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jtag-server",
			Port:  1309,
		},
	},
	"jute": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jute",
			Port:  5883,
		},
	},
	"juxml-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "juxml-port",
			Port:  3642,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "juxml-port",
			Port:  3642,
		},
	},
	"jvclient": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jvclient",
			Port:  1940,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jvclient",
			Port:  1940,
		},
	},
	"jvl-mactalk": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "jvl-mactalk",
			Port:  47100,
		},
	},
	"jvserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jvserver",
			Port:  1939,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jvserver",
			Port:  1939,
		},
	},
	"jwalkserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jwalkserver",
			Port:  1289,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jwalkserver",
			Port:  1289,
		},
	},
	"jwclient": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jwclient",
			Port:  1938,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jwclient",
			Port:  1938,
		},
	},
	"jwpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jwpc",
			Port:  16020,
		},
	},
	"jwpc-bin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jwpc-bin",
			Port:  16021,
		},
	},
	"jwserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "jwserver",
			Port:  1937,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "jwserver",
			Port:  1937,
		},
	},
	"k-block": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "k-block",
			Port:  287,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "k-block",
			Port:  287,
		},
	},
	"k3software-cli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "k3software-cli",
			Port:  26263,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "k3software-cli",
			Port:  26263,
		},
	},
	"k3software-svr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "k3software-svr",
			Port:  26262,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "k3software-svr",
			Port:  26262,
		},
	},
	"ka0wuc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ka0wuc",
			Port:  2822,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ka0wuc",
			Port:  2822,
		},
	},
	"kali": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kali",
			Port:  2213,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kali",
			Port:  2213,
		},
	},
	"kamanda": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kamanda",
			Port:  10081,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kamanda",
			Port:  10081,
		},
	},
	"kana": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kana",
			Port:  2656,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kana",
			Port:  2656,
		},
	},
	"kar2ouche": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kar2ouche",
			Port:  4661,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kar2ouche",
			Port:  4661,
		},
	},
	"kastenchasepad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kastenchasepad",
			Port:  2918,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kastenchasepad",
			Port:  2918,
		},
	},
	"kastenxpipe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kastenxpipe",
			Port:  36865,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kastenxpipe",
			Port:  36865,
		},
	},
	"kazaa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kazaa",
			Port:  1214,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kazaa",
			Port:  1214,
		},
	},
	"kca-service": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "kca-service",
			Port:  9878,
		},
	},
	"kdm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kdm",
			Port:  2115,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kdm",
			Port:  2115,
		},
	},
	"kentrox-prot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kentrox-prot",
			Port:  2502,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kentrox-prot",
			Port:  2502,
		},
	},
	"kerberos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kerberos",
			Port:  88,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kerberos",
			Port:  88,
		},
	},
	"kerberos-adm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kerberos-adm",
			Port:  749,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kerberos-adm",
			Port:  749,
		},
	},
	"kerberos-iv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kerberos-iv",
			Port:  750,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kerberos-iv",
			Port:  750,
		},
	},
	"kerberos_master": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kerberos_master",
			Port:  751,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kerberos_master",
			Port:  751,
		},
	},
	"kermit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kermit",
			Port:  1649,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kermit",
			Port:  1649,
		},
	},
	"keyserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "keyserver",
			Port:  584,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "keyserver",
			Port:  584,
		},
	},
	"keyshadow": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "keyshadow",
			Port:  19315,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "keyshadow",
			Port:  19315,
		},
	},
	"keysrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "keysrvr",
			Port:  19283,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "keysrvr",
			Port:  19283,
		},
	},
	"kfserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kfserver",
			Port:  5343,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kfserver",
			Port:  5343,
		},
	},
	"kftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kftp",
			Port:  6621,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kftp",
			Port:  6621,
		},
	},
	"kftp-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kftp-data",
			Port:  6620,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kftp-data",
			Port:  6620,
		},
	},
	"kfxaclicensing": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kfxaclicensing",
			Port:  3581,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kfxaclicensing",
			Port:  3581,
		},
	},
	"kingdomsonline": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kingdomsonline",
			Port:  30260,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kingdomsonline",
			Port:  30260,
		},
	},
	"kingfisher": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kingfisher",
			Port:  4058,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kingfisher",
			Port:  4058,
		},
	},
	"kink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kink",
			Port:  910,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kink",
			Port:  910,
		},
	},
	"kiosk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kiosk",
			Port:  1061,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kiosk",
			Port:  1061,
		},
	},
	"kis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kis",
			Port:  186,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kis",
			Port:  186,
		},
	},
	"kitim": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kitim",
			Port:  35354,
		},
	},
	"kjtsiteserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kjtsiteserver",
			Port:  1339,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kjtsiteserver",
			Port:  1339,
		},
	},
	"klio": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "klio",
			Port:  7697,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "klio",
			Port:  7697,
		},
	},
	"klogin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "klogin",
			Port:  543,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "klogin",
			Port:  543,
		},
	},
	"kme-trap-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kme-trap-port",
			Port:  2081,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kme-trap-port",
			Port:  2081,
		},
	},
	"kmip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kmip",
			Port:  5696,
		},
	},
	"kmscontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kmscontrol",
			Port:  1773,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kmscontrol",
			Port:  1773,
		},
	},
	"knet-cmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "knet-cmp",
			Port:  157,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "knet-cmp",
			Port:  157,
		},
	},
	"knetd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "knetd",
			Port:  2053,
		},
	},
	"kofax-svr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kofax-svr",
			Port:  2424,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kofax-svr",
			Port:  2424,
		},
	},
	"konshus-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "konshus-lm",
			Port:  2294,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "konshus-lm",
			Port:  2294,
		},
	},
	"konspire2b": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "konspire2b",
			Port:  6085,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "konspire2b",
			Port:  6085,
		},
	},
	"kopek-httphead": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kopek-httphead",
			Port:  27504,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kopek-httphead",
			Port:  27504,
		},
	},
	"kpasswd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kpasswd",
			Port:  464,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kpasswd",
			Port:  464,
		},
	},
	"kpdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kpdp",
			Port:  5253,
		},
	},
	"kpn-icw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kpn-icw",
			Port:  3699,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kpn-icw",
			Port:  3699,
		},
	},
	"kpop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kpop",
			Port:  1109,
		},
	},
	"krb524": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "krb524",
			Port:  4444,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "krb524",
			Port:  4444,
		},
	},
	"krb5_prop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "krb5_prop",
			Port:  754,
		},
	},
	"krb5gatekeeper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "krb5gatekeeper",
			Port:  1318,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "krb5gatekeeper",
			Port:  1318,
		},
	},
	"krbupdate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "krbupdate",
			Port:  760,
		},
	},
	"kryptolan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kryptolan",
			Port:  398,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kryptolan",
			Port:  398,
		},
	},
	"kshell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kshell",
			Port:  544,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kshell",
			Port:  544,
		},
	},
	"ksysguard": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ksysguard",
			Port:  3112,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ksysguard",
			Port:  3112,
		},
	},
	"ktelnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ktelnet",
			Port:  6623,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ktelnet",
			Port:  6623,
		},
	},
	"kti-icad-srvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kti-icad-srvr",
			Port:  6701,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kti-icad-srvr",
			Port:  6701,
		},
	},
	"kv-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kv-agent",
			Port:  3361,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kv-agent",
			Port:  3361,
		},
	},
	"kv-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kv-server",
			Port:  3360,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kv-server",
			Port:  3360,
		},
	},
	"kvm-via-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kvm-via-ip",
			Port:  1132,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kvm-via-ip",
			Port:  1132,
		},
	},
	"kwdb-commn": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "kwdb-commn",
			Port:  1127,
		},
	},
	"kwtc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kwtc",
			Port:  4566,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kwtc",
			Port:  4566,
		},
	},
	"kyoceranetdev": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "kyoceranetdev",
			Port:  1063,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "kyoceranetdev",
			Port:  1063,
		},
	},
	"l-acoustics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l-acoustics",
			Port:  4432,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l-acoustics",
			Port:  4432,
		},
	},
	"l2c-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l2c-control",
			Port:  4371,
		},
	},
	"l2c-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l2c-data",
			Port:  4372,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l2c-data",
			Port:  4372,
		},
	},
	"l2c-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "l2c-disc",
			Port:  4371,
		},
	},
	"l2tp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l2tp",
			Port:  1701,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l2tp",
			Port:  1701,
		},
	},
	"l3-exprt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l3-exprt",
			Port:  2840,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l3-exprt",
			Port:  2840,
		},
	},
	"l3-hawk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l3-hawk",
			Port:  2842,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l3-hawk",
			Port:  2842,
		},
	},
	"l3-hbmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l3-hbmon",
			Port:  2370,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l3-hbmon",
			Port:  2370,
		},
	},
	"l3-ranger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l3-ranger",
			Port:  2841,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l3-ranger",
			Port:  2841,
		},
	},
	"l3t-at-an": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l3t-at-an",
			Port:  4591,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l3t-at-an",
			Port:  4591,
		},
	},
	"l5nas-parchan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "l5nas-parchan",
			Port:  9747,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "l5nas-parchan",
			Port:  9747,
		},
	},
	"la-maint": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "la-maint",
			Port:  51,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "la-maint",
			Port:  51,
		},
	},
	"labrat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "labrat",
			Port:  2560,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "labrat",
			Port:  2560,
		},
	},
	"laes-bf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "laes-bf",
			Port:  9536,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "laes-bf",
			Port:  9536,
		},
	},
	"lam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lam",
			Port:  2040,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lam",
			Port:  2040,
		},
	},
	"lan900_remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lan900_remote",
			Port:  2395,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lan900_remote",
			Port:  2395,
		},
	},
	"landmarks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "landmarks",
			Port:  3969,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "landmarks",
			Port:  3969,
		},
	},
	"lanmessenger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lanmessenger",
			Port:  2372,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lanmessenger",
			Port:  2372,
		},
	},
	"lanner-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lanner-lm",
			Port:  4547,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lanner-lm",
			Port:  4547,
		},
	},
	"lanrevagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lanrevagent",
			Port:  3970,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lanrevagent",
			Port:  3970,
		},
	},
	"lanrevserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lanrevserver",
			Port:  3971,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lanrevserver",
			Port:  3971,
		},
	},
	"lanschool": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lanschool",
			Port:  11796,
		},
	},
	"lanschool-mpt": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "lanschool-mpt",
			Port:  11796,
		},
	},
	"lanserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lanserver",
			Port:  637,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lanserver",
			Port:  637,
		},
	},
	"lansource": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lansource",
			Port:  1485,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lansource",
			Port:  1485,
		},
	},
	"lansurveyor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lansurveyor",
			Port:  4347,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lansurveyor",
			Port:  4347,
		},
	},
	"lansurveyorxml": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lansurveyorxml",
			Port:  3815,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lansurveyorxml",
			Port:  3815,
		},
	},
	"lanyon-lantern": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lanyon-lantern",
			Port:  1682,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lanyon-lantern",
			Port:  1682,
		},
	},
	"laplink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "laplink",
			Port:  1547,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "laplink",
			Port:  1547,
		},
	},
	"launchbird-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "launchbird-lm",
			Port:  3739,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "launchbird-lm",
			Port:  3739,
		},
	},
	"lavenir-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lavenir-lm",
			Port:  3373,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lavenir-lm",
			Port:  3373,
		},
	},
	"lazy-ptop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lazy-ptop",
			Port:  7099,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lazy-ptop",
			Port:  7099,
		},
	},
	"lbc-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lbc-control",
			Port:  2780,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lbc-control",
			Port:  2780,
		},
	},
	"lbc-measure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lbc-measure",
			Port:  2815,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lbc-measure",
			Port:  2815,
		},
	},
	"lbc-sync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lbc-sync",
			Port:  2779,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lbc-sync",
			Port:  2779,
		},
	},
	"lbc-watchdog": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lbc-watchdog",
			Port:  2816,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lbc-watchdog",
			Port:  2816,
		},
	},
	"lbf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lbf",
			Port:  2466,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lbf",
			Port:  2466,
		},
	},
	"lbm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lbm",
			Port:  2465,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lbm",
			Port:  2465,
		},
	},
	"lcm-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lcm-server",
			Port:  7365,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lcm-server",
			Port:  7365,
		},
	},
	"lcs-ap": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "lcs-ap",
			Port:  9082,
		},
	},
	"ldap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldap",
			Port:  389,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ldap",
			Port:  389,
		},
	},
	"ldap-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldap-admin",
			Port:  3407,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ldap-admin",
			Port:  3407,
		},
	},
	"ldaps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldaps",
			Port:  636,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ldaps",
			Port:  636,
		},
	},
	"ldgateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldgateway",
			Port:  9592,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ldgateway",
			Port:  9592,
		},
	},
	"ldoms-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldoms-mgmt",
			Port:  6482,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ldoms-mgmt",
			Port:  6482,
		},
	},
	"ldoms-migr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldoms-migr",
			Port:  8101,
		},
	},
	"ldp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldp",
			Port:  646,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ldp",
			Port:  646,
		},
	},
	"lds-distrib": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lds-distrib",
			Port:  6543,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lds-distrib",
			Port:  6543,
		},
	},
	"lds-dump": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lds-dump",
			Port:  6544,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lds-dump",
			Port:  6544,
		},
	},
	"ldss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldss",
			Port:  6087,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ldss",
			Port:  6087,
		},
	},
	"ldxp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ldxp",
			Port:  4042,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ldxp",
			Port:  4042,
		},
	},
	"lecroy-vicp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lecroy-vicp",
			Port:  1861,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lecroy-vicp",
			Port:  1861,
		},
	},
	"leecoposserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "leecoposserver",
			Port:  2212,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "leecoposserver",
			Port:  2212,
		},
	},
	"legent-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "legent-1",
			Port:  373,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "legent-1",
			Port:  373,
		},
	},
	"legent-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "legent-2",
			Port:  374,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "legent-2",
			Port:  374,
		},
	},
	"leoip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "leoip",
			Port:  1886,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "leoip",
			Port:  1886,
		},
	},
	"lhtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lhtp",
			Port:  1983,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lhtp",
			Port:  1983,
		},
	},
	"liberty-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "liberty-lm",
			Port:  1496,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "liberty-lm",
			Port:  1496,
		},
	},
	"licensedaemon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "licensedaemon",
			Port:  1986,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "licensedaemon",
			Port:  1986,
		},
	},
	"light": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "light",
			Port:  4670,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "light",
			Port:  4670,
		},
	},
	"link": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "link",
			Port:  245,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "link",
			Port:  245,
		},
	},
	"linkname": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "linkname",
			Port:  1903,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "linkname",
			Port:  1903,
		},
	},
	"linktest": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "linktest",
			Port:  3746,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "linktest",
			Port:  3746,
		},
	},
	"linktest-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "linktest-s",
			Port:  3747,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "linktest-s",
			Port:  3747,
		},
	},
	"linogridengine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "linogridengine",
			Port:  12300,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "linogridengine",
			Port:  12300,
		},
	},
	"linx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "linx",
			Port:  1361,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "linx",
			Port:  1361,
		},
	},
	"lionhead": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lionhead",
			Port:  2611,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lionhead",
			Port:  2611,
		},
	},
	"lipsinc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lipsinc",
			Port:  1968,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lipsinc",
			Port:  1968,
		},
	},
	"lipsinc1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lipsinc1",
			Port:  1969,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lipsinc1",
			Port:  1969,
		},
	},
	"lisp-cons": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lisp-cons",
			Port:  4342,
		},
	},
	"lisp-control": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "lisp-control",
			Port:  4342,
		},
	},
	"lisp-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lisp-data",
			Port:  4341,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lisp-data",
			Port:  4341,
		},
	},
	"lispworks-orb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lispworks-orb",
			Port:  3672,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lispworks-orb",
			Port:  3672,
		},
	},
	"listcrt-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "listcrt-port",
			Port:  3913,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "listcrt-port",
			Port:  3913,
		},
	},
	"listcrt-port-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "listcrt-port-2",
			Port:  3914,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "listcrt-port-2",
			Port:  3914,
		},
	},
	"listmgr-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "listmgr-port",
			Port:  3767,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "listmgr-port",
			Port:  3767,
		},
	},
	"livelan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "livelan",
			Port:  1555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "livelan",
			Port:  1555,
		},
	},
	"livestats": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "livestats",
			Port:  2795,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "livestats",
			Port:  2795,
		},
	},
	"ljk-login": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ljk-login",
			Port:  472,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ljk-login",
			Port:  472,
		},
	},
	"lkcmserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lkcmserver",
			Port:  3278,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lkcmserver",
			Port:  3278,
		},
	},
	"llm-csv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "llm-csv",
			Port:  2814,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "llm-csv",
			Port:  2814,
		},
	},
	"llm-pass": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "llm-pass",
			Port:  2813,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "llm-pass",
			Port:  2813,
		},
	},
	"llrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "llrp",
			Port:  5084,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "llrp",
			Port:  5084,
		},
	},
	"llsurfup-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "llsurfup-http",
			Port:  1183,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "llsurfup-http",
			Port:  1183,
		},
	},
	"llsurfup-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "llsurfup-https",
			Port:  1184,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "llsurfup-https",
			Port:  1184,
		},
	},
	"lm-dta": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lm-dta",
			Port:  8206,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lm-dta",
			Port:  8206,
		},
	},
	"lm-instmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lm-instmgr",
			Port:  8205,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lm-instmgr",
			Port:  8205,
		},
	},
	"lm-mon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lm-mon",
			Port:  31620,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lm-mon",
			Port:  31620,
		},
	},
	"lm-perfworks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lm-perfworks",
			Port:  8204,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lm-perfworks",
			Port:  8204,
		},
	},
	"lm-sserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lm-sserver",
			Port:  8207,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lm-sserver",
			Port:  8207,
		},
	},
	"lm-webwatcher": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lm-webwatcher",
			Port:  8208,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lm-webwatcher",
			Port:  8208,
		},
	},
	"lm-x": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lm-x",
			Port:  6200,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lm-x",
			Port:  6200,
		},
	},
	"lmcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lmcs",
			Port:  4877,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lmcs",
			Port:  4877,
		},
	},
	"lmdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lmdp",
			Port:  2623,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lmdp",
			Port:  2623,
		},
	},
	"lmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lmp",
			Port:  701,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lmp",
			Port:  701,
		},
	},
	"lms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lms",
			Port:  4056,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lms",
			Port:  4056,
		},
	},
	"lmsocialserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lmsocialserver",
			Port:  1111,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lmsocialserver",
			Port:  1111,
		},
	},
	"lmtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lmtp",
			Port:  24,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lmtp",
			Port:  24,
		},
	},
	"lnvalarm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lnvalarm",
			Port:  2282,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lnvalarm",
			Port:  2282,
		},
	},
	"lnvconsole": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lnvconsole",
			Port:  2281,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lnvconsole",
			Port:  2281,
		},
	},
	"lnvmailmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lnvmailmon",
			Port:  2285,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lnvmailmon",
			Port:  2285,
		},
	},
	"lnvmaps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lnvmaps",
			Port:  2284,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lnvmaps",
			Port:  2284,
		},
	},
	"lnvpoller": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lnvpoller",
			Port:  2280,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lnvpoller",
			Port:  2280,
		},
	},
	"lnvstatus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lnvstatus",
			Port:  2283,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lnvstatus",
			Port:  2283,
		},
	},
	"loaprobe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "loaprobe",
			Port:  1634,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "loaprobe",
			Port:  1634,
		},
	},
	"localinfosrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "localinfosrvr",
			Port:  1487,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "localinfosrvr",
			Port:  1487,
		},
	},
	"lockstep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lockstep",
			Port:  2125,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lockstep",
			Port:  2125,
		},
	},
	"locus-con": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "locus-con",
			Port:  127,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "locus-con",
			Port:  127,
		},
	},
	"locus-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "locus-disc",
			Port:  5058,
		},
	},
	"locus-map": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "locus-map",
			Port:  125,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "locus-map",
			Port:  125,
		},
	},
	"lofr-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lofr-lm",
			Port:  1752,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lofr-lm",
			Port:  1752,
		},
	},
	"login": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "login",
			Port:  513,
		},
	},
	"lonewolf-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lonewolf-lm",
			Port:  6146,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lonewolf-lm",
			Port:  6146,
		},
	},
	"lontalk-norm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lontalk-norm",
			Port:  1628,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lontalk-norm",
			Port:  1628,
		},
	},
	"lontalk-urgnt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lontalk-urgnt",
			Port:  1629,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lontalk-urgnt",
			Port:  1629,
		},
	},
	"lonworks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lonworks",
			Port:  2540,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lonworks",
			Port:  2540,
		},
	},
	"lonworks2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lonworks2",
			Port:  2541,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lonworks2",
			Port:  2541,
		},
	},
	"lorica-in": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lorica-in",
			Port:  4080,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lorica-in",
			Port:  4080,
		},
	},
	"lorica-in-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lorica-in-sec",
			Port:  4081,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lorica-in-sec",
			Port:  4081,
		},
	},
	"lorica-out": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lorica-out",
			Port:  4082,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lorica-out",
			Port:  4082,
		},
	},
	"lorica-out-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lorica-out-sec",
			Port:  4083,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lorica-out-sec",
			Port:  4083,
		},
	},
	"lot105-ds-upd": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "lot105-ds-upd",
			Port:  2053,
		},
	},
	"lotusmtap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lotusmtap",
			Port:  3007,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lotusmtap",
			Port:  3007,
		},
	},
	"lotusnote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lotusnote",
			Port:  1352,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lotusnote",
			Port:  1352,
		},
	},
	"lpcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lpcp",
			Port:  1298,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lpcp",
			Port:  1298,
		},
	},
	"lpdg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lpdg",
			Port:  10805,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lpdg",
			Port:  10805,
		},
	},
	"lpsrecommender": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lpsrecommender",
			Port:  2620,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lpsrecommender",
			Port:  2620,
		},
	},
	"lrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lrp",
			Port:  2090,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lrp",
			Port:  2090,
		},
	},
	"lrs-paging": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lrs-paging",
			Port:  3700,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lrs-paging",
			Port:  3700,
		},
	},
	"ls3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ls3",
			Port:  3069,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ls3",
			Port:  3069,
		},
	},
	"ls3bcast": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ls3bcast",
			Port:  3068,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ls3bcast",
			Port:  3068,
		},
	},
	"lsi-bobcat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lsi-bobcat",
			Port:  5574,
		},
	},
	"lsi-raid-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lsi-raid-mgmt",
			Port:  2463,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lsi-raid-mgmt",
			Port:  2463,
		},
	},
	"lsp-ping": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lsp-ping",
			Port:  3503,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lsp-ping",
			Port:  3503,
		},
	},
	"lstp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lstp",
			Port:  2559,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lstp",
			Port:  2559,
		},
	},
	"ltctcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ltctcp",
			Port:  3487,
		},
	},
	"ltcudp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ltcudp",
			Port:  3487,
		},
	},
	"ltp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ltp",
			Port:  4044,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ltp",
			Port:  4044,
		},
	},
	"ltp-deepspace": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ltp-deepspace",
			Port:  1113,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ltp-deepspace",
			Port:  1113,
		},
	},
	"lumimgrd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lumimgrd",
			Port:  4741,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lumimgrd",
			Port:  4741,
		},
	},
	"lupa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lupa",
			Port:  1212,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lupa",
			Port:  1212,
		},
	},
	"lutap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lutap",
			Port:  4912,
		},
	},
	"lutcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lutcp",
			Port:  4913,
		},
	},
	"lv-auth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lv-auth",
			Port:  2147,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lv-auth",
			Port:  2147,
		},
	},
	"lv-ffx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lv-ffx",
			Port:  2144,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lv-ffx",
			Port:  2144,
		},
	},
	"lv-frontpanel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lv-frontpanel",
			Port:  3079,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lv-frontpanel",
			Port:  3079,
		},
	},
	"lv-jc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lv-jc",
			Port:  2143,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lv-jc",
			Port:  2143,
		},
	},
	"lv-not": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lv-not",
			Port:  2146,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lv-not",
			Port:  2146,
		},
	},
	"lv-pici": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lv-pici",
			Port:  2145,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lv-pici",
			Port:  2145,
		},
	},
	"lvision-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lvision-lm",
			Port:  6471,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lvision-lm",
			Port:  6471,
		},
	},
	"lxi-evntsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lxi-evntsvc",
			Port:  5044,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lxi-evntsvc",
			Port:  5044,
		},
	},
	"lyskom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "lyskom",
			Port:  4894,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "lyskom",
			Port:  4894,
		},
	},
	"m-wnn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "m-wnn",
			Port:  3732,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "m-wnn",
			Port:  3732,
		},
	},
	"m2ap": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "m2ap",
			Port:  36443,
		},
	},
	"m2mservices": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "m2mservices",
			Port:  8383,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "m2mservices",
			Port:  8383,
		},
	},
	"m2pa": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "m2pa",
			Port:  3565,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "m2pa",
			Port:  3565,
		},
	},
	"m2ua": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "m2ua",
			Port:  2904,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "m2ua",
			Port:  2904,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "m2ua",
			Port:  2904,
		},
	},
	"m3ap": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "m3ap",
			Port:  36444,
		},
	},
	"m3da": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "m3da",
			Port:  44900,
		},
	},
	"m3da-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "m3da-disc",
			Port:  44900,
		},
	},
	"m3ua": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "m3ua",
			Port:  2905,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "m3ua",
			Port:  2905,
		},
	},
	"m4-network-as": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "m4-network-as",
			Port:  4345,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "m4-network-as",
			Port:  4345,
		},
	},
	"mac-srvr-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mac-srvr-admin",
			Port:  660,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mac-srvr-admin",
			Port:  660,
		},
	},
	"macbak": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "macbak",
			Port:  4181,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "macbak",
			Port:  4181,
		},
	},
	"macon-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "macon-tcp",
			Port:  456,
		},
	},
	"macon-udp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "macon-udp",
			Port:  456,
		},
	},
	"macromedia-fcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "macromedia-fcs",
			Port:  1935,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "macromedia-fcs",
			Port:  1935,
		},
	},
	"madcap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "madcap",
			Port:  2535,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "madcap",
			Port:  2535,
		},
	},
	"madge-ltd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "madge-ltd",
			Port:  2453,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "madge-ltd",
			Port:  2453,
		},
	},
	"magaya-network": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "magaya-network",
			Port:  3691,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "magaya-network",
			Port:  3691,
		},
	},
	"magbind": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "magbind",
			Port:  3194,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "magbind",
			Port:  3194,
		},
	},
	"magenta-logic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "magenta-logic",
			Port:  313,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "magenta-logic",
			Port:  313,
		},
	},
	"magiccontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "magiccontrol",
			Port:  4902,
		},
	},
	"magicnotes": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "magicnotes",
			Port:  3023,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "magicnotes",
			Port:  3023,
		},
	},
	"magicom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "magicom",
			Port:  2243,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "magicom",
			Port:  2243,
		},
	},
	"magpie": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "magpie",
			Port:  5092,
		},
	},
	"mailbox": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mailbox",
			Port:  2004,
		},
	},
	"mailbox-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mailbox-lm",
			Port:  505,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mailbox-lm",
			Port:  505,
		},
	},
	"mailprox": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mailprox",
			Port:  3936,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mailprox",
			Port:  3936,
		},
	},
	"mailq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mailq",
			Port:  174,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mailq",
			Port:  174,
		},
	},
	"maincontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "maincontrol",
			Port:  2516,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "maincontrol",
			Port:  2516,
		},
	},
	"mainsoft-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mainsoft-lm",
			Port:  1593,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mainsoft-lm",
			Port:  1593,
		},
	},
	"maitrd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "maitrd",
			Port:  997,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "maitrd",
			Port:  997,
		},
	},
	"manage-exec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "manage-exec",
			Port:  2342,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "manage-exec",
			Port:  2342,
		},
	},
	"mandelspawn": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mandelspawn",
			Port:  9359,
		},
	},
	"manet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "manet",
			Port:  269,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "manet",
			Port:  269,
		},
	},
	"manyone-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "manyone-http",
			Port:  8910,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "manyone-http",
			Port:  8910,
		},
	},
	"manyone-xml": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "manyone-xml",
			Port:  8911,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "manyone-xml",
			Port:  8911,
		},
	},
	"mao": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mao",
			Port:  2908,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mao",
			Port:  2908,
		},
	},
	"mapper-mapethd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mapper-mapethd",
			Port:  3985,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mapper-mapethd",
			Port:  3985,
		},
	},
	"mapper-nodemgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mapper-nodemgr",
			Port:  3984,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mapper-nodemgr",
			Port:  3984,
		},
	},
	"mapper-ws_ethd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mapper-ws_ethd",
			Port:  3986,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mapper-ws_ethd",
			Port:  3986,
		},
	},
	"marcam-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "marcam-lm",
			Port:  1444,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "marcam-lm",
			Port:  1444,
		},
	},
	"markem-dcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "markem-dcp",
			Port:  3836,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "markem-dcp",
			Port:  3836,
		},
	},
	"masc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "masc",
			Port:  2587,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "masc",
			Port:  2587,
		},
	},
	"masqdialer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "masqdialer",
			Port:  224,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "masqdialer",
			Port:  224,
		},
	},
	"matahari": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "matahari",
			Port:  49000,
		},
	},
	"matip-type-a": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "matip-type-a",
			Port:  350,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "matip-type-a",
			Port:  350,
		},
	},
	"matip-type-b": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "matip-type-b",
			Port:  351,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "matip-type-b",
			Port:  351,
		},
	},
	"matrix_vnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "matrix_vnet",
			Port:  4360,
		},
	},
	"max": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "max",
			Port:  6074,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "max",
			Port:  6074,
		},
	},
	"maxim-asics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "maxim-asics",
			Port:  3276,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "maxim-asics",
			Port:  3276,
		},
	},
	"maytagshuffle": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "maytagshuffle",
			Port:  2591,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "maytagshuffle",
			Port:  2591,
		},
	},
	"mbg-ctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mbg-ctrl",
			Port:  3569,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mbg-ctrl",
			Port:  3569,
		},
	},
	"mbl-battd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mbl-battd",
			Port:  4153,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mbl-battd",
			Port:  4153,
		},
	},
	"mbus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mbus",
			Port:  47000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mbus",
			Port:  47000,
		},
	},
	"mc-appserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mc-appserver",
			Port:  8763,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mc-appserver",
			Port:  8763,
		},
	},
	"mc-brk-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mc-brk-srv",
			Port:  3180,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mc-brk-srv",
			Port:  3180,
		},
	},
	"mc-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mc-client",
			Port:  1180,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mc-client",
			Port:  1180,
		},
	},
	"mc-comm": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mc-comm",
			Port:  9632,
		},
	},
	"mc-gt-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mc-gt-srv",
			Port:  2180,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mc-gt-srv",
			Port:  2180,
		},
	},
	"mc2studios": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mc2studios",
			Port:  1899,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mc2studios",
			Port:  1899,
		},
	},
	"mc3ss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mc3ss",
			Port:  3521,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mc3ss",
			Port:  3521,
		},
	},
	"mcagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcagent",
			Port:  1820,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcagent",
			Port:  1820,
		},
	},
	"mccwebsvr-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mccwebsvr-port",
			Port:  3570,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mccwebsvr-port",
			Port:  3570,
		},
	},
	"mcer-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcer-port",
			Port:  6510,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcer-port",
			Port:  6510,
		},
	},
	"mcftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcftp",
			Port:  6622,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcftp",
			Port:  6622,
		},
	},
	"mcidas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcidas",
			Port:  112,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcidas",
			Port:  112,
		},
	},
	"mck-ivpip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mck-ivpip",
			Port:  2698,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mck-ivpip",
			Port:  2698,
		},
	},
	"mcns-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcns-sec",
			Port:  638,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcns-sec",
			Port:  638,
		},
	},
	"mcns-tel-ret": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcns-tel-ret",
			Port:  3311,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcns-tel-ret",
			Port:  3311,
		},
	},
	"mcntp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcntp",
			Port:  5418,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcntp",
			Port:  5418,
		},
	},
	"mcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcp",
			Port:  4458,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcp",
			Port:  4458,
		},
	},
	"mcp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcp-port",
			Port:  3558,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcp-port",
			Port:  3558,
		},
	},
	"mcreport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcreport",
			Port:  8003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcreport",
			Port:  8003,
		},
	},
	"mcs-calypsoicf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcs-calypsoicf",
			Port:  3330,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcs-calypsoicf",
			Port:  3330,
		},
	},
	"mcs-fastmail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcs-fastmail",
			Port:  3302,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcs-fastmail",
			Port:  3302,
		},
	},
	"mcs-mailsvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcs-mailsvr",
			Port:  3332,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcs-mailsvr",
			Port:  3332,
		},
	},
	"mcs-messaging": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mcs-messaging",
			Port:  3331,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mcs-messaging",
			Port:  3331,
		},
	},
	"mctet-gateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mctet-gateway",
			Port:  3116,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mctet-gateway",
			Port:  3116,
		},
	},
	"mctet-jserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mctet-jserv",
			Port:  3117,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mctet-jserv",
			Port:  3117,
		},
	},
	"mctet-master": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mctet-master",
			Port:  3115,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mctet-master",
			Port:  3115,
		},
	},
	"mctfeed": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mctfeed",
			Port:  5598,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mctfeed",
			Port:  5598,
		},
	},
	"mctp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mctp",
			Port:  1100,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mctp",
			Port:  1100,
		},
	},
	"md-cg-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "md-cg-http",
			Port:  2688,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "md-cg-http",
			Port:  2688,
		},
	},
	"mdap-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mdap-port",
			Port:  3235,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mdap-port",
			Port:  3235,
		},
	},
	"mdbs_daemon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mdbs_daemon",
			Port:  800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mdbs_daemon",
			Port:  800,
		},
	},
	"mdc-portmapper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mdc-portmapper",
			Port:  685,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mdc-portmapper",
			Port:  685,
		},
	},
	"mdm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mdm",
			Port:  7871,
		},
	},
	"mdns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mdns",
			Port:  5353,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mdns",
			Port:  5353,
		},
	},
	"mdnsresponder": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mdnsresponder",
			Port:  5354,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mdnsresponder",
			Port:  5354,
		},
	},
	"mdqs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mdqs",
			Port:  666,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mdqs",
			Port:  666,
		},
	},
	"mdtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mdtp",
			Port:  3232,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mdtp",
			Port:  3232,
		},
	},
	"mecomm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mecomm",
			Port:  668,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mecomm",
			Port:  668,
		},
	},
	"med-ci": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "med-ci",
			Port:  24005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "med-ci",
			Port:  24005,
		},
	},
	"med-fsp-rx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "med-fsp-rx",
			Port:  24001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "med-fsp-rx",
			Port:  24001,
		},
	},
	"med-fsp-tx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "med-fsp-tx",
			Port:  24002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "med-fsp-tx",
			Port:  24002,
		},
	},
	"med-ltp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "med-ltp",
			Port:  24000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "med-ltp",
			Port:  24000,
		},
	},
	"med-net-svc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "med-net-svc",
			Port:  24006,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "med-net-svc",
			Port:  24006,
		},
	},
	"med-ovw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "med-ovw",
			Port:  24004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "med-ovw",
			Port:  24004,
		},
	},
	"med-supp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "med-supp",
			Port:  24003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "med-supp",
			Port:  24003,
		},
	},
	"medevolve": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "medevolve",
			Port:  13930,
		},
	},
	"media-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "media-agent",
			Port:  2789,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "media-agent",
			Port:  2789,
		},
	},
	"mediabox": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mediabox",
			Port:  46999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mediabox",
			Port:  46999,
		},
	},
	"mediacntrlnfsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mediacntrlnfsd",
			Port:  2363,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mediacntrlnfsd",
			Port:  2363,
		},
	},
	"mediaspace": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mediaspace",
			Port:  3594,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mediaspace",
			Port:  3594,
		},
	},
	"mediat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mediat",
			Port:  5157,
		},
	},
	"mediavault-gui": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mediavault-gui",
			Port:  3673,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mediavault-gui",
			Port:  3673,
		},
	},
	"medimageportal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "medimageportal",
			Port:  7720,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "medimageportal",
			Port:  7720,
		},
	},
	"megaco-h248": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "megaco-h248",
			Port:  2944,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "megaco-h248",
			Port:  2944,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "megaco-h248",
			Port:  2944,
		},
	},
	"megardsvr-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "megardsvr-port",
			Port:  3571,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "megardsvr-port",
			Port:  3571,
		},
	},
	"megaregsvrport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "megaregsvrport",
			Port:  3572,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "megaregsvrport",
			Port:  3572,
		},
	},
	"memcache": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "memcache",
			Port:  11211,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "memcache",
			Port:  11211,
		},
	},
	"menandmice-dns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "menandmice-dns",
			Port:  1337,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "menandmice-dns",
			Port:  1337,
		},
	},
	"menandmice-lpm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "menandmice-lpm",
			Port:  1231,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "menandmice-lpm",
			Port:  1231,
		},
	},
	"menandmice-mon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "menandmice-mon",
			Port:  4552,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "menandmice-mon",
			Port:  4552,
		},
	},
	"menandmice-upg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "menandmice-upg",
			Port:  4603,
		},
	},
	"menandmice_noh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "menandmice_noh",
			Port:  4151,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "menandmice_noh",
			Port:  4151,
		},
	},
	"mentaclient": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mentaclient",
			Port:  2117,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mentaclient",
			Port:  2117,
		},
	},
	"mentaserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mentaserver",
			Port:  2118,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mentaserver",
			Port:  2118,
		},
	},
	"mercantile": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mercantile",
			Port:  3398,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mercantile",
			Port:  3398,
		},
	},
	"mercury-disc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mercury-disc",
			Port:  9596,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mercury-disc",
			Port:  9596,
		},
	},
	"meregister": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "meregister",
			Port:  669,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "meregister",
			Port:  669,
		},
	},
	"mesavistaco": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mesavistaco",
			Port:  1249,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mesavistaco",
			Port:  1249,
		},
	},
	"messageasap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "messageasap",
			Port:  6070,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "messageasap",
			Port:  6070,
		},
	},
	"messageservice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "messageservice",
			Port:  2311,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "messageservice",
			Port:  2311,
		},
	},
	"meta-corp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "meta-corp",
			Port:  6141,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "meta-corp",
			Port:  6141,
		},
	},
	"meta5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "meta5",
			Port:  393,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "meta5",
			Port:  393,
		},
	},
	"metaagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metaagent",
			Port:  1897,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metaagent",
			Port:  1897,
		},
	},
	"metaconsole": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metaconsole",
			Port:  2850,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metaconsole",
			Port:  2850,
		},
	},
	"metaedit-mu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metaedit-mu",
			Port:  6360,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metaedit-mu",
			Port:  6360,
		},
	},
	"metaedit-se": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metaedit-se",
			Port:  6370,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metaedit-se",
			Port:  6370,
		},
	},
	"metaedit-ws": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metaedit-ws",
			Port:  6390,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metaedit-ws",
			Port:  6390,
		},
	},
	"metagram": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metagram",
			Port:  99,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metagram",
			Port:  99,
		},
	},
	"metalbend": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metalbend",
			Port:  7172,
		},
	},
	"metasage": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metasage",
			Port:  1207,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metasage",
			Port:  1207,
		},
	},
	"metastorm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metastorm",
			Port:  2511,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metastorm",
			Port:  2511,
		},
	},
	"metasys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metasys",
			Port:  11001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metasys",
			Port:  11001,
		},
	},
	"metatude-mds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metatude-mds",
			Port:  6382,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metatude-mds",
			Port:  6382,
		},
	},
	"meter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "meter",
			Port:  570,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "meter",
			Port:  570,
		},
	},
	"metricadbc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metricadbc",
			Port:  2622,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metricadbc",
			Port:  2622,
		},
	},
	"metrics-pas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "metrics-pas",
			Port:  1824,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "metrics-pas",
			Port:  1824,
		},
	},
	"mevent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mevent",
			Port:  7900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mevent",
			Port:  7900,
		},
	},
	"mfcobol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mfcobol",
			Port:  86,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mfcobol",
			Port:  86,
		},
	},
	"mfserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mfserver",
			Port:  2266,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mfserver",
			Port:  2266,
		},
	},
	"mftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mftp",
			Port:  349,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mftp",
			Port:  349,
		},
	},
	"mgcp-callagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mgcp-callagent",
			Port:  2727,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mgcp-callagent",
			Port:  2727,
		},
	},
	"mgcp-gateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mgcp-gateway",
			Port:  2427,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mgcp-gateway",
			Port:  2427,
		},
	},
	"mgcs-mfp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mgcs-mfp-port",
			Port:  6509,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mgcs-mfp-port",
			Port:  6509,
		},
	},
	"mgemanagement": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mgemanagement",
			Port:  4680,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mgemanagement",
			Port:  4680,
		},
	},
	"mgesupervision": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mgesupervision",
			Port:  4679,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mgesupervision",
			Port:  4679,
		},
	},
	"mgxswitch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mgxswitch",
			Port:  3070,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mgxswitch",
			Port:  3070,
		},
	},
	"miami-bcast": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "miami-bcast",
			Port:  6083,
		},
	},
	"mib-streaming": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mib-streaming",
			Port:  2292,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mib-streaming",
			Port:  2292,
		},
	},
	"mice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mice",
			Port:  5022,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mice",
			Port:  5022,
		},
	},
	"micom-pfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "micom-pfs",
			Port:  490,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "micom-pfs",
			Port:  490,
		},
	},
	"microcom-sbp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "microcom-sbp",
			Port:  1680,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "microcom-sbp",
			Port:  1680,
		},
	},
	"micromuse-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "micromuse-lm",
			Port:  1534,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "micromuse-lm",
			Port:  1534,
		},
	},
	"micromuse-ncps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "micromuse-ncps",
			Port:  7979,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "micromuse-ncps",
			Port:  7979,
		},
	},
	"micromuse-ncpw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "micromuse-ncpw",
			Port:  9600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "micromuse-ncpw",
			Port:  9600,
		},
	},
	"microsan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "microsan",
			Port:  20001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "microsan",
			Port:  20001,
		},
	},
	"microsoft-ds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "microsoft-ds",
			Port:  445,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "microsoft-ds",
			Port:  445,
		},
	},
	"microtalon-com": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "microtalon-com",
			Port:  7014,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "microtalon-com",
			Port:  7014,
		},
	},
	"microtalon-dis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "microtalon-dis",
			Port:  7013,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "microtalon-dis",
			Port:  7013,
		},
	},
	"midnight-tech": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "midnight-tech",
			Port:  3008,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "midnight-tech",
			Port:  3008,
		},
	},
	"mikey": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mikey",
			Port:  2269,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mikey",
			Port:  2269,
		},
	},
	"mil-2045-47001": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mil-2045-47001",
			Port:  1581,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mil-2045-47001",
			Port:  1581,
		},
	},
	"miles-apart": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "miles-apart",
			Port:  2621,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "miles-apart",
			Port:  2621,
		},
	},
	"mimer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mimer",
			Port:  1360,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mimer",
			Port:  1360,
		},
	},
	"mindarray-ca": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mindarray-ca",
			Port:  9445,
		},
	},
	"mindfilesys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mindfilesys",
			Port:  7391,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mindfilesys",
			Port:  7391,
		},
	},
	"mindprint": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mindprint",
			Port:  8033,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mindprint",
			Port:  8033,
		},
	},
	"minger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "minger",
			Port:  4069,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "minger",
			Port:  4069,
		},
	},
	"mini-sql": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mini-sql",
			Port:  1114,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mini-sql",
			Port:  1114,
		},
	},
	"minilock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "minilock",
			Port:  3798,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "minilock",
			Port:  3798,
		},
	},
	"minipay": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "minipay",
			Port:  2105,
		},
	},
	"minivend": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "minivend",
			Port:  7786,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "minivend",
			Port:  7786,
		},
	},
	"minotaur-sa": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "minotaur-sa",
			Port:  5136,
		},
	},
	"mipv6tls": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mipv6tls",
			Port:  7872,
		},
	},
	"mira": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mira",
			Port:  3454,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mira",
			Port:  3454,
		},
	},
	"miroconnect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "miroconnect",
			Port:  1532,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "miroconnect",
			Port:  1532,
		},
	},
	"mirrtex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mirrtex",
			Port:  4310,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mirrtex",
			Port:  4310,
		},
	},
	"mit-dov": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mit-dov",
			Port:  91,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mit-dov",
			Port:  91,
		},
	},
	"mit-ml-dev": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mit-ml-dev",
			Port:  83,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mit-ml-dev",
			Port:  83,
		},
	},
	"miteksys-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "miteksys-lm",
			Port:  1482,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "miteksys-lm",
			Port:  1482,
		},
	},
	"miva-mqs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "miva-mqs",
			Port:  1277,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "miva-mqs",
			Port:  1277,
		},
	},
	"mkm-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mkm-discovery",
			Port:  3837,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mkm-discovery",
			Port:  3837,
		},
	},
	"ml-svnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ml-svnet",
			Port:  4171,
		},
	},
	"mle": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mle",
			Port:  19788,
		},
	},
	"mloadd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mloadd",
			Port:  1427,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mloadd",
			Port:  1427,
		},
	},
	"mlsn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mlsn",
			Port:  32801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mlsn",
			Port:  32801,
		},
	},
	"mma-discovery": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mma-discovery",
			Port:  4173,
		},
	},
	"mmacomm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mmacomm",
			Port:  4667,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mmacomm",
			Port:  4667,
		},
	},
	"mmaeds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mmaeds",
			Port:  4668,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mmaeds",
			Port:  4668,
		},
	},
	"mmcal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mmcal",
			Port:  2272,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mmcal",
			Port:  2272,
		},
	},
	"mmcals": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mmcals",
			Port:  2271,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mmcals",
			Port:  2271,
		},
	},
	"mmcc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mmcc",
			Port:  5050,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mmcc",
			Port:  5050,
		},
	},
	"mmpft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mmpft",
			Port:  1815,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mmpft",
			Port:  1815,
		},
	},
	"mnet-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mnet-discovery",
			Port:  5237,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mnet-discovery",
			Port:  5237,
		},
	},
	"mngsuite": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mngsuite",
			Port:  9535,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mngsuite",
			Port:  9535,
		},
	},
	"mni-prot-rout": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mni-prot-rout",
			Port:  3764,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mni-prot-rout",
			Port:  3764,
		},
	},
	"mnp-exchange": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mnp-exchange",
			Port:  2197,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mnp-exchange",
			Port:  2197,
		},
	},
	"mns-mail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mns-mail",
			Port:  2593,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mns-mail",
			Port:  2593,
		},
	},
	"mobile-file-dl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mobile-file-dl",
			Port:  2926,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mobile-file-dl",
			Port:  2926,
		},
	},
	"mobile-p2p": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mobile-p2p",
			Port:  4688,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mobile-p2p",
			Port:  4688,
		},
	},
	"mobileanalyzer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mobileanalyzer",
			Port:  7869,
		},
	},
	"mobileip-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mobileip-agent",
			Port:  434,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mobileip-agent",
			Port:  434,
		},
	},
	"mobilip-mn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mobilip-mn",
			Port:  435,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mobilip-mn",
			Port:  435,
		},
	},
	"mobrien-chat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mobrien-chat",
			Port:  2031,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mobrien-chat",
			Port:  2031,
		},
	},
	"moldflow-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "moldflow-lm",
			Port:  1576,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "moldflow-lm",
			Port:  1576,
		},
	},
	"molly": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "molly",
			Port:  1374,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "molly",
			Port:  1374,
		},
	},
	"mon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mon",
			Port:  2583,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mon",
			Port:  2583,
		},
	},
	"mondex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mondex",
			Port:  471,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mondex",
			Port:  471,
		},
	},
	"monetra": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "monetra",
			Port:  8665,
		},
	},
	"monetra-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "monetra-admin",
			Port:  8666,
		},
	},
	"monitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "monitor",
			Port:  561,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "monitor",
			Port:  561,
		},
	},
	"monkeycom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "monkeycom",
			Port:  9898,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "monkeycom",
			Port:  9898,
		},
	},
	"monp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "monp",
			Port:  3445,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "monp",
			Port:  3445,
		},
	},
	"montage-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "montage-lm",
			Port:  6147,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "montage-lm",
			Port:  6147,
		},
	},
	"mortgageware": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mortgageware",
			Port:  367,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mortgageware",
			Port:  367,
		},
	},
	"mosaicsyssvc1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mosaicsyssvc1",
			Port:  1235,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mosaicsyssvc1",
			Port:  1235,
		},
	},
	"mosaixcc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mosaixcc",
			Port:  2561,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mosaixcc",
			Port:  2561,
		},
	},
	"moshebeeri": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "moshebeeri",
			Port:  2627,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "moshebeeri",
			Port:  2627,
		},
	},
	"mountd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mountd",
			Port:  20048,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mountd",
			Port:  20048,
		},
	},
	"movaz-ssc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "movaz-ssc",
			Port:  5252,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "movaz-ssc",
			Port:  5252,
		},
	},
	"moy-corp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "moy-corp",
			Port:  2488,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "moy-corp",
			Port:  2488,
		},
	},
	"mpc-lifenet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpc-lifenet",
			Port:  1213,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpc-lifenet",
			Port:  1213,
		},
	},
	"mpfoncl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpfoncl",
			Port:  2579,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpfoncl",
			Port:  2579,
		},
	},
	"mpfwsas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpfwsas",
			Port:  2952,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpfwsas",
			Port:  2952,
		},
	},
	"mphlpdmc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mphlpdmc",
			Port:  9344,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mphlpdmc",
			Port:  9344,
		},
	},
	"mpidcagt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpidcagt",
			Port:  9397,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpidcagt",
			Port:  9397,
		},
	},
	"mpidcmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpidcmgr",
			Port:  9343,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpidcmgr",
			Port:  9343,
		},
	},
	"mpl-gprs-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpl-gprs-port",
			Port:  3924,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpl-gprs-port",
			Port:  3924,
		},
	},
	"mpm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpm",
			Port:  45,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpm",
			Port:  45,
		},
	},
	"mpm-flags": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpm-flags",
			Port:  44,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpm-flags",
			Port:  44,
		},
	},
	"mpm-snd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpm-snd",
			Port:  46,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpm-snd",
			Port:  46,
		},
	},
	"mpnjsc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpnjsc",
			Port:  1952,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpnjsc",
			Port:  1952,
		},
	},
	"mpnjsocl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpnjsocl",
			Port:  2685,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpnjsocl",
			Port:  2685,
		},
	},
	"mpnjsomb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpnjsomb",
			Port:  2681,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpnjsomb",
			Port:  2681,
		},
	},
	"mpnjsomg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpnjsomg",
			Port:  2686,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpnjsomg",
			Port:  2686,
		},
	},
	"mpnjsosv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpnjsosv",
			Port:  2684,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpnjsosv",
			Port:  2684,
		},
	},
	"mpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpp",
			Port:  218,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpp",
			Port:  218,
		},
	},
	"mppolicy-mgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mppolicy-mgr",
			Port:  5969,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mppolicy-mgr",
			Port:  5969,
		},
	},
	"mppolicy-v5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mppolicy-v5",
			Port:  5968,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mppolicy-v5",
			Port:  5968,
		},
	},
	"mps-raft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mps-raft",
			Port:  1700,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mps-raft",
			Port:  1700,
		},
	},
	"mpshrsv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpshrsv",
			Port:  1261,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpshrsv",
			Port:  1261,
		},
	},
	"mpsserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpsserver",
			Port:  6106,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpsserver",
			Port:  6106,
		},
	},
	"mpsysrmsvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mpsysrmsvr",
			Port:  3358,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mpsysrmsvr",
			Port:  3358,
		},
	},
	"mptn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mptn",
			Port:  397,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mptn",
			Port:  397,
		},
	},
	"mqe-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mqe-agent",
			Port:  3958,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mqe-agent",
			Port:  3958,
		},
	},
	"mqe-broker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mqe-broker",
			Port:  3957,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mqe-broker",
			Port:  3957,
		},
	},
	"mrip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mrip",
			Port:  4986,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mrip",
			Port:  4986,
		},
	},
	"mrm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mrm",
			Port:  679,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mrm",
			Port:  679,
		},
	},
	"mrssrendezvous": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mrssrendezvous",
			Port:  7392,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mrssrendezvous",
			Port:  7392,
		},
	},
	"ms-alerter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-alerter",
			Port:  5359,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-alerter",
			Port:  5359,
		},
	},
	"ms-cluster-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-cluster-net",
			Port:  3343,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-cluster-net",
			Port:  3343,
		},
	},
	"ms-ilm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-ilm",
			Port:  5725,
		},
	},
	"ms-ilm-sts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-ilm-sts",
			Port:  5726,
		},
	},
	"ms-la": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-la",
			Port:  3535,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-la",
			Port:  3535,
		},
	},
	"ms-licensing": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-licensing",
			Port:  5720,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-licensing",
			Port:  5720,
		},
	},
	"ms-olap1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-olap1",
			Port:  2393,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-olap1",
			Port:  2393,
		},
	},
	"ms-olap2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-olap2",
			Port:  2394,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-olap2",
			Port:  2394,
		},
	},
	"ms-olap3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-olap3",
			Port:  2382,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-olap3",
			Port:  2382,
		},
	},
	"ms-olap4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-olap4",
			Port:  2383,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-olap4",
			Port:  2383,
		},
	},
	"ms-rome": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-rome",
			Port:  569,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-rome",
			Port:  569,
		},
	},
	"ms-rule-engine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-rule-engine",
			Port:  3132,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-rule-engine",
			Port:  3132,
		},
	},
	"ms-s-sideshow": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-s-sideshow",
			Port:  5361,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-s-sideshow",
			Port:  5361,
		},
	},
	"ms-shuttle": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-shuttle",
			Port:  568,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-shuttle",
			Port:  568,
		},
	},
	"ms-sideshow": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-sideshow",
			Port:  5360,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-sideshow",
			Port:  5360,
		},
	},
	"ms-smlbiz": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-smlbiz",
			Port:  5356,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-smlbiz",
			Port:  5356,
		},
	},
	"ms-sna-base": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-sna-base",
			Port:  1478,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-sna-base",
			Port:  1478,
		},
	},
	"ms-sna-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-sna-server",
			Port:  1477,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-sna-server",
			Port:  1477,
		},
	},
	"ms-sql-m": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-sql-m",
			Port:  1434,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-sql-m",
			Port:  1434,
		},
	},
	"ms-sql-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-sql-s",
			Port:  1433,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-sql-s",
			Port:  1433,
		},
	},
	"ms-streaming": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-streaming",
			Port:  1755,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-streaming",
			Port:  1755,
		},
	},
	"ms-theater": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-theater",
			Port:  2460,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-theater",
			Port:  2460,
		},
	},
	"ms-v-worlds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-v-worlds",
			Port:  2525,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-v-worlds",
			Port:  2525,
		},
	},
	"ms-wbt-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ms-wbt-server",
			Port:  3389,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ms-wbt-server",
			Port:  3389,
		},
	},
	"msdfsr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msdfsr",
			Port:  5722,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msdfsr",
			Port:  5722,
		},
	},
	"msdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msdp",
			Port:  639,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msdp",
			Port:  639,
		},
	},
	"msdts1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msdts1",
			Port:  3882,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msdts1",
			Port:  3882,
		},
	},
	"msexch-routing": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msexch-routing",
			Port:  691,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msexch-routing",
			Port:  691,
		},
	},
	"msfrs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msfrs",
			Port:  4554,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msfrs",
			Port:  4554,
		},
	},
	"msft-dpm-cert": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msft-dpm-cert",
			Port:  6076,
		},
	},
	"msft-gc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msft-gc",
			Port:  3268,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msft-gc",
			Port:  3268,
		},
	},
	"msft-gc-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msft-gc-ssl",
			Port:  3269,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msft-gc-ssl",
			Port:  3269,
		},
	},
	"msfw-array": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msfw-array",
			Port:  2174,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msfw-array",
			Port:  2174,
		},
	},
	"msfw-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msfw-control",
			Port:  3847,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msfw-control",
			Port:  3847,
		},
	},
	"msfw-replica": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msfw-replica",
			Port:  2173,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msfw-replica",
			Port:  2173,
		},
	},
	"msfw-s-storage": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msfw-s-storage",
			Port:  2172,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msfw-s-storage",
			Port:  2172,
		},
	},
	"msfw-storage": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msfw-storage",
			Port:  2171,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msfw-storage",
			Port:  2171,
		},
	},
	"msg-auth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msg-auth",
			Port:  31,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msg-auth",
			Port:  31,
		},
	},
	"msg-icp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msg-icp",
			Port:  29,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msg-icp",
			Port:  29,
		},
	},
	"msgclnt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msgclnt",
			Port:  8786,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msgclnt",
			Port:  8786,
		},
	},
	"msgsrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msgsrvr",
			Port:  8787,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msgsrvr",
			Port:  8787,
		},
	},
	"msgsys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msgsys",
			Port:  9594,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msgsys",
			Port:  9594,
		},
	},
	"mshvlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mshvlm",
			Port:  6600,
		},
	},
	"msi-cps-rm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msi-cps-rm",
			Port:  8675,
		},
	},
	"msi-cps-rm-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "msi-cps-rm-disc",
			Port:  8675,
		},
	},
	"msi-selectplay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msi-selectplay",
			Port:  2871,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msi-selectplay",
			Port:  2871,
		},
	},
	"msiccp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msiccp",
			Port:  1731,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msiccp",
			Port:  1731,
		},
	},
	"msims": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msims",
			Port:  1582,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msims",
			Port:  1582,
		},
	},
	"msl_lmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msl_lmd",
			Port:  1464,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msl_lmd",
			Port:  1464,
		},
	},
	"msmq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msmq",
			Port:  1801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msmq",
			Port:  1801,
		},
	},
	"msnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msnp",
			Port:  1863,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msnp",
			Port:  1863,
		},
	},
	"msolap-ptp2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msolap-ptp2",
			Port:  2725,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msolap-ptp2",
			Port:  2725,
		},
	},
	"msp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msp",
			Port:  18,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msp",
			Port:  18,
		},
	},
	"msp-os": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msp-os",
			Port:  4686,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msp-os",
			Port:  4686,
		},
	},
	"msr-plugin-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msr-plugin-port",
			Port:  3931,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msr-plugin-port",
			Port:  3931,
		},
	},
	"msrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msrp",
			Port:  2855,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msrp",
			Port:  2855,
		},
	},
	"msss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msss",
			Port:  7742,
		},
	},
	"mstmg-sstp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mstmg-sstp",
			Port:  6601,
		},
	},
	"msync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "msync",
			Port:  2072,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "msync",
			Port:  2072,
		},
	},
	"mt-scaleserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mt-scaleserver",
			Port:  2305,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mt-scaleserver",
			Port:  2305,
		},
	},
	"mtcevrunqman": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mtcevrunqman",
			Port:  4558,
		},
	},
	"mtcevrunqss": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mtcevrunqss",
			Port:  4557,
		},
	},
	"mtftp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mtftp",
			Port:  1759,
		},
	},
	"mti-tcs-comm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mti-tcs-comm",
			Port:  2469,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mti-tcs-comm",
			Port:  2469,
		},
	},
	"mtl8000-matrix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mtl8000-matrix",
			Port:  8115,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mtl8000-matrix",
			Port:  8115,
		},
	},
	"mtn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mtn",
			Port:  4691,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mtn",
			Port:  4691,
		},
	},
	"mtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mtp",
			Port:  1911,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mtp",
			Port:  1911,
		},
	},
	"mtport-regist": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mtport-regist",
			Port:  2791,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mtport-regist",
			Port:  2791,
		},
	},
	"mtportmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mtportmon",
			Port:  7421,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mtportmon",
			Port:  7421,
		},
	},
	"mtqp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mtqp",
			Port:  1038,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mtqp",
			Port:  1038,
		},
	},
	"mtrgtrans": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mtrgtrans",
			Port:  19398,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mtrgtrans",
			Port:  19398,
		},
	},
	"mtsserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mtsserver",
			Port:  4602,
		},
	},
	"multicast-ping": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "multicast-ping",
			Port:  9903,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "multicast-ping",
			Port:  9903,
		},
	},
	"multiling-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "multiling-http",
			Port:  777,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "multiling-http",
			Port:  777,
		},
	},
	"multip-msg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "multip-msg",
			Port:  3733,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "multip-msg",
			Port:  3733,
		},
	},
	"multiplex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "multiplex",
			Port:  171,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "multiplex",
			Port:  171,
		},
	},
	"mumps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mumps",
			Port:  188,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mumps",
			Port:  188,
		},
	},
	"munin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "munin",
			Port:  4949,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "munin",
			Port:  4949,
		},
	},
	"mupdate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mupdate",
			Port:  3905,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mupdate",
			Port:  3905,
		},
	},
	"murray": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "murray",
			Port:  1123,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "murray",
			Port:  1123,
		},
	},
	"murx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "murx",
			Port:  2743,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "murx",
			Port:  2743,
		},
	},
	"muse": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "muse",
			Port:  6888,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "muse",
			Port:  6888,
		},
	},
	"musiconline": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "musiconline",
			Port:  1806,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "musiconline",
			Port:  1806,
		},
	},
	"must-backplane": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "must-backplane",
			Port:  3515,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "must-backplane",
			Port:  3515,
		},
	},
	"must-p2p": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "must-p2p",
			Port:  3514,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "must-p2p",
			Port:  3514,
		},
	},
	"mvel-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mvel-lm",
			Port:  1574,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mvel-lm",
			Port:  1574,
		},
	},
	"mvs-capacity": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mvs-capacity",
			Port:  10007,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mvs-capacity",
			Port:  10007,
		},
	},
	"mvx-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mvx-lm",
			Port:  1510,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mvx-lm",
			Port:  1510,
		},
	},
	"mxi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mxi",
			Port:  8005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mxi",
			Port:  8005,
		},
	},
	"mxit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mxit",
			Port:  9119,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mxit",
			Port:  9119,
		},
	},
	"mxodbc-connect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mxodbc-connect",
			Port:  6632,
		},
	},
	"mxomss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mxomss",
			Port:  1141,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mxomss",
			Port:  1141,
		},
	},
	"mxxrlogin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mxxrlogin",
			Port:  1035,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mxxrlogin",
			Port:  1035,
		},
	},
	"myblast": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "myblast",
			Port:  3795,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "myblast",
			Port:  3795,
		},
	},
	"mylex-mapd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mylex-mapd",
			Port:  467,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mylex-mapd",
			Port:  467,
		},
	},
	"mylxamport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mylxamport",
			Port:  2981,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mylxamport",
			Port:  2981,
		},
	},
	"mynahautostart": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mynahautostart",
			Port:  2388,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mynahautostart",
			Port:  2388,
		},
	},
	"myrtle": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "myrtle",
			Port:  1831,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "myrtle",
			Port:  1831,
		},
	},
	"mysql": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mysql",
			Port:  3306,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mysql",
			Port:  3306,
		},
	},
	"mysql-cluster": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mysql-cluster",
			Port:  1186,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mysql-cluster",
			Port:  1186,
		},
	},
	"mysql-cm-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mysql-cm-agent",
			Port:  1862,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mysql-cm-agent",
			Port:  1862,
		},
	},
	"mysql-im": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mysql-im",
			Port:  2273,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mysql-im",
			Port:  2273,
		},
	},
	"mysql-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mysql-proxy",
			Port:  6446,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mysql-proxy",
			Port:  6446,
		},
	},
	"mzap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mzap",
			Port:  2106,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "mzap",
			Port:  2106,
		},
	},
	"mzca-action": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "mzca-action",
			Port:  7282,
		},
	},
	"mzca-alert": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "mzca-alert",
			Port:  7282,
		},
	},
	"n1-fwp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "n1-fwp",
			Port:  4446,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "n1-fwp",
			Port:  4446,
		},
	},
	"n1-rmgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "n1-rmgmt",
			Port:  4447,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "n1-rmgmt",
			Port:  4447,
		},
	},
	"n2h2server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "n2h2server",
			Port:  9285,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "n2h2server",
			Port:  9285,
		},
	},
	"n2nremote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "n2nremote",
			Port:  1685,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "n2nremote",
			Port:  1685,
		},
	},
	"n2receive": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "n2receive",
			Port:  9286,
		},
	},
	"na-er-tip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "na-er-tip",
			Port:  3725,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "na-er-tip",
			Port:  3725,
		},
	},
	"na-localise": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "na-localise",
			Port:  5062,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "na-localise",
			Port:  5062,
		},
	},
	"naap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "naap",
			Port:  1340,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "naap",
			Port:  1340,
		},
	},
	"nacagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nacagent",
			Port:  4407,
		},
	},
	"nacnl": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "nacnl",
			Port:  4361,
		},
	},
	"namemunge": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "namemunge",
			Port:  3950,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "namemunge",
			Port:  3950,
		},
	},
	"nameserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nameserver",
			Port:  42,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nameserver",
			Port:  42,
		},
	},
	"namp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "namp",
			Port:  167,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "namp",
			Port:  167,
		},
	},
	"nani": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nani",
			Port:  2236,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nani",
			Port:  2236,
		},
	},
	"nas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nas",
			Port:  991,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nas",
			Port:  991,
		},
	},
	"nas-metering": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nas-metering",
			Port:  2286,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nas-metering",
			Port:  2286,
		},
	},
	"nasmanager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nasmanager",
			Port:  1960,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nasmanager",
			Port:  1960,
		},
	},
	"natdataservice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "natdataservice",
			Port:  3927,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "natdataservice",
			Port:  3927,
		},
	},
	"nati-dstp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nati-dstp",
			Port:  3015,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nati-dstp",
			Port:  3015,
		},
	},
	"nati-logos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nati-logos",
			Port:  2343,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nati-logos",
			Port:  2343,
		},
	},
	"nati-svrloc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nati-svrloc",
			Port:  3580,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nati-svrloc",
			Port:  3580,
		},
	},
	"nati-vi-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nati-vi-server",
			Port:  3363,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nati-vi-server",
			Port:  3363,
		},
	},
	"nattyserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nattyserver",
			Port:  3753,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nattyserver",
			Port:  3753,
		},
	},
	"natuslink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "natuslink",
			Port:  2895,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "natuslink",
			Port:  2895,
		},
	},
	"nav-data": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "nav-data",
			Port:  6317,
		},
	},
	"nav-data-cmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nav-data-cmd",
			Port:  6317,
		},
	},
	"nav-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nav-port",
			Port:  3859,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nav-port",
			Port:  3859,
		},
	},
	"navbuddy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "navbuddy",
			Port:  1288,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "navbuddy",
			Port:  1288,
		},
	},
	"navegaweb-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "navegaweb-port",
			Port:  3159,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "navegaweb-port",
			Port:  3159,
		},
	},
	"navisphere": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "navisphere",
			Port:  2162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "navisphere",
			Port:  2162,
		},
	},
	"navisphere-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "navisphere-sec",
			Port:  2163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "navisphere-sec",
			Port:  2163,
		},
	},
	"nbd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nbd",
			Port:  10809,
		},
	},
	"nbdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nbdb",
			Port:  13785,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nbdb",
			Port:  13785,
		},
	},
	"nbp": map[string]Service{
		"ddp": Service{
			Proto: "ddp",
			Name:  "nbp",
			Port:  2,
		},
	},
	"nbt-pc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nbt-pc",
			Port:  5133,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nbt-pc",
			Port:  5133,
		},
	},
	"nbt-wol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nbt-wol",
			Port:  6133,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nbt-wol",
			Port:  6133,
		},
	},
	"nburn_id": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nburn_id",
			Port:  20034,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nburn_id",
			Port:  20034,
		},
	},
	"nbx-au": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nbx-au",
			Port:  2094,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nbx-au",
			Port:  2094,
		},
	},
	"nbx-cc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nbx-cc",
			Port:  2093,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nbx-cc",
			Port:  2093,
		},
	},
	"nbx-dir": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nbx-dir",
			Port:  2096,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nbx-dir",
			Port:  2096,
		},
	},
	"nbx-ser": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nbx-ser",
			Port:  2095,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nbx-ser",
			Port:  2095,
		},
	},
	"ncacn-ip-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncacn-ip-tcp",
			Port:  3062,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncacn-ip-tcp",
			Port:  3062,
		},
	},
	"ncadg-ip-udp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncadg-ip-udp",
			Port:  3063,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncadg-ip-udp",
			Port:  3063,
		},
	},
	"ncconfig": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncconfig",
			Port:  1888,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncconfig",
			Port:  1888,
		},
	},
	"ncdloadbalance": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncdloadbalance",
			Port:  2683,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncdloadbalance",
			Port:  2683,
		},
	},
	"ncdmirroring": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncdmirroring",
			Port:  2706,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncdmirroring",
			Port:  2706,
		},
	},
	"nced": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nced",
			Port:  404,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nced",
			Port:  404,
		},
	},
	"ncl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncl",
			Port:  2397,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncl",
			Port:  2397,
		},
	},
	"ncld": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncld",
			Port:  405,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncld",
			Port:  405,
		},
	},
	"ncp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncp",
			Port:  524,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncp",
			Port:  524,
		},
	},
	"ncpm-ft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncpm-ft",
			Port:  1744,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncpm-ft",
			Port:  1744,
		},
	},
	"ncpm-hip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncpm-hip",
			Port:  1683,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncpm-hip",
			Port:  1683,
		},
	},
	"ncpm-pm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncpm-pm",
			Port:  1591,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncpm-pm",
			Port:  1591,
		},
	},
	"ncr_ccl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncr_ccl",
			Port:  2528,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncr_ccl",
			Port:  2528,
		},
	},
	"ncu-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncu-1",
			Port:  3195,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncu-1",
			Port:  3195,
		},
	},
	"ncu-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncu-2",
			Port:  3196,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncu-2",
			Port:  3196,
		},
	},
	"ncube-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncube-lm",
			Port:  1521,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncube-lm",
			Port:  1521,
		},
	},
	"ncxcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ncxcp",
			Port:  5681,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ncxcp",
			Port:  5681,
		},
	},
	"ndl-aas": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ndl-aas",
			Port:  3128,
		},
	},
	"ndl-ahp-svc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndl-ahp-svc",
			Port:  6064,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndl-ahp-svc",
			Port:  6064,
		},
	},
	"ndl-als": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndl-als",
			Port:  3431,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndl-als",
			Port:  3431,
		},
	},
	"ndl-aps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndl-aps",
			Port:  3096,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndl-aps",
			Port:  3096,
		},
	},
	"ndl-tcp-ois-gw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndl-tcp-ois-gw",
			Port:  2738,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndl-tcp-ois-gw",
			Port:  2738,
		},
	},
	"ndm-agent-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndm-agent-port",
			Port:  43189,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndm-agent-port",
			Port:  43189,
		},
	},
	"ndm-requester": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndm-requester",
			Port:  1363,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndm-requester",
			Port:  1363,
		},
	},
	"ndm-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndm-server",
			Port:  1364,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndm-server",
			Port:  1364,
		},
	},
	"ndmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndmp",
			Port:  10000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndmp",
			Port:  10000,
		},
	},
	"ndmps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndmps",
			Port:  30000,
		},
	},
	"ndnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndnp",
			Port:  2883,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndnp",
			Port:  2883,
		},
	},
	"nds_sso": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nds_sso",
			Port:  3024,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nds_sso",
			Port:  3024,
		},
	},
	"ndsauth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndsauth",
			Port:  353,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndsauth",
			Port:  353,
		},
	},
	"ndsconnect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndsconnect",
			Port:  3890,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndsconnect",
			Port:  3890,
		},
	},
	"ndsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndsp",
			Port:  2881,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndsp",
			Port:  2881,
		},
	},
	"ndtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ndtp",
			Port:  2882,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ndtp",
			Port:  2882,
		},
	},
	"nec-raidplus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nec-raidplus",
			Port:  2730,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nec-raidplus",
			Port:  2730,
		},
	},
	"neckar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neckar",
			Port:  37475,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neckar",
			Port:  37475,
		},
	},
	"necp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "necp",
			Port:  3262,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "necp",
			Port:  3262,
		},
	},
	"nei-management": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nei-management",
			Port:  3886,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nei-management",
			Port:  3886,
		},
	},
	"neo4j": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neo4j",
			Port:  7474,
		},
	},
	"neod1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neod1",
			Port:  1047,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neod1",
			Port:  1047,
		},
	},
	"neod2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neod2",
			Port:  1048,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neod2",
			Port:  1048,
		},
	},
	"neoiface": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neoiface",
			Port:  1285,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neoiface",
			Port:  1285,
		},
	},
	"neon24x7": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neon24x7",
			Port:  3213,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neon24x7",
			Port:  3213,
		},
	},
	"nerv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nerv",
			Port:  1222,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nerv",
			Port:  1222,
		},
	},
	"nesh-broker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nesh-broker",
			Port:  3507,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nesh-broker",
			Port:  3507,
		},
	},
	"nessus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nessus",
			Port:  1241,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nessus",
			Port:  1241,
		},
	},
	"nest-protocol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nest-protocol",
			Port:  489,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nest-protocol",
			Port:  489,
		},
	},
	"net-assistant": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "net-assistant",
			Port:  3283,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "net-assistant",
			Port:  3283,
		},
	},
	"net-device": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "net-device",
			Port:  4350,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "net-device",
			Port:  4350,
		},
	},
	"net-projection": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "net-projection",
			Port:  5363,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "net-projection",
			Port:  5363,
		},
	},
	"net-steward": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "net-steward",
			Port:  2128,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "net-steward",
			Port:  2128,
		},
	},
	"net2display": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "net2display",
			Port:  9086,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "net2display",
			Port:  9086,
		},
	},
	"net8-cman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "net8-cman",
			Port:  1830,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "net8-cman",
			Port:  1830,
		},
	},
	"netadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netadmin",
			Port:  2450,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netadmin",
			Port:  2450,
		},
	},
	"netagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netagent",
			Port:  5771,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netagent",
			Port:  5771,
		},
	},
	"netangel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netangel",
			Port:  2442,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netangel",
			Port:  2442,
		},
	},
	"netapp-icdata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netapp-icdata",
			Port:  11105,
		},
	},
	"netapp-icmgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netapp-icmgmt",
			Port:  11104,
		},
	},
	"netarx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netarx",
			Port:  1040,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netarx",
			Port:  1040,
		},
	},
	"netaspi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netaspi",
			Port:  2902,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netaspi",
			Port:  2902,
		},
	},
	"netattachsdmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netattachsdmp",
			Port:  3066,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netattachsdmp",
			Port:  3066,
		},
	},
	"netbill-auth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbill-auth",
			Port:  1615,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbill-auth",
			Port:  1615,
		},
	},
	"netbill-cred": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbill-cred",
			Port:  1614,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbill-cred",
			Port:  1614,
		},
	},
	"netbill-keyrep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbill-keyrep",
			Port:  1613,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbill-keyrep",
			Port:  1613,
		},
	},
	"netbill-prod": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbill-prod",
			Port:  1616,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbill-prod",
			Port:  1616,
		},
	},
	"netbill-trans": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbill-trans",
			Port:  1612,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbill-trans",
			Port:  1612,
		},
	},
	"netbios-dgm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbios-dgm",
			Port:  138,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbios-dgm",
			Port:  138,
		},
	},
	"netbios-ns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbios-ns",
			Port:  137,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbios-ns",
			Port:  137,
		},
	},
	"netbios-ssn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbios-ssn",
			Port:  139,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbios-ssn",
			Port:  139,
		},
	},
	"netblox": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "netblox",
			Port:  4441,
		},
	},
	"netbookmark": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netbookmark",
			Port:  3131,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netbookmark",
			Port:  3131,
		},
	},
	"netboot-pxe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netboot-pxe",
			Port:  3928,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netboot-pxe",
			Port:  3928,
		},
	},
	"netcabinet-com": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netcabinet-com",
			Port:  4409,
		},
	},
	"netcelera": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netcelera",
			Port:  3701,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netcelera",
			Port:  3701,
		},
	},
	"netchat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netchat",
			Port:  2451,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netchat",
			Port:  2451,
		},
	},
	"netcheque": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netcheque",
			Port:  4008,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netcheque",
			Port:  4008,
		},
	},
	"netclip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netclip",
			Port:  2971,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netclip",
			Port:  2971,
		},
	},
	"netcomm1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netcomm1",
			Port:  1676,
		},
	},
	"netcomm2": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "netcomm2",
			Port:  1676,
		},
	},
	"netconf-beep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netconf-beep",
			Port:  831,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netconf-beep",
			Port:  831,
		},
	},
	"netconf-ssh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netconf-ssh",
			Port:  830,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netconf-ssh",
			Port:  830,
		},
	},
	"netconf-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netconf-tls",
			Port:  6513,
		},
	},
	"netconfsoapbeep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netconfsoapbeep",
			Port:  833,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netconfsoapbeep",
			Port:  833,
		},
	},
	"netconfsoaphttp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netconfsoaphttp",
			Port:  832,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netconfsoaphttp",
			Port:  832,
		},
	},
	"netcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netcp",
			Port:  395,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netcp",
			Port:  395,
		},
	},
	"netdb-export": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netdb-export",
			Port:  1329,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netdb-export",
			Port:  1329,
		},
	},
	"neteh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neteh",
			Port:  3828,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neteh",
			Port:  3828,
		},
	},
	"neteh-ext": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neteh-ext",
			Port:  3829,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neteh-ext",
			Port:  3829,
		},
	},
	"netgw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netgw",
			Port:  741,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netgw",
			Port:  741,
		},
	},
	"netinfo-local": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netinfo-local",
			Port:  1033,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netinfo-local",
			Port:  1033,
		},
	},
	"netiq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netiq",
			Port:  2220,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netiq",
			Port:  2220,
		},
	},
	"netiq-endpoint": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netiq-endpoint",
			Port:  10113,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netiq-endpoint",
			Port:  10113,
		},
	},
	"netiq-endpt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netiq-endpt",
			Port:  10115,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netiq-endpt",
			Port:  10115,
		},
	},
	"netiq-mc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netiq-mc",
			Port:  2735,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netiq-mc",
			Port:  2735,
		},
	},
	"netiq-ncap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netiq-ncap",
			Port:  2219,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netiq-ncap",
			Port:  2219,
		},
	},
	"netiq-qcheck": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netiq-qcheck",
			Port:  10114,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netiq-qcheck",
			Port:  10114,
		},
	},
	"netiq-voipa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netiq-voipa",
			Port:  10116,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netiq-voipa",
			Port:  10116,
		},
	},
	"netlabs-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netlabs-lm",
			Port:  1406,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netlabs-lm",
			Port:  1406,
		},
	},
	"netmagic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netmagic",
			Port:  1196,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netmagic",
			Port:  1196,
		},
	},
	"netmap_lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netmap_lm",
			Port:  1493,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netmap_lm",
			Port:  1493,
		},
	},
	"netml": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netml",
			Port:  2288,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netml",
			Port:  2288,
		},
	},
	"netmo-default": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netmo-default",
			Port:  6841,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netmo-default",
			Port:  6841,
		},
	},
	"netmo-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netmo-http",
			Port:  6842,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netmo-http",
			Port:  6842,
		},
	},
	"netmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netmon",
			Port:  2606,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netmon",
			Port:  2606,
		},
	},
	"netmount": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netmount",
			Port:  2061,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netmount",
			Port:  2061,
		},
	},
	"netmpi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netmpi",
			Port:  3827,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netmpi",
			Port:  3827,
		},
	},
	"netnews": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netnews",
			Port:  532,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netnews",
			Port:  532,
		},
	},
	"neto-dcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neto-dcs",
			Port:  3814,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neto-dcs",
			Port:  3814,
		},
	},
	"neto-wol-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neto-wol-server",
			Port:  3812,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neto-wol-server",
			Port:  3812,
		},
	},
	"netobjects1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netobjects1",
			Port:  2485,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netobjects1",
			Port:  2485,
		},
	},
	"netobjects2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netobjects2",
			Port:  2486,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netobjects2",
			Port:  2486,
		},
	},
	"netop-rc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netop-rc",
			Port:  1970,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netop-rc",
			Port:  1970,
		},
	},
	"netop-school": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netop-school",
			Port:  1971,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netop-school",
			Port:  1971,
		},
	},
	"netopia-vo1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netopia-vo1",
			Port:  1839,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netopia-vo1",
			Port:  1839,
		},
	},
	"netopia-vo2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netopia-vo2",
			Port:  1840,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netopia-vo2",
			Port:  1840,
		},
	},
	"netopia-vo3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netopia-vo3",
			Port:  1841,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netopia-vo3",
			Port:  1841,
		},
	},
	"netopia-vo4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netopia-vo4",
			Port:  1842,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netopia-vo4",
			Port:  1842,
		},
	},
	"netopia-vo5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netopia-vo5",
			Port:  1843,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netopia-vo5",
			Port:  1843,
		},
	},
	"netops-broker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netops-broker",
			Port:  5465,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netops-broker",
			Port:  5465,
		},
	},
	"netperf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netperf",
			Port:  12865,
		},
	},
	"netplan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netplan",
			Port:  2983,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netplan",
			Port:  2983,
		},
	},
	"netplay-port1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netplay-port1",
			Port:  3640,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netplay-port1",
			Port:  3640,
		},
	},
	"netplay-port2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netplay-port2",
			Port:  3641,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netplay-port2",
			Port:  3641,
		},
	},
	"netport-id": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netport-id",
			Port:  3129,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netport-id",
			Port:  3129,
		},
	},
	"netrcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrcs",
			Port:  742,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrcs",
			Port:  742,
		},
	},
	"netrek": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrek",
			Port:  2592,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrek",
			Port:  2592,
		},
	},
	"netrisk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrisk",
			Port:  1799,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrisk",
			Port:  1799,
		},
	},
	"netrix-sftm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrix-sftm",
			Port:  2328,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrix-sftm",
			Port:  2328,
		},
	},
	"netrjs-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrjs-1",
			Port:  71,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrjs-1",
			Port:  71,
		},
	},
	"netrjs-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrjs-2",
			Port:  72,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrjs-2",
			Port:  72,
		},
	},
	"netrjs-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrjs-3",
			Port:  73,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrjs-3",
			Port:  73,
		},
	},
	"netrjs-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrjs-4",
			Port:  74,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrjs-4",
			Port:  74,
		},
	},
	"netrockey6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netrockey6",
			Port:  4425,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netrockey6",
			Port:  4425,
		},
	},
	"netsc-dev": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netsc-dev",
			Port:  155,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netsc-dev",
			Port:  155,
		},
	},
	"netsc-prod": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netsc-prod",
			Port:  154,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netsc-prod",
			Port:  154,
		},
	},
	"netscript": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netscript",
			Port:  4118,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netscript",
			Port:  4118,
		},
	},
	"netserialext1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netserialext1",
			Port:  16360,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netserialext1",
			Port:  16360,
		},
	},
	"netserialext2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netserialext2",
			Port:  16361,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netserialext2",
			Port:  16361,
		},
	},
	"netserialext3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netserialext3",
			Port:  16367,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netserialext3",
			Port:  16367,
		},
	},
	"netserialext4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netserialext4",
			Port:  16368,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netserialext4",
			Port:  16368,
		},
	},
	"netspeak-acd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netspeak-acd",
			Port:  21848,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netspeak-acd",
			Port:  21848,
		},
	},
	"netspeak-cps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netspeak-cps",
			Port:  21849,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netspeak-cps",
			Port:  21849,
		},
	},
	"netspeak-cs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netspeak-cs",
			Port:  21847,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netspeak-cs",
			Port:  21847,
		},
	},
	"netspeak-is": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netspeak-is",
			Port:  21846,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netspeak-is",
			Port:  21846,
		},
	},
	"netstat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netstat",
			Port:  15,
		},
	},
	"netsteward": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netsteward",
			Port:  2810,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netsteward",
			Port:  2810,
		},
	},
	"netsupport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netsupport",
			Port:  5405,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netsupport",
			Port:  5405,
		},
	},
	"netsupport2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netsupport2",
			Port:  5421,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netsupport2",
			Port:  5421,
		},
	},
	"nettest": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nettest",
			Port:  4138,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nettest",
			Port:  4138,
		},
	},
	"nettgain-nms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nettgain-nms",
			Port:  1879,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nettgain-nms",
			Port:  1879,
		},
	},
	"netuitive": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netuitive",
			Port:  1286,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netuitive",
			Port:  1286,
		},
	},
	"netview-aix-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-1",
			Port:  1661,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-1",
			Port:  1661,
		},
	},
	"netview-aix-10": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-10",
			Port:  1670,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-10",
			Port:  1670,
		},
	},
	"netview-aix-11": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-11",
			Port:  1671,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-11",
			Port:  1671,
		},
	},
	"netview-aix-12": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-12",
			Port:  1672,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-12",
			Port:  1672,
		},
	},
	"netview-aix-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-2",
			Port:  1662,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-2",
			Port:  1662,
		},
	},
	"netview-aix-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-3",
			Port:  1663,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-3",
			Port:  1663,
		},
	},
	"netview-aix-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-4",
			Port:  1664,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-4",
			Port:  1664,
		},
	},
	"netview-aix-5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-5",
			Port:  1665,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-5",
			Port:  1665,
		},
	},
	"netview-aix-6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-6",
			Port:  1666,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-6",
			Port:  1666,
		},
	},
	"netview-aix-7": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-7",
			Port:  1667,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-7",
			Port:  1667,
		},
	},
	"netview-aix-8": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-8",
			Port:  1668,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-8",
			Port:  1668,
		},
	},
	"netview-aix-9": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netview-aix-9",
			Port:  1669,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netview-aix-9",
			Port:  1669,
		},
	},
	"netviewdm1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netviewdm1",
			Port:  729,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netviewdm1",
			Port:  729,
		},
	},
	"netviewdm2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netviewdm2",
			Port:  730,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netviewdm2",
			Port:  730,
		},
	},
	"netviewdm3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netviewdm3",
			Port:  731,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netviewdm3",
			Port:  731,
		},
	},
	"netwall": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netwall",
			Port:  533,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netwall",
			Port:  533,
		},
	},
	"netware-csp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netware-csp",
			Port:  1366,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netware-csp",
			Port:  1366,
		},
	},
	"netware-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netware-ip",
			Port:  396,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netware-ip",
			Port:  396,
		},
	},
	"netwatcher-db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netwatcher-db",
			Port:  3204,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netwatcher-db",
			Port:  3204,
		},
	},
	"netwatcher-mon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netwatcher-mon",
			Port:  3203,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netwatcher-mon",
			Port:  3203,
		},
	},
	"netwave-ap-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netwave-ap-mgmt",
			Port:  2411,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netwave-ap-mgmt",
			Port:  2411,
		},
	},
	"netwkpathengine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netwkpathengine",
			Port:  3209,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netwkpathengine",
			Port:  3209,
		},
	},
	"networklens": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "networklens",
			Port:  3409,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "networklens",
			Port:  3409,
		},
	},
	"networklenss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "networklenss",
			Port:  3410,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "networklenss",
			Port:  3410,
		},
	},
	"netx-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netx-agent",
			Port:  2586,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netx-agent",
			Port:  2586,
		},
	},
	"netx-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netx-server",
			Port:  2585,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netx-server",
			Port:  2585,
		},
	},
	"netxms-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netxms-agent",
			Port:  4700,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netxms-agent",
			Port:  4700,
		},
	},
	"netxms-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netxms-mgmt",
			Port:  4701,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netxms-mgmt",
			Port:  4701,
		},
	},
	"netxms-sync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "netxms-sync",
			Port:  4702,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "netxms-sync",
			Port:  4702,
		},
	},
	"neveroffline": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "neveroffline",
			Port:  2614,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "neveroffline",
			Port:  2614,
		},
	},
	"new-rwho": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "new-rwho",
			Port:  550,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "new-rwho",
			Port:  550,
		},
	},
	"newacct": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newacct",
			Port:  100,
		},
	},
	"newbay-snc-mc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newbay-snc-mc",
			Port:  16900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newbay-snc-mc",
			Port:  16900,
		},
	},
	"newgenpay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newgenpay",
			Port:  3165,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newgenpay",
			Port:  3165,
		},
	},
	"newheights": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newheights",
			Port:  2114,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newheights",
			Port:  2114,
		},
	},
	"newlixconfig": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newlixconfig",
			Port:  2076,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newlixconfig",
			Port:  2076,
		},
	},
	"newlixengine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newlixengine",
			Port:  2075,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newlixengine",
			Port:  2075,
		},
	},
	"newlixreg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newlixreg",
			Port:  2671,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newlixreg",
			Port:  2671,
		},
	},
	"newoak": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newoak",
			Port:  4001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newoak",
			Port:  4001,
		},
	},
	"news": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "news",
			Port:  2009,
		},
	},
	"newton-dock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newton-dock",
			Port:  3679,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newton-dock",
			Port:  3679,
		},
	},
	"newwavesearch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "newwavesearch",
			Port:  2058,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "newwavesearch",
			Port:  2058,
		},
	},
	"nexentamv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nexentamv",
			Port:  8457,
		},
	},
	"nexgen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nexgen",
			Port:  6627,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nexgen",
			Port:  6627,
		},
	},
	"nexstorindltd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nexstorindltd",
			Port:  2360,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nexstorindltd",
			Port:  2360,
		},
	},
	"nextstep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nextstep",
			Port:  178,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nextstep",
			Port:  178,
		},
	},
	"nexus-portal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nexus-portal",
			Port:  4021,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nexus-portal",
			Port:  4021,
		},
	},
	"nfa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nfa",
			Port:  1155,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nfa",
			Port:  1155,
		},
	},
	"nfoldman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nfoldman",
			Port:  7393,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nfoldman",
			Port:  7393,
		},
	},
	"nfs": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "nfs",
			Port:  2049,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "nfs",
			Port:  2049,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nfs",
			Port:  2049,
		},
	},
	"nfsd-keepalive": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "nfsd-keepalive",
			Port:  1110,
		},
	},
	"nfsrdma": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "nfsrdma",
			Port:  20049,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "nfsrdma",
			Port:  20049,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nfsrdma",
			Port:  20049,
		},
	},
	"ng-umds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ng-umds",
			Port:  1690,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ng-umds",
			Port:  1690,
		},
	},
	"nhci": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nhci",
			Port:  3842,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nhci",
			Port:  3842,
		},
	},
	"nhserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nhserver",
			Port:  2672,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nhserver",
			Port:  2672,
		},
	},
	"ni-ftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ni-ftp",
			Port:  47,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ni-ftp",
			Port:  47,
		},
	},
	"ni-mail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ni-mail",
			Port:  61,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ni-mail",
			Port:  61,
		},
	},
	"ni-visa-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ni-visa-remote",
			Port:  3537,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ni-visa-remote",
			Port:  3537,
		},
	},
	"nicelink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nicelink",
			Port:  1095,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nicelink",
			Port:  1095,
		},
	},
	"nicetec-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nicetec-mgmt",
			Port:  2557,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nicetec-mgmt",
			Port:  2557,
		},
	},
	"nicetec-nmsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nicetec-nmsvc",
			Port:  2556,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nicetec-nmsvc",
			Port:  2556,
		},
	},
	"nicname": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nicname",
			Port:  43,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nicname",
			Port:  43,
		},
	},
	"nifty-hmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nifty-hmi",
			Port:  4134,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nifty-hmi",
			Port:  4134,
		},
	},
	"nilinkanalyst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nilinkanalyst",
			Port:  25902,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nilinkanalyst",
			Port:  25902,
		},
	},
	"nim": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nim",
			Port:  1058,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nim",
			Port:  1058,
		},
	},
	"nim-vdrshell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nim-vdrshell",
			Port:  6420,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nim-vdrshell",
			Port:  6420,
		},
	},
	"nim-wan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nim-wan",
			Port:  6421,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nim-wan",
			Port:  6421,
		},
	},
	"nimaux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimaux",
			Port:  3902,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nimaux",
			Port:  3902,
		},
	},
	"nimbusdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimbusdb",
			Port:  48004,
		},
	},
	"nimbusdbctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimbusdbctrl",
			Port:  48005,
		},
	},
	"nimcontroller": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimcontroller",
			Port:  48000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nimcontroller",
			Port:  48000,
		},
	},
	"nimgtw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimgtw",
			Port:  48003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nimgtw",
			Port:  48003,
		},
	},
	"nimhub": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimhub",
			Port:  48002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nimhub",
			Port:  48002,
		},
	},
	"nimreg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimreg",
			Port:  1059,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nimreg",
			Port:  1059,
		},
	},
	"nimrod-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimrod-agent",
			Port:  1617,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nimrod-agent",
			Port:  1617,
		},
	},
	"nimsh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimsh",
			Port:  3901,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nimsh",
			Port:  3901,
		},
	},
	"nimspooler": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nimspooler",
			Port:  48001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nimspooler",
			Port:  48001,
		},
	},
	"ninaf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ninaf",
			Port:  5627,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ninaf",
			Port:  5627,
		},
	},
	"ninstall": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ninstall",
			Port:  2150,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ninstall",
			Port:  2150,
		},
	},
	"niobserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "niobserver",
			Port:  25901,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "niobserver",
			Port:  25901,
		},
	},
	"nip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nip",
			Port:  376,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nip",
			Port:  376,
		},
	},
	"niprobe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "niprobe",
			Port:  25903,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "niprobe",
			Port:  25903,
		},
	},
	"nirp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nirp",
			Port:  4043,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nirp",
			Port:  4043,
		},
	},
	"nitrogen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nitrogen",
			Port:  7725,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nitrogen",
			Port:  7725,
		},
	},
	"njenet-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "njenet-ssl",
			Port:  2252,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "njenet-ssl",
			Port:  2252,
		},
	},
	"nkd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nkd",
			Port:  1650,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nkd",
			Port:  1650,
		},
	},
	"nlg-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nlg-data",
			Port:  5299,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nlg-data",
			Port:  5299,
		},
	},
	"nlogin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nlogin",
			Port:  758,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nlogin",
			Port:  758,
		},
	},
	"nls-tl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nls-tl",
			Port:  7549,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nls-tl",
			Port:  7549,
		},
	},
	"nm-asses-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nm-asses-admin",
			Port:  3150,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nm-asses-admin",
			Port:  3150,
		},
	},
	"nm-assessor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nm-assessor",
			Port:  3151,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nm-assessor",
			Port:  3151,
		},
	},
	"nm-game-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nm-game-admin",
			Port:  3148,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nm-game-admin",
			Port:  3148,
		},
	},
	"nm-game-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nm-game-server",
			Port:  3149,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nm-game-server",
			Port:  3149,
		},
	},
	"nmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nmap",
			Port:  689,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nmap",
			Port:  689,
		},
	},
	"nmasoverip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nmasoverip",
			Port:  1242,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nmasoverip",
			Port:  1242,
		},
	},
	"nmc-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "nmc-disc",
			Port:  10810,
		},
	},
	"nmea-0183": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nmea-0183",
			Port:  10110,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nmea-0183",
			Port:  10110,
		},
	},
	"nmea-onenet": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "nmea-onenet",
			Port:  10111,
		},
	},
	"nmmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nmmp",
			Port:  3649,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nmmp",
			Port:  3649,
		},
	},
	"nms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nms",
			Port:  1429,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nms",
			Port:  1429,
		},
	},
	"nms-dpnss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nms-dpnss",
			Port:  2503,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nms-dpnss",
			Port:  2503,
		},
	},
	"nms_topo_serv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nms_topo_serv",
			Port:  1486,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nms_topo_serv",
			Port:  1486,
		},
	},
	"nmsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nmsd",
			Port:  1239,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nmsd",
			Port:  1239,
		},
	},
	"nmsigport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nmsigport",
			Port:  2817,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nmsigport",
			Port:  2817,
		},
	},
	"nmsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nmsp",
			Port:  537,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nmsp",
			Port:  537,
		},
	},
	"nmsserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nmsserver",
			Port:  2244,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nmsserver",
			Port:  2244,
		},
	},
	"nnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nnp",
			Port:  3780,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nnp",
			Port:  3780,
		},
	},
	"nnsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nnsp",
			Port:  433,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nnsp",
			Port:  433,
		},
	},
	"nntp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nntp",
			Port:  119,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nntp",
			Port:  119,
		},
	},
	"nntps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nntps",
			Port:  563,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nntps",
			Port:  563,
		},
	},
	"noaaport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "noaaport",
			Port:  2210,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "noaaport",
			Port:  2210,
		},
	},
	"noadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "noadmin",
			Port:  1921,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "noadmin",
			Port:  1921,
		},
	},
	"noagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "noagent",
			Port:  1917,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "noagent",
			Port:  1917,
		},
	},
	"noit-transport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "noit-transport",
			Port:  43191,
		},
	},
	"nokia-ann-ch1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nokia-ann-ch1",
			Port:  3405,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nokia-ann-ch1",
			Port:  3405,
		},
	},
	"nokia-ann-ch2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nokia-ann-ch2",
			Port:  3406,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nokia-ann-ch2",
			Port:  3406,
		},
	},
	"nomad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nomad",
			Port:  5209,
		},
	},
	"nomdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nomdb",
			Port:  13786,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nomdb",
			Port:  13786,
		},
	},
	"norton-lambert": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "norton-lambert",
			Port:  2338,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "norton-lambert",
			Port:  2338,
		},
	},
	"notateit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "notateit",
			Port:  4803,
		},
	},
	"notateit-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "notateit-disc",
			Port:  4803,
		},
	},
	"noteit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "noteit",
			Port:  4663,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "noteit",
			Port:  4663,
		},
	},
	"noteshare": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "noteshare",
			Port:  8474,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "noteshare",
			Port:  8474,
		},
	},
	"notify": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "notify",
			Port:  773,
		},
	},
	"notify_srvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "notify_srvr",
			Port:  3016,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "notify_srvr",
			Port:  3016,
		},
	},
	"novar-alarm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "novar-alarm",
			Port:  23401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "novar-alarm",
			Port:  23401,
		},
	},
	"novar-dbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "novar-dbase",
			Port:  23400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "novar-dbase",
			Port:  23400,
		},
	},
	"novar-global": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "novar-global",
			Port:  23402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "novar-global",
			Port:  23402,
		},
	},
	"novastorbakcup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "novastorbakcup",
			Port:  308,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "novastorbakcup",
			Port:  308,
		},
	},
	"novation": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "novation",
			Port:  1322,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "novation",
			Port:  1322,
		},
	},
	"novell-ipx-cmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "novell-ipx-cmd",
			Port:  2645,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "novell-ipx-cmd",
			Port:  2645,
		},
	},
	"novell-lu6.2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "novell-lu6.2",
			Port:  1416,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "novell-lu6.2",
			Port:  1416,
		},
	},
	"novell-zen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "novell-zen",
			Port:  2544,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "novell-zen",
			Port:  2544,
		},
	},
	"nowcontact": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nowcontact",
			Port:  3167,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nowcontact",
			Port:  3167,
		},
	},
	"npdbgmngr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npdbgmngr",
			Port:  2293,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "npdbgmngr",
			Port:  2293,
		},
	},
	"npds-tracker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npds-tracker",
			Port:  3680,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "npds-tracker",
			Port:  3680,
		},
	},
	"npep-messaging": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npep-messaging",
			Port:  2868,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "npep-messaging",
			Port:  2868,
		},
	},
	"npmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npmp",
			Port:  8450,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "npmp",
			Port:  8450,
		},
	},
	"npmp-gui": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npmp-gui",
			Port:  611,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "npmp-gui",
			Port:  611,
		},
	},
	"npmp-local": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npmp-local",
			Port:  610,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "npmp-local",
			Port:  610,
		},
	},
	"npmp-trap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npmp-trap",
			Port:  609,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "npmp-trap",
			Port:  609,
		},
	},
	"npp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npp",
			Port:  92,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "npp",
			Port:  92,
		},
	},
	"nppmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nppmp",
			Port:  3476,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nppmp",
			Port:  3476,
		},
	},
	"npqes-test": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npqes-test",
			Port:  4703,
		},
	},
	"npsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "npsp",
			Port:  4088,
		},
	},
	"nqs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nqs",
			Port:  607,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nqs",
			Port:  607,
		},
	},
	"nrcabq-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nrcabq-lm",
			Port:  1458,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nrcabq-lm",
			Port:  1458,
		},
	},
	"ns": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ns",
			Port:  760,
		},
	},
	"ns-cfg-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ns-cfg-server",
			Port:  3266,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ns-cfg-server",
			Port:  3266,
		},
	},
	"ns-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ns-server",
			Port:  5415,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ns-server",
			Port:  5415,
		},
	},
	"nsc-ccs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsc-ccs",
			Port:  2604,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsc-ccs",
			Port:  2604,
		},
	},
	"nsc-posa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsc-posa",
			Port:  2605,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsc-posa",
			Port:  2605,
		},
	},
	"nsdeepfreezectl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsdeepfreezectl",
			Port:  7724,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsdeepfreezectl",
			Port:  7724,
		},
	},
	"nsesrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsesrvr",
			Port:  9988,
		},
	},
	"nsiiops": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsiiops",
			Port:  261,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsiiops",
			Port:  261,
		},
	},
	"nsjtp-ctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsjtp-ctrl",
			Port:  1687,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsjtp-ctrl",
			Port:  1687,
		},
	},
	"nsjtp-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsjtp-data",
			Port:  1688,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsjtp-data",
			Port:  1688,
		},
	},
	"nsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsp",
			Port:  5012,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsp",
			Port:  5012,
		},
	},
	"nsrmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsrmp",
			Port:  359,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsrmp",
			Port:  359,
		},
	},
	"nsrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsrp",
			Port:  7170,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsrp",
			Port:  7170,
		},
	},
	"nss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nss",
			Port:  4159,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nss",
			Port:  4159,
		},
	},
	"nss-routing": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nss-routing",
			Port:  159,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nss-routing",
			Port:  159,
		},
	},
	"nssagentmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nssagentmgr",
			Port:  4454,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nssagentmgr",
			Port:  4454,
		},
	},
	"nssalertmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nssalertmgr",
			Port:  4453,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nssalertmgr",
			Port:  4453,
		},
	},
	"nssocketport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nssocketport",
			Port:  3522,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nssocketport",
			Port:  3522,
		},
	},
	"nsstp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsstp",
			Port:  1036,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsstp",
			Port:  1036,
		},
	},
	"nst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nst",
			Port:  4687,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nst",
			Port:  4687,
		},
	},
	"nsw-fe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsw-fe",
			Port:  27,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsw-fe",
			Port:  27,
		},
	},
	"nsws": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nsws",
			Port:  3049,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nsws",
			Port:  3049,
		},
	},
	"nta-ds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nta-ds",
			Port:  7544,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nta-ds",
			Port:  7544,
		},
	},
	"nta-us": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nta-us",
			Port:  7545,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nta-us",
			Port:  7545,
		},
	},
	"ntalk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ntalk",
			Port:  518,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ntalk",
			Port:  518,
		},
	},
	"ntp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ntp",
			Port:  123,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ntp",
			Port:  123,
		},
	},
	"nuauth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nuauth",
			Port:  4129,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nuauth",
			Port:  4129,
		},
	},
	"nucleus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nucleus",
			Port:  1463,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nucleus",
			Port:  1463,
		},
	},
	"nucleus-sand": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nucleus-sand",
			Port:  1201,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nucleus-sand",
			Port:  1201,
		},
	},
	"nufw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nufw",
			Port:  4128,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nufw",
			Port:  4128,
		},
	},
	"nupaper-ss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nupaper-ss",
			Port:  12121,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nupaper-ss",
			Port:  12121,
		},
	},
	"nut": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nut",
			Port:  3493,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nut",
			Port:  3493,
		},
	},
	"nuts_bootp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nuts_bootp",
			Port:  4133,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nuts_bootp",
			Port:  4133,
		},
	},
	"nuts_dem": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nuts_dem",
			Port:  4132,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nuts_dem",
			Port:  4132,
		},
	},
	"nuxsl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nuxsl",
			Port:  5991,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nuxsl",
			Port:  5991,
		},
	},
	"nvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nvc",
			Port:  8711,
		},
	},
	"nvcnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nvcnet",
			Port:  3999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nvcnet",
			Port:  3999,
		},
	},
	"nvd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nvd",
			Port:  2184,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nvd",
			Port:  2184,
		},
	},
	"nvmsgd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nvmsgd",
			Port:  3519,
		},
	},
	"nw-license": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nw-license",
			Port:  3697,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nw-license",
			Port:  3697,
		},
	},
	"nxedit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nxedit",
			Port:  126,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nxedit",
			Port:  126,
		},
	},
	"nxlmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "nxlmd",
			Port:  28000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "nxlmd",
			Port:  28000,
		},
	},
	"o2server-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "o2server-port",
			Port:  1894,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "o2server-port",
			Port:  1894,
		},
	},
	"oa-system": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oa-system",
			Port:  8022,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oa-system",
			Port:  8022,
		},
	},
	"obex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "obex",
			Port:  650,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "obex",
			Port:  650,
		},
	},
	"objcall": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "objcall",
			Port:  94,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "objcall",
			Port:  94,
		},
	},
	"objective-dbc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "objective-dbc",
			Port:  1388,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "objective-dbc",
			Port:  1388,
		},
	},
	"objectmanager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "objectmanager",
			Port:  2038,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "objectmanager",
			Port:  2038,
		},
	},
	"obrpd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "obrpd",
			Port:  1092,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "obrpd",
			Port:  1092,
		},
	},
	"oc-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oc-lm",
			Port:  1448,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oc-lm",
			Port:  1448,
		},
	},
	"ocbinder": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ocbinder",
			Port:  183,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ocbinder",
			Port:  183,
		},
	},
	"oce-snmp-trap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oce-snmp-trap",
			Port:  2697,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oce-snmp-trap",
			Port:  2697,
		},
	},
	"oceansoft-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oceansoft-lm",
			Port:  1466,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oceansoft-lm",
			Port:  1466,
		},
	},
	"ocs_amu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ocs_amu",
			Port:  429,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ocs_amu",
			Port:  429,
		},
	},
	"ocs_cmu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ocs_cmu",
			Port:  428,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ocs_cmu",
			Port:  428,
		},
	},
	"ocserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ocserver",
			Port:  184,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ocserver",
			Port:  184,
		},
	},
	"octopus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "octopus",
			Port:  10008,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "octopus",
			Port:  10008,
		},
	},
	"odbcpathway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odbcpathway",
			Port:  9628,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odbcpathway",
			Port:  9628,
		},
	},
	"odette-ftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odette-ftp",
			Port:  3305,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odette-ftp",
			Port:  3305,
		},
	},
	"odette-ftps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odette-ftps",
			Port:  6619,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odette-ftps",
			Port:  6619,
		},
	},
	"odeumservlink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odeumservlink",
			Port:  3523,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odeumservlink",
			Port:  3523,
		},
	},
	"odi-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odi-port",
			Port:  3187,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odi-port",
			Port:  3187,
		},
	},
	"odmr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odmr",
			Port:  366,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odmr",
			Port:  366,
		},
	},
	"odn-castraq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odn-castraq",
			Port:  2498,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odn-castraq",
			Port:  2498,
		},
	},
	"odnsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odnsp",
			Port:  9966,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odnsp",
			Port:  9966,
		},
	},
	"odsi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "odsi",
			Port:  1308,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "odsi",
			Port:  1308,
		},
	},
	"oem-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oem-agent",
			Port:  3872,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oem-agent",
			Port:  3872,
		},
	},
	"oemcacao-jmxmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oemcacao-jmxmp",
			Port:  11172,
		},
	},
	"oemcacao-rmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oemcacao-rmi",
			Port:  11174,
		},
	},
	"oemcacao-websvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oemcacao-websvc",
			Port:  11175,
		},
	},
	"office-tools": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "office-tools",
			Port:  7789,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "office-tools",
			Port:  7789,
		},
	},
	"officelink2000": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "officelink2000",
			Port:  3320,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "officelink2000",
			Port:  3320,
		},
	},
	"ofsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ofsd",
			Port:  2322,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ofsd",
			Port:  2322,
		},
	},
	"ogs-client": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ogs-client",
			Port:  9007,
		},
	},
	"ogs-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ogs-server",
			Port:  9008,
		},
	},
	"ohimsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ohimsrv",
			Port:  506,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ohimsrv",
			Port:  506,
		},
	},
	"ohmtrigger": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ohmtrigger",
			Port:  4732,
		},
	},
	"ohsc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ohsc",
			Port:  18186,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ohsc",
			Port:  18186,
		},
	},
	"oi-2000": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oi-2000",
			Port:  2364,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oi-2000",
			Port:  2364,
		},
	},
	"oidocsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oidocsvc",
			Port:  4142,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oidocsvc",
			Port:  4142,
		},
	},
	"oidsr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oidsr",
			Port:  4143,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oidsr",
			Port:  4143,
		},
	},
	"oirtgsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oirtgsvc",
			Port:  4141,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oirtgsvc",
			Port:  4141,
		},
	},
	"olhost": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "olhost",
			Port:  2661,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "olhost",
			Port:  2661,
		},
	},
	"olsr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "olsr",
			Port:  698,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "olsr",
			Port:  698,
		},
	},
	"olsv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "olsv",
			Port:  1160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "olsv",
			Port:  1160,
		},
	},
	"oma-dcdocbs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oma-dcdocbs",
			Port:  7278,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oma-dcdocbs",
			Port:  7278,
		},
	},
	"oma-ilp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oma-ilp",
			Port:  7276,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oma-ilp",
			Port:  7276,
		},
	},
	"oma-ilp-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oma-ilp-s",
			Port:  7277,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oma-ilp-s",
			Port:  7277,
		},
	},
	"oma-mlp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oma-mlp",
			Port:  9210,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oma-mlp",
			Port:  9210,
		},
	},
	"oma-mlp-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oma-mlp-s",
			Port:  9211,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oma-mlp-s",
			Port:  9211,
		},
	},
	"oma-rlp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oma-rlp",
			Port:  7273,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oma-rlp",
			Port:  7273,
		},
	},
	"oma-rlp-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oma-rlp-s",
			Port:  7274,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oma-rlp-s",
			Port:  7274,
		},
	},
	"oma-ulp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oma-ulp",
			Port:  7275,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oma-ulp",
			Port:  7275,
		},
	},
	"omabcastltkm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omabcastltkm",
			Port:  4359,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omabcastltkm",
			Port:  4359,
		},
	},
	"omasgport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omasgport",
			Port:  4090,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omasgport",
			Port:  4090,
		},
	},
	"omginitialrefs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omginitialrefs",
			Port:  900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omginitialrefs",
			Port:  900,
		},
	},
	"omhs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omhs",
			Port:  5723,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omhs",
			Port:  5723,
		},
	},
	"omirr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omirr",
			Port:  808,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omirr",
			Port:  808,
		},
	},
	"omnilink-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omnilink-port",
			Port:  3904,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omnilink-port",
			Port:  3904,
		},
	},
	"omnisky": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omnisky",
			Port:  2056,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omnisky",
			Port:  2056,
		},
	},
	"omnivision": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omnivision",
			Port:  1135,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omnivision",
			Port:  1135,
		},
	},
	"omnivisionesx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omnivisionesx",
			Port:  4395,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omnivisionesx",
			Port:  4395,
		},
	},
	"oms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oms",
			Port:  4662,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oms",
			Port:  4662,
		},
	},
	"oms-nonsecure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oms-nonsecure",
			Port:  5102,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oms-nonsecure",
			Port:  5102,
		},
	},
	"omscontact": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omscontact",
			Port:  4161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omscontact",
			Port:  4161,
		},
	},
	"omsdk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omsdk",
			Port:  5724,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omsdk",
			Port:  5724,
		},
	},
	"omserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omserv",
			Port:  764,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omserv",
			Port:  764,
		},
	},
	"omstopology": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omstopology",
			Port:  4162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "omstopology",
			Port:  4162,
		},
	},
	"omviagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omviagent",
			Port:  4429,
		},
	},
	"omviserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "omviserver",
			Port:  4428,
		},
	},
	"onbase-dds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "onbase-dds",
			Port:  2185,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "onbase-dds",
			Port:  2185,
		},
	},
	"onehome-help": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "onehome-help",
			Port:  2199,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "onehome-help",
			Port:  2199,
		},
	},
	"onehome-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "onehome-remote",
			Port:  2198,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "onehome-remote",
			Port:  2198,
		},
	},
	"onesaf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "onesaf",
			Port:  3244,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "onesaf",
			Port:  3244,
		},
	},
	"onmux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "onmux",
			Port:  417,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "onmux",
			Port:  417,
		},
	},
	"onpsocket": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "onpsocket",
			Port:  5014,
		},
	},
	"onscreen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "onscreen",
			Port:  5080,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "onscreen",
			Port:  5080,
		},
	},
	"ontime": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ontime",
			Port:  1622,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ontime",
			Port:  1622,
		},
	},
	"ontobroker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ontobroker",
			Port:  2267,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ontobroker",
			Port:  2267,
		},
	},
	"oob-ws-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oob-ws-http",
			Port:  623,
		},
	},
	"oob-ws-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oob-ws-https",
			Port:  664,
		},
	},
	"op-probe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "op-probe",
			Port:  7030,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "op-probe",
			Port:  7030,
		},
	},
	"opalis-rbt-ipc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opalis-rbt-ipc",
			Port:  5314,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opalis-rbt-ipc",
			Port:  5314,
		},
	},
	"opalis-rdv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opalis-rdv",
			Port:  536,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opalis-rdv",
			Port:  536,
		},
	},
	"opalis-robot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opalis-robot",
			Port:  314,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opalis-robot",
			Port:  314,
		},
	},
	"opc-job-start": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opc-job-start",
			Port:  423,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opc-job-start",
			Port:  423,
		},
	},
	"opc-job-track": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opc-job-track",
			Port:  424,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opc-job-track",
			Port:  424,
		},
	},
	"opcon-xps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opcon-xps",
			Port:  3100,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opcon-xps",
			Port:  3100,
		},
	},
	"opcua-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opcua-tcp",
			Port:  4840,
		},
	},
	"opcua-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opcua-tls",
			Port:  4843,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opcua-tls",
			Port:  4843,
		},
	},
	"opcua-udp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "opcua-udp",
			Port:  4840,
		},
	},
	"opencm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opencm",
			Port:  3434,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opencm",
			Port:  3434,
		},
	},
	"opencore": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opencore",
			Port:  4089,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opencore",
			Port:  4089,
		},
	},
	"opendeploy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opendeploy",
			Port:  20014,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opendeploy",
			Port:  20014,
		},
	},
	"openhpid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openhpid",
			Port:  4743,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openhpid",
			Port:  4743,
		},
	},
	"openmail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openmail",
			Port:  5729,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openmail",
			Port:  5729,
		},
	},
	"openmailg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openmailg",
			Port:  5755,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openmailg",
			Port:  5755,
		},
	},
	"openmailns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openmailns",
			Port:  5766,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openmailns",
			Port:  5766,
		},
	},
	"openmailpxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openmailpxy",
			Port:  5768,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openmailpxy",
			Port:  5768,
		},
	},
	"openmath": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openmath",
			Port:  1473,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openmath",
			Port:  1473,
		},
	},
	"opennl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opennl",
			Port:  1258,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opennl",
			Port:  1258,
		},
	},
	"opennl-voice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opennl-voice",
			Port:  1259,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opennl-voice",
			Port:  1259,
		},
	},
	"openport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openport",
			Port:  260,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openport",
			Port:  260,
		},
	},
	"openqueue": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openqueue",
			Port:  8764,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openqueue",
			Port:  8764,
		},
	},
	"openremote-ctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openremote-ctrl",
			Port:  8688,
		},
	},
	"openstack-id": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openstack-id",
			Port:  35357,
		},
	},
	"opentable": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opentable",
			Port:  2368,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opentable",
			Port:  2368,
		},
	},
	"opentrac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opentrac",
			Port:  3855,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opentrac",
			Port:  3855,
		},
	},
	"openvms-sysipc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openvms-sysipc",
			Port:  557,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openvms-sysipc",
			Port:  557,
		},
	},
	"openvpn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openvpn",
			Port:  1194,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openvpn",
			Port:  1194,
		},
	},
	"openwebnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "openwebnet",
			Port:  20005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "openwebnet",
			Port:  20005,
		},
	},
	"opequus-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opequus-server",
			Port:  2400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opequus-server",
			Port:  2400,
		},
	},
	"opi-sock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opi-sock",
			Port:  7429,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opi-sock",
			Port:  7429,
		},
	},
	"opnet-smp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opnet-smp",
			Port:  3433,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opnet-smp",
			Port:  3433,
		},
	},
	"opsec-cvp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsec-cvp",
			Port:  18181,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsec-cvp",
			Port:  18181,
		},
	},
	"opsec-ela": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsec-ela",
			Port:  18187,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsec-ela",
			Port:  18187,
		},
	},
	"opsec-lea": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsec-lea",
			Port:  18184,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsec-lea",
			Port:  18184,
		},
	},
	"opsec-omi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsec-omi",
			Port:  18185,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsec-omi",
			Port:  18185,
		},
	},
	"opsec-sam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsec-sam",
			Port:  18183,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsec-sam",
			Port:  18183,
		},
	},
	"opsec-uaa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsec-uaa",
			Port:  19191,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsec-uaa",
			Port:  19191,
		},
	},
	"opsec-ufp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsec-ufp",
			Port:  18182,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsec-ufp",
			Port:  18182,
		},
	},
	"opsession-clnt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsession-clnt",
			Port:  3303,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsession-clnt",
			Port:  3303,
		},
	},
	"opsession-prxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsession-prxy",
			Port:  3307,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsession-prxy",
			Port:  3307,
		},
	},
	"opsession-srvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsession-srvr",
			Port:  3304,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsession-srvr",
			Port:  3304,
		},
	},
	"opsmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsmgr",
			Port:  1270,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsmgr",
			Port:  1270,
		},
	},
	"opsview-envoy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opsview-envoy",
			Port:  4125,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opsview-envoy",
			Port:  4125,
		},
	},
	"opswagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opswagent",
			Port:  3976,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opswagent",
			Port:  3976,
		},
	},
	"opswmanager": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opswmanager",
			Port:  3977,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opswmanager",
			Port:  3977,
		},
	},
	"optech-port1-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optech-port1-lm",
			Port:  2237,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optech-port1-lm",
			Port:  2237,
		},
	},
	"optika-emedia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optika-emedia",
			Port:  1829,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optika-emedia",
			Port:  1829,
		},
	},
	"optilogic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optilogic",
			Port:  2435,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optilogic",
			Port:  2435,
		},
	},
	"optima-vnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optima-vnet",
			Port:  1051,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optima-vnet",
			Port:  1051,
		},
	},
	"optiwave-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optiwave-lm",
			Port:  2524,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optiwave-lm",
			Port:  2524,
		},
	},
	"optocontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optocontrol",
			Port:  22001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optocontrol",
			Port:  22001,
		},
	},
	"optohost002": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optohost002",
			Port:  22002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optohost002",
			Port:  22002,
		},
	},
	"optohost003": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optohost003",
			Port:  22003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optohost003",
			Port:  22003,
		},
	},
	"optohost004": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optohost004",
			Port:  22004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optohost004",
			Port:  22004,
		},
	},
	"optohost005": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "optohost005",
			Port:  22005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "optohost005",
			Port:  22005,
		},
	},
	"opus-services": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "opus-services",
			Port:  3718,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "opus-services",
			Port:  3718,
		},
	},
	"ora-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ora-lm",
			Port:  1446,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ora-lm",
			Port:  1446,
		},
	},
	"ora-oap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ora-oap",
			Port:  5575,
		},
	},
	"oracle": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "oracle",
			Port:  2005,
		},
	},
	"oracle-em1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oracle-em1",
			Port:  1748,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oracle-em1",
			Port:  1748,
		},
	},
	"oracle-em2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oracle-em2",
			Port:  1754,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oracle-em2",
			Port:  1754,
		},
	},
	"oracle-oms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oracle-oms",
			Port:  1159,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oracle-oms",
			Port:  1159,
		},
	},
	"oracle-vp1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oracle-vp1",
			Port:  1809,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oracle-vp1",
			Port:  1809,
		},
	},
	"oracle-vp2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oracle-vp2",
			Port:  1808,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oracle-vp2",
			Port:  1808,
		},
	},
	"oracleas-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oracleas-https",
			Port:  7443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oracleas-https",
			Port:  7443,
		},
	},
	"oraclenames": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oraclenames",
			Port:  1575,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oraclenames",
			Port:  1575,
		},
	},
	"oraclenet8cman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oraclenet8cman",
			Port:  1630,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oraclenet8cman",
			Port:  1630,
		},
	},
	"orbiter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orbiter",
			Port:  2398,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orbiter",
			Port:  2398,
		},
	},
	"orbix-cfg-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orbix-cfg-ssl",
			Port:  3078,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orbix-cfg-ssl",
			Port:  3078,
		},
	},
	"orbix-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orbix-config",
			Port:  3076,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orbix-config",
			Port:  3076,
		},
	},
	"orbix-loc-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orbix-loc-ssl",
			Port:  3077,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orbix-loc-ssl",
			Port:  3077,
		},
	},
	"orbix-locator": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orbix-locator",
			Port:  3075,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orbix-locator",
			Port:  3075,
		},
	},
	"orbixd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orbixd",
			Port:  1570,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orbixd",
			Port:  1570,
		},
	},
	"orbplus-iiop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orbplus-iiop",
			Port:  1597,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orbplus-iiop",
			Port:  1597,
		},
	},
	"ordinox-dbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ordinox-dbase",
			Port:  3355,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ordinox-dbase",
			Port:  3355,
		},
	},
	"ordinox-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ordinox-server",
			Port:  3274,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ordinox-server",
			Port:  3274,
		},
	},
	"origo-native": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "origo-native",
			Port:  3001,
		},
	},
	"origo-sync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "origo-sync",
			Port:  11103,
		},
	},
	"orion": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orion",
			Port:  2407,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orion",
			Port:  2407,
		},
	},
	"orion-rmi-reg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "orion-rmi-reg",
			Port:  2413,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "orion-rmi-reg",
			Port:  2413,
		},
	},
	"ortec-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ortec-disc",
			Port:  40853,
		},
	},
	"os-licman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "os-licman",
			Port:  1384,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "os-licman",
			Port:  1384,
		},
	},
	"osaut": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "osaut",
			Port:  6679,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "osaut",
			Port:  6679,
		},
	},
	"osb-sd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "osb-sd",
			Port:  400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "osb-sd",
			Port:  400,
		},
	},
	"osdcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "osdcp",
			Port:  3432,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "osdcp",
			Port:  3432,
		},
	},
	"osm-appsrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "osm-appsrvr",
			Port:  9990,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "osm-appsrvr",
			Port:  9990,
		},
	},
	"osm-oev": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "osm-oev",
			Port:  9991,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "osm-oev",
			Port:  9991,
		},
	},
	"osmosis-aeea": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "osmosis-aeea",
			Port:  3034,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "osmosis-aeea",
			Port:  3034,
		},
	},
	"osp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "osp",
			Port:  5045,
		},
	},
	"ospf-lite": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ospf-lite",
			Port:  8899,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ospf-lite",
			Port:  8899,
		},
	},
	"osu-nms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "osu-nms",
			Port:  192,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "osu-nms",
			Port:  192,
		},
	},
	"otlp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "otlp",
			Port:  6951,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "otlp",
			Port:  6951,
		},
	},
	"otmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "otmp",
			Port:  29167,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "otmp",
			Port:  29167,
		},
	},
	"otp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "otp",
			Port:  9390,
		},
	},
	"otpatch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "otpatch",
			Port:  2936,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "otpatch",
			Port:  2936,
		},
	},
	"ott": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ott",
			Port:  2428,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ott",
			Port:  2428,
		},
	},
	"ottp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ottp",
			Port:  2951,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ottp",
			Port:  2951,
		},
	},
	"otv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "otv",
			Port:  8472,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "otv",
			Port:  8472,
		},
	},
	"outlaws": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "outlaws",
			Port:  5310,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "outlaws",
			Port:  5310,
		},
	},
	"ov-nnm-websrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ov-nnm-websrv",
			Port:  3443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ov-nnm-websrv",
			Port:  3443,
		},
	},
	"ovalarmsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovalarmsrv",
			Port:  2953,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovalarmsrv",
			Port:  2953,
		},
	},
	"ovalarmsrv-cmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovalarmsrv-cmd",
			Port:  2954,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovalarmsrv-cmd",
			Port:  2954,
		},
	},
	"ovbus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovbus",
			Port:  7501,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovbus",
			Port:  7501,
		},
	},
	"oveadmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "oveadmgr",
			Port:  7427,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "oveadmgr",
			Port:  7427,
		},
	},
	"ovhpas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovhpas",
			Port:  7510,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovhpas",
			Port:  7510,
		},
	},
	"ovladmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovladmgr",
			Port:  7428,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovladmgr",
			Port:  7428,
		},
	},
	"ovobs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovobs",
			Port:  30999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovobs",
			Port:  30999,
		},
	},
	"ovrimosdbman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovrimosdbman",
			Port:  2956,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovrimosdbman",
			Port:  2956,
		},
	},
	"ovsam-d-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovsam-d-agent",
			Port:  3870,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovsam-d-agent",
			Port:  3870,
		},
	},
	"ovsam-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovsam-mgmt",
			Port:  3869,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovsam-mgmt",
			Port:  3869,
		},
	},
	"ovsessionmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovsessionmgr",
			Port:  2389,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovsessionmgr",
			Port:  2389,
		},
	},
	"ovtopmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovtopmd",
			Port:  2532,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovtopmd",
			Port:  2532,
		},
	},
	"ovwdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ovwdb",
			Port:  2447,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ovwdb",
			Port:  2447,
		},
	},
	"owamp-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "owamp-control",
			Port:  861,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "owamp-control",
			Port:  861,
		},
	},
	"owserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "owserver",
			Port:  4304,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "owserver",
			Port:  4304,
		},
	},
	"p-net-local": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "p-net-local",
			Port:  34378,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "p-net-local",
			Port:  34378,
		},
	},
	"p-net-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "p-net-remote",
			Port:  34379,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "p-net-remote",
			Port:  34379,
		},
	},
	"p25cai": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "p25cai",
			Port:  6082,
		},
	},
	"p2pcommunity": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "p2pcommunity",
			Port:  3955,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "p2pcommunity",
			Port:  3955,
		},
	},
	"p2pgroup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "p2pgroup",
			Port:  3587,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "p2pgroup",
			Port:  3587,
		},
	},
	"p2pq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "p2pq",
			Port:  1981,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "p2pq",
			Port:  1981,
		},
	},
	"p4p-portal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "p4p-portal",
			Port:  6671,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "p4p-portal",
			Port:  6671,
		},
	},
	"p6ssmc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "p6ssmc",
			Port:  4311,
		},
	},
	"pacerforum": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pacerforum",
			Port:  1480,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pacerforum",
			Port:  1480,
		},
	},
	"pacmand": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pacmand",
			Port:  1307,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pacmand",
			Port:  1307,
		},
	},
	"pacom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pacom",
			Port:  3435,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pacom",
			Port:  3435,
		},
	},
	"padl2sim": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "padl2sim",
			Port:  5236,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "padl2sim",
			Port:  5236,
		},
	},
	"pads": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pads",
			Port:  7237,
		},
	},
	"pafec-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pafec-lm",
			Port:  7511,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pafec-lm",
			Port:  7511,
		},
	},
	"paging-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "paging-port",
			Port:  3771,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "paging-port",
			Port:  3771,
		},
	},
	"pago-services1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pago-services1",
			Port:  30001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pago-services1",
			Port:  30001,
		},
	},
	"pago-services2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pago-services2",
			Port:  30002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pago-services2",
			Port:  30002,
		},
	},
	"palace-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "palace-1",
			Port:  9992,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "palace-1",
			Port:  9992,
		},
	},
	"palace-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "palace-2",
			Port:  9993,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "palace-2",
			Port:  9993,
		},
	},
	"palace-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "palace-3",
			Port:  9994,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "palace-3",
			Port:  9994,
		},
	},
	"palace-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "palace-4",
			Port:  9995,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "palace-4",
			Port:  9995,
		},
	},
	"palace-5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "palace-5",
			Port:  9996,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "palace-5",
			Port:  9996,
		},
	},
	"palace-6": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "palace-6",
			Port:  9997,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "palace-6",
			Port:  9997,
		},
	},
	"palcom-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "palcom-disc",
			Port:  6657,
		},
	},
	"pammratc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pammratc",
			Port:  1632,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pammratc",
			Port:  1632,
		},
	},
	"pammrpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pammrpc",
			Port:  1633,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pammrpc",
			Port:  1633,
		},
	},
	"pana": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "pana",
			Port:  716,
		},
	},
	"panagolin-ident": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "panagolin-ident",
			Port:  9021,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "panagolin-ident",
			Port:  9021,
		},
	},
	"panasas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "panasas",
			Port:  3095,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "panasas",
			Port:  3095,
		},
	},
	"pando-pub": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pando-pub",
			Port:  7680,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pando-pub",
			Port:  7680,
		},
	},
	"pando-sec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pando-sec",
			Port:  8276,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pando-sec",
			Port:  8276,
		},
	},
	"pangolin-laser": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pangolin-laser",
			Port:  3348,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pangolin-laser",
			Port:  3348,
		},
	},
	"paradym-31port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "paradym-31port",
			Port:  1864,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "paradym-31port",
			Port:  1864,
		},
	},
	"paragent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "paragent",
			Port:  9022,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "paragent",
			Port:  9022,
		},
	},
	"parallel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "parallel",
			Port:  4989,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "parallel",
			Port:  4989,
		},
	},
	"park-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "park-agent",
			Port:  5431,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "park-agent",
			Port:  5431,
		},
	},
	"parliant": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "parliant",
			Port:  4681,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "parliant",
			Port:  4681,
		},
	},
	"parsec-game": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "parsec-game",
			Port:  6582,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "parsec-game",
			Port:  6582,
		},
	},
	"parsec-master": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "parsec-master",
			Port:  6580,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "parsec-master",
			Port:  6580,
		},
	},
	"parsec-peer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "parsec-peer",
			Port:  6581,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "parsec-peer",
			Port:  6581,
		},
	},
	"partimage": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "partimage",
			Port:  4025,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "partimage",
			Port:  4025,
		},
	},
	"passgo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "passgo",
			Port:  511,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "passgo",
			Port:  511,
		},
	},
	"passgo-tivoli": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "passgo-tivoli",
			Port:  627,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "passgo-tivoli",
			Port:  627,
		},
	},
	"passwd_server": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "passwd_server",
			Port:  752,
		},
	},
	"password-chg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "password-chg",
			Port:  586,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "password-chg",
			Port:  586,
		},
	},
	"passwrd-policy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "passwrd-policy",
			Port:  1333,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "passwrd-policy",
			Port:  1333,
		},
	},
	"patrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "patrol",
			Port:  8160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "patrol",
			Port:  8160,
		},
	},
	"patrol-coll": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "patrol-coll",
			Port:  6162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "patrol-coll",
			Port:  6162,
		},
	},
	"patrol-ism": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "patrol-ism",
			Port:  6161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "patrol-ism",
			Port:  6161,
		},
	},
	"patrol-mq-gm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "patrol-mq-gm",
			Port:  2664,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "patrol-mq-gm",
			Port:  2664,
		},
	},
	"patrol-mq-nm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "patrol-mq-nm",
			Port:  2665,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "patrol-mq-nm",
			Port:  2665,
		},
	},
	"patrol-snmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "patrol-snmp",
			Port:  8161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "patrol-snmp",
			Port:  8161,
		},
	},
	"patrolview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "patrolview",
			Port:  4097,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "patrolview",
			Port:  4097,
		},
	},
	"pawserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pawserv",
			Port:  345,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pawserv",
			Port:  345,
		},
	},
	"pay-per-view": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pay-per-view",
			Port:  1564,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pay-per-view",
			Port:  1564,
		},
	},
	"paycash-online": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "paycash-online",
			Port:  8128,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "paycash-online",
			Port:  8128,
		},
	},
	"paycash-wbp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "paycash-wbp",
			Port:  8129,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "paycash-wbp",
			Port:  8129,
		},
	},
	"payrouter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "payrouter",
			Port:  1246,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "payrouter",
			Port:  1246,
		},
	},
	"pc-mta-addrmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pc-mta-addrmap",
			Port:  2246,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pc-mta-addrmap",
			Port:  2246,
		},
	},
	"pc-telecommute": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pc-telecommute",
			Port:  2299,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pc-telecommute",
			Port:  2299,
		},
	},
	"pcanywheredata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcanywheredata",
			Port:  5631,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcanywheredata",
			Port:  5631,
		},
	},
	"pcanywherestat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcanywherestat",
			Port:  5632,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcanywherestat",
			Port:  5632,
		},
	},
	"pcc-image-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcc-image-port",
			Port:  3892,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcc-image-port",
			Port:  3892,
		},
	},
	"pcc-mfp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcc-mfp",
			Port:  2256,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcc-mfp",
			Port:  2256,
		},
	},
	"pcep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcep",
			Port:  4189,
		},
	},
	"pcia-rxp-b": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcia-rxp-b",
			Port:  1332,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcia-rxp-b",
			Port:  1332,
		},
	},
	"pciarray": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pciarray",
			Port:  1552,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pciarray",
			Port:  1552,
		},
	},
	"pcihreq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcihreq",
			Port:  3085,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcihreq",
			Port:  3085,
		},
	},
	"pcle-infex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcle-infex",
			Port:  3189,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcle-infex",
			Port:  3189,
		},
	},
	"pclemultimedia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pclemultimedia",
			Port:  2558,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pclemultimedia",
			Port:  2558,
		},
	},
	"pcmail-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcmail-srv",
			Port:  158,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcmail-srv",
			Port:  158,
		},
	},
	"pcmk-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcmk-remote",
			Port:  3121,
		},
	},
	"pcoip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcoip",
			Port:  4172,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcoip",
			Port:  4172,
		},
	},
	"pconnectmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pconnectmgr",
			Port:  1562,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pconnectmgr",
			Port:  1562,
		},
	},
	"pcp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "pcp",
			Port:  5350,
		},
	},
	"pcp-multicast": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcp-multicast",
			Port:  5350,
		},
	},
	"pcptcpservice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcptcpservice",
			Port:  4182,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcptcpservice",
			Port:  4182,
		},
	},
	"pcs-pcw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcs-pcw",
			Port:  2566,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcs-pcw",
			Port:  2566,
		},
	},
	"pcs-sf-ui-man": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcs-sf-ui-man",
			Port:  6655,
		},
	},
	"pcsync-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcsync-http",
			Port:  8444,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcsync-http",
			Port:  8444,
		},
	},
	"pcsync-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcsync-https",
			Port:  8443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcsync-https",
			Port:  8443,
		},
	},
	"pctrader": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pctrader",
			Port:  3048,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pctrader",
			Port:  3048,
		},
	},
	"pcttunnell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pcttunnell",
			Port:  2274,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pcttunnell",
			Port:  2274,
		},
	},
	"pd-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pd-admin",
			Port:  9597,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pd-admin",
			Port:  9597,
		},
	},
	"pda-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pda-data",
			Port:  3253,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pda-data",
			Port:  3253,
		},
	},
	"pda-gate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pda-gate",
			Port:  4012,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pda-gate",
			Port:  4012,
		},
	},
	"pda-sys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pda-sys",
			Port:  3254,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pda-sys",
			Port:  3254,
		},
	},
	"pdap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdap",
			Port:  344,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdap",
			Port:  344,
		},
	},
	"pdap-np": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdap-np",
			Port:  1526,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdap-np",
			Port:  1526,
		},
	},
	"pdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdb",
			Port:  3033,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdb",
			Port:  3033,
		},
	},
	"pdefmns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdefmns",
			Port:  16311,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdefmns",
			Port:  16311,
		},
	},
	"pdnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdnet",
			Port:  2843,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdnet",
			Port:  2843,
		},
	},
	"pdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdp",
			Port:  1675,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdp",
			Port:  1675,
		},
	},
	"pdps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdps",
			Port:  1314,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdps",
			Port:  1314,
		},
	},
	"pdrncs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdrncs",
			Port:  3299,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdrncs",
			Port:  3299,
		},
	},
	"pds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pds",
			Port:  9595,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pds",
			Port:  9595,
		},
	},
	"pdtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pdtp",
			Port:  6086,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pdtp",
			Port:  6086,
		},
	},
	"pduncs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pduncs",
			Port:  16310,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pduncs",
			Port:  16310,
		},
	},
	"pe-mike": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pe-mike",
			Port:  1305,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pe-mike",
			Port:  1305,
		},
	},
	"pearldoc-xact": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pearldoc-xact",
			Port:  1980,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pearldoc-xact",
			Port:  1980,
		},
	},
	"peerbook-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "peerbook-port",
			Port:  3135,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "peerbook-port",
			Port:  3135,
		},
	},
	"peerwire": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "peerwire",
			Port:  9104,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "peerwire",
			Port:  9104,
		},
	},
	"pegasus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pegasus",
			Port:  9278,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pegasus",
			Port:  9278,
		},
	},
	"pegasus-ctl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pegasus-ctl",
			Port:  9279,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pegasus-ctl",
			Port:  9279,
		},
	},
	"pegboard": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pegboard",
			Port:  1357,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pegboard",
			Port:  1357,
		},
	},
	"pehelp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pehelp",
			Port:  2307,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pehelp",
			Port:  2307,
		},
	},
	"pentbox-sim": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pentbox-sim",
			Port:  6817,
		},
	},
	"peocoll": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "peocoll",
			Port:  9631,
		},
	},
	"peoctlr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "peoctlr",
			Port:  9630,
		},
	},
	"peport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "peport",
			Port:  1449,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "peport",
			Port:  1449,
		},
	},
	"perf-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "perf-port",
			Port:  1995,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "perf-port",
			Port:  1995,
		},
	},
	"perfd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "perfd",
			Port:  5227,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "perfd",
			Port:  5227,
		},
	},
	"perimlan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "perimlan",
			Port:  4075,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "perimlan",
			Port:  4075,
		},
	},
	"periscope": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "periscope",
			Port:  1230,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "periscope",
			Port:  1230,
		},
	},
	"permabit-cs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "permabit-cs",
			Port:  5312,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "permabit-cs",
			Port:  5312,
		},
	},
	"perrla": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "perrla",
			Port:  4313,
		},
	},
	"persona": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "persona",
			Port:  1916,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "persona",
			Port:  1916,
		},
	},
	"personal-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "personal-agent",
			Port:  5555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "personal-agent",
			Port:  5555,
		},
	},
	"personal-link": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "personal-link",
			Port:  281,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "personal-link",
			Port:  281,
		},
	},
	"personalos-001": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "personalos-001",
			Port:  3557,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "personalos-001",
			Port:  3557,
		},
	},
	"personnel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "personnel",
			Port:  3109,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "personnel",
			Port:  3109,
		},
	},
	"pftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pftp",
			Port:  662,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pftp",
			Port:  662,
		},
	},
	"pfu-prcallback": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pfu-prcallback",
			Port:  3208,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pfu-prcallback",
			Port:  3208,
		},
	},
	"pgbouncer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pgbouncer",
			Port:  6432,
		},
	},
	"pgpkeyserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pgpkeyserver",
			Port:  11371,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pgpkeyserver",
			Port:  11371,
		},
	},
	"pgps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pgps",
			Port:  9280,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pgps",
			Port:  9280,
		},
	},
	"ph": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ph",
			Port:  481,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ph",
			Port:  481,
		},
	},
	"pharmasoft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pharmasoft",
			Port:  1779,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pharmasoft",
			Port:  1779,
		},
	},
	"pharos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pharos",
			Port:  4443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pharos",
			Port:  4443,
		},
	},
	"philips-vc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "philips-vc",
			Port:  583,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "philips-vc",
			Port:  583,
		},
	},
	"phoenix-rpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "phoenix-rpc",
			Port:  3347,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "phoenix-rpc",
			Port:  3347,
		},
	},
	"phonebook": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "phonebook",
			Port:  767,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "phonebook",
			Port:  767,
		},
	},
	"phonex-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "phonex-port",
			Port:  3177,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "phonex-port",
			Port:  3177,
		},
	},
	"photuris": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "photuris",
			Port:  468,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "photuris",
			Port:  468,
		},
	},
	"phrelay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "phrelay",
			Port:  4868,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "phrelay",
			Port:  4868,
		},
	},
	"phrelaydbg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "phrelaydbg",
			Port:  4869,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "phrelaydbg",
			Port:  4869,
		},
	},
	"piccolo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "piccolo",
			Port:  2787,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "piccolo",
			Port:  2787,
		},
	},
	"pichat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pichat",
			Port:  9009,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pichat",
			Port:  9009,
		},
	},
	"picknfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "picknfs",
			Port:  1598,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "picknfs",
			Port:  1598,
		},
	},
	"picodbc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "picodbc",
			Port:  1603,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "picodbc",
			Port:  1603,
		},
	},
	"pictrography": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pictrography",
			Port:  1280,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pictrography",
			Port:  1280,
		},
	},
	"pim-port": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "pim-port",
			Port:  8471,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "pim-port",
			Port:  8471,
		},
	},
	"pim-rp-disc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pim-rp-disc",
			Port:  496,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pim-rp-disc",
			Port:  496,
		},
	},
	"ping-pong": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ping-pong",
			Port:  3010,
		},
	},
	"pinghgl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pinghgl",
			Port:  4306,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pinghgl",
			Port:  4306,
		},
	},
	"pip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pip",
			Port:  321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pip",
			Port:  321,
		},
	},
	"pipe_server": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "pipe_server",
			Port:  2010,
		},
	},
	"pipes": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pipes",
			Port:  1465,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pipes",
			Port:  1465,
		},
	},
	"piranha1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "piranha1",
			Port:  4600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "piranha1",
			Port:  4600,
		},
	},
	"piranha2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "piranha2",
			Port:  4601,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "piranha2",
			Port:  4601,
		},
	},
	"pirp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pirp",
			Port:  553,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pirp",
			Port:  553,
		},
	},
	"pit-vpn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pit-vpn",
			Port:  2865,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pit-vpn",
			Port:  2865,
		},
	},
	"pjlink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pjlink",
			Port:  4352,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pjlink",
			Port:  4352,
		},
	},
	"pk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pk",
			Port:  5272,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pk",
			Port:  5272,
		},
	},
	"pk-electronics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pk-electronics",
			Port:  2634,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pk-electronics",
			Port:  2634,
		},
	},
	"pkagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pkagent",
			Port:  3118,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pkagent",
			Port:  3118,
		},
	},
	"pkix-3-ca-ra": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pkix-3-ca-ra",
			Port:  829,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pkix-3-ca-ra",
			Port:  829,
		},
	},
	"pkix-cmc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pkix-cmc",
			Port:  5318,
		},
	},
	"pkix-timestamp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pkix-timestamp",
			Port:  318,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pkix-timestamp",
			Port:  318,
		},
	},
	"pkt-krb-ipsec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pkt-krb-ipsec",
			Port:  1293,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pkt-krb-ipsec",
			Port:  1293,
		},
	},
	"pktcable-cops": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pktcable-cops",
			Port:  2126,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pktcable-cops",
			Port:  2126,
		},
	},
	"pktcablemmcops": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pktcablemmcops",
			Port:  3918,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pktcablemmcops",
			Port:  3918,
		},
	},
	"plato": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "plato",
			Port:  3285,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "plato",
			Port:  3285,
		},
	},
	"plato-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "plato-lm",
			Port:  1819,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "plato-lm",
			Port:  1819,
		},
	},
	"playsta2-app": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "playsta2-app",
			Port:  4658,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "playsta2-app",
			Port:  4658,
		},
	},
	"playsta2-lob": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "playsta2-lob",
			Port:  4659,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "playsta2-lob",
			Port:  4659,
		},
	},
	"plbserve-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "plbserve-port",
			Port:  3933,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "plbserve-port",
			Port:  3933,
		},
	},
	"plcy-net-svcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "plcy-net-svcs",
			Port:  4351,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "plcy-net-svcs",
			Port:  4351,
		},
	},
	"plethora": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "plethora",
			Port:  3480,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "plethora",
			Port:  3480,
		},
	},
	"plgproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "plgproxy",
			Port:  2790,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "plgproxy",
			Port:  2790,
		},
	},
	"pluribus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pluribus",
			Port:  3469,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pluribus",
			Port:  3469,
		},
	},
	"plysrv-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "plysrv-http",
			Port:  6770,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "plysrv-http",
			Port:  6770,
		},
	},
	"plysrv-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "plysrv-https",
			Port:  6771,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "plysrv-https",
			Port:  6771,
		},
	},
	"pm-cmdsvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pm-cmdsvr",
			Port:  5112,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pm-cmdsvr",
			Port:  5112,
		},
	},
	"pmas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmas",
			Port:  4066,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmas",
			Port:  4066,
		},
	},
	"pmcd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmcd",
			Port:  44321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmcd",
			Port:  44321,
		},
	},
	"pmcdproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmcdproxy",
			Port:  44322,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmcdproxy",
			Port:  44322,
		},
	},
	"pmcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmcp",
			Port:  3821,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmcp",
			Port:  3821,
		},
	},
	"pmcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmcs",
			Port:  6355,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmcs",
			Port:  6355,
		},
	},
	"pmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmd",
			Port:  7431,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmd",
			Port:  7431,
		},
	},
	"pmdfmgt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmdfmgt",
			Port:  7633,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmdfmgt",
			Port:  7633,
		},
	},
	"pmdmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmdmgr",
			Port:  7426,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmdmgr",
			Port:  7426,
		},
	},
	"pmip6-cntl": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "pmip6-cntl",
			Port:  5436,
		},
	},
	"pmip6-data": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "pmip6-data",
			Port:  5437,
		},
	},
	"pmsm-webrctl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pmsm-webrctl",
			Port:  2972,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pmsm-webrctl",
			Port:  2972,
		},
	},
	"pn-requester": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pn-requester",
			Port:  2717,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pn-requester",
			Port:  2717,
		},
	},
	"pn-requester2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pn-requester2",
			Port:  2718,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pn-requester2",
			Port:  2718,
		},
	},
	"pnaconsult-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pnaconsult-lm",
			Port:  2937,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pnaconsult-lm",
			Port:  2937,
		},
	},
	"pnbs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pnbs",
			Port:  6124,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pnbs",
			Port:  6124,
		},
	},
	"pnbscada": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pnbscada",
			Port:  3875,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pnbscada",
			Port:  3875,
		},
	},
	"pnet-conn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pnet-conn",
			Port:  7797,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pnet-conn",
			Port:  7797,
		},
	},
	"pnet-enc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pnet-enc",
			Port:  7798,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pnet-enc",
			Port:  7798,
		},
	},
	"pnrp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pnrp-port",
			Port:  3540,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pnrp-port",
			Port:  3540,
		},
	},
	"pns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pns",
			Port:  2487,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pns",
			Port:  2487,
		},
	},
	"polestar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "polestar",
			Port:  1060,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "polestar",
			Port:  1060,
		},
	},
	"policyserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "policyserver",
			Port:  3055,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "policyserver",
			Port:  3055,
		},
	},
	"pop2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pop2",
			Port:  109,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pop2",
			Port:  109,
		},
	},
	"pop3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pop3",
			Port:  110,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pop3",
			Port:  110,
		},
	},
	"pop3s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pop3s",
			Port:  995,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pop3s",
			Port:  995,
		},
	},
	"poppassd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "poppassd",
			Port:  106,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "poppassd",
			Port:  106,
		},
	},
	"popup-reminders": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "popup-reminders",
			Port:  7787,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "popup-reminders",
			Port:  7787,
		},
	},
	"portgate-auth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "portgate-auth",
			Port:  3710,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "portgate-auth",
			Port:  3710,
		},
	},
	"postgres": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "postgres",
			Port:  5432,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "postgres",
			Port:  5432,
		},
	},
	"pov-ray": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pov-ray",
			Port:  494,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pov-ray",
			Port:  494,
		},
	},
	"powerburst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "powerburst",
			Port:  485,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "powerburst",
			Port:  485,
		},
	},
	"powerclientcsf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "powerclientcsf",
			Port:  2443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "powerclientcsf",
			Port:  2443,
		},
	},
	"powerexchange": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "powerexchange",
			Port:  2480,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "powerexchange",
			Port:  2480,
		},
	},
	"powergemplus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "powergemplus",
			Port:  2899,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "powergemplus",
			Port:  2899,
		},
	},
	"powerguardian": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "powerguardian",
			Port:  1777,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "powerguardian",
			Port:  1777,
		},
	},
	"poweronnud": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "poweronnud",
			Port:  3168,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "poweronnud",
			Port:  3168,
		},
	},
	"powerschool": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "powerschool",
			Port:  5071,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "powerschool",
			Port:  5071,
		},
	},
	"powwow-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "powwow-client",
			Port:  13223,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "powwow-client",
			Port:  13223,
		},
	},
	"powwow-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "powwow-server",
			Port:  13224,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "powwow-server",
			Port:  13224,
		},
	},
	"ppactivation": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ppactivation",
			Port:  5134,
		},
	},
	"ppcontrol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ppcontrol",
			Port:  2505,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ppcontrol",
			Port:  2505,
		},
	},
	"ppsms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ppsms",
			Port:  3967,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ppsms",
			Port:  3967,
		},
	},
	"ppsuitemsg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ppsuitemsg",
			Port:  5863,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ppsuitemsg",
			Port:  5863,
		},
	},
	"pptconference": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pptconference",
			Port:  1711,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pptconference",
			Port:  1711,
		},
	},
	"pptp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pptp",
			Port:  1723,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pptp",
			Port:  1723,
		},
	},
	"pq-lic-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pq-lic-mgmt",
			Port:  2687,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pq-lic-mgmt",
			Port:  2687,
		},
	},
	"pqsflows": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pqsflows",
			Port:  9640,
		},
	},
	"pqsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pqsp",
			Port:  28001,
		},
	},
	"prRegister": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prRegister",
			Port:  4457,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prRegister",
			Port:  4457,
		},
	},
	"pra_elmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pra_elmd",
			Port:  1587,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pra_elmd",
			Port:  1587,
		},
	},
	"prat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prat",
			Port:  1264,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prat",
			Port:  1264,
		},
	},
	"prchat-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prchat-server",
			Port:  4456,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prchat-server",
			Port:  4456,
		},
	},
	"prchat-user": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prchat-user",
			Port:  4455,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prchat-user",
			Port:  4455,
		},
	},
	"precise-comm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "precise-comm",
			Port:  5630,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "precise-comm",
			Port:  5630,
		},
	},
	"precise-i3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "precise-i3",
			Port:  3607,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "precise-i3",
			Port:  3607,
		},
	},
	"precise-sft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "precise-sft",
			Port:  2315,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "precise-sft",
			Port:  2315,
		},
	},
	"precise-vip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "precise-vip",
			Port:  2924,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "precise-vip",
			Port:  2924,
		},
	},
	"predatar-comms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "predatar-comms",
			Port:  1753,
		},
	},
	"prelude": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prelude",
			Port:  4690,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prelude",
			Port:  4690,
		},
	},
	"presence": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "presence",
			Port:  5298,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "presence",
			Port:  5298,
		},
	},
	"press": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "press",
			Port:  3582,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "press",
			Port:  3582,
		},
	},
	"prex-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prex-tcp",
			Port:  4487,
		},
	},
	"prgp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prgp",
			Port:  7747,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prgp",
			Port:  7747,
		},
	},
	"primaserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "primaserver",
			Port:  6105,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "primaserver",
			Port:  6105,
		},
	},
	"print-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "print-srv",
			Port:  170,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "print-srv",
			Port:  170,
		},
	},
	"printer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "printer",
			Port:  515,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "printer",
			Port:  515,
		},
	},
	"printer_agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "printer_agent",
			Port:  3396,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "printer_agent",
			Port:  3396,
		},
	},
	"printopia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "printopia",
			Port:  10631,
		},
	},
	"priority-e-com": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "priority-e-com",
			Port:  2618,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "priority-e-com",
			Port:  2618,
		},
	},
	"prism-deploy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prism-deploy",
			Port:  3133,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prism-deploy",
			Port:  3133,
		},
	},
	"prismiq-plugin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prismiq-plugin",
			Port:  3650,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prismiq-plugin",
			Port:  3650,
		},
	},
	"privateark": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "privateark",
			Port:  1858,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "privateark",
			Port:  1858,
		},
	},
	"privatechat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "privatechat",
			Port:  1735,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "privatechat",
			Port:  1735,
		},
	},
	"privatewire": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "privatewire",
			Port:  4449,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "privatewire",
			Port:  4449,
		},
	},
	"privilege": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "privilege",
			Port:  2588,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "privilege",
			Port:  2588,
		},
	},
	"privoxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "privoxy",
			Port:  8118,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "privoxy",
			Port:  8118,
		},
	},
	"prizma": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prizma",
			Port:  2039,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prizma",
			Port:  2039,
		},
	},
	"prm-nm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prm-nm",
			Port:  409,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prm-nm",
			Port:  409,
		},
	},
	"prm-nm-np": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prm-nm-np",
			Port:  1403,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prm-nm-np",
			Port:  1403,
		},
	},
	"prm-sm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prm-sm",
			Port:  408,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prm-sm",
			Port:  408,
		},
	},
	"prm-sm-np": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prm-sm-np",
			Port:  1402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prm-sm-np",
			Port:  1402,
		},
	},
	"prnrequest": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prnrequest",
			Port:  3910,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prnrequest",
			Port:  3910,
		},
	},
	"prnstatus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prnstatus",
			Port:  3911,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prnstatus",
			Port:  3911,
		},
	},
	"pro-ed": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pro-ed",
			Port:  8032,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pro-ed",
			Port:  8032,
		},
	},
	"proactivate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proactivate",
			Port:  24678,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proactivate",
			Port:  24678,
		},
	},
	"proactivesrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proactivesrvr",
			Port:  2722,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proactivesrvr",
			Port:  2722,
		},
	},
	"proaxess": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proaxess",
			Port:  3961,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proaxess",
			Port:  3961,
		},
	},
	"procos-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "procos-lm",
			Port:  3248,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "procos-lm",
			Port:  3248,
		},
	},
	"prodigy-intrnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prodigy-intrnet",
			Port:  1778,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prodigy-intrnet",
			Port:  1778,
		},
	},
	"productinfo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "productinfo",
			Port:  1283,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "productinfo",
			Port:  1283,
		},
	},
	"profile": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "profile",
			Port:  136,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "profile",
			Port:  136,
		},
	},
	"profilemac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "profilemac",
			Port:  4749,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "profilemac",
			Port:  4749,
		},
	},
	"profinet-cm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "profinet-cm",
			Port:  34964,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "profinet-cm",
			Port:  34964,
		},
	},
	"profinet-rt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "profinet-rt",
			Port:  34962,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "profinet-rt",
			Port:  34962,
		},
	},
	"profinet-rtm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "profinet-rtm",
			Port:  34963,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "profinet-rtm",
			Port:  34963,
		},
	},
	"progistics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "progistics",
			Port:  3973,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "progistics",
			Port:  3973,
		},
	},
	"programmar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "programmar",
			Port:  15999,
		},
	},
	"prolink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prolink",
			Port:  1678,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prolink",
			Port:  1678,
		},
	},
	"proofd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proofd",
			Port:  1093,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proofd",
			Port:  1093,
		},
	},
	"propel-msgsys": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "propel-msgsys",
			Port:  1268,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "propel-msgsys",
			Port:  1268,
		},
	},
	"proremote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proremote",
			Port:  8183,
		},
	},
	"proshare-mc-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proshare-mc-1",
			Port:  1673,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proshare-mc-1",
			Port:  1673,
		},
	},
	"proshare-mc-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proshare-mc-2",
			Port:  1674,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proshare-mc-2",
			Port:  1674,
		},
	},
	"proshare1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proshare1",
			Port:  1459,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proshare1",
			Port:  1459,
		},
	},
	"proshare2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proshare2",
			Port:  1460,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proshare2",
			Port:  1460,
		},
	},
	"proshareaudio": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proshareaudio",
			Port:  5713,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proshareaudio",
			Port:  5713,
		},
	},
	"prosharedata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prosharedata",
			Port:  5715,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prosharedata",
			Port:  5715,
		},
	},
	"prosharenotify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prosharenotify",
			Port:  5717,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prosharenotify",
			Port:  5717,
		},
	},
	"prosharerequest": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prosharerequest",
			Port:  5716,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prosharerequest",
			Port:  5716,
		},
	},
	"prosharevideo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prosharevideo",
			Port:  5714,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prosharevideo",
			Port:  5714,
		},
	},
	"prospero": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prospero",
			Port:  191,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prospero",
			Port:  191,
		},
	},
	"prospero-np": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prospero-np",
			Port:  1525,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prospero-np",
			Port:  1525,
		},
	},
	"proxim": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proxim",
			Port:  1732,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proxim",
			Port:  1732,
		},
	},
	"proxima-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proxima-lm",
			Port:  1445,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proxima-lm",
			Port:  1445,
		},
	},
	"proxy-gateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "proxy-gateway",
			Port:  2303,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "proxy-gateway",
			Port:  2303,
		},
	},
	"prp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prp",
			Port:  2091,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prp",
			Port:  2091,
		},
	},
	"prsvp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "prsvp",
			Port:  3455,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "prsvp",
			Port:  3455,
		},
	},
	"ps-ams": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ps-ams",
			Port:  3658,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ps-ams",
			Port:  3658,
		},
	},
	"psbserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "psbserver",
			Port:  2350,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "psbserver",
			Port:  2350,
		},
	},
	"pscl-mgt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pscl-mgt",
			Port:  4312,
		},
	},
	"pscribe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pscribe",
			Port:  6163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pscribe",
			Port:  6163,
		},
	},
	"pscupd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pscupd",
			Port:  3453,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pscupd",
			Port:  3453,
		},
	},
	"psdbserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "psdbserver",
			Port:  2355,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "psdbserver",
			Port:  2355,
		},
	},
	"pserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pserver",
			Port:  3662,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pserver",
			Port:  3662,
		},
	},
	"psi-ptt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "psi-ptt",
			Port:  4374,
		},
	},
	"pslicser": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pslicser",
			Port:  4168,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pslicser",
			Port:  4168,
		},
	},
	"pslserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pslserver",
			Port:  2352,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pslserver",
			Port:  2352,
		},
	},
	"psmond": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "psmond",
			Port:  1788,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "psmond",
			Port:  1788,
		},
	},
	"psprserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "psprserver",
			Port:  2354,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "psprserver",
			Port:  2354,
		},
	},
	"pspserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pspserver",
			Port:  2353,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pspserver",
			Port:  2353,
		},
	},
	"psrserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "psrserver",
			Port:  2351,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "psrserver",
			Port:  2351,
		},
	},
	"pss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pss",
			Port:  7880,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pss",
			Port:  7880,
		},
	},
	"pssc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pssc",
			Port:  645,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pssc",
			Port:  645,
		},
	},
	"pt2-discover": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pt2-discover",
			Port:  1101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pt2-discover",
			Port:  1101,
		},
	},
	"ptcnameservice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ptcnameservice",
			Port:  597,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ptcnameservice",
			Port:  597,
		},
	},
	"ptk-alink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ptk-alink",
			Port:  3089,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ptk-alink",
			Port:  3089,
		},
	},
	"ptp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ptp",
			Port:  15740,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ptp",
			Port:  15740,
		},
	},
	"ptp-event": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ptp-event",
			Port:  319,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ptp-event",
			Port:  319,
		},
	},
	"ptp-general": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ptp-general",
			Port:  320,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ptp-general",
			Port:  320,
		},
	},
	"publiqare-sync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "publiqare-sync",
			Port:  4329,
		},
	},
	"pulseaudio": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pulseaudio",
			Port:  4713,
		},
	},
	"pulsonixnls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pulsonixnls",
			Port:  6140,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pulsonixnls",
			Port:  6140,
		},
	},
	"puparp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "puparp",
			Port:  998,
		},
	},
	"purenoise": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "purenoise",
			Port:  663,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "purenoise",
			Port:  663,
		},
	},
	"pushns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pushns",
			Port:  7997,
		},
	},
	"pvaccess": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pvaccess",
			Port:  5075,
		},
	},
	"pvsw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pvsw",
			Port:  2520,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pvsw",
			Port:  2520,
		},
	},
	"pvsw-inet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pvsw-inet",
			Port:  2441,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pvsw-inet",
			Port:  2441,
		},
	},
	"pvuniwien": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pvuniwien",
			Port:  1081,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pvuniwien",
			Port:  1081,
		},
	},
	"pvxpluscs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pvxpluscs",
			Port:  4093,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pvxpluscs",
			Port:  4093,
		},
	},
	"pvxplusio": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pvxplusio",
			Port:  4193,
		},
	},
	"pwdgen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pwdgen",
			Port:  129,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pwdgen",
			Port:  129,
		},
	},
	"pwdis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pwdis",
			Port:  3735,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pwdis",
			Port:  3735,
		},
	},
	"pwgippfax": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pwgippfax",
			Port:  3951,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pwgippfax",
			Port:  3951,
		},
	},
	"pwgpsi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pwgpsi",
			Port:  3800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pwgpsi",
			Port:  3800,
		},
	},
	"pwgwims": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pwgwims",
			Port:  4951,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pwgwims",
			Port:  4951,
		},
	},
	"pwrsevent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pwrsevent",
			Port:  2694,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pwrsevent",
			Port:  2694,
		},
	},
	"pxc-epmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-epmap",
			Port:  2434,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-epmap",
			Port:  2434,
		},
	},
	"pxc-ntfy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-ntfy",
			Port:  3009,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-ntfy",
			Port:  3009,
		},
	},
	"pxc-pin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-pin",
			Port:  4005,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-pin",
			Port:  4005,
		},
	},
	"pxc-roid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-roid",
			Port:  4004,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-roid",
			Port:  4004,
		},
	},
	"pxc-sapxom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-sapxom",
			Port:  2680,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-sapxom",
			Port:  2680,
		},
	},
	"pxc-splr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-splr",
			Port:  4007,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-splr",
			Port:  4007,
		},
	},
	"pxc-splr-ft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-splr-ft",
			Port:  4003,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-splr-ft",
			Port:  4003,
		},
	},
	"pxc-spvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-spvr",
			Port:  4006,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-spvr",
			Port:  4006,
		},
	},
	"pxc-spvr-ft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pxc-spvr-ft",
			Port:  4002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pxc-spvr-ft",
			Port:  4002,
		},
	},
	"pxe": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "pxe",
			Port:  4011,
		},
	},
	"pyrrho": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "pyrrho",
			Port:  5433,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "pyrrho",
			Port:  5433,
		},
	},
	"q3ade": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "q3ade",
			Port:  7794,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "q3ade",
			Port:  7794,
		},
	},
	"q55-pcc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "q55-pcc",
			Port:  1253,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "q55-pcc",
			Port:  1253,
		},
	},
	"qadmifevent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qadmifevent",
			Port:  2462,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qadmifevent",
			Port:  2462,
		},
	},
	"qadmifoper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qadmifoper",
			Port:  2461,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qadmifoper",
			Port:  2461,
		},
	},
	"qb-db-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qb-db-server",
			Port:  10160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qb-db-server",
			Port:  10160,
		},
	},
	"qbdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qbdb",
			Port:  8019,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qbdb",
			Port:  8019,
		},
	},
	"qbikgdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qbikgdp",
			Port:  368,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qbikgdp",
			Port:  368,
		},
	},
	"qcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qcp",
			Port:  5082,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qcp",
			Port:  5082,
		},
	},
	"qdb2service": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qdb2service",
			Port:  45825,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qdb2service",
			Port:  45825,
		},
	},
	"qencp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qencp",
			Port:  2120,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qencp",
			Port:  2120,
		},
	},
	"qfp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qfp",
			Port:  5083,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qfp",
			Port:  5083,
		},
	},
	"qft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qft",
			Port:  189,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qft",
			Port:  189,
		},
	},
	"qftest-lookup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qftest-lookup",
			Port:  3543,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qftest-lookup",
			Port:  3543,
		},
	},
	"qip-audup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qip-audup",
			Port:  2765,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qip-audup",
			Port:  2765,
		},
	},
	"qip-login": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qip-login",
			Port:  2366,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qip-login",
			Port:  2366,
		},
	},
	"qip-msgd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qip-msgd",
			Port:  2468,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qip-msgd",
			Port:  2468,
		},
	},
	"qip-qdhcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qip-qdhcp",
			Port:  2490,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qip-qdhcp",
			Port:  2490,
		},
	},
	"qke-llc-v3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qke-llc-v3",
			Port:  2523,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qke-llc-v3",
			Port:  2523,
		},
	},
	"qmqp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qmqp",
			Port:  628,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qmqp",
			Port:  628,
		},
	},
	"qmtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qmtp",
			Port:  209,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qmtp",
			Port:  209,
		},
	},
	"qmvideo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qmvideo",
			Port:  5689,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qmvideo",
			Port:  5689,
		},
	},
	"qnts-orb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qnts-orb",
			Port:  1262,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qnts-orb",
			Port:  1262,
		},
	},
	"qnxnetman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qnxnetman",
			Port:  3385,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qnxnetman",
			Port:  3385,
		},
	},
	"qo-secure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qo-secure",
			Port:  7913,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qo-secure",
			Port:  7913,
		},
	},
	"qotd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qotd",
			Port:  17,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qotd",
			Port:  17,
		},
	},
	"qotps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qotps",
			Port:  2724,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qotps",
			Port:  2724,
		},
	},
	"qpasa-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qpasa-agent",
			Port:  2612,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qpasa-agent",
			Port:  2612,
		},
	},
	"qptlmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qptlmd",
			Port:  10055,
		},
	},
	"qrh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qrh",
			Port:  752,
		},
	},
	"qsm-gui": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsm-gui",
			Port:  1165,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsm-gui",
			Port:  1165,
		},
	},
	"qsm-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsm-proxy",
			Port:  1164,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsm-proxy",
			Port:  1164,
		},
	},
	"qsm-remote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsm-remote",
			Port:  1166,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsm-remote",
			Port:  1166,
		},
	},
	"qsnet-assist": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsnet-assist",
			Port:  4356,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsnet-assist",
			Port:  4356,
		},
	},
	"qsnet-cond": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsnet-cond",
			Port:  4357,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsnet-cond",
			Port:  4357,
		},
	},
	"qsnet-nucl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsnet-nucl",
			Port:  4358,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsnet-nucl",
			Port:  4358,
		},
	},
	"qsnet-trans": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsnet-trans",
			Port:  4354,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsnet-trans",
			Port:  4354,
		},
	},
	"qsnet-workst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsnet-workst",
			Port:  4355,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsnet-workst",
			Port:  4355,
		},
	},
	"qsoft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qsoft",
			Port:  3059,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qsoft",
			Port:  3059,
		},
	},
	"qt-serveradmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qt-serveradmin",
			Port:  1220,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qt-serveradmin",
			Port:  1220,
		},
	},
	"qtms-bootstrap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qtms-bootstrap",
			Port:  3850,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qtms-bootstrap",
			Port:  3850,
		},
	},
	"qtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qtp",
			Port:  2935,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qtp",
			Port:  2935,
		},
	},
	"quaddb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quaddb",
			Port:  2497,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quaddb",
			Port:  2497,
		},
	},
	"quailnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quailnet",
			Port:  5464,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quailnet",
			Port:  5464,
		},
	},
	"quake": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quake",
			Port:  26000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quake",
			Port:  26000,
		},
	},
	"quantastor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quantastor",
			Port:  8153,
		},
	},
	"quartus-tcl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quartus-tcl",
			Port:  2589,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quartus-tcl",
			Port:  2589,
		},
	},
	"quasar-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quasar-server",
			Port:  3599,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quasar-server",
			Port:  3599,
		},
	},
	"qubes": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qubes",
			Port:  1341,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qubes",
			Port:  1341,
		},
	},
	"quest-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quest-agent",
			Port:  3843,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quest-agent",
			Port:  3843,
		},
	},
	"quest-data-hub": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quest-data-hub",
			Port:  3566,
		},
	},
	"quest-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "quest-disc",
			Port:  7040,
		},
	},
	"quest-vista": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quest-vista",
			Port:  7980,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quest-vista",
			Port:  7980,
		},
	},
	"questdb2-lnchr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "questdb2-lnchr",
			Port:  5677,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "questdb2-lnchr",
			Port:  5677,
		},
	},
	"questnotify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "questnotify",
			Port:  3554,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "questnotify",
			Port:  3554,
		},
	},
	"queueadm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "queueadm",
			Port:  2230,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "queueadm",
			Port:  2230,
		},
	},
	"quickbooksrds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quickbooksrds",
			Port:  3790,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quickbooksrds",
			Port:  3790,
		},
	},
	"quicksuite": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quicksuite",
			Port:  2900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quicksuite",
			Port:  2900,
		},
	},
	"quosa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quosa",
			Port:  4841,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quosa",
			Port:  4841,
		},
	},
	"quotad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "quotad",
			Port:  762,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "quotad",
			Port:  762,
		},
	},
	"qvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qvr",
			Port:  5028,
		},
	},
	"qwave": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "qwave",
			Port:  2177,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "qwave",
			Port:  2177,
		},
	},
	"raadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raadmin",
			Port:  5676,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "raadmin",
			Port:  5676,
		},
	},
	"racf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "racf",
			Port:  18136,
		},
	},
	"radan-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radan-http",
			Port:  8088,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radan-http",
			Port:  8088,
		},
	},
	"radclientport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radclientport",
			Port:  3178,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radclientport",
			Port:  3178,
		},
	},
	"radec-corp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radec-corp",
			Port:  5430,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radec-corp",
			Port:  5430,
		},
	},
	"radio": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radio",
			Port:  1595,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radio",
			Port:  1595,
		},
	},
	"radio-bc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "radio-bc",
			Port:  1596,
		},
	},
	"radio-sm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radio-sm",
			Port:  1596,
		},
	},
	"radius": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radius",
			Port:  1812,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radius",
			Port:  1812,
		},
	},
	"radius-acct": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radius-acct",
			Port:  1813,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radius-acct",
			Port:  1813,
		},
	},
	"radius-dynauth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radius-dynauth",
			Port:  3799,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radius-dynauth",
			Port:  3799,
		},
	},
	"radix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radix",
			Port:  2872,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radix",
			Port:  2872,
		},
	},
	"radmin-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radmin-port",
			Port:  4899,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radmin-port",
			Port:  4899,
		},
	},
	"radmind": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radmind",
			Port:  6222,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radmind",
			Port:  6222,
		},
	},
	"radpdf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radpdf",
			Port:  18104,
		},
	},
	"rads": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rads",
			Port:  12302,
		},
	},
	"radsec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radsec",
			Port:  2083,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radsec",
			Port:  2083,
		},
	},
	"radware-rpm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radware-rpm",
			Port:  2188,
		},
	},
	"radware-rpm-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radware-rpm-s",
			Port:  2189,
		},
	},
	"radwiz-nms-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "radwiz-nms-srv",
			Port:  2736,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "radwiz-nms-srv",
			Port:  2736,
		},
	},
	"raid-ac": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "raid-ac",
			Port:  2012,
		},
	},
	"raid-am": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raid-am",
			Port:  2013,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "raid-am",
			Port:  2007,
		},
	},
	"raid-cc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raid-cc",
			Port:  2011,
		},
	},
	"raid-cd": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "raid-cd",
			Port:  2006,
		},
	},
	"raid-cs": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "raid-cs",
			Port:  2015,
		},
	},
	"raid-sf": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "raid-sf",
			Port:  2014,
		},
	},
	"railgun-webaccl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "railgun-webaccl",
			Port:  2408,
		},
	},
	"ramp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ramp",
			Port:  7227,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ramp",
			Port:  7227,
		},
	},
	"rap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rap",
			Port:  38,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rap",
			Port:  38,
		},
	},
	"rap-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rap-ip",
			Port:  3813,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rap-ip",
			Port:  3813,
		},
	},
	"rap-listen": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rap-listen",
			Port:  1531,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rap-listen",
			Port:  1531,
		},
	},
	"rap-service": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rap-service",
			Port:  1530,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rap-service",
			Port:  1530,
		},
	},
	"rapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rapi",
			Port:  2176,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rapi",
			Port:  2176,
		},
	},
	"rapidbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rapidbase",
			Port:  1953,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rapidbase",
			Port:  1953,
		},
	},
	"rapidmq-center": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rapidmq-center",
			Port:  3093,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rapidmq-center",
			Port:  3093,
		},
	},
	"rapidmq-reg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rapidmq-reg",
			Port:  3094,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rapidmq-reg",
			Port:  3094,
		},
	},
	"rapido-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rapido-ip",
			Port:  2457,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rapido-ip",
			Port:  2457,
		},
	},
	"raqmon-pdu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raqmon-pdu",
			Port:  7744,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "raqmon-pdu",
			Port:  7744,
		},
	},
	"rasadv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rasadv",
			Port:  9753,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rasadv",
			Port:  9753,
		},
	},
	"ratio-adp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ratio-adp",
			Port:  1108,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ratio-adp",
			Port:  1108,
		},
	},
	"ratl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ratl",
			Port:  2449,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ratl",
			Port:  2449,
		},
	},
	"ravehd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ravehd",
			Port:  4037,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ravehd",
			Port:  4037,
		},
	},
	"raven-rdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raven-rdp",
			Port:  3533,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "raven-rdp",
			Port:  3533,
		},
	},
	"raven-rmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raven-rmp",
			Port:  3532,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "raven-rmp",
			Port:  3532,
		},
	},
	"raventbs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raventbs",
			Port:  2713,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "raventbs",
			Port:  2713,
		},
	},
	"raventdm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raventdm",
			Port:  2714,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "raventdm",
			Port:  2714,
		},
	},
	"raw-serial": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raw-serial",
			Port:  2167,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "raw-serial",
			Port:  2167,
		},
	},
	"raxa-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "raxa-mgmt",
			Port:  6099,
		},
	},
	"razor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "razor",
			Port:  3555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "razor",
			Port:  3555,
		},
	},
	"rbakcup1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rbakcup1",
			Port:  2773,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rbakcup1",
			Port:  2773,
		},
	},
	"rbakcup2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rbakcup2",
			Port:  2774,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rbakcup2",
			Port:  2774,
		},
	},
	"rblcheckd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rblcheckd",
			Port:  3768,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rblcheckd",
			Port:  3768,
		},
	},
	"rbr-debug": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rbr-debug",
			Port:  44553,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rbr-debug",
			Port:  44553,
		},
	},
	"rbr-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rbr-discovery",
			Port:  3553,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rbr-discovery",
			Port:  3553,
		},
	},
	"rbsystem": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rbsystem",
			Port:  5693,
		},
	},
	"rbt-smc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rbt-smc",
			Port:  7870,
		},
	},
	"rbt-wanopt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rbt-wanopt",
			Port:  7810,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rbt-wanopt",
			Port:  7810,
		},
	},
	"rcc-host": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rcc-host",
			Port:  2332,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rcc-host",
			Port:  2332,
		},
	},
	"rcip-itu": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "rcip-itu",
			Port:  2225,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "rcip-itu",
			Port:  2225,
		},
	},
	"rcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rcp",
			Port:  469,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rcp",
			Port:  469,
		},
	},
	"rcst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rcst",
			Port:  3467,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rcst",
			Port:  3467,
		},
	},
	"rcts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rcts",
			Port:  2258,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rcts",
			Port:  2258,
		},
	},
	"rda": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rda",
			Port:  630,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rda",
			Port:  630,
		},
	},
	"rdb-dbs-disp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rdb-dbs-disp",
			Port:  1571,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rdb-dbs-disp",
			Port:  1571,
		},
	},
	"rdc-wh-eos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rdc-wh-eos",
			Port:  3142,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rdc-wh-eos",
			Port:  3142,
		},
	},
	"rdlap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rdlap",
			Port:  2321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rdlap",
			Port:  2321,
		},
	},
	"rdm-tfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rdm-tfs",
			Port:  21553,
		},
	},
	"rdmnet-ctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rdmnet-ctrl",
			Port:  5569,
		},
	},
	"rdmnet-device": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "rdmnet-device",
			Port:  5569,
		},
	},
	"rdrmshc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rdrmshc",
			Port:  1075,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rdrmshc",
			Port:  1075,
		},
	},
	"rds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rds",
			Port:  1540,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rds",
			Port:  1540,
		},
	},
	"rds-ib": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rds-ib",
			Port:  18634,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rds-ib",
			Port:  18634,
		},
	},
	"rds-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rds-ip",
			Port:  18635,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rds-ip",
			Port:  18635,
		},
	},
	"rds2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rds2",
			Port:  1541,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rds2",
			Port:  1541,
		},
	},
	"re-conn-proto": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "re-conn-proto",
			Port:  1306,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "re-conn-proto",
			Port:  1306,
		},
	},
	"re-mail-ck": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "re-mail-ck",
			Port:  50,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "re-mail-ck",
			Port:  50,
		},
	},
	"re101": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "re101",
			Port:  1343,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "re101",
			Port:  1343,
		},
	},
	"reachout": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "reachout",
			Port:  43188,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "reachout",
			Port:  43188,
		},
	},
	"realm-rusd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "realm-rusd",
			Port:  688,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "realm-rusd",
			Port:  688,
		},
	},
	"realsecure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "realsecure",
			Port:  2998,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "realsecure",
			Port:  2998,
		},
	},
	"rebol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rebol",
			Port:  2997,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rebol",
			Port:  2997,
		},
	},
	"recipe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "recipe",
			Port:  2240,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "recipe",
			Port:  2240,
		},
	},
	"recvr-rc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "recvr-rc",
			Port:  43000,
		},
	},
	"recvr-rc-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "recvr-rc-disc",
			Port:  43000,
		},
	},
	"redstone-cpss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "redstone-cpss",
			Port:  2928,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "redstone-cpss",
			Port:  2928,
		},
	},
	"redstorm_diag": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "redstorm_diag",
			Port:  2349,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "redstorm_diag",
			Port:  2349,
		},
	},
	"redstorm_find": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "redstorm_find",
			Port:  2347,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "redstorm_find",
			Port:  2347,
		},
	},
	"redstorm_info": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "redstorm_info",
			Port:  2348,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "redstorm_info",
			Port:  2348,
		},
	},
	"redstorm_join": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "redstorm_join",
			Port:  2346,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "redstorm_join",
			Port:  2346,
		},
	},
	"redwood-chat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "redwood-chat",
			Port:  3032,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "redwood-chat",
			Port:  3032,
		},
	},
	"reftek": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "reftek",
			Port:  2543,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "reftek",
			Port:  2543,
		},
	},
	"registrar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "registrar",
			Port:  1712,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "registrar",
			Port:  1712,
		},
	},
	"relief": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "relief",
			Port:  1353,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "relief",
			Port:  1353,
		},
	},
	"rellpack": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "rellpack",
			Port:  2018,
		},
	},
	"reload-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "reload-config",
			Port:  6084,
		},
	},
	"remcap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remcap",
			Port:  4731,
		},
	},
	"remctl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remctl",
			Port:  4373,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "remctl",
			Port:  4373,
		},
	},
	"remographlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remographlm",
			Port:  2373,
		},
	},
	"remote-as": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remote-as",
			Port:  1053,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "remote-as",
			Port:  1053,
		},
	},
	"remote-collab": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remote-collab",
			Port:  2250,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "remote-collab",
			Port:  2250,
		},
	},
	"remote-kis": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remote-kis",
			Port:  185,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "remote-kis",
			Port:  185,
		},
	},
	"remote-winsock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remote-winsock",
			Port:  1745,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "remote-winsock",
			Port:  1745,
		},
	},
	"remotedeploy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remotedeploy",
			Port:  3789,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "remotedeploy",
			Port:  3789,
		},
	},
	"remotefs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remotefs",
			Port:  556,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "remotefs",
			Port:  556,
		},
	},
	"remoteware-un": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "remoteware-un",
			Port:  2999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "remoteware-un",
			Port:  2999,
		},
	},
	"repcmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "repcmd",
			Port:  641,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "repcmd",
			Port:  641,
		},
	},
	"repliweb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "repliweb",
			Port:  2837,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "repliweb",
			Port:  2837,
		},
	},
	"repscmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "repscmd",
			Port:  653,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "repscmd",
			Port:  653,
		},
	},
	"repsvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "repsvc",
			Port:  6320,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "repsvc",
			Port:  6320,
		},
	},
	"res": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "res",
			Port:  1942,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "res",
			Port:  1942,
		},
	},
	"res-sap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "res-sap",
			Port:  3163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "res-sap",
			Port:  3163,
		},
	},
	"resacommunity": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "resacommunity",
			Port:  1154,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "resacommunity",
			Port:  1154,
		},
	},
	"rescap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rescap",
			Port:  283,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rescap",
			Port:  283,
		},
	},
	"resorcs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "resorcs",
			Port:  4733,
		},
	},
	"resource_mgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "resource_mgr",
			Port:  3019,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "resource_mgr",
			Port:  3019,
		},
	},
	"responselogic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "responselogic",
			Port:  2886,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "responselogic",
			Port:  2886,
		},
	},
	"responsenet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "responsenet",
			Port:  3045,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "responsenet",
			Port:  3045,
		},
	},
	"retp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "retp",
			Port:  32811,
		},
	},
	"retrospect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "retrospect",
			Port:  497,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "retrospect",
			Port:  497,
		},
	},
	"rets": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rets",
			Port:  6103,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rets",
			Port:  6103,
		},
	},
	"rets-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rets-ssl",
			Port:  12109,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rets-ssl",
			Port:  12109,
		},
	},
	"reversion": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "reversion",
			Port:  5842,
		},
	},
	"rexecj": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rexecj",
			Port:  8230,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rexecj",
			Port:  8230,
		},
	},
	"rfa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rfa",
			Port:  4672,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rfa",
			Port:  4672,
		},
	},
	"rfb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rfb",
			Port:  5900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rfb",
			Port:  5900,
		},
	},
	"rfe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rfe",
			Port:  5002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rfe",
			Port:  5002,
		},
	},
	"rfid-rp1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rfid-rp1",
			Port:  4684,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rfid-rp1",
			Port:  4684,
		},
	},
	"rfio": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rfio",
			Port:  3147,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rfio",
			Port:  3147,
		},
	},
	"rfmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rfmp",
			Port:  2249,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rfmp",
			Port:  2249,
		},
	},
	"rfx-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rfx-lm",
			Port:  1497,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rfx-lm",
			Port:  1497,
		},
	},
	"rgtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rgtp",
			Port:  1431,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rgtp",
			Port:  1431,
		},
	},
	"rhp-iibp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rhp-iibp",
			Port:  1912,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rhp-iibp",
			Port:  1912,
		},
	},
	"rib-slm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rib-slm",
			Port:  3296,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rib-slm",
			Port:  3296,
		},
	},
	"ricardo-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ricardo-lm",
			Port:  1522,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ricardo-lm",
			Port:  1522,
		},
	},
	"rich-cp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rich-cp",
			Port:  2057,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rich-cp",
			Port:  2057,
		},
	},
	"rid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rid",
			Port:  4590,
		},
	},
	"ridgeway1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ridgeway1",
			Port:  2776,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ridgeway1",
			Port:  2776,
		},
	},
	"ridgeway2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ridgeway2",
			Port:  2777,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ridgeway2",
			Port:  2777,
		},
	},
	"rimf-ps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rimf-ps",
			Port:  2209,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rimf-ps",
			Port:  2209,
		},
	},
	"rimsl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rimsl",
			Port:  2044,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rimsl",
			Port:  2044,
		},
	},
	"ripng": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ripng",
			Port:  521,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ripng",
			Port:  521,
		},
	},
	"ris": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ris",
			Port:  180,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ris",
			Port:  180,
		},
	},
	"ris-cm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ris-cm",
			Port:  748,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ris-cm",
			Port:  748,
		},
	},
	"rise": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rise",
			Port:  7473,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rise",
			Port:  7473,
		},
	},
	"rjcdb-vcards": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rjcdb-vcards",
			Port:  9208,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rjcdb-vcards",
			Port:  9208,
		},
	},
	"rje": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rje",
			Port:  5,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rje",
			Port:  5,
		},
	},
	"rkb-oscs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rkb-oscs",
			Port:  1817,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rkb-oscs",
			Port:  1817,
		},
	},
	"rlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rlm",
			Port:  5053,
		},
	},
	"rlm-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rlm-admin",
			Port:  5054,
		},
	},
	"rlm-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "rlm-disc",
			Port:  5053,
		},
	},
	"rlp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rlp",
			Port:  39,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rlp",
			Port:  39,
		},
	},
	"rlzdbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rlzdbase",
			Port:  635,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rlzdbase",
			Port:  635,
		},
	},
	"rmc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmc",
			Port:  657,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmc",
			Port:  657,
		},
	},
	"rmiactivation": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmiactivation",
			Port:  1098,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmiactivation",
			Port:  1098,
		},
	},
	"rmiaux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmiaux",
			Port:  10990,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmiaux",
			Port:  10990,
		},
	},
	"rmiregistry": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmiregistry",
			Port:  1099,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmiregistry",
			Port:  1099,
		},
	},
	"rmlnk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmlnk",
			Port:  2818,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmlnk",
			Port:  2818,
		},
	},
	"rmonitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmonitor",
			Port:  560,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmonitor",
			Port:  560,
		},
	},
	"rmonitor_secure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmonitor_secure",
			Port:  5145,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmonitor_secure",
			Port:  5145,
		},
	},
	"rmopagt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmopagt",
			Port:  2959,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmopagt",
			Port:  2959,
		},
	},
	"rmpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmpp",
			Port:  1121,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmpp",
			Port:  1121,
		},
	},
	"rmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmt",
			Port:  411,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmt",
			Port:  411,
		},
	},
	"rmtserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rmtserver",
			Port:  2416,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rmtserver",
			Port:  2416,
		},
	},
	"rna": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "rna",
			Port:  25471,
		},
	},
	"rndc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rndc",
			Port:  953,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rndc",
			Port:  953,
		},
	},
	"rnm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rnm",
			Port:  3844,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rnm",
			Port:  3844,
		},
	},
	"rnmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rnmap",
			Port:  3418,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rnmap",
			Port:  3418,
		},
	},
	"rnrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rnrp",
			Port:  2423,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rnrp",
			Port:  2423,
		},
	},
	"robcad-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "robcad-lm",
			Port:  1509,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "robcad-lm",
			Port:  1509,
		},
	},
	"robix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "robix",
			Port:  9599,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "robix",
			Port:  9599,
		},
	},
	"roboeda": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "roboeda",
			Port:  2920,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "roboeda",
			Port:  2920,
		},
	},
	"roboer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "roboer",
			Port:  2919,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "roboer",
			Port:  2919,
		},
	},
	"rockwell-csp1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rockwell-csp1",
			Port:  2221,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rockwell-csp1",
			Port:  2221,
		},
	},
	"rockwell-csp2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rockwell-csp2",
			Port:  2223,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rockwell-csp2",
			Port:  2223,
		},
	},
	"rocrail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rocrail",
			Port:  8051,
		},
	},
	"roketz": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "roketz",
			Port:  1730,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "roketz",
			Port:  1730,
		},
	},
	"rootd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rootd",
			Port:  1094,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rootd",
			Port:  1094,
		},
	},
	"routematch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "routematch",
			Port:  1287,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "routematch",
			Port:  1287,
		},
	},
	"router": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "router",
			Port:  520,
		},
	},
	"roverlog": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "roverlog",
			Port:  3677,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "roverlog",
			Port:  3677,
		},
	},
	"rp-reputation": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "rp-reputation",
			Port:  6568,
		},
	},
	"rpasswd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rpasswd",
			Port:  774,
		},
	},
	"rpc2portmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rpc2portmap",
			Port:  369,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rpc2portmap",
			Port:  369,
		},
	},
	"rpi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rpi",
			Port:  2214,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rpi",
			Port:  2214,
		},
	},
	"rpki-rtr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rpki-rtr",
			Port:  323,
		},
	},
	"rpki-rtr-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rpki-rtr-tls",
			Port:  324,
		},
	},
	"rprt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rprt",
			Port:  3064,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rprt",
			Port:  3064,
		},
	},
	"rquotad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rquotad",
			Port:  875,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rquotad",
			Port:  875,
		},
	},
	"rrac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrac",
			Port:  5678,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrac",
			Port:  5678,
		},
	},
	"rrdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrdp",
			Port:  5313,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrdp",
			Port:  5313,
		},
	},
	"rrh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrh",
			Port:  753,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrh",
			Port:  753,
		},
	},
	"rrifmm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrifmm",
			Port:  1696,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrifmm",
			Port:  1696,
		},
	},
	"rrilwm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrilwm",
			Port:  1695,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrilwm",
			Port:  1695,
		},
	},
	"rrimwm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrimwm",
			Port:  1694,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrimwm",
			Port:  1694,
		},
	},
	"rrirtr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrirtr",
			Port:  1693,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrirtr",
			Port:  1693,
		},
	},
	"rrisat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrisat",
			Port:  1697,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrisat",
			Port:  1697,
		},
	},
	"rrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rrp",
			Port:  648,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rrp",
			Port:  648,
		},
	},
	"rs-pias": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rs-pias",
			Port:  13217,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rs-pias",
			Port:  13217,
		},
	},
	"rs-rmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rs-rmi",
			Port:  3736,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rs-rmi",
			Port:  3736,
		},
	},
	"rsap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsap",
			Port:  1647,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsap",
			Port:  1647,
		},
	},
	"rsc-robot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsc-robot",
			Port:  1793,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsc-robot",
			Port:  1793,
		},
	},
	"rscd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rscd",
			Port:  5750,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rscd",
			Port:  5750,
		},
	},
	"rscs": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "rscs",
			Port:  10201,
		},
	},
	"rsf-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsf-1",
			Port:  1195,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsf-1",
			Port:  1195,
		},
	},
	"rsh-spx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsh-spx",
			Port:  222,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsh-spx",
			Port:  222,
		},
	},
	"rsip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsip",
			Port:  4555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsip",
			Port:  4555,
		},
	},
	"rsisysaccess": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsisysaccess",
			Port:  2752,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsisysaccess",
			Port:  2752,
		},
	},
	"rsms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsms",
			Port:  10201,
		},
	},
	"rsmtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsmtp",
			Port:  2390,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsmtp",
			Port:  2390,
		},
	},
	"rsom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsom",
			Port:  2889,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsom",
			Port:  2889,
		},
	},
	"rsqlserver": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "rsqlserver",
			Port:  4430,
		},
	},
	"rsvd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsvd",
			Port:  168,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsvd",
			Port:  168,
		},
	},
	"rsvp-encap-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsvp-encap-1",
			Port:  1698,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsvp-encap-1",
			Port:  1698,
		},
	},
	"rsvp-encap-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsvp-encap-2",
			Port:  1699,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsvp-encap-2",
			Port:  1699,
		},
	},
	"rsvp_tunnel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsvp_tunnel",
			Port:  363,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsvp_tunnel",
			Port:  363,
		},
	},
	"rsync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rsync",
			Port:  873,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rsync",
			Port:  873,
		},
	},
	"rt-event": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rt-event",
			Port:  3706,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rt-event",
			Port:  3706,
		},
	},
	"rt-event-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rt-event-s",
			Port:  3707,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rt-event-s",
			Port:  3707,
		},
	},
	"rtc-pm-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtc-pm-port",
			Port:  3891,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtc-pm-port",
			Port:  3891,
		},
	},
	"rtcm-sc104": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtcm-sc104",
			Port:  2101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtcm-sc104",
			Port:  2101,
		},
	},
	"rtelnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtelnet",
			Port:  107,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtelnet",
			Port:  107,
		},
	},
	"rtip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtip",
			Port:  771,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtip",
			Port:  771,
		},
	},
	"rtmp": map[string]Service{
		"ddp": Service{
			Proto: "ddp",
			Name:  "rtmp",
			Port:  1,
		},
	},
	"rtmp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtmp-port",
			Port:  3500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtmp-port",
			Port:  3500,
		},
	},
	"rtnt-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtnt-1",
			Port:  3137,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtnt-1",
			Port:  3137,
		},
	},
	"rtnt-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtnt-2",
			Port:  3138,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtnt-2",
			Port:  3138,
		},
	},
	"rtps-dd-mt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtps-dd-mt",
			Port:  7402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtps-dd-mt",
			Port:  7402,
		},
	},
	"rtps-dd-ut": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtps-dd-ut",
			Port:  7401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtps-dd-ut",
			Port:  7401,
		},
	},
	"rtps-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtps-discovery",
			Port:  7400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtps-discovery",
			Port:  7400,
		},
	},
	"rtraceroute": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtraceroute",
			Port:  3765,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtraceroute",
			Port:  3765,
		},
	},
	"rtsclient": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtsclient",
			Port:  2501,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtsclient",
			Port:  2501,
		},
	},
	"rtsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtsp",
			Port:  554,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtsp",
			Port:  554,
		},
	},
	"rtsp-alt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtsp-alt",
			Port:  8554,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtsp-alt",
			Port:  8554,
		},
	},
	"rtsps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtsps",
			Port:  322,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtsps",
			Port:  322,
		},
	},
	"rtsserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rtsserv",
			Port:  2500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rtsserv",
			Port:  2500,
		},
	},
	"rugameonline": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rugameonline",
			Port:  5156,
		},
	},
	"rusb-sys-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rusb-sys-port",
			Port:  3422,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rusb-sys-port",
			Port:  3422,
		},
	},
	"rushd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rushd",
			Port:  696,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rushd",
			Port:  696,
		},
	},
	"rvs-isdn-dcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rvs-isdn-dcp",
			Port:  2578,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rvs-isdn-dcp",
			Port:  2578,
		},
	},
	"rwhois": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rwhois",
			Port:  4321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rwhois",
			Port:  4321,
		},
	},
	"rxapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rxapi",
			Port:  10010,
		},
	},
	"rxe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rxe",
			Port:  761,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rxe",
			Port:  761,
		},
	},
	"rxmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "rxmon",
			Port:  1311,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "rxmon",
			Port:  1311,
		},
	},
	"s-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "s-net",
			Port:  166,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "s-net",
			Port:  166,
		},
	},
	"s-openmail": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "s-openmail",
			Port:  5767,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "s-openmail",
			Port:  5767,
		},
	},
	"s1-control": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "s1-control",
			Port:  36412,
		},
	},
	"s102": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "s102",
			Port:  23272,
		},
	},
	"s3db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "s3db",
			Port:  2278,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "s3db",
			Port:  2278,
		},
	},
	"s8-client-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "s8-client-port",
			Port:  3153,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "s8-client-port",
			Port:  3153,
		},
	},
	"sa-msg-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sa-msg-port",
			Port:  1646,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sa-msg-port",
			Port:  1646,
		},
	},
	"sabams": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sabams",
			Port:  2760,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sabams",
			Port:  2760,
		},
	},
	"sabarsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sabarsd",
			Port:  8401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sabarsd",
			Port:  8401,
		},
	},
	"sabp-signal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sabp-signal",
			Port:  3452,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sabp-signal",
			Port:  3452,
		},
	},
	"sac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sac",
			Port:  8097,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sac",
			Port:  8097,
		},
	},
	"sacred": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sacred",
			Port:  1118,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sacred",
			Port:  1118,
		},
	},
	"safetynetp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "safetynetp",
			Port:  40000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "safetynetp",
			Port:  40000,
		},
	},
	"saft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "saft",
			Port:  487,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "saft",
			Port:  487,
		},
	},
	"sage-best-com1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sage-best-com1",
			Port:  14033,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sage-best-com1",
			Port:  14033,
		},
	},
	"sage-best-com2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sage-best-com2",
			Port:  14034,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sage-best-com2",
			Port:  14034,
		},
	},
	"sagectlpanel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sagectlpanel",
			Port:  3698,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sagectlpanel",
			Port:  3698,
		},
	},
	"sagxtsds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sagxtsds",
			Port:  4952,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sagxtsds",
			Port:  4952,
		},
	},
	"sah-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sah-lm",
			Port:  3291,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sah-lm",
			Port:  3291,
		},
	},
	"sai_sentlm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sai_sentlm",
			Port:  2640,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sai_sentlm",
			Port:  2640,
		},
	},
	"sais": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sais",
			Port:  1426,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sais",
			Port:  1426,
		},
	},
	"saiscm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "saiscm",
			Port:  1501,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "saiscm",
			Port:  1501,
		},
	},
	"saiseh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "saiseh",
			Port:  1644,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "saiseh",
			Port:  1644,
		},
	},
	"saism": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "saism",
			Port:  1436,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "saism",
			Port:  1436,
		},
	},
	"salient-dtasrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "salient-dtasrv",
			Port:  5409,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "salient-dtasrv",
			Port:  5409,
		},
	},
	"salient-mux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "salient-mux",
			Port:  5422,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "salient-mux",
			Port:  5422,
		},
	},
	"salient-usrmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "salient-usrmgr",
			Port:  5410,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "salient-usrmgr",
			Port:  5410,
		},
	},
	"samd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "samd",
			Port:  3275,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "samd",
			Port:  3275,
		},
	},
	"samsung-unidex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "samsung-unidex",
			Port:  4010,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "samsung-unidex",
			Port:  4010,
		},
	},
	"sanavigator": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sanavigator",
			Port:  4033,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sanavigator",
			Port:  4033,
		},
	},
	"sane-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sane-port",
			Port:  6566,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sane-port",
			Port:  6566,
		},
	},
	"sanity": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sanity",
			Port:  643,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sanity",
			Port:  643,
		},
	},
	"santak-ups": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "santak-ups",
			Port:  3038,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "santak-ups",
			Port:  3038,
		},
	},
	"saphostctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "saphostctrl",
			Port:  1128,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "saphostctrl",
			Port:  1128,
		},
	},
	"saphostctrls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "saphostctrls",
			Port:  1129,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "saphostctrls",
			Port:  1129,
		},
	},
	"sapv1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sapv1",
			Port:  9875,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sapv1",
			Port:  9875,
		},
	},
	"saratoga": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "saratoga",
			Port:  7542,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "saratoga",
			Port:  7542,
		},
	},
	"saris": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "saris",
			Port:  4442,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "saris",
			Port:  4442,
		},
	},
	"sas-remote-hlp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sas-remote-hlp",
			Port:  3755,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sas-remote-hlp",
			Port:  3755,
		},
	},
	"sasg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sasg",
			Port:  3744,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sasg",
			Port:  3744,
		},
	},
	"sasggprs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sasggprs",
			Port:  3964,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sasggprs",
			Port:  3964,
		},
	},
	"sasp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sasp",
			Port:  3860,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sasp",
			Port:  3860,
		},
	},
	"sauterdongle": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sauterdongle",
			Port:  25576,
		},
	},
	"savant": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "savant",
			Port:  3391,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "savant",
			Port:  3391,
		},
	},
	"sbackup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sbackup",
			Port:  5163,
		},
	},
	"sbcap": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "sbcap",
			Port:  29168,
		},
	},
	"sbi-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sbi-agent",
			Port:  3962,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sbi-agent",
			Port:  3962,
		},
	},
	"sbl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sbl",
			Port:  1039,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sbl",
			Port:  1039,
		},
	},
	"sbook": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sbook",
			Port:  1349,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sbook",
			Port:  1349,
		},
	},
	"scan-change": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scan-change",
			Port:  2719,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scan-change",
			Port:  2719,
		},
	},
	"scanstat-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scanstat-1",
			Port:  1215,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scanstat-1",
			Port:  1215,
		},
	},
	"scc-security": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scc-security",
			Port:  582,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scc-security",
			Port:  582,
		},
	},
	"sccip-media": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sccip-media",
			Port:  3499,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sccip-media",
			Port:  3499,
		},
	},
	"sceanics": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sceanics",
			Port:  5435,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sceanics",
			Port:  5435,
		},
	},
	"scenccs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scenccs",
			Port:  7129,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scenccs",
			Port:  7129,
		},
	},
	"scenidm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scenidm",
			Port:  7128,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scenidm",
			Port:  7128,
		},
	},
	"scientia-sdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scientia-sdb",
			Port:  1811,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scientia-sdb",
			Port:  1811,
		},
	},
	"scientia-ssdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scientia-ssdb",
			Port:  2121,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scientia-ssdb",
			Port:  2121,
		},
	},
	"scinet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scinet",
			Port:  7708,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scinet",
			Port:  7708,
		},
	},
	"scipticslsrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scipticslsrvr",
			Port:  2577,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scipticslsrvr",
			Port:  2577,
		},
	},
	"sco-aip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sco-aip",
			Port:  5307,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sco-aip",
			Port:  5307,
		},
	},
	"sco-dtmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sco-dtmgr",
			Port:  617,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sco-dtmgr",
			Port:  617,
		},
	},
	"sco-inetmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sco-inetmgr",
			Port:  615,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sco-inetmgr",
			Port:  615,
		},
	},
	"sco-peer-tta": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sco-peer-tta",
			Port:  5427,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sco-peer-tta",
			Port:  5427,
		},
	},
	"sco-sysmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sco-sysmgr",
			Port:  616,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sco-sysmgr",
			Port:  616,
		},
	},
	"sco-websrvrmg3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sco-websrvrmg3",
			Port:  598,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sco-websrvrmg3",
			Port:  598,
		},
	},
	"sco-websrvrmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sco-websrvrmgr",
			Port:  620,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sco-websrvrmgr",
			Port:  620,
		},
	},
	"scohelp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scohelp",
			Port:  457,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scohelp",
			Port:  457,
		},
	},
	"scoi2odialog": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scoi2odialog",
			Port:  360,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scoi2odialog",
			Port:  360,
		},
	},
	"scol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scol",
			Port:  1200,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scol",
			Port:  1200,
		},
	},
	"scoremgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scoremgr",
			Port:  2034,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scoremgr",
			Port:  2034,
		},
	},
	"scotty-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "scotty-disc",
			Port:  14002,
		},
	},
	"scotty-ft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scotty-ft",
			Port:  14000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scotty-ft",
			Port:  14000,
		},
	},
	"scp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scp",
			Port:  3820,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scp",
			Port:  3820,
		},
	},
	"scp-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scp-config",
			Port:  10001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scp-config",
			Port:  10001,
		},
	},
	"scpi-raw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scpi-raw",
			Port:  5025,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scpi-raw",
			Port:  5025,
		},
	},
	"scpi-telnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scpi-telnet",
			Port:  5024,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scpi-telnet",
			Port:  5024,
		},
	},
	"scrabble": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scrabble",
			Port:  2026,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scrabble",
			Port:  2026,
		},
	},
	"screencast": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "screencast",
			Port:  1368,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "screencast",
			Port:  1368,
		},
	},
	"scriptview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scriptview",
			Port:  7741,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scriptview",
			Port:  7741,
		},
	},
	"scscp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scscp",
			Port:  26133,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scscp",
			Port:  26133,
		},
	},
	"scservp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scservp",
			Port:  3637,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scservp",
			Port:  3637,
		},
	},
	"scte104": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scte104",
			Port:  5167,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scte104",
			Port:  5167,
		},
	},
	"scte30": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scte30",
			Port:  5168,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scte30",
			Port:  5168,
		},
	},
	"sctp-tunneling": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "sctp-tunneling",
			Port:  9899,
		},
	},
	"scup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scup",
			Port:  6315,
		},
	},
	"scup-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "scup-disc",
			Port:  6315,
		},
	},
	"scx-proxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "scx-proxy",
			Port:  470,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "scx-proxy",
			Port:  470,
		},
	},
	"sd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sd",
			Port:  9876,
		},
	},
	"sd-capacity": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "sd-capacity",
			Port:  2384,
		},
	},
	"sd-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sd-data",
			Port:  2385,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sd-data",
			Port:  2385,
		},
	},
	"sd-elmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sd-elmd",
			Port:  1681,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sd-elmd",
			Port:  1681,
		},
	},
	"sd-request": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sd-request",
			Port:  2384,
		},
	},
	"sdbproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdbproxy",
			Port:  3562,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdbproxy",
			Port:  3562,
		},
	},
	"sdclient": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdclient",
			Port:  2310,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdclient",
			Port:  2310,
		},
	},
	"sddp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sddp",
			Port:  1163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sddp",
			Port:  1163,
		},
	},
	"sde-discovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sde-discovery",
			Port:  5152,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sde-discovery",
			Port:  5152,
		},
	},
	"sdfunc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdfunc",
			Port:  2046,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdfunc",
			Port:  2046,
		},
	},
	"sdhelp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdhelp",
			Port:  2308,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdhelp",
			Port:  2308,
		},
	},
	"sdl-ets": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdl-ets",
			Port:  5081,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdl-ets",
			Port:  5081,
		},
	},
	"sdmmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdmmp",
			Port:  5573,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdmmp",
			Port:  5573,
		},
	},
	"sdnskmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdnskmp",
			Port:  558,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdnskmp",
			Port:  558,
		},
	},
	"sdo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdo",
			Port:  3635,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdo",
			Port:  3635,
		},
	},
	"sdo-ssh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdo-ssh",
			Port:  3897,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdo-ssh",
			Port:  3897,
		},
	},
	"sdo-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdo-tls",
			Port:  3896,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdo-tls",
			Port:  3896,
		},
	},
	"sdp-id-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdp-id-port",
			Port:  3242,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdp-id-port",
			Port:  3242,
		},
	},
	"sdp-portmapper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdp-portmapper",
			Port:  3935,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdp-portmapper",
			Port:  3935,
		},
	},
	"sdproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdproxy",
			Port:  1297,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdproxy",
			Port:  1297,
		},
	},
	"sdr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdr",
			Port:  9010,
		},
	},
	"sds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sds",
			Port:  5059,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sds",
			Port:  5059,
		},
	},
	"sds-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sds-admin",
			Port:  2705,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sds-admin",
			Port:  2705,
		},
	},
	"sdsc-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdsc-lm",
			Port:  1537,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdsc-lm",
			Port:  1537,
		},
	},
	"sdserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdserver",
			Port:  2309,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdserver",
			Port:  2309,
		},
	},
	"sdt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdt",
			Port:  5568,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdt",
			Port:  5568,
		},
	},
	"sdt-lmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sdt-lmd",
			Port:  3319,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sdt-lmd",
			Port:  3319,
		},
	},
	"seagull-ais": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "seagull-ais",
			Port:  1208,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "seagull-ais",
			Port:  1208,
		},
	},
	"seagulllms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "seagulllms",
			Port:  1291,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "seagulllms",
			Port:  1291,
		},
	},
	"seaodbc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "seaodbc",
			Port:  2471,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "seaodbc",
			Port:  2471,
		},
	},
	"search": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "search",
			Port:  2010,
		},
	},
	"search-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "search-agent",
			Port:  1234,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "search-agent",
			Port:  1234,
		},
	},
	"seaview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "seaview",
			Port:  3143,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "seaview",
			Port:  3143,
		},
	},
	"sec-ntb-clnt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sec-ntb-clnt",
			Port:  32635,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sec-ntb-clnt",
			Port:  32635,
		},
	},
	"sec-pc2fax-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sec-pc2fax-srv",
			Port:  9402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sec-pc2fax-srv",
			Port:  9402,
		},
	},
	"sec-t4net-clt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sec-t4net-clt",
			Port:  9401,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sec-t4net-clt",
			Port:  9401,
		},
	},
	"sec-t4net-srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sec-t4net-srv",
			Port:  9400,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sec-t4net-srv",
			Port:  9400,
		},
	},
	"seclayer-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "seclayer-tcp",
			Port:  3495,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "seclayer-tcp",
			Port:  3495,
		},
	},
	"seclayer-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "seclayer-tls",
			Port:  3496,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "seclayer-tls",
			Port:  3496,
		},
	},
	"secrmmsafecopya": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "secrmmsafecopya",
			Port:  38865,
		},
	},
	"secure-cfg-svr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "secure-cfg-svr",
			Port:  3978,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "secure-cfg-svr",
			Port:  3978,
		},
	},
	"secure-mqtt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "secure-mqtt",
			Port:  8883,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "secure-mqtt",
			Port:  8883,
		},
	},
	"secure-ts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "secure-ts",
			Port:  9318,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "secure-ts",
			Port:  9318,
		},
	},
	"securitychase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "securitychase",
			Port:  5399,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "securitychase",
			Port:  5399,
		},
	},
	"seispoc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "seispoc",
			Port:  2254,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "seispoc",
			Port:  2254,
		},
	},
	"semantix": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "semantix",
			Port:  361,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "semantix",
			Port:  361,
		},
	},
	"semaphore": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "semaphore",
			Port:  3255,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "semaphore",
			Port:  3255,
		},
	},
	"send": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "send",
			Port:  169,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "send",
			Port:  169,
		},
	},
	"senip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senip",
			Port:  3898,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senip",
			Port:  3898,
		},
	},
	"senomix01": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senomix01",
			Port:  8052,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senomix01",
			Port:  8052,
		},
	},
	"senomix02": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senomix02",
			Port:  8053,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senomix02",
			Port:  8053,
		},
	},
	"senomix03": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senomix03",
			Port:  8054,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senomix03",
			Port:  8054,
		},
	},
	"senomix04": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senomix04",
			Port:  8055,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senomix04",
			Port:  8055,
		},
	},
	"senomix05": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senomix05",
			Port:  8056,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senomix05",
			Port:  8056,
		},
	},
	"senomix06": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senomix06",
			Port:  8057,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senomix06",
			Port:  8057,
		},
	},
	"senomix07": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senomix07",
			Port:  8058,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senomix07",
			Port:  8058,
		},
	},
	"senomix08": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "senomix08",
			Port:  8059,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "senomix08",
			Port:  8059,
		},
	},
	"sent-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sent-lm",
			Port:  2316,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sent-lm",
			Port:  2316,
		},
	},
	"sentinel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sentinel",
			Port:  3588,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sentinel",
			Port:  3588,
		},
	},
	"sentinel-ent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sentinel-ent",
			Port:  3712,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sentinel-ent",
			Port:  3712,
		},
	},
	"sentinel-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sentinel-lm",
			Port:  5093,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sentinel-lm",
			Port:  5093,
		},
	},
	"sentinelsrm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sentinelsrm",
			Port:  1947,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sentinelsrm",
			Port:  1947,
		},
	},
	"sentlm-srv2srv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sentlm-srv2srv",
			Port:  5099,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sentlm-srv2srv",
			Port:  5099,
		},
	},
	"sep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sep",
			Port:  2089,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sep",
			Port:  2089,
		},
	},
	"seraph": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "seraph",
			Port:  4076,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "seraph",
			Port:  4076,
		},
	},
	"sercomm-scadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sercomm-scadmin",
			Port:  6108,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sercomm-scadmin",
			Port:  6108,
		},
	},
	"sercomm-wlink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sercomm-wlink",
			Port:  2235,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sercomm-wlink",
			Port:  2235,
		},
	},
	"serialgateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serialgateway",
			Port:  1243,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serialgateway",
			Port:  1243,
		},
	},
	"server-find": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "server-find",
			Port:  8351,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "server-find",
			Port:  8351,
		},
	},
	"servergraph": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "servergraph",
			Port:  1251,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "servergraph",
			Port:  1251,
		},
	},
	"serverstart": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serverstart",
			Port:  9213,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serverstart",
			Port:  9213,
		},
	},
	"serverview-as": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serverview-as",
			Port:  3169,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serverview-as",
			Port:  3169,
		},
	},
	"serverview-asn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serverview-asn",
			Port:  3170,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serverview-asn",
			Port:  3170,
		},
	},
	"serverview-gf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serverview-gf",
			Port:  3171,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serverview-gf",
			Port:  3171,
		},
	},
	"serverview-icc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serverview-icc",
			Port:  3173,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serverview-icc",
			Port:  3173,
		},
	},
	"serverview-rm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serverview-rm",
			Port:  3172,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serverview-rm",
			Port:  3172,
		},
	},
	"serverviewdbms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serverviewdbms",
			Port:  9212,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serverviewdbms",
			Port:  9212,
		},
	},
	"serverwsd2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "serverwsd2",
			Port:  5362,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "serverwsd2",
			Port:  5362,
		},
	},
	"servexec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "servexec",
			Port:  2021,
		},
	},
	"service-ctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "service-ctrl",
			Port:  2367,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "service-ctrl",
			Port:  2367,
		},
	},
	"servicemeter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "servicemeter",
			Port:  2603,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "servicemeter",
			Port:  2603,
		},
	},
	"servicetags": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "servicetags",
			Port:  6481,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "servicetags",
			Port:  6481,
		},
	},
	"servistaitsm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "servistaitsm",
			Port:  3636,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "servistaitsm",
			Port:  3636,
		},
	},
	"servserv": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "servserv",
			Port:  2011,
		},
	},
	"servstat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "servstat",
			Port:  633,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "servstat",
			Port:  633,
		},
	},
	"sesi-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sesi-lm",
			Port:  1714,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sesi-lm",
			Port:  1714,
		},
	},
	"set": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "set",
			Port:  257,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "set",
			Port:  257,
		},
	},
	"sf-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sf-lm",
			Port:  4546,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sf-lm",
			Port:  4546,
		},
	},
	"sflm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sflm",
			Port:  3162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sflm",
			Port:  3162,
		},
	},
	"sflow": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sflow",
			Port:  6343,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sflow",
			Port:  6343,
		},
	},
	"sfm-db-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sfm-db-server",
			Port:  5636,
		},
	},
	"sfmsso": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sfmsso",
			Port:  5635,
		},
	},
	"sfs-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sfs-config",
			Port:  452,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sfs-config",
			Port:  452,
		},
	},
	"sfs-smp-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sfs-smp-net",
			Port:  451,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sfs-smp-net",
			Port:  451,
		},
	},
	"sftdst-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sftdst-port",
			Port:  3230,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sftdst-port",
			Port:  3230,
		},
	},
	"sftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sftp",
			Port:  115,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sftp",
			Port:  115,
		},
	},
	"sftsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sftsrv",
			Port:  1303,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sftsrv",
			Port:  1303,
		},
	},
	"sftu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sftu",
			Port:  3326,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sftu",
			Port:  3326,
		},
	},
	"sg-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sg-lm",
			Port:  1659,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sg-lm",
			Port:  1659,
		},
	},
	"sgcip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgcip",
			Port:  16950,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgcip",
			Port:  16950,
		},
	},
	"sgcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgcp",
			Port:  440,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgcp",
			Port:  440,
		},
	},
	"sge_execd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sge_execd",
			Port:  6445,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sge_execd",
			Port:  6445,
		},
	},
	"sge_qmaster": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sge_qmaster",
			Port:  6444,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sge_qmaster",
			Port:  6444,
		},
	},
	"sgi-arrayd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgi-arrayd",
			Port:  5434,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgi-arrayd",
			Port:  5434,
		},
	},
	"sgi-dgl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgi-dgl",
			Port:  5232,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgi-dgl",
			Port:  5232,
		},
	},
	"sgi-dmfmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgi-dmfmgr",
			Port:  11109,
		},
	},
	"sgi-esphttp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgi-esphttp",
			Port:  5554,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgi-esphttp",
			Port:  5554,
		},
	},
	"sgi-eventmond": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgi-eventmond",
			Port:  5553,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgi-eventmond",
			Port:  5553,
		},
	},
	"sgi-lk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgi-lk",
			Port:  11106,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgi-lk",
			Port:  11106,
		},
	},
	"sgi-soap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgi-soap",
			Port:  11110,
		},
	},
	"sgi-storman": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "sgi-storman",
			Port:  1178,
		},
	},
	"sgmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgmp",
			Port:  153,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgmp",
			Port:  153,
		},
	},
	"sgmp-traps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sgmp-traps",
			Port:  160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sgmp-traps",
			Port:  160,
		},
	},
	"sgsap": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "sgsap",
			Port:  29118,
		},
	},
	"shadowserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shadowserver",
			Port:  2027,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "shadowserver",
			Port:  2027,
		},
	},
	"shaperai": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shaperai",
			Port:  43210,
		},
	},
	"shaperai-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "shaperai-disc",
			Port:  43210,
		},
	},
	"shareapp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shareapp",
			Port:  3595,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "shareapp",
			Port:  3595,
		},
	},
	"sharp-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sharp-server",
			Port:  3617,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sharp-server",
			Port:  3617,
		},
	},
	"shell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shell",
			Port:  514,
		},
	},
	"shiva_confsrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shiva_confsrvr",
			Port:  1651,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "shiva_confsrvr",
			Port:  1651,
		},
	},
	"shivadiscovery": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shivadiscovery",
			Port:  1502,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "shivadiscovery",
			Port:  1502,
		},
	},
	"shivahose": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shivahose",
			Port:  1549,
		},
	},
	"shivasound": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "shivasound",
			Port:  1549,
		},
	},
	"shockwave": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shockwave",
			Port:  1626,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "shockwave",
			Port:  1626,
		},
	},
	"shockwave2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shockwave2",
			Port:  1257,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "shockwave2",
			Port:  1257,
		},
	},
	"shofar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shofar",
			Port:  4105,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "shofar",
			Port:  4105,
		},
	},
	"shrinkwrap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "shrinkwrap",
			Port:  358,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "shrinkwrap",
			Port:  358,
		},
	},
	"sia-ctrl-plane": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sia-ctrl-plane",
			Port:  4787,
		},
	},
	"siam": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "siam",
			Port:  498,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "siam",
			Port:  498,
		},
	},
	"sicct": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sicct",
			Port:  4742,
		},
	},
	"sicct-sdp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "sicct-sdp",
			Port:  4742,
		},
	},
	"siebel-ns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "siebel-ns",
			Port:  2320,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "siebel-ns",
			Port:  2320,
		},
	},
	"siemensgsm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "siemensgsm",
			Port:  28240,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "siemensgsm",
			Port:  28240,
		},
	},
	"sieve": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sieve",
			Port:  4190,
		},
	},
	"sieve-filter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sieve-filter",
			Port:  2000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sieve-filter",
			Port:  2000,
		},
	},
	"sift-uft": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sift-uft",
			Port:  608,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sift-uft",
			Port:  608,
		},
	},
	"sigma-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sigma-port",
			Port:  3614,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sigma-port",
			Port:  3614,
		},
	},
	"signacert-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "signacert-agent",
			Port:  5032,
		},
	},
	"signal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "signal",
			Port:  2974,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "signal",
			Port:  2974,
		},
	},
	"signet-ctf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "signet-ctf",
			Port:  2733,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "signet-ctf",
			Port:  2733,
		},
	},
	"siipat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "siipat",
			Port:  1733,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "siipat",
			Port:  1733,
		},
	},
	"silc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silc",
			Port:  706,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silc",
			Port:  706,
		},
	},
	"silhouette": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silhouette",
			Port:  7500,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silhouette",
			Port:  7500,
		},
	},
	"silkmeter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silkmeter",
			Port:  5461,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silkmeter",
			Port:  5461,
		},
	},
	"silkp1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silkp1",
			Port:  2829,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silkp1",
			Port:  2829,
		},
	},
	"silkp2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silkp2",
			Port:  2830,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silkp2",
			Port:  2830,
		},
	},
	"silkp3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silkp3",
			Port:  2831,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silkp3",
			Port:  2831,
		},
	},
	"silkp4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silkp4",
			Port:  2832,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silkp4",
			Port:  2832,
		},
	},
	"silverpeakcomm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silverpeakcomm",
			Port:  4164,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silverpeakcomm",
			Port:  4164,
		},
	},
	"silverpeakpeer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silverpeakpeer",
			Port:  4163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silverpeakpeer",
			Port:  4163,
		},
	},
	"silverplatter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "silverplatter",
			Port:  416,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "silverplatter",
			Port:  416,
		},
	},
	"sim-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sim-control",
			Port:  3110,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sim-control",
			Port:  3110,
		},
	},
	"simba-cs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simba-cs",
			Port:  1543,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simba-cs",
			Port:  1543,
		},
	},
	"simbaexpress": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simbaexpress",
			Port:  1583,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simbaexpress",
			Port:  1583,
		},
	},
	"simbaservices": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simbaservices",
			Port:  1599,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simbaservices",
			Port:  1599,
		},
	},
	"simco": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "simco",
			Port:  7626,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "simco",
			Port:  7626,
		},
	},
	"simctlp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simctlp",
			Port:  2857,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simctlp",
			Port:  2857,
		},
	},
	"simon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simon",
			Port:  4753,
		},
	},
	"simon-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "simon-disc",
			Port:  4753,
		},
	},
	"simp-all": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simp-all",
			Port:  1959,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simp-all",
			Port:  1959,
		},
	},
	"simple-push": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simple-push",
			Port:  3687,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simple-push",
			Port:  3687,
		},
	},
	"simple-push-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simple-push-s",
			Port:  3688,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simple-push-s",
			Port:  3688,
		},
	},
	"simple-tx-rx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simple-tx-rx",
			Port:  2257,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simple-tx-rx",
			Port:  2257,
		},
	},
	"simplement-tie": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simplement-tie",
			Port:  2756,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simplement-tie",
			Port:  2756,
		},
	},
	"simplifymedia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simplifymedia",
			Port:  8087,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simplifymedia",
			Port:  8087,
		},
	},
	"simslink": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "simslink",
			Port:  2676,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "simslink",
			Port:  2676,
		},
	},
	"sip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sip",
			Port:  5060,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sip",
			Port:  5060,
		},
	},
	"sip-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sip-tls",
			Port:  5061,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sip-tls",
			Port:  5061,
		},
	},
	"sis-emt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sis-emt",
			Port:  2545,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sis-emt",
			Port:  2545,
		},
	},
	"sitaradir": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sitaradir",
			Port:  2631,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sitaradir",
			Port:  2631,
		},
	},
	"sitaramgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sitaramgmt",
			Port:  2630,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sitaramgmt",
			Port:  2630,
		},
	},
	"sitaraserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sitaraserver",
			Port:  2629,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sitaraserver",
			Port:  2629,
		},
	},
	"sitewatch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sitewatch",
			Port:  3792,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sitewatch",
			Port:  3792,
		},
	},
	"six-degrees": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "six-degrees",
			Port:  3611,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "six-degrees",
			Port:  3611,
		},
	},
	"sixnetudr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sixnetudr",
			Port:  1658,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sixnetudr",
			Port:  1658,
		},
	},
	"sixtrak": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sixtrak",
			Port:  1594,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sixtrak",
			Port:  1594,
		},
	},
	"sixxsconfig": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sixxsconfig",
			Port:  3874,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sixxsconfig",
			Port:  3874,
		},
	},
	"skip-cert-recv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "skip-cert-recv",
			Port:  6455,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "skip-cert-recv",
			Port:  6455,
		},
	},
	"skip-cert-send": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "skip-cert-send",
			Port:  6456,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "skip-cert-send",
			Port:  6456,
		},
	},
	"skip-mc-gikreq": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "skip-mc-gikreq",
			Port:  1660,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "skip-mc-gikreq",
			Port:  1660,
		},
	},
	"skkserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "skkserv",
			Port:  1178,
		},
	},
	"skronk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "skronk",
			Port:  460,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "skronk",
			Port:  460,
		},
	},
	"sky-transport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sky-transport",
			Port:  3556,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sky-transport",
			Port:  3556,
		},
	},
	"skytelnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "skytelnet",
			Port:  1618,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "skytelnet",
			Port:  1618,
		},
	},
	"slc-ctrlrloops": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slc-ctrlrloops",
			Port:  2827,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slc-ctrlrloops",
			Port:  2827,
		},
	},
	"slc-systemlog": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slc-systemlog",
			Port:  2826,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slc-systemlog",
			Port:  2826,
		},
	},
	"slim-devices": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slim-devices",
			Port:  3483,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slim-devices",
			Port:  3483,
		},
	},
	"slingshot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slingshot",
			Port:  1705,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slingshot",
			Port:  1705,
		},
	},
	"slinkysearch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slinkysearch",
			Port:  1225,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slinkysearch",
			Port:  1225,
		},
	},
	"slinterbase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slinterbase",
			Port:  3065,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slinterbase",
			Port:  3065,
		},
	},
	"slm-api": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slm-api",
			Port:  1606,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slm-api",
			Port:  1606,
		},
	},
	"slp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slp",
			Port:  1605,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slp",
			Port:  1605,
		},
	},
	"slp-notify": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slp-notify",
			Port:  1847,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slp-notify",
			Port:  1847,
		},
	},
	"slscc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slscc",
			Port:  4408,
		},
	},
	"slslavemon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slslavemon",
			Port:  3102,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slslavemon",
			Port:  3102,
		},
	},
	"slush": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "slush",
			Port:  1966,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "slush",
			Port:  1966,
		},
	},
	"sm-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "sm-disc",
			Port:  4174,
		},
	},
	"sm-pas-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sm-pas-1",
			Port:  2938,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sm-pas-1",
			Port:  2938,
		},
	},
	"sm-pas-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sm-pas-2",
			Port:  2939,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sm-pas-2",
			Port:  2939,
		},
	},
	"sm-pas-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sm-pas-3",
			Port:  2940,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sm-pas-3",
			Port:  2940,
		},
	},
	"sm-pas-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sm-pas-4",
			Port:  2941,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sm-pas-4",
			Port:  2941,
		},
	},
	"sm-pas-5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sm-pas-5",
			Port:  2942,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sm-pas-5",
			Port:  2942,
		},
	},
	"sma-spw": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "sma-spw",
			Port:  9522,
		},
	},
	"smaclmgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smaclmgr",
			Port:  4660,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smaclmgr",
			Port:  4660,
		},
	},
	"smakynet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smakynet",
			Port:  122,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smakynet",
			Port:  122,
		},
	},
	"smap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smap",
			Port:  3731,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smap",
			Port:  3731,
		},
	},
	"smar-se-port1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smar-se-port1",
			Port:  4987,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smar-se-port1",
			Port:  4987,
		},
	},
	"smar-se-port2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smar-se-port2",
			Port:  4988,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smar-se-port2",
			Port:  4988,
		},
	},
	"smart-diagnose": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smart-diagnose",
			Port:  2721,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smart-diagnose",
			Port:  2721,
		},
	},
	"smart-install": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smart-install",
			Port:  4786,
		},
	},
	"smart-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smart-lm",
			Port:  1608,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smart-lm",
			Port:  1608,
		},
	},
	"smartcard-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smartcard-port",
			Port:  3516,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smartcard-port",
			Port:  3516,
		},
	},
	"smartcard-tls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smartcard-tls",
			Port:  4116,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smartcard-tls",
			Port:  4116,
		},
	},
	"smartpackets": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smartpackets",
			Port:  3218,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smartpackets",
			Port:  3218,
		},
	},
	"smartsdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smartsdp",
			Port:  426,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smartsdp",
			Port:  426,
		},
	},
	"smauth-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smauth-port",
			Port:  3929,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smauth-port",
			Port:  3929,
		},
	},
	"smbdirect": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "smbdirect",
			Port:  5445,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "smbdirect",
			Port:  5445,
		},
	},
	"smc-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smc-admin",
			Port:  6787,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smc-admin",
			Port:  6787,
		},
	},
	"smc-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smc-http",
			Port:  6788,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smc-http",
			Port:  6788,
		},
	},
	"smc-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smc-https",
			Port:  6789,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smc-https",
			Port:  6789,
		},
	},
	"smc-jmx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smc-jmx",
			Port:  6786,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smc-jmx",
			Port:  6786,
		},
	},
	"smcluster": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smcluster",
			Port:  4174,
		},
	},
	"smile": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smile",
			Port:  3670,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smile",
			Port:  3670,
		},
	},
	"smip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smip",
			Port:  7734,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smip",
			Port:  7734,
		},
	},
	"smntubootstrap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smntubootstrap",
			Port:  2613,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smntubootstrap",
			Port:  2613,
		},
	},
	"smpnameres": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "smpnameres",
			Port:  901,
		},
	},
	"smpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smpp",
			Port:  2775,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smpp",
			Port:  2775,
		},
	},
	"smpppd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smpppd",
			Port:  3185,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smpppd",
			Port:  3185,
		},
	},
	"smpte": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smpte",
			Port:  420,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smpte",
			Port:  420,
		},
	},
	"sms-chat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sms-chat",
			Port:  2703,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sms-chat",
			Port:  2703,
		},
	},
	"sms-rcinfo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sms-rcinfo",
			Port:  2701,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sms-rcinfo",
			Port:  2701,
		},
	},
	"sms-remctrl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sms-remctrl",
			Port:  2704,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sms-remctrl",
			Port:  2704,
		},
	},
	"sms-xfer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sms-xfer",
			Port:  2702,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sms-xfer",
			Port:  2702,
		},
	},
	"smsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smsd",
			Port:  596,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smsd",
			Port:  596,
		},
	},
	"smsp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smsp",
			Port:  413,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smsp",
			Port:  413,
		},
	},
	"smsqp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smsqp",
			Port:  11201,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smsqp",
			Port:  11201,
		},
	},
	"smtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smtp",
			Port:  25,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smtp",
			Port:  25,
		},
	},
	"smux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smux",
			Port:  199,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smux",
			Port:  199,
		},
	},
	"smwan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "smwan",
			Port:  3979,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "smwan",
			Port:  3979,
		},
	},
	"sna-cs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sna-cs",
			Port:  1553,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sna-cs",
			Port:  1553,
		},
	},
	"snac": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snac",
			Port:  3536,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snac",
			Port:  3536,
		},
	},
	"snagas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snagas",
			Port:  108,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snagas",
			Port:  108,
		},
	},
	"snap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snap",
			Port:  4752,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snap",
			Port:  4752,
		},
	},
	"snapd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snapd",
			Port:  2599,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snapd",
			Port:  2599,
		},
	},
	"snapenetio": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snapenetio",
			Port:  22000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snapenetio",
			Port:  22000,
		},
	},
	"snapp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snapp",
			Port:  2333,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snapp",
			Port:  2333,
		},
	},
	"snare": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snare",
			Port:  509,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snare",
			Port:  509,
		},
	},
	"snaresecure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snaresecure",
			Port:  1684,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snaresecure",
			Port:  1684,
		},
	},
	"sncp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sncp",
			Port:  7560,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sncp",
			Port:  7560,
		},
	},
	"snifferclient": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snifferclient",
			Port:  2452,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snifferclient",
			Port:  2452,
		},
	},
	"snifferdata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snifferdata",
			Port:  2892,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snifferdata",
			Port:  2892,
		},
	},
	"snifferserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snifferserver",
			Port:  2533,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snifferserver",
			Port:  2533,
		},
	},
	"snip-slave": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snip-slave",
			Port:  33656,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snip-slave",
			Port:  33656,
		},
	},
	"snmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snmp",
			Port:  161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snmp",
			Port:  161,
		},
	},
	"snmp-tcp-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snmp-tcp-port",
			Port:  1993,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snmp-tcp-port",
			Port:  1993,
		},
	},
	"snmpdtls": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "snmpdtls",
			Port:  10161,
		},
	},
	"snmpdtls-trap": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "snmpdtls-trap",
			Port:  10162,
		},
	},
	"snmpssh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snmpssh",
			Port:  5161,
		},
	},
	"snmpssh-trap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snmpssh-trap",
			Port:  5162,
		},
	},
	"snmptls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snmptls",
			Port:  10161,
		},
	},
	"snmptls-trap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snmptls-trap",
			Port:  10162,
		},
	},
	"snmptrap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snmptrap",
			Port:  162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snmptrap",
			Port:  162,
		},
	},
	"snpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "snpp",
			Port:  444,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "snpp",
			Port:  444,
		},
	},
	"sns-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sns-admin",
			Port:  2658,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sns-admin",
			Port:  2658,
		},
	},
	"sns-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sns-agent",
			Port:  5417,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sns-agent",
			Port:  5417,
		},
	},
	"sns-channels": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sns-channels",
			Port:  3380,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sns-channels",
			Port:  3380,
		},
	},
	"sns-dispatcher": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sns-dispatcher",
			Port:  2657,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sns-dispatcher",
			Port:  2657,
		},
	},
	"sns-gateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sns-gateway",
			Port:  5416,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sns-gateway",
			Port:  5416,
		},
	},
	"sns-protocol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sns-protocol",
			Port:  2409,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sns-protocol",
			Port:  2409,
		},
	},
	"sns-query": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sns-query",
			Port:  2659,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sns-query",
			Port:  2659,
		},
	},
	"sns-quote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sns-quote",
			Port:  1967,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sns-quote",
			Port:  1967,
		},
	},
	"snss": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "snss",
			Port:  11171,
		},
	},
	"sntlkeyssrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sntlkeyssrvr",
			Port:  9450,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sntlkeyssrvr",
			Port:  9450,
		},
	},
	"sntp-heartbeat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sntp-heartbeat",
			Port:  580,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sntp-heartbeat",
			Port:  580,
		},
	},
	"soagateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "soagateway",
			Port:  5250,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "soagateway",
			Port:  5250,
		},
	},
	"soap-beep": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "soap-beep",
			Port:  605,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "soap-beep",
			Port:  605,
		},
	},
	"soap-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "soap-http",
			Port:  7627,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "soap-http",
			Port:  7627,
		},
	},
	"socalia": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "socalia",
			Port:  5100,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "socalia",
			Port:  5100,
		},
	},
	"social-alarm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "social-alarm",
			Port:  5146,
		},
	},
	"socks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "socks",
			Port:  1080,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "socks",
			Port:  1080,
		},
	},
	"socorfs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "socorfs",
			Port:  3379,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "socorfs",
			Port:  3379,
		},
	},
	"socp-c": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "socp-c",
			Port:  4882,
		},
	},
	"socp-t": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "socp-t",
			Port:  4881,
		},
	},
	"softaudit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "softaudit",
			Port:  3419,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "softaudit",
			Port:  3419,
		},
	},
	"softcm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "softcm",
			Port:  6110,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "softcm",
			Port:  6110,
		},
	},
	"softdataphone": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "softdataphone",
			Port:  1621,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "softdataphone",
			Port:  1621,
		},
	},
	"softpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "softpc",
			Port:  215,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "softpc",
			Port:  215,
		},
	},
	"softrack-meter": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "softrack-meter",
			Port:  3884,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "softrack-meter",
			Port:  3884,
		},
	},
	"solaris-audit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "solaris-audit",
			Port:  16162,
		},
	},
	"solera-epmap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "solera-epmap",
			Port:  2132,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "solera-epmap",
			Port:  2132,
		},
	},
	"solera-lpn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "solera-lpn",
			Port:  4738,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "solera-lpn",
			Port:  4738,
		},
	},
	"solid-e-engine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "solid-e-engine",
			Port:  1964,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "solid-e-engine",
			Port:  1964,
		},
	},
	"solid-mux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "solid-mux",
			Port:  1029,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "solid-mux",
			Port:  1029,
		},
	},
	"solve": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "solve",
			Port:  2636,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "solve",
			Port:  2636,
		},
	},
	"sonar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sonar",
			Port:  572,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sonar",
			Port:  572,
		},
	},
	"sonardata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sonardata",
			Port:  2863,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sonardata",
			Port:  2863,
		},
	},
	"soniqsync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "soniqsync",
			Port:  3803,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "soniqsync",
			Port:  3803,
		},
	},
	"sonus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sonus",
			Port:  2653,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sonus",
			Port:  2653,
		},
	},
	"sonus-logging": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sonus-logging",
			Port:  2290,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sonus-logging",
			Port:  2290,
		},
	},
	"sonuscallsig": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sonuscallsig",
			Port:  2569,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sonuscallsig",
			Port:  2569,
		},
	},
	"sophia-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sophia-lm",
			Port:  1408,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sophia-lm",
			Port:  1408,
		},
	},
	"sops": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sops",
			Port:  3944,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sops",
			Port:  3944,
		},
	},
	"sor-update": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sor-update",
			Port:  3922,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sor-update",
			Port:  3922,
		},
	},
	"sos": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sos",
			Port:  3838,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sos",
			Port:  3838,
		},
	},
	"sossd-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sossd-agent",
			Port:  7982,
		},
	},
	"sossd-collect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sossd-collect",
			Port:  7981,
		},
	},
	"sossd-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "sossd-disc",
			Port:  7982,
		},
	},
	"sossecollector": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sossecollector",
			Port:  3166,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sossecollector",
			Port:  3166,
		},
	},
	"soundsvirtual": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "soundsvirtual",
			Port:  17185,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "soundsvirtual",
			Port:  17185,
		},
	},
	"sp-remotetablet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sp-remotetablet",
			Port:  46998,
		},
	},
	"spamtrap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spamtrap",
			Port:  2568,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spamtrap",
			Port:  2568,
		},
	},
	"spandataport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spandataport",
			Port:  3193,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spandataport",
			Port:  3193,
		},
	},
	"spc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spc",
			Port:  6111,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spc",
			Port:  6111,
		},
	},
	"spcsdlobby": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spcsdlobby",
			Port:  2888,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spcsdlobby",
			Port:  2888,
		},
	},
	"spdp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "spdp",
			Port:  5794,
		},
	},
	"spdy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spdy",
			Port:  6121,
		},
	},
	"spearway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spearway",
			Port:  2440,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spearway",
			Port:  2440,
		},
	},
	"spectardata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spectardata",
			Port:  3834,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spectardata",
			Port:  3834,
		},
	},
	"spectardb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spectardb",
			Port:  3835,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spectardb",
			Port:  3835,
		},
	},
	"spectraport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spectraport",
			Port:  3851,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spectraport",
			Port:  3851,
		},
	},
	"speedtrace": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "speedtrace",
			Port:  33334,
		},
	},
	"speedtrace-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "speedtrace-disc",
			Port:  33334,
		},
	},
	"sphinxapi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sphinxapi",
			Port:  9312,
		},
	},
	"sphinxql": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sphinxql",
			Port:  9306,
		},
	},
	"spice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spice",
			Port:  1923,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spice",
			Port:  1923,
		},
	},
	"spike": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spike",
			Port:  4683,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spike",
			Port:  4683,
		},
	},
	"spiral-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spiral-admin",
			Port:  3438,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spiral-admin",
			Port:  3438,
		},
	},
	"splitlock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "splitlock",
			Port:  3606,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "splitlock",
			Port:  3606,
		},
	},
	"splitlock-gw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "splitlock-gw",
			Port:  3647,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "splitlock-gw",
			Port:  3647,
		},
	},
	"spmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spmp",
			Port:  656,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spmp",
			Port:  656,
		},
	},
	"spock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spock",
			Port:  2507,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spock",
			Port:  2507,
		},
	},
	"spocp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spocp",
			Port:  4751,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spocp",
			Port:  4751,
		},
	},
	"spramsca": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spramsca",
			Port:  5769,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spramsca",
			Port:  5769,
		},
	},
	"spramsd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spramsd",
			Port:  5770,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spramsd",
			Port:  5770,
		},
	},
	"sps-tunnel": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sps-tunnel",
			Port:  2876,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sps-tunnel",
			Port:  2876,
		},
	},
	"spsc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spsc",
			Port:  478,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spsc",
			Port:  478,
		},
	},
	"spss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spss",
			Port:  5443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spss",
			Port:  5443,
		},
	},
	"spss-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spss-lm",
			Port:  1759,
		},
	},
	"spt-automation": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spt-automation",
			Port:  5814,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spt-automation",
			Port:  5814,
		},
	},
	"spugna": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spugna",
			Port:  3807,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spugna",
			Port:  3807,
		},
	},
	"spw-dialer": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spw-dialer",
			Port:  3796,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spw-dialer",
			Port:  3796,
		},
	},
	"spw-dnspreload": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spw-dnspreload",
			Port:  3849,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spw-dnspreload",
			Port:  3849,
		},
	},
	"spytechphone": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "spytechphone",
			Port:  8192,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "spytechphone",
			Port:  8192,
		},
	},
	"sqdr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sqdr",
			Port:  2728,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sqdr",
			Port:  2728,
		},
	},
	"sql*net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sql*net",
			Port:  66,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sql*net",
			Port:  66,
		},
	},
	"sql-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sql-net",
			Port:  150,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sql-net",
			Port:  150,
		},
	},
	"sqlexec": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sqlexec",
			Port:  9088,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sqlexec",
			Port:  9088,
		},
	},
	"sqlexec-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sqlexec-ssl",
			Port:  9089,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sqlexec-ssl",
			Port:  9089,
		},
	},
	"sqlserv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sqlserv",
			Port:  118,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sqlserv",
			Port:  118,
		},
	},
	"sqlserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sqlserver",
			Port:  4430,
		},
	},
	"sqlsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sqlsrv",
			Port:  156,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sqlsrv",
			Port:  156,
		},
	},
	"squid": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "squid",
			Port:  3128,
		},
	},
	"src": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "src",
			Port:  200,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "src",
			Port:  200,
		},
	},
	"srcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "srcp",
			Port:  4303,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "srcp",
			Port:  4303,
		},
	},
	"srdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "srdp",
			Port:  3942,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "srdp",
			Port:  3942,
		},
	},
	"srmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "srmp",
			Port:  193,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "srmp",
			Port:  193,
		},
	},
	"srp-feedback": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "srp-feedback",
			Port:  2737,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "srp-feedback",
			Port:  2737,
		},
	},
	"srssend": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "srssend",
			Port:  362,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "srssend",
			Port:  362,
		},
	},
	"sruth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sruth",
			Port:  38800,
		},
	},
	"srvc_registry": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "srvc_registry",
			Port:  3018,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "srvc_registry",
			Port:  3018,
		},
	},
	"ss-idi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ss-idi",
			Port:  20013,
		},
	},
	"ss-idi-disc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ss-idi-disc",
			Port:  20012,
		},
	},
	"ss7ns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ss7ns",
			Port:  477,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ss7ns",
			Port:  477,
		},
	},
	"ssad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssad",
			Port:  4750,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssad",
			Port:  4750,
		},
	},
	"ssc-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssc-agent",
			Port:  2967,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssc-agent",
			Port:  2967,
		},
	},
	"sscan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sscan",
			Port:  3853,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sscan",
			Port:  3853,
		},
	},
	"ssdispatch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssdispatch",
			Port:  3430,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssdispatch",
			Port:  3430,
		},
	},
	"ssdp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssdp",
			Port:  1900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssdp",
			Port:  1900,
		},
	},
	"ssdtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssdtp",
			Port:  6071,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssdtp",
			Port:  6071,
		},
	},
	"sse-app-config": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sse-app-config",
			Port:  3852,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sse-app-config",
			Port:  3852,
		},
	},
	"ssh": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "ssh",
			Port:  22,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssh",
			Port:  22,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssh",
			Port:  22,
		},
	},
	"ssh-mgmt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssh-mgmt",
			Port:  17235,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssh-mgmt",
			Port:  17235,
		},
	},
	"sshell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sshell",
			Port:  614,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sshell",
			Port:  614,
		},
	},
	"sslp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sslp",
			Port:  1750,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sslp",
			Port:  1750,
		},
	},
	"ssm-cssps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssm-cssps",
			Port:  2478,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssm-cssps",
			Port:  2478,
		},
	},
	"ssm-cvs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssm-cvs",
			Port:  2477,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssm-cvs",
			Port:  2477,
		},
	},
	"ssm-els": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssm-els",
			Port:  2479,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssm-els",
			Port:  2479,
		},
	},
	"ssmc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssmc",
			Port:  2187,
		},
	},
	"ssmd": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "ssmd",
			Port:  2187,
		},
	},
	"ssmpp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssmpp",
			Port:  3550,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssmpp",
			Port:  3550,
		},
	},
	"sso-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sso-control",
			Port:  2711,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sso-control",
			Port:  2711,
		},
	},
	"sso-service": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sso-service",
			Port:  2710,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sso-service",
			Port:  2710,
		},
	},
	"ssowatch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssowatch",
			Port:  3644,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssowatch",
			Port:  3644,
		},
	},
	"ssp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssp",
			Port:  3249,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssp",
			Port:  3249,
		},
	},
	"ssp-client": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssp-client",
			Port:  7801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssp-client",
			Port:  7801,
		},
	},
	"ssql": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssql",
			Port:  3352,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssql",
			Port:  3352,
		},
	},
	"ssr-servermgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssr-servermgr",
			Port:  45966,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssr-servermgr",
			Port:  45966,
		},
	},
	"ssrip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssrip",
			Port:  3318,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssrip",
			Port:  3318,
		},
	},
	"ssslic-mgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssslic-mgr",
			Port:  1203,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssslic-mgr",
			Port:  1203,
		},
	},
	"ssslog-mgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ssslog-mgr",
			Port:  1204,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ssslog-mgr",
			Port:  1204,
		},
	},
	"sst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sst",
			Port:  266,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sst",
			Port:  266,
		},
	},
	"sstp-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sstp-1",
			Port:  7743,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sstp-1",
			Port:  7743,
		},
	},
	"sstp-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sstp-2",
			Port:  9801,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sstp-2",
			Port:  9801,
		},
	},
	"sstsys-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sstsys-lm",
			Port:  1692,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sstsys-lm",
			Port:  1692,
		},
	},
	"stanag-5066": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stanag-5066",
			Port:  5066,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stanag-5066",
			Port:  5066,
		},
	},
	"starbot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "starbot",
			Port:  2838,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "starbot",
			Port:  2838,
		},
	},
	"starfish": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "starfish",
			Port:  3981,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "starfish",
			Port:  3981,
		},
	},
	"stargatealerts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stargatealerts",
			Port:  1654,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stargatealerts",
			Port:  1654,
		},
	},
	"starquiz-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "starquiz-port",
			Port:  3526,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "starquiz-port",
			Port:  3526,
		},
	},
	"stars": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stars",
			Port:  4131,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stars",
			Port:  4131,
		},
	},
	"starschool": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "starschool",
			Port:  2270,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "starschool",
			Port:  2270,
		},
	},
	"start-network": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "start-network",
			Port:  3615,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "start-network",
			Port:  3615,
		},
	},
	"startron": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "startron",
			Port:  1057,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "startron",
			Port:  1057,
		},
	},
	"stat-cc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stat-cc",
			Port:  4158,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stat-cc",
			Port:  4158,
		},
	},
	"stat-results": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stat-results",
			Port:  4156,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stat-results",
			Port:  4156,
		},
	},
	"stat-scanner": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stat-scanner",
			Port:  4157,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stat-scanner",
			Port:  4157,
		},
	},
	"statsci1-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "statsci1-lm",
			Port:  6144,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "statsci1-lm",
			Port:  6144,
		},
	},
	"statsci2-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "statsci2-lm",
			Port:  6145,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "statsci2-lm",
			Port:  6145,
		},
	},
	"statsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "statsrv",
			Port:  133,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "statsrv",
			Port:  133,
		},
	},
	"statusd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "statusd",
			Port:  5414,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "statusd",
			Port:  5414,
		},
	},
	"stdptc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stdptc",
			Port:  2154,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stdptc",
			Port:  2154,
		},
	},
	"ste-smsc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ste-smsc",
			Port:  1836,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ste-smsc",
			Port:  1836,
		},
	},
	"stgxfws": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stgxfws",
			Port:  1226,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stgxfws",
			Port:  1226,
		},
	},
	"sti-envision": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sti-envision",
			Port:  1312,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sti-envision",
			Port:  1312,
		},
	},
	"stm_pproc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stm_pproc",
			Port:  3080,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stm_pproc",
			Port:  3080,
		},
	},
	"stmf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stmf",
			Port:  501,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stmf",
			Port:  501,
		},
	},
	"stone-design-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stone-design-1",
			Port:  1492,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stone-design-1",
			Port:  1492,
		},
	},
	"stonefalls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stonefalls",
			Port:  2986,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stonefalls",
			Port:  2986,
		},
	},
	"storman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "storman",
			Port:  4178,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "storman",
			Port:  4178,
		},
	},
	"storview": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "storview",
			Port:  9293,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "storview",
			Port:  9293,
		},
	},
	"streamcomm-ds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "streamcomm-ds",
			Port:  9612,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "streamcomm-ds",
			Port:  9612,
		},
	},
	"street-stream": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "street-stream",
			Port:  1736,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "street-stream",
			Port:  1736,
		},
	},
	"streetperfect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "streetperfect",
			Port:  1330,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "streetperfect",
			Port:  1330,
		},
	},
	"streettalk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "streettalk",
			Port:  566,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "streettalk",
			Port:  566,
		},
	},
	"stresstester": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stresstester",
			Port:  5397,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stresstester",
			Port:  5397,
		},
	},
	"strexec-d": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "strexec-d",
			Port:  5026,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "strexec-d",
			Port:  5026,
		},
	},
	"strexec-s": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "strexec-s",
			Port:  5027,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "strexec-s",
			Port:  5027,
		},
	},
	"stryker-com": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stryker-com",
			Port:  3854,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stryker-com",
			Port:  3854,
		},
	},
	"stss": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stss",
			Port:  3090,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stss",
			Port:  3090,
		},
	},
	"stt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stt",
			Port:  1607,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stt",
			Port:  1607,
		},
	},
	"stun": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stun",
			Port:  3478,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stun",
			Port:  3478,
		},
	},
	"stun-p1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stun-p1",
			Port:  1990,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stun-p1",
			Port:  1990,
		},
	},
	"stun-p2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stun-p2",
			Port:  1991,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stun-p2",
			Port:  1991,
		},
	},
	"stun-p3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stun-p3",
			Port:  1992,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stun-p3",
			Port:  1992,
		},
	},
	"stun-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stun-port",
			Port:  1994,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stun-port",
			Port:  1994,
		},
	},
	"stuns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stuns",
			Port:  5349,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stuns",
			Port:  5349,
		},
	},
	"stvp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stvp",
			Port:  3158,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stvp",
			Port:  3158,
		},
	},
	"stx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "stx",
			Port:  527,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "stx",
			Port:  527,
		},
	},
	"su-mit-tg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "su-mit-tg",
			Port:  89,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "su-mit-tg",
			Port:  89,
		},
	},
	"sua": map[string]Service{
		"sctp": Service{
			Proto: "sctp",
			Name:  "sua",
			Port:  14001,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "sua",
			Port:  14001,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sua",
			Port:  14001,
		},
	},
	"submission": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "submission",
			Port:  587,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "submission",
			Port:  587,
		},
	},
	"submit": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "submit",
			Port:  773,
		},
	},
	"submitserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "submitserver",
			Port:  2028,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "submitserver",
			Port:  2028,
		},
	},
	"subntbcst_tftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "subntbcst_tftp",
			Port:  247,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "subntbcst_tftp",
			Port:  247,
		},
	},
	"sugp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sugp",
			Port:  1905,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sugp",
			Port:  1905,
		},
	},
	"suitcase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "suitcase",
			Port:  2903,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "suitcase",
			Port:  2903,
		},
	},
	"suitjd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "suitjd",
			Port:  3354,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "suitjd",
			Port:  3354,
		},
	},
	"sum": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sum",
			Port:  6551,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sum",
			Port:  6551,
		},
	},
	"sun-as-iiops": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-as-iiops",
			Port:  3708,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-as-iiops",
			Port:  3708,
		},
	},
	"sun-as-iiops-ca": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-as-iiops-ca",
			Port:  3808,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-as-iiops-ca",
			Port:  3808,
		},
	},
	"sun-as-jmxrmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-as-jmxrmi",
			Port:  8686,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-as-jmxrmi",
			Port:  8686,
		},
	},
	"sun-as-jpda": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-as-jpda",
			Port:  9191,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-as-jpda",
			Port:  9191,
		},
	},
	"sun-as-nodeagt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-as-nodeagt",
			Port:  4850,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-as-nodeagt",
			Port:  4850,
		},
	},
	"sun-dr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-dr",
			Port:  665,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-dr",
			Port:  665,
		},
	},
	"sun-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-lm",
			Port:  7588,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-lm",
			Port:  7588,
		},
	},
	"sun-mc-grp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-mc-grp",
			Port:  5306,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-mc-grp",
			Port:  5306,
		},
	},
	"sun-sea-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sea-port",
			Port:  16161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sea-port",
			Port:  16161,
		},
	},
	"sun-sr-admin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sr-admin",
			Port:  6489,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sr-admin",
			Port:  6489,
		},
	},
	"sun-sr-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sr-http",
			Port:  6480,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sr-http",
			Port:  6480,
		},
	},
	"sun-sr-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sr-https",
			Port:  6443,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sr-https",
			Port:  6443,
		},
	},
	"sun-sr-iiop": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sr-iiop",
			Port:  6485,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sr-iiop",
			Port:  6485,
		},
	},
	"sun-sr-iiop-aut": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sr-iiop-aut",
			Port:  6487,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sr-iiop-aut",
			Port:  6487,
		},
	},
	"sun-sr-iiops": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sr-iiops",
			Port:  6486,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sr-iiops",
			Port:  6486,
		},
	},
	"sun-sr-jms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sr-jms",
			Port:  6484,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sr-jms",
			Port:  6484,
		},
	},
	"sun-sr-jmx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-sr-jmx",
			Port:  6488,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-sr-jmx",
			Port:  6488,
		},
	},
	"sun-user-https": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sun-user-https",
			Port:  7677,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sun-user-https",
			Port:  7677,
		},
	},
	"suncacao-csa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "suncacao-csa",
			Port:  11164,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "suncacao-csa",
			Port:  11164,
		},
	},
	"suncacao-jmxmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "suncacao-jmxmp",
			Port:  11162,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "suncacao-jmxmp",
			Port:  11162,
		},
	},
	"suncacao-rmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "suncacao-rmi",
			Port:  11163,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "suncacao-rmi",
			Port:  11163,
		},
	},
	"suncacao-snmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "suncacao-snmp",
			Port:  11161,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "suncacao-snmp",
			Port:  11161,
		},
	},
	"suncacao-websvc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "suncacao-websvc",
			Port:  11165,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "suncacao-websvc",
			Port:  11165,
		},
	},
	"sunclustergeo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunclustergeo",
			Port:  2084,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunclustergeo",
			Port:  2084,
		},
	},
	"sunclustermgr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunclustermgr",
			Port:  1097,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunclustermgr",
			Port:  1097,
		},
	},
	"sunfm-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunfm-port",
			Port:  3934,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunfm-port",
			Port:  3934,
		},
	},
	"sunlps-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunlps-http",
			Port:  3816,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunlps-http",
			Port:  3816,
		},
	},
	"sunrpc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunrpc",
			Port:  111,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunrpc",
			Port:  111,
		},
	},
	"sunscalar-dns": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunscalar-dns",
			Port:  1870,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunscalar-dns",
			Port:  1870,
		},
	},
	"sunscalar-svc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunscalar-svc",
			Port:  1860,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunscalar-svc",
			Port:  1860,
		},
	},
	"sunwebadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunwebadmin",
			Port:  8800,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunwebadmin",
			Port:  8800,
		},
	},
	"sunwebadmins": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sunwebadmins",
			Port:  8989,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sunwebadmins",
			Port:  8989,
		},
	},
	"supdup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "supdup",
			Port:  95,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "supdup",
			Port:  95,
		},
	},
	"supercell": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "supercell",
			Port:  7967,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "supercell",
			Port:  7967,
		},
	},
	"supermon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "supermon",
			Port:  2709,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "supermon",
			Port:  2709,
		},
	},
	"supfiledbg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "supfiledbg",
			Port:  1127,
		},
	},
	"supfilesrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "supfilesrv",
			Port:  871,
		},
	},
	"support": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "support",
			Port:  1529,
		},
	},
	"sur-meas": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sur-meas",
			Port:  243,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sur-meas",
			Port:  243,
		},
	},
	"surebox": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "surebox",
			Port:  5453,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "surebox",
			Port:  5453,
		},
	},
	"surf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "surf",
			Port:  1010,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "surf",
			Port:  1010,
		},
	},
	"surfcontrolcpa": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "surfcontrolcpa",
			Port:  3909,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "surfcontrolcpa",
			Port:  3909,
		},
	},
	"surfpass": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "surfpass",
			Port:  5030,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "surfpass",
			Port:  5030,
		},
	},
	"surveyinst": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "surveyinst",
			Port:  3212,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "surveyinst",
			Port:  3212,
		},
	},
	"suucp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "suucp",
			Port:  4031,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "suucp",
			Port:  4031,
		},
	},
	"svbackup": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "svbackup",
			Port:  8405,
		},
	},
	"svcloud": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "svcloud",
			Port:  8404,
		},
	},
	"svdrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "svdrp",
			Port:  6419,
		},
	},
	"svn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "svn",
			Port:  3690,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "svn",
			Port:  3690,
		},
	},
	"svnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "svnet",
			Port:  3413,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "svnet",
			Port:  3413,
		},
	},
	"svnetworks": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "svnetworks",
			Port:  2973,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "svnetworks",
			Port:  2973,
		},
	},
	"svrloc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "svrloc",
			Port:  427,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "svrloc",
			Port:  427,
		},
	},
	"svs-omagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "svs-omagent",
			Port:  1625,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "svs-omagent",
			Port:  1625,
		},
	},
	"sw-orion": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sw-orion",
			Port:  17777,
		},
	},
	"swa-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swa-1",
			Port:  9023,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swa-1",
			Port:  9023,
		},
	},
	"swa-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swa-2",
			Port:  9024,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swa-2",
			Port:  9024,
		},
	},
	"swa-3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swa-3",
			Port:  9025,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swa-3",
			Port:  9025,
		},
	},
	"swa-4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swa-4",
			Port:  9026,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swa-4",
			Port:  9026,
		},
	},
	"swat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swat",
			Port:  901,
		},
	},
	"swdtp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swdtp",
			Port:  10104,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swdtp",
			Port:  10104,
		},
	},
	"swdtp-sv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swdtp-sv",
			Port:  10009,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swdtp-sv",
			Port:  10009,
		},
	},
	"sweetware-apps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sweetware-apps",
			Port:  1221,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sweetware-apps",
			Port:  1221,
		},
	},
	"swift-rvf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swift-rvf",
			Port:  97,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swift-rvf",
			Port:  97,
		},
	},
	"swiftnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swiftnet",
			Port:  1751,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swiftnet",
			Port:  1751,
		},
	},
	"swismgr1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swismgr1",
			Port:  6963,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swismgr1",
			Port:  6963,
		},
	},
	"swismgr2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swismgr2",
			Port:  6964,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swismgr2",
			Port:  6964,
		},
	},
	"swispol": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swispol",
			Port:  6966,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swispol",
			Port:  6966,
		},
	},
	"swistrap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swistrap",
			Port:  6965,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swistrap",
			Port:  6965,
		},
	},
	"swldy-sias": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swldy-sias",
			Port:  1250,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swldy-sias",
			Port:  1250,
		},
	},
	"swr-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swr-port",
			Port:  3491,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swr-port",
			Port:  3491,
		},
	},
	"swrmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swrmi",
			Port:  1866,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swrmi",
			Port:  1866,
		},
	},
	"swtp-port1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swtp-port1",
			Port:  9281,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swtp-port1",
			Port:  9281,
		},
	},
	"swtp-port2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swtp-port2",
			Port:  9282,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swtp-port2",
			Port:  9282,
		},
	},
	"swx-gate": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swx-gate",
			Port:  4538,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swx-gate",
			Port:  4538,
		},
	},
	"swxadmin": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "swxadmin",
			Port:  5043,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "swxadmin",
			Port:  5043,
		},
	},
	"sxmp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sxmp",
			Port:  3273,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sxmp",
			Port:  3273,
		},
	},
	"sxuptp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sxuptp",
			Port:  19540,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sxuptp",
			Port:  19540,
		},
	},
	"syam-agent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "syam-agent",
			Port:  3894,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "syam-agent",
			Port:  3894,
		},
	},
	"syam-smc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "syam-smc",
			Port:  3895,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "syam-smc",
			Port:  3895,
		},
	},
	"syam-webserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "syam-webserver",
			Port:  3930,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "syam-webserver",
			Port:  3930,
		},
	},
	"sybase-sqlany": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sybase-sqlany",
			Port:  1498,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sybase-sqlany",
			Port:  1498,
		},
	},
	"sybaseanywhere": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sybaseanywhere",
			Port:  2638,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sybaseanywhere",
			Port:  2638,
		},
	},
	"sybasedbsynch": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sybasedbsynch",
			Port:  2439,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sybasedbsynch",
			Port:  2439,
		},
	},
	"sybasesrvmon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sybasesrvmon",
			Port:  4950,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sybasesrvmon",
			Port:  4950,
		},
	},
	"sychrond": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sychrond",
			Port:  3723,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sychrond",
			Port:  3723,
		},
	},
	"symantec-sfdb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "symantec-sfdb",
			Port:  5629,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "symantec-sfdb",
			Port:  5629,
		},
	},
	"symantec-sim": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "symantec-sim",
			Port:  3547,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "symantec-sim",
			Port:  3547,
		},
	},
	"symb-sb-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "symb-sb-port",
			Port:  3923,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "symb-sb-port",
			Port:  3923,
		},
	},
	"symplex": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "symplex",
			Port:  1507,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "symplex",
			Port:  1507,
		},
	},
	"synapse": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synapse",
			Port:  2880,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synapse",
			Port:  2880,
		},
	},
	"synapse-nhttp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synapse-nhttp",
			Port:  8280,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synapse-nhttp",
			Port:  8280,
		},
	},
	"synapse-nhttps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synapse-nhttps",
			Port:  8243,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synapse-nhttps",
			Port:  8243,
		},
	},
	"synapsis-edge": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synapsis-edge",
			Port:  5008,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synapsis-edge",
			Port:  5008,
		},
	},
	"sync-em7": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sync-em7",
			Port:  7707,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sync-em7",
			Port:  7707,
		},
	},
	"synchromesh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synchromesh",
			Port:  4548,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synchromesh",
			Port:  4548,
		},
	},
	"synchronet-db": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synchronet-db",
			Port:  6100,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synchronet-db",
			Port:  6100,
		},
	},
	"synchronet-rtc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synchronet-rtc",
			Port:  6101,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synchronet-rtc",
			Port:  6101,
		},
	},
	"synchronet-upd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synchronet-upd",
			Port:  6102,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synchronet-upd",
			Port:  6102,
		},
	},
	"synchronite": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synchronite",
			Port:  4106,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synchronite",
			Port:  4106,
		},
	},
	"syncserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "syncserver",
			Port:  2647,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "syncserver",
			Port:  2647,
		},
	},
	"syncserverssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "syncserverssl",
			Port:  2679,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "syncserverssl",
			Port:  2679,
		},
	},
	"synctest": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synctest",
			Port:  45045,
		},
	},
	"synel-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synel-data",
			Port:  3734,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synel-data",
			Port:  3734,
		},
	},
	"synoptics-trap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synoptics-trap",
			Port:  412,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synoptics-trap",
			Port:  412,
		},
	},
	"synotics-broker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synotics-broker",
			Port:  392,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synotics-broker",
			Port:  392,
		},
	},
	"synotics-relay": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "synotics-relay",
			Port:  391,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "synotics-relay",
			Port:  391,
		},
	},
	"sype-transport": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sype-transport",
			Port:  9911,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sype-transport",
			Port:  9911,
		},
	},
	"syscomlan": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "syscomlan",
			Port:  1065,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "syscomlan",
			Port:  1065,
		},
	},
	"syserverremote": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "syserverremote",
			Port:  6418,
		},
	},
	"sysinfo-sp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sysinfo-sp",
			Port:  11967,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sysinfo-sp",
			Port:  11967,
		},
	},
	"syslog": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "syslog",
			Port:  514,
		},
	},
	"syslog-conn": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "syslog-conn",
			Port:  601,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "syslog-conn",
			Port:  601,
		},
	},
	"syslog-tls": map[string]Service{
		"dccp": Service{
			Proto: "dccp",
			Name:  "syslog-tls",
			Port:  6514,
		},
		"tcp": Service{
			Proto: "tcp",
			Name:  "syslog-tls",
			Port:  6514,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "syslog-tls",
			Port:  6514,
		},
	},
	"sysopt": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sysopt",
			Port:  3281,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sysopt",
			Port:  3281,
		},
	},
	"sysorb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sysorb",
			Port:  3241,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sysorb",
			Port:  3241,
		},
	},
	"sysrqd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sysrqd",
			Port:  4094,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sysrqd",
			Port:  4094,
		},
	},
	"sysscanner": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "sysscanner",
			Port:  3251,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "sysscanner",
			Port:  3251,
		},
	},
	"systat": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "systat",
			Port:  11,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "systat",
			Port:  11,
		},
	},
	"system-monitor": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "system-monitor",
			Port:  2609,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "system-monitor",
			Port:  2609,
		},
	},
	"systemics-sox": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "systemics-sox",
			Port:  5406,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "systemics-sox",
			Port:  5406,
		},
	},
	"t1-e1-over-ip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "t1-e1-over-ip",
			Port:  3175,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "t1-e1-over-ip",
			Port:  3175,
		},
	},
	"t128-gateway": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "t128-gateway",
			Port:  1627,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "t128-gateway",
			Port:  1627,
		},
	},
	"t1distproc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "t1distproc",
			Port:  1274,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "t1distproc",
			Port:  1274,
		},
	},
	"t1distproc60": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "t1distproc60",
			Port:  32249,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "t1distproc60",
			Port:  32249,
		},
	},
	"t2-brm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "t2-brm",
			Port:  7933,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "t2-brm",
			Port:  7933,
		},
	},
	"t2-drm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "t2-drm",
			Port:  7932,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "t2-drm",
			Port:  7932,
		},
	},
	"t5-straton": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "t5-straton",
			Port:  11173,
		},
	},
	"tabula": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tabula",
			Port:  1437,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tabula",
			Port:  1437,
		},
	},
	"tacacs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tacacs",
			Port:  49,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tacacs",
			Port:  49,
		},
	},
	"tacacs-ds": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tacacs-ds",
			Port:  65,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tacacs-ds",
			Port:  65,
		},
	},
	"tacnews": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tacnews",
			Port:  98,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tacnews",
			Port:  98,
		},
	},
	"tacticalauth": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tacticalauth",
			Port:  2392,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tacticalauth",
			Port:  2392,
		},
	},
	"taep-as-svc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "taep-as-svc",
			Port:  5111,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "taep-as-svc",
			Port:  5111,
		},
	},
	"tag-pm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tag-pm",
			Port:  5073,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tag-pm",
			Port:  5073,
		},
	},
	"tag-ups-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tag-ups-1",
			Port:  3573,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tag-ups-1",
			Port:  3573,
		},
	},
	"taiclock": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "taiclock",
			Port:  4014,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "taiclock",
			Port:  4014,
		},
	},
	"tal-pod": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tal-pod",
			Port:  6149,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tal-pod",
			Port:  6149,
		},
	},
	"talarian-mcast1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talarian-mcast1",
			Port:  4015,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talarian-mcast1",
			Port:  4015,
		},
	},
	"talarian-mcast2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talarian-mcast2",
			Port:  4016,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talarian-mcast2",
			Port:  4016,
		},
	},
	"talarian-mcast3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talarian-mcast3",
			Port:  4017,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talarian-mcast3",
			Port:  4017,
		},
	},
	"talarian-mcast4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talarian-mcast4",
			Port:  4018,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talarian-mcast4",
			Port:  4018,
		},
	},
	"talarian-mcast5": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talarian-mcast5",
			Port:  4019,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talarian-mcast5",
			Port:  4019,
		},
	},
	"talarian-mqs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talarian-mqs",
			Port:  2493,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talarian-mqs",
			Port:  2493,
		},
	},
	"talarian-tcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talarian-tcp",
			Port:  5101,
		},
	},
	"talarian-udp": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "talarian-udp",
			Port:  5101,
		},
	},
	"taligent-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "taligent-lm",
			Port:  1475,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "taligent-lm",
			Port:  1475,
		},
	},
	"talikaserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talikaserver",
			Port:  22763,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talikaserver",
			Port:  22763,
		},
	},
	"talk": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talk",
			Port:  517,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talk",
			Port:  517,
		},
	},
	"talnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talnet",
			Port:  1838,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talnet",
			Port:  1838,
		},
	},
	"talon-disc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talon-disc",
			Port:  7011,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talon-disc",
			Port:  7011,
		},
	},
	"talon-engine": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talon-engine",
			Port:  7012,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talon-engine",
			Port:  7012,
		},
	},
	"talon-webserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "talon-webserver",
			Port:  7015,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "talon-webserver",
			Port:  7015,
		},
	},
	"tambora": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tambora",
			Port:  9020,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tambora",
			Port:  9020,
		},
	},
	"tams": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tams",
			Port:  2726,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tams",
			Port:  2726,
		},
	},
	"tapestry": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tapestry",
			Port:  1922,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tapestry",
			Port:  1922,
		},
	},
	"tapeware": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tapeware",
			Port:  3817,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tapeware",
			Port:  3817,
		},
	},
	"tappi-boxnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tappi-boxnet",
			Port:  2306,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tappi-boxnet",
			Port:  2306,
		},
	},
	"tarantella": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tarantella",
			Port:  3144,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tarantella",
			Port:  3144,
		},
	},
	"targus-getdata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "targus-getdata",
			Port:  5200,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "targus-getdata",
			Port:  5200,
		},
	},
	"targus-getdata1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "targus-getdata1",
			Port:  5201,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "targus-getdata1",
			Port:  5201,
		},
	},
	"targus-getdata2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "targus-getdata2",
			Port:  5202,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "targus-getdata2",
			Port:  5202,
		},
	},
	"targus-getdata3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "targus-getdata3",
			Port:  5203,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "targus-getdata3",
			Port:  5203,
		},
	},
	"taserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "taserver",
			Port:  3552,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "taserver",
			Port:  3552,
		},
	},
	"taskman-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "taskman-port",
			Port:  2470,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "taskman-port",
			Port:  2470,
		},
	},
	"taskmaster2000": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "taskmaster2000",
			Port:  2402,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "taskmaster2000",
			Port:  2402,
		},
	},
	"tasp-net": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tasp-net",
			Port:  25900,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tasp-net",
			Port:  25900,
		},
	},
	"taurus-wh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "taurus-wh",
			Port:  1610,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "taurus-wh",
			Port:  1610,
		},
	},
	"tbrpf": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tbrpf",
			Port:  712,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tbrpf",
			Port:  712,
		},
	},
	"tcc-http": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcc-http",
			Port:  24680,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcc-http",
			Port:  24680,
		},
	},
	"tcim-control": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcim-control",
			Port:  2729,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcim-control",
			Port:  2729,
		},
	},
	"tclprodebugger": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tclprodebugger",
			Port:  2576,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tclprodebugger",
			Port:  2576,
		},
	},
	"tcoaddressbook": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcoaddressbook",
			Port:  1977,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcoaddressbook",
			Port:  1977,
		},
	},
	"tcoflashagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcoflashagent",
			Port:  1975,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcoflashagent",
			Port:  1975,
		},
	},
	"tcoregagent": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcoregagent",
			Port:  1976,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcoregagent",
			Port:  1976,
		},
	},
	"tcp-id-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcp-id-port",
			Port:  1999,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcp-id-port",
			Port:  1999,
		},
	},
	"tcpdataserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcpdataserver",
			Port:  3805,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcpdataserver",
			Port:  3805,
		},
	},
	"tcpmux": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcpmux",
			Port:  1,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcpmux",
			Port:  1,
		},
	},
	"tcpnethaspsrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tcpnethaspsrv",
			Port:  475,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tcpnethaspsrv",
			Port:  475,
		},
	},
	"td-postman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "td-postman",
			Port:  1049,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "td-postman",
			Port:  1049,
		},
	},
	"td-replica": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "td-replica",
			Port:  268,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "td-replica",
			Port:  268,
		},
	},
	"td-service": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "td-service",
			Port:  267,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "td-service",
			Port:  267,
		},
	},
	"tdaccess": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tdaccess",
			Port:  2910,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tdaccess",
			Port:  2910,
		},
	},
	"tdmoip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tdmoip",
			Port:  2142,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tdmoip",
			Port:  2142,
		},
	},
	"tdp-suite": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tdp-suite",
			Port:  1814,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tdp-suite",
			Port:  1814,
		},
	},
	"teamcoherence": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "teamcoherence",
			Port:  9222,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "teamcoherence",
			Port:  9222,
		},
	},
	"tec5-sdctp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tec5-sdctp",
			Port:  9668,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tec5-sdctp",
			Port:  9668,
		},
	},
	"teedtap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "teedtap",
			Port:  559,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "teedtap",
			Port:  559,
		},
	},
	"tekpls": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tekpls",
			Port:  1946,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tekpls",
			Port:  1946,
		},
	},
	"telaconsole": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telaconsole",
			Port:  5428,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telaconsole",
			Port:  5428,
		},
	},
	"telefinder": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telefinder",
			Port:  1474,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telefinder",
			Port:  1474,
		},
	},
	"telelpathattack": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telelpathattack",
			Port:  5011,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telelpathattack",
			Port:  5011,
		},
	},
	"telelpathstart": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telelpathstart",
			Port:  5010,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telelpathstart",
			Port:  5010,
		},
	},
	"teleniumdaemon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "teleniumdaemon",
			Port:  2060,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "teleniumdaemon",
			Port:  2060,
		},
	},
	"telesis-licman": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telesis-licman",
			Port:  1380,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telesis-licman",
			Port:  1380,
		},
	},
	"telindus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telindus",
			Port:  1728,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telindus",
			Port:  1728,
		},
	},
	"tell": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "tell",
			Port:  754,
		},
	},
	"tellumat-nms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tellumat-nms",
			Port:  3549,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tellumat-nms",
			Port:  3549,
		},
	},
	"telnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telnet",
			Port:  23,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telnet",
			Port:  23,
		},
	},
	"telnetcpcd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telnetcpcd",
			Port:  3696,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telnetcpcd",
			Port:  3696,
		},
	},
	"telnets": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telnets",
			Port:  992,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telnets",
			Port:  992,
		},
	},
	"telops-lmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "telops-lmd",
			Port:  7491,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "telops-lmd",
			Port:  7491,
		},
	},
	"tempest-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tempest-port",
			Port:  11600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tempest-port",
			Port:  11600,
		},
	},
	"tempo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tempo",
			Port:  526,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tempo",
			Port:  526,
		},
	},
	"tenfold": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tenfold",
			Port:  658,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tenfold",
			Port:  658,
		},
	},
	"tentacle": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tentacle",
			Port:  41121,
		},
	},
	"terabase": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "terabase",
			Port:  4000,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "terabase",
			Port:  4000,
		},
	},
	"teradataordbms": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "teradataordbms",
			Port:  8002,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "teradataordbms",
			Port:  8002,
		},
	},
	"teredo": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "teredo",
			Port:  3544,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "teredo",
			Port:  3544,
		},
	},
	"terminaldb": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "terminaldb",
			Port:  2018,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "terminaldb",
			Port:  2008,
		},
	},
	"tesla-sys-msg": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tesla-sys-msg",
			Port:  7631,
		},
	},
	"tetrinet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tetrinet",
			Port:  31457,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tetrinet",
			Port:  31457,
		},
	},
	"texai": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "texai",
			Port:  5048,
		},
	},
	"texar": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "texar",
			Port:  333,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "texar",
			Port:  333,
		},
	},
	"tfido": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tfido",
			Port:  60177,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tfido",
			Port:  60177,
		},
	},
	"tftp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tftp",
			Port:  69,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tftp",
			Port:  69,
		},
	},
	"tftp-mcast": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tftp-mcast",
			Port:  1758,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tftp-mcast",
			Port:  1758,
		},
	},
	"tftps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tftps",
			Port:  3713,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tftps",
			Port:  3713,
		},
	},
	"tgcconnect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tgcconnect",
			Port:  4146,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tgcconnect",
			Port:  4146,
		},
	},
	"tgp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tgp",
			Port:  1223,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tgp",
			Port:  1223,
		},
	},
	"thermo-calc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "thermo-calc",
			Port:  6201,
		},
	},
	"theta-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "theta-lm",
			Port:  2296,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "theta-lm",
			Port:  2296,
		},
	},
	"thrp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "thrp",
			Port:  3963,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "thrp",
			Port:  3963,
		},
	},
	"thrtx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "thrtx",
			Port:  4139,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "thrtx",
			Port:  4139,
		},
	},
	"tht-treasure": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tht-treasure",
			Port:  1832,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tht-treasure",
			Port:  1832,
		},
	},
	"ticf-1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ticf-1",
			Port:  492,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ticf-1",
			Port:  492,
		},
	},
	"ticf-2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "ticf-2",
			Port:  493,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "ticf-2",
			Port:  493,
		},
	},
	"tick-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tick-port",
			Port:  3200,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tick-port",
			Port:  3200,
		},
	},
	"tidp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tidp",
			Port:  7548,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tidp",
			Port:  7548,
		},
	},
	"tig": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tig",
			Port:  3943,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tig",
			Port:  3943,
		},
	},
	"tigv2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tigv2",
			Port:  4124,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tigv2",
			Port:  4124,
		},
	},
	"timbuktu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timbuktu",
			Port:  407,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timbuktu",
			Port:  407,
		},
	},
	"timbuktu-srv1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timbuktu-srv1",
			Port:  1417,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timbuktu-srv1",
			Port:  1417,
		},
	},
	"timbuktu-srv2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timbuktu-srv2",
			Port:  1418,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timbuktu-srv2",
			Port:  1418,
		},
	},
	"timbuktu-srv3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timbuktu-srv3",
			Port:  1419,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timbuktu-srv3",
			Port:  1419,
		},
	},
	"timbuktu-srv4": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timbuktu-srv4",
			Port:  1420,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timbuktu-srv4",
			Port:  1420,
		},
	},
	"time": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "time",
			Port:  37,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "time",
			Port:  37,
		},
	},
	"timed": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timed",
			Port:  525,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timed",
			Port:  525,
		},
	},
	"timeflies": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timeflies",
			Port:  1362,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timeflies",
			Port:  1362,
		},
	},
	"timelot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timelot",
			Port:  3243,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timelot",
			Port:  3243,
		},
	},
	"timestenbroker": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "timestenbroker",
			Port:  3754,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "timestenbroker",
			Port:  3754,
		},
	},
	"tinc": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tinc",
			Port:  655,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tinc",
			Port:  655,
		},
	},
	"tinymessage": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "tinymessage",
			Port:  5104,
		},
	},
	"tip-app-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tip-app-server",
			Port:  3160,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tip-app-server",
			Port:  3160,
		},
	},
	"tip2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tip2",
			Port:  3372,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tip2",
			Port:  3372,
		},
	},
	"tipc": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "tipc",
			Port:  6118,
		},
	},
	"tircproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tircproxy",
			Port:  7666,
		},
	},
	"tivoconnect": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tivoconnect",
			Port:  2190,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tivoconnect",
			Port:  2190,
		},
	},
	"tivoli-npm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tivoli-npm",
			Port:  1965,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tivoli-npm",
			Port:  1965,
		},
	},
	"tksocket": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tksocket",
			Port:  2915,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tksocket",
			Port:  2915,
		},
	},
	"tl-ipcproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tl-ipcproxy",
			Port:  4176,
		},
	},
	"tl1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tl1",
			Port:  2361,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tl1",
			Port:  2361,
		},
	},
	"tl1-lv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tl1-lv",
			Port:  3081,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tl1-lv",
			Port:  3081,
		},
	},
	"tl1-raw": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tl1-raw",
			Port:  3082,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tl1-raw",
			Port:  3082,
		},
	},
	"tl1-raw-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tl1-raw-ssl",
			Port:  6251,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tl1-raw-ssl",
			Port:  6251,
		},
	},
	"tl1-ssh": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tl1-ssh",
			Port:  6252,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tl1-ssh",
			Port:  6252,
		},
	},
	"tl1-telnet": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tl1-telnet",
			Port:  3083,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tl1-telnet",
			Port:  3083,
		},
	},
	"tlisrv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tlisrv",
			Port:  1527,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tlisrv",
			Port:  1527,
		},
	},
	"tmesis-upshot": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tmesis-upshot",
			Port:  2798,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tmesis-upshot",
			Port:  2798,
		},
	},
	"tmi": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tmi",
			Port:  8300,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tmi",
			Port:  8300,
		},
	},
	"tmo-icon-sync": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tmo-icon-sync",
			Port:  5583,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tmo-icon-sync",
			Port:  5583,
		},
	},
	"tmophl7mts": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tmophl7mts",
			Port:  20046,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tmophl7mts",
			Port:  20046,
		},
	},
	"tmosms0": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tmosms0",
			Port:  5580,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tmosms0",
			Port:  5580,
		},
	},
	"tmosms1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tmosms1",
			Port:  5581,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tmosms1",
			Port:  5581,
		},
	},
	"tn-timing": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tn-timing",
			Port:  2739,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tn-timing",
			Port:  2739,
		},
	},
	"tn-tl-fd1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tn-tl-fd1",
			Port:  476,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tn-tl-fd1",
			Port:  476,
		},
	},
	"tn-tl-fd2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tn-tl-fd2",
			Port:  1584,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tn-tl-fd2",
			Port:  1584,
		},
	},
	"tn-tl-r1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tn-tl-r1",
			Port:  1580,
		},
	},
	"tn-tl-r2": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "tn-tl-r2",
			Port:  1580,
		},
	},
	"tn-tl-w1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tn-tl-w1",
			Port:  474,
		},
	},
	"tn-tl-w2": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "tn-tl-w2",
			Port:  474,
		},
	},
	"tnETOS": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tnETOS",
			Port:  377,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tnETOS",
			Port:  377,
		},
	},
	"tnmpv2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tnmpv2",
			Port:  3686,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tnmpv2",
			Port:  3686,
		},
	},
	"tnos-dp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tnos-dp",
			Port:  7902,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tnos-dp",
			Port:  7902,
		},
	},
	"tnos-dps": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tnos-dps",
			Port:  7903,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tnos-dps",
			Port:  7903,
		},
	},
	"tnos-sp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tnos-sp",
			Port:  7901,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tnos-sp",
			Port:  7901,
		},
	},
	"tnp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tnp",
			Port:  8321,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tnp",
			Port:  8321,
		},
	},
	"tnp-discover": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tnp-discover",
			Port:  8320,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tnp-discover",
			Port:  8320,
		},
	},
	"tnp1-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tnp1-port",
			Port:  4024,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tnp1-port",
			Port:  4024,
		},
	},
	"tns-adv": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tns-adv",
			Port:  3309,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tns-adv",
			Port:  3309,
		},
	},
	"tns-cml": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tns-cml",
			Port:  590,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tns-cml",
			Port:  590,
		},
	},
	"tns-server": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tns-server",
			Port:  3308,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tns-server",
			Port:  3308,
		},
	},
	"toad": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "toad",
			Port:  2669,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "toad",
			Port:  2669,
		},
	},
	"tolfab": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tolfab",
			Port:  20167,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tolfab",
			Port:  20167,
		},
	},
	"tolteces": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tolteces",
			Port:  4375,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tolteces",
			Port:  4375,
		},
	},
	"tomato-springs": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tomato-springs",
			Port:  3040,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tomato-springs",
			Port:  3040,
		},
	},
	"tonidods": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tonidods",
			Port:  24465,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tonidods",
			Port:  24465,
		},
	},
	"topflow": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "topflow",
			Port:  2885,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "topflow",
			Port:  2885,
		},
	},
	"topflow-ssl": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "topflow-ssl",
			Port:  3885,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "topflow-ssl",
			Port:  3885,
		},
	},
	"topovista-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "topovista-data",
			Port:  3906,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "topovista-data",
			Port:  3906,
		},
	},
	"topx": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "topx",
			Port:  2436,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "topx",
			Port:  2436,
		},
	},
	"toruxserver": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "toruxserver",
			Port:  5153,
		},
	},
	"touchnetplus": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "touchnetplus",
			Port:  2158,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "touchnetplus",
			Port:  2158,
		},
	},
	"tpcsrvr": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tpcsrvr",
			Port:  2078,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tpcsrvr",
			Port:  2078,
		},
	},
	"tpdu": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tpdu",
			Port:  1430,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tpdu",
			Port:  1430,
		},
	},
	"tpip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tpip",
			Port:  594,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tpip",
			Port:  594,
		},
	},
	"tpmd": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tpmd",
			Port:  1906,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tpmd",
			Port:  1906,
		},
	},
	"tproxy": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tproxy",
			Port:  8081,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tproxy",
			Port:  8081,
		},
	},
	"tqdata": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tqdata",
			Port:  2700,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tqdata",
			Port:  2700,
		},
	},
	"tr-rsrb-p1": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tr-rsrb-p1",
			Port:  1987,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tr-rsrb-p1",
			Port:  1987,
		},
	},
	"tr-rsrb-p2": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tr-rsrb-p2",
			Port:  1988,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tr-rsrb-p2",
			Port:  1988,
		},
	},
	"tr-rsrb-p3": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tr-rsrb-p3",
			Port:  1989,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tr-rsrb-p3",
			Port:  1989,
		},
	},
	"tr-rsrb-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tr-rsrb-port",
			Port:  1996,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tr-rsrb-port",
			Port:  1996,
		},
	},
	"traceroute": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "traceroute",
			Port:  33434,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "traceroute",
			Port:  33434,
		},
	},
	"track": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "track",
			Port:  20670,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "track",
			Port:  20670,
		},
	},
	"tragic": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tragic",
			Port:  2642,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tragic",
			Port:  2642,
		},
	},
	"traingpsdata": map[string]Service{
		"udp": Service{
			Proto: "udp",
			Name:  "traingpsdata",
			Port:  9277,
		},
	},
	"tram": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tram",
			Port:  4567,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tram",
			Port:  4567,
		},
	},
	"transact": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "transact",
			Port:  1869,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "transact",
			Port:  1869,
		},
	},
	"transmit-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "transmit-port",
			Port:  5282,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "transmit-port",
			Port:  5282,
		},
	},
	"trap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trap",
			Port:  4020,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trap",
			Port:  4020,
		},
	},
	"trap-daemon": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trap-daemon",
			Port:  3600,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trap-daemon",
			Port:  3600,
		},
	},
	"trap-port": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trap-port",
			Port:  3857,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trap-port",
			Port:  3857,
		},
	},
	"trap-port-mom": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trap-port-mom",
			Port:  3858,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trap-port-mom",
			Port:  3858,
		},
	},
	"traversal": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "traversal",
			Port:  4678,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "traversal",
			Port:  4678,
		},
	},
	"travsoft-ipx-t": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "travsoft-ipx-t",
			Port:  2644,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "travsoft-ipx-t",
			Port:  2644,
		},
	},
	"trc-netpoll": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trc-netpoll",
			Port:  2405,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trc-netpoll",
			Port:  2405,
		},
	},
	"treehopper": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "treehopper",
			Port:  3959,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "treehopper",
			Port:  3959,
		},
	},
	"trendchip-dcp": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trendchip-dcp",
			Port:  3608,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trendchip-dcp",
			Port:  3608,
		},
	},
	"tributary": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tributary",
			Port:  2580,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tributary",
			Port:  2580,
		},
	},
	"trident-data": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trident-data",
			Port:  7727,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trident-data",
			Port:  7727,
		},
	},
	"trim": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trim",
			Port:  1137,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trim",
			Port:  1137,
		},
	},
	"trim-event": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trim-event",
			Port:  4322,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trim-event",
			Port:  4322,
		},
	},
	"trim-ice": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trim-ice",
			Port:  4323,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trim-ice",
			Port:  4323,
		},
	},
	"triomotion": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "triomotion",
			Port:  3240,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "triomotion",
			Port:  3240,
		},
	},
	"trip": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trip",
			Port:  6069,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trip",
			Port:  6069,
		},
	},
	"tripe": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tripe",
			Port:  4070,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tripe",
			Port:  4070,
		},
	},
	"tripwire": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "tripwire",
			Port:  1169,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "tripwire",
			Port:  1169,
		},
	},
	"triquest-lm": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "triquest-lm",
			Port:  1588,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "triquest-lm",
			Port:  1588,
		},
	},
	"trisoap": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trisoap",
			Port:  10200,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trisoap",
			Port:  10200,
		},
	},
	"trispen-sra": map[string]Service{
		"tcp": Service{
			Proto: "tcp",
			Name:  "trispen-sra",
			Port:  9555,
		},
		"udp": Service{
			Proto: "udp",
			Name:  "trispen-sra",
			Port:  9555,
		},
	},
		"tcp": Service{
		},
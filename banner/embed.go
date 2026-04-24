package banner

import _ "embed"

//go:embed rpc.txt
var Rpc string

//go:embed rpc-mirror.txt
var RpcMirror string

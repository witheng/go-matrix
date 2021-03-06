// Copyright (c) 2018-2019 The MATRIX Authors
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php
// +build cgo
// +build !appengine

package metrics

import "runtime"

func numCgoCall() int64 {
	return runtime.NumCgoCall()
}

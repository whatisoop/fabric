/*
 * Copyright Greg Haskins All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * See github.com/whatisoop/fabric/test/chaincodes/AutoVendor/chaincode/main.go for details
 */
package indirectdep

import "fmt"

func PointlessFunction() {
	fmt.Printf("Successfully invoked pointless function\n")
}

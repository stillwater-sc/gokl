/*
 File		:	$File: //depot/stillwater-sc/gokl/initialize.go $

 Authors	:	E. Theodore L. Omtzigt
 Date		:	6 April 2014

 Source Control Information:
 Version	:	$Revision: #1 $
 Latest		:	$Date: 2014/04/06 $
 Location	:	$Id: //depot/stillwater-sc/gokl/initialize.go#1 $

 Organization:
		Stillwater Supercomputing, Inc.
		P.O Box 720
		South Freeport, ME 04078-0720

Copyright (c) 2014-2017 E. Theodore L. Omtzigt.  All rights reserved.

Licence      : MIT license as defined in this directory

 */
package gokl

import "fmt"

func Initialize() (rm *ResourceManager) {
	rm = &ResourceManager{}
	rm.Initialize("Default RM")
	return
}

func PrettyPrint(mat [][]float32) string {
	var str string
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			str = str + fmt.Sprintf("%8.3f ", mat[i][j])
		}
		str = str + "\n"
	}
	return str
}

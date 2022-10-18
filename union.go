// Copyright union-pdd-go Author(https://houseme.github.io/union-pdd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-pdd-go.

package pdd

const (
	version = "0.01"
)

// UnionPDD is a union service.
type UnionPDD struct {
	// contains filtered or unexported fields
}

// NewUnionPDD returns a union service.
func NewUnionPDD() *UnionPDD {
	return &UnionPDD{}
}

// Version returns the version of the union service.
func (u *UnionPDD) Version() string {
	return version
}

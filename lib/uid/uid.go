package uidlib

import gonanoid "github.com/matoous/go-nanoid"

type UidStruct struct{}

var UidLib = new(UidStruct)

const UidAlphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
const UidLen = 16

func (r UidStruct) GenUid() (string, error) {
	return gonanoid.Generate(UidAlphabet, UidLen)
}

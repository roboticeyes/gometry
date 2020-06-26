package geom

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestCreateInvalid(t *testing.T) {

	r := bytes.NewReader([]byte("anything"))
	i := NewImage("img", r)
	if i.Valid() {
		t.Fatal("Image should not be valid")
	}
}

func TestCreateValid(t *testing.T) {

	b, err := base64.StdEncoding.DecodeString(pngImage)
	if err != nil {
		t.Fatal("Image not valid")
	}
	r := bytes.NewReader(b)
	i := NewImage("image", r)
	if !i.Valid() {
		t.Fatal("Image should be valid")
	}

	if i.Name() != "image" {
		t.Fatal("Image name does not fit")
	}

	if i.Width() != 16 {
		t.Fatal("Image width is not correct")
	}
	if i.Height() != 16 {
		t.Fatal("Image height is not correct")
	}
}

const pngImage = `
iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAIAAACQkWg2AAABhGlDQ1BJQ0MgcHJvZmlsZQAAKJF9
kT1Iw0AcxV9btUUqDnaQ4pChOlmQKuKoVShChVArtOpgcukXNGlIUlwcBdeCgx+LVQcXZ10dXAVB
8APEydFJ0UVK/F9SaBHjwXE/3t173L0D/M0qU82eCUDVLCOTSgq5/KoQfEUfogghjITETH1OFNPw
HF/38PH1Ls6zvM/9OQaUgskAn0A8y3TDIt4gnt60dM77xBFWlhTic+Jxgy5I/Mh12eU3ziWH/Twz
YmQz88QRYqHUxXIXs7KhEk8RxxRVo3x/zmWF8xZntVpn7XvyF4YL2soy12mOIIVFLEGEABl1VFCF
hTitGikmMrSf9PBHHb9ILplcFTByLKAGFZLjB/+D392axcmEmxROAr0vtv0xCgR3gVbDtr+Pbbt1
AgSegSut4681gZlP0hsdLXYEDG4DF9cdTd4DLneA4SddMiRHCtD0F4vA+xl9Ux4YugX619ze2vs4
fQCy1FX6Bjg4BMZKlL3u8e5Qd2//nmn39wNo8XKjCrGeBQAAAAlwSFlzAAAuIwAALiMBeKU/dgAA
AAd0SU1FB+QGGggNK+bkZ9oAAAAZdEVYdENvbW1lbnQAQ3JlYXRlZCB3aXRoIEdJTVBXgQ4XAAAB
EElEQVQoz7WSO2vDQBCEd+5OSLaMdZaMVAg3Lvz/f40KFwYbN35iC/SCu9sUCUokCIkJmXLZD3Zn
BsxMr0jQi1LjATM5x8YQEQFQigACvgGc467junZlSUQkpdAaQQDP6xk12G4ae7nY49HdbmQtJhOR
ZTLPxWLRM58AG2Pvd7Pd2qLg65WbBlqLPCdj4HnQegg4x23rzme727micIcDNQ2ShKsKUSSSBGEI
3/8CMHPXcVny48GnE+33VNfcthTH/Hy6qpLW/tFWAL6P+RxaI8u4695PwnKJKBJhSFIOASEQBCJN
5XpNbYs07Z+Wq9WHSyNboZSMY9psMJ2ObMVs1ueAQZd+ERzG5fupGvj3tr4BPW2XItubnPoAAAAA
SUVORK5CYII=
`

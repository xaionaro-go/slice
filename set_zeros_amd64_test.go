// +build amd64

package slice

func init() {
	implementations["words"] = setZerosWords
}

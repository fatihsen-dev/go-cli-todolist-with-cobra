package test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/fatihsen-dev/go-cli-todolist-with-cobra/cmd"
)

func TestTypeLocal(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd.Root.SetOut(buf)
	cmd.Root.SetArgs([]string{"create", "--name=test"})

	err := cmd.Root.Execute()
	if err != nil {
		fmt.Println(err)
	}
	if buf.String() != "test" {
		t.Errorf("Create komutunun testi başarısız oldu - Çıktı: %s", buf.String())
	}

}

package iterative

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScript(t *testing.T) {
	t.Run("AWS known region should not add the NVIDA drivers", func(t *testing.T) {
		data := make(map[string]interface{})
		data["ami"] = isAMIAvailable("aws", "us-east-1")

		script, _ := renderScript(data)
		assert.Equal(t, strings.Contains(script, "sudo ubuntu-drivers autoinstall"), false)
	})

	t.Run("AWS unknown region should add the NVIDA drivers", func(t *testing.T) {
		data := make(map[string]interface{})
		data["ami"] = isAMIAvailable("aws", "us-east-99")

		script, _ := renderScript(data)
		assert.Equal(t, strings.Contains(script, "sudo ubuntu-drivers autoinstall"), true)
	})

	t.Run("Azure known region should add the NVIDA drivers", func(t *testing.T) {
		data := make(map[string]interface{})
		data["ami"] = isAMIAvailable("azure", "westus")

		script, _ := renderScript(data)
		assert.Equal(t, strings.Contains(script, "sudo ubuntu-drivers autoinstall"), true)
	})

	t.Run("Azure unknown region should add the NVIDA drivers", func(t *testing.T) {
		data := make(map[string]interface{})
		data["ami"] = isAMIAvailable("azure", "us-east-99")

		script, _ := renderScript(data)
		assert.Equal(t, strings.Contains(script, "sudo ubuntu-drivers autoinstall"), true)
	})

	t.Run("Runner Startup Script", func(t *testing.T) {
		data := make(map[string]interface{})
		startupScript, _ := base64.StdEncoding.DecodeString("ZWNobyAiaGVsbG8gd29ybGQiCmVjaG8gImJ5ZSB3b3JsZCI=")
		data["runner_startup_script"] = string(startupScript)

		script, _ := renderScript(data)
		assert.Equal(t, strings.Contains(script, "echo \"hello world\"\necho \"bye world\""), true)
	})
}

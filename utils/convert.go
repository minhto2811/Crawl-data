package utils

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var count,max = 0,20

func KillAllSoffice() {
	exec.Command("taskkill", "/IM", "soffice.exe", "/F").Run()
}



func ConvertDocxToPDF(input, outputDir string, workerID int) (string, error) {
	count++
	if count >= max {
		KillAllSoffice()
		count = 0
	}

	profileDir := fmt.Sprintf("C:/temp/lo-profile-%d", workerID)
	_ = os.MkdirAll(profileDir, 0755)

	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		`C:\Program Files\LibreOffice\program\soffice.exe`,
		"--headless",
		"--nologo",
		"--nodefault",
		"--nolockcheck",
		"--norestore",
		"-env:UserInstallation=file:///" + strings.ReplaceAll(profileDir, `\`, `/`),
		"--convert-to", "pdf",
		input,
		"--outdir", outputDir,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("timeout: soffice bị treo")
		}
		return "", fmt.Errorf("lỗi convert pdf: %v", err)
	}

	base := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
	return filepath.Join(outputDir, base+".pdf"), nil
}


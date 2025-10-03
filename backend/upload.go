package backend

import (
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// testing error handling remains here
func (a *App) UploadFile(filePath string) {
	a.logger.Info(filePath)
	conns := make([]net.Conn, len(a.ips))
	for idx, ip := range a.ips {
		var err error
		conns[idx], err = a.handshake(ip)
		if err != nil {
			a.logger.Warn(err.Error())
		}
		conns[idx].Write([]byte{1 << 2})
	}
}

func (a *App) SelectFilesToUpload() ([]string, error) {
	filePaths, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a file to upload",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "All Files",
				Pattern:     "*",
			},
		},
	})

	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	return filePaths, nil
}

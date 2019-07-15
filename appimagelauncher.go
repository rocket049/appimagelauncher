//Help appimage program to create .desktop file into ~/.local/share/applications
package appimagelauncher

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	fp1, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fp1.Close()
	fp2, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fp2.Close()
	_, err = io.Copy(fp1, fp2)
	log.Println("copy", src, dst)
	return err
}

func isNewer(fileNew, fileOld string) bool {
	infoNew, err := os.Lstat(fileNew)
	if err != nil {
		return false
	}
	infoOld, err := os.Lstat(fileOld)
	if err != nil {
		return true
	}
	if infoNew.ModTime().Unix() > infoOld.ModTime().Unix() {
		return true
	}
	return false
}

//Create Copy .desktop file from APPDIR to ~/.local/share/applications.
//Copy icon file from APPDIR to ~/.local/share/icons
//Replace Exec value with APPIMAGE while copying.
//Skip if the launcher is newer than APPIMAGE file when 'force' is false.
func Create(desktopFile, iconFile string, force bool) error {
	desktop := desktopFile
	icon := iconFile

	appimage := os.Getenv("APPIMAGE")
	appdir := os.Getenv("APPDIR")
	if len(appimage) == 0 || len(appdir) == 0 {
		return errors.New("Not Appimage")
	}

	home, _ := os.UserHomeDir()

	src := filepath.Join(appdir, desktop)
	dst := filepath.Join(home, ".local", "share", "applications", desktop)
	if isNewer(dst, appimage) && force == false {
		return nil
	}

	iconSrc := filepath.Join(appdir, icon)

	iconDir := filepath.Join(home, ".local", "share", "icons")
	iconDst := filepath.Join(iconDir, icon)

	os.MkdirAll(iconDir, os.ModePerm)

	err := copyFile(iconSrc, iconDst)
	if err != nil {
		log.Println(err)
	}

	srcFp, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return err
	}
	defer srcFp.Close()
	reader := bufio.NewReader(srcFp)

	fp, err := os.Create(dst)
	if err != nil {
		log.Println(err)
		return err
	}
	defer fp.Close()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if bytes.HasPrefix(line, []byte("Exec=")) {
			fp.WriteString("Exec=" + appimage + "\n")
		} else {
			fp.Write(line)
			fp.WriteString("\n")
		}
	}
	return nil
}

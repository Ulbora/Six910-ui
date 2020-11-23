package bkupsrv

/*
 Six910 is a shopping cart and E-commerce system.
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.
 Copyright (C) 2020 Ken Williamson
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"archive/tar"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"encoding/json"

	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	contentStore  = "contentStore"
	templateStore = "templateStore"
	imageFiles    = "imageFiles"
	templateFiles = "templateFiles"
)

//BackupFiles BackupFiles
type BackupFiles struct {
	CarouselStoreFiles *[]BackupFile
	ContentStoreFiles  *[]BackupFile
	CountryStoreFiles  *[]BackupFile
	CSSStoreFiles      *[]BackupFile
	MenuStoreFiles     *[]BackupFile
	StateStoreFiles    *[]BackupFile
	TemplateStoreFiles *[]BackupFile
	ImageFiles         *[]BackupFile
	TemplateFiles      *BackupFile
}

//BackupFile BackupFile
type BackupFile struct {
	FilesLocation string
	Name          string
	FileData      []byte
}

//UploadBackups UploadBackups
func (c *Six910BackupService) UploadBackups(bk *[]byte) bool {
	var rtn bool
	var bkfs BackupFiles
	var b bytes.Buffer
	b.Write(*bk)
	r, err := zlib.NewReader(&b)
	if err == nil {
		var out bytes.Buffer
		io.Copy(&out, r)
		r.Close()
		bout := out.Bytes()
		//c.Log.Debug("content file in upload: ", string(bout))
		umerr := json.Unmarshal(bout, &bkfs)
		c.Log.Debug("BackupFiles file unmarshal err: ", umerr)
		c.Log.Debug("BackupFiles file unmarshal : ", bkfs)

		// content store files
		os.RemoveAll(c.ContentStorePath)
		os.Mkdir(c.ContentStorePath, os.FileMode(0777))

		for _, cf := range *bkfs.ContentStoreFiles {
			c.Log.Debug("BackupFile content file name: ", c.ContentStorePath+string(filepath.Separator)+cf.Name)
			//c.Log.Debug("BackupFile content file: ", cf)
			werr := ioutil.WriteFile(c.ContentStorePath+string(filepath.Separator)+cf.Name, cf.FileData, 0666)
			c.Log.Debug("BackupFile content file write err : ", werr)
		}

		c.Store.Reload()

		// template store files
		os.RemoveAll(c.TemplateStorePath)
		os.Mkdir(c.TemplateStorePath, os.FileMode(0777))

		for _, cf := range *bkfs.TemplateStoreFiles {
			c.Log.Debug("BackupFile template file name: ", c.TemplateStorePath+string(filepath.Separator)+cf.Name)
			//c.Log.Debug("BackupFile template file: ", cf)
			werr := ioutil.WriteFile(c.TemplateStorePath+string(filepath.Separator)+cf.Name, cf.FileData, 0666)
			c.Log.Debug("BackupFile template file write err : ", werr)
		}

		c.TemplateStore.Reload()

		//carousel store files
		os.RemoveAll(c.CarouselStorePath)
		os.Mkdir(c.CarouselStorePath, os.FileMode(0777))

		for _, cf := range *bkfs.CarouselStoreFiles {
			c.Log.Debug("BackupFile Carousel file name: ", c.CarouselStorePath+string(filepath.Separator)+cf.Name)
			//c.Log.Debug("BackupFile Carousel file: ", cf)
			werr := ioutil.WriteFile(c.CarouselStorePath+string(filepath.Separator)+cf.Name, cf.FileData, 0666)
			c.Log.Debug("BackupFile Carousel file write err : ", werr)
		}

		c.CarouselStore.Reload()

		// country store files
		os.RemoveAll(c.CountryStorePath)
		os.Mkdir(c.CountryStorePath, os.FileMode(0777))

		for _, cf := range *bkfs.CountryStoreFiles {
			c.Log.Debug("BackupFile country file name: ", c.CountryStorePath+string(filepath.Separator)+cf.Name)
			//c.Log.Debug("BackupFile country file: ", cf)
			werr := ioutil.WriteFile(c.CountryStorePath+string(filepath.Separator)+cf.Name, cf.FileData, 0666)
			c.Log.Debug("BackupFile country file write err : ", werr)
		}

		c.CountryStore.Reload()

		// css store files
		os.RemoveAll(c.CSSStorePath)
		os.Mkdir(c.CSSStorePath, os.FileMode(0777))

		for _, cf := range *bkfs.CSSStoreFiles {
			c.Log.Debug("BackupFile css file name: ", c.CSSStorePath+string(filepath.Separator)+cf.Name)
			//c.Log.Debug("BackupFile css file: ", cf)
			werr := ioutil.WriteFile(c.CSSStorePath+string(filepath.Separator)+cf.Name, cf.FileData, 0666)
			c.Log.Debug("BackupFile css file write err : ", werr)
		}

		c.CSSStore.Reload()

		// menu store files
		os.RemoveAll(c.MenuStorePath)
		os.Mkdir(c.MenuStorePath, os.FileMode(0777))

		for _, cf := range *bkfs.MenuStoreFiles {
			c.Log.Debug("BackupFile menu file name: ", c.MenuStorePath+string(filepath.Separator)+cf.Name)
			//c.Log.Debug("BackupFile menu file: ", cf)
			werr := ioutil.WriteFile(c.MenuStorePath+string(filepath.Separator)+cf.Name, cf.FileData, 0666)
			c.Log.Debug("BackupFile menu file write err : ", werr)
		}

		c.MenuStore.Reload()

		// state store files
		os.RemoveAll(c.StateStorePath)
		os.Mkdir(c.StateStorePath, os.FileMode(0777))

		for _, cf := range *bkfs.StateStoreFiles {
			c.Log.Debug("BackupFile state file name: ", c.StateStorePath+string(filepath.Separator)+cf.Name)
			//c.Log.Debug("BackupFile state file: ", cf)
			werr := ioutil.WriteFile(c.StateStorePath+string(filepath.Separator)+cf.Name, cf.FileData, 0666)
			c.Log.Debug("BackupFile state file write err : ", werr)
		}

		c.StateStore.Reload()

		// image files
		os.RemoveAll(c.ImagePath)
		os.Mkdir(c.ImagePath, os.FileMode(0777))

		for _, cf := range *bkfs.ImageFiles {
			c.Log.Debug("BackupFile image file name: ", c.ImagePath+string(filepath.Separator)+cf.Name)
			//c.Log.Debug("BackupFile image file: ", cf)
			werr := ioutil.WriteFile(c.ImagePath+string(filepath.Separator)+cf.Name, cf.FileData, 0666)
			c.Log.Debug("BackupFile image file write err : ", werr)
		}

		// template files
		os.RemoveAll(c.TemplateFilePath)
		os.Mkdir(c.TemplateFilePath, os.FileMode(0777))

		r := bytes.NewReader(bkfs.TemplateFiles.FileData)
		gzf, gzerr := gzip.NewReader(r)
		c.Log.Debug("BackupFile template file reader err : ", gzerr)
		tr := tar.NewReader(gzf)
		for {
			hdr, err := tr.Next()
			c.Log.Debug("new reader next error in extract file: ", err)
			if err == io.EOF {
				break
			}
			eName := c.TemplateFilePath
			c.Log.Debug("eName in extract file: ", eName)
			err2 := c.extractTarGzFile(tr, hdr, eName)
			c.Log.Debug("extractTarGzFile error in extract file: ", err2)
			defer gzf.Close()
		}
		rtn = true
	}
	return rtn
}

//DownloadBackups DownloadBackups
func (c *Six910BackupService) DownloadBackups() (bool, *[]byte) {
	var rtn bool
	var bkfs BackupFiles

	//contentStore

	c.processCarouselFiles(&bkfs)
	c.processContentFiles(&bkfs)
	c.processCountryFiles(&bkfs)
	c.processCSSFiles(&bkfs)
	c.processMenuFiles(&bkfs)
	c.processStateFiles(&bkfs)
	c.processTemplateFiles(&bkfs)
	c.processImageFiles(&bkfs)

	//zip template files
	cwpath, _ := os.Getwd()
	//fmt.Println("current dir: ", cwpath)
	var buf bytes.Buffer
	zr := gzip.NewWriter(&buf)
	tw := tar.NewWriter(zr)
	files, err := ioutil.ReadDir(c.TemplateFilePath)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				os.Chdir(c.TemplateFilePath + string(filepath.Separator))
				c.compress(file.Name(), tw)
				os.Chdir(cwpath)
			} else {
				c.Log.Debug("template file: ", c.TemplateFilePath+string(filepath.Separator)+file.Name())
				//c.compress(file.Name(), tw)
				data, oerr := os.Open(c.TemplateFilePath + string(filepath.Separator) + file.Name())
				c.Log.Debug("oerr: ", oerr)
				if oerr == nil {
					info, err := data.Stat()
					c.Log.Debug("err: ", err)
					if err == nil {
						header, err := tar.FileInfoHeader(info, info.Name())
						c.Log.Debug("err FileInfoHeader: ", err)
						if err == nil {
							header.Name = file.Name()
							err = tw.WriteHeader(header)
							c.Log.Debug("err tw.WriteHeader: ", err)
							_, cerr := io.Copy(tw, data)
							c.Log.Debug("cerr io.Copy: ", cerr)
						}

					}
				}
				//fileData, rerr := ioutil.ReadFile(c.TemplateFilePath + string(filepath.Separator) + file.Name())
				//c.Log.Debug("template file data: ", fileData)
				// if rerr == nil {
				// 	var cbk BackupFile
				// 	cbk.Name = file.Name()
				// 	cbk.FilesLocation = c.TemplateFilePath
				// 	cbk.FileData = fileData
				// 	//templateStoreFiles = append(TemplateFilePath, cbk)
				// }
			}
		}
		rtn = true
	}
	zr.Close()
	tw.Close()

	var tbkf BackupFile
	tbkf.Name = "templates.tar.gz"
	tbkf.FilesLocation = c.TemplateFilePath
	tbkf.FileData = buf.Bytes()
	bkfs.TemplateFiles = &tbkf

	// //-----------test
	// fileToWrite, err := os.OpenFile("./testBackupZips/testtemplates.tar.gz", os.O_CREATE|os.O_RDWR, os.FileMode(777))
	// fmt.Println("fileToWrite err : ", err)
	// var buf2 = bytes.NewBuffer(buf.Bytes())

	// //fmt.Println("compress file data: ", buf.Bytes())
	// _, err2 := io.Copy(fileToWrite, buf2)
	// fmt.Println("io.copy err : ", err2)
	// os.Chmod("./testBackupZips/compress.dat", os.FileMode(0666))
	// //

	c.Log.Debug("backup file: ", bkfs)

	bts, _ := json.Marshal(bkfs)

	var b bytes.Buffer
	w, err := zlib.NewWriterLevel(&b, flate.BestCompression)
	if err == nil {
		w.Write(bts)
		w.Close()
	}

	compressedData := b.Bytes()
	//c.Log.Debug("backup file compressedData: ", compressedData)

	return rtn, &compressedData
}

func (c *Six910BackupService) compress(dir string, tw *tar.Writer) {
	//os.Chdir(dir)
	filepath.Walk(dir, func(file string, fi os.FileInfo, err error) error {
		var errr error
		c.Log.Debug("file in walk : ", file)
		header, herr := tar.FileInfoHeader(fi, file)
		errr = herr
		if herr == nil {
			header.Name = filepath.ToSlash(file)
			hrerr := tw.WriteHeader(header)
			errr = hrerr
			if hrerr == nil {
				if !fi.IsDir() {
					data, oerr := os.Open(file)
					errr = oerr
					if oerr == nil {
						_, cerr := io.Copy(tw, data)
						errr = cerr
					}
				}
			}
		}
		return errr
	})
}

func (c *Six910BackupService) processContentFiles(bkfs *BackupFiles) {
	var contStoreFiles []BackupFile
	cntfiles, err := ioutil.ReadDir(c.ContentStorePath)
	if err == nil {
		for _, sfile := range cntfiles {
			if !sfile.IsDir() {
				c.Log.Debug("content store file: ", c.ContentStorePath+string(filepath.Separator)+sfile.Name())
				fileData, rerr := ioutil.ReadFile(c.ContentStorePath + string(filepath.Separator) + sfile.Name())
				//c.Log.Debug("content store file data: ", fileData)
				if rerr == nil {
					var cbk BackupFile
					cbk.Name = sfile.Name()
					cbk.FilesLocation = c.ContentStorePath
					cbk.FileData = fileData
					contStoreFiles = append(contStoreFiles, cbk)
				}
			}
		}
		//c.Log.Debug("content store file list: ", contStoreFiles)
		bkfs.ContentStoreFiles = &contStoreFiles
	}
}

func (c *Six910BackupService) processCarouselFiles(bkfs *BackupFiles) {
	var carStoreFiles []BackupFile
	carfiles, err := ioutil.ReadDir(c.CarouselStorePath)
	if err == nil {
		for _, sfile := range carfiles {
			if !sfile.IsDir() {
				c.Log.Debug("carousel store file: ", c.CarouselStorePath+string(filepath.Separator)+sfile.Name())
				fileData, rerr := ioutil.ReadFile(c.CarouselStorePath + string(filepath.Separator) + sfile.Name())
				//c.Log.Debug("content store file data: ", fileData)
				if rerr == nil {
					var cbk BackupFile
					cbk.Name = sfile.Name()
					cbk.FilesLocation = c.CarouselStorePath
					cbk.FileData = fileData
					carStoreFiles = append(carStoreFiles, cbk)
				}
			}
		}
		//c.Log.Debug("content store file list: ", contStoreFiles)
		bkfs.CarouselStoreFiles = &carStoreFiles
	}
}

func (c *Six910BackupService) processCountryFiles(bkfs *BackupFiles) {
	var ctryStoreFiles []BackupFile
	ctryfiles, err := ioutil.ReadDir(c.CountryStorePath)
	if err == nil {
		for _, sfile := range ctryfiles {
			if !sfile.IsDir() {
				c.Log.Debug("country store file: ", c.CountryStorePath+string(filepath.Separator)+sfile.Name())
				fileData, rerr := ioutil.ReadFile(c.CountryStorePath + string(filepath.Separator) + sfile.Name())
				//c.Log.Debug("content store file data: ", fileData)
				if rerr == nil {
					var cbk BackupFile
					cbk.Name = sfile.Name()
					cbk.FilesLocation = c.CountryStorePath
					cbk.FileData = fileData
					ctryStoreFiles = append(ctryStoreFiles, cbk)
				}
			}
		}
		//c.Log.Debug("content store file list: ", contStoreFiles)
		bkfs.CountryStoreFiles = &ctryStoreFiles
	}
}

func (c *Six910BackupService) processCSSFiles(bkfs *BackupFiles) {
	var cssStoreFiles []BackupFile
	cssfiles, err := ioutil.ReadDir(c.CSSStorePath)
	if err == nil {
		for _, sfile := range cssfiles {
			if !sfile.IsDir() {
				c.Log.Debug("css store file: ", c.CSSStorePath+string(filepath.Separator)+sfile.Name())
				fileData, rerr := ioutil.ReadFile(c.CSSStorePath + string(filepath.Separator) + sfile.Name())
				//c.Log.Debug("content store file data: ", fileData)
				if rerr == nil {
					var cbk BackupFile
					cbk.Name = sfile.Name()
					cbk.FilesLocation = c.CSSStorePath
					cbk.FileData = fileData
					cssStoreFiles = append(cssStoreFiles, cbk)
				}
			}
		}
		//c.Log.Debug("content store file list: ", contStoreFiles)
		bkfs.CSSStoreFiles = &cssStoreFiles
	}
}

func (c *Six910BackupService) processMenuFiles(bkfs *BackupFiles) {
	var muStoreFiles []BackupFile
	mufiles, err := ioutil.ReadDir(c.MenuStorePath)
	if err == nil {
		for _, sfile := range mufiles {
			if !sfile.IsDir() {
				c.Log.Debug("menu store file: ", c.MenuStorePath+string(filepath.Separator)+sfile.Name())
				fileData, rerr := ioutil.ReadFile(c.MenuStorePath + string(filepath.Separator) + sfile.Name())
				//c.Log.Debug("content store file data: ", fileData)
				if rerr == nil {
					var cbk BackupFile
					cbk.Name = sfile.Name()
					cbk.FilesLocation = c.MenuStorePath
					cbk.FileData = fileData
					muStoreFiles = append(muStoreFiles, cbk)
				}
			}
		}
		//c.Log.Debug("content store file list: ", contStoreFiles)
		bkfs.MenuStoreFiles = &muStoreFiles
	}
}

func (c *Six910BackupService) processStateFiles(bkfs *BackupFiles) {
	var stStoreFiles []BackupFile
	stfiles, err := ioutil.ReadDir(c.StateStorePath)
	if err == nil {
		for _, sfile := range stfiles {
			if !sfile.IsDir() {
				c.Log.Debug("state store file: ", c.StateStorePath+string(filepath.Separator)+sfile.Name())
				fileData, rerr := ioutil.ReadFile(c.StateStorePath + string(filepath.Separator) + sfile.Name())
				//c.Log.Debug("content store file data: ", fileData)
				if rerr == nil {
					var cbk BackupFile
					cbk.Name = sfile.Name()
					cbk.FilesLocation = c.StateStorePath
					cbk.FileData = fileData
					stStoreFiles = append(stStoreFiles, cbk)
				}
			}
		}
		//c.Log.Debug("content store file list: ", contStoreFiles)
		bkfs.StateStoreFiles = &stStoreFiles
	}
}

func (c *Six910BackupService) processTemplateFiles(bkfs *BackupFiles) {
	var templateStoreFiles []BackupFile
	tempfiles, err := ioutil.ReadDir(c.TemplateStorePath)
	if err == nil {
		for _, sfile := range tempfiles {
			if !sfile.IsDir() {
				c.Log.Debug("template store file: ", c.TemplateStorePath+string(filepath.Separator)+sfile.Name())
				fileData, rerr := ioutil.ReadFile(c.TemplateStorePath + string(filepath.Separator) + sfile.Name())
				c.Log.Debug("template store  file data: ", fileData)
				if rerr == nil {
					var cbk BackupFile
					cbk.Name = sfile.Name()
					cbk.FilesLocation = c.TemplateStorePath
					cbk.FileData = fileData
					templateStoreFiles = append(templateStoreFiles, cbk)
				}
			}
		}
		c.Log.Debug("template file list: ", templateStoreFiles)
		bkfs.TemplateStoreFiles = &templateStoreFiles
	}
}

func (c *Six910BackupService) processImageFiles(bkfs *BackupFiles) {
	var imageFiles []BackupFile
	imgfiles, err := ioutil.ReadDir(c.ImagePath)
	if err == nil {
		c.Log.Debug("imgfiles: ", imgfiles)
		for _, sfile := range imgfiles {
			if !sfile.IsDir() {
				c.Log.Debug("image file: ", c.ImagePath+string(filepath.Separator)+sfile.Name())
				fileData, rerr := ioutil.ReadFile(c.ImagePath + string(filepath.Separator) + sfile.Name())
				if rerr == nil {
					var cbk BackupFile
					cbk.Name = sfile.Name()
					cbk.FilesLocation = c.ImagePath
					cbk.FileData = fileData
					imageFiles = append(imageFiles, cbk)
				}
			}
		}
		bkfs.ImageFiles = &imageFiles
	}
}

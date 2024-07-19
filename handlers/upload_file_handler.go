package handlers

type File interface {
	UploadFile()
	DownloadFile()
}
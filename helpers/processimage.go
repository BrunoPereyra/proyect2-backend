package helpers

import (
	"backend/config"
	"context"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Processimage(fileHeader *multipart.FileHeader, PostImageChanel chan string, errChanel chan error) {
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()
		cldService, errcloudinary := cloudinary.NewFromURL(config.CLOUDINARY_URL())
		if errcloudinary != nil {
			errChanel <- errcloudinary
		}
		resp, errcldService := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		if errcldService != nil || !strings.HasPrefix(resp.SecureURL, "https://") {
			errChanel <- errcldService
		}
		PostImageChanel <- resp.SecureURL
	} else {
		PostImageChanel <- ""
	}
}

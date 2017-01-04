package cloudinary

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/HomesNZ/go-common/env"
	"github.com/nicday/go-cloudinary"
)

var (
	// Service is a singleton Cloudinary Service instance that be modified during testing.
	service *CDNService
)

func cloudinaryURI() string {
	return fmt.Sprintf(
		"cloudinary://%s:%s@%s",
		env.MustGetString("CLOUDINARY_API_KEY"),
		env.MustGetString("CLOUDINARY_API_SECRET"),
		env.MustGetString("CLOUDINARY_CLOUD_NAME"),
	)
}

// Service returns a Cloudinary Service singleton.
func Service() *CDNService {
	if service != nil {
		return service
	}

	c, err := cloudinary.Dial(cloudinaryURI())
	if err != nil {
		log.Fatal(err)
	}

	service = &CDNService{
		service: c,
	}

	return service
}

// CDNService is a Cloudinary concrete implementation of cdn.Interface
type CDNService struct {
	service *cloudinary.Service
}

// UploadImage uploads a new image to the Cloudinary CDN, with the supplied name and reader data. The public URL will be
// returned when successful, otherwise an error will be returned.
func (c CDNService) UploadImage(name string, data io.Reader) (string, error) {
	now := time.Now()
	fileName := fmt.Sprintf("%s_%d", name, now.UnixNano())

	_, err := c.service.UploadStaticImage(fileName, data, "")
	if err != nil {
		return "", err
	}

	return c.service.Url(fileName, cloudinary.ImageType), nil
}

// UploadURI sets the URI used when uploading assets to the CDN.
func (c *CDNService) UploadURI(uri string) error {
	return c.service.UploadURI(uri)
}

// RemoveAsset removes an asset from Cloudinary.
func (c CDNService) RemoveAsset(uri string) error {
	id, err := c.service.PublicID(uri)
	if err != nil {
		return err
	}

	return c.service.Delete(id, "", cloudinary.ImageType)
}

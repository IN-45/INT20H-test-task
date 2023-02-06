package cloudstorage

type Config struct {
	CloudCredentialsPath string `env:"CLOUD_CREDENTIALS_PATH" envDefault:"./config/cloud_credentials.json"`
	ImagesBucketName     string `env:"IMAGES_BUCKET_NAME" envDefault:"in_45_images_bucket"`
}

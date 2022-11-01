package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a GCP resource (Storage Bucket)
		_, err := storage.NewBucket(ctx, "my-bucket-0927", &storage.BucketArgs{
			Name:                     pulumi.String("my-bucket-0927"),
			Location:                 pulumi.String("ASIA-NORTHEAST1"),
			StorageClass:             pulumi.String("STANDARD"),
			UniformBucketLevelAccess: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		return nil
	})
}

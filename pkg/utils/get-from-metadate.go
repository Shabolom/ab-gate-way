package utils

import "google.golang.org/grpc/metadata"

func GetFromMetadata(md metadata.MD, key string) string {
	values := md.Get(key)
	if len(values) == 0 {
		return ""
	}

	return values[0]
}

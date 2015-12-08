// +build ignore

/*
 * Minio Go Library for Amazon S3 Compatible Cloud Storage (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"
	"time"

	"github.com/minio/minio-go"
)

func main() {
	config := minio.Config{
		Endpoint:        "https://s3.amazonaws.com",
		AccessKeyID:     "YOUR-ACCESS-KEY-HERE",
		SecretAccessKey: "YOUR-PASSWORD-HERE",
	}

	// Default is Signature Version 4. To enable Signature Version 2 do the following.
	// config.Signature = minio.SignatureV2

	s3Client, err := minio.New(config)
	if err != nil {
		log.Fatalln(err)
	}
	presignedURL, err := s3Client.PresignedPutObject("bucketName", "objectName", time.Duration(1000)*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)
}
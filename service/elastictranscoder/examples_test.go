package elastictranscoder_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/elastictranscoder"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleElasticTranscoder_CancelJob() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.CancelJobInput{
		ID: aws.String("Id"), // Required
	}
	resp, err := svc.CancelJob(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_CreateJob() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.CreateJobInput{
		Input: &elastictranscoder.JobInput{ // Required
			AspectRatio: aws.String("AspectRatio"),
			Container:   aws.String("JobContainer"),
			Encryption: &elastictranscoder.Encryption{
				InitializationVector: aws.String("ZeroTo255String"),
				Key:                  aws.String("Base64EncodedString"),
				KeyMD5:               aws.String("Base64EncodedString"),
				Mode:                 aws.String("EncryptionMode"),
			},
			FrameRate:  aws.String("FrameRate"),
			Interlaced: aws.String("Interlaced"),
			Key:        aws.String("Key"),
			Resolution: aws.String("Resolution"),
		},
		PipelineID: aws.String("Id"), // Required
		Output: &elastictranscoder.CreateJobOutput{
			AlbumArt: &elastictranscoder.JobAlbumArt{
				Artwork: []*elastictranscoder.Artwork{
					&elastictranscoder.Artwork{ // Required
						AlbumArtFormat: aws.String("JpgOrPng"),
						Encryption: &elastictranscoder.Encryption{
							InitializationVector: aws.String("ZeroTo255String"),
							Key:                  aws.String("Base64EncodedString"),
							KeyMD5:               aws.String("Base64EncodedString"),
							Mode:                 aws.String("EncryptionMode"),
						},
						InputKey:      aws.String("WatermarkKey"),
						MaxHeight:     aws.String("DigitsOrAuto"),
						MaxWidth:      aws.String("DigitsOrAuto"),
						PaddingPolicy: aws.String("PaddingPolicy"),
						SizingPolicy:  aws.String("SizingPolicy"),
					},
					// More values...
				},
				MergePolicy: aws.String("MergePolicy"),
			},
			Captions: &elastictranscoder.Captions{
				CaptionFormats: []*elastictranscoder.CaptionFormat{
					&elastictranscoder.CaptionFormat{ // Required
						Encryption: &elastictranscoder.Encryption{
							InitializationVector: aws.String("ZeroTo255String"),
							Key:                  aws.String("Base64EncodedString"),
							KeyMD5:               aws.String("Base64EncodedString"),
							Mode:                 aws.String("EncryptionMode"),
						},
						Format:  aws.String("CaptionFormatFormat"),
						Pattern: aws.String("CaptionFormatPattern"),
					},
					// More values...
				},
				CaptionSources: []*elastictranscoder.CaptionSource{
					&elastictranscoder.CaptionSource{ // Required
						Encryption: &elastictranscoder.Encryption{
							InitializationVector: aws.String("ZeroTo255String"),
							Key:                  aws.String("Base64EncodedString"),
							KeyMD5:               aws.String("Base64EncodedString"),
							Mode:                 aws.String("EncryptionMode"),
						},
						Key:        aws.String("Key"),
						Label:      aws.String("Name"),
						Language:   aws.String("Key"),
						TimeOffset: aws.String("TimeOffset"),
					},
					// More values...
				},
				MergePolicy: aws.String("CaptionMergePolicy"),
			},
			Composition: []*elastictranscoder.Clip{
				&elastictranscoder.Clip{ // Required
					TimeSpan: &elastictranscoder.TimeSpan{
						Duration:  aws.String("Time"),
						StartTime: aws.String("Time"),
					},
				},
				// More values...
			},
			Encryption: &elastictranscoder.Encryption{
				InitializationVector: aws.String("ZeroTo255String"),
				Key:                  aws.String("Base64EncodedString"),
				KeyMD5:               aws.String("Base64EncodedString"),
				Mode:                 aws.String("EncryptionMode"),
			},
			Key:             aws.String("Key"),
			PresetID:        aws.String("Id"),
			Rotate:          aws.String("Rotate"),
			SegmentDuration: aws.String("Float"),
			ThumbnailEncryption: &elastictranscoder.Encryption{
				InitializationVector: aws.String("ZeroTo255String"),
				Key:                  aws.String("Base64EncodedString"),
				KeyMD5:               aws.String("Base64EncodedString"),
				Mode:                 aws.String("EncryptionMode"),
			},
			ThumbnailPattern: aws.String("ThumbnailPattern"),
			Watermarks: []*elastictranscoder.JobWatermark{
				&elastictranscoder.JobWatermark{ // Required
					Encryption: &elastictranscoder.Encryption{
						InitializationVector: aws.String("ZeroTo255String"),
						Key:                  aws.String("Base64EncodedString"),
						KeyMD5:               aws.String("Base64EncodedString"),
						Mode:                 aws.String("EncryptionMode"),
					},
					InputKey:          aws.String("WatermarkKey"),
					PresetWatermarkID: aws.String("PresetWatermarkId"),
				},
				// More values...
			},
		},
		OutputKeyPrefix: aws.String("Key"),
		Outputs: []*elastictranscoder.CreateJobOutput{
			&elastictranscoder.CreateJobOutput{ // Required
				AlbumArt: &elastictranscoder.JobAlbumArt{
					Artwork: []*elastictranscoder.Artwork{
						&elastictranscoder.Artwork{ // Required
							AlbumArtFormat: aws.String("JpgOrPng"),
							Encryption: &elastictranscoder.Encryption{
								InitializationVector: aws.String("ZeroTo255String"),
								Key:                  aws.String("Base64EncodedString"),
								KeyMD5:               aws.String("Base64EncodedString"),
								Mode:                 aws.String("EncryptionMode"),
							},
							InputKey:      aws.String("WatermarkKey"),
							MaxHeight:     aws.String("DigitsOrAuto"),
							MaxWidth:      aws.String("DigitsOrAuto"),
							PaddingPolicy: aws.String("PaddingPolicy"),
							SizingPolicy:  aws.String("SizingPolicy"),
						},
						// More values...
					},
					MergePolicy: aws.String("MergePolicy"),
				},
				Captions: &elastictranscoder.Captions{
					CaptionFormats: []*elastictranscoder.CaptionFormat{
						&elastictranscoder.CaptionFormat{ // Required
							Encryption: &elastictranscoder.Encryption{
								InitializationVector: aws.String("ZeroTo255String"),
								Key:                  aws.String("Base64EncodedString"),
								KeyMD5:               aws.String("Base64EncodedString"),
								Mode:                 aws.String("EncryptionMode"),
							},
							Format:  aws.String("CaptionFormatFormat"),
							Pattern: aws.String("CaptionFormatPattern"),
						},
						// More values...
					},
					CaptionSources: []*elastictranscoder.CaptionSource{
						&elastictranscoder.CaptionSource{ // Required
							Encryption: &elastictranscoder.Encryption{
								InitializationVector: aws.String("ZeroTo255String"),
								Key:                  aws.String("Base64EncodedString"),
								KeyMD5:               aws.String("Base64EncodedString"),
								Mode:                 aws.String("EncryptionMode"),
							},
							Key:        aws.String("Key"),
							Label:      aws.String("Name"),
							Language:   aws.String("Key"),
							TimeOffset: aws.String("TimeOffset"),
						},
						// More values...
					},
					MergePolicy: aws.String("CaptionMergePolicy"),
				},
				Composition: []*elastictranscoder.Clip{
					&elastictranscoder.Clip{ // Required
						TimeSpan: &elastictranscoder.TimeSpan{
							Duration:  aws.String("Time"),
							StartTime: aws.String("Time"),
						},
					},
					// More values...
				},
				Encryption: &elastictranscoder.Encryption{
					InitializationVector: aws.String("ZeroTo255String"),
					Key:                  aws.String("Base64EncodedString"),
					KeyMD5:               aws.String("Base64EncodedString"),
					Mode:                 aws.String("EncryptionMode"),
				},
				Key:             aws.String("Key"),
				PresetID:        aws.String("Id"),
				Rotate:          aws.String("Rotate"),
				SegmentDuration: aws.String("Float"),
				ThumbnailEncryption: &elastictranscoder.Encryption{
					InitializationVector: aws.String("ZeroTo255String"),
					Key:                  aws.String("Base64EncodedString"),
					KeyMD5:               aws.String("Base64EncodedString"),
					Mode:                 aws.String("EncryptionMode"),
				},
				ThumbnailPattern: aws.String("ThumbnailPattern"),
				Watermarks: []*elastictranscoder.JobWatermark{
					&elastictranscoder.JobWatermark{ // Required
						Encryption: &elastictranscoder.Encryption{
							InitializationVector: aws.String("ZeroTo255String"),
							Key:                  aws.String("Base64EncodedString"),
							KeyMD5:               aws.String("Base64EncodedString"),
							Mode:                 aws.String("EncryptionMode"),
						},
						InputKey:          aws.String("WatermarkKey"),
						PresetWatermarkID: aws.String("PresetWatermarkId"),
					},
					// More values...
				},
			},
			// More values...
		},
		Playlists: []*elastictranscoder.CreateJobPlaylist{
			&elastictranscoder.CreateJobPlaylist{ // Required
				Format: aws.String("PlaylistFormat"),
				HLSContentProtection: &elastictranscoder.HLSContentProtection{
					InitializationVector:  aws.String("ZeroTo255String"),
					Key:                   aws.String("Base64EncodedString"),
					KeyMD5:                aws.String("Base64EncodedString"),
					KeyStoragePolicy:      aws.String("KeyStoragePolicy"),
					LicenseAcquisitionURL: aws.String("ZeroTo512String"),
					Method:                aws.String("HlsContentProtectionMethod"),
				},
				Name: aws.String("Filename"),
				OutputKeys: []*string{
					aws.String("Key"), // Required
					// More values...
				},
			},
			// More values...
		},
		UserMetadata: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateJob(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_CreatePipeline() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.CreatePipelineInput{
		InputBucket:  aws.String("BucketName"), // Required
		Name:         aws.String("Name"),       // Required
		Role:         aws.String("Role"),       // Required
		AWSKMSKeyARN: aws.String("KeyArn"),
		ContentConfig: &elastictranscoder.PipelineOutputConfig{
			Bucket: aws.String("BucketName"),
			Permissions: []*elastictranscoder.Permission{
				&elastictranscoder.Permission{ // Required
					Access: []*string{
						aws.String("AccessControl"), // Required
						// More values...
					},
					Grantee:     aws.String("Grantee"),
					GranteeType: aws.String("GranteeType"),
				},
				// More values...
			},
			StorageClass: aws.String("StorageClass"),
		},
		Notifications: &elastictranscoder.Notifications{
			Completed:   aws.String("SnsTopic"),
			Error:       aws.String("SnsTopic"),
			Progressing: aws.String("SnsTopic"),
			Warning:     aws.String("SnsTopic"),
		},
		OutputBucket: aws.String("BucketName"),
		ThumbnailConfig: &elastictranscoder.PipelineOutputConfig{
			Bucket: aws.String("BucketName"),
			Permissions: []*elastictranscoder.Permission{
				&elastictranscoder.Permission{ // Required
					Access: []*string{
						aws.String("AccessControl"), // Required
						// More values...
					},
					Grantee:     aws.String("Grantee"),
					GranteeType: aws.String("GranteeType"),
				},
				// More values...
			},
			StorageClass: aws.String("StorageClass"),
		},
	}
	resp, err := svc.CreatePipeline(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_CreatePreset() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.CreatePresetInput{
		Container: aws.String("PresetContainer"), // Required
		Name:      aws.String("Name"),            // Required
		Audio: &elastictranscoder.AudioParameters{
			BitRate:  aws.String("AudioBitRate"),
			Channels: aws.String("AudioChannels"),
			Codec:    aws.String("AudioCodec"),
			CodecOptions: &elastictranscoder.AudioCodecOptions{
				Profile: aws.String("AudioCodecProfile"),
			},
			SampleRate: aws.String("AudioSampleRate"),
		},
		Description: aws.String("Description"),
		Thumbnails: &elastictranscoder.Thumbnails{
			AspectRatio:   aws.String("AspectRatio"),
			Format:        aws.String("JpgOrPng"),
			Interval:      aws.String("Digits"),
			MaxHeight:     aws.String("DigitsOrAuto"),
			MaxWidth:      aws.String("DigitsOrAuto"),
			PaddingPolicy: aws.String("PaddingPolicy"),
			Resolution:    aws.String("ThumbnailResolution"),
			SizingPolicy:  aws.String("SizingPolicy"),
		},
		Video: &elastictranscoder.VideoParameters{
			AspectRatio: aws.String("AspectRatio"),
			BitRate:     aws.String("VideoBitRate"),
			Codec:       aws.String("VideoCodec"),
			CodecOptions: &map[string]*string{
				"Key": aws.String("CodecOption"), // Required
				// More values...
			},
			DisplayAspectRatio: aws.String("AspectRatio"),
			FixedGOP:           aws.String("FixedGOP"),
			FrameRate:          aws.String("FrameRate"),
			KeyframesMaxDist:   aws.String("KeyframesMaxDist"),
			MaxFrameRate:       aws.String("MaxFrameRate"),
			MaxHeight:          aws.String("DigitsOrAuto"),
			MaxWidth:           aws.String("DigitsOrAuto"),
			PaddingPolicy:      aws.String("PaddingPolicy"),
			Resolution:         aws.String("Resolution"),
			SizingPolicy:       aws.String("SizingPolicy"),
			Watermarks: []*elastictranscoder.PresetWatermark{
				&elastictranscoder.PresetWatermark{ // Required
					HorizontalAlign:  aws.String("HorizontalAlign"),
					HorizontalOffset: aws.String("PixelsOrPercent"),
					ID:               aws.String("PresetWatermarkId"),
					MaxHeight:        aws.String("PixelsOrPercent"),
					MaxWidth:         aws.String("PixelsOrPercent"),
					Opacity:          aws.String("Opacity"),
					SizingPolicy:     aws.String("WatermarkSizingPolicy"),
					Target:           aws.String("Target"),
					VerticalAlign:    aws.String("VerticalAlign"),
					VerticalOffset:   aws.String("PixelsOrPercent"),
				},
				// More values...
			},
		},
	}
	resp, err := svc.CreatePreset(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_DeletePipeline() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.DeletePipelineInput{
		ID: aws.String("Id"), // Required
	}
	resp, err := svc.DeletePipeline(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_DeletePreset() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.DeletePresetInput{
		ID: aws.String("Id"), // Required
	}
	resp, err := svc.DeletePreset(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_ListJobsByPipeline() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.ListJobsByPipelineInput{
		PipelineID: aws.String("Id"), // Required
		Ascending:  aws.String("Ascending"),
		PageToken:  aws.String("Id"),
	}
	resp, err := svc.ListJobsByPipeline(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_ListJobsByStatus() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.ListJobsByStatusInput{
		Status:    aws.String("JobStatus"), // Required
		Ascending: aws.String("Ascending"),
		PageToken: aws.String("Id"),
	}
	resp, err := svc.ListJobsByStatus(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_ListPipelines() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.ListPipelinesInput{
		Ascending: aws.String("Ascending"),
		PageToken: aws.String("Id"),
	}
	resp, err := svc.ListPipelines(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_ListPresets() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.ListPresetsInput{
		Ascending: aws.String("Ascending"),
		PageToken: aws.String("Id"),
	}
	resp, err := svc.ListPresets(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_ReadJob() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.ReadJobInput{
		ID: aws.String("Id"), // Required
	}
	resp, err := svc.ReadJob(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_ReadPipeline() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.ReadPipelineInput{
		ID: aws.String("Id"), // Required
	}
	resp, err := svc.ReadPipeline(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_ReadPreset() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.ReadPresetInput{
		ID: aws.String("Id"), // Required
	}
	resp, err := svc.ReadPreset(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_TestRole() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.TestRoleInput{
		InputBucket:  aws.String("BucketName"), // Required
		OutputBucket: aws.String("BucketName"), // Required
		Role:         aws.String("Role"),       // Required
		Topics: []*string{ // Required
			aws.String("SnsTopic"), // Required
			// More values...
		},
	}
	resp, err := svc.TestRole(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_UpdatePipeline() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.UpdatePipelineInput{
		ID:           aws.String("Id"), // Required
		AWSKMSKeyARN: aws.String("KeyArn"),
		ContentConfig: &elastictranscoder.PipelineOutputConfig{
			Bucket: aws.String("BucketName"),
			Permissions: []*elastictranscoder.Permission{
				&elastictranscoder.Permission{ // Required
					Access: []*string{
						aws.String("AccessControl"), // Required
						// More values...
					},
					Grantee:     aws.String("Grantee"),
					GranteeType: aws.String("GranteeType"),
				},
				// More values...
			},
			StorageClass: aws.String("StorageClass"),
		},
		InputBucket: aws.String("BucketName"),
		Name:        aws.String("Name"),
		Notifications: &elastictranscoder.Notifications{
			Completed:   aws.String("SnsTopic"),
			Error:       aws.String("SnsTopic"),
			Progressing: aws.String("SnsTopic"),
			Warning:     aws.String("SnsTopic"),
		},
		Role: aws.String("Role"),
		ThumbnailConfig: &elastictranscoder.PipelineOutputConfig{
			Bucket: aws.String("BucketName"),
			Permissions: []*elastictranscoder.Permission{
				&elastictranscoder.Permission{ // Required
					Access: []*string{
						aws.String("AccessControl"), // Required
						// More values...
					},
					Grantee:     aws.String("Grantee"),
					GranteeType: aws.String("GranteeType"),
				},
				// More values...
			},
			StorageClass: aws.String("StorageClass"),
		},
	}
	resp, err := svc.UpdatePipeline(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_UpdatePipelineNotifications() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.UpdatePipelineNotificationsInput{
		ID: aws.String("Id"), // Required
		Notifications: &elastictranscoder.Notifications{ // Required
			Completed:   aws.String("SnsTopic"),
			Error:       aws.String("SnsTopic"),
			Progressing: aws.String("SnsTopic"),
			Warning:     aws.String("SnsTopic"),
		},
	}
	resp, err := svc.UpdatePipelineNotifications(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleElasticTranscoder_UpdatePipelineStatus() {
	svc := elastictranscoder.New(nil)

	params := &elastictranscoder.UpdatePipelineStatusInput{
		ID:     aws.String("Id"),             // Required
		Status: aws.String("PipelineStatus"), // Required
	}
	resp, err := svc.UpdatePipelineStatus(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
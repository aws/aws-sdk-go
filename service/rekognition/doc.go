// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rekognition provides the client and types for making API
// requests to Amazon Rekognition.
//
// This is the API Reference for Amazon Rekognition Image (https://docs.aws.amazon.com/rekognition/latest/dg/images.html),
// Amazon Rekognition Custom Labels (https://docs.aws.amazon.com/rekognition/latest/customlabels-dg/what-is.html),
// Amazon Rekognition Stored Video (https://docs.aws.amazon.com/rekognition/latest/dg/video.html),
// Amazon Rekognition Streaming Video (https://docs.aws.amazon.com/rekognition/latest/dg/streaming-video.html).
// It provides descriptions of actions, data types, common parameters, and common
// errors.
//
// Amazon Rekognition Image
//
//   - AssociateFaces (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_AssociateFaces.html)
//
//   - CompareFaces (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CompareFaces.html)
//
//   - CreateCollection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateCollection.html)
//
//   - CreateUser (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateUser.html)
//
//   - DeleteCollection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DeleteCollection.html)
//
//   - DeleteFaces (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DeleteFaces.html)
//
//   - DeleteUser (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DeleteUser.html)
//
//   - DescribeCollection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DescribeCollection.html)
//
//   - DetectFaces (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DetectFaces.html)
//
//   - DetectLabels (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DetectLabels.html)
//
//   - DetectModerationLabels (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DetectModerationLabels.html)
//
//   - DetectProtectiveEquipment (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DetectProtectiveEquipment.html)
//
//   - DetectText (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DetectText.html)
//
//   - DisassociateFaces (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DisassociateFaces.html)
//
//   - GetCelebrityInfo (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetCelebrityInfo.html)
//
//   - GetMediaAnalysisJob (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetMediaAnalysisJob.html)
//
//   - IndexFaces (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_IndexFaces.html)
//
//   - ListCollections (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_ListCollections.html)
//
//   - ListMediaAnalysisJob (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_ListMediaAnalysisJob.html)
//
//   - ListFaces (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_ListFaces.html)
//
//   - ListUsers (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_ListFaces.html)
//
//   - RecognizeCelebrities (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_RecognizeCelebrities.html)
//
//   - SearchFaces (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_SearchFaces.html)
//
//   - SearchFacesByImage (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_SearchFacesByImage.html)
//
//   - SearchUsers (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_SearchUsers.html)
//
//   - SearchUsersByImage (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_SearchUsersByImage.html)
//
//   - StartMediaAnalysisJob (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartMediaAnalysisJob.html)
//
// Amazon Rekognition Custom Labels
//
//   - CopyProjectVersion (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CopyProjectVersion.html)
//
//   - CreateDataset (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateDataset.html)
//
//   - CreateProject (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateProject.html)
//
//   - CreateProjectVersion (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateProjectVersion.html)
//
//   - DeleteDataset (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DeleteDataset.html)
//
//   - DeleteProject (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DeleteProject.html)
//
//   - DeleteProjectPolicy (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DeleteProjectPolicy.html)
//
//   - DeleteProjectVersion (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DeleteProjectVersion.html)
//
//   - DescribeDataset (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DescribeDataset.html)
//
//   - DescribeProjects (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DescribeProjects.html)
//
//   - DescribeProjectVersions (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DescribeProjectVersions.html)
//
//   - DetectCustomLabels (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DetectCustomLabels.html)
//
//   - DistributeDatasetEntries (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DistributeDatasetEntries.html)
//
//   - ListDatasetEntries (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_ListDatasetEntries.html)
//
//   - ListDatasetLabels (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_ListDatasetLabels.html)
//
//   - ListProjectPolicies (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_ListProjectPolicies.html)
//
//   - PutProjectPolicy (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_PutProjectPolicy.html)
//
//   - StartProjectVersion (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartProjectVersion.html)
//
//   - StopProjectVersion (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StopProjectVersion.html)
//
//   - UpdateDatasetEntries (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_UpdateDatasetEntries.html)
//
// Amazon Rekognition Video Stored Video
//
//   - GetCelebrityRecognition (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetCelebrityRecognition.html)
//
//   - GetContentModeration (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetContentModeration.html)
//
//   - GetFaceDetection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetFaceDetection.html)
//
//   - GetFaceSearch (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetFaceSearch.html)
//
//   - GetLabelDetection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetLabelDetection.html)
//
//   - GetPersonTracking (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetPersonTracking.html)
//
//   - GetSegmentDetection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetSegmentDetection.html)
//
//   - GetTextDetection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_GetTextDetection.html)
//
//   - StartCelebrityRecognition (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartCelebrityRecognition.html)
//
//   - StartContentModeration (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartContentModeration.html)
//
//   - StartFaceDetection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartFaceDetection.html)
//
//   - StartFaceSearch (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartFaceSearch.html)
//
//   - StartLabelDetection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartLabelDetection.html)
//
//   - StartPersonTracking (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartPersonTracking.html)
//
//   - StartSegmentDetection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartSegmentDetection.html)
//
//   - StartTextDetection (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartTextDetection.html)
//
// Amazon Rekognition Video Streaming Video
//
//   - CreateStreamProcessor (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor.html)
//
//   - DeleteStreamProcessor (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DeleteStreamProcessor.html)
//
//   - DescribeStreamProcessor (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_DescribeStreamProcessor.html)
//
//   - ListStreamProcessors (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_ListStreamProcessors.html)
//
//   - StartStreamProcessor (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StartStreamProcessor.html)
//
//   - StopStreamProcessor (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StopStreamProcessor.html)
//
//   - UpdateStreamProcessor (https://docs.aws.amazon.com/rekognition/latest/APIReference/API_UpdateStreamProcessor.html)
//
// See rekognition package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/rekognition/
//
// # Using the Client
//
// To contact Amazon Rekognition with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Rekognition client Rekognition for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/rekognition/#New
package rekognition

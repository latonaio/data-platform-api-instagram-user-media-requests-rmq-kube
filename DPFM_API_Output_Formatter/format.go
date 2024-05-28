package dpfm_api_output_formatter

func ConvertToInstagramUserMediaRequestsFromResponse(
	instagramUserMediaRequestsResponseBody InstagramUserMediaResponseBody,
) InstagramUserMediaResponse {
	var instagramUserMediaRequestsResponse InstagramUserMediaResponse

	instagramUserMediaRequestsResponse.InstagramID 						= instagramUserMediaRequestsResponseBody.ID
	instagramUserMediaRequestsResponse.InstagramMediaType				= instagramUserMediaRequestsResponseBody.MediaType
	instagramUserMediaRequestsResponse.InstagramMediaURL				= instagramUserMediaRequestsResponseBody.MediaURL
	instagramUserMediaRequestsResponse.InstagramMediaVideoThumbnailURL	= instagramUserMediaRequestsResponseBody.ThumbnailURL
	instagramUserMediaRequestsResponse.InstagramMediaTimeStamp			= instagramUserMediaRequestsResponseBody.TimeStamp
	instagramUserMediaRequestsResponse.InstagramMediaUserName			= instagramUserMediaRequestsResponseBody.UserName

	return instagramUserMediaRequestsResponse
}
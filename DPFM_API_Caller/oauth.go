package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-instagram-user-media-requests-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-instagram-user-media-requests-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-instagram-user-media-requests-rmq-kube/config"
	"encoding/json"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
)

func (c *DPFMAPICaller) InstagramUserMedia(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) *[]dpfm_api_output_formatter.InstagramUserMediaResponse {
	var instagramUserMedia []dpfm_api_output_formatter.InstagramUserMediaResponse

	accessToken := input.InstagramUserMedia.AccessToken

	userMediaBaseURL := conf.OAuth.UserMediaURL
	userMediaURL := fmt.Sprintf(
		"%s?access_token=%s&fields=id,username",
		userMediaBaseURL,
		accessToken,
	)

	req, err := http.NewRequest("GET", userMediaURL, nil)

	if err != nil {
		*errs = append(*errs, xerrors.Errorf("NewRequest error: %d", err))
		return nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("User media request error: %d", err))
		return nil
	}
	defer resp.Body.Close()

	userMediaBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("User media request response read error: %d", err))
		return nil
	}

	var response map[string]interface{}
	err = json.Unmarshal(userMediaBody, &response)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("Response response error: %d", err))
		return nil
	}

	errorObj, ok := response["error"].(map[string]interface{})
	if ok {
		code, ok := errorObj["code"].(float64)
		if ok {
			errMsg, _ := errorObj["message"].(string)
			*errs = append(*errs, xerrors.Errorf("Status code error: %v %v", code, errMsg))
			return nil
		}
	}

	var instagramUserMediaResponseBody dpfm_api_output_formatter.InstagramUserMediaResponseBody
	err = json.Unmarshal(userMediaBody, &instagramUserMediaResponseBody)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("User media request response unmarshal error: %d", err))
		return nil
	}

	userMedia := dpfm_api_output_formatter.ConvertToInstagramUserMediaRequestsFromResponse(instagramUserMediaResponseBody)

	instagramUserMedia = append(
		instagramUserMedia,
		userMedia,
	)

	return &instagramUserMedia
}

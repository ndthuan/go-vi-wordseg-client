package apiv1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const segmentationEndpoint string = "/v1/segmenting"
const taggingEndpoint string = "/v1/tagging"

type request struct {
	Text string `json:"text"`
}

//SegmentationResult holds result of a segmentation request
type SegmentationResult struct {
	Sentences []string `json:"sentences"`
}

//Word holds a segmented word with extra tagging labels
type Word struct {
	Form string `json:"form"` // The segmented word/phrase, might contain underscores
	Pos  string `json:"pos"`  // Part-of-Speech tag, noun or verb etc...
	Ner  string `json:"ner"`  // Named Entity Recognition tag
	Dep  string `json:"dep"`  // Dependency label, role of the word in the sentences
}

//TaggingResult holds result of a tagging request
type TaggingResult struct {
	Sentences [][]Word `json:"sentences"`
}

//Client represents the client
type Client struct {
	client *http.Client
	URL    string
}

//Segment segments a string into meaningful phrases
func (c *Client) Segment(text string, shouldSkipPunct bool) (*SegmentationResult, error) {
	reqBody, err := json.Marshal(&request{Text: text})
	if err != nil {
		return nil, fmt.Errorf("unable to encode request: %v", err)
	}

	skipPunct := 0
	if shouldSkipPunct {
		skipPunct = 1
	}

	url := fmt.Sprintf("%s%s?skipPunct=%d", c.URL, segmentationEndpoint, skipPunct)

	resBody, err := c.sendJSONRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var result SegmentationResult

	if err = json.Unmarshal(resBody, &result); err != nil {
		return nil, fmt.Errorf("unable to decode response: %v", err)
	}

	return &result, nil
}

//Tag segments a string into meaningful phrases, also labels their responsibilities in the sentences
func (c *Client) Tag(text string) (*TaggingResult, error) {
	reqBody, err := json.Marshal(&request{Text: text})
	if err != nil {
		return nil, fmt.Errorf("unable to encode request: %v", err)
	}

	url := fmt.Sprintf("%s%s", c.URL, taggingEndpoint)

	resBody, err := c.sendJSONRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var result TaggingResult

	if err = json.Unmarshal(resBody, &result); err != nil {
		return nil, fmt.Errorf("unable to decode response: %v", err)
	}

	return &result, nil
}

//NewClient creates a new service client
func NewClient(URL string) *Client {
	c := &Client{URL: URL}
	c.client = &http.Client{}

	return c
}

func (c *Client) sendJSONRequest(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to do request: %v", err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %v", err)
	}

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response: %s", resBody)
	}

	return resBody, nil
}

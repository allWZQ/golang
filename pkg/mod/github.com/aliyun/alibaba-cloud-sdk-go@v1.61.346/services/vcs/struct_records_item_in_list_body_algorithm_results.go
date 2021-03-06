package vcs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// RecordsItemInListBodyAlgorithmResults is a nested struct in vcs response
type RecordsItemInListBodyAlgorithmResults struct {
	CapStyle         string  `json:"CapStyle" xml:"CapStyle"`
	CorpId           string  `json:"CorpId" xml:"CorpId"`
	DataSourceId     string  `json:"DataSourceId" xml:"DataSourceId"`
	PersonId         string  `json:"PersonId" xml:"PersonId"`
	GenderCode       string  `json:"GenderCode" xml:"GenderCode"`
	HairStyle        string  `json:"HairStyle" xml:"HairStyle"`
	LeftTopX         float64 `json:"LeftTopX" xml:"LeftTopX"`
	LeftTopY         float64 `json:"LeftTopY" xml:"LeftTopY"`
	MaxAge           string  `json:"MaxAge" xml:"MaxAge"`
	MinAge           string  `json:"MinAge" xml:"MinAge"`
	PicUrlPath       string  `json:"PicUrlPath" xml:"PicUrlPath"`
	RightBottomX     float64 `json:"RightBottomX" xml:"RightBottomX"`
	RightBottomY     float64 `json:"RightBottomY" xml:"RightBottomY"`
	ShotTime         string  `json:"ShotTime" xml:"ShotTime"`
	TargetPicUrlPath string  `json:"TargetPicUrlPath" xml:"TargetPicUrlPath"`
	CoatLength       string  `json:"CoatLength" xml:"CoatLength"`
	CoatStyle        string  `json:"CoatStyle" xml:"CoatStyle"`
	TrousersLength   string  `json:"TrousersLength" xml:"TrousersLength"`
	TrousersStyle    string  `json:"TrousersStyle" xml:"TrousersStyle"`
	CoatColor        string  `json:"CoatColor" xml:"CoatColor"`
	TrousersColor    string  `json:"TrousersColor" xml:"TrousersColor"`
}

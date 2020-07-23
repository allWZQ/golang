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

// Data is a nested struct in vcs response
type Data struct {
	TotalPage    int                `json:"TotalPage" xml:"TotalPage"`
	DataSourceId string             `json:"DataSourceId" xml:"DataSourceId"`
	MaxId        string             `json:"MaxId" xml:"MaxId"`
	PersonId     string             `json:"PersonId" xml:"PersonId"`
	KafkaTopic   string             `json:"KafkaTopic" xml:"KafkaTopic"`
	PageSize     int                `json:"PageSize" xml:"PageSize"`
	OssPath      string             `json:"OssPath" xml:"OssPath"`
	TaskId       string             `json:"TaskId" xml:"TaskId"`
	TotalCount   int                `json:"TotalCount" xml:"TotalCount"`
	QualityScore string             `json:"QualityScore" xml:"QualityScore"`
	PageNumber   int                `json:"PageNumber" xml:"PageNumber"`
	PicUrl       string             `json:"PicUrl" xml:"PicUrl"`
	Description  string             `json:"Description" xml:"Description"`
	Attributes   Attributes         `json:"Attributes" xml:"Attributes"`
	Records      []RecordsItem      `json:"Records" xml:"Records"`
	ResultObject []ResultObjectItem `json:"ResultObject" xml:"ResultObject"`
	TagList      []TagListItem      `json:"TagList" xml:"TagList"`
	FaceList     []Face             `json:"FaceList" xml:"FaceList"`
	BodyList     []Body             `json:"BodyList" xml:"BodyList"`
}

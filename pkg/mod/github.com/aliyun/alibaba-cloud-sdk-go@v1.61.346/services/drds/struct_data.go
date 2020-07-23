package drds

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

// Data is a nested struct in drds response
type Data struct {
	DbShardingFunction   string                  `json:"DbShardingFunction" xml:"DbShardingFunction"`
	OrderId              int64                   `json:"OrderId" xml:"OrderId"`
	NewestVersion        string                  `json:"NewestVersion" xml:"NewestVersion"`
	CreateTime           string                  `json:"CreateTime" xml:"CreateTime"`
	Mode                 string                  `json:"Mode" xml:"Mode"`
	IsShard              bool                    `json:"IsShard" xml:"IsShard"`
	DbComputeLength      int                     `json:"DbComputeLength" xml:"DbComputeLength"`
	InstRole             string                  `json:"InstRole" xml:"InstRole"`
	TbRightShiftOffset   int                     `json:"TbRightShiftOffset" xml:"TbRightShiftOffset"`
	ShardTbKey           string                  `json:"ShardTbKey" xml:"ShardTbKey"`
	Expired              string                  `json:"Expired" xml:"Expired"`
	IsActive             bool                    `json:"IsActive" xml:"IsActive"`
	Schema               string                  `json:"Schema" xml:"Schema"`
	DbInstType           string                  `json:"DbInstType" xml:"DbInstType"`
	IsSyncType           bool                    `json:"IsSyncType" xml:"IsSyncType"`
	SourceTableName      string                  `json:"SourceTableName" xml:"SourceTableName"`
	TbComputeLength      int                     `json:"TbComputeLength" xml:"TbComputeLength"`
	ShardDbKey           string                  `json:"ShardDbKey" xml:"ShardDbKey"`
	TableName            string                  `json:"TableName" xml:"TableName"`
	DbName               string                  `json:"DbName" xml:"DbName"`
	Stage                string                  `json:"Stage" xml:"Stage"`
	DbRightShiftOffset   int                     `json:"DbRightShiftOffset" xml:"DbRightShiftOffset"`
	Progress             string                  `json:"Progress" xml:"Progress"`
	InstanceVersion      string                  `json:"InstanceVersion" xml:"InstanceVersion"`
	TbShardingFunction   string                  `json:"TbShardingFunction" xml:"TbShardingFunction"`
	RandomCode           string                  `json:"RandomCode" xml:"RandomCode"`
	TargetTableName      string                  `json:"TargetTableName" xml:"TargetTableName"`
	TbPartitions         int                     `json:"TbPartitions" xml:"TbPartitions"`
	Msg                  string                  `json:"Msg" xml:"Msg"`
	Status               string                  `json:"Status" xml:"Status"`
	DbShardingColumnList DbShardingColumnList    `json:"DbShardingColumnList" xml:"DbShardingColumnList"`
	TbShardingColumnList TbShardingColumnList    `json:"TbShardingColumnList" xml:"TbShardingColumnList"`
	DrdsInstanceIdList   DrdsInstanceIdList      `json:"DrdsInstanceIdList" xml:"DrdsInstanceIdList"`
	FullRevise           FullRevise              `json:"FullRevise" xml:"FullRevise"`
	Increment            Increment               `json:"Increment" xml:"Increment"`
	FullCheck            FullCheck               `json:"FullCheck" xml:"FullCheck"`
	Full                 Full                    `json:"Full" xml:"Full"`
	Review               Review                  `json:"Review" xml:"Review"`
	List                 ListInDescribeHotDbList `json:"List" xml:"List"`
	ColumnList           []Column                `json:"ColumnList" xml:"ColumnList"`
}

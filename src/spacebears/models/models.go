package models

type BucketMetadata struct {
	Credentials []BucketCredentials `json:"credentials"`
	//Context map[string]interface{} `json:"context"`
}

type BucketCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

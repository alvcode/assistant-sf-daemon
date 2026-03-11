package dto

type Config struct {
	DriveDomain       string `json:"assistant_domain"`
	AssistantLogin    string `json:"assistant_login"`
	AssistantPassword string `json:"assistant_password"`
	FolderPath        string `json:"folder_path"`
}

package dict

const AppName = "ast-sync-folder"
const (
	StructTypeFolder = 0
	StructTypeFile   = 1
)
const ChunkSize = 45 * 1024 * 1024

const DBName = "app.db"
const (
	ConfigKeyDomain       = "drive_domain"
	ConfigKeyToken        = "token"
	ConfigKeyRefreshToken = "refresh_token"
	ConfigKeyFolderPath   = "folder_path"
)

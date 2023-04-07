package conts

import (
	"time"
)

const (
	AccessTokenExpiredSeconds = 24 * 3600
	RefreshTokenExpiredDays   = 30
	TokenAccessCachePrefix    = "admin_access_token_"
	TokenRefreshCachePrefix   = "admin_refresh_token_"
	AdminTokenHeaderName      = "Admin-Authorization"
	AuthorizedUser            = "authorized_user"
	CodePrefix                = "code_"
	CodeValidDuration         = time.Second
	OneTimeTokenQueryName     = "ott"
	SessionID                 = "session_id"
	AccessPermissionKeyPrefix = "access_permission_"
)

const (
	SonicBackupPrefix         = "sonic-backup-"
	SonicDataExportPrefix     = "sonic-data-export-"
	SonicBackupMarkdownPrefix = "sonic-backup-markdown-"
	SonicDefaultTagColor      = "#cfd3d7"
	SonicUploadDir            = "upload"
	SonicDefaultThemeDirName  = "default-theme-anatole"
)

var (
	ThemePropertyFilenames = [2]string{"theme.yaml", "theme.yml"}
	ThemeSettingFilenames  = [2]string{"settings.yaml", "settings.yml"}
)

const (
	DefaultThemeID         = "caicai_anatole"
	ThemeScreenshotsName   = "screenshot"
	ThemeCustomSheetPrefix = "sheet_"
	ThemeCustomPostPrefix  = "post_"
)

var (
	StartTime       time.Time = time.Now()
	DatabaseVersion string
	SonicVersion    = "v1.0.0"
	BuildTime       string
	BuildCommit     string
)

// 公式
const (
	VERIFY_EXP_USERNAME = `^[a-zA-Z0-9_]{3,15}$`
	VERIFY_EXP_PASSWORD = `^[a-zA-Z0-9_\.\&\@]{6,16}$`
)
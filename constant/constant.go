package constant

import "time"

const (
	CLIGRPCAddress             = "/tmp/simple-panel.sock"
	UserCLIGRPCAddress         = "/var/user_cli/cli.sock"
	HTTPAddress                = "127.0.0.1:28091"
	GrpcProtocol               = "unix"
	MinUnixID                  = 1000
	DefaultAutoSSLPath         = "/var/spanel/autossl"
	DefaultLSWSHomePath        = "/usr/local/lsws/DEFAULT/html"
	PanelDataPath              = "/var/spanel"
	BackupVersion              = "1.0.0"
	FullBackupDir              = "/backup"
	IncBackupDir               = "/inc_backup"
	DefaultLswsRestartInterval = 10 * time.Second
	MaxUserPerDB               = 10
	MaxACLsPerUser             = 10
	Unlimited                  = -1
	MaxRetryPerDay             = 5
	DayDuration                = 24 * time.Hour
	MaxDownloadFileSize        = 10 * 1024 * 1024 * 1024 // 10GB
	MaxThumbnailSize           = 1024 * 1024 * 50        // 50MB
	MaxPreviewSize             = 1024 * 1024 * 100       // 100MB
	MaxUpdateFileSize          = 1024 * 1024 * 10        // 10MB
	MaxUploadFileSize          = 10 * 1024 * 1024 * 1024 // 10GB
	MaxCronJobsPerUser         = 100
	RootUser                   = "root"
	CSFPanelPath               = "/var/lib/csf/csf.pl"
	MaxTraversalFile           = 3000
	TrashFolder                = "trash"
	TrashRestore               = ".trash_restore"
	RcloneExcludeFile          = "/root/.config/rclone/exclude.txt"
	RcloneExcludeRestore       = "/root/.config/rclone/exclude_restore.txt"
	RcloneConfigFile           = "/root/.config/rclone/rclone.conf"
	MaxBufferSizeBytes         = 512
	Redis                      = "redis"
	Memcached                  = "memcached"
	CronTab                    = "crontab"
	MetaFileName               = "meta.json"
	MaxFilesInAVWorker         = 10000
	MaxKeepBackupDays          = 31
	TeleChannelName            = "@zema_update"
	SSLPort                    = 443
)

// refresh csrf token when in get user

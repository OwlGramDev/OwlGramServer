package consts

import "OwlGramServer/github_bot/types"

const BasePathOwlGram = "/home/owlgram/"
const BaseWWWPath = "/var/www/"
const WebServerPath = BasePathOwlGram + "app.owlgram.org/"
const ServerFilesFolder = BaseWWWPath + "files.owlgram.org/"
const WebsiteHome = BaseWWWPath + "owlgram.org/"
const LibsPath = WebServerPath + "libs/"

const CacheFolder = WebServerPath + "cache/"
const CacheFolderBundles = CacheFolder + "bundles/"
const CacheFolderExtractedApks = CacheFolder + "extracted_apks/"
const UploadsFolder = CacheFolder + "uploads/"
const ProGuardFolder = CacheFolder + "proguard/"
const ConfigFilesFolder = WebServerPath + "config/"
const WebAppFiles = ServerFilesFolder + "webapp/"
const DebugFilesFolder = CacheFolder + "test/"
const NotFoundFile = WebsiteHome + "404.html"
const ForbiddenFile = WebsiteHome + "403.html"

const CrowdinApiLink = "https://owlgram.crowdin.com/api/v2/projects/"
const CrowdinProjectId = "8"

const GoogleApiJsonPath = ConfigFilesFolder + "login.json"
const GooglePlayConsoleLink = "https://play.google.com/console/u/0/developers/5671039640982520074/app/4974658660264210304/app-dashboard"

const OwlGramTGChannelLink = "https://t.me/s/OwlGram"
const OwlGramTGChannelBetaLink = "https://t.me/s/OwlGramBeta"
const UpdateFileDescription = ConfigFilesFolder + "update.json"
const OwlGramFilesServer = "https://files.owlgram.org/"

const DcCheckerCacheFile = CacheFolder + "dc_checker_cache.json"
const ReviewsCacheFile = CacheFolder + "reviews_cache.json"
const TelegramSessionFile = CacheFolder + "tg_session.dat"

var AliasSupported = []string{"/", ";", ".", "+", "!"}

const StaffGroupID = -1001267698171
const Tappo03UserID = 225117913
const BotImagesCache = CacheFolder + "bot_images_cache/"
const WebAppAuthTimeout = 3600

var GithubGroups = []types.Group{
	{
		ID:              -1001631666002, // OwlGramIT
		ForumID:         43580,
		AllowedBranches: []string{"master"},
	},
	{
		ID:              -1001672748705, // OwlGram
		ForumID:         38534,
		AllowedBranches: []string{"master"},
	},
	{
		ID:              -1001703872347, // Internal testing
		AllowedBranches: []string{"master", "develop"},
	},
}

const GithubRepo = "OwlGram"
const GithubRepoOwner = "OwlGramDev"
const GithubPemFile = ConfigFilesFolder + "owlgram-deploy.2022-01-07.private-key.pem"
const AndroidPackageName = "it.owlgram.android"

const BundleToolPath = LibsPath + "bundletool-all-1.9.1.jar"
const AAPT2ToolPath = LibsPath + "aapt2"
const RetraceToolPath = LibsPath + "proguard-retrace-6.0.3.jar"
const PythonLibApkSenderPath = LibsPath + "apks_sender"

const AppGalleryApi = "https://connect-api.cloud.huawei.com/api/"
const AppGalleryAppID = 105849965

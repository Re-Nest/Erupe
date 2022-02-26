package mhfpacket

import "github.com/Solenataris/Erupe/network"

// FromOpcode gets a packet struct that fulfills the MHFPacket interface by it's opcode.
func FromOpcode(opcode network.PacketID) MHFPacket {
	switch opcode {
	case network.MSG_HEAD:
		return &MsgHead{}
	case network.MSG_SYS_reserve01:
		return &MsgSysReserve01{}
	case network.MSG_SYS_reserve02:
		return &MsgSysReserve02{}
	case network.MSG_SYS_reserve03:
		return &MsgSysReserve03{}
	case network.MSG_SYS_reserve04:
		return &MsgSysReserve04{}
	case network.MSG_SYS_reserve05:
		return &MsgSysReserve05{}
	case network.MSG_SYS_reserve06:
		return &MsgSysReserve06{}
	case network.MSG_SYS_reserve07:
		return &MsgSysReserve07{}
	case network.MSG_SYS_ADD_OBJECT:
		return &MsgSysAddObject{}
	case network.MSG_SYS_DEL_OBJECT:
		return &MsgSysDelObject{}
	case network.MSG_SYS_DISP_OBJECT:
		return &MsgSysDispObject{}
	case network.MSG_SYS_HIDE_OBJECT:
		return &MsgSysHideObject{}
	case network.MSG_SYS_reserve0C:
		return &MsgSysReserve0C{}
	case network.MSG_SYS_reserve0D:
		return &MsgSysReserve0D{}
	case network.MSG_SYS_reserve0E:
		return &MsgSysReserve0E{}
	case network.MSG_SYS_EXTEND_THRESHOLD:
		return &MsgSysExtendThreshold{}
	case network.MSG_SYS_END:
		return &MsgSysEnd{}
	case network.MSG_SYS_NOP:
		return &MsgSysNop{}
	case network.MSG_SYS_ACK:
		return &MsgSysAck{}
	case network.MSG_SYS_TERMINAL_LOG:
		return &MsgSysTerminalLog{}
	case network.MSG_SYS_LOGIN:
		return &MsgSysLogin{}
	case network.MSG_SYS_LOGOUT:
		return &MsgSysLogout{}
	case network.MSG_SYS_SET_STATUS:
		return &MsgSysSetStatus{}
	case network.MSG_SYS_PING:
		return &MsgSysPing{}
	case network.MSG_SYS_CAST_BINARY:
		return &MsgSysCastBinary{}
	case network.MSG_SYS_HIDE_CLIENT:
		return &MsgSysHideClient{}
	case network.MSG_SYS_TIME:
		return &MsgSysTime{}
	case network.MSG_SYS_CASTED_BINARY:
		return &MsgSysCastedBinary{}
	case network.MSG_SYS_GET_FILE:
		return &MsgSysGetFile{}
	case network.MSG_SYS_ISSUE_LOGKEY:
		return &MsgSysIssueLogkey{}
	case network.MSG_SYS_RECORD_LOG:
		return &MsgSysRecordLog{}
	case network.MSG_SYS_ECHO:
		return &MsgSysEcho{}
	case network.MSG_SYS_CREATE_STAGE:
		return &MsgSysCreateStage{}
	case network.MSG_SYS_STAGE_DESTRUCT:
		return &MsgSysStageDestruct{}
	case network.MSG_SYS_ENTER_STAGE:
		return &MsgSysEnterStage{}
	case network.MSG_SYS_BACK_STAGE:
		return &MsgSysBackStage{}
	case network.MSG_SYS_MOVE_STAGE:
		return &MsgSysMoveStage{}
	case network.MSG_SYS_LEAVE_STAGE:
		return &MsgSysLeaveStage{}
	case network.MSG_SYS_LOCK_STAGE:
		return &MsgSysLockStage{}
	case network.MSG_SYS_UNLOCK_STAGE:
		return &MsgSysUnlockStage{}
	case network.MSG_SYS_RESERVE_STAGE:
		return &MsgSysReserveStage{}
	case network.MSG_SYS_UNRESERVE_STAGE:
		return &MsgSysUnreserveStage{}
	case network.MSG_SYS_SET_STAGE_PASS:
		return &MsgSysSetStagePass{}
	case network.MSG_SYS_WAIT_STAGE_BINARY:
		return &MsgSysWaitStageBinary{}
	case network.MSG_SYS_SET_STAGE_BINARY:
		return &MsgSysSetStageBinary{}
	case network.MSG_SYS_GET_STAGE_BINARY:
		return &MsgSysGetStageBinary{}
	case network.MSG_SYS_ENUMERATE_CLIENT:
		return &MsgSysEnumerateClient{}
	case network.MSG_SYS_ENUMERATE_STAGE:
		return &MsgSysEnumerateStage{}
	case network.MSG_SYS_CREATE_MUTEX:
		return &MsgSysCreateMutex{}
	case network.MSG_SYS_CREATE_OPEN_MUTEX:
		return &MsgSysCreateOpenMutex{}
	case network.MSG_SYS_DELETE_MUTEX:
		return &MsgSysDeleteMutex{}
	case network.MSG_SYS_OPEN_MUTEX:
		return &MsgSysOpenMutex{}
	case network.MSG_SYS_CLOSE_MUTEX:
		return &MsgSysCloseMutex{}
	case network.MSG_SYS_CREATE_SEMAPHORE:
		return &MsgSysCreateSemaphore{}
	case network.MSG_SYS_CREATE_ACQUIRE_SEMAPHORE:
		return &MsgSysCreateAcquireSemaphore{}
	case network.MSG_SYS_DELETE_SEMAPHORE:
		return &MsgSysDeleteSemaphore{}
	case network.MSG_SYS_ACQUIRE_SEMAPHORE:
		return &MsgSysAcquireSemaphore{}
	case network.MSG_SYS_RELEASE_SEMAPHORE:
		return &MsgSysReleaseSemaphore{}
	case network.MSG_SYS_LOCK_GLOBAL_SEMA:
		return &MsgSysLockGlobalSema{}
	case network.MSG_SYS_UNLOCK_GLOBAL_SEMA:
		return &MsgSysUnlockGlobalSema{}
	case network.MSG_SYS_CHECK_SEMAPHORE:
		return &MsgSysCheckSemaphore{}
	case network.MSG_SYS_OPERATE_REGISTER:
		return &MsgSysOperateRegister{}
	case network.MSG_SYS_LOAD_REGISTER:
		return &MsgSysLoadRegister{}
	case network.MSG_SYS_NOTIFY_REGISTER:
		return &MsgSysNotifyRegister{}
	case network.MSG_SYS_CREATE_OBJECT:
		return &MsgSysCreateObject{}
	case network.MSG_SYS_DELETE_OBJECT:
		return &MsgSysDeleteObject{}
	case network.MSG_SYS_POSITION_OBJECT:
		return &MsgSysPositionObject{}
	case network.MSG_SYS_ROTATE_OBJECT:
		return &MsgSysRotateObject{}
	case network.MSG_SYS_DUPLICATE_OBJECT:
		return &MsgSysDuplicateObject{}
	case network.MSG_SYS_SET_OBJECT_BINARY:
		return &MsgSysSetObjectBinary{}
	case network.MSG_SYS_GET_OBJECT_BINARY:
		return &MsgSysGetObjectBinary{}
	case network.MSG_SYS_GET_OBJECT_OWNER:
		return &MsgSysGetObjectOwner{}
	case network.MSG_SYS_UPDATE_OBJECT_BINARY:
		return &MsgSysUpdateObjectBinary{}
	case network.MSG_SYS_CLEANUP_OBJECT:
		return &MsgSysCleanupObject{}
	case network.MSG_SYS_reserve4A:
		return &MsgSysReserve4A{}
	case network.MSG_SYS_reserve4B:
		return &MsgSysReserve4B{}
	case network.MSG_SYS_reserve4C:
		return &MsgSysReserve4C{}
	case network.MSG_SYS_reserve4D:
		return &MsgSysReserve4D{}
	case network.MSG_SYS_reserve4E:
		return &MsgSysReserve4E{}
	case network.MSG_SYS_reserve4F:
		return &MsgSysReserve4F{}
	case network.MSG_SYS_INSERT_USER:
		return &MsgSysInsertUser{}
	case network.MSG_SYS_DELETE_USER:
		return &MsgSysDeleteUser{}
	case network.MSG_SYS_SET_USER_BINARY:
		return &MsgSysSetUserBinary{}
	case network.MSG_SYS_GET_USER_BINARY:
		return &MsgSysGetUserBinary{}
	case network.MSG_SYS_NOTIFY_USER_BINARY:
		return &MsgSysNotifyUserBinary{}
	case network.MSG_SYS_reserve55:
		return &MsgSysReserve55{}
	case network.MSG_SYS_reserve56:
		return &MsgSysReserve56{}
	case network.MSG_SYS_reserve57:
		return &MsgSysReserve57{}
	case network.MSG_SYS_UPDATE_RIGHT:
		return &MsgSysUpdateRight{}
	case network.MSG_SYS_AUTH_QUERY:
		return &MsgSysAuthQuery{}
	case network.MSG_SYS_AUTH_DATA:
		return &MsgSysAuthData{}
	case network.MSG_SYS_AUTH_TERMINAL:
		return &MsgSysAuthTerminal{}
	case network.MSG_SYS_reserve5C:
		return &MsgSysReserve5C{}
	case network.MSG_SYS_RIGHTS_RELOAD:
		return &MsgSysRightsReload{}
	case network.MSG_SYS_reserve5E:
		return &MsgSysReserve5E{}
	case network.MSG_SYS_reserve5F:
		return &MsgSysReserve5F{}
	case network.MSG_MHF_SAVEDATA:
		return &MsgMhfSavedata{}
	case network.MSG_MHF_LOADDATA:
		return &MsgMhfLoaddata{}
	case network.MSG_MHF_LIST_MEMBER:
		return &MsgMhfListMember{}
	case network.MSG_MHF_OPR_MEMBER:
		return &MsgMhfOprMember{}
	case network.MSG_MHF_ENUMERATE_DIST_ITEM:
		return &MsgMhfEnumerateDistItem{}
	case network.MSG_MHF_APPLY_DIST_ITEM:
		return &MsgMhfApplyDistItem{}
	case network.MSG_MHF_ACQUIRE_DIST_ITEM:
		return &MsgMhfAcquireDistItem{}
	case network.MSG_MHF_GET_DIST_DESCRIPTION:
		return &MsgMhfGetDistDescription{}
	case network.MSG_MHF_SEND_MAIL:
		return &MsgMhfSendMail{}
	case network.MSG_MHF_READ_MAIL:
		return &MsgMhfReadMail{}
	case network.MSG_MHF_LIST_MAIL:
		return &MsgMhfListMail{}
	case network.MSG_MHF_OPRT_MAIL:
		return &MsgMhfOprtMail{}
	case network.MSG_MHF_LOAD_FAVORITE_QUEST:
		return &MsgMhfLoadFavoriteQuest{}
	case network.MSG_MHF_SAVE_FAVORITE_QUEST:
		return &MsgMhfSaveFavoriteQuest{}
	case network.MSG_MHF_REGISTER_EVENT:
		return &MsgMhfRegisterEvent{}
	case network.MSG_MHF_RELEASE_EVENT:
		return &MsgMhfReleaseEvent{}
	case network.MSG_MHF_TRANSIT_MESSAGE:
		return &MsgMhfTransitMessage{}
	case network.MSG_SYS_reserve71:
		return &MsgSysReserve71{}
	case network.MSG_SYS_reserve72:
		return &MsgSysReserve72{}
	case network.MSG_SYS_reserve73:
		return &MsgSysReserve73{}
	case network.MSG_SYS_reserve74:
		return &MsgSysReserve74{}
	case network.MSG_SYS_reserve75:
		return &MsgSysReserve75{}
	case network.MSG_SYS_reserve76:
		return &MsgSysReserve76{}
	case network.MSG_SYS_reserve77:
		return &MsgSysReserve77{}
	case network.MSG_SYS_reserve78:
		return &MsgSysReserve78{}
	case network.MSG_SYS_reserve79:
		return &MsgSysReserve79{}
	case network.MSG_SYS_reserve7A:
		return &MsgSysReserve7A{}
	case network.MSG_SYS_reserve7B:
		return &MsgSysReserve7B{}
	case network.MSG_SYS_reserve7C:
		return &MsgSysReserve7C{}
	case network.MSG_CA_EXCHANGE_ITEM:
		return &MsgCaExchangeItem{}
	case network.MSG_SYS_reserve7E:
		return &MsgSysReserve7E{}
	case network.MSG_MHF_PRESENT_BOX:
		return &MsgMhfPresentBox{}
	case network.MSG_MHF_SERVER_COMMAND:
		return &MsgMhfServerCommand{}
	case network.MSG_MHF_SHUT_CLIENT:
		return &MsgMhfShutClient{}
	case network.MSG_MHF_ANNOUNCE:
		return &MsgMhfAnnounce{}
	case network.MSG_MHF_SET_LOGINWINDOW:
		return &MsgMhfSetLoginwindow{}
	case network.MSG_SYS_TRANS_BINARY:
		return &MsgSysTransBinary{}
	case network.MSG_SYS_COLLECT_BINARY:
		return &MsgSysCollectBinary{}
	case network.MSG_SYS_GET_STATE:
		return &MsgSysGetState{}
	case network.MSG_SYS_SERIALIZE:
		return &MsgSysSerialize{}
	case network.MSG_SYS_ENUMLOBBY:
		return &MsgSysEnumlobby{}
	case network.MSG_SYS_ENUMUSER:
		return &MsgSysEnumuser{}
	case network.MSG_SYS_INFOKYSERVER:
		return &MsgSysInfokyserver{}
	case network.MSG_MHF_GET_CA_UNIQUE_ID:
		return &MsgMhfGetCaUniqueID{}
	case network.MSG_MHF_SET_CA_ACHIEVEMENT:
		return &MsgMhfSetCaAchievement{}
	case network.MSG_MHF_CARAVAN_MY_SCORE:
		return &MsgMhfCaravanMyScore{}
	case network.MSG_MHF_CARAVAN_RANKING:
		return &MsgMhfCaravanRanking{}
	case network.MSG_MHF_CARAVAN_MY_RANK:
		return &MsgMhfCaravanMyRank{}
	case network.MSG_MHF_CREATE_GUILD:
		return &MsgMhfCreateGuild{}
	case network.MSG_MHF_OPERATE_GUILD:
		return &MsgMhfOperateGuild{}
	case network.MSG_MHF_OPERATE_GUILD_MEMBER:
		return &MsgMhfOperateGuildMember{}
	case network.MSG_MHF_INFO_GUILD:
		return &MsgMhfInfoGuild{}
	case network.MSG_MHF_ENUMERATE_GUILD:
		return &MsgMhfEnumerateGuild{}
	case network.MSG_MHF_UPDATE_GUILD:
		return &MsgMhfUpdateGuild{}
	case network.MSG_MHF_ARRANGE_GUILD_MEMBER:
		return &MsgMhfArrangeGuildMember{}
	case network.MSG_MHF_ENUMERATE_GUILD_MEMBER:
		return &MsgMhfEnumerateGuildMember{}
	case network.MSG_MHF_ENUMERATE_CAMPAIGN:
		return &MsgMhfEnumerateCampaign{}
	case network.MSG_MHF_STATE_CAMPAIGN:
		return &MsgMhfStateCampaign{}
	case network.MSG_MHF_APPLY_CAMPAIGN:
		return &MsgMhfApplyCampaign{}
	case network.MSG_MHF_ENUMERATE_ITEM:
		return &MsgMhfEnumerateItem{}
	case network.MSG_MHF_ACQUIRE_ITEM:
		return &MsgMhfAcquireItem{}
	case network.MSG_MHF_TRANSFER_ITEM:
		return &MsgMhfTransferItem{}
	case network.MSG_MHF_MERCENARY_HUNTDATA:
		return &MsgMhfMercenaryHuntdata{}
	case network.MSG_MHF_ENTRY_ROOKIE_GUILD:
		return &MsgMhfEntryRookieGuild{}
	case network.MSG_MHF_ENUMERATE_QUEST:
		return &MsgMhfEnumerateQuest{}
	case network.MSG_MHF_ENUMERATE_EVENT:
		return &MsgMhfEnumerateEvent{}
	case network.MSG_MHF_ENUMERATE_PRICE:
		return &MsgMhfEnumeratePrice{}
	case network.MSG_MHF_ENUMERATE_RANKING:
		return &MsgMhfEnumerateRanking{}
	case network.MSG_MHF_ENUMERATE_ORDER:
		return &MsgMhfEnumerateOrder{}
	case network.MSG_MHF_ENUMERATE_SHOP:
		return &MsgMhfEnumerateShop{}
	case network.MSG_MHF_GET_EXTRA_INFO:
		return &MsgMhfGetExtraInfo{}
	case network.MSG_MHF_UPDATE_INTERIOR:
		return &MsgMhfUpdateInterior{}
	case network.MSG_MHF_ENUMERATE_HOUSE:
		return &MsgMhfEnumerateHouse{}
	case network.MSG_MHF_UPDATE_HOUSE:
		return &MsgMhfUpdateHouse{}
	case network.MSG_MHF_LOAD_HOUSE:
		return &MsgMhfLoadHouse{}
	case network.MSG_MHF_OPERATE_WAREHOUSE:
		return &MsgMhfOperateWarehouse{}
	case network.MSG_MHF_ENUMERATE_WAREHOUSE:
		return &MsgMhfEnumerateWarehouse{}
	case network.MSG_MHF_UPDATE_WAREHOUSE:
		return &MsgMhfUpdateWarehouse{}
	case network.MSG_MHF_ACQUIRE_TITLE:
		return &MsgMhfAcquireTitle{}
	case network.MSG_MHF_ENUMERATE_TITLE:
		return &MsgMhfEnumerateTitle{}
	case network.MSG_MHF_ENUMERATE_GUILD_ITEM:
		return &MsgMhfEnumerateGuildItem{}
	case network.MSG_MHF_UPDATE_GUILD_ITEM:
		return &MsgMhfUpdateGuildItem{}
	case network.MSG_MHF_ENUMERATE_UNION_ITEM:
		return &MsgMhfEnumerateUnionItem{}
	case network.MSG_MHF_UPDATE_UNION_ITEM:
		return &MsgMhfUpdateUnionItem{}
	case network.MSG_MHF_CREATE_JOINT:
		return &MsgMhfCreateJoint{}
	case network.MSG_MHF_OPERATE_JOINT:
		return &MsgMhfOperateJoint{}
	case network.MSG_MHF_INFO_JOINT:
		return &MsgMhfInfoJoint{}
	case network.MSG_MHF_UPDATE_GUILD_ICON:
		return &MsgMhfUpdateGuildIcon{}
	case network.MSG_MHF_INFO_FESTA:
		return &MsgMhfInfoFesta{}
	case network.MSG_MHF_ENTRY_FESTA:
		return &MsgMhfEntryFesta{}
	case network.MSG_MHF_CHARGE_FESTA:
		return &MsgMhfChargeFesta{}
	case network.MSG_MHF_ACQUIRE_FESTA:
		return &MsgMhfAcquireFesta{}
	case network.MSG_MHF_STATE_FESTA_U:
		return &MsgMhfStateFestaU{}
	case network.MSG_MHF_STATE_FESTA_G:
		return &MsgMhfStateFestaG{}
	case network.MSG_MHF_ENUMERATE_FESTA_MEMBER:
		return &MsgMhfEnumerateFestaMember{}
	case network.MSG_MHF_VOTE_FESTA:
		return &MsgMhfVoteFesta{}
	case network.MSG_MHF_ACQUIRE_CAFE_ITEM:
		return &MsgMhfAcquireCafeItem{}
	case network.MSG_MHF_UPDATE_CAFEPOINT:
		return &MsgMhfUpdateCafepoint{}
	case network.MSG_MHF_CHECK_DAILY_CAFEPOINT:
		return &MsgMhfCheckDailyCafepoint{}
	case network.MSG_MHF_GET_COG_INFO:
		return &MsgMhfGetCogInfo{}
	case network.MSG_MHF_CHECK_MONTHLY_ITEM:
		return &MsgMhfCheckMonthlyItem{}
	case network.MSG_MHF_ACQUIRE_MONTHLY_ITEM:
		return &MsgMhfAcquireMonthlyItem{}
	case network.MSG_MHF_CHECK_WEEKLY_STAMP:
		return &MsgMhfCheckWeeklyStamp{}
	case network.MSG_MHF_EXCHANGE_WEEKLY_STAMP:
		return &MsgMhfExchangeWeeklyStamp{}
	case network.MSG_MHF_CREATE_MERCENARY:
		return &MsgMhfCreateMercenary{}
	case network.MSG_MHF_SAVE_MERCENARY:
		return &MsgMhfSaveMercenary{}
	case network.MSG_MHF_READ_MERCENARY_W:
		return &MsgMhfReadMercenaryW{}
	case network.MSG_MHF_READ_MERCENARY_M:
		return &MsgMhfReadMercenaryM{}
	case network.MSG_MHF_CONTRACT_MERCENARY:
		return &MsgMhfContractMercenary{}
	case network.MSG_MHF_ENUMERATE_MERCENARY_LOG:
		return &MsgMhfEnumerateMercenaryLog{}
	case network.MSG_MHF_ENUMERATE_GUACOT:
		return &MsgMhfEnumerateGuacot{}
	case network.MSG_MHF_UPDATE_GUACOT:
		return &MsgMhfUpdateGuacot{}
	case network.MSG_MHF_INFO_TOURNAMENT:
		return &MsgMhfInfoTournament{}
	case network.MSG_MHF_ENTRY_TOURNAMENT:
		return &MsgMhfEntryTournament{}
	case network.MSG_MHF_ENTER_TOURNAMENT_QUEST:
		return &MsgMhfEnterTournamentQuest{}
	case network.MSG_MHF_ACQUIRE_TOURNAMENT:
		return &MsgMhfAcquireTournament{}
	case network.MSG_MHF_GET_ACHIEVEMENT:
		return &MsgMhfGetAchievement{}
	case network.MSG_MHF_RESET_ACHIEVEMENT:
		return &MsgMhfResetAchievement{}
	case network.MSG_MHF_ADD_ACHIEVEMENT:
		return &MsgMhfAddAchievement{}
	case network.MSG_MHF_PAYMENT_ACHIEVEMENT:
		return &MsgMhfPaymentAchievement{}
	case network.MSG_MHF_DISPLAYED_ACHIEVEMENT:
		return &MsgMhfDisplayedAchievement{}
	case network.MSG_MHF_INFO_SCENARIO_COUNTER:
		return &MsgMhfInfoScenarioCounter{}
	case network.MSG_MHF_SAVE_SCENARIO_DATA:
		return &MsgMhfSaveScenarioData{}
	case network.MSG_MHF_LOAD_SCENARIO_DATA:
		return &MsgMhfLoadScenarioData{}
	case network.MSG_MHF_GET_BBS_SNS_STATUS:
		return &MsgMhfGetBbsSnsStatus{}
	case network.MSG_MHF_APPLY_BBS_ARTICLE:
		return &MsgMhfApplyBbsArticle{}
	case network.MSG_MHF_GET_ETC_POINTS:
		return &MsgMhfGetEtcPoints{}
	case network.MSG_MHF_UPDATE_ETC_POINT:
		return &MsgMhfUpdateEtcPoint{}
	case network.MSG_MHF_GET_MYHOUSE_INFO:
		return &MsgMhfGetMyhouseInfo{}
	case network.MSG_MHF_UPDATE_MYHOUSE_INFO:
		return &MsgMhfUpdateMyhouseInfo{}
	case network.MSG_MHF_GET_WEEKLY_SCHEDULE:
		return &MsgMhfGetWeeklySchedule{}
	case network.MSG_MHF_ENUMERATE_INV_GUILD:
		return &MsgMhfEnumerateInvGuild{}
	case network.MSG_MHF_OPERATION_INV_GUILD:
		return &MsgMhfOperationInvGuild{}
	case network.MSG_MHF_STAMPCARD_STAMP:
		return &MsgMhfStampcardStamp{}
	case network.MSG_MHF_STAMPCARD_PRIZE:
		return &MsgMhfStampcardPrize{}
	case network.MSG_MHF_UNRESERVE_SRG:
		return &MsgMhfUnreserveSrg{}
	case network.MSG_MHF_LOAD_PLATE_DATA:
		return &MsgMhfLoadPlateData{}
	case network.MSG_MHF_SAVE_PLATE_DATA:
		return &MsgMhfSavePlateData{}
	case network.MSG_MHF_LOAD_PLATE_BOX:
		return &MsgMhfLoadPlateBox{}
	case network.MSG_MHF_SAVE_PLATE_BOX:
		return &MsgMhfSavePlateBox{}
	case network.MSG_MHF_READ_GUILDCARD:
		return &MsgMhfReadGuildcard{}
	case network.MSG_MHF_UPDATE_GUILDCARD:
		return &MsgMhfUpdateGuildcard{}
	case network.MSG_MHF_READ_BEAT_LEVEL:
		return &MsgMhfReadBeatLevel{}
	case network.MSG_MHF_UPDATE_BEAT_LEVEL:
		return &MsgMhfUpdateBeatLevel{}
	case network.MSG_MHF_READ_BEAT_LEVEL_ALL_RANKING:
		return &MsgMhfReadBeatLevelAllRanking{}
	case network.MSG_MHF_READ_BEAT_LEVEL_MY_RANKING:
		return &MsgMhfReadBeatLevelMyRanking{}
	case network.MSG_MHF_READ_LAST_WEEK_BEAT_RANKING:
		return &MsgMhfReadLastWeekBeatRanking{}
	case network.MSG_MHF_ACCEPT_READ_REWARD:
		return &MsgMhfAcceptReadReward{}
	case network.MSG_MHF_GET_ADDITIONAL_BEAT_REWARD:
		return &MsgMhfGetAdditionalBeatReward{}
	case network.MSG_MHF_GET_FIXED_SEIBATU_RANKING_TABLE:
		return &MsgMhfGetFixedSeibatuRankingTable{}
	case network.MSG_MHF_GET_BBS_USER_STATUS:
		return &MsgMhfGetBbsUserStatus{}
	case network.MSG_MHF_KICK_EXPORT_FORCE:
		return &MsgMhfKickExportForce{}
	case network.MSG_MHF_GET_BREAK_SEIBATU_LEVEL_REWARD:
		return &MsgMhfGetBreakSeibatuLevelReward{}
	case network.MSG_MHF_GET_WEEKLY_SEIBATU_RANKING_REWARD:
		return &MsgMhfGetWeeklySeibatuRankingReward{}
	case network.MSG_MHF_GET_EARTH_STATUS:
		return &MsgMhfGetEarthStatus{}
	case network.MSG_MHF_LOAD_PARTNER:
		return &MsgMhfLoadPartner{}
	case network.MSG_MHF_SAVE_PARTNER:
		return &MsgMhfSavePartner{}
	case network.MSG_MHF_GET_GUILD_MISSION_LIST:
		return &MsgMhfGetGuildMissionList{}
	case network.MSG_MHF_GET_GUILD_MISSION_RECORD:
		return &MsgMhfGetGuildMissionRecord{}
	case network.MSG_MHF_ADD_GUILD_MISSION_COUNT:
		return &MsgMhfAddGuildMissionCount{}
	case network.MSG_MHF_SET_GUILD_MISSION_TARGET:
		return &MsgMhfSetGuildMissionTarget{}
	case network.MSG_MHF_CANCEL_GUILD_MISSION_TARGET:
		return &MsgMhfCancelGuildMissionTarget{}
	case network.MSG_MHF_LOAD_OTOMO_AIROU:
		return &MsgMhfLoadOtomoAirou{}
	case network.MSG_MHF_SAVE_OTOMO_AIROU:
		return &MsgMhfSaveOtomoAirou{}
	case network.MSG_MHF_ENUMERATE_GUILD_TRESURE:
		return &MsgMhfEnumerateGuildTresure{}
	case network.MSG_MHF_ENUMERATE_AIROULIST:
		return &MsgMhfEnumerateAiroulist{}
	case network.MSG_MHF_REGIST_GUILD_TRESURE:
		return &MsgMhfRegistGuildTresure{}
	case network.MSG_MHF_ACQUIRE_GUILD_TRESURE:
		return &MsgMhfAcquireGuildTresure{}
	case network.MSG_MHF_OPERATE_GUILD_TRESURE_REPORT:
		return &MsgMhfOperateGuildTresureReport{}
	case network.MSG_MHF_GET_GUILD_TRESURE_SOUVENIR:
		return &MsgMhfGetGuildTresureSouvenir{}
	case network.MSG_MHF_ACQUIRE_GUILD_TRESURE_SOUVENIR:
		return &MsgMhfAcquireGuildTresureSouvenir{}
	case network.MSG_MHF_ENUMERATE_FESTA_INTERMEDIATE_PRIZE:
		return &MsgMhfEnumerateFestaIntermediatePrize{}
	case network.MSG_MHF_ACQUIRE_FESTA_INTERMEDIATE_PRIZE:
		return &MsgMhfAcquireFestaIntermediatePrize{}
	case network.MSG_MHF_LOAD_DECO_MYSET:
		return &MsgMhfLoadDecoMyset{}
	case network.MSG_MHF_SAVE_DECO_MYSET:
		return &MsgMhfSaveDecoMyset{}
	case network.MSG_MHF_reserve010F:
		return &MsgMhfReserve010F{}
	case network.MSG_MHF_LOAD_GUILD_COOKING:
		return &MsgMhfLoadGuildCooking{}
	case network.MSG_MHF_REGIST_GUILD_COOKING:
		return &MsgMhfRegistGuildCooking{}
	case network.MSG_MHF_LOAD_GUILD_ADVENTURE:
		return &MsgMhfLoadGuildAdventure{}
	case network.MSG_MHF_REGIST_GUILD_ADVENTURE:
		return &MsgMhfRegistGuildAdventure{}
	case network.MSG_MHF_ACQUIRE_GUILD_ADVENTURE:
		return &MsgMhfAcquireGuildAdventure{}
	case network.MSG_MHF_CHARGE_GUILD_ADVENTURE:
		return &MsgMhfChargeGuildAdventure{}
	case network.MSG_MHF_LOAD_LEGEND_DISPATCH:
		return &MsgMhfLoadLegendDispatch{}
	case network.MSG_MHF_LOAD_HUNTER_NAVI:
		return &MsgMhfLoadHunterNavi{}
	case network.MSG_MHF_SAVE_HUNTER_NAVI:
		return &MsgMhfSaveHunterNavi{}
	case network.MSG_MHF_REGIST_SPABI_TIME:
		return &MsgMhfRegistSpabiTime{}
	case network.MSG_MHF_GET_GUILD_WEEKLY_BONUS_MASTER:
		return &MsgMhfGetGuildWeeklyBonusMaster{}
	case network.MSG_MHF_GET_GUILD_WEEKLY_BONUS_ACTIVE_COUNT:
		return &MsgMhfGetGuildWeeklyBonusActiveCount{}
	case network.MSG_MHF_ADD_GUILD_WEEKLY_BONUS_EXCEPTIONAL_USER:
		return &MsgMhfAddGuildWeeklyBonusExceptionalUser{}
	case network.MSG_MHF_GET_TOWER_INFO:
		return &MsgMhfGetTowerInfo{}
	case network.MSG_MHF_POST_TOWER_INFO:
		return &MsgMhfPostTowerInfo{}
	case network.MSG_MHF_GET_GEM_INFO:
		return &MsgMhfGetGemInfo{}
	case network.MSG_MHF_POST_GEM_INFO:
		return &MsgMhfPostGemInfo{}
	case network.MSG_MHF_GET_EARTH_VALUE:
		return &MsgMhfGetEarthValue{}
	case network.MSG_MHF_DEBUG_POST_VALUE:
		return &MsgMhfDebugPostValue{}
	case network.MSG_MHF_GET_PAPER_DATA:
		return &MsgMhfGetPaperData{}
	case network.MSG_MHF_GET_NOTICE:
		return &MsgMhfGetNotice{}
	case network.MSG_MHF_POST_NOTICE:
		return &MsgMhfPostNotice{}
	case network.MSG_MHF_GET_BOOST_TIME:
		return &MsgMhfGetBoostTime{}
	case network.MSG_MHF_POST_BOOST_TIME:
		return &MsgMhfPostBoostTime{}
	case network.MSG_MHF_GET_BOOST_TIME_LIMIT:
		return &MsgMhfGetBoostTimeLimit{}
	case network.MSG_MHF_POST_BOOST_TIME_LIMIT:
		return &MsgMhfPostBoostTimeLimit{}
	case network.MSG_MHF_ENUMERATE_FESTA_PERSONAL_PRIZE:
		return &MsgMhfEnumerateFestaPersonalPrize{}
	case network.MSG_MHF_ACQUIRE_FESTA_PERSONAL_PRIZE:
		return &MsgMhfAcquireFestaPersonalPrize{}
	case network.MSG_MHF_GET_RAND_FROM_TABLE:
		return &MsgMhfGetRandFromTable{}
	case network.MSG_MHF_GET_CAFE_DURATION:
		return &MsgMhfGetCafeDuration{}
	case network.MSG_MHF_GET_CAFE_DURATION_BONUS_INFO:
		return &MsgMhfGetCafeDurationBonusInfo{}
	case network.MSG_MHF_RECEIVE_CAFE_DURATION_BONUS:
		return &MsgMhfReceiveCafeDurationBonus{}
	case network.MSG_MHF_POST_CAFE_DURATION_BONUS_RECEIVED:
		return &MsgMhfPostCafeDurationBonusReceived{}
	case network.MSG_MHF_GET_GACHA_POINT:
		return &MsgMhfGetGachaPoint{}
	case network.MSG_MHF_USE_GACHA_POINT:
		return &MsgMhfUseGachaPoint{}
	case network.MSG_MHF_EXCHANGE_FPOINT_2_ITEM:
		return &MsgMhfExchangeFpoint2Item{}
	case network.MSG_MHF_EXCHANGE_ITEM_2_FPOINT:
		return &MsgMhfExchangeItem2Fpoint{}
	case network.MSG_MHF_GET_FPOINT_EXCHANGE_LIST:
		return &MsgMhfGetFpointExchangeList{}
	case network.MSG_MHF_PLAY_STEPUP_GACHA:
		return &MsgMhfPlayStepupGacha{}
	case network.MSG_MHF_RECEIVE_GACHA_ITEM:
		return &MsgMhfReceiveGachaItem{}
	case network.MSG_MHF_GET_STEPUP_STATUS:
		return &MsgMhfGetStepupStatus{}
	case network.MSG_MHF_PLAY_FREE_GACHA:
		return &MsgMhfPlayFreeGacha{}
	case network.MSG_MHF_GET_TINY_BIN:
		return &MsgMhfGetTinyBin{}
	case network.MSG_MHF_POST_TINY_BIN:
		return &MsgMhfPostTinyBin{}
	case network.MSG_MHF_GET_SENYU_DAILY_COUNT:
		return &MsgMhfGetSenyuDailyCount{}
	case network.MSG_MHF_GET_GUILD_TARGET_MEMBER_NUM:
		return &MsgMhfGetGuildTargetMemberNum{}
	case network.MSG_MHF_GET_BOOST_RIGHT:
		return &MsgMhfGetBoostRight{}
	case network.MSG_MHF_START_BOOST_TIME:
		return &MsgMhfStartBoostTime{}
	case network.MSG_MHF_POST_BOOST_TIME_QUEST_RETURN:
		return &MsgMhfPostBoostTimeQuestReturn{}
	case network.MSG_MHF_GET_BOX_GACHA_INFO:
		return &MsgMhfGetBoxGachaInfo{}
	case network.MSG_MHF_PLAY_BOX_GACHA:
		return &MsgMhfPlayBoxGacha{}
	case network.MSG_MHF_RESET_BOX_GACHA_INFO:
		return &MsgMhfResetBoxGachaInfo{}
	case network.MSG_MHF_GET_SEIBATTLE:
		return &MsgMhfGetSeibattle{}
	case network.MSG_MHF_POST_SEIBATTLE:
		return &MsgMhfPostSeibattle{}
	case network.MSG_MHF_GET_RYOUDAMA:
		return &MsgMhfGetRyoudama{}
	case network.MSG_MHF_POST_RYOUDAMA:
		return &MsgMhfPostRyoudama{}
	case network.MSG_MHF_GET_TENROUIRAI:
		return &MsgMhfGetTenrouirai{}
	case network.MSG_MHF_POST_TENROUIRAI:
		return &MsgMhfPostTenrouirai{}
	case network.MSG_MHF_POST_GUILD_SCOUT:
		return &MsgMhfPostGuildScout{}
	case network.MSG_MHF_CANCEL_GUILD_SCOUT:
		return &MsgMhfCancelGuildScout{}
	case network.MSG_MHF_ANSWER_GUILD_SCOUT:
		return &MsgMhfAnswerGuildScout{}
	case network.MSG_MHF_GET_GUILD_SCOUT_LIST:
		return &MsgMhfGetGuildScoutList{}
	case network.MSG_MHF_GET_GUILD_MANAGE_RIGHT:
		return &MsgMhfGetGuildManageRight{}
	case network.MSG_MHF_SET_GUILD_MANAGE_RIGHT:
		return &MsgMhfSetGuildManageRight{}
	case network.MSG_MHF_PLAY_NORMAL_GACHA:
		return &MsgMhfPlayNormalGacha{}
	case network.MSG_MHF_GET_DAILY_MISSION_MASTER:
		return &MsgMhfGetDailyMissionMaster{}
	case network.MSG_MHF_GET_DAILY_MISSION_PERSONAL:
		return &MsgMhfGetDailyMissionPersonal{}
	case network.MSG_MHF_SET_DAILY_MISSION_PERSONAL:
		return &MsgMhfSetDailyMissionPersonal{}
	case network.MSG_MHF_GET_GACHA_PLAY_HISTORY:
		return &MsgMhfGetGachaPlayHistory{}
	case network.MSG_MHF_GET_REJECT_GUILD_SCOUT:
		return &MsgMhfGetRejectGuildScout{}
	case network.MSG_MHF_SET_REJECT_GUILD_SCOUT:
		return &MsgMhfSetRejectGuildScout{}
	case network.MSG_MHF_GET_CA_ACHIEVEMENT_HIST:
		return &MsgMhfGetCaAchievementHist{}
	case network.MSG_MHF_SET_CA_ACHIEVEMENT_HIST:
		return &MsgMhfSetCaAchievementHist{}
	case network.MSG_MHF_GET_KEEP_LOGIN_BOOST_STATUS:
		return &MsgMhfGetKeepLoginBoostStatus{}
	case network.MSG_MHF_USE_KEEP_LOGIN_BOOST:
		return &MsgMhfUseKeepLoginBoost{}
	case network.MSG_MHF_GET_UD_SCHEDULE:
		return &MsgMhfGetUdSchedule{}
	case network.MSG_MHF_GET_UD_INFO:
		return &MsgMhfGetUdInfo{}
	case network.MSG_MHF_GET_KIJU_INFO:
		return &MsgMhfGetKijuInfo{}
	case network.MSG_MHF_SET_KIJU:
		return &MsgMhfSetKiju{}
	case network.MSG_MHF_ADD_UD_POINT:
		return &MsgMhfAddUdPoint{}
	case network.MSG_MHF_GET_UD_MY_POINT:
		return &MsgMhfGetUdMyPoint{}
	case network.MSG_MHF_GET_UD_TOTAL_POINT_INFO:
		return &MsgMhfGetUdTotalPointInfo{}
	case network.MSG_MHF_GET_UD_BONUS_QUEST_INFO:
		return &MsgMhfGetUdBonusQuestInfo{}
	case network.MSG_MHF_GET_UD_SELECTED_COLOR_INFO:
		return &MsgMhfGetUdSelectedColorInfo{}
	case network.MSG_MHF_GET_UD_MONSTER_POINT:
		return &MsgMhfGetUdMonsterPoint{}
	case network.MSG_MHF_GET_UD_DAILY_PRESENT_LIST:
		return &MsgMhfGetUdDailyPresentList{}
	case network.MSG_MHF_GET_UD_NORMA_PRESENT_LIST:
		return &MsgMhfGetUdNormaPresentList{}
	case network.MSG_MHF_GET_UD_RANKING_REWARD_LIST:
		return &MsgMhfGetUdRankingRewardList{}
	case network.MSG_MHF_ACQUIRE_UD_ITEM:
		return &MsgMhfAcquireUdItem{}
	case network.MSG_MHF_GET_REWARD_SONG:
		return &MsgMhfGetRewardSong{}
	case network.MSG_MHF_USE_REWARD_SONG:
		return &MsgMhfUseRewardSong{}
	case network.MSG_MHF_ADD_REWARD_SONG_COUNT:
		return &MsgMhfAddRewardSongCount{}
	case network.MSG_MHF_GET_UD_RANKING:
		return &MsgMhfGetUdRanking{}
	case network.MSG_MHF_GET_UD_MY_RANKING:
		return &MsgMhfGetUdMyRanking{}
	case network.MSG_MHF_ACQUIRE_MONTHLY_REWARD:
		return &MsgMhfAcquireMonthlyReward{}
	case network.MSG_MHF_GET_UD_GUILD_MAP_INFO:
		return &MsgMhfGetUdGuildMapInfo{}
	case network.MSG_MHF_GENERATE_UD_GUILD_MAP:
		return &MsgMhfGenerateUdGuildMap{}
	case network.MSG_MHF_GET_UD_TACTICS_POINT:
		return &MsgMhfGetUdTacticsPoint{}
	case network.MSG_MHF_ADD_UD_TACTICS_POINT:
		return &MsgMhfAddUdTacticsPoint{}
	case network.MSG_MHF_GET_UD_TACTICS_RANKING:
		return &MsgMhfGetUdTacticsRanking{}
	case network.MSG_MHF_GET_UD_TACTICS_REWARD_LIST:
		return &MsgMhfGetUdTacticsRewardList{}
	case network.MSG_MHF_GET_UD_TACTICS_LOG:
		return &MsgMhfGetUdTacticsLog{}
	case network.MSG_MHF_GET_EQUIP_SKIN_HIST:
		return &MsgMhfGetEquipSkinHist{}
	case network.MSG_MHF_UPDATE_EQUIP_SKIN_HIST:
		return &MsgMhfUpdateEquipSkinHist{}
	case network.MSG_MHF_GET_UD_TACTICS_FOLLOWER:
		return &MsgMhfGetUdTacticsFollower{}
	case network.MSG_MHF_SET_UD_TACTICS_FOLLOWER:
		return &MsgMhfSetUdTacticsFollower{}
	case network.MSG_MHF_GET_UD_SHOP_COIN:
		return &MsgMhfGetUdShopCoin{}
	case network.MSG_MHF_USE_UD_SHOP_COIN:
		return &MsgMhfUseUdShopCoin{}
	case network.MSG_MHF_GET_ENHANCED_MINIDATA:
		return &MsgMhfGetEnhancedMinidata{}
	case network.MSG_MHF_SET_ENHANCED_MINIDATA:
		return &MsgMhfSetEnhancedMinidata{}
	case network.MSG_MHF_SEX_CHANGER:
		return &MsgMhfSexChanger{}
	case network.MSG_MHF_GET_LOBBY_CROWD:
		return &MsgMhfGetLobbyCrowd{}
	case network.MSG_SYS_reserve180:
		return &MsgSysReserve180{}
	case network.MSG_MHF_GUILD_HUNTDATA:
		return &MsgMhfGuildHuntdata{}
	case network.MSG_MHF_ADD_KOURYOU_POINT:
		return &MsgMhfAddKouryouPoint{}
	case network.MSG_MHF_GET_KOURYOU_POINT:
		return &MsgMhfGetKouryouPoint{}
	case network.MSG_MHF_EXCHANGE_KOURYOU_POINT:
		return &MsgMhfExchangeKouryouPoint{}
	case network.MSG_MHF_GET_UD_TACTICS_BONUS_QUEST:
		return &MsgMhfGetUdTacticsBonusQuest{}
	case network.MSG_MHF_GET_UD_TACTICS_FIRST_QUEST_BONUS:
		return &MsgMhfGetUdTacticsFirstQuestBonus{}
	case network.MSG_MHF_GET_UD_TACTICS_REMAINING_POINT:
		return &MsgMhfGetUdTacticsRemainingPoint{}
	case network.MSG_SYS_reserve188:
		return &MsgSysReserve188{}
	case network.MSG_MHF_LOAD_PLATE_MYSET:
		return &MsgMhfLoadPlateMyset{}
	case network.MSG_MHF_SAVE_PLATE_MYSET:
		return &MsgMhfSavePlateMyset{}
	case network.MSG_SYS_reserve18B:
		return &MsgSysReserve18B{}
	case network.MSG_MHF_GET_RESTRICTION_EVENT:
		return &MsgMhfGetRestrictionEvent{}
	case network.MSG_MHF_SET_RESTRICTION_EVENT:
		return &MsgMhfSetRestrictionEvent{}
	case network.MSG_SYS_reserve18E:
		return &MsgSysReserve18E{}
	case network.MSG_SYS_reserve18F:
		return &MsgSysReserve18F{}
	case network.MSG_MHF_GET_TREND_WEAPON:
		return &MsgMhfGetTrendWeapon{}
	case network.MSG_MHF_UPDATE_USE_TREND_WEAPON_LOG:
		return &MsgMhfUpdateUseTrendWeaponLog{}
	case network.MSG_SYS_reserve192:
		return &MsgSysReserve192{}
	case network.MSG_SYS_reserve193:
		return &MsgSysReserve193{}
	case network.MSG_SYS_reserve194:
		return &MsgSysReserve194{}
	case network.MSG_MHF_SAVE_RENGOKU_DATA:
		return &MsgMhfSaveRengokuData{}
	case network.MSG_MHF_LOAD_RENGOKU_DATA:
		return &MsgMhfLoadRengokuData{}
	case network.MSG_MHF_GET_RENGOKU_BINARY:
		return &MsgMhfGetRengokuBinary{}
	case network.MSG_MHF_ENUMERATE_RENGOKU_RANKING:
		return &MsgMhfEnumerateRengokuRanking{}
	case network.MSG_MHF_GET_RENGOKU_RANKING_RANK:
		return &MsgMhfGetRengokuRankingRank{}
	case network.MSG_MHF_ACQUIRE_EXCHANGE_SHOP:
		return &MsgMhfAcquireExchangeShop{}
	case network.MSG_SYS_reserve19B:
		return &MsgSysReserve19B{}
	case network.MSG_MHF_SAVE_MEZFES_DATA:
		return &MsgMhfSaveMezfesData{}
	case network.MSG_MHF_LOAD_MEZFES_DATA:
		return &MsgMhfLoadMezfesData{}
	case network.MSG_SYS_reserve19E:
		return &MsgSysReserve19E{}
	case network.MSG_SYS_reserve19F:
		return &MsgSysReserve19F{}
	case network.MSG_MHF_UPDATE_FORCE_GUILD_RANK:
		return &MsgMhfUpdateForceGuildRank{}
	case network.MSG_MHF_RESET_TITLE:
		return &MsgMhfResetTitle{}
	case network.MSG_SYS_reserve202:
		return &MsgSysReserve202{}
	case network.MSG_SYS_reserve203:
		return &MsgSysReserve203{}
	case network.MSG_SYS_reserve204:
		return &MsgSysReserve204{}
	case network.MSG_SYS_reserve205:
		return &MsgSysReserve205{}
	case network.MSG_SYS_reserve206:
		return &MsgSysReserve206{}
	case network.MSG_SYS_reserve207:
		return &MsgSysReserve207{}
	case network.MSG_SYS_reserve208:
		return &MsgSysReserve208{}
	case network.MSG_SYS_reserve209:
		return &MsgSysReserve209{}
	case network.MSG_SYS_reserve20A:
		return &MsgSysReserve20A{}
	case network.MSG_SYS_reserve20B:
		return &MsgSysReserve20B{}
	case network.MSG_SYS_reserve20C:
		return &MsgSysReserve20C{}
	case network.MSG_SYS_reserve20D:
		return &MsgSysReserve20D{}
	case network.MSG_SYS_reserve20E:
		return &MsgSysReserve20E{}
	case network.MSG_SYS_reserve20F:
		return &MsgSysReserve20F{}
	}
	return nil
}
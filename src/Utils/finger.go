package Utils

func LoadFingers(typ string) string {
	if typ == "tcp" {
		return `[{"name": "Mysql_unauthorized", "level": 0, "default_port": {"3306": true}, "isVuln": true, "protocol": "tcp", "regexps": {"regexp": ["Host .* is not allowed to connect to this MySQL server"]}}, {"name": "MariaDB_unauthorized", "level": 0, "default_port": {"3306": true}, "isVuln": true, "protocol": "tcp", "regexps": {"regexp": ["Host .* is not allowed to connect to this MariaDB server"]}}, {"name": "MySQL", "level": 0, "default_port": {"3306": true}, "protocol": "tcp", "regexps": {"regexp": ["^.\u0000\u0000\u0000\n(.\\.[-_~.+\\w]+)\u0000", "^.\u0000\u0000\u0000\u00ffj\u0004'[\\d.]+' .* MySQL"]}}, {"name": "MariaDB", "level": 0, "default_port": {"3306": true}, "protocol": "tcp", "regexps": {"regexp": ["^.\u0000\u0000\u0000\n(5\\.[-_~.+:\\w]+MariaDB-[-_~.+:\\w]+)\u0000"]}}, {"name": "redis", "level": 0, "vuln": "redis_unauthorized", "default_port": {"6379": true}, "protocol": "tcp", "send_data": "info\n", "regexps": {"regexp": ["redis_version:(.*)"], "info": ["-NOAUTH Authentication"]}}]`
	} else if typ == "http" {
		return `[{"name": "Portainer(Docker\u7ba1\u7406)", "protocol": "http", "level": 0, "regexps": {"html": ["portainer.updatePassword", "portainer.init.admin"]}}, {"name": "Gogs\u7b80\u6613Git\u670d\u52a1", "protocol": "http", "level": 0, "regexps": {"cookie": ["i_like_gogs", "i_like_gogits", "i_like_gitea"]}}, {"name": "\u5b9d\u5854-BT.cn", "protocol": "http", "level": 0, "regexps": {"html": ["app.bt.cn/static/app.png", "\u5b89\u5168\u5165\u53e3\u6821\u9a8c\u5931\u8d25"]}}, {"name": "Nexus", "protocol": "http", "level": 0, "regexps": {"html": ["Nexus Repository Manager"], "cookie": ["NX-ANTI-CSRF-TOKEN"]}}, {"name": "Harbor", "protocol": "http", "level": 0, "regexps": {"html": ["<title>Harbor</title>"], "cookie": ["harbor-lang"]}}, {"name": "\u7985\u9053", "protocol": "http", "level": 0, "regexps": {"html": ["/theme/default/images/main/zt-logo.png", "/zentao/theme/zui/css/min.css"], "cookie": ["zentaosid"]}}, {"name": "\u5fae\u64ce", "protocol": "http", "level": 0, "regexps": {"html": ["/index.php?c=user&a=find-password"]}}, {"name": "\u534f\u4f17OA", "protocol": "http", "level": 0, "regexps": {"html": ["Powered by \u534f\u4f17OA"], "cookie": ["CNOAOASESSID"]}}, {"name": "xxl-job", "protocol": "http", "level": 0, "regexps": {"html": ["\u5206\u5e03\u5f0f\u4efb\u52a1\u8c03\u5ea6\u5e73\u53f0XXL-JOB"]}}, {"name": "atmail-WebMail", "protocol": "http", "level": 0, "regexps": {"cookie": ["atmail6"], "html": ["/index.php/mail/auth/processlogin"]}}, {"name": "weblogic", "protocol": "http", "level": 0, "regexps": {"html": ["/console/framework/skins/wlsconsole/images/login_WebLogic_branding.png", "Welcome to Weblogic Application Server", "<i>Hypertext Transfer Protocol -- HTTP/1.1</i>"]}}, {"name": "WebSphere", "protocol": "http", "level": 0, "regexps": {"html": ["SRVE0255E", "ibm.ws.webcontainer.servlet", "IBM WebSphere Application Server"]}}, {"name": "\u5b89\u7b56\u7ee9\u6548\u7ba1\u7406\u7cfb\u7edf", "protocol": "http", "level": 0, "regexps": {"html": ["\u5b89\u7b56\u7ee9\u6548\u4fe1\u606f\u79d1\u6280"]}}, {"name": "\u7528\u53cb\u81f4\u8fdcOA", "protocol": "http", "level": 0, "regexps": {"html": ["/seeyon/USER-DATA/IMAGES/LOGIN/login.gif", "/seeyon/common/"]}}, {"name": "\u7528\u53cb NC", "protocol": "http", "level": 0, "regexps": {"html": ["/nc/servlet/nc.ui.iufo.login.Index"]}}, {"name": "\u7528\u53cb GRP-u8", "protocol": "http", "level": 0, "regexps": {"html": ["/R9iPortal/js/jquery.js"]}}, {"name": "\u7528\u53cb GRP-R9i", "protocol": "http", "level": 0, "regexps": {"html": ["R9iPortal/index.jsp"]}}, {"name": "\u7528\u53cb-UFIDA", "protocol": "http", "level": 0, "regexps": {"html": ["/System/Login/Login.asp?"]}}, {"name": "\u7528\u53cb U8-Cloud ", "protocol": "http", "level": 0, "regexps": {"html": ["uclient.yonyou.com/api/uclient/public/download"]}}, {"name": "\u7528\u53cb-RMIS", "protocol": "http", "level": 0, "regexps": {"html": ["css/commonlogincomp.min.css"]}}, {"name": "\u7528\u53cb-\u65f6\u7a7a\u4f01\u4e1a\u4fe1\u606f\u7ba1\u7406\u95e8\u6237", "protocol": "http", "level": 0, "regexps": {"html": ["mkbh = getCookie("]}}, {"name": "\u7528\u53cb-\u667a\u80fd\u5de5\u5382|\u667a\u80fd\u7f51\u5173", "protocol": "http", "level": 0, "regexps": {"html": ["/modules/core/client/views/header.client.view.html"]}}, {"name": "\u84dd\u51ccoa EKP", "protocol": "http", "level": 0, "regexps": {"html": ["/resource/customization/images/login_single_random"]}}, {"name": "\u534e\u5929\u52a8\u529boa", "protocol": "http", "level": 0, "regexps": {"html": ["/htoa/css/CN/BLUE/login.css"]}}, {"name": "\u7545\u6377\u901a", "protocol": "http", "level": 0, "regexps": {"html": ["\u7545\u6377\u901a T+", "tp.control.css.ashx"]}}, {"name": "discuz", "protocol": "http", "level": 0, "regexps": {"html": ["content=\"Discuz! X\""]}}, {"name": "Typecho", "protocol": "http", "level": 0, "regexps": {"html": ["Typecho</a>"]}}, {"name": "\u91d1\u8776EAS", "protocol": "http", "level": 0, "regexps": {"html": ["easSessionId"]}}, {"name": "phpMyAdmin", "protocol": "http", "level": 0, "regexps": {"cookie": ["pma_lang", "phpMyAdmin"], "html": ["/themes/pmahomme/img/logo_right.png"]}}, {"name": "H3C-AM8000", "protocol": "http", "level": 0, "regexps": {"html": ["AM8000"]}}, {"name": "360\u4f01\u4e1a\u7248", "protocol": "http", "level": 0, "regexps": {"html": ["360EntWebAdminMD5Secret"]}}, {"name": "H3C\u516c\u53f8\u4ea7\u54c1", "protocol": "http", "level": 0, "regexps": {"html": ["service@h3c.com"]}}, {"name": "H3C ICG 1000", "protocol": "http", "level": 0, "regexps": {"html": ["ICG 1000\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "Citrix-Metaframe", "protocol": "http", "level": 0, "regexps": {"html": ["window.location=\"/Citrix/MetaFrame"]}}, {"name": "H3C ER5100", "protocol": "http", "level": 0, "regexps": {"html": ["ER5100\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "CISCO_EPC3925", "protocol": "http", "level": 0, "regexps": {"html": ["Docsis_system"]}}, {"name": "CISCO ASR", "protocol": "http", "level": 0, "regexps": {"html": ["CISCO ASR"]}}, {"name": "H3C ER3200", "protocol": "http", "level": 0, "regexps": {"html": ["ER3200\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "\u4e07\u6237ezOFFICE", "protocol": "http", "level": 0, "regexps": {"header": ["LocLan"]}}, {"name": "\u4e07\u6237\u7f51\u7edc", "protocol": "http", "level": 0, "regexps": {"html": ["css/css_whir.css"]}}, {"name": "Spark_Master", "protocol": "http", "level": 0, "regexps": {"html": ["Spark Master at"]}}, {"name": "\u534e\u4e3a_HUAWEI_SRG2220", "protocol": "http", "level": 0, "regexps": {"html": ["HUAWEI SRG2220"]}}, {"name": "\u84dd\u51ccEIS\u667a\u6167\u534f\u540c\u5e73\u53f0", "protocol": "http", "level": 0, "regexps": {"html": ["/scripts/jquery.landray.common.js"]}}, {"name": "\u6df1\u4fe1\u670dssl-vpn", "protocol": "http", "level": 0, "regexps": {"html": ["login_psw.csp"]}}, {"name": "\u534e\u4e3a NetOpen", "protocol": "http", "level": 0, "regexps": {"html": ["/netopen/theme/css/inFrame.css"]}}, {"name": "Citrix-Web-PN-Server", "protocol": "http", "level": 0, "regexps": {"html": ["Citrix Web PN Server"]}}, {"name": "juniper_vpn", "protocol": "http", "level": 0, "regexps": {"html": ["welcome.cgi?p=logo", "/images/logo_juniper_reversed.gif"]}}, {"name": "360\u4e3b\u673a\u536b\u58eb", "protocol": "http", "level": 0, "regexps": {"header": ["zhuji.360.cn"]}}, {"name": "Nagios", "protocol": "http", "level": 0, "regexps": {"header": ["Nagios Access"]}}, {"name": "H3C ER8300", "protocol": "http", "level": 0, "regexps": {"html": ["ER8300\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "Citrix-Access-Gateway", "protocol": "http", "level": 0, "regexps": {"html": ["Citrix Access Gateway"]}}, {"name": "\u534e\u4e3a MCU", "protocol": "http", "level": 0, "regexps": {"html": ["McuR5-min.js"]}}, {"name": "TP-LINK Wireless WDR3600", "protocol": "http", "level": 0, "regexps": {"html": ["TP-LINK Wireless WDR3600"]}}, {"name": "\u6cdb\u5faeOA E-cology", "protocol": "http", "level": 0, "regexps": {"header": ["ecology_JSessionid"], "html": ["/wui/theme/ecology", "/cloudstore/resource/pc/com"]}}, {"name": "\u6cdb\u5fae E-Mobile", "protocol": "http", "level": 0, "regexps": {"html": ["Weaver E-Mobile"]}}, {"name": "\u6cdb\u5fae E-Office", "protocol": "http", "level": 0, "regexps": {"html": ["/lang/cn/javascript.lang.js"]}}, {"name": "\u6cdb\u5fae\u4e91\u6865 E-Bridge", "protocol": "http", "level": 0, "regexps": {"html": ["/main/assets/images/favicon.ico"]}}, {"name": "TP-LINK Wireless WDR3600", "protocol": "http", "level": 0, "regexps": {"html": ["TP-LINK Wireless WDR3600"]}}, {"name": "\u534e\u4e3a_HUAWEI_ASG2050", "protocol": "http", "level": 0, "regexps": {"html": ["HUAWEI ASG2050"]}}, {"name": "360\u7f51\u7ad9\u536b\u58eb", "protocol": "http", "level": 0, "regexps": {"html": ["360wzb"]}}, {"name": "Citrix-XenServer", "protocol": "http", "level": 0, "regexps": {"html": ["Citrix Systems, Inc. XenServer"]}}, {"name": "H3C ER2100V2", "protocol": "http", "level": 0, "regexps": {"html": ["ER2100V2\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "zabbix", "protocol": "http", "level": 0, "regexps": {"cookie": ["zbx_sessionid"], "html": ["images/general/zabbix.ico", "Zabbix SIA"]}}, {"name": "CISCO_VPN", "protocol": "http", "level": 0, "regexps": {"header": ["webvpn"]}}, {"name": "360\u7ad9\u957f\u5e73\u53f0", "protocol": "http", "level": 0, "regexps": {"html": ["360-site-verification"]}}, {"name": "H3C ER3108GW", "protocol": "http", "level": 0, "regexps": {"html": ["ER3108GW\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "o2security_vpn", "protocol": "http", "level": 0, "regexps": {"header": ["client_param=install_active"]}}, {"name": "H3C ER3260G2", "protocol": "http", "level": 0, "regexps": {"html": ["ER3260G2\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "H3C ICG1000", "protocol": "http", "level": 0, "regexps": {"html": ["ICG1000\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "CISCO-CX20", "protocol": "http", "level": 0, "regexps": {"html": ["CISCO-CX20"]}}, {"name": "H3C ER5200", "protocol": "http", "level": 0, "regexps": {"html": ["ER5200\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "linksys-vpn-bragap14-parintins", "protocol": "http", "level": 0, "regexps": {"html": ["linksys-vpn-bragap14-parintins"]}}, {"name": "360\u7f51\u7ad9\u536b\u58eb\u5e38\u7528\u524d\u7aef\u516c\u5171\u5e93", "protocol": "http", "level": 0, "regexps": {"html": ["libs.useso.com"]}}, {"name": "H3C ER3100", "protocol": "http", "level": 0, "regexps": {"html": ["ER3100\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "H3C-SecBlade-FireWall", "protocol": "http", "level": 0, "regexps": {"html": ["js/MulPlatAPI.js"]}}, {"name": "360webfacil_360WebManager", "protocol": "http", "level": 0, "regexps": {"html": ["publico/template/"]}}, {"name": "Citrix_Netscaler", "protocol": "http", "level": 0, "regexps": {"html": ["ns_af"]}}, {"name": "H3C ER6300G2", "protocol": "http", "level": 0, "regexps": {"html": ["ER6300G2\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "H3C ER3260", "protocol": "http", "level": 0, "regexps": {"html": ["ER3260\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "\u534e\u4e3a_HUAWEI_SRG3250", "protocol": "http", "level": 0, "regexps": {"html": ["HUAWEI SRG3250"]}}, {"name": "exchange", "protocol": "http", "level": 0, "regexps": {"html": ["/owa/auth.owa"]}}, {"name": "Spark_Worker", "protocol": "http", "level": 0, "regexps": {"html": ["Spark Worker at"]}}, {"name": "H3C ER3108G", "protocol": "http", "level": 0, "regexps": {"html": ["ER3108G\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "\u6df1\u4fe1\u670d\u9632\u706b\u5899\u7c7b\u4ea7\u54c1", "protocol": "http", "level": 0, "regexps": {"html": ["SANGFOR FW"]}}, {"name": "Citrix-ConfProxy", "protocol": "http", "level": 0, "regexps": {"html": ["confproxy"]}}, {"name": "360\u7f51\u7ad9\u5b89\u5168\u68c0\u6d4b", "protocol": "http", "level": 0, "regexps": {"html": ["webscan.360.cn/status/pai/hash"]}}, {"name": "H3C ER5200G2", "protocol": "http", "level": 0, "regexps": {"html": ["ER5200G2\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "\u534e\u4e3a\uff08HUAWEI\uff09\u5b89\u5168\u8bbe\u5907", "protocol": "http", "level": 0, "regexps": {"html": ["sweb-lib/resource/"]}}, {"name": "H3C ER6300", "protocol": "http", "level": 0, "regexps": {"html": ["ER6300\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "\u534e\u4e3a_HUAWEI_ASG2100", "protocol": "http", "level": 0, "regexps": {"html": ["HUAWEI ASG2100"]}}, {"name": "TP-Link 3600 DD-WRT", "protocol": "http", "level": 0, "regexps": {"html": ["TP-Link 3600 DD-WRT"]}}, {"name": "NETGEAR WNDR3600", "protocol": "http", "level": 0, "regexps": {"html": ["NETGEAR WNDR3600"]}}, {"name": "H3C ER2100", "protocol": "http", "level": 0, "regexps": {"html": ["ER2100\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "\u7eff\u76df\u4e0b\u4e00\u4ee3\u9632\u706b\u5899", "protocol": "http", "level": 0, "regexps": {"html": ["NSFOCUS NF"]}}, {"name": "jira", "protocol": "http", "level": 0, "regexps": {"html": ["jira.webresources"]}}, {"name": "\u91d1\u548c\u534f\u540c\u7ba1\u7406\u5e73\u53f0", "protocol": "http", "level": 0, "regexps": {"html": ["\u91d1\u548c\u534f\u540c\u7ba1\u7406\u5e73\u53f0", "JHSoft.UI.Lib"]}}, {"name": "Citrix-NetScaler", "protocol": "http", "level": 0, "regexps": {"html": ["NS-CACHE"]}}, {"name": "linksys-vpn", "protocol": "http", "level": 0, "regexps": {"header": ["linksys-vpn"]}}, {"name": "\u901a\u8fbeOA", "protocol": "http", "level": 0, "regexps": {"html": ["/static/images/tongda.ico"]}}, {"name": "\u534e\u4e3a\uff08HUAWEI\uff09Secoway\u8bbe\u5907", "protocol": "http", "level": 0, "regexps": {"html": ["Secoway"]}}, {"name": "\u534e\u4e3a_HUAWEI_SRG1220", "protocol": "http", "level": 0, "regexps": {"html": ["HUAWEI SRG1220"]}}, {"name": "H3C ER2100n", "protocol": "http", "level": 0, "regexps": {"html": ["ER2100n\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "H3C ER8300G2", "protocol": "http", "level": 0, "regexps": {"html": ["ER8300G2\u7cfb\u7edf\u7ba1\u7406"]}}, {"name": "\u91d1\u8776\u653f\u52a1GSiS", "protocol": "http", "level": 0, "regexps": {"html": ["/kdgs/script/kdgs.js"]}}, {"name": "Jboss", "protocol": "http", "level": 0, "regexps": {"html": ["Welcome to JBoss", "jboss.css"], "header": ["JBoss"]}}, {"name": "\u6cdb\u5faeE-mobile", "protocol": "http", "level": 0, "regexps": {"html": ["Weaver E-mobile", "weaver,e-mobile"]}}, {"name": "\u9f50\u6cbb\u5821\u5792\u673a", "protocol": "http", "level": 0, "regexps": {"html": ["logo-icon-ico72.png"]}}, {"name": "ThinkPHP", "protocol": "http", "level": 0, "regexps": {"header": ["ThinkPHP"], "html": ["/Public/static/js/"]}}, {"name": "weaver-ebridge", "protocol": "http", "level": 0, "regexps": {"html": ["e-Bridge,http://wx.weaver"]}}, {"name": "Laravel", "protocol": "http", "level": 0, "regexps": {"header": ["laravel_session"]}}, {"name": "DWR", "protocol": "http", "level": 0, "regexps": {"html": ["dwr/engine.js"]}}, {"name": "swagger_ui", "protocol": "http", "level": 0, "regexps": {"html": ["swagger-ui/css", "\"swagger\":", "swagger-ui.min.js"]}}, {"name": "\u5927\u6c49\u7248\u901a\u53d1\u5e03\u7cfb\u7edf", "protocol": "http", "level": 0, "regexps": {"html": ["\u5927\u6c49\u7248\u901a\u53d1\u5e03\u7cfb\u7edf", "\u5927\u6c49\u7f51\u7edc"]}}, {"name": "\u5927\u6c49Web", "protocol": "http", "level": 0, "regexps": {"html": ["hanweb.css"]}}, {"name": "\u5947\u5b89\u4fe1VPN/\u7f51\u5eb7", "protocol": "http", "level": 0, "regexps": {"html": ["new_style/deepblue.css"]}}, {"name": "IIS", "protocol": "http", "level": 0, "regexps": {"html": ["Bad Request (Invalid Hostname)"]}}, {"name": "JeeSite", "protocol": "http", "level": 0, "regexps": {"html": ["jeesite.min.css"]}}, {"name": "JeeCMS", "protocol": "http", "level": 0, "regexps": {"html": ["/r/cms/www/default/css/jSlider.css", "/r/cms/jquery.js"]}}, {"name": "\u6df1\u4fe1\u670d\u8bbe\u5907", "protocol": "http", "level": 0, "regexps": {"html": ["at sangforsvpn@sangfor.com.cn"]}}, {"name": "\u6df1\u4fe1\u670d\u4e0a\u7f51\u884c\u4e3a\u7ba1\u7406\u7cfb\u7edf", "protocol": "http", "level": 1, "send_data": "/login/img/product_logo.png", "regexps": {"md5": ["d8de64b0ca4281380295aa23500267ca"]}}]`
	} else if typ == "md5" {
		return `{"04d9541338e525258daf47cc844d59f3": "BIG-IP", "302464c3f6207d57240649926cfc7bd4": "\u84dd\u51ccOA", "799f70b71314a7508326d1d2f68f7519": "JBOSS", "d8d7c9138e93d43579ebf2e384745ba8": "\u9510\u6377\u7f51\u5173", "0b24d4d5c7d300d50ee1cd96059a9e85": "\u6df1\u4fe1\u670dedr", "cdc85452665e7708caed3009ecb7d4e2": "\u81f4\u8fdcOA", "17ac348fcce0b320e7bfab3fe2858dfa": "\u81f4\u8fdcOA-A8+", "57f307ad3764553df84e7b14b7a85432": "\u81f4\u8fdcOA", "3c8df395ec2cbd72782286d18a286a9a": "\u81f4\u8fdcOA", "2f761c27b6b7f9386bbd61403635dc42": "\u81f4\u8fdcOA", "48ee373f098d8e96e53b7dd778f09ff4": "\u9f50\u6cbb\u5821\u5792\u673a", "0488faca4c19046b94d07c3ee83cf9d6": "SprintBoot", "f49c4a4bde1eec6c0b80c2277c76e3db": "ThinkPHP", "ed0044587917c76d08573577c8b72883": "\u901a\u8fbeOA", "41eca7a9245394106a09b2534d8030df": "\u6cdb\u5fae", "c27547e27e1d2c7514545cd8d5988946": "\u6cdb\u5faeOA", "9b1d3f08ede38dbe699d6b2e72a8febb": "\u6cdb\u5faeOA", "281348dd57383c1f214ffb8aed3a1210": "\u6cdb\u5faeOA", "23e8c7bd78e8cd826c5a6073b15068b1": "Jenkins", "4644f2d45601037b8423d45e13194c93": "Tomcat", "4859e39ae6c0f1f428f2126a6bb32bd9": "ASP.NET", "f7e3d97f404e71d302b3239eef48d5f2": "GitLab", "85c754581e1d4b628be5b7712c042224": "GitLab", "e49fd30ea870c7a820464ca56a113e6e": "\u82e5\u4f9d", "4eeb8a8eb30b70af511dcc28c11a3216": "\u82e5\u4f9d", "3c3804ddad747313df58112a3e1bc680": "\u5b9d\u5854", "6eb4a43cb64c97f76562af703893c8fd": "XMAPP", "531b63a51234bb06c9d77f219eb25553": "PHPMyAdmin", "825af13371930eeb2f85cf075fa25b68": "IIS", "219732cc524bb3fa67f8ac25784b640e": "WDcp linux\u7ba1\u7406\u7cfb\u7edf", "c028c4822428e83a358c60a93ef65381": "Discuz!", "fe4792d482196a50cf9ae0d9d90b6493": "Vue", "97c6417ed01bdc0ae3ef32ae4894fd03": "Jupyter", "5a508149f630bfd801df6e04369527f9": "Harbor", "902bfddccca904dd2b5f5da05fe9f226": "egg.js", "3031fb045188ef19ad81f39cb063da19": "EL-admin\u7ba1\u7406\u7cfb\u7edf", "a976d227e5d1dcf62f5f7e623211dd1b": "Yii", "7ef1f0a0093460fe46bb691578c07c95": "DedeCMS", "effadc8bcbed022ebaeb16ae33a20ca5": "CRMEB \u7535\u5546\u9500\u8425\u7cfb\u7edf", "b9aa7c338693424aae99599bec875b5f": "Angular.js", "fa8833ac684d6949f480ae3bc9a51b4a": "\u5fae\u64ce", "0c295d21ff53bdd0a95499db97acf0dd": "Nexux", "966e60f8eb85b7ea43a7b0095f3e2336": "Confluence", "20334371817c7368907b5ea52aab2d9e": "JumpServer", "293a3a7763eee19a068ef3dff98a7b8a": "\u6df1\u4fe1\u670d\u8bbe\u5907"}`
	} else if typ == "port" {
		return `[{"name": "top1", "ports": ["80", "443", "8080"], "type": ["http"]}, {"name": "top2", "ports": ["80-90", "443", "1080", "2000", "2001", "3000", "4443", "4430", "5000", "6000", "7000-7003", "9000-9003", "8080-8090", "8000-8020", "8443", "8787", "7080", "8070", "7070", "9080-9083", "5555", "6666", "7777", "9999", "8888", "8889", "9090", "8091", "8099", "8848", "8060", "8899", "800", "801", "10000", "10080"], "type": ["http"]}, {"name": "top3", "ports": ["9443", "6080", "6443", "9091", "7003-7010", "9003-9010", "8100-8110", "8161", "8021-8030", "8880-8890", "8010-8020", "8090-8100", "8180-8181", "8363", "8800", "8761", "8873", "8866", "8900", "8282", "8999", "8989", "8066", "8200", "8111", "8030", "8040", "8060", "8180", "10800"], "type": ["http"]}, {"name": "iis", "ports": ["80", "443", "47001"], "type": ["http"]}, {"name": "jboss", "ports": ["45566", "8080", "80", "8443"], "type": ["http"]}, {"name": "postgresql", "ports": ["5432"], "type": ["db"]}, {"name": "mssql", "ports": ["1433"], "type": ["db"]}, {"name": "mysql", "ports": ["3300-3310"], "type": ["db"]}, {"name": "oracle", "ports": ["1158", "1521", "210"], "type": ["db"]}, {"name": "counchdb", "ports": ["5984"], "type": ["db"]}, {"name": "redis", "ports": ["6379"], "type": ["db", "rce"]}, {"name": "memcache", "ports": ["11211"], "type": ["db"]}, {"name": "sybase", "ports": ["5000", "4100"], "type": ["db"]}, {"name": "mongodb", "ports": ["27017"], "type": ["db"]}, {"name": "hbase", "ports": ["16000", "16010", "16020", "16030"], "type": ["db"]}, {"name": "jndi", "ports": ["1098-1101", "1000-1001", "4444-4447", "10999", "19001", "9999", "8083", "8686", "10001", "11099"], "type": ["rce"]}, {"name": "jdwp", "ports": ["8000", "8080"], "type": ["rce"]}, {"name": "jmx", "ports": ["8686", "8093", "9010-9012", "50500", "61616"], "type": ["rce"]}, {"name": "php-xdebug", "ports": ["9000"], "type": ["rce"]}, {"name": "nodejs-debug", "ports": ["5858", "9229"], "type": ["rce"]}, {"name": "snmp", "ports": ["1161-1162"], "type": ["info"]}, {"name": "glassfish", "ports": ["4848"], "type": ["rce"]}, {"name": "rocketmq", "ports": ["9876", "10909", "10911", "10912"], "type": ["rce"]}, {"name": "activemq", "ports": ["8161"], "type": ["rce"]}, {"name": "kafka", "ports": ["9092"], "type": ["rce"]}, {"name": "cisco", "ports": ["4786"], "type": ["rce"]}, {"name": "rlogin", "ports": ["512-514"], "type": ["rce"]}, {"name": "hp", "ports": ["5555", "5556"], "type": ["rce"]}, {"name": "docker", "ports": ["2375-2377"], "type": ["rce"]}, {"name": "solr", "ports": ["8983"], "type": ["rce"]}, {"name": "zoho", "ports": ["8383"], "type": ["rce"]}, {"name": "altassian", "ports": ["4990"], "type": ["rce"]}, {"name": "portainer", "ports": ["9000"], "type": ["rce"]}, {"name": "hashicorp", "ports": ["8500"], "type": ["rce"]}, {"name": "flink", "ports": ["8081"], "type": ["rce"]}, {"name": "ajp", "ports": ["8009"], "type": ["rce"]}, {"name": "elasticsearch", "ports": ["9200", "9300"], "type": ["rce"]}, {"name": "zabbix", "ports": ["8069"], "type": ["rce"]}, {"name": "windows", "ports": ["88", "135", "137", "139", "389", "445", "1080", "3389", "5985"], "type": ["win"]}, {"name": "other", "ports": ["21-23", "69", "161", "901-902", "50000"], "type": ["info"]}, {"name": "email", "ports": ["25,", "109", "110", "143", "465", "995", "993"], "type": ["info"]}, {"name": "zoomkeeper", "ports": ["2181", "2888", "3888"], "type": ["info"]}, {"name": "rsync", "ports": ["873"], "type": ["info", "brute"]}, {"name": "lotus", "ports": ["1352"], "type": ["info"]}, {"name": "nfs", "ports": ["2049"], "type": ["rce"]}, {"name": "oracle-ftp", "ports": ["2100"], "type": ["rce"]}, {"name": "squid", "ports": ["Squid"], "type": ["rce"]}, {"name": "pcanywhere", "ports": ["5632"], "type": ["info"]}, {"name": "vnc", "ports": ["5900"], "type": ["brute"]}, {"name": "websphere", "ports": ["9043", "9060", "9080-9083", "9443", "2809", "8880", "9401-9402", "9100", "9353", "7276", "7286", "5558", "5578"], "type": ["info"]}, {"name": "hadoop", "ports": ["8020", "9000", "50070", "50010", "50020", "50075", "50470", "50475", "50090", "9083", "8485", "8480", "8040-8042", "19888", "41414"], "type": ["info"]}]`
	}
	return ""
}

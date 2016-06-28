package ftp

var replies = map[int]string{
	200: "Command okay.",
	202: "Command not implemented, superfluous at this site.",
	211: "System status, or system help reply.",
	212: "Directory status.",
	213: "File status.",
	214: "Help message.",
	215: "NAME system type.",
	220: "Service ready for new user.",
	221: "Service closing control connection.",
	225: "Data connection open; no transfer in progress.",
	226: "Closing data connection.",
	227: "Entering Passive Mode.",
	230: "User logged in, proceed.",
	250: "Requested file action okay, completed.",
	257: "\"PATHNAME\" created.",

	331: "User name okay, need password.",
	332: "Need account for login.",
	350: "Requested file action pending further information.",

	421: "Service not available, closing control connection.",
	425: "Can't open data connection.",
	426: "Connection closed; transfer aborted.",
	450: "File unavailable.",
	451: "Local error in processing.",
	452: "Insufficient storage space in system.",

	500: "Syntax error, command unrecognized.",
	501: "Syntax error in parameters or arguments.",
	502: "Command not implemented.",
	503: "Bad sequence of commands.",
	504: "Command not implemented for that parameter.",
	530: "Not logged in.",
	532: "Need account for storing files.",
	550: "File unavailable.",
	551: "Page type unknown.",
	552: "Exceeded storage allocation.",
	553: "File name not allowed.",
}

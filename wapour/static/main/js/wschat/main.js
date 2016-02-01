define(
	"main",
	[
		"MessageList"
	],
	function(MessageList) {
		var ws = new WebSocket("ws://10.10.111.143:8090/entry");
		var list = new MessageList(ws);
		ko.applyBindings(list);
	}
);

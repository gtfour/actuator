define(
	"main",
	[
		"MessageList"
	],
	function(MessageList) {
		var ws = new WebSocket("ws://127.0.0.1:8090/entry");
		var list = new MessageList(ws);
		ko.applyBindings(list);
	}
);

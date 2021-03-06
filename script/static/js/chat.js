var messageTxt;
var messages;

$(function () {

	 messageTxt = $("#messageTxt");
	 messages = $("#messages");


	w = new Ws("ws://" + HOST + "/sync");
	w.OnConnect(function () {
		console.log("Websocket connection established");
        console.log(HOST)
	});

	w.OnDisconnect(function () {
		appendMessage($("<div><center><h3>Disconnected</h3></center></div>"));
        console.log("Websocket connection destroied")
	});

    var count = 0
    /*
    var message = {
        image: "busybox",
        tag: "v"+String(count)
    }
    */

    // var m = ""

    w.On("sync", function (m) {
        appendMessage($("<div>" + m + "</div>")); 
        //console.log(m)
        //console.log(message)
    });

	$("#sendBtn").click(function () {
      //  message.tag = messageTxt
        count += 1
        var message = {
            image: "busybox",
            tag: "v"+String(count) 
        }        

        m = JSON.stringify(message)
        console.log(m)
        w.Emit("sync", m);
	});

})


function appendMessage(messageDiv) {
    var theDiv = messages[0];
    var doScroll = theDiv.scrollTop == theDiv.scrollHeight - theDiv.clientHeight;
    messageDiv.appendTo(messages);
    if (doScroll) {
        theDiv.scrollTop = theDiv.scrollHeight - theDiv.clientHeight;
    }
}
